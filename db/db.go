package db

import (
	"fmt"

	"github.com/kristiyanpd/interview/antimoneylaundering/account"
	"github.com/kristiyanpd/interview/antimoneylaundering/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db: open: %w", err)
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&account.AccountDAO{},
		&transaction.TransactionDAO{},
	)
}
