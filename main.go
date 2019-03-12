package main


import (
	"github.com/rogercoll/blockchainGo/block"
	"github.com/rogercoll/blockchainGo/blockchain"
	"fmt")


func main(){
	b := block.NewBlock([]string{"one ", " twwo", " three "},"fdsfasfdsa")
	bc := blockchain.NewBlockchain()
	fmt.Println(b)
	fmt.Println(bc)
}
