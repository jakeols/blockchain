package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./uri"
	"./uri/handlers"
)

func main() {
	blockchain := new(BlockChain)
	blockchain.GenesisBlock()

	InitialBlock := new(Block)
	InitialBlock.Initial(0, "none", "none")
	blockchain.Insert(*InitialBlock)

	SecondBlock := new(Block)
	SecondBlock.Initial(1, blockchain.Chain[blockchain.Length-1][0].Header.Hash, "tx vals")
	blockchain.Insert(*SecondBlock)

	ThirdBlock := new(Block)
	ThirdBlock.Initial(2, blockchain.Chain[blockchain.Length-1][0].Header.Hash, "more tx vals")
	blockchain.Insert(*ThirdBlock)

	JSONBlockchain, _ := blockchain.EncodeToJSON()
	fmt.Println(JSONBlockchain)

	// communication

	router := uri.NewRouter()

	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "6689"
	}

	handlers.InitSelfAddress(port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
