package handlers

import (
	"encoding/json"
	"go-api/database"
	"go-api/models"
	"go-api/utils"
	"net/http"
)

func MakeRedemption(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	var totalPoints int
	var transactionDetails []models.TransactionDetail

	for _, detail := range transaction.Details {
		var voucher models.Voucher
		if err := database.DB.First(&voucher, detail.VoucherID).Error; err != nil {
			utils.SendError(w, http.StatusBadRequest, "Voucher not found")
			return
		}

		totalCost := voucher.CostInPoint * detail.Quantity
		totalPoints += totalCost

		transactionDetails = append(transactionDetails, models.TransactionDetail{
			VoucherID: voucher.ID,
			Quantity:  detail.Quantity,
			TotalCost: totalCost,
		})
	}

	transaction.TotalPoints = totalPoints
	if err := database.DB.Create(&transaction).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create transaction")
		return
	}

	for i := range transactionDetails {
		transactionDetails[i].TransactionID = transaction.ID
	}
	if err := database.DB.Create(&transactionDetails).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create transaction details")
		return
	}

	utils.SendSuccess(w, transaction, http.StatusCreated)
}
func GetTransactionDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("transactionId")
	if id == "" {
		utils.SendError(w, http.StatusBadRequest, "Transaction ID is required")
		return
	}

	var transaction models.Transaction
	if err := database.DB.Preload("Details").
		Preload("Details.Voucher").
		Preload("Details.Voucher.Brand"). // Add preload for Brand here
		First(&transaction, id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Transaction not found")
		return
	}
	utils.SendSuccess(w, transaction, http.StatusOK)
}
