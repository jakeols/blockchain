package data

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Header Header
	Value  string
}

type Header struct {
	Height     int32
	Timestamp  int64
	Hash       string
	ParentHash string
	Size       int32
	Nonce      int
}

// initializes a new block, uses hash of JSON value as header hash
func (b *Block) Initial(height int32, parentHash string, value string) {

	b.Value = value

	b.Header = Header{
		Height:     height,
		Timestamp:  time.Now().UnixNano(),
		ParentHash: parentHash,
		Size:       32,
		Nonce:      FindNonce(parentHash, value),
	}

	// set hash
	hash := sha256.New()
	JSONString, _ := b.EncodeToJSON()
	hash.Write([]byte(JSONString))
	md := hash.Sum(nil)
	b.Header.Hash = hex.EncodeToString(md)

}

func FindNonce(parentHash string, value string) int {
	var nonceFound bool = false
	//var hashString string
	var counter int = 0 // maybe make more random in future
	for nonceFound == false {

		hash := sha256.New()
		hash.Write([]byte(parentHash + strconv.Itoa(counter) + value))
		sha256Bytes := sha256.Sum256(hash.Sum(nil))

		s := fmt.Sprintf("%08b", sha256Bytes[:])
		n := strings.Trim(s, "[]")
		t := strings.Replace(n, " ", "", -1)
		// determine if it starts with 10 0's
		if strings.HasPrefix(t, strings.Repeat("0", 10)) {
			nonceFound = true
			break
		}
		counter++

		hash.Reset()

	}
	fmt.Println("nonce found")
	return counter
}

func CheckNonce(nonce int, parentHash string, value string) bool {
	hash := sha256.New()
	hash.Write([]byte(parentHash + strconv.Itoa(nonce) + value))
	sha256Bytes := sha256.Sum256(hash.Sum(nil))

	s := fmt.Sprintf("%08b", sha256Bytes[:])
	n := strings.Trim(s, "[]")
	t := strings.Replace(n, " ", "", -1)

	if strings.HasPrefix(t, strings.Repeat("0", 10)) {
		return true
	}
	return false

}

// decodes a block from JSON
func (b Block) DecodeFromJSON(data string) (Block, error) {
	var NewBlock Block
	err := json.Unmarshal([]byte(data), &NewBlock)
	return NewBlock, err
}

// encodes a block to JSON
func (b Block) EncodeToJSON() (string, error) {
	var JSONData []byte
	JSONData, err := json.Marshal(b)
	return string(JSONData), err
}
