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

func (a Account) CreateAccount() {
	connectToDB := db.ConnectToDB()
	query := "INSERT INTO accounts (first_name, second_name, email, password) values ($1,$2,$3,$4);"
	stmt, _ := connectToDB.Prepare(query)
	res, _ := stmt.Exec(a.FirstName, a.SecondName, a.Login, a.Password)

	fmt.Println(res.RowsAffected())

}

func (a Account) ReadAccount() Account {
	var acc Account
	connectToDB := db.ConnectToDB()
	query := "SELECT * FROM accounts WHERE email = $1"
	stmt, _ := connectToDB.Prepare(query)

	stmt.QueryRow(a.Login).Scan(&a.ID, &a.FirstName, &a.SecondName, &a.Login, &a.Password)

	return acc
}

func (a Account) UpdateAccount() {
	connectToDB := db.ConnectToDB()
	query := "UPDATE accounts SET first_name = $2, second_name = $3, password = $4 WHERE email = $1;"
	stmt, _ := connectToDB.Prepare(query)
	stmt.Exec(a.Login, a.FirstName, a.SecondName, a.Password)

}

func (a Account) DeleteAccount(id int) {
	connectToDB := db.ConnectToDB()
	query := "DELETE FROM accounts WHERE account_id = $1;"
	stmt, _ := connectToDB.Prepare(query)
	stmt.Exec(id)
}
