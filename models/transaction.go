package models

type Transaction struct {
	ID          int    `json:"id"`
	IDWallet    int    `json:"id_wallet"`
	Datetime    string `json:"datetime"`
	Amount      int    `json:"amount"`
	Description string `json:"desc"`
}
