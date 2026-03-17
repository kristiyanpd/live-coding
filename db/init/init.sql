CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL PRIMARY KEY,
    balance BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    amount BIGINT NOT NULL,
    from_account_id BIGINT NOT NULL REFERENCES accounts (id),
    to_account_id BIGINT NOT NULL REFERENCES accounts (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_transactions_from_account_created_at
    ON transactions (from_account_id, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_transactions_to_account_created_at
    ON transactions (to_account_id, created_at DESC);

INSERT INTO accounts (balance) VALUES
    (1500),
    (500);
