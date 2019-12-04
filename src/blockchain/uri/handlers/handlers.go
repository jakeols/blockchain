package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"../../data"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var SELF_ADDRESS string

func InitSelfAddress(port string) {
	SELF_ADDRESS = "http://localhost:" + port
}

func SimpleGet(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("just a simple get"))
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func AnotherGet(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		num := vars["number"]

		w.WriteHeader(http.StatusOK)
		str := "Another get api, with paramerter, number as: " + num
		_, _ = w.Write([]byte(str))
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func SimplePost(w http.ResponseWriter, r *http.Request) {

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

func AskOddOrEven(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		vars := mux.Vars(r)
		number, err := strconv.Atoi(vars["number"]) //getting number from request
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		ques := data.Question{
			Value: number,
		}
		quesJson, _ := ques.ToJson() //creating question json

		//// forming url for post request ////
		uri := SELF_ADDRESS + "/oddoreven"

		//// sending post request and taking resp ////
		resp, err := http.Post(uri, "application/json", bytes.NewBuffer(quesJson))
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		defer resp.Body.Close()

		//reading response of post request
		recvResponse, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		//writing response back to initial get request
		w.Write(recvResponse)

	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func OddOrEven(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		reqBody, err := readRequestBody(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		ques, err := data.QuestionFromJson([]byte(reqBody))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		answer := data.Answer{
			Value:  ques.Value,
			Result: "",
		}
		zeroOroOne := ques.Value % 2
		if zeroOroOne == 0 {
			answer.Result = "Even"
		} else {
			answer.Result = "Odd"
		}
		ansJson, _ := answer.ToJson()

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(ansJson)

	}
	w.WriteHeader(http.StatusMethodNotAllowed)

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
