package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI)updateAsset(contract *client.Contract, id, user, info string)  {
	fmt.Println("UpdateAsset")
	UpdateAsset(contract, id, user, info)
}
