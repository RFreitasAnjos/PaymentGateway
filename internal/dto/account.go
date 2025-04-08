/*
 * Ira receber o Json e trata para encaminhar para o repository
 */

package dto

import (
	"time"

	"github.com/RFreitasAnjos/imersao22FC/go-gateway/internal/domain"
)

// 
type CreateAccountInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Retorna dados de uma conta
type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `jsono:"updated_at"`
}

// 
func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) AccountOutput {
	return AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		APIKey:    account.APIkey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
