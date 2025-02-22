package handlers

import (
	"encoding/json"
	"net/http"

	"go-api/database"
	"go-api/models"
	"go-api/utils"
)

func CreateBrand(w http.ResponseWriter, r *http.Request) {
	var brand models.Brand
	if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := database.DB.Create(&brand).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create brand")
		return
	}

	utils.SendSuccess(w, brand, http.StatusCreated)
}
