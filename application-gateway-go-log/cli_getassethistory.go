package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI) getAssetHistory(contract *client.Contract, id string) {
	fmt.Println("GetAssetHistory:")
	GetAssetHistory(contract, id)
}
