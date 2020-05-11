SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	abigen --abi BasAccountant.abi --pkg eth --type BasAccountant --out BasAccountant.go
	abigen --abi BasDomainConf.abi --pkg eth --type BasDomainConf --out BasDomainConf.go
	abigen --abi BasExpiredOwnership.abi --pkg eth --type BasExpiredOwnership --out BasExpiredOwnership.go
	abigen --abi BasMail.abi --pkg eth --type BasMail --out BasMail.go
	abigen --abi BasMailManager.abi --pkg eth --type BasMailManager --out BasMailManager.go
	abigen --abi BasMarket.abi --pkg eth --type BasMarket --out BasMarket.go
	abigen --abi BasMiner.abi --pkg eth --type BasMiner --out BasMiner.go
	abigen --abi BasOANN.abi --pkg eth --type BasOANN --out BasOANN.go
	abigen --abi BasRootDomain.abi --pkg eth --type BasRootDomain --out BasRootDomain.go
	abigen --abi BasSubDomain.abi --pkg eth --type BasSubDomain --out BasSubDomain.go
	abigen --abi BasToken.abi --pkg eth --type BasToken --out BasToken.go
	abigen --abi BasTradableOwnership.abi --pkg eth --type BasTradableOwnership --out BasTradableOwnership.go
	abigen --abi BasView.abi --pkg eth --type BasView --out BasView.go