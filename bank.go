package main

import (
	"fmt"
)

type Bank struct {
	AccountIds map[string]float64
}

func (this *Bank) Add_Account(account_id string, balance float64) bool {
	_, exists := this.AccountIds[account_id]
	if exists {
		return false
	}
	this.AccountIds[account_id] = balance
	return true
}

func main4() {
	var b Bank
	b = Bank{
		AccountIds: make(map[string]float64),
	}
	bank_pointer := &b
	fmt.Println(bank_pointer.Add_Account("First", 10))
	fmt.Println(bank_pointer.Add_Account("First", 10))
	fmt.Println(bank_pointer.Add_Account("Second", 20))
	for key, val := range bank_pointer.AccountIds {
		fmt.Println("Account Name=", key, ", Balance=", val)
	}
}
