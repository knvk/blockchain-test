package main

import (
    "fmt"
    "time"
    "strconv"
    "encoding/hex"
    "crypto/sha256"
)

type Blockchain []*Block

type Block struct {
    Index     int
    Timestamp int64
    Data    []byte
    Hash    []byte
    PrevHash   []byte
}

func (b *Block) GenerateHash() {
    header := strconv.Itoa(b.Index) + strconv.FormatInt(b.Timestamp,10) + string(b.Data) + string(b.PrevHash)
    hashed := sha256.Sum256([]byte(header))
    b.Hash = hashed[:]
}

func (bc *Blockchain) GenNewBlock(data string) {
    t := time.Now().UnixNano()
    prevBlock := (*bc)[len(*bc)-1]
    block := &Block{prevBlock.Index+1, t, []byte(data), []byte{}, prevBlock.Hash}
    block.GenerateHash()
    *bc = append((*bc), block)
}

func main() {
    genesisBlock := &Block{0, time.Now().UnixNano(), []byte("GENESIS"), []byte{}, []byte{}}
    genesisBlock.GenerateHash()
    chain := Blockchain{genesisBlock}
    testArr := []string{"DATA1", "DATA2", "DATA3"}
    for _, v := range testArr {
        chain.GenNewBlock(v)
    }
    for _, bl := range chain {
        fmt.Printf("id=%d;data=%s;hash=%s;old_hash=%s\n",bl.Index, string(bl.Data), hex.EncodeToString(bl.Hash), hex.EncodeToString(bl.PrevHash))
    }
}
