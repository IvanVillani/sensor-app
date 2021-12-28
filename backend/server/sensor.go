package server

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seeis/sensor-app/backend/logger"
)

// Sensor structure: holds data of a specific sensor
type Sensor struct {
	ID           string        `json:"id" yaml:"id"`
	Name         string        `json:"name" yaml:"name"`
	Description  string        `json:"description" yaml:"description"`
	Unit         string        `json:"unit" yaml:"unit"`
	DeviceID     string        `json:"deviceId" yaml:"deviceId"`
	Groups       Group         `json:"sensorGroups" yaml:"sensorGroups"`
	Measurements []Measurement `json:"measurements" yaml:"measurements"`
}

func (s *Server) ReturnAllSensors(w http.ResponseWriter, r *http.Request) {
	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      GetAllSensorsType,
		ReplyChan: replyChan}

	reply := <-replyChan

	if reply.Err != nil {
		logger.Error.Printf("Failed to return endpoint (all sensors): %s\n", reply.Err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, returned all sensors.")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply.JSON)
}

func (s *Server) ReturnSensorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      GetSensorByIdType,
		ID:        id,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	switch {
	case err == sql.ErrNoRows:
		logger.Error.Printf("Failed to return endpoint (sensor): %s\n", err)

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	case err != nil:
		logger.Error.Printf("Failed to return endpoint (sensor): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, returned sensor.")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply.JSON)
}

func (s *Server) CreateNewSensor(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var sensor Sensor

	json.Unmarshal(reqBody, &sensor)

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      CreateNewSensorType,
		Data:      reqBody,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	if err != nil {
		logger.Error.Printf("Failed to create endpoint (sensor): %s\n", err)

		if _, ok := err.(ValidationError); ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	logger.Info.Println("Endpoint reached, created new sensor.")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensor)
}

func (s *Server) DeleteSensorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      DeleteSensorByIdType,
		ID:        id,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	if err != nil {
		logger.Error.Printf("Failed to delete endpoint (sensor): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, deleted device.")

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) UpdateSensorById(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var sensor Sensor

	json.Unmarshal(reqBody, &sensor)

	vars := mux.Vars(r)

	id := vars["id"]

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      UpdateSensorByIdType,
		ID:        id,
		Data:      reqBody,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	switch {
	case err == sql.ErrNoRows:
		logger.Error.Printf("Failed to update endpoint (sensor): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusNotFound)

		return
	case err != nil:
		logger.Error.Printf("Failed to update endpoint (sensor): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, updated device.")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensor)
}
