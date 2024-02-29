package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/PedrobyJoao/libp2p-test-network/libp2p"
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

func handleErrorResponse(w http.ResponseWriter, err error) {
	log.Printf("%v", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dht/get_providers/{key}", getProvidersDHT).Methods("GET")
	router.HandleFunc("/dht/advertise", advertiseKeyValueDHT).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetItem handles GET requests for a specific item
// func GetItem(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for _, item := range items {
// 		if item.ID == params["id"] {
// 			// json.NewEncoder(w).Encode(item) [item must be a ]
// 			return
// 		}
// 	}
