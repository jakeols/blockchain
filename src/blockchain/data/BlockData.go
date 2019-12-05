package data

import "encoding/json"

type BlockData struct {
	SenderId      int32
	SenderAddress int32
	BlockJSON     string
}

// decodes block data from JSON
func (b BlockData) DecodeFromJSON(data string) (BlockData, error) {
	var NewBlock BlockData
	err := json.Unmarshal([]byte(data), &NewBlock)
	return NewBlock, err
}

// encodes a block to JSON
func (b BlockData) EncodeToJSON() (string, error) {
	var JSONData []byte
	JSONData, err := json.Marshal(b)
	return string(JSONData), err
}
