package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI)readAssetByID(contract *client.Contract, id string)  {
	fmt.Println("readAsset")
	ReadAssetByID(contract, id)
}
