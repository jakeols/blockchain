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
	Nonce      string
}

// initializes a new block, uses hash of JSON value as header hash
func (b *Block) Initial(height int32, parentHash string, value string, nonce string) {

	b.Value = value

	// find the nonce here

	b.Header = Header{
		Height:     height,
		Timestamp:  time.Now().UnixNano(),
		ParentHash: parentHash,
		Size:       32,
		Nonce:      nonce,
	}

	// set hash
	hash := sha256.New()
	JSONString, _ := b.EncodeToJSON()
	hash.Write([]byte(JSONString))
	md := hash.Sum(nil)
	b.Header.Hash = hex.EncodeToString(md)

}

func FindNonce(parentHash string) {
	// should have 10 0's
	var nonceFound bool = false

	var counter int = 1 // we should start this at a random value later
	for nonceFound == false {
		hash := sha256.New()
		var testString string = "12345"
		hash.Write([]byte(strconv.Itoa(counter)))

		hash.Write([]byte(testString))
		// also write block hash here maybe
		md := hash.Sum(nil)
		fmt.Println("New hash generated: ")
		fmt.Println(len(md))

		// determine if it starts with 10 0's
		// if bits.LeadingZeros8(md) == 10 {
		// 	fmt.Println("Nonce found")
		// 	nonceFound = true
		// }
		counter++

	}

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
	InitialBlock.Initial(0, "none", "none", "none")
	blockchain.Insert(*InitialBlock)

	SecondBlock := new(Block)
	SecondBlock.Initial(1, blockchain.Chain[blockchain.Length-1][0].Header.Hash, "tx vals", "nonce")
	blockchain.Insert(*SecondBlock)

	ThirdBlock := new(Block)
	ThirdBlock.Initial(2, blockchain.Chain[blockchain.Length-1][0].Header.Hash, "more tx vals", "nonce")
	blockchain.Insert(*ThirdBlock)

	JSONBlockchain, _ := blockchain.EncodeToJSON()
	fmt.Println(JSONBlockchain)

	FindNonce("10")

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
