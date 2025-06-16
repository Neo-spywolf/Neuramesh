package routing
import "github.com/cooper/neuramesh/internal/model"

type ScoreBasedRouting struct {
    MinScore int
}

func (s ScoreBasedRouting) Select(peers []*model.PeerHealth) []*model.Peer {
    var selected []*model.Peer
    for _, ph := range peers {
        if ph.Score >= s.MinScore {
            selected = append(selected, &ph.Peer)
        }
    }
    return selected
}
