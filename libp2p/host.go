package libp2p

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/routing"
	"github.com/multiformats/go-multiaddr"
)

var h host.Host
var idht *dht.IpfsDHT

// NewHost creates a new libp2p host
func NewHost(port int, bootstrapPeers ...string) (host.Host, error) {

	ctx := context.Background()

	// Creates a new RSA key pair for this host.
	r := rand.Reader

	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	h, err = libp2p.New(
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port),
			fmt.Sprintf("/ip4/127.0.0.1/udp/%d/quic", port),
		),
		libp2p.Identity(prvKey),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			idht, err = dht.New(ctx, h)
			return idht, err
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new p2p host: %v", err)
	}

	if len(bootstrapPeers) != 0 {
		err = connectToBootstrapNodes(ctx, h, bootstrapPeers)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to bootstrap nodes: %v", err)
		}
	}

	// prints host PeerID and listen addresses
	fmt.Printf("PeerID: %s\n", h.ID())
	fmt.Printf("Listen addresses: %s\n", h.Addrs())

	return h, nil
}

func connectToBootstrapNodes(ctx context.Context, h host.Host, bootstrapPeers []string) error {
	// TODO: dht bootstrap?
	for _, bp := range bootstrapPeers {
		addr, err := multiaddr.NewMultiaddr(bp)
		if err != nil {
			return fmt.Errorf("failed to parse bootstrap peer address to multiaddr: %v", err)
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return fmt.Errorf("failed to parse peer address (%v) to peerInfo: %v", addr, err)
		}

		err = h.Connect(ctx, *peerInfo)
		if err != nil {
			return fmt.Errorf("failed to connect to bootstrap peer: %v", err)
		} else {
			log.Println("connected to bootstrap peer: ", peerInfo.ID)
		}
	}
	return nil
}
