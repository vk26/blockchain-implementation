package main

import "fmt"

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) newGenesis() {
	block := NewBlock("Genesis block", []byte{})
	bc.blocks = append(bc.blocks, block)
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.newGenesis()
	return bc
}

func (bc *Blockchain) PrintLastBlock() {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	fmt.Printf("Hash: %x\n", lastBlock.Hash)
	fmt.Println("Data:", string(lastBlock.Data))
	fmt.Printf("PrevHash: %x\n", lastBlock.PrevHash)
	fmt.Println()
}
