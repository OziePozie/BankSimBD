package Components

import (
	"BankSimBD/db"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        float64   `json:"balance"`
	CurrencyTag    string    `json:"CurrencyTag"`
	History        []History `json:"history"`
	IsCardActive   bool      `json:"isCardActive"`
}

func (bill Bill) CreateCard(currencyTag string) bool {
	connectToDB := db.ConnectToDB()
	fmt.Println(currencyTag)
	currencyId := FindCurrencyIdByTag(currencyTag)
	fmt.Println(currencyId)
	query := `INSERT INTO cards (bill_id, number, cvv, expiration_date, iscardactive, currency_id)
            values ($1,$2,$3, $4, $5, $6);`

	stmt, _ := connectToDB.Prepare(query)

	cvv, number := randomCVVAndNumber()

	res, _ := stmt.Exec(bill.ID,
		number,
		cvv,
		time.Now().AddDate(4, 0, 0),
		true,
		currencyId)
	defer stmt.Close()

	fmt.Println(res.RowsAffected())
	return true

}

func (bill *Bill) FindAllCardsByBillId() []Card {
	var Cards []Card
	var card Card

	connectToDB := db.ConnectToDB()
	query := "SELECT number, cvv, expiration_date, iscardactive, currency_tag FROM Cards JOIN public.currency c on c.currency_id = Cards.currency_id WHERE bill_id = ($1)"
	stmt, _ := connectToDB.Prepare(query)
	rows, err := stmt.Query(bill.ID)
	if err != nil {
		return nil
	}
	for rows.Next() {
		rows.Scan(&card.Number,
			&card.Cvv,
			&card.ExpirationDate,
			&card.IsCardActive,
			&card.CurrencyTag)
		Cards = append(Cards, card)
	}
	bill.Cards = Cards
	defer stmt.Close()
	return Cards
}

//func UpdateCard(a Card) {
//	connectToDB := db.ConnectToDB()
//	query := "UPDATE Cards SET first_name = $2, second_name = $3, password = $4 WHERE Card_id = $1;"
//	stmt, _ := connectToDB.Prepare(query)
//	exec, err := stmt.Exec(a.ID, a.FirstName, a.SecondName, a.Password)
//	if err != nil {
//		return
//	}
//	defer stmt.Close()
//	fmt.Println(exec.RowsAffected())
//
//}

func DeleteCard(id int) {
	connectToDB := db.ConnectToDB()
	query := "DELETE FROM Cards WHERE Card_id = $1;"
	stmt, _ := connectToDB.Prepare(query)

	stmt.Exec(id)
	defer stmt.Close()
}

func randomCVVAndNumber() (string, string) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var number string
	for i := 0; i < 16; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	var cvv string
	for i := 0; i < 3; i++ {
		cvv += strconv.Itoa(r.Intn(10))
	}
	return cvv, number
}
