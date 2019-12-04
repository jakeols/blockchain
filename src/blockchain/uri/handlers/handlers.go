package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"

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

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("getting peerlist"))
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
