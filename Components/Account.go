package Components

import (
	"BankSimBD/db"
	"fmt"
)

type Account struct {
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
