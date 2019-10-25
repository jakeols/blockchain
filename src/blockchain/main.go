package main

type Block struct {
	
	Header struct {
		height int32
		timestamp int64
		hash string
	}

	parent_hash string
	size int32
	value string
}

func (b Block ) Initial() {

}

func (b Block) DecodeFromJSON(json string) Block {

}

func (b Block) EncodeToJSON() string {
	return "yo";
}

type BlockChain struct {
	chain map[int32][]Block
	length int32
}

func (c BlockChain) Get(height string) []Block {

}

func (c BlockChain) Insert(block Block) {

}

func (c BlockChain) EncodeToJSON() string {
	return "json";
}

func (c BlockChain) DecodeFromJSON(json string) {

}

func main(){

}
