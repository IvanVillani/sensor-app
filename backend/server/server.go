package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seeis/sensor-app/backend/db"
	"github.com/seeis/sensor-app/backend/logger"
)

type ActionType int

const (
	GetAllDevicesType = iota
	GetDeviceByIdType
	CreateNewDeviceType
	DeleteDeviceByIdType
	UpdateDeviceByIdType

	GetAllSensorsType
	GetSensorByIdType
	CreateNewSensorType
	DeleteSensorByIdType
	UpdateSensorByIdType

	GetAllMeasurementsType
	CreateNewMeasurementType

	GetSensorAverageValueType
)

type Result struct {
	JSON []byte
	Err  error
}

type Action struct {
	Type      ActionType
	ID        string
	Data      []byte
	ReplyChan chan Result
}

type Server struct {
	ActionChan chan<- Action
}

type ValidationError struct {
	Msg string
}

func (err ValidationError) Error() string {
	return err.Msg
}

func HandleRequests() {
	db.InitializePostgreSQLConnection()
	db.InitializeInfluxDBConnection()

	defer db.PostgreSQL.Close()
	defer db.InfluxClient.Close()

	r := mux.NewRouter().StrictSlash(true)

	server := &Server{ActionChan: StartServerManager()}

	// Handle device-related paths
	r.HandleFunc("/api/v1/device", server.ReturnAllDevices).Methods("GET")
	r.HandleFunc("/api/v1/device/{id}", server.ReturnDeviceById).Methods("GET")
	r.HandleFunc("/api/v1/device", server.CreateNewDevice).Methods("POST")
	r.HandleFunc("/api/v1/device/{id}", server.DeleteDeviceById).Methods("DELETE")
	r.HandleFunc("/api/v1/device/{id}", server.UpdateDeviceById).Methods("PUT")

	// Handle sensor-related paths
	r.HandleFunc("/api/v1/sensor", server.ReturnAllSensors).Methods("GET")
	r.HandleFunc("/api/v1/sensor/{id}", server.ReturnSensorByID).Methods("GET")
	r.HandleFunc("/api/v1/sensor", server.CreateNewSensor).Methods("POST")
	r.HandleFunc("/api/v1/sensor/{id}", server.DeleteSensorById).Methods("DELETE")
	r.HandleFunc("/api/v1/sensor/{id}", server.UpdateSensorById).Methods("PUT")

	// Handle measurement-related paths
	r.HandleFunc("/api/v1/measurement", server.ReturnAllMeasurements).Methods("GET")
	r.HandleFunc("/api/v1/measurement", server.CreateNewMeasurement).Methods("POST")
	r.HandleFunc("/api/v1/measurement", server.DeleteMeasurement).Methods("DELETE")
	r.HandleFunc("/api/v1/measurement", server.UpdateMeasurement).Methods("PUT")

	// Handle other paths
	r.HandleFunc("/api/v1/sensorAvarageValue", server.ReturnSensorAverageValue).Methods("GET")

	logger.Fatal.Fatal(http.ListenAndServe(":8000", r))
}
