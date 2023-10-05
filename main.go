package main

import (
	"BankSimBD/Components"
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var account Components.Account

func main() {
	//db.ConnectToDB()

	//account := Components.GetAccount("213116111")
	fmt.Println("Выберите что вы хотите сделать:" +
		"1. Регистрация" +
		"2. Вход")
	scanner.Scan()
	choice := scanner.Text()
	mainScenario(choice)
	//fmt.Println(account.StructToJSON())
}
func mainScenario(choice string) {

	fmt.Println("Выберите что вы хотите сделать:" +
		"1. Регистрация" +
		"2. Вход")

	switch choice {
	case "1":
		Components.RegisterAccount(scanner)
		mainScenario("2")
	case "2":
		fmt.Println("Введите логин")
		scanner.Scan()
		login := scanner.Text()
		fmt.Println("Введите пароль")
		scanner.Scan()
		password := scanner.Text()
		account, _ = Components.AuthAccount(login, password)
		logicListener()
	}
}
func logicListener() {
	fmt.Println("Выберите что вы хотите сделать:" +
		"1. Карты" +
		"2. Счета" +
		"3. История")
	scanner.Scan()
	switcher := scanner.Text()
	switch switcher {
	case "1":
		cardListener()
	case "2":
		billListener()
	case "3":
		historyListener()
	default:

	}

}

func billListener() {
	if account.FindAllBillsByAccountId() == nil {
		fmt.Println("У вас нет счетов, создать ?" +
			"1. Da" +
			"2. Net")
		scanner.Scan()
		switcher := scanner.Text()
		switch switcher {
		case "1":
			account.CreateBill()
		case "2":
			logicListener()
		}

	}
	fmt.Println("Выберите что вы хотите сделать:" +
		"1. Посмотреть список карт" +
		"2. Создать новый счет " +
		"3. Установить лимит для счета" +
		"4. Посмотреть баланс на счете" +
		"5. Закрыть счет")
	scanner.Scan()
	switcher := scanner.Text()

	switch switcher {

	}
}

func cardListener() {

}
func historyListener() {

}
