package eth

import (
	"Lost_and_Found/solitdy"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

// Err 定义全局变量
var Err error

// Ins 定义全局变量

var Ins *solitdy.LoAfo

// Opts 定义全局变量
var Opts *bind.TransactOpts

// Client 定义全局变量
var Client *ethclient.Client

// Key 定义全局变量
var Key *keystore.Key

// ChainID 定义全局变量
var ChainID *big.Int

// GasPrice 定义全局变量
var GasPrice *big.Int

func InitEth() {
	fmt.Println("Hello")
	b, err := ioutil.ReadFile("./keytstore/UTC--2024-03-01T12-09-22.028453400Z--ea14174e9e9b428bac158cbc566e3e9151818203")
	if err != nil {
		log.Fatal(err)
	}

	Key, err = keystore.DecryptKey(b, "123456")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("key:", Key)

	Client, err = ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	defer Client.Close()

	ChainID, err = Client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ChainID)

	GasPrice, err = Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(GasPrice)

	Ins, err = solitdy.NewLoAfo(common.HexToAddress("0xb9b2b724fA8c32b7072C4718C39bc8A905FaaD3d"), Client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Ins)

	keyfile := "./keytstore/UTC--2024-03-01T12-09-22.028453400Z--ea14174e9e9b428bac158cbc566e3e9151818203"
	reader, _ := os.Open(keyfile)
	Opts, err = bind.NewTransactorWithChainID(reader, "123456", ChainID)
	if err != nil {
		log.Fatal("NewTransactor", err)
	}
	Opts.GasLimit = 3000000
	Opts.GasPrice = GasPrice

	fmt.Println("opts:", Opts)

}
