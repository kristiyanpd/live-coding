package transaction

import (
	"time"

	"github.com/kristiyanpd/interview/antimoneylaundering/account"
)

type TransactionDAO struct {
	ID            uint64
	Amount        int64              `gorm:"not null"`
	FromAccountID uint64             `gorm:"not null"`
	FromAccount   account.AccountDAO `gorm:"foreignKey:FromAccountID"`
	ToAccountID   uint64             `gorm:"not null"`
	ToAccount     account.AccountDAO `gorm:"foreignKey:ToAccountID"`
	CreatedAt     time.Time
}

func (TransactionDAO) TableName() string { return "transactions" }

func ToDAO(t *Transaction) *TransactionDAO {
	return &TransactionDAO{
		ID:            t.ID,
		Amount:        t.Amount,
		FromAccountID: t.FromAccountID,
		ToAccountID:   t.ToAccountID,
		CreatedAt:     t.CreatedAt,
	}
}

func ToDomain(d *TransactionDAO) *Transaction {
	return &Transaction{
		ID:            d.ID,
		Amount:        d.Amount,
		FromAccountID: d.FromAccountID,
		ToAccountID:   d.ToAccountID,
		CreatedAt:     d.CreatedAt,
	}
}
