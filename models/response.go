package models

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type TransactionsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}
type TransactionResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}
type WalletsResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Wallet `json:"data"`
}
type WalletResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Wallet `json:"data"`
}
