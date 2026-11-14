package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/eduardotashiro/paycrash/internal/merchant/entity"
	"github.com/google/uuid"
)

type createMerchantRequest struct {
	Document     string `json:"document"`
	DocumentType string `json:"documentType"`
	FantasyName  string `json:"fantasyName"`
}

type merchantResponse struct {
	Code string               `json:"code"`
	Data merchantResponseData `json:"data"`
}

type merchantResponseData struct {
	ID           uuid.UUID `json:"id"`
	Document     string    `json:"document"`
	DocumentType string    `json:"documentType"`
	FantasyName  string    `json:"fantasyName"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
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

	id := uuid.New()

	merchant, err := entity.NewMerchant(id, req.Document, req.DocumentType, req.FantasyName)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	responseData := merchantResponseData{
		ID:           merchant.Id,
		Document:     merchant.Document,
		DocumentType: merchant.DocumentType,
		FantasyName:  merchant.FantasyName,
		Status:       merchant.Status,
		CreatedAt:    merchant.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(merchantResponse{
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
