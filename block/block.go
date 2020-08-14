package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Hash      string
	PrevHash  string
}

func GenerateBlock(prevBlock Block) (*Block, error) {

	var newBlock *Block
	timestamp := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = timestamp.String()
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = ComputeHash(newBlock)

	return newBlock, nil
}

func ComputeHash(block *Block) string {
	concatBlockAttrs := string(block.Index) + block.Timestamp + block.PrevHash

	blockBytes := sha256.New()
	blockBytes.Write([]byte(concatBlockAttrs))

	blockHash := blockBytes.Sum(nil)

	return hex.EncodeToString(blockHash)
}
