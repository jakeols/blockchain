package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	data "./data"
	"github.com/pkg/errors"
)

var SELF_ADDRESS string

func InitSelfAddress(port string) {
	SELF_ADDRESS = "http://localhost:" + port
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
	// receive block
	if r.Method == http.MethodPost {

		requestBody, err := readRequestBody(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		sb := strings.Builder{}
		sb.WriteString("In simple post, received request body is: \n")
		sb.WriteString(requestBody)
		_, _ = w.Write([]byte(sb.String()))
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// register to peer list
func Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	sb := strings.Builder{}
	sb.WriteString(r.RemoteAddr)
	// create a new peer
	newPeer := new(data.Peer)
	newPeer.ID = r.RemoteAddr

	CurrentPeerList = append(CurrentPeerList, *newPeer)
	// add this to my peerlist
	_, _ = w.Write([]byte(sb.String()))

}

func Start(w http.ResponseWriter, r *http.Request) {

}

// returns blockchain
func Upload(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(CurrentBlockChain); err != nil {
		panic(err)
	}

}

func ReturnBlock(w http.ResponseWriter, r *http.Request) {

}
