package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/seeis/sensor-app/backend/logger"
)

// Measurement structure: holds data of a specific measurement
type Measurement struct {
	MeasuredAt time.Time `json:"measuredAt" yaml:"measuredAt"`
	Value      string    `json:"value" yaml:"value"`
	SensorID   string    `json:"sensorId" yaml:"sensorId"`
	DeviceID   string    `json:"deviceId" yaml:"deviceId"`
}

func (s *Server) ReturnAllMeasurements(w http.ResponseWriter, r *http.Request) {
	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      GetAllMeasurementsType,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	if err != nil {
		logger.Error.Printf("Failed to return endpoint (all measurements): %s\n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	logger.Info.Println("Endpoint reached, returned all measurements.")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply.JSON)
}

func (s *Server) CreateNewMeasurement(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var measurement Measurement

	json.Unmarshal(reqBody, &measurement)

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      CreateNewMeasurementType,
		Data:      reqBody,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	if err != nil {
		logger.Error.Printf("Failed to create endpoint (measurement): %s\n", reply.Err)

		if _, ok := err.(ValidationError); ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	logger.Info.Println("Endpoint reached, created new measurement.")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurement)
}

func (s *Server) ReturnSensorAverageValue(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	replyChan := make(chan Result)

	s.ActionChan <- Action{
		Type:      GetSensorAverageValueType,
		Data:      reqBody,
		ReplyChan: replyChan}

	reply := <-replyChan

	err := reply.Err

	if err != nil {
		logger.Error.Printf("Failed to return endpoint (sensorAverageValue): %s\n", err)

		if _, ok := err.(ValidationError); ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	logger.Info.Println("Endpoint reached, returned sensorAverageValue.")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply.JSON)
}

func (s *Server) DeleteMeasurement(w http.ResponseWriter, r *http.Request) {
	logger.Error.Println("Failed to execute DELETE request (measurement).")

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Server) UpdateMeasurement(w http.ResponseWriter, r *http.Request) {
	logger.Error.Println("Failed to execute PUT request (measurement).")

	w.WriteHeader(http.StatusMethodNotAllowed)
}
