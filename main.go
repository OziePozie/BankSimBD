package main

import (
	"BankSimBD/Components"
	"BankSimBD/db"
	"fmt"
)

func main() {
	db.ConnectToDB()

	//var scanner = bufio.NewScanner(os.Stdin)

	//Components.RegisterAccount(scanner)

	account := Components.FindAccountById(23)
	bill := account.FindAllBillsByAccountId()[0]

	fmt.Println(account.StructToJSON())
	fmt.Println(bill)
	bill.CreateCard("RU")
	fmt.Println(bill.FindAllCardsByAccountId())
}
