package account

import (
	"time"
)

type AccountDAO struct {
	ID        uint64
	Balance   int64 `gorm:"not null"`
	CreatedAt time.Time
}

func (AccountDAO) TableName() string { return "accounts" }

func ToDAO(a *Account) *AccountDAO {
	return &AccountDAO{
		ID:        a.ID,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
	}
}

func ToDomain(d *AccountDAO) *Account {
	return &Account{
		ID:        d.ID,
		Balance:   d.Balance,
		CreatedAt: d.CreatedAt,
	}
}
