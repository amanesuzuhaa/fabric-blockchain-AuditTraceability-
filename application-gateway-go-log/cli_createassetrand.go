package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI) createAssetRand(contract *client.Contract, id string)  {
	fmt.Println("CreateAssetRand")
	CreateAsset(contract, id, randUser(), randInfo())
}
