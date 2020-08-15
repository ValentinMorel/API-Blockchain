package blockchain

import (
	"API-example/block"
	"time"
)

type Blockchain struct {
	Blocks []*block.Block
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*block.Block{}}
}

func (b *Blockchain) AddBlock(block *block.Block) {
	b.Blocks = append(b.Blocks, block)
}

func (b *Blockchain) GenerateNewBlock(data string) (*block.Block, error) {

	var newBlock block.Block
	prevBlock := b.Blocks[len(b.Blocks)-1]
	timestamp := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = timestamp.String()
	newBlock.Data = data
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = block.ComputeHash(&newBlock)

	return &newBlock, nil
}
