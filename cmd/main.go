// File: cmd/main.go
package main

import (
    "fmt"
    "os"
 "github.com/cooper/neuramesh/internal/model"

    "github.com/cooper/neuramesh/internal/wg"
    "github.com/cooper/neuramesh/internal/peer"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: neuramesh <command>")
        return
    }

    switch os.Args[1] {
    case "status":
        fmt.Println("✅ NeuraMesh running. 0 peers connected.")

    case "init":
        priv, pub, err := wg.GenerateKeys()
        if err != nil {
            fmt.Println("❌ Failed to generate keys:", err)
            return
        }
        err = wg.SaveToFile(priv, pub, "configs/self.conf")
        if err != nil {
            fmt.Println("❌ Failed to save config:", err)
            return
        }
        fmt.Println("✅ Keys generated and saved to configs/self.conf")

    case "add-peer":
        if len(os.Args) < 5 {
            fmt.Println("Usage: neuramesh add-peer <name> <pubkey> <ip>")
            return
        }
        p :=model.Peer{
            Name:      os.Args[2],
            PublicKey: os.Args[3],
            IP:        os.Args[4],
        }
        err := peer.AddPeer(p, "configs/peers.json")
        if err != nil {
            fmt.Println("❌ Failed to add peer:", err)
            return
        }
        fmt.Println("✅ Peer added:", p.Name)
case "edit-peer":
    if len(os.Args) < 6 {
        fmt.Println("Usage: neuramesh edit-peer <existing-name> <new-name> <new-pubkey> <new-ip>")
        return
    }

    updated := model.Peer{
        Name:      os.Args[3],
        PublicKey: os.Args[4],
        IP:        os.Args[5],
    }

    err := peer.EditPeer(os.Args[2], updated, "configs/peers.json")
    if err != nil {
        fmt.Println("❌ Failed to edit peer:", err)
        return
    }
    fmt.Println("✅ Peer updated:", updated.Name)

case "monitor":
    peer.MonitorPeers("configs/peers.json")

    case "connect":
        if len(os.Args) < 3 {
            fmt.Println("Usage: neuramesh connect <self-ip>")
            return
        }
        err := wg.Connect(os.Args[2])
        if err != nil {
            fmt.Println("❌ Connect failed:", err)
            return
        }

    default:
        fmt.Println("❌ Unknown command:", os.Args[1])
    }
}

