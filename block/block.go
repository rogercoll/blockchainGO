package block

import (
    "encoding/base64"
    "crypto/sha256"
    "strings"
    "time"
    "strconv")

type Block struct{
	timestamp int64
	transactions []string
	previousHash string
	hash string
	nonce int
}

func NewBlock( transactions []string, previousHash string) *Block {
    b := &Block{time.Now().Unix(), transactions, previousHash, "", 0}
    hash := b.CalculateHash()
    b.hash = hash
    return b
}

func NewGenesisBlock() *Block {
    b := &Block{time.Now().Unix(), []string{"This is the Genesis Block"}, "", "", 0}
    hash := b.CalculateHash()
    b.hash = hash
    return b
}

func (b *Block) CalculateHash() string {
	s := b.transactions[:]
	s = append(s,strconv.FormatInt(b.timestamp, 10),b.previousHash,string(b.nonce))
	stringByte := strings.Join(s, " ")
	bv := []byte(stringByte) 
	hasher := sha256.New()
    hasher.Write(bv)
    return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

