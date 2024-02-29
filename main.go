package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/PedrobyJoao/libp2p-test-network/libp2p"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port      = flag.Int("port", 0, "port number to listen on libp2p host")
	bootstrap = flag.Bool("bootstrap", true, "connect to bootstrap nodes")
)

func main() {
	flag.Parse()

	var bootstrapPeers []string
	if *bootstrap {
		log.Println("Selected to connect to bootstrap nodes...")
		bootstrapPeers = []string{
			"/ip4/127.0.0.1/tcp/8080/p2p/QmdWGyw79h4KVK6eiwzvqgaTYV1mZTco9epPHhfeJwPBwz",
		}
	}

	host, err := libp2p.NewHost(*port, bootstrapPeers...)
	if err != nil {
		panic(err)
	}

	// start a web server to expose metrics
	// serveMetricsToPrometheus()

	fmt.Println(host.Peerstore().Peers())

	// wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")

	// shut the node down
	if err := host.Close(); err != nil {
		panic(err)
	}
}

func serveMetricsToPrometheus() {
	go func() {
		http.Handle("/debug/metrics/prometheus", promhttp.Handler())
		log.Fatal(http.ListenAndServe("localhost:0", nil))
	}()
}
