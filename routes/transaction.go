package routes

import (
	"waysbucks-api/handlers"
	"waysbucks-api/pkg/middleware"
	"waysbucks-api/pkg/mysql"
	"waysbucks-api/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.FindTransaction)).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction-user", middleware.Auth(h.GetTransactionByUser)).Methods("GET")
	
	// r.HandleFunc("/transaction/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")
	r.HandleFunc("/my-order", middleware.Auth(h.GetOrderByID)).Methods("GET")
	r.HandleFunc("/transUpdate/{id}", middleware.Auth(h.UpdateAdmin)).Methods("PATCH")
}