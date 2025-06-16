package peer

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/cooper/neuramesh/internal/model"
)

func AddPeer(p model.Peer, filename string) error {
    peers, err := LoadPeers(filename)
    if err != nil {
        return err
    }

    for _, existing := range peers {
        if existing.Name == p.Name {
            return fmt.Errorf("peer with name '%s' already exists", p.Name)
        }
    }

    peers = append(peers, p)
    return SavePeers(peers, filename)
}

func LoadPeers(filename string) ([]model.Peer, error) {
    var peers []model.Peer
    data, err := os.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return peers, nil
        }
        return nil, err
    }

    err = json.Unmarshal(data, &peers)
    return peers, err
}

func SavePeers(peers []model.Peer, filename string) error {
    data, err := json.MarshalIndent(peers, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0600)
}
func EditPeer(name string, updated model.Peer, filename string) error {
    peers, err := LoadPeers(filename)
    if err != nil {
        return err
    }

    updatedList := []model.Peer{}
    found := false

    for _, p := range peers {
        if p.Name == name {
            updatedList = append(updatedList, updated)
            found = true
        } else {
            updatedList = append(updatedList, p)
        }
    }

    if !found {
        return fmt.Errorf("peer '%s' not found", name)
    }

    return SavePeers(updatedList, filename)
}

