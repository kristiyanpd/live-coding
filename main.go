package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/kristiyanpd/interview/antimoneylaundering/account"
	"github.com/kristiyanpd/interview/antimoneylaundering/db"
	"github.com/kristiyanpd/interview/antimoneylaundering/processor"
	"github.com/kristiyanpd/interview/antimoneylaundering/transaction"
)

const dsn = "host=localhost user=aml password=aml dbname=aml port=5432 sslmode=disable"

var (
	sampleAccounts = []*account.Account{
		{ID: 1, Balance: 1500},
		{ID: 2, Balance: 500},
	}

	sampleTransactions = []*transaction.Transaction{
		transaction.NewTransaction(350, 1, 2),  // Fail: exceeds single-transaction limit (300)
		transaction.NewTransaction(200, 99, 2), // Fail: source account 99 does not exist
		transaction.NewTransaction(200, 1, 98), // Fail: destination account 98 does not exist
		transaction.NewTransaction(800, 2, 1),  // Fail: account 2 only has 500 balance (insufficient funds)
	}
)

func main() {
	gormDB, err := db.Open(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.Migrate(gormDB)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	ctx := context.Background()

	for i, tx := range sampleTransactions {
		slog.Info("processing sample transaction",
			"index", i,
			"from", tx.FromAccountID,
			"to", tx.ToAccountID,
			"amount", tx.Amount,
		)

		err = processor.Process(ctx, gormDB, tx)
		if err != nil {
			slog.Warn("sample transaction failed (expected)", "index", i, "error", err)
			continue
		}

		slog.Info("sample transaction processed successfully (unexpected)", "index", i)
	}
}
