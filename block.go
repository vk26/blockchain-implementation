package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
	Data      []byte
}

func NewBlock(data string, prevHash []byte) *Block {
	b := &Block{
		Timestamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      []byte(data),
	}
	b.setHash()
	return b
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	header := bytes.Join([][]byte{timestamp, b.Hash, b.Data}, []byte{})
	sha256.Sum256(header)
}
