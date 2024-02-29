package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/PedrobyJoao/libp2p-test-network/libp2p"
	"github.com/PedrobyJoao/libp2p-test-network/utils"
)

func advertiseKeyValueDHT(w http.ResponseWriter, r *http.Request) {
	var key string

	err := json.NewDecoder(r.Body).Decode(&key)
	if err != nil {
		handleErrorResponse(
			w,
			fmt.Errorf("Error decoding params: %s", err),
		)
		return
	}

	err = libp2p.AdvertiseKeyValue(r.Context(), key)
	if err != nil {
		handleErrorResponse(
			w,
			fmt.Errorf("Error advertising key: %s", err),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getProvidersDHT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["key"]

	peers, err := libp2p.GetProvidersForKey(r.Context(), key)
	if err != nil {
		handleErrorResponse(
			w,
			fmt.Errorf("Error getting providers: %s", err),
		)
		return
	}

	json.NewEncoder(w).Encode(peers)
}

func getPeersFromPeerStore(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(libp2p.GetPeersFromPeerStore())
}

func handleErrorResponse(w http.ResponseWriter, err error) {
	log.Printf("%v", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/peer_store/peers", getPeersFromPeerStore).Methods("GET")
	router.HandleFunc("/dht/get_providers/{key}", getProvidersDHT).Methods("GET")
	router.HandleFunc("/dht/advertise", advertiseKeyValueDHT).Methods("POST")

	port, err := utils.FindFreePort()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("API Listening on port %d", port)

	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(port), router))
}
