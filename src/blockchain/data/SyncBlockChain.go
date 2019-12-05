package data

import "sync"

type SyncBlockChain struct {
	bc  BlockChain
	mux sync.Mutex
}

func (blockchain *SyncBlockChain) GetParentBlock(block Block) Block {
	blockchain.mux.Lock()

	b, _ := blockchain.bc.GetParentBlock(block)
	blockchain.mux.Unlock()
	return b
}

func (blockchain *SyncBlockChain) GetChainHeight() int32 {
	blockchain.mux.Lock()
	height := blockchain.bc.Length
	blockchain.mux.Unlock()
	return height
}

func (blockchain *SyncBlockChain) Show() string {
	blockchain.mux.Lock()
	chain := blockchain.bc.Show()
	blockchain.mux.Unlock()
	return chain
}

func (blockchain *SyncBlockChain) Insert(block Block) error {
	blockchain.mux.Lock()
	err := blockchain.bc.Insert(block)
	blockchain.mux.Unlock()
	return err
}

func (blockchain *SyncBlockChain) GetLatestBlocks(height int32) ([]Block, error) {
	blockchain.mux.Lock()
	blocks, err := blockchain.bc.GetLatestBlocks(height)
	blockchain.mux.Unlock()
	return blocks, err
}

func (c *SyncBlockChain) Get(height int32) []Block {
	c.mux.Lock()
	blocks := c.bc.Chain[height]
	c.mux.Unlock()
	return blocks
}

func (blockchain *SyncBlockChain) EncodeToJSON() (string, error) {
	blockchain.mux.Lock()
	json, err := blockchain.bc.EncodeToJSON()
	blockchain.mux.Unlock()
	return json, err
}

func (blockchain *SyncBlockChain) DecodeFromJSON(data string) error {
	blockchain.mux.Lock()
	_, err := blockchain.bc.DecodeFromJSON(data)
	blockchain.mux.Unlock()
	return err
}
