package limit

// Config controls AML guardrails.
type Config struct {
	MaxTransactionAmount   int64
	MaxDailyTransactions   int
	MaxDailyAmount         int64
	MaxMonthlyTransactions int
	MaxMonthlyAmount       int64
}

func DefaultConfig() Config {
	return Config{
		MaxTransactionAmount:   300,
		MaxDailyTransactions:   10,
		MaxDailyAmount:         5_000,
		MaxMonthlyTransactions: 50,
		MaxMonthlyAmount:       15_000,
	}
}
