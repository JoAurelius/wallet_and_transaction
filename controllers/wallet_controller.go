package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quiz_pbp_1/models"
)

func GetAllWalletTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	vars := mux.Vars(r)
	query := WalletTransactionQuery() + " WHERE wallet.id = " + vars["wallet_id"]

	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(404, "Query Error", w)
		return
	}
	var wallet = GetWallet(vars["wallet_id"], w)
	var transaction models.Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction.ID,
			&transaction.IDWallet,
			&transaction.Datetime,
			&transaction.Amount,
			&transaction.Description); err != nil {
			SendErrorResponse(404, "Wallet Not Found", w)
			return
		} else {
			wallet.Transactions = append(wallet.Transactions, transaction)
		}
	}
	if len(wallet.Transactions) > 0 {
		SendSuccessWalletResponse(200, "Success", wallet, w)
	} else {
		SendErrorResponse(400, "No wallet or transaction from that ID", w)
	}

}
func GetWallet(walletId string, w http.ResponseWriter) models.Wallet {
	db := Connect()
	defer db.Close()
	query := "SELECT * from wallet WHERE ID = " + walletId
	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(404, "Wallet Not Found", w)
	}
	var wallet models.Wallet
	for rows.Next() {
		if err := rows.Scan(&wallet.ID,
			&wallet.Currency,
			&wallet.Username,
			&wallet.Password,
			&wallet.DisableUser); err != nil {
			SendErrorResponse(404, "Wallet not found", w)
		}
	}
	return wallet
}
func WalletTransactionQuery() string {
	return `SELECT 
    transaction.id,
    transaction.idWallet,
    transaction.datetime,
    transaction.amount,
    transaction.description
FROM wallet JOIN transaction
ON wallet.id = transaction.idWallet `
}
func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	vars := mux.Vars(r)
	var wallet = GetWallet(vars["wallet_id"], w)
	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(404, "Form is empty!", w)
		return
	}
	if (r.Form.Get("currency")) != "" {
		wallet.Currency = r.Form.Get("currency")
	}
	if (r.Form.Get("username")) != "" {
		wallet.Username = r.Form.Get("username")
	}
	if (r.Form.Get("password")) != "" {
		wallet.Password = r.Form.Get("password")
	}
	result, _ := db.Exec("UPDATE wallet SET username = ?, currency = ?, password =  ? WHERE id = ?",
		wallet.Username, wallet.Currency, wallet.Password, wallet.ID)

	num, _ := result.RowsAffected()

	if num != 0 {
		SendSuccessWalletResponse(200, "Update Success", wallet, w)
	} else {
		SendErrorResponse(404, "Update Failed", w)
	}
}
func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	vars := mux.Vars(r)
	query := WalletTransactionQuery() + " WHERE wallet.id = " + vars["wallet_id"]

	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(404, "Query Error", w)
		return
	}
	var wallet = GetWallet(vars["wallet_id"], w)
	var transaction models.Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction.ID,
			&transaction.IDWallet,
			&transaction.Datetime,
			&transaction.Amount,
			&transaction.Description); err != nil {
			SendErrorResponse(404, "Wallet Not Found", w)
			return
		} else {
			wallet.Transactions = append(wallet.Transactions, transaction)
		}
	}

	if len(wallet.Transactions) == 0 {
		result, _ := db.Exec("DELETE FROM wallet WHERE id = ?",
			wallet.ID)
		num, _ := result.RowsAffected()

		if num != 0 {
			SendSuccessResponse(200, "Hard Delete Success", w)
		} else {
			SendErrorResponse(404, "Hard Delete Failed", w)
		}
	} else {
		result, _ := db.Exec("UPDATE wallet SET disableUser = 1 WHERE id = ?",
			wallet.ID)
		num, _ := result.RowsAffected()

		if num != 0 {
			SendSuccessResponse(200, "Delete Success", w)
		} else {
			SendErrorResponse(404, "Delete Failed", w)
		}
	}

}
func Login(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}

	var username = r.Form.Get("username")
	var password = r.Form.Get("password")
	vars := mux.Vars(r)
	var wallet = GetWallet(vars["login_id"], w)
	if wallet.Password == password && wallet.Username == username {
		SendSuccessResponse(200, "Login Success", w)
	} else {
		SendErrorResponse(400, "Login Failed.\n Recheck your credential", w)
	}
}
