package http

import (
	"encoding/json"
	"net/http"

	"github.com/cheemx5395/personal-expense-tracker/src/service"
)

func HealthRoute(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := struct {
			Message string `json:"message"`
		}{
			Message: "Healthy Connection Established!",
		}

		body, _ := json.Marshal(res)

		w.Write(body)
		w.WriteHeader(200)
	}
}

func CreateExpense(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetExpenses(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func DeleteExpense(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
