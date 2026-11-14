package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	isActive  = "Active"
	isBlocked = "Blocked"
)

type Merchant struct {
	Id           uuid.UUID
	Document     string
	DocumentType string
	FantasyName  string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewMerchant(id uuid.UUID, document, documentType, fantasyName string) (*Merchant, error) {
	now := time.Now()
	m := &Merchant{
		Id:           id,
		Document:     document,
		DocumentType: documentType,
		FantasyName:  fantasyName,
		Status:       isActive,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err := m.Validate()

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Merchant) Validate() error {
	var errs []error

	if m.Document == "" {
		errs = append(errs, errors.New("document is required"))
	}
	if m.DocumentType == "" {
		errs = append(errs, errors.New("documentType is required"))
	}
	if m.FantasyName == "" {
		errs = append(errs, errors.New("fantasyName is required"))
	}
	return errors.Join(errs...)
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
