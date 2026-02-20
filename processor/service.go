package processor

import (
	"context"
	"errors"

	"github.com/kristiyanpd/interview/antimoneylaundering/transaction"
	"gorm.io/gorm"
)

var (
	ErrNilTransaction = errors.New("processor: transaction is nil")
	ErrNotImplemented = errors.New("processor: implement AML orchestration")
)

// Process should enforce limits, perform money movement atomically, and store
// the resulting transaction. Replace the stubbed implementation as you build
// out the service.
func Process(ctx context.Context, db *gorm.DB, tx *transaction.Transaction) error {
	if tx == nil {
		return ErrNilTransaction
	}

	return ErrNotImplemented
}
