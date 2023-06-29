package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/usecase"
)

type TransactionController struct {
	usecases *usecase.TransactionUsecase
}

func NewTransactionController(usecases *usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{
		usecases: usecases,
	}
}


// Process a new Credit Cart Transaction
func (c *TransactionController) Process(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.process")

	var tr dto.TransactionCreateDto
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	transaction, err := c.usecases.Create(&tr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&transaction)
}


// Return a specific Transaction
func (c *TransactionController) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.get")

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tr, err := c.usecases.Find(id)
	if tr == nil || err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tr)
}


// Update the value of the Transaction
// No validation need but transaction ID must exists
func (c *TransactionController) Modify(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.modify")

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tr dto.TransactionUpdateDto
	err := json.NewDecoder(r.Body).Decode(&tr)
	if err != nil {
		json.NewEncoder(w).Encode("error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tr.ID = id
	err = c.usecases.Update(&tr)
	if err != nil {
		json.NewEncoder(w).Encode("error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tr)
}


// Delete an existing Transaction
// No validation need but transaction ID must exists
func (c *TransactionController) Remove(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.remove")

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := c.usecases.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}


// Return all existing Transactions (paginated) 
func (c *TransactionController) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.getall")
	
	page, err :=  strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	
	limit, err :=  strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}

	var trs dto.TransactionFindAllResponseDto
	trs, err = c.usecases.FindAll(page, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&trs)
}
