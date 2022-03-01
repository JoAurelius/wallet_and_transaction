package controllers

import (
	"net/http"
	"strconv"

	"github.com/quiz_pbp_1/models"
)

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}
	var transaction models.Transaction
	transaction.IDWallet, _ = strconv.Atoi(r.Form.Get("id_wallet"))
	transaction.Datetime = r.Form.Get("datetime")
	transaction.Amount, _ = strconv.Atoi(r.Form.Get("amount"))
	transaction.Description = r.Form.Get("desc")

	result, _ := db.Exec("insert into transaction (IDWallet, datetime, amount, description) values (?, ?, ?, ?)",
		transaction.IDWallet, transaction.Datetime, transaction.Amount, transaction.Description)

	num, _ := result.RowsAffected()

	if num != 0 {
		SendSuccessTransactionResponse(200, "Insert Transaction Sucess", transaction, w)
	} else {
		SendErrorResponse(404, "Insert failed", w)
	}
}
