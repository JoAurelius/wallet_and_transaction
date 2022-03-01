package models

type Wallet struct {
	ID           int           `json:"id"`
	Currency     string        `json:"currency"`
	Username     string        `json:"username"`
	Password     string        `json:"password"`
	DisableUser  bool          `json:"disable_user"`
	Transactions []Transaction `json:"transactions"`
}
