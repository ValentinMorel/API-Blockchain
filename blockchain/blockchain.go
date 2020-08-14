package blockchain

import (
	"API-example/block"
)

type Blockchain struct {
	Blocks []*block.Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block}}
}

func (b *Blockchain) AddBlock(block *block.Block) {
	b.Blocks = append(b.Blocks, *block)
}
