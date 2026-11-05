package entity

import (
	"testing"
)

func TestNewPartner(t *testing.T) {
	testes := []struct {
		testName        string
		id              string
		document        string
		documentType    string
		fantasyName     string
		waitingTheError bool
	}{
		{testName: "partner created successfully", id: "uuid", document: "123.456.789-99", documentType: "cpf", fantasyName: "bar dos coroa", waitingTheError: false},
		{testName: "error creating the partner", document: "999.876.543-21", documentType: "fpc", fantasyName: "coroa dos bar", waitingTheError: true},
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
				if partner.Id != tt.id {
					t.Errorf("esperava o ID %s, mas he came %s", tt.id, partner.Id)
				}

				if partner.FantasyName != tt.fantasyName {
					t.Errorf("esperava o name %s, mas veio %s", tt.fantasyName, partner.FantasyName)
				}
				if partner.Status != "Active" {
					t.Errorf("expected status Active, mas veio %s", partner.Status)
				}
			}

		})
	}
}
