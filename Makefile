server:
	@go build -o personal-expense-tracker-rest ./cmd/server && ./personal-expense-tracker-rest

cli:
	@go build -o personal-expense-tracker-cli ./cmd/cli && ./personal-expense-tracker-cli

migrationUp:
	@goose -dir src/repository/migrations sqlite3 expenses.db up

migrationDown:
	@goose -dir src/repository/migrations sqlite3 expenses.db down