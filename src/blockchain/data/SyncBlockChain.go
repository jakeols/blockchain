package data

import "sync"

type SyncBlockchain struct {
	bc  BlockChain
	mux sync.Mutex
}

func (blockchain *SyncBlockchain) GetParentBlock(block Block) Block {
	blockchain.mux.Lock()

	b, _ := blockchain.bc.GetParentBlock(block)
	blockchain.mux.Unlock()
	return b
}

func (blockchain *SyncBlockchain) Insert(block Block) error {
	blockchain.mux.Lock()
	err := blockchain.bc.Insert(block)
	blockchain.mux.Unlock()
	return err
}

func (blockchain *SyncBlockchain) GetLatestBlocks(height int32) ([]Block, error) {
	blockchain.mux.Lock()
	blocks, err := blockchain.bc.GetLatestBlocks(height)
	blockchain.mux.Unlock()
	return blocks, err
}

func (blockchain *SyncBlockchain) EncodeToJSON() (string, error) {
	blockchain.mux.Lock()
	json, err := blockchain.bc.EncodeToJSON()
	blockchain.mux.Unlock()
	return json, err
}

func (blockchain *SyncBlockchain) DecodeFromJSON(data string) error {
	blockchain.mux.Lock()
	_, err := blockchain.bc.DecodeFromJSON(data)
	blockchain.mux.Unlock()
	return err
}
