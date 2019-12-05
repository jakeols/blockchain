package data

import (
	"encoding/json"
	"errors"
)

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

		if block.Header.Height > c.Length {
			c.Length = block.Header.Height
		}
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

func (c BlockChain) GetLatestBlocks(height int32) ([]Block, error) {

	return c.Chain[height], nil

}

func (c BlockChain) GetParentBlock(b Block) (Block, error) {
	parentHash := b.Header.ParentHash
	parentHeight := b.Header.Height - 1
	blocks := c.Get(parentHeight)
	if blocks == nil {
		return Block{}, errors.New("can't find parent blocks ")
	}

	for _, b := range blocks {
		if b.Header.Hash == parentHash {
			return b, nil
		}
	}

	return Block{}, errors.New("genesis block")
}
