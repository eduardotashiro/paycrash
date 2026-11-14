package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewPartner(t *testing.T) {
	testes := []struct {
		testName        string
		id              uuid.UUID
		document        string
		documentType    string
		fantasyName     string
		waitingTheError bool
	}{
		{testName: "partner created successfully", document: "123.456.789-99", documentType: "cpf", fantasyName: "bar dos coroa", waitingTheError: false},
		{testName: "error creating the partner", documentType: "fpc", fantasyName: "coroa dos bar", waitingTheError: true},
	}

	for _, tt := range testes {
		t.Run(tt.testName, func(t *testing.T) {
			partner, err := NewPartner(tt.id, tt.document, tt.documentType, tt.fantasyName)

			if tt.waitingTheError && err == nil {
				t.Errorf("esperava um erro, mas não retornou nothing")
			}

			if !tt.waitingTheError && err != nil {
				t.Errorf("não esperava erro, but retornou: %v", err)
			}

			if !tt.waitingTheError && partner != nil {
				if partner.Document != tt.document {
					t.Errorf("esperava o DOC %s, mas he came %s", tt.id, partner.Id)
				}

				if partner.FantasyName != tt.fantasyName {
					t.Errorf("esperava o name %s, mas veio %s", tt.fantasyName, partner.FantasyName)
				}
				if partner.Status != "Active" {
					t.Errorf("expected status Active, but veio %s", partner.Status)
				}
			}

		})
	}
}
