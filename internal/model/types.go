package model

type Peer struct {
    Name      string `json:"name"`
    PublicKey string `json:"public_key"`
    IP        string `json:"ip"`
}

type PeerHealth struct {
    Peer   Peer
    Status string
    Fails  int
    Score  int
}
