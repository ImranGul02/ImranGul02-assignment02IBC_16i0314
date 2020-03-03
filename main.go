package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	hash   []byte
	data   []byte
	p_hash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.data, b.p_hash}, []byte{})
	Hash := sha256.Sum256(info)
	b.hash = Hash[:]
}
func CreateBlock(data string, p_hash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), p_hash}
	block.DeriveHash()
	return block

}

func (chain *BlockChain) Add_Block(data string) {
	p_block := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, p_block.hash)
	chain.blocks = append(chain.blocks, new)
}

func satoshi() *Block {
	return CreateBlock("satoshi", []byte{})
}

func initBC() *BlockChain {
	return &BlockChain{[]*Block{satoshi()}}

}
func main() {
	chain := initBC()
	chain.Add_Block("first")
	chain.Add_Block("second")
	chain.Add_Block("third")

	for _, block := range chain.blocks {
		fmt.Printf("previous hash %x\n", block.p_hash)
		fmt.Printf("data %s\n", block.data)
		fmt.Printf("hash %x\n", block.hash)
	}
}
