package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func GenesisBlock(data string) *Block {
	var genesisBlock Block
	t := time.Now()

	genesisBlock.Index = 0
	genesisBlock.Timestamp = t.String()
	genesisBlock.Data = data
	genesisBlock.Hash = ComputeHash(&genesisBlock)
	genesisBlock.PrevHash = ""
	return &genesisBlock
}

func ComputeHash(block *Block) string {
	concatBlockAttrs := string(block.Index) + block.Timestamp + block.Data + block.PrevHash

	blockBytes := sha256.New()
	blockBytes.Write([]byte(concatBlockAttrs))

	blockHash := blockBytes.Sum(nil)

	return hex.EncodeToString(blockHash)
}
