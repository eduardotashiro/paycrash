package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eduardotashiro/paycrash/internal/merchant/entity"
)

type createMerchantRequest struct {
	PartnerId    string `json:"partnerId"`
	Document     string `json:"document"`
	DocumentType string `json:"documentType"`
	FantasyName  string `json:"fantasyName"`
}

type merchantResponse struct {
	Code string           `json:"code"`
	Data *entity.Merchant `json:"data"`
}

type errorResponse struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}

func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var req createMerchantRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	id := "42"
	merchant, err := entity.NewMerchant(id, req.PartnerId, req.Document, req.DocumentType, req.FantasyName)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(merchantResponse{
		Code: http.StatusText(http.StatusCreated),
		Data: merchant,
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
