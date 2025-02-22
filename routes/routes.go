package routes

import (
	"go-api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/brand", handlers.CreateBrand).Methods(http.MethodPost)
	r.HandleFunc("/voucher", handlers.CreateVoucher).Methods(http.MethodPost)
	r.HandleFunc("/voucher", handlers.GetVoucher).Methods(http.MethodGet)
	r.HandleFunc("/voucher/brand", handlers.GetVouchersByBrand).Methods(http.MethodGet)
	r.HandleFunc("/transaction/redemption", handlers.MakeRedemption).Methods(http.MethodPost)
	r.HandleFunc("/transaction/redemption", handlers.GetTransactionDetail).Methods(http.MethodGet)

	return r
}
