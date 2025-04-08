package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RFreitasAnjos/imersao22FC/go-gateway/internal/dto"
	"github.com/RFreitasAnjos/imersao22FC/go-gateway/internal/service"
)

type AccountHanlder struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHanlder{
	return &AccountHanlder{accountService: accountService}
}

func (h *AccountHanlder) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateAccountInput 
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(),http.StatusBadRequest)
		return
	}

	output, err := h.accountService.CreateAccount(input)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *AccountHanlder) Get(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w,"x-API-Key is required", http.StatusUnauthorized)
		return
	}
	output, err := h.accountService.FindByAPIKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}