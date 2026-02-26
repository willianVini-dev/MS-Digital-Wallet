package createaccount

import (
	"wv-ms-wallet/internal/entity"
	"wv-ms-wallet/internal/gateway"
)

type CreateAccountInput struct {
	ClientID string
}

type CreateAccountOutput struct {
	ID string
}

type CreateAccountUseCase struct {
	accountGateway gateway.AccountGateway
	clientGateway  gateway.ClientGateway
}

func NewAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountGateway: a,
		clientGateway:  c,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInput) (*CreateAccountOutput, error) {

	// Verificar se o cliente existe
	client, err := uc.clientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	// Criar a conta
	account := entity.NewAccount(client)
	err = uc.accountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutput{ID: account.ID}, nil
}
