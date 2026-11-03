package entity

import (
	"errors"
	"time"
)

const (
	isActive  = "Active"
	isBlocked = "Blocked"
)

type Partner struct {
	Id           string
	Status       string
	Document     string
	DocumentType string
	FantasyName  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewPartner(id, document, documentType, fantasyName string) (*Partner, error) {
	now := time.Now()
	p := &Partner{
		Id:           id,
		Document:     document,
		DocumentType: documentType,
		FantasyName:  fantasyName,
		Status:       isActive,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err := p.Validate()

	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Partner) Validate() error {
	var errs []error

	if p.Id == "" {
		errs = append(errs, errors.New("partnerId is required"))
	}
	if p.Document == "" {
		errs = append(errs, errors.New("document is required"))
	}
	if p.DocumentType == "" {
		errs = append(errs, errors.New("documentType is required"))
	}
	if p.FantasyName == "" {
		errs = append(errs, errors.New("fantasyName is required"))
	}
	return errors.Join(errs...)
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
