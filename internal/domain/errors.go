package domain

import "errors"

var (
	// 'E retornado quando uma conta nao e encontrada.
	ErrAccountNotFound = errors.New("Account not found.")
	// ErrDuplicatedAPIKey é retornado quando há tentativa de criar uma conta com API key duplicada
	ErrDuplicateApiKey = errors.New("Account duplicate.")
	// ErrInvoiceNotFound é retornado quando uma fatura não é encontrada.
	ErrInvoiceNotFound = errors.New("invoice not found.")
	// ErrUnauthorizedAcess é retornado quando há tentativa de acesso não autorizado a um recurso
	ErrUnauthorizedAcess = errors.New("unauthorized acess.")

	
	ErroInvalidAmount = errors.New("Invalid amount")

	ErrInvalidStatus = errors.New("invalid status")
)