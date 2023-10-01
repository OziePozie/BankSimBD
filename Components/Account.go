package Components

import (
	"BankSimBD/db"
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

func CreateAccount(a Account) bool {
	connectToDB := db.ConnectToDB()
	query := "INSERT INTO accounts (first_name, second_name, email, password) values ($1,$2,$3,$4);"
	stmt, _ := connectToDB.Prepare(query)
	res, err := stmt.Exec(a.FirstName, a.SecondName, a.Login, a.Password)
	fmt.Println(err)
	fmt.Println(res.RowsAffected())
	return true

}

func FindAll() []Account {
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
		rows.Scan(&account.ID, &account.FirstName, &account.SecondName, &account.Login, &account.Password)
		accounts = append(accounts, account)
	}

	return accounts
}

func UpdateAccount(a Account) {
	connectToDB := db.ConnectToDB()
	query := "UPDATE accounts SET first_name = $2, second_name = $3, password = $4 WHERE account_id = $1;"
	stmt, _ := connectToDB.Prepare(query)
	exec, err := stmt.Exec(a.ID, a.FirstName, a.SecondName, a.Password)
	if err != nil {
		return
	}
	fmt.Println(exec.RowsAffected())

}

func DeleteAccount(id int) {
	connectToDB := db.ConnectToDB()
	query := "DELETE FROM accounts WHERE account_id = $1;"
	stmt, _ := connectToDB.Prepare(query)
	stmt.Exec(id)
}
