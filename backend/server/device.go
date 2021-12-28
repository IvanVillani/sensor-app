package server

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seeis/sensor-app/backend/logger"
)

// Device structure: holds data of a specific device
type Device struct {
	ID          string   `json:"id" yaml:"id"`
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description" yaml:"description"`
	Sensors     []Sensor `json:"sensors" yaml:"sensors"`
}

func (s *Server) ReturnAllDevices(w http.ResponseWriter, r *http.Request) {
	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      GetAllDevicesType,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	if err != nil {
		logger.Error.Printf("Failed to return endpoint (all devices): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, returned all devices.")

	w.Header().Set("Content-Type", "application/json")
	w.Write(reply.JSON)
}

func (s *Server) ReturnDeviceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      GetDeviceByIdType,
		ID:        id,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	switch {
	case err == sql.ErrNoRows:
		logger.Error.Printf("Failed to return endpoint (device): %s\n", err)

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	case err != nil:
		logger.Error.Printf("Failed to return endpoint (device): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, returned device.")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply.JSON)
}

func (s *Server) CreateNewDevice(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var device Device

	json.Unmarshal(reqBody, &device)

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      CreateNewDeviceType,
		Data:      reqBody,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	if err != nil {
		logger.Error.Printf("Failed to create endpoint (device): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, created new device.")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}

func (s *Server) DeleteDeviceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      DeleteDeviceByIdType,
		ID:        id,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	if err != nil {
		logger.Error.Printf("Failed to delete endpoint (device): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, deleted device.")

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) UpdateDeviceById(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var device Device

	json.Unmarshal(reqBody, &device)

	vars := mux.Vars(r)

	id := vars["id"]

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      UpdateDeviceByIdType,
		ID:        id,
		Data:      reqBody,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	switch {
	case err == sql.ErrNoRows:
		logger.Error.Printf("Failed to update endpoint (device): %s\n", err)

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	case err != nil:
		logger.Error.Printf("Failed to update endpoint (device): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, updated device.")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}
