package main

import (
	"log"
	"net/http"
	"os"

	"./data/"
)

var CurrentPeerList = data.PeerList{}
var CurrentBlockChain = data.SyncBlockChain{}

func main() {

	// create and start server
	router := NewRouter()
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "6689"
	}
	InitSelfAddress(port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
