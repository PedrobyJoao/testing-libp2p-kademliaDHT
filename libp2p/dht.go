package libp2p

import (
	"context"
	"fmt"
	"log"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	mh "github.com/multiformats/go-multihash"
)

func main() {
	// Your random string
	randomString := "your-random-string"

	// Create a multihash from the random string
	hash, _ := mh.Sum([]byte(randomString), mh.SHA2_256, -1)

	// Create a CID using the multihash
	c := cid.NewCidV1(cid.Raw, hash)

	// Print out the CID
	fmt.Println(c)
}

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

	// Create a multihash from the random string
	hash, err := mh.Sum([]byte(s), mh.SHA2_256, -1)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("couldn't create multihash from string: %w", err)
	}

	// Create a CID using the multihash
	c := cid.NewCidV1(cid.Raw, hash)

	return c, nil
}
