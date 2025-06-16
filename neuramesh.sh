#!/bin/bash

CONFIG="configs/self.conf"
PEERS="configs/peers.json"

function show_menu() {
  clear
  echo "🚀 NeuraMesh Dashboard"
  echo "------------------------"
  echo "1) Connect to VPN"
  echo "2) Monitor peers"
  echo "3) Add peer"
  echo "4) Show peer config"
  echo "5) Edit peer"
  echo "6) Exit"
  echo ""
  read -p "Choose an option: " choice
  case $choice in
    1) connect ;;
    2) monitor ;;
    3) add_peer ;;
    4) cat "$PEERS"; read -p "Press Enter to return..." ;;
    5) edit_peer ;;
    6) exit 0 ;;
    *) echo "Invalid option"; sleep 1 ;;
  esac
  show_menu
}

function connect() {
  read -p "Enter your self IP (e.g. 10.0.0.1): " selfip
  go build -o neuramesh cmd/main.go
  sudo ./neuramesh connect "$selfip"
  read -p "Press Enter to return..."
}

function monitor() {
  sudo ./neuramesh monitor
}

function add_peer() {
  read -p "Peer name: " name
  read -p "Peer public key: " pubkey
  read -p "Press Enter to return..."
}
function edit_peer() {
  read -p "Existing peer name: " oldname
  read -p "New name: " newname
  read -p "New public key: " pubkey
  read -p "New IP: " ip
  ./neuramesh edit-peer "$oldname" "$newname" "$pubkey" "$ip"
  read -p "Press Enter to return..."
}

show_menu
