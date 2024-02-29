package libp2p

import (
	"context"
	"fmt"
	"log"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
)

func AdvertiseKeyValue(ctx context.Context, key string) error {
	log.Printf("Advertising key: %s. Parsing it to CID first...", key)
	keyCID, err := parseStringIntoCID(key)
	if err != nil {
		return fmt.Errorf("couldn't parse key: %w", err)
	}

	log.Printf("CID to be advertised based on key: %s", keyCID)

	err = idht.Provide(ctx, keyCID, true)
	if err != nil {
		return fmt.Errorf("couldn't advertise key: %w", err)
	}

	log.Printf("Key advertised successfully: %s", key)

	return nil
}

// GetProvidersForKey returns the providers for a given key
func GetProvidersForKey(ctx context.Context, key string) ([]peer.AddrInfo, error) {
	log.Printf("Searching for providers for key: %s. Parsing it to CID first...", key)
	keyCID, err := parseStringIntoCID(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse key: %w", err)
	}

	log.Printf("CID to be searched based on key: %s", keyCID)

	peers, err := idht.FindProviders(ctx, keyCID)
	if err != nil {
		return nil, fmt.Errorf("error looking for providers: %w", err)
	}

	log.Printf("Found %d providers for key: %s\n Providers: %v", len(peers), key, peers)

	return peers, nil
}

// parseStringIntoCID "generates" a CID based on a given string. In this case,
// the string is the data.
func parseStringIntoCID(s string) (cid.Cid, error) {
	if s == "" {
		return cid.Cid{}, nil
	}

	_, cidForm, err := cid.CidFromBytes([]byte(s))
	if err != nil {
		return cid.Cid{}, fmt.Errorf("couldn't convert bytes to CID: %w", err)
	}

	return cidForm, nil
}
