package service

import (
	"fmt"

	"github.com/cheemx5395/personal-expense-tracker/src/repository"
)

type Config struct {
	DB *repository.Queries
}

func Load() (*Config, error) {
	dbQueries, err := repository.OpenDB()
	if err != nil {
		return &Config{}, err
	}

	fmt.Printf("DB Connected successfully: %v\n", dbQueries)

	cfg := &Config{DB: &dbQueries}
	return cfg, nil
}
