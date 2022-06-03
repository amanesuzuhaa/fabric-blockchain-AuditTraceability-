/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	rand2 "math/rand"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	rand2.Seed(time.Now().UnixNano())

	assetChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
