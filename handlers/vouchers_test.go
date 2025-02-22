package handlers

import (
	"bytes"
	"encoding/json"
	"go-api/database"
	"go-api/models"
	"go-api/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateVouchers(t *testing.T) {
	database.SetupTestDB()

	voucher := models.Voucher{
		BrandID:     1,
		Name:        "voucher-1",
		CostInPoint: 5000,
	}

	voucherData, err := json.Marshal(voucher)
	if err != nil {
		t.Fatalf("Error marshalling voucher data: %v", err)
	}

	req := httptest.NewRequest("POST", "/brand", bytes.NewReader(voucherData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateVoucher)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, status)
	}

	var response utils.Response
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	log.Println(response)

	if response.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", response.Status)
	}

	if response.Data == nil {
		t.Fatalf("Expected response data, got nil")
	}

	responseVoucher, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected response data to be a map, got %T", response.Data)
	}

	if responseVoucher["name"] != voucher.Name {
		t.Errorf("Expected Voucher name '%s', got '%s'", voucher.Name, responseVoucher["name"])
	}

	var dbVoucher models.Voucher
	if err := database.DB.First(&dbVoucher, "name = ?", voucher.Name).Error; err != nil {
		t.Fatalf("Error fetching voucher from the database: %v", err)
	}

	if dbVoucher.ID == 0 {
		t.Error("Expected Voucher ID to be generated, but got 0")
	}
}
