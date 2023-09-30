package main

import (
	"BankSimBD/Components"
	"BankSimBD/db"
)

func main() {
	db.ConnectToDB()

	Components.Account.CreateAccount(Components.Account{
		FirstName:  "123",
		SecondName: "123",
		Login:      "123",
		Password:   "123",
	})
}
