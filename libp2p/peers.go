package libp2p

import (
	"log"

	"github.com/libp2p/go-libp2p/core/peer"
)

func GetPeersFromPeerStore() peer.IDSlice {
	log.Printf(
		"printing connected peers just in case it's different from peerstore\n Connected: %v\n",
		h.Network().Peers(),
	)
	return h.Peerstore().Peers()
}
