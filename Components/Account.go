package Components

import (
	"BankSimBD/db"
	"bufio"
	"encoding/json"
	"fmt"
)

type Account struct {
	ID         int    `json:"ID"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Bill       []Bill `json:"bills"`
}

type AccountDetails struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

func CreateAccount(a AccountDetails) bool {
	connectToDB := db.ConnectToDB()
	query := "INSERT INTO accounts (first_name, second_name, email, password) values ($1,$2,$3,$4);"
	stmt, _ := connectToDB.Prepare(query)
	res, _ := stmt.Exec(a.FirstName, a.SecondName, a.Login, a.Password)
	defer stmt.Close()

	fmt.Println(res.RowsAffected())
	return true

}

func FindAllAccounts() []Account {
	var accounts []Account
	var account Account
	connectToDB := db.ConnectToDB()
	query := "SELECT * FROM accounts"
	stmt, _ := connectToDB.Prepare(query)
	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	for rows.Next() {
		rows.Scan(&account.ID,
			&account.FirstName,
			&account.SecondName,
			&account.Login,
			&account.Password)
		accounts = append(accounts, account)
	}
	defer stmt.Close()
	return accounts
}

func FindAccountById(id int) Account {
	var account Account
	connectToDB := db.ConnectToDB()
	query := "SELECT * FROM accounts WHERE account_id = ($1)"
	stmt, _ := connectToDB.Prepare(query)
	rows, _ := stmt.Query(id)

	for rows.Next() {
		rows.Scan(&account.ID,
			&account.FirstName,
			&account.SecondName,
			&account.Login,
			&account.Password)

	}
	defer stmt.Close()
	return account
}
func FindAccountByLogin(email string) Account {
	var account Account
	connectToDB := db.ConnectToDB()
	query := "SELECT * FROM accounts WHERE email = ($1)"
	stmt, _ := connectToDB.Prepare(query)
	rows, _ := stmt.Query(email)

	for rows.Next() {
		rows.Scan(&account.ID,
			&account.FirstName,
			&account.SecondName,
			&account.Login,
			&account.Password)

	}
	defer stmt.Close()
	return account
}

func UpdateAccount(a Account) {
	connectToDB := db.ConnectToDB()
	query := "UPDATE accounts SET first_name = $2, second_name = $3, password = $4 WHERE account_id = $1;"
	stmt, _ := connectToDB.Prepare(query)
	exec, err := stmt.Exec(a.ID, a.FirstName, a.SecondName, a.Password)
	if err != nil {
		return
	}
	defer stmt.Close()
	fmt.Println(exec.RowsAffected())

}

func DeleteAccount(id int) {
	connectToDB := db.ConnectToDB()
	query := "DELETE FROM accounts WHERE account_id = $1;"
	stmt, _ := connectToDB.Prepare(query)

	stmt.Exec(id)
	defer stmt.Close()
}
func (account Account) JSONToStruct(b []byte) Account {

	_ = json.Unmarshal(b, &account)
	return account
}

func (account Account) StructToJSON() string {
	b, err := json.Marshal(account)
	if err != nil {

	}
	return string(b)
}

func RegisterAccount(scanner *bufio.Scanner) {
	fmt.Println("Регистрация пользователя")
	var account AccountDetails
	fmt.Print("Введите Имя: ")
	scanner.Scan()
	account.FirstName = scanner.Text()

	fmt.Print("Введите Фамилию: ")
	scanner.Scan()
	account.SecondName = scanner.Text()

	fmt.Print("Введите логин: ")
	scanner.Scan()
	account.Login = scanner.Text()

	fmt.Print("Введите пароль: ")
	scanner.Scan()
	account.Password = scanner.Text()

	account.createAccount(account.Login, account.Password,
		account.FirstName, account.SecondName)

	fmt.Println("Регистрация успешно завершена!")
}

func GetAccount(email string) Account {
	var acc Account
	acc = FindAccountByLogin(email)
	bills := acc.FindAllBillsByAccountId()
	for _, bill := range bills {
		cards := bill.FindAllCardsByBillId()
		bill.Cards = cards
		bills = append(bills, bill)
	}
	acc.Bill = bills
	return acc
}

func (account AccountDetails) createAccount(login, password, firstName, secondName string) bool {

	account.FirstName = firstName

	account.SecondName = secondName
	account.Login = login

	account.Password = password

	CreateAccount(account)

	return true
}
func AuthAccount(login, password string) (Account, error) {
	account := FindAccountByLogin(login)
	if account.Password == password {
		return account, nil
	} else {

		panic("Неправильные креды")

	}
}
