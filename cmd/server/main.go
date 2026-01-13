package main

import (
	"log"
	"net/http"

	"github.com/cheemx5395/personal-expense-tracker/src/service"
	rest "github.com/cheemx5395/personal-expense-tracker/src/transport/http"
	"github.com/cheemx5395/personal-expense-tracker/src/transport/middleware"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := service.Load()
	if err != nil {
		log.Printf("Error connecting with the Database: %v", err)
	}
	r := &mux.Router{}
	r.Use(middleware.LoggerMiddleware)

	r.HandleFunc("/health", rest.HealthRoute(cfg)).Methods(http.MethodGet)
	r.HandleFunc("/expenses", rest.GetExpenses(cfg)).Methods(http.MethodGet)
	r.HandleFunc("/expense", rest.CreateExpense(cfg)).Methods(http.MethodPost)
	r.HandleFunc("/expense/{id}", rest.GetExpense(cfg)).Methods(http.MethodGet)
	r.HandleFunc("/expense/{id}", rest.UpdateExpense(cfg)).Methods(http.MethodPut)
	r.HandleFunc("/expense/{id}", rest.DeleteExpense(cfg)).Methods(http.MethodDelete)

	port := ":8080"
	log.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server Failed: %v", err)
	}
}
