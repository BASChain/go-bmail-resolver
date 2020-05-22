package resolver

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Config struct {
	A          []byte
	AAAA       []byte
	MX         []byte
	BlockChain []byte
	IOTA       []byte
	CName      []byte
	MXBCA      []byte
}

func ConvertToConfig(s struct {
	A          []byte
	AAAA       []byte
	MX         []byte
	BlockChain []byte
	IOTA       []byte
	CName      []byte
	MXBCA      []byte
}) Config {
	return Config{
		A:          s.A,
		AAAA:       s.AAAA,
		MX:         s.MX,
		BlockChain: s.BlockChain,
		IOTA:       s.IOTA,
		CName:      s.CName,
		MXBCA:      s.MXBCA,
	}
}

type MailInfo struct {
	Owner      common.Address
	Expiration *big.Int
	DomainHash [32]byte
	IsValid    bool
	AliasName  []byte
	BcAddress  []byte
}

func ConvertToMailInfo(s struct {
	Owner      common.Address
	Expiration *big.Int
	DomainHash [32]byte
	IsValid    bool
	AliasName  []byte
	BcAddress  []byte
}) MailInfo {
	return MailInfo{
		Owner:      s.Owner,
		Expiration: s.Expiration,
		DomainHash: s.DomainHash,
		IsValid:    s.IsValid,
		AliasName:  s.AliasName,
		BcAddress:  s.BcAddress,
	}
}
