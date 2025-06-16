# 📘 NeuraMesh Usage Guide

Welcome to NeuraMesh — a self-healing VPN mesh powered by Go and WireGuard.

This guide shows how to install, configure, and run NeuraMesh securely.

---

## 🔧 Setup

1. **Build the CLI**

   ```bash
   go build -o neuramesh cmd/main.go
   chmod +x neuramesh.sh
   ```

2. **Initialize Your Node**

   ```bash
   ./neuramesh init
   ```

   - Creates `configs/self.conf` with your private key
   - Prints your public key (used by other peers)

3. **Add or Edit Peers**

   Add peer:

   ```bash
   ./neuramesh add-peer bob <public_key> 10.0.0.2
   ```

   Edit peer:

   ```bash
   ./neuramesh edit-peer bob bob2 <new_key> 10.0.0.3
   ```

   Reset peer list:

   ```bash
   echo "[]" > configs/peers.json
   ```

4. **Connect the Mesh**

   ```bash
   sudo ./neuramesh connect 10.0.0.1
   ```

5. **Monitor Peer Health**

   ```bash
   sudo ./neuramesh monitor
   ```

   - Removes dead peers after 3 fails
   - Re-adds recovered peers
   - Score-based selection (default: ≥ 80)

6. **Use the Dashboard**

   ```bash
   ./neuramesh.sh
   ```

   Options:
   - 1: Connect to VPN
   - 2: Monitor peers
   - 3: Add peer
   - 4: Show peers
   - 5: Edit peer
   - 6: Exit

7. **Test Multi-Node Setup**

   - Set up 2+ machines or VMs
   - Exchange public keys + IPs
   - Run `connect` + `monitor` on both
   - Use `ping 10.0.0.X` to confirm VPN is working

---

## 📁 Requirements

- Go 1.21+
- WireGuard CLI (`wg`)
- Linux/macOS/WSL with `ip`, `ping`
