package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Resp struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test-welcome", Welcome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func Welcome(w http.ResponseWriter, r *http.Request) {
	var resp = Resp{Code: "200", Data: "Ok"}
	json.NewEncoder(w).Encode(resp)
}
