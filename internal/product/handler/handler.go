package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eduardotashiro/paycrash/internal/product/entity"
)

type createProductRequest struct {
	MerchantId         string `json:"merchantId"`
	ProductDescription string `json:"productDescription"`
	Amount             int64  `json:"amount"`
	Quantity           int64  `json:"itemQuantity"`
}

type partnerResponse struct {
	Code string          `json:"code"`
	Data *entity.Product `json:"data"`
}

type errorResponse struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req createProductRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	id := "42"
	product, err := entity.NewProduct(id, req.MerchantId, req.ProductDescription, req.Amount, req.Quantity)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(partnerResponse{
		Code: http.StatusText(http.StatusCreated),
		Data: product,
	})
}

func writeError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse{
		Code:   http.StatusText(status),
		Errors: strings.Split(err.Error(), "\n"),
	})
}
