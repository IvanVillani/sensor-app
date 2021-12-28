package server

import (
	"database/sql"
	"encoding/json"

	"github.com/seeis/sensor-app/backend/db"
)

func GetAllDevices() ([]byte, error) {
	entries, err := db.PostgreSQL.Query("SELECT * FROM device;")

	if err != nil {
		return nil, err
	}

	defer entries.Close()

	devices := make([]Device, 0)

	for entries.Next() {
		device, err := extractDevice(entries)

		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	if err = entries.Err(); err != nil {
		return nil, err
	}

	return json.Marshal(devices)
}

func GetDeviceById(id string) ([]byte, error) {
	entry := db.PostgreSQL.QueryRow("SELECT * FROM device WHERE id = $1;", id)

	device := Device{}

	err := entry.Scan(&device.ID, &device.Name, &device.Description)

	if err != nil {
		return nil, err
	}

	var sensors []Sensor

	err = GetAllSensorsByDeviceId(device.ID, &sensors)

	if err != nil {
		return nil, err
	}

	device.Sensors = sensors

	return json.Marshal(device)
}

func CreateNewDevice(data []byte) error {
	device := Device{}

	json.Unmarshal(data, &device)

	_, err := db.PostgreSQL.Exec("INSERT INTO device (id, name, description) VALUES ($1, $2, $3);",
		device.ID, device.Name, device.Description)

	return err
}

func DeleteDeviceById(id string) error {
	if _, err := GetDeviceById(id); err != nil {
		return err
	}

	err := DeleteSensorsByDeviceId(id)

	if err != nil {
		return err
	}

	_, err = db.PostgreSQL.Exec("DELETE FROM device WHERE id=$1;", id)

	return err
}

func UpdateDeviceById(id string, data []byte) error {
	entry := db.PostgreSQL.QueryRow("SELECT * FROM device WHERE id = $1;", id)

	device := Device{}

	err := entry.Scan(&device.ID, &device.Name, &device.Description)

	if err != nil {
		return err
	}

	json.Unmarshal(data, &device)

	if device.Name != "" {
		_, err = db.PostgreSQL.Exec("UPDATE device SET name=$1 WHERE id=$2;", device.Name, id)

		if err != nil {
			return err
		}
	}

	if device.Description != "" {
		_, err = db.PostgreSQL.Exec("UPDATE device SET description=$1 WHERE id=$2;", device.Description, id)

		if err != nil {
			return err
		}
	}

	return err
}

func extractDevice(entries *sql.Rows) (Device, error) {
	device := Device{}

	err := entries.Scan(&device.ID, &device.Name, &device.Description)

	if err != nil {
		return device, err
	}

	var sensors []Sensor

	err = GetAllSensorsByDeviceId(device.ID, &sensors)

	if err != nil {
		return device, err
	}

	device.Sensors = sensors

	return device, nil
}
