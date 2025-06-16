# 🏗️ NeuraMesh Architecture

---

## 📁 Project Structure

```
neuramesh/
├── cmd/           # Entry point
├── configs/       # self.conf, peers.json
├── internal/
│   ├── model/     # Peer, PeerHealth structs
│   ├── peer/      # Registry, monitor logic
│   ├── routing/   # Score-based strategy
│   └── wg/        # WireGuard connection logic
├── neuramesh.sh   # Dashboard script
```

---

## 🔁 Connect Flow

1. Reads `self.conf` (private key)
2. Loads peers from `peers.json`
3. Applies `ScoreBasedRouting` strategy
4. Builds `wg0` tunnel using:
   - Private key
   - Allowed IPs + public keys of best peers

---

## 📡 Monitor Flow

1. Pings all peers every 5s
2. If 3 fails → removes from wg0
3. If ping returns → re-adds to wg0
4. Score adjusts dynamically

---

## 🔌 Routing Strategy Interface

```go
type RoutingStrategy interface {
    Select(peers []*PeerHealth) []*Peer
}
```

Plug-and-play logic for:
- Static rules
- Score threshold
- ML/AI models

---

## 🤖 AI-Ready Design

You can add:
- Python gRPC/REST model for scoring
- Logs of uptime/latency
- Custom routing strategy

---

## 🔄 Self-Healing Table

| Condition        | Action         |
|------------------|----------------|
| ping fails x3     | remove peer    |
| ping returns      | re-add peer    |
| score ≥ threshold | include in wg0 |

---

## 📡 Tech Stack

- Go 1.21
- WireGuard CLI
- Linux/macOS/WSL
