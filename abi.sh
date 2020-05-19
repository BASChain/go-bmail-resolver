rm -rf ~/workspace/go/src/github.com/BASChain/go-bmail-resolver/eth

solc ~/workspace/contract/BASChain/smart_contract/contracts/BasView.sol --abi -o ~/workspace/go/src/github.com/BASChain/go-bmail-resolver/eth
cd  ~/workspace/go/src/github.com/BASChain/go-bmail-resolver/eth

abigen --abi BasView.abi --pkg eth --type BasView --out BasView.go