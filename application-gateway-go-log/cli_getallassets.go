package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI) getAllAssets(contract *client.Contract) {
	fmt.Println("GetAllAssets:")
	GetAllAssets(contract)
}
