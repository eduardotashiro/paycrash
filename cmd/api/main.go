package main

import (
	"net/http"

	merchantHandler "github.com/eduardotashiro/paycrash/internal/merchant/handler"
	partnerHandler "github.com/eduardotashiro/paycrash/internal/partner/handler"
	productHandler "github.com/eduardotashiro/paycrash/internal/product/handler"
)

func main() {

	http.HandleFunc("POST /merchant", merchantHandler.CreateMerchant)
	http.HandleFunc("POST /partner", partnerHandler.CreatePartner)
	http.HandleFunc("POST /partner", productHandler.CreateProduct)

	http.ListenAndServe(":8090", nil)
}
