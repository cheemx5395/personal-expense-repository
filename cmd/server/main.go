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

	r.HandleFunc("/health", rest.HealthRoute(cfg)).Methods("GET")
	r.HandleFunc("/expenses", rest.GetExpenses(cfg)).Methods("GET")
	r.HandleFunc("/expense", rest.CreateExpense(cfg)).Methods("POST")
	r.HandleFunc("/expense/{id}", rest.DeleteExpense(cfg)).Methods("DELETE")

	log.Println("Server is running on port :8080")
	log.Fatalf("Server crashed: %v", http.ListenAndServe(":8080", r))
}
