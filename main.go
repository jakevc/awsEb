package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// respond with data
func respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application-json")
	json.NewEncoder(w).Encode(data)
}

// message formats data for response 
func message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func main() {
	// define port and router 
	PORT := ":8080"
	r := mux.NewRouter()

	// handle all routes 
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := message(true, "The test app is live!")
		respond(w, resp)
	})

	// log and serve
	log.Printf("Listening on port: %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}