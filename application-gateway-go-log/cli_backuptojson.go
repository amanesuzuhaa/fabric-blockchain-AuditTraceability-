package main

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func (cli *CLI)backupToJson(contract *client.Contract)  {
	fmt.Println("BackupToJson")
	BackupToJson(contract)
}