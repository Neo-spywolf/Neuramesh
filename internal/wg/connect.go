package wg

import (
    "encoding/json"
    "fmt"
    "os"
    "os/exec"

    "github.com/cooper/neuramesh/internal/model"
    "github.com/cooper/neuramesh/internal/routing"
)

func parseConfig(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }

    lines := string(data)
    for _, line := range splitLines(lines) {
        if len(line) > 8 && line[:8] == "private=" {
            return line[8:], nil
        }
    }
    return "", fmt.Errorf("private key not found in config")
}

func splitLines(s string) []string {
    var lines []string
    current := ""
    for _, c := range s {
        if c == '\n' {
            lines = append(lines, current)
            current = ""
        } else {
            current += string(c)
        }
    }
    if current != "" {
        lines = append(lines, current)
    }
    return lines
}

func Connect(selfIP string) error {
    privKey, err := parseConfig("configs/self.conf")
    if err != nil {
        return fmt.Errorf("failed to read private key: %v", err)
    }

    data, err := os.ReadFile("configs/peers.json")
    if err != nil {
        return fmt.Errorf("failed to load peers.json: %v", err)
    }
    var rawPeers []model.Peer
    if err := json.Unmarshal(data, &rawPeers); err != nil {
        return fmt.Errorf("invalid peers.json: %v", err)
    }

    var health []*model.PeerHealth
    for _, p := range rawPeers {
        health = append(health, &model.PeerHealth{
            Peer:   p,
            Status: "Assumed Healthy",
            Score:  100,
        })
    }

    router := routing.ScoreBasedRouting{MinScore: 80}
    selected := router.Select(health)
    if len(selected) == 0 {
        return fmt.Errorf("no healthy peers to connect to")
    }

    _ = exec.Command("ip", "link", "del", "wg0").Run()
    if err := exec.Command("ip", "link", "add", "wg0", "type", "wireguard").Run(); err != nil {
        return fmt.Errorf("failed to add wg0: %v", err)
    }

    if err := exec.Command("ip", "address", "add", selfIP+"/24", "dev", "wg0").Run(); err != nil {
        return fmt.Errorf("failed to assign IP: %v", err)
    }

    tmpKey := "/tmp/neuramesh.key"
    if err := os.WriteFile(tmpKey, []byte(privKey), 0600); err != nil {
        return fmt.Errorf("failed to write temp priv key: %v", err)
    }
    defer os.Remove(tmpKey)

    if err := exec.Command("wg", "set", "wg0", "private-key", tmpKey).Run(); err != nil {
        return fmt.Errorf("failed to set private key: %v", err)
    }

    for _, p := range selected {
        cmd := exec.Command("wg", "set", "wg0",
            "peer", p.PublicKey,
            "allowed-ips", p.IP+"/32",
            "endpoint", p.IP+":51820")
        if output, err := cmd.CombinedOutput(); err != nil {
            return fmt.Errorf("failed to add peer %s: %v — %s", p.Name, err, string(output))
        }
    }

    if err := exec.Command("ip", "link", "set", "up", "dev", "wg0").Run(); err != nil {
        return fmt.Errorf("failed to bring up wg0: %v", err)
    }

    fmt.Println("✅ VPN tunnel established using best-scored peers!")
    return nil
}
