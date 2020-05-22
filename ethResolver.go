package resolver

import (
	"github.com/BASChain/go-bmail-account"
	"github.com/BASChain/go-bmail-resolver/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"net"
	"time"
)

type EthResolverConf struct {
	AccessPoints []string
	BasViewAddr  common.Address
}

var conf = []*EthResolverConf{
	{
		AccessPoints: []string{"https://infura.io/v3/f3245cef90ed440897e43efc6b3dd0f7"},
		BasViewAddr:  common.HexToAddress("0x58af099F693efb2b907b1450Bb0268C45bCB6b5D"),
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
	conf := QueryDomainConfigs(GetHash(domain), 0)
	ipStrings := Split(conf.A, Separator)
	var r []net.IP
	for _, t := range ipStrings {
		r = append(r, net.ParseIP(t))
	}
	return r
}

//								   				  MX，		 MXBCA
func (er *EthResolver) DomainMX(domain string) ([]net.IP, []bmail.Address) {
	conf := QueryDomainConfigs(GetHash(domain), 0)
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
	info := QueryEmailInfo(GetHash(mailName), 0)
	return bmail.Address(string(info.BcAddress)), string(info.AliasName)
}

func QueryEmailInfo(hash Hash, tryTimes int) *MailInfo {
	opts := GetCallOpts(0)
	conn := connect()
	defer conn.Close()
	result, err := BasView(conn).QueryEmailInfo(opts, hash)
	if err != nil {
		tryTimes += 1
		if tryTimes > 3 {
			logger.Error("can't query mail info after many retries", err)
			return nil
		} else {
			time.Sleep(time.Duration(RetryRule[tryTimes]) * time.Second)
			return QueryEmailInfo(hash, tryTimes)
		}
	} else {
		r := ConvertToMailInfo(result)
		return &r
	}
}

func QueryDomainConfigs(hash Hash, tryTimes int) *Config {
	opts := GetCallOpts(0)
	conn := connect()
	defer conn.Close()
	result, err := BasView(conn).QueryDomainConfigs(opts, hash)
	if err != nil {
		tryTimes += 1
		if tryTimes > 3 {
			logger.Error("can't query domain config after many retries", err)
			return nil
		} else {
			time.Sleep(time.Duration(RetryRule[tryTimes]) * time.Second)
			return QueryDomainConfigs(hash, tryTimes)
		}
	} else {
		r := ConvertToConfig(result)
		return &r
	}
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
