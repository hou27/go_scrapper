package main

import (
	"fmt"

	"github.com/hou27/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("hou27")
	account.Deposit(10)
	fmt.Println(account)
	fmt.Println(account.Balance())
}