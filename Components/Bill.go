package Components

import (
	"BankSimBD/db"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type BillDetails struct {
}

type Bill struct {
	ID           int       `json:"ID"`
	Number       string    `json:"number"`
	Limit        int       `json:"limit"`
	Balance      Balance   `json:"balance"`
	Cards        []Card    `json:"cards"`
	History      []History `json:"history"`
	IsBillActive bool      `json:"isBillActive"`
}

func (account Account) CreateBill() bool {
	connectToDB := db.ConnectToDB()
	query := "INSERT INTO Bills (account_id, number, sum_limit) values ($1,$2,$3);"
	stmt, _ := connectToDB.Prepare(query)

	number := randomNumberBill()

	res, _ := stmt.Exec(account.ID, number, 0)
	defer stmt.Close()

	fmt.Println(res.RowsAffected())
	return true

}

func randomNumberBill() string {
	var number string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < 20; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	return number
}

func (account Account) FindAllBillsByAccountId() []Bill {
	var Bills []Bill
	var bill Bill

	connectToDB := db.ConnectToDB()
	query := "SELECT bill_id, number, sum_limit FROM bills WHERE account_id = ($1)"
	stmt, _ := connectToDB.Prepare(query)
	rows, err := stmt.Query(account.ID)
	if err != nil {
		return nil
	}
	for rows.Next() {
		rows.Scan(&bill.ID,
			&bill.Number,
			&bill.Limit)
		Bills = append(Bills, bill)
	}
	defer stmt.Close()
	return Bills
}

//func UpdateBill(a Bill) {
//	connectToDB := db.ConnectToDB()
//	query := "UPDATE Bills SET first_name = $2, second_name = $3, password = $4 WHERE Bill_id = $1;"
//	stmt, _ := connectToDB.Prepare(query)
//	exec, err := stmt.Exec(a.ID, a.FirstName, a.SecondName, a.Password)
//	if err != nil {
//		return
//	}
//	defer stmt.Close()
//	fmt.Println(exec.RowsAffected())
//
//}

func DeleteBill(id int) {
	connectToDB := db.ConnectToDB()
	query := "DELETE FROM Bills WHERE Bill_id = $1;"
	stmt, _ := connectToDB.Prepare(query)

	stmt.Exec(id)
	defer stmt.Close()
}
