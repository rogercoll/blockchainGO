package blockchain

import(
"github.com/rogercoll/blockchainGo/block")	


type Blockchain struct {
	blocks []*block.Block
}

func NewBlockchain() *Blockchain{
	return &Blockchain{[]*block.Block{block.NewGenesisBlock()}}
}


func (bc *Blockchain) AddBlock(){

}
