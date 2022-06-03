package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI)initLedger(contract *client.Contract)  {
	fmt.Println("InitLedger")
	InitLedger(contract)
}
