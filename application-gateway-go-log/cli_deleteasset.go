package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI)deleteAsset(contract *client.Contract, id string)  {
	fmt.Println("DeleteAsset")
	DeleteAsset(contract, id)
}