package main

import "net/http"

func ReplyEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(Reply("World")))
}
