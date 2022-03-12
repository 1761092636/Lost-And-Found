package main

//func NewAcc(pass string) {
//	cli, err := ethclient.Dial("http://localhost:8545")
//	if err != nil {
//		log.Fatal("faild to connect to geth", err)
//	}
//	defer cli.Close()
//	ins, err := contract.(common.HexToAddress("0xd9145CCE52D386f254917e481eB44e9943F39138"), cli)
//	if err != nil {
//		log.Fatal("faild to new account personal", err)
//	}
//
//}

import (
	"CDC/CDC-/solitdy"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Hello")
	b, err := ioutil.ReadFile("D:\\Workspace\\MyGo\\src\\CDC\\CDC-\\keytstore\\UTC--2022-02-27T12-52-27.002839400Z--6fd75d526050fefe56e297af413d0af9249ef174")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "123456")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("key:", key)

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(chainID)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gasPrice)

	ins, err := solitdy.NewTest(common.HexToAddress("0xd9145CCE52D386f254917e481eB44e9943F39138"), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ins)
	val, err := ins.Nowacc(nil)
	fmt.Println("返回:", val)

	//keyfile := "UTC--2021-09-15T11-25-16.664077600Z--a7385378bdbc97951cb5e534305c82b9b615510d"
	//reader, _ := os.Open(keyfile)
	//opts, err := bind.bindMain(reader, "123456")
	//if err != nil {
	// log.Fatal("NewTransactor", err)
	//}
	//opts.GasLimit = 3000000
	//opts.GasPrice = gasPrice

	//fmt.Println("opts:", opts)

}
