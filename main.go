package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", ReplyEndpoint).Methods("GET")
	return router
}

func main() {
	router := Router()
	log.Fatal(http.ListenAndServe(":8888", router))
}
