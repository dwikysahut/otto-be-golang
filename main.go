package main

import (
	"log"
	"net/http"

	"go-api/database"
	"go-api/models"
	"go-api/routes"
)

func main() {

	database.Connect()
	database.DB.AutoMigrate(&models.Brand{}, &models.Voucher{}, &models.Transaction{}, &models.TransactionDetail{})

	r := routes.RegisterRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
