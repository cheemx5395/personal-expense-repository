run:
	@go build -o personal-expense-tracker ./cmd/server && ./personal-expense-tracker

migrationUp:
	@goose -dir src/repository/migrations sqlite3 expenses.db up