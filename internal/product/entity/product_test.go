package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewProduct(t *testing.T) {
	testes := []struct {
		testName           string
		id                 uuid.UUID
		merchantId         uuid.UUID
		productDescription string
		amount             int64
		quantity           int64
		waitingTheError    bool
	}{
		{testName: "product created successfully", id: uuid.New(), merchantId: uuid.New(), productDescription: "Jhonsons Baby 1L", amount: 1322, quantity: 42, waitingTheError: false},
		{testName: "error creating the product", id: uuid.New(), merchantId: uuid.New(), productDescription: "Sabonete de enxofre", amount: 0, quantity: 0, waitingTheError: true},
	}

	for _, tt := range testes {
		t.Run(tt.testName, func(t *testing.T) {
			product, err := NewProduct(tt.id, tt.merchantId, tt.productDescription, tt.amount, tt.quantity)

			if tt.waitingTheError && err == nil {
				t.Errorf("esperava um erro, mas não retornou nothing")
			}

			if !tt.waitingTheError && err != nil {
				t.Errorf("não esperava erro, but retornou: %v", err)
			}

			if !tt.waitingTheError && product != nil {
				if product.Id != tt.id {
					t.Errorf("esperava o ID %s, mas he came %s", tt.id, product.Id)
				}

				if product.Amount != tt.amount {
					t.Errorf("esperava o name %d, mas veio %d", tt.amount, product.Amount)
				}
				if product.Status != "Available" {
					t.Errorf("expected status Active, mas veio %s", product.Status)
				}
			}

		})
	}
}
