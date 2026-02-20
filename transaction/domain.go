package transaction

import "time"

type Transaction struct {
	ID            uint64
	Amount        int64
	FromAccountID uint64
	ToAccountID   uint64
	CreatedAt     time.Time
}

func NewTransaction(amount int64, fromAccountID uint64, toAccountID uint64) *Transaction {
	return &Transaction{
		Amount:        amount,
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		CreatedAt:     time.Now().UTC(),
	}
}
