package eth

import (
	"fmt"
	"log"
	"net/rpc"
)

func NewAcc(pass string) {
	cli, err := rpc.Dial("tcp", "localhost:8545")
	if err != nil {
		log.Fatal("faild to connect to geth", err)
	}
	defer cli.Close()
	var account string
	err = cli.Call("personal_newAccount", &account, pass)
	if err != nil {
		log.Fatal("faild to new account personal", err)
	}
	fmt.Println("account ======", account)
}
