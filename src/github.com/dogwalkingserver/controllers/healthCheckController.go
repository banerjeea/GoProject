package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//HealthCheck :
type HealthCheck struct{}

//NewHealthCheckController :
func NewHealthCheckController() *HealthCheck {

	return &HealthCheck{}

}

//GetHealth :
func (hc HealthCheck) GetHealth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode("Test server!!")
}
