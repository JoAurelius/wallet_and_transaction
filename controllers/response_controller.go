package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/quiz_pbp_1/models"
)

func SendErrorResponse(status int, errMsg string, w http.ResponseWriter) {
	var response models.GeneralResponse
	response.Status = status
	response.Message = errMsg
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponse(status int, msg string, w http.ResponseWriter) {
	var response models.GeneralResponse
	response.Status = status
	response.Message = msg
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func SendSuccessTransactionsResponse(status int, msg string, transaction []models.Transaction, w http.ResponseWriter) {
	var response models.TransactionsResponse
	response.Status = status
	response.Message = msg
	response.Data = transaction
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
func SendSuccessTransactionResponse(status int, msg string, transaction models.Transaction, w http.ResponseWriter) {
	var response models.TransactionResponse
	response.Status = status
	response.Message = msg
	response.Data = transaction
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func SendSuccessWalletResponse(status int, msg string, wallet models.Wallet, w http.ResponseWriter) {
	var response models.WalletResponse
	response.Status = status
	response.Message = msg
	response.Data = wallet
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
func SendSuccessWalletsResponse(status int, msg string, wallet []models.Wallet, w http.ResponseWriter) {
	var response models.WalletsResponse
	response.Status = status
	response.Message = msg
	response.Data = wallet
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
