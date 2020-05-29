package resolver

import (
	"fmt"
	"github.com/BASChain/go-bmail-account"
	"github.com/BASChain/go-bmail-resolver/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"net"
)

type EthResolverConf struct {
	AccessPoints []string
	BasViewAddr  common.Address
}

var conf = []*EthResolverConf{
	{
		AccessPoints: []string{"https://infura.io/v3/f3245cef90ed440897e43efc6b3dd0f7"},
		BasViewAddr:  common.HexToAddress("0x331c08bBd8493d190906aFFcF134691846A4957F"),
	},
	{
		AccessPoints: []string{"https://ropsten.infura.io/v3/f3245cef90ed440897e43efc6b3dd0f7",
			"https://ropsten.infura.io/v3/831ab04fa4964991b5fba5c52106d7b0"},
		BasViewAddr: common.HexToAddress("0xf3e0222FC99897E3569F4490026D914A9421572a"),
	},
}
var ResConf *EthResolverConf

const Separator = 0x7f

type EthResolver struct {
}

func (er *EthResolver) DomainA(domain string) []net.IP {
	conf, err := QueryDomainConfigs(GetHash(domain))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ipStrings := Split(conf.A, Separator)
	var r []net.IP
	for _, t := range ipStrings {
		r = append(r, net.ParseIP(t))
	}
	return r
}

func (er *EthResolver) DomainMX(domain string) ([]net.IP, []bmail.Address) {
	conf, err := QueryDomainConfigs(GetHash(domain))
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	mx := Split(conf.MX, Separator)
	var ips []net.IP
	for _, t := range mx {
		ips = append(ips, net.ParseIP(t))
	}

	mxbca := Split(conf.MXBCA, Separator)
	var bca []bmail.Address
	for _, t := range mxbca {
		bca = append(bca, bmail.Address(t))
	}
	return ips, bca
}

func (er *EthResolver) BMailBCA(mailName string) (bmail.Address, string) {
	info, err := QueryEmailInfo(GetHash(mailName))
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	return info.BcAddress, info.AliasName
}

func QueryEmailInfo(hash Hash) (*MailInfo, error) {
	opts := GetCallOpts(0)
	conn := connect()
	defer conn.Close()
	result, err := BasView(conn).QueryEmailInfo(opts, hash)
	if err != nil {
		return nil, err
	}
	r := ConvertToMailInfo(result)
	return &r, nil
}

func QueryDomainConfigs(hash Hash) (*Config, error) {
	opts := GetCallOpts(0)
	conn := connect()
	defer conn.Close()
	result, err := BasView(conn).QueryDomainConfigs(opts, hash)
	if err != nil {
		return nil, err
	}
	r := ConvertToConfig(result)
	return &r, nil
}

func NewEthResolver(debug bool) NameResolver {
	obj := &EthResolver{}
	if debug {
		ResConf = conf[1]
	} else {
		ResConf = conf[0]
	}
	return obj
}

func BasView(conn *ethclient.Client) *eth.BasView {
	if instance, err := eth.NewBasView(ResConf.BasViewAddr, conn); err == nil {
		return instance
	} else {
		logger.Error("can't recover BasView instance, ", err)
		return nil
	}
}

func connect() *ethclient.Client {
	for _, s := range ResConf.AccessPoints {
		c, err := ethclient.Dial(s)
		if err != nil {
			continue
		} else {
			return c
		}
	}
	logger.Error("all access points failed, please check network!!!")
	return nil
}

func GetCallOpts(blockNumber uint64) *bind.CallOpts {
	var opts = new(bind.CallOpts)
	if blockNumber == 0 {
		opts = nil
	} else {
		opts.BlockNumber = new(big.Int).SetUint64(blockNumber)
	}
	return opts
}
