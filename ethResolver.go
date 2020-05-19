package resolver

import (
	"fmt"
	"github.com/BASChain/go-bmail-account"
	"github.com/BASChain/go-bmail-resolver/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"net"
)

type EthResolverConf struct {
	ApiUrl  string
	MailDNS common.Address
}

var conf = []*EthResolverConf{
	&EthResolverConf{
		ApiUrl:  "https://infura.io/v3/f3245cef90ed440897e43efc6b3dd0f7",
		MailDNS: common.HexToAddress("0x58af099F693efb2b907b1450Bb0268C45bCB6b5D"),
	},
	&EthResolverConf{
		ApiUrl:  "https://ropsten.infura.io/v3/f3245cef90ed440897e43efc6b3dd0f7",
		MailDNS: common.HexToAddress("0x58af099F693efb2b907b1450Bb0268C45bCB6b5D"),
	},
}
var ResConf *EthResolverConf

type EthResolver struct {
}

func (er *EthResolver) DomainA(domain string) []net.IP {
	fmt.Println("implement me")
	return nil
}

func (er *EthResolver) DomainMX(domainMX string) ([]net.IP,[]bmail.Address) {
	fmt.Println("implement me")
	return nil,nil
}

func (er *EthResolver) BMailBCA(mailHash string) (bmail.Address, string) {

	hash := common.HexToHash(mailHash) //crypto.Keccak256Hash([]byte(mailName))
	fmt.Println(mailHash)
	conn, err := connect()
	if err != nil {
		fmt.Println("[BMailBCA]: connect err:", err.Error())
		return "", ""
	}
	res, err := conn.QueryEmailInfo(nil, hash)
	if err != nil {
		fmt.Println("[BMailBCA]: connect err:", err.Error())
		return "", ""
	}
	return bmail.Address(res.BcAddress), string(res.AliasName)
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

func connect() (*eth.BasView, error) {
	conn, err := ethclient.Dial(ResConf.ApiUrl)
	if err != nil {
		return nil, err
	}
	return eth.NewBasView(ResConf.MailDNS, conn)
}
