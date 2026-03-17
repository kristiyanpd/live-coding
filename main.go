package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/kristiyanpd/interview/antimoneylaundering/db"
	"github.com/kristiyanpd/interview/antimoneylaundering/processor"
	"github.com/kristiyanpd/interview/antimoneylaundering/transaction"
)

const dsn = "postgres://aml:aml@localhost:5432/aml?sslmode=disable"

func main() {
	database, err := db.Open(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	ctx := context.Background()

	tx := transaction.NewTransaction(100, 1, 2)

	err = processor.Process(ctx, database, tx)
	if err != nil {
		slog.Error("sample transaction failed", "error", err)
	}
}
