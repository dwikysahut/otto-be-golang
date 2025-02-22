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

func TestCreateBrand(t *testing.T) {
	database.SetupTestDB()

	brand := models.Brand{
		Name: "Test Brand2",
	}

	brandData, err := json.Marshal(brand)
	if err != nil {
		t.Fatalf("Error marshalling brand data: %v", err)
	}

	req := httptest.NewRequest("POST", "/brand", bytes.NewReader(brandData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateBrand)
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

	responseBrand, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected response data to be a map, got %T", response.Data)
	}

	if responseBrand["name"] != brand.Name {
		t.Errorf("Expected brand name '%s', got '%s'", brand.Name, responseBrand["name"])
	}

	var dbBrand models.Brand
	if err := database.DB.First(&dbBrand, "name = ?", brand.Name).Error; err != nil {
		t.Fatalf("Error fetching brand from the database: %v", err)
	}

	if dbBrand.ID == 0 {
		t.Error("Expected brand ID to be generated, but got 0")
	}
}
