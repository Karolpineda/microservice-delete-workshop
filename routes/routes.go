package routes

import (
	"microservicedeleteteworkshops/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Ruta Health
	r.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	r.HandleFunc("/workshops/{id}", controllers.DeleteWorkshop).Methods("DELETE")

}
