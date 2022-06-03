package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI) createAsset(contract *client.Contract, id, user, info string)  {
	fmt.Println("CreateAsset")
	CreateAsset(contract, id, user,info)
}
