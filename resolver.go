package resolver

import (
	"github.com/realbmail/go-bmail-account"
	"github.com/op/go-logging"
	"net"
)

var logger, _ = logging.GetLogger("resolver")

type NameResolver interface {
	DomainA(domainName string) []net.IP
	DomainA2(domainName string) ([]net.IP,error)
	DomainA3(domain string) ([]net.IP, []string, error)
	DomainMX(domainName string) ([]net.IP, []bmail.Address)
	BMailBCA(mailName string) (address bmail.Address, cname string)
}
