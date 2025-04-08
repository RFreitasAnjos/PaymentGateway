package domain

import "errors"

var (
	// 'E retornado quando uma conta nao e encontrada.
	ErrAccountNotFound = errors.New("Account not found.")
	ErrDuplicateApiKey = errors.New("Account duplicate.")
)