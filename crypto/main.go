package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	PrevHash []byte
	Data     []byte
}

func (b *Block) GetHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash := hash[:]
	fmt.Println(hash)
}

func main() {

}
