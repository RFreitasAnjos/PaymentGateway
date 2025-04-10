package domain 

import "time"

type Status string 

const (
	StatusPending Status = "pending"
	StatusApproved Status = "aprroved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID string
	AccountID string
	Amount float64
	Status string
	Description string
	PaymenteType string
	CardLastDigits string
	CreatedAtt time.Time
	UpdatedAt time.Time
}

type CreditCard struct {
	Number string
	CVV string
	ExpiryMonth int
	ExpiryYear int
	CardholderName string
}

func NewInvoice(accountID string, amount float64, description string, paymenteType string, card CreditCard) (*Invoice, error){
	if amount <= 0 {
		return nil, ErroInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:]
} 