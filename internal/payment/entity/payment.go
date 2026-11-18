package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	isAuthorized               = "AUTHORIZED"
	isCaptured                 = "CAPTURED"
	isPartialAmountCancelation = "PARTIAL_CANCELED"
)

type PaymentItem struct {
	Id         uuid.UUID
	ProductId  uuid.UUID
	MerchantId uuid.UUID
	Amount     int64
	Quantity   int64
}

// pagamento vem do merchant nao do partner
type Payment struct {
	Id          uuid.UUID
	PartnerId   uuid.UUID
	TotalAmount int64
	Status      string
	Items       []PaymentItem
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewPayment(partnerId uuid.UUID, totalAmount int64, items []PaymentItem) (*Payment, error) {
	now := time.Now()
	uuid := uuid.New()
	payment := &Payment{
		Id:          uuid,
		PartnerId:   partnerId,
		TotalAmount: totalAmount,
		Status:      isAuthorized,
		Items:       items,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := payment.Validate()

	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (p *Payment) Validate() error {
	var errs []error

	if p.Id == uuid.Nil {
		errs = append(errs, errors.New("id is required"))
	}
	if p.PartnerId == uuid.Nil {
		errs = append(errs, errors.New("partner_id is required"))
	}
	if p.TotalAmount <= 0 {
		errs = append(errs, errors.New("totalAmount is required"))
	}
	if len(p.Items) == 0 {
		errs = append(errs, errors.New("items is required"))
	}

	return errors.Join(errs...)
}

// 2. Validações básicas:
//    - O Id não pode ser uuid.Nil
//    - O PartnerId não pode ser uuid.Nil
//    - A lista de Items não pode estar vazia (len(p.Items) == 0)

// 3. A Matemática (O Chefão):
//    - Crie uma variável: var somaCalculada int64 = 0
//    - Faça um 'for _, item := range p.Items'
//    - Dentro do for, acumule: somaCalculada += (item.Amount * item.Quantity)
//    - Importante: aproveite o for para validar se o item.Amount <= 0 ou item.Quantity <= 0 e adicionar erro se for.

// 4. O Checkmate:
//    - Fora do for, faça um if comparando se a 'somaCalculada' é diferente do 'p.TotalAmount'
//    - Se for diferente, adicione um erro dizendo que os valores não batem.

// 5. Retorne errors.Join(errs...)
