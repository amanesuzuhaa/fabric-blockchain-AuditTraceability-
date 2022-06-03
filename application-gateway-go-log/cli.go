package main

import (
	"flag"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
	"os"
)

// CLI responsible for processing command line arguments
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  InitLedge - init assets")
	fmt.Println("  InitAssetsFromMysql - init assets from mysql")
	fmt.Println("  BackupToJson - backup assets and asset's history to a json file")
	fmt.Println("  GetAllAssets - get all assets from the ledge")

	//for test
	fmt.Println("  CreateAssetRand -id ID - create a now Asset randomly by id")
	fmt.Println("  UpdateAssetRand -id ID - update an asset randomly by id")
	//for test
	fmt.Println("  CreateAsset -id ID -user USER -info -INFO - create a now Asset by id, user, info")
	fmt.Println("  UpdateAsset -id ID -user USER -info -INFO - update a now Asset by id, user, info")
	fmt.Println("  DeleteAsset -id ID - delete an asset by id")
	fmt.Println("  ReadAssetByID -id ID - read an asset by id")
	fmt.Println("  GetAssetHistory -id ID - get an asset history by id")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run(contract *client.Contract) {
	cli.validateArgs()

	initLedgerCmd := flag.NewFlagSet("InitLedger", flag.ExitOnError)
	initAssetsFromMysqlCmd := flag.NewFlagSet("InitAssetsFromMysql", flag.ExitOnError)
	backupToJsonCmd := flag.NewFlagSet("BackupToJson", flag.ExitOnError)
	getAllAssetsCmd := flag.NewFlagSet("GetAllAssets", flag.ExitOnError)
	//generate rand for test
	createAssetRandCmd := flag.NewFlagSet("CreateAssetRand", flag.ExitOnError)
	updateAssetRandCmd := flag.NewFlagSet("UpdateAssetRand", flag.ExitOnError)
	//for test
	createAssetCmd := flag.NewFlagSet("CreateAssetCmd",flag.ExitOnError)
	updateAssetCmd := flag.NewFlagSet("UpdateAssetCmd",flag.ExitOnError)
	deleteAssetCmd := flag.NewFlagSet("DeleteAsset",flag.ExitOnError)
	readAssetByIDCmd := flag.NewFlagSet("ReadAssetByID", flag.ExitOnError)
	getAssetHistoryCmd := flag.NewFlagSet("GetAssetHistory", flag.ExitOnError)

	//for test
	createAssetRandID := createAssetRandCmd.String("id", "", "create asset randomly with id")
	updateAssetRandID := updateAssetRandCmd.String("id", "", "update asset randomly with id")
	//for test
	createAssetID := createAssetCmd.String("id","","create asset")
	createAssetUser := createAssetCmd.String("user","","create asset")
	createAssetInfo := createAssetCmd.String("info", "", "create asset")

	updateAssetID := updateAssetCmd.String("id","","update asset")
	updateAssetUser := updateAssetCmd.String("user","","update asset")
	updateAssetInfo := updateAssetCmd.String("info","","update asset")

	deleteAssetID := deleteAssetCmd.String("id", "", "delete asset by id")
	readAssetID := readAssetByIDCmd.String("id", "", "read asset by id")
	getAssetHistoryID := getAssetHistoryCmd.String("id", "", "get asset history by id")

	switch os.Args[1] {
	case "GetAllAssets":
		err := getAllAssetsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "CreateAssetRand":
		err := createAssetRandCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "UpdateAssetRand":
		err := updateAssetRandCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "CreateAsset":
		err := createAssetCmd.Parse(os.Args[2: ])
		if err != nil {
			log.Panic(err)
		}
	case "UpdateAsset":
		err := updateAssetCmd.Parse(os.Args[2: ])
		if err != nil {
			log.Panic(err)
		}

	case "ReadAssetByID":
		err := readAssetByIDCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "GetAssetHistory":
		err := getAssetHistoryCmd.Parse(os.Args[2:])
		if err != nil{
			log.Panic(err)
		}
	case "InitLedger":
		err := initLedgerCmd.Parse(os.Args[2:])
		if err != nil{
			log.Panic(err)
		}
	case "InitAssetsFromMysql":
		err := initAssetsFromMysqlCmd.Parse(os.Args[2:])
		if err != nil{
			log.Panic(err)
		}
	case "BackupToJson":
		err := backupToJsonCmd.Parse(os.Args[2:])
		if err != nil{
			log.Panic(err)
		}
	case "DeleteAsset":
		err := deleteAssetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		cli.printUsage()
		os.Exit(1)
	}

	if initLedgerCmd.Parsed(){
		cli.initLedger(contract)
	}
	
	if getAllAssetsCmd.Parsed() {
		cli.getAllAssets(contract)
	}

	if createAssetCmd.Parsed(){
		if *createAssetID == "" || *createAssetUser == "" || *createAssetInfo == ""{
			createAssetCmd.Usage()
			os.Exit(1)
		}
		cli.createAsset(contract, *createAssetID, *createAssetUser,*createAssetInfo)
	}

	if updateAssetCmd.Parsed(){
		if *updateAssetID == "" || *updateAssetUser == "" || *updateAssetInfo == ""{
			updateAssetCmd.Usage()
			os.Exit(1)
		}
		cli.updateAsset(contract, *updateAssetID, *updateAssetUser,*updateAssetInfo)
	}

	if createAssetRandCmd.Parsed() {
		if *createAssetRandID == ""{
			createAssetRandCmd.Usage()
			os.Exit(1)
		}
		cli.createAssetRand(contract, *createAssetRandID)
	}

	if updateAssetRandCmd.Parsed(){
		if *updateAssetRandID == ""{
			updateAssetRandCmd.Usage()
			os.Exit(1)
		}
		cli.updateAssetRand(contract, *updateAssetRandID)
	}

	if deleteAssetCmd.Parsed(){
		if *deleteAssetID == ""{
			deleteAssetCmd.Usage()
			os.Exit(1)
		}
		cli.deleteAsset(contract, *deleteAssetID)
	}

	if readAssetByIDCmd.Parsed() {
		if *readAssetID == ""{
			readAssetByIDCmd.Usage()
			os.Exit(1)
		}
		cli.readAssetByID(contract, *readAssetID)
	}

	if getAssetHistoryCmd.Parsed(){
		if *getAssetHistoryID == ""{
			getAssetHistoryCmd.Usage()
			os.Exit(1)
		}
		cli.getAssetHistory(contract, *getAssetHistoryID)
	}
	if initAssetsFromMysqlCmd.Parsed(){
		cli.initAssetsFromMysql(contract)
	}
	if backupToJsonCmd.Parsed(){
		cli.backupToJson(contract)
	}
}