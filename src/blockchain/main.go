package main

import (
	"log"
	"net/http"
	"os"

	"./data/"
)

var CurrentPeerList = data.PeerList{}
var CurrentBlockChain = data.BlockChain{}

func main() {

	// now test the setup

	// create a random  block for testing
	genesis := new(data.Block)
	genesis.Initial(0, "genesis", "genesis")
	// turn it into json
	jsonGenesis, _ := genesis.EncodeToJSON()
	// create this as blockdata
	testBlockData := new(data.BlockData)
	testBlockData.SenderId = 12
	testBlockData.SenderAddress = 10
	testBlockData.BlockJSON = jsonGenesis
	//	jsonBlockData, _ := testBlockData.EncodeToJSON()

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
