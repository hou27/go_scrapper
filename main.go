package main

import (
	"fmt"
	"log"

	"github.com/hou27/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jalapeno")
	account.Deposit(10)
	err := account.Withdraw(6)
	if err != nil { // err handling
		log.Fatalln(err)
	}
	fmt.Println(account)
}