package resolver

import (
	"github.com/BASChain/go-bmail-account"
	"github.com/op/go-logging"
	"net"
)

var logger, _ = logging.GetLogger("resolver")

type NameResolver interface {
	DomainA(domainName string) []net.IP
	DomainMX(domainName string) ([]net.IP, []bmail.Address)
	BMailBCA(mailName string) (address bmail.Address, cname string)
}
