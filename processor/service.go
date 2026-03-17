package processor

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kristiyanpd/interview/antimoneylaundering/transaction"
)

var (
	ErrNilTransaction = errors.New("processor: transaction is nil")
	ErrNotImplemented = errors.New("processor: implement AML orchestration")
)

// Process should enforce limits, perform money movement atomically, and store
// the resulting transaction. Replace the stubbed implementation as you build
// out the service.
func Process(ctx context.Context, db *sql.DB, tx *transaction.Transaction) error {
	if tx == nil {
		return ErrNilTransaction
	}

	return ErrNotImplemented
}
