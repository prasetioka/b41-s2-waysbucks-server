package routes

import (
	"waysbucks-api/handlers"
	"waysbucks-api/pkg/mysql"
	"waysbucks-api/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	ProductRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(ProductRepository)

	r.HandleFunc("/products", h.FindProducts).Methods("GET")
	// r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	// r.HandleFunc("/user", h.CreateUser).Methods("POST")
	// r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	// r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}