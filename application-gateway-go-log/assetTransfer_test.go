package main

import (
	"fmt"
	"testing"
)

func BenchmarkReadAssetByID(b *testing.B) {
	//clientConnection := newGrpcConnection()
	//defer clientConnection.Close()
	//
	//id := newIdentity()
	//sign := newSign()
	//
	//// Create a Gateway connection for a specific client identity
	//gateway, err := client.Connect(
	//	id,
	//	client.WithSign(sign),
	//	client.WithClientConnection(clientConnection),
	//	// Default timeouts for different gRPC calls
	//	client.WithEvaluateTimeout(5*time.Second),
	//	client.WithEndorseTimeout(15*time.Second),
	//	client.WithSubmitTimeout(5*time.Second),
	//	client.WithCommitStatusTimeout(1*time.Minute),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//defer gateway.Close()
	//
	//network := gateway.GetNetwork(channelName)
	//contract := network.GetContract(chaincodeName)


	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		//需要进行测试性能的代码段
		ReadAssetByID(contract,"9")
	}
}

func BenchmarkGetAssetHistory(b *testing.B) {
	//clientConnection := newGrpcConnection()
	//defer clientConnection.Close()
	//
	//id := newIdentity()
	//sign := newSign()
	//
	//// Create a Gateway connection for a specific client identity
	//gateway, err := client.Connect(
	//	id,
	//	client.WithSign(sign),
	//	client.WithClientConnection(clientConnection),
	//	// Default timeouts for different gRPC calls
	//	client.WithEvaluateTimeout(5*time.Second),
	//	client.WithEndorseTimeout(15*time.Second),
	//	client.WithSubmitTimeout(5*time.Second),
	//	client.WithCommitStatusTimeout(1*time.Minute),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//defer gateway.Close()
	//
	//network := gateway.GetNetwork(channelName)
	//contract := network.GetContract(chaincodeName)

	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		//需要进行测试性能的代码段
		GetAssetHistory(contract,"9")
	}
}

func BenchmarkCreateAsset(b *testing.B) {
	//clientConnection := newGrpcConnection()
	//defer clientConnection.Close()
	//
	//id := newIdentity()
	//sign := newSign()
	//
	//// Create a Gateway connection for a specific client identity
	//gateway, err := client.Connect(
	//	id,
	//	client.WithSign(sign),
	//	client.WithClientConnection(clientConnection),
	//	// Default timeouts for different gRPC calls
	//	client.WithEvaluateTimeout(5*time.Second),
	//	client.WithEndorseTimeout(15*time.Second),
	//	client.WithSubmitTimeout(5*time.Second),
	//	client.WithCommitStatusTimeout(1*time.Minute),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//defer gateway.Close()
	//
	//network := gateway.GetNetwork(channelName)
	//contract := network.GetContract(chaincodeName)

	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		//需要进行测试性能的代码段
		CreateAsset(contract,fmt.Sprint(i+300),"","")
	}
}

func BenchmarkUpdateAsset(b *testing.B) {
	//clientConnection := newGrpcConnection()
	//defer clientConnection.Close()
	//
	//id := newIdentity()
	//sign := newSign()
	//
	//// Create a Gateway connection for a specific client identity
	//gateway, err := client.Connect(
	//	id,
	//	client.WithSign(sign),
	//	client.WithClientConnection(clientConnection),
	//	// Default timeouts for different gRPC calls
	//	client.WithEvaluateTimeout(5*time.Second),
	//	client.WithEndorseTimeout(15*time.Second),
	//	client.WithSubmitTimeout(5*time.Second),
	//	client.WithCommitStatusTimeout(1*time.Minute),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//defer gateway.Close()
	//
	//network := gateway.GetNetwork(channelName)
	//contract := network.GetContract(chaincodeName)

	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		//需要进行测试性能的代码段
		UpdateAsset(contract,"1","","")
	}
}

func BenchmarkDeleteAsset(b *testing.B) {
	//clientConnection := newGrpcConnection()
	//defer clientConnection.Close()
	//
	//id := newIdentity()
	//sign := newSign()
	//
	//// Create a Gateway connection for a specific client identity
	//gateway, err := client.Connect(
	//	id,
	//	client.WithSign(sign),
	//	client.WithClientConnection(clientConnection),
	//	// Default timeouts for different gRPC calls
	//	client.WithEvaluateTimeout(5*time.Second),
	//	client.WithEndorseTimeout(15*time.Second),
	//	client.WithSubmitTimeout(5*time.Second),
	//	client.WithCommitStatusTimeout(1*time.Minute),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//defer gateway.Close()
	//
	//network := gateway.GetNetwork(channelName)
	//contract := network.GetContract(chaincodeName)

	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		//需要进行测试性能的代码段
		CreateAsset(contract,"1","","")
		DeleteAsset(contract,"1")
	}
}