## Testing content routing of Kademlia DHT (by libp2p)

**IMPORTANT**: this code is just to **test** a libp2p feature based on Kademlia DHT,
it's **not** a full implementation of a p2p network.

`go run main.go --port=8080 --bootstrap=false` to deploy a bootstrap node 

> I know, just noticed that is a lil bit strange to have `--bootstrap=false` for a bootstrap node. The idea was to
connect to bootstrap nodes in case it's `=true`. I won't change it because I already tested
what I need to test.

`go run main.go` to deploy normal peers

**Note 1**: Peers are running **locally**.

**Note 2**: **bootstrap node** must be hardcoded on `main.go`.

**Why I'm testing?** see context on https://gitlab.com/nunet/research/core-platform/-/issues/15 

### To see more about libp2p implementation of Kademlia DHT:

- https://pl-launchpad.io/curriculum/libp2p/dht/
- https://docs.libp2p.io/concepts/discovery-routing/kaddht/
- https://github.com/libp2p/specs/blob/master/kad-dht/README.md
