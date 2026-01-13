package service

import (
	"context"
	"fmt"

	"github.com/cheemx5395/personal-expense-tracker/src/repository"
)

type Config struct {
	DB  *repository.Queries
	Ctx context.Context
}

func Load() (*Config, error) {
	dbQueries, err := repository.OpenDB()
	if err != nil {
		return &Config{}, err
	}

	fmt.Printf("DB Connected successfully: %v\n", dbQueries)

	cfg := &Config{
		DB:  &dbQueries,
		Ctx: context.Background(),
	}
	return cfg, nil
}

func (cfg *Config) CreateExpense(expense repository.CreateExpenseParams) (repository.Expense, error) {
	exp, err := cfg.DB.CreateExpense(cfg.Ctx, expense)
	if err != nil {
		return repository.Expense{}, err
	}
	return exp, nil
}

func (cfg *Config) GetExpense(id int64) (repository.Expense, error) {
	exp, err := cfg.DB.GetExpenseByID(cfg.Ctx, id)
	if err != nil {
		return repository.Expense{}, err
	}
	return exp, nil
}

func (cfg *Config) GetExpenses() ([]repository.Expense, error) {
	res, err := cfg.DB.GetExpenses(cfg.Ctx)
	if err != nil {
		return []repository.Expense{}, err
	}
	return res, nil
}

func (cfg *Config) UpdateExpense(params repository.UpdateExpenseByIDParams) (repository.Expense, error) {
	expense, err := cfg.DB.UpdateExpenseByID(cfg.Ctx, params)
	if err != nil {
		return repository.Expense{}, err
	}
	return expense, nil
}

func (cfg *Config) DeleteExpense(id int64) error {
	return cfg.DB.DeleteExpenseByID(cfg.Ctx, id)
}
