package routing

import "github.com/cooper/neuramesh/internal/model"

type RoutingStrategy interface {
    Select(peers []*model.PeerHealth) []*model.Peer
}
