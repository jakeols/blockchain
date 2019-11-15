package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"./uri"
	"./uri/handlers"
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
		//hashString = hex.EncodeToString(hash.Sum(nil))
		sha256Bytes := sha256.Sum256(hash.Sum(nil))

		//fmt.Println("SHA256 String is ", hex.EncodeToString(sha256Bytes[:]))

		s := fmt.Sprintf("%08b", sha256Bytes[:])
		n := strings.Trim(s, "[\t]")
		fmt.Println(n)

		//fmt.Println(bytes.Equal([]byte(hashString), []byte("0000000000")))

		// determine if it starts with 10 0's
		if strings.HasPrefix(n, strings.Repeat("0", 8)) {
			nonceFound = true
			break
		}
		counter++

		hash.Reset()

	}
	return counter
}

func CheckNonce(nonce int, parentHash string, value string) bool {
	hash := sha256.New()
	hash.Write([]byte(parentHash))
	hash.Write([]byte(strconv.Itoa(nonce)))
	hash.Write([]byte(value))

	hashString := hex.EncodeToString(hash.Sum(nil))

	if strings.HasPrefix(hashString, strings.Repeat("0", 2)) {
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

type BlockChain struct {
	Chain  map[int32][]Block
	Length int32
}

func (c BlockChain) Get(height int32) []Block {
	return c.Chain[height]
}

func (c *BlockChain) Insert(block Block) error {

	if c.Length == 0 {
		newChain := make(map[int32][]Block)
		c.Chain = newChain
		c.Length = 0
	}

	// check nonce here
	// need nonce, parentHash, and value
	if CheckNonce(block.Header.Nonce, block.Header.ParentHash, block.Value) == true || c.Length == 0 {
		height := block.Header.Height
		LastBlock := c.Chain[height]

		for _, v := range LastBlock {
			if v.Header.Hash == block.Header.Hash {
				return errors.New("cannot insert block.")
			}
		}

		location := c.Chain[height]

		c.Chain[height] = append(location, block)

		c.Length = int32(len(c.Chain))
		return nil

	}
	// don't accept the block
	return errors.New("nonce invalid")

}

func (c BlockChain) EncodeToJSON() (string, error) {

	blockchain := make([]Block, 0)

	for index := 0; index < int(c.Length); index++ {

		block := c.Chain[int32(index)][0]
		blockchain = append(blockchain, block)
	}

	var JSONData []byte
	JSONData, err := json.Marshal(blockchain)
	return string(JSONData), err
}

func (c BlockChain) DecodeFromJSON(data string) (BlockChain, error) {

	chain := new([]Block)

	err := json.Unmarshal([]byte(data), &chain)

	for _, block := range *chain {
		c.Insert(block)
	}

	return c, err
}

func main() {
	blockchain := new(BlockChain)

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
