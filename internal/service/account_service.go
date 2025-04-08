/*
* service = Ele permitira que os metodos necessarios para criar nova conta peguem os dados passado pelo usuario
 */

package service

import (
	"github.com/RFreitasAnjos/imersao22FC/go-gateway/internal/domain"
	"github.com/RFreitasAnjos/imersao22FC/go-gateway/internal/dto"
)

/*
 * Recebe os dados do banco de dados do repositorio
 */
type AccountService struct {
	repository domain.AccountRepository
}

/*
 * Funcao que constrora que recebera um novo AccountService
 */

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount (input dto.CreateAccountInput) (*dto.AccountOutput, error){
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIkey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil{
		return nil, domain.ErrDuplicateApiKey
	}

	err = s.repository.Save(account)
	if err != nil{
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error){
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil{
		return nil,err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func( s*AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error){
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}