package entity

import (
	"errors"
	"time"
)

const (
	isAvailable   = "Available"
	isUnavailable = "Unavailable"
)

type Product struct {
	Id                 string
	Status             string
	MerchantId         string
	ProductDescription string
	Amount             int64
	Quantity           int64
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func NewProduct(id, merchantId, productDescription string, amount, quantity int64) (*Product, error) {
	now := time.Now()
	p := &Product{
		Id:                 id,
		MerchantId:         merchantId,
		ProductDescription: productDescription,
		Amount:             amount,
		Quantity:           quantity,
		Status:             isAvailable,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	err := p.Validate()

	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) Validate() error {
	var errs []error

	if p.Id == "" {
		errs = append(errs, errors.New("product_id is required"))
	}
	if p.MerchantId == "" {
		errs = append(errs, errors.New("merchant_id is required"))
	}
	if p.ProductDescription == "" {
		errs = append(errs, errors.New("product_description is required"))
	}
	if p.Amount <= 0 {
		errs = append(errs, errors.New("amount is required"))
	}
	if p.Quantity == 0 {
		errs = append(errs, errors.New("quantity is required"))
	}

	return nil
}

func (p *Product) Block() {
	now := time.Now()
	p.Status = isUnavailable
	p.UpdatedAt = now
}

func (p *Product) Unblock() {
	now := time.Now()
	p.Status = isAvailable
	p.UpdatedAt = now
}
