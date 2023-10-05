package Components

import "BankSimBD/db"

type Currency struct {
	CurrencyId     int     `json:"currencyId"`
	CurrencyTag    string  `json:"currencyTag"`
	CourseToDollar float64 `json:"courseToDollar"`
}

func FindCurrencyIdByTag(tag string) int {
	var id int
	connectToDB := db.ConnectToDB()
	query := "SELECT currency_id FROM currency WHERE currency_tag = ($1)"

	stmt, _ := connectToDB.Prepare(query)
	row := stmt.QueryRow(tag)
	row.Scan(&id)
	return id
}
