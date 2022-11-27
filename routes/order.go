package routes

import (
	"waysbucks-api/handlers"
	"waysbucks-api/pkg/middleware"
	"waysbucks-api/pkg/mysql"
	"waysbucks-api/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerOrder(orderRepository)
	r.HandleFunc("/orders", middleware.Auth(h.FindOrders)).Methods("GET")
	r.HandleFunc("/order/{id}", middleware.Auth(h.GetOrder)).Methods("GET")
	r.HandleFunc("/order/{id}", middleware.Auth(h.AddOrder)).Methods("POST")
	r.HandleFunc("/order/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
}
