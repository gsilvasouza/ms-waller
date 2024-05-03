package create_account

import (
	"fmt"

	"github.com/gsilvasouza/ms-waller/internal/entity"
	"github.com/gsilvasouza/ms-waller/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(clientGateway gateway.ClientGateway, accountGateway gateway.AccountGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		ClientGateway:  clientGateway,
		AccountGateway: accountGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	fmt.Println(input.ClientID)
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
	fmt.Println(client)
	account := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
