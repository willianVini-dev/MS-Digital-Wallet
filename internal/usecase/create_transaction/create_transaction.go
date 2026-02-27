package createtransaction

import (
	"wv-ms-wallet/internal/entity"
	"wv-ms-wallet/internal/gateway"
)

type CreateTransactionInput struct {
	AccountFromID string
	AccountToID   string
	Amount        float64
}

type CreateTransactionOutput struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInput) (*CreateTransactionOutput, error) {

	accountFrom, err := uc.AccountGateway.FindById(input.AccountFromID)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.FindById(input.AccountToID)
	if err != nil {
		return nil, err
	}

	// CRIANDO A TRANSAÇÃO
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	// SALVANDO A TRANSAÇÃO
	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutput{
		ID: transaction.ID,
	}, nil

}
