package main

import (
	"fmt"
	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("niko")
	account.Deposit(100)
	fmt.Println(account)
	err := account.Withdraw(200)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account)
}
