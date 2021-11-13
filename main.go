package main

import (
	"fmt"

	"github.com/hou27/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("hou27")

	fmt.Println(account)
}