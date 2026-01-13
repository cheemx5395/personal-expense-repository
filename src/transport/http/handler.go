package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cheemx5395/personal-expense-tracker/src/repository"
	"github.com/cheemx5395/personal-expense-tracker/src/service"
	"github.com/gorilla/mux"
)

// WriteHeader must be called before Write when writing response

func HealthRoute(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := struct {
			Message string `json:"message"`
		}{
			Message: "Healthy Connection Established!",
		}

		body, _ := json.Marshal(res)

		w.WriteHeader(200)
		w.Write(body)
	}
}

func CreateExpense(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req repository.CreateExpenseParams
		incomingReq := struct {
			Title       string `json:"title"`
			Amount      int64  `json:"amount"`
			Description string `json:"description"`
		}{}
		res := struct {
			Error string `json:"error"`
		}{}

		err := json.NewDecoder(r.Body).Decode(&incomingReq)
		if err != nil {
			res.Error = "Bad Request"
			log.Printf("Error creating expense in REST: %v", err)
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
			return
		}

		req.Amount = incomingReq.Amount
		req.Title = sql.NullString{String: incomingReq.Title, Valid: true}
		req.Description = sql.NullString{String: incomingReq.Description, Valid: true}

		exp, err := cfg.CreateExpense(req)
		if err != nil {
			res.Error = "Internal Server Error"
			log.Printf("Error creating expense in REST: %v\n", err)
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
			return
		}

		body, _ := json.Marshal(exp)
		w.WriteHeader(http.StatusCreated)
		w.Write(body)
	}
}

func GetExpenses(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := struct {
			Error string `json:"error"`
		}{}

		expenses, err := cfg.GetExpenses()
		if err != nil {
			res.Error = "Internal Server Error"
			log.Printf("Error getting expenses from DB: %v\n", err)
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
			return
		}

		body, _ := json.Marshal(expenses)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func DeleteExpense(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		res := struct {
			Error string `json:"error"`
		}{}
		if idStr == "" {
			res.Error = "Bad Request"
			log.Println("No id parameter available in path")
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			res.Error = "Internal Server Error "
			log.Println("Error parsing the id into int64")
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
			return
		}

		err = cfg.DeleteExpense(id)
		if err != nil {
			res.Error = "Internal Server Error"
			log.Printf("Error deleting expense in DB: %v\n", err)
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateExpense(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		res := struct {
			Error string `json:"error"`
		}{}
		if idStr == "" {
			res.Error = "Bad Request"
			log.Println("Error updating the expense, no id parameter available in path")
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			res.Error = "Internal Server Error "
			log.Println("Error parsing the id into int64")
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
			return
		}

		req := struct {
			Title       string `json:"title"`
			Amount      int64  `json:"amount"`
			Description string `json:"description"`
		}{}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			res.Error = "Internal Server Error"
			log.Printf("Error upating the expense %v\n", err)
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
			return
		}

		var param repository.UpdateExpenseByIDParams
		param.Amount = req.Amount
		param.Description.String = req.Description
		param.Description.Valid = true
		param.Title.String = req.Title
		param.Title.Valid = true
		param.ID = id

		updatedExpense, err := cfg.UpdateExpense(param)
		if err != nil {
			res.Error = "Internal Server Error"
			log.Printf("Error updating expense in DB: %v\n", err)
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
			return
		}

		body, _ := json.Marshal(updatedExpense)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func GetExpense(cfg *service.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		res := struct {
			Error string `json:"error"`
		}{}
		if idStr == "" {
			res.Error = "Bad Request"
			log.Printf("Error getting the expense, no id parameter available in path- idStr: %s\n", idStr)
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			res.Error = "Internal Server Error"
			log.Println("Error parsing the id into int64")
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
			return
		}

		expense, err := cfg.GetExpense(id)
		if err != nil {
			res.Error = "Internal Server Error"
			log.Printf("Error getting expense from DB: %v\n", err)
			body, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
			return
		}

		body, _ := json.Marshal(expense)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}
