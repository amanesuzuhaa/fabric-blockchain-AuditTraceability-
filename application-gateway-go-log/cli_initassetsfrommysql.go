package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI)initAssetsFromMysql(contract *client.Contract)  {
	fmt.Println("InitAssetsFromMysql")
	InitAssetsFromMysql(contract)
}