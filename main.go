package main

import (
	"fmt"
	"log"

	"github.com/hou27/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jalapeno")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil { // err handling
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}