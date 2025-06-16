\# NeuraMesh

🧠 **NeuraMesh** is a cross-platform, self-healing VPN mesh built with Go and WireGuard.  
It monitors peer health in real time, routes traffic only through the most reliable nodes, and supports intelligent routing strategies — AI-ready by design.

---

## 🌐 Features

- 🔐 Encrypted VPN using WireGuard
- 🧠 Intelligent, pluggable routing (score-based by default)
- 🔄 Auto-removal of dead peers
- ♻️ Auto-recovery of revived peers
- 🧪 Peer scoring and health monitoring
- 🧰 Interactive terminal dashboard (`neuramesh.sh`)
- 🧱 Modular, production-grade Go code

---

## 📦 Quick Start

```bash
# Clone and build
git clone https://github.com/YOUR_USERNAME/neuramesh.git
cd neuramesh
go build -o neuramesh cmd/main.go
chmod +x neuramesh.sh

-----------------------------------------------------------------------------------------------------------------------------------------------

🚀 CLI Commands

| Command                    | Description                  |
| -------------------------- | ---------------------------- |
| `./neuramesh init`         | Generate self key pair       |
| `./neuramesh add-peer`     | Register a new peer          |
| `./neuramesh edit-peer`    | Modify existing peer         |
| `./neuramesh connect <ip>` | Start VPN interface          |
| `./neuramesh monitor`      | Start health monitor         |
| `./neuramesh.sh`           | Launch interactive dashboard |

-----------------------------------------------------------------------------------------------------------------------------------------------

📁 File Structure

configs/           # Contains self.conf, peers.json
cmd/main.go        # CLI entry point
internal/
├── model/         # Shared Peer, PeerHealth types
├── peer/          # Registry + Monitor logic
├── routing/       # Routing strategies
├── wg/            # WireGuard connect logic
neuramesh.sh       # Interactive dashboard

-----------------------------------------------------------------------------------------------------------------------------------------------

🧠 How It Works

1) Each node runs neuramesh init to generate keys

2) Nodes exchange public keys + IPs

3) neuramesh add-peer adds them to the mesh

4) connect brings up wg0 tunnel only to healthy peers

5) monitor scores peers and heals broken tunnels

6) Pluggable strategy (default: score ≥ 80) controls routing
 
-----------------------------------------------------------------------------------------------------------------------------------------------

## 📚 Full Documentation

- [Usage Guide](docs/usage.md)
- [Architecture Overview](docs/architecture.md)

