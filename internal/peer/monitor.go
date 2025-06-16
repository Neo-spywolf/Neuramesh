package peer

import (
    "encoding/json"
    "fmt"
    "os"
    "os/exec"
    "time"

    "github.com/cooper/neuramesh/internal/model"
    "github.com/cooper/neuramesh/internal/routing"
)

func ping(ip string) bool {
    cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
    return cmd.Run() == nil
}

func MonitorPeers(filename string) {
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("❌ Failed to read peers:", err)
        os.Exit(1)
    }

    var peers []model.Peer
    err = json.Unmarshal(data, &peers)
    if err != nil {
        fmt.Println("❌ Failed to parse peers.json:", err)
        os.Exit(1)
    }

    health := make(map[string]*model.PeerHealth)
    revived := make(map[string]bool)
    for _, p := range peers {
        health[p.Name] = &model.PeerHealth{Peer: p, Status: "Checking", Fails: 0, Score: 60}
    }

    router := routing.ScoreBasedRouting{MinScore: 80}

    fmt.Println("📡 Starting peer monitor... (Ctrl+C to stop)")
    for {
        fmt.Println("----------- Peer Health -----------")
        var all []*model.PeerHealth
        for name, ph := range health {
            alive := ping(ph.Peer.IP)
            if alive {
                if ph.Status == "❌ Dead" && !revived[ph.Peer.Name] {
                    fmt.Printf("🧠 Re-adding peer %s to wg0\n", ph.Peer.Name)
                    _ = exec.Command("wg", "set", "wg0", "peer", ph.Peer.PublicKey, "allowed-ips", ph.Peer.IP+"/32", "endpoint", ph.Peer.IP+":51820").Run()
                    revived[ph.Peer.Name] = true
                }
                ph.Status = "✅ Healthy"
                ph.Fails = 0
                if ph.Score < 100 {
                    ph.Score += 10
                }
            } else {
                ph.Fails++
                if ph.Fails >= 3 && ph.Status != "❌ Dead" {
                    ph.Status = "❌ Dead"
                    ph.Score = 0
                    fmt.Printf("⚠️ Removing peer %s from wg0\n", ph.Peer.Name)
                    _ = exec.Command("wg", "set", "wg0", "peer", ph.Peer.PublicKey, "remove").Run()
                    revived[ph.Peer.Name] = false
                } else if ph.Status != "❌ Dead" {
                    ph.Status = "⚠️ Unreachable"
                    ph.Score -= 15
                    if ph.Score < 0 {
                        ph.Score = 0
                    }
                }
            }
            all = append(all, ph)
            fmt.Printf("%s [%s] → %s | Score: %d\n", name, ph.Peer.IP, ph.Status, ph.Score)
        }

        fmt.Println("🧭 Selected Peers Based on Strategy:")
        for _, p := range router.Select(all) {
            fmt.Printf("➡️  %s [%s]\n", p.Name, p.IP)
        }
        fmt.Println("-----------------------------------")
        time.Sleep(5 * time.Second)
    }
}
