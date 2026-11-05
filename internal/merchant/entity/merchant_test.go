package entity

import (
	"testing"
)

func TestNewMerchant(t *testing.T) {
	testes := []struct {
		testName        string
		id              string
		partnerId       string
		document        string
		documentType    string
		fantasyName     string
		waitingTheError bool
	}{
		{testName: "merchant created successfully", id: "uuid", partnerId: "idDoPartner", document: "123.456.789-99", documentType: "cpf", fantasyName: "bar dos coroa", waitingTheError: false},
		{testName: "error creating the merchant", partnerId: "PartnerDoId", document: "999.876.543-21", documentType: "fpc", fantasyName: "coroa dos bar", waitingTheError: true},
	}

	for _, tt := range testes {
		t.Run(tt.testName, func(t *testing.T) {
			merchant, err := NewMerchant(tt.id, tt.partnerId, tt.document, tt.documentType, tt.fantasyName)

			if tt.waitingTheError && err == nil {
				t.Errorf("esperava um erro, mas não retornou nothing")
			}

			if !tt.waitingTheError && err != nil {
				t.Errorf("não esperava erro, but retornou: %v", err)
			}

			if !tt.waitingTheError && merchant != nil {
				if merchant.Id != tt.id {
					t.Errorf("esperava o ID %s, mas he came %s", tt.id, merchant.Id)
				}

				if merchant.FantasyName != tt.fantasyName {
					t.Errorf("esperava o name %s, mas veio %s", tt.fantasyName, merchant.FantasyName)
				}
				if merchant.Status != "Active" {
					t.Errorf("expected status Active, mas veio %s", merchant.Status)
				}
			}

		})
	}
}
