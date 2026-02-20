package account

import "time"

type Account struct {
	ID        uint64
	Balance   int64
	CreatedAt time.Time
}
