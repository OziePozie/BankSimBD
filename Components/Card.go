package Components

import (
	"time"
)

type Card struct {
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        Balance   `json:"balance"`
	History        []History `json:"history"`
	IsCardActive   bool      `json:"isCardActive"`
}
