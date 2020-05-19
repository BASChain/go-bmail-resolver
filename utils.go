package resolver

import (
	"encoding/hex"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type Hash [32]byte

func (hash Hash)String() string{
	return "0x"+hex.EncodeToString(hash[:])
}

func GetHash(key string) Hash{
	hash := solsha3.SoliditySHA3(solsha3.String(key))
	var ret [32]byte
	for i:=0;i<32;i++ {
		ret[i] = hash[i]
	}
	return ret
}

var RetryRule = map[int]int{
	1:1,
	2:2,
	3:3,
}

func Split(buffer []byte, s byte) []string {
	var recovered []string
	start  :=  0
	for i,  e:=range buffer  {
		var temp []byte
		if e  == s {
			temp = buffer[start:i]
			recovered = append(recovered, string(temp))
			start  = i+1
		}
	}
	recovered = append(recovered, string(buffer[start:]))
	return recovered
}