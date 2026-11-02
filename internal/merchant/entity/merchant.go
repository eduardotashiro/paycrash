package entity

import (
	"errors"
	"time"
)

const (
	isActive  = "ACTIVE"
	isBlocked = "BLOCKED"
)

type Merchant struct {
	Id             string
	PartnerId      string
	DocumentNumber string
	DocumentType   string
	NameFantasy    string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewMerchant(id, partnerId, documentNumber, documentType, nameFantasy string) (*Merchant, error) {
	now := time.Now()
	m := &Merchant{
		Id:             id,
		PartnerId:      partnerId,
		DocumentNumber: documentNumber,
		DocumentType:   documentType,
		NameFantasy:    nameFantasy,
		Status:         isActive,
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
	if m.Id == "" {
		return errors.New("merchant_id is required")
	}
	if m.PartnerId == "" {
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

	return nil
}

func (m *Merchant) Block() {
	now := time.Now()
	m.Status = isBlocked
	m.UpdatedAt = now
}

func (p *Merchant) Unblock() {
	now := time.Now()
	p.Status = isActive
	p.UpdatedAt = now
}
