package handlers

import (
	"encoding/json"
	"net/http"

	"go-api/database"
	"go-api/models"
	"go-api/utils"
)

func CreateVoucher(w http.ResponseWriter, r *http.Request) {
	var voucher models.Voucher
	if err := json.NewDecoder(r.Body).Decode(&voucher); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := database.DB.Create(&voucher).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create voucher")
		return
	}

	utils.SendSuccess(w, voucher, http.StatusCreated)
}
func GetVoucher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.SendError(w, http.StatusBadRequest, "Voucher ID is required")
		return
	}

	var voucher models.Voucher
	if err := database.DB.Preload("Brand").First(&voucher, id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Voucher not found")
		return
	}

	utils.SendSuccess(w, voucher, http.StatusOK)
}
func GetVouchersByBrand(w http.ResponseWriter, r *http.Request) {
	brandID := r.URL.Query().Get("id")
	if brandID == "" {
		utils.SendError(w, http.StatusBadRequest, "Brand ID is required")
		return
	}

	var vouchers []models.Voucher
	if err := database.DB.Where("brand_id = ?", brandID).Find(&vouchers).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch vouchers")
		return
	}

	utils.SendSuccess(w, vouchers, http.StatusOK)
}
