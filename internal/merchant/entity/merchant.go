package entity

import (
	"errors"
	"time"
)

const (
	statusActive  = "ACTIVE"
	statusBlocked = "BLOCKED"
)

type Merchant struct {
	ID             string
	PartnerID      string
	DocumentNumber string
	DocumentType   string
	NameFantasy    string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewMerchant(id, partnerID, documentNumber, documentType, nameFantasy string) (*Merchant, error) {
	now := time.Now()
	m := &Merchant{
		ID:             id,
		PartnerID:      partnerID,
		DocumentNumber: documentNumber,
		DocumentType:   documentType,
		NameFantasy:    nameFantasy,
		Status:         statusActive,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	err := m.Validate()

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Merchant) Validate() error {
	if m.ID == "" {
		return errors.New("merchant_id is required")
	}
	if m.PartnerID == "" {
		return errors.New("partner_id is required")
	}
	if m.DocumentNumber == "" {
		return errors.New("document_number is required")
	}
	if m.DocumentType == "" {
		return errors.New("document_type is required")
	}
	if m.NameFantasy == "" {
		return errors.New("name_fantasy is required")
	}

	if m.Status == statusBlocked {
		return errors.New("merchant_blocked")
	}
	return nil
}

func (m *Merchant) Block() {
	now := time.Now()
	m.Status = statusBlocked
	m.UpdatedAt = now
}
