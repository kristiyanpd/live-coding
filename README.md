## Anti–Money-Laundering Coding Exercise

You are building a very small anti–money-laundering (AML) engine for a fictional payments company. The goal is to design a simple, well-structured service that can ingest transfer requests, apply limits, and persist the results.

⏰ **Estimated duration: 1.5 hours**

### What’s already in place

- Domain models for `Account` and `Transaction`
- Limit configuration defaults in `limit/config.go`
- Docker Compose stack with PostgreSQL 18
- Auto-migration via `db/init.sql`

### Your primary tasks

1. **Database layer**  
   - The application already connects to PostgreSQL (running via Docker) and auto-migrates the schema.  

2. **AML limits**  
   Enforce the following for outgoing transfers (per `FromAccountID`):
   - Maximum single transaction amount: `300`
   - Maximum number of transactions per day: `10`
   - Maximum amount sent per day: `5000`
   - Maximum number of transactions per month: `50`
   - Maximum amount sent per month: `15000`

3. **Concurrent processing** *(discuss or implement as time allows)*  
   - Design a worker/concurrency pipeline that processes requests with multiple workers.  
   - Be prepared to discuss how jobs are queued, how errors propagate, and back-pressure strategy — even if you only sketch the design.

### Stretch ideas (if time allows)

- Implement the concurrency pipeline end-to-end.
- Add structured logging, metrics, or tracing for AML rejections.
- Build automated tests covering limit enforcement and data integrity under concurrent load.

### Getting started

```bash
# Spin up Postgres (default creds: aml/aml)
docker compose up -d db

# Run tests / lint (add more as you build)
go test ./...

# Launch the sample main program
go run ./...
```

### Repository layout

- `account/` – domain types 
- `transaction/` – domain types 
- `limit/` – limit configuration helpers
- `processor/` – orchestration logic
- `db/` – database connection and auto-migration

### What we look for

- **Atomic, correct money movement** — transfers must not leave balances inconsistent
- **Sound AML limit enforcement** — clean design that's easy to extend with new rules
- Clear, testable design with well-separated layers
- Thoughtful reasoning about concurrency, error handling, and observability

### Need interviewer-only material?

See `README.interviewer.md` for scoring guidelines, hints, and runbook notes. Please stop reading there if you are a candidate. Good luck! 👊
