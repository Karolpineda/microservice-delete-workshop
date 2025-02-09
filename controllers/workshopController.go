package controllers

import (
	"encoding/json"
	"microservicedeleteteworkshops/config"
	"microservicedeleteteworkshops/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// DeleteWorkshop
// @Summary Elimina un workshop existente
// @Description Elimina un workshop basado en su ID
// @Tags Workshops
// @Accept json
// @Produce json
// @Param id path string true "ID del Workshop"
// @Success 200 {string} string "Workshop deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Workshop not found"
// @Router /workshops/{id} [delete]
func DeleteWorkshop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		http.Error(w, "Missing workshop ID", http.StatusBadRequest)
		return
	}

	workshopID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid workshop ID format", http.StatusBadRequest)
		return
	}

	db := config.SetupDatabase()
	err = models.DeleteWorkshop(db, workshopID)
	if err != nil {
		http.Error(w, "Workshop not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Workshop deleted successfully"})
}

// HealthCheck
// @Summary Verifica el estado del microservicio
// @Description Retorna un mensaje que indica que el microservicio está en funcionamiento
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Microservice is up and running"})
}
