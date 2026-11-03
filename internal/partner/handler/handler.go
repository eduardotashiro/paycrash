package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eduardotashiro/paycrash/internal/partner/entity"
)

type createPartnerRequest struct {
	Document     string `json:"document"`
	DocumentType string `json:"documentType"`
	FantasyName  string `json:"fantasyName"`
}

type partnerResponse struct {
	Code string          `json:"code"`
	Data *entity.Partner `json:"data"`
}

type errorResponse struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}

func CreatePartner(w http.ResponseWriter, r *http.Request) {
	var req createPartnerRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	id := "42"
	partner, err := entity.NewPartner(id, req.Document, req.DocumentType, req.FantasyName)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(partnerResponse{
		Code: http.StatusText(http.StatusCreated),
		Data: partner,
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
