package examples

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Item ... dummy item type
type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// StartHTTPJSONResponseServer ... starts http server
func StartHTTPJSONResponseServer() {
	http.HandleFunc("/endpoint", EndpointRequestHandler)
	fmt.Print("server started on port 8000")
	http.ListenAndServe(":8000", nil)
}

// EndpointRequestHandler ... handles endpoint
func EndpointRequestHandler(w http.ResponseWriter, r *http.Request) {
	item := Item{
		Name:  "some-item",
		Price: 123,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
