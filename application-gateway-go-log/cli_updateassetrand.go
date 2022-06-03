package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI) updateAssetRand(contract *client.Contract, id string)  {
	fmt.Println("UpdateAssetRand")
	UpdateAsset(contract, id, randUser(),randInfo())
}