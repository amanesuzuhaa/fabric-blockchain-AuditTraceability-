/*
Copyright 2021 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

const (
	mspID         = "Org1MSP"
	cryptoPath    = "../../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com"
	certPath      = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
	keyPath       = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	tlsCertPath   = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint  = "localhost:7051"
	gatewayPeer   = "peer0.org1.example.com"
	channelName   = "mychannel"
	chaincodeName = "log"
)

var now = time.Now()


// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func InitLedger(contract *client.Contract) {
	fmt.Printf("Submit Transaction: InitLedger, function creates the initial set of assets on the ledger \n")

	_, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

func GetAllAssets(contract *client.Contract) {
	fmt.Println("Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}

func BackupToJson(contract *client.Contract)  {
	fmt.Println("Evaluate Transaction: BackupToJson, function returns all the current assets")

	evaluateResult, err := contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	//创建备份目录
	folderName := "Backup" + time.Now().Local().Format(time.RFC3339)
	os.Mkdir(folderName, os.ModePerm)

	//备份assets
	fileName := folderName+"/Assets.json"
	AssetFile,err  := os.Create(fileName)  //创建该文件，返回句柄
	defer AssetFile.Close()      //确保文件在该函数执行完以后关闭
	if err != nil {
		log.Fatalln("open file error !")
	}

	debugJson := log.New(AssetFile,"",log.Lmsgprefix)
	debugJson.Print(string(evaluateResult))

	//备份 history
	var assets []Asset
	json.Unmarshal(evaluateResult, &assets)

	for _, asset := range assets{
		evaluateResult2, err := contract.EvaluateTransaction("GetAssetHistory", asset.ID)

		if err != nil {
			panic(fmt.Errorf("failed to evaluate transaction: %w", err))
		}

		fileName2 := folderName+"/History-id==<"+asset.ID+">.json"
		HistoryFile,err  := os.Create(fileName2)  //创建该文件，返回句柄
		defer HistoryFile.Close()      //确保文件在该函数执行完以后关闭
		if err != nil {
			log.Fatalln("open file error !")
		}
		debugHistory := log.New(HistoryFile,"",log.Lmsgprefix)
		debugHistory.Print(string(evaluateResult2))
	}

}

func InitAssetsFromMysql(contract *client.Contract)  {
	fmt.Printf("Submit Transaction: CreateAsset, creates new asset from mysql \n")
	assets,_ := getAssetFromMysql(sql)
	for _,asset := range assets{
		fmt.Println(asset.ID, asset.USER, asset.INFO)
		_, err := contract.SubmitTransaction("CreateAsset", asset.ID, asset.USER, asset.INFO)
		if err != nil {
			panic(fmt.Errorf("failed to submit transaction: %w", err))
		}
	}
	fmt.Printf("*** Transaction committed successfully\n")
}

func CreateAsset(contract *client.Contract, id string, user string, info string) {
	fmt.Printf("Submit Transaction: CreateAsset, creates new asset with ID, USER and INFO arguments \n")

	_, err := contract.SubmitTransaction("CreateAsset", id, user, info)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

func UpdateAsset(contract *client.Contract, id string, user string, info string){
	fmt.Printf("Submit Transaction: UpdateAsset, update an asset by id \n")

	_, err := contract.SubmitTransaction("UpdateAsset", id, user, info)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

func DeleteAsset(contract *client.Contract, id string){
	fmt.Printf("Submit Transaction: DeleteAsset, delete an asset by id\n")

	_, err := contract.SubmitTransaction("DeleteAsset", id)

	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

func ReadAssetByID(contract *client.Contract, id string) {
	fmt.Printf("Evaluate Transaction: ReadAsset, function returns asset attributes\n")

	evaluateResult, err := contract.EvaluateTransaction("ReadAsset", id)
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}

// GetAssetHistory query the evaluated history of an asset by id
func GetAssetHistory(contract *client.Contract, id string)  {
	fmt.Printf("Evaluate Transaction: GetAssetHistory, function returns asset history\n")

	evaluateResult, err := contract.EvaluateTransaction("GetAssetHistory", id)

	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}
