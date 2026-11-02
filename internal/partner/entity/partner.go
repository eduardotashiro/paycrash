package entity

import (
	"errors"
	"time"
)

const (
	isActive  = "ACTIVE"
	isBlocked = "BLOCKED"
)

type Partner struct {
	Id             string
	Status         string
	DocumentNumber string
	DocumentType   string
	NameFantasy    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewPartner(id, documentNumber, documentType, nameFantasy string) (*Partner, error) {
	now := time.Now()
	p := &Partner{
		Id:             id,
		DocumentNumber: documentNumber,
		DocumentType:   documentType,
		NameFantasy:    nameFantasy,
		Status:         isActive,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	err := p.Validate()

	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Partner) Validate() error {
	if p.Id == "" {
		return errors.New("partner_id is required")
	}
	if p.DocumentNumber == "" {
		return errors.New("document_number is required")
	}
	if p.DocumentType == "" {
		return errors.New("document_type is required")
	}
	if p.NameFantasy == "" {
		return errors.New("name_fantasy is required")
	}

	return nil
}

func (p *Partner) Block() {
	now := time.Now()
	p.Status = isBlocked
	p.UpdatedAt = now
}

func (p *Partner) Unblock() {
	now := time.Now()
	p.Status = isActive
	p.UpdatedAt = now
}
