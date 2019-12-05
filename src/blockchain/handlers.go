package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	data "./data"
	"github.com/pkg/errors"
)

var SELF_ADDRESS string

var running bool

func InitSelfAddress(port string) {
	SELF_ADDRESS = "http://localhost:" + port
}

// the main stuff here, starts up with tryinc nonces
func StartTryingNonces() {
	running = true
	for {
		startingBlocks := CurrentBlockChain.Get(CurrentBlockChain.Length)
		if len(startingBlocks) > 0 {
			// get the parent block hash
			parentBlock, _ := CurrentBlockChain.GetParentBlock(startingBlocks[len(startingBlocks)-1])
			fmt.Println(parentBlock.Header.Timestamp)
			// create a new block
			newBlock := new(data.Block)
			newBlock.Initial(CurrentBlockChain.Length+1, parentBlock.Header.ParentHash, "value")
			CurrentBlockChain.Insert(*newBlock)

		} else {
			// create a new block
			newBlock := new(data.Block)
			newBlock.Initial(1, "test", "value")
			CurrentBlockChain.Insert(*newBlock)
		}

	}

}
func readRequestBody(r *http.Request) (string, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", errors.New("cannot read request body")
	}
	defer r.Body.Close()
	return string(reqBody), nil
}

func ReceiveBlock(w http.ResponseWriter, r *http.Request) {
	// get new block JSON
	reqBody, err := readRequestBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	newBlockData := new(data.BlockData)
	newBlockData.DecodeFromJSON(reqBody)
	// insert it into our blockchain
	newBlock := new(data.Block)
	newBlock.DecodeFromJSON(newBlockData.BlockJSON)

	error := CurrentBlockChain.Insert(*newBlock)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// register to peer list
func Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	sb := strings.Builder{}
	sb.WriteString(r.RemoteAddr)
	// create a new peer
	newPeer := new(data.Peer)
	newPeer.ID = r.RemoteAddr
	// add this to my peerlist
	CurrentPeerList = append(CurrentPeerList, *newPeer)

	// write it back
	_, _ = w.Write([]byte(sb.String()))

}

func Start(w http.ResponseWriter, r *http.Request) {
	if !running {
		go StartTryingNonces()
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// returns blockchain
func Upload(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(CurrentBlockChain); err != nil {
		panic(err)
	}

}

// gets block at /block/{height}/{hash}
func ReturnBlock(w http.ResponseWriter, r *http.Request) {

}
