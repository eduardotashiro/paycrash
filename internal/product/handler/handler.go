package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/eduardotashiro/paycrash/internal/product/entity"

	"github.com/google/uuid"
)

type createProductRequest struct {
	MerchantId         uuid.UUID `json:"merchantId"`
	ProductDescription string    `json:"productDescription"`
	Amount             int64     `json:"amount"`
	Quantity           int64     `json:"itemQuantity"`
}

type ProductResponse struct {
	Code string              `json:"code"`
	Data productResponseData `json:"data"`
}

type productResponseData struct {
	ID                 uuid.UUID `json:"id"`
	Status             string    `json:"status"`
	MerchantId         uuid.UUID `json:"merchantId"`
	ProductDescription string    `json:"productDescription"`
	Amount             int64     `json:"amount"`
	Quantity           int64     `json:"quantity"`
	CreatedAt          time.Time `json:"createdAt"`
}

type errorResponse struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}

// TODO: validar se o seller existe antes de criar o Product..
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req createProductRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	id := uuid.New()

	product, err := entity.NewProduct(id, req.MerchantId, req.ProductDescription, req.Amount, req.Quantity)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	responseData := productResponseData{
		ID:                 product.Id,
		MerchantId:         product.MerchantId,
		ProductDescription: product.ProductDescription,
		Amount:             product.Amount,
		Quantity:           product.Quantity,
		Status:             product.Status,
		CreatedAt:          product.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ProductResponse{
		Code: http.StatusText(http.StatusCreated),
		Data: responseData,
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
