package Components

type Bill struct {
	Number       int64     `json:"number"`
	Balance      Balance   `json:"balance"`
	Cards        []Card    `json:"cards"`
	History      []History `json:"history"`
	Limit        int       `json:"limit"`
	IsBillActive bool      `json:"isBillActive"`
}
