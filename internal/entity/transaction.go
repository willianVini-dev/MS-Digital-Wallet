package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	AccountFrom *Account  `json:"account_from"`
	AccountTo   *Account  `json:"account_to"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		Amount:      amount,
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		CreatedAt:   time.Now(),
	}

	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	transaction.Commit()

	return transaction, nil
}

func (t *Transaction) Validate() error {

	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds in the source account")
	}

	return nil
}

func (t *Transaction) Commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}
