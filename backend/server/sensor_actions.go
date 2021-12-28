package server

import (
	"database/sql"
	"encoding/json"

	"github.com/seeis/sensor-app/backend/db"
)

func GetAllSensors() ([]byte, error) {
	entries, err := db.PostgreSQL.Query("SELECT * FROM sensor;")

	if err != nil {
		return nil, err
	}

	defer entries.Close()

	sensors := make([]Sensor, 0)

	for entries.Next() {
		sensor := Sensor{}

		err := entries.Scan(&sensor.ID, &sensor.Name, &sensor.Description, &sensor.Unit, &sensor.Groups, &sensor.DeviceID)

		if err != nil {
			return nil, err
		}

		var measurements []Measurement

		err = GetAllMeasurementsBySensorId(sensor.ID, &measurements)

		if err != nil {
			return nil, err
		}

		sensor.Measurements = measurements

		sensors = append(sensors, sensor)
	}

	if err = entries.Err(); err != nil {
		return nil, err
	}

	return json.Marshal(sensors)
}

func GetAllSensorsByDeviceId(deviceId string, sensors *[]Sensor) error {
	entries, err := db.PostgreSQL.Query("SELECT * FROM sensor WHERE device_id = $1;", deviceId)

	if err != nil {
		return err
	}

	defer entries.Close()

	for entries.Next() {
		sensor, err := extractSensor(entries)

		if err != nil {
			return err
		}

		*sensors = append(*sensors, sensor)
	}

	return entries.Err()
}

func GetSensorById(id string) ([]byte, error) {
	entry := db.PostgreSQL.QueryRow("SELECT * FROM sensor WHERE id = $1;", id)

	sensor := Sensor{}

	err := entry.Scan(&sensor.ID, &sensor.Name, &sensor.Description, &sensor.Unit, &sensor.Groups, &sensor.DeviceID)

	if err != nil {
		return nil, err
	}

	var measurements []Measurement

	err = GetAllMeasurementsBySensorId(sensor.ID, &measurements)

	if err != nil {
		return nil, err
	}

	sensor.Measurements = measurements

	return json.Marshal(sensor)
}

func CreateNewSensor(data []byte) error {
	sensor := Sensor{}

	json.Unmarshal(data, &sensor)

	err := validateDeviceID(sensor.DeviceID)

	if err != nil {
		return err
	}

	_, err = db.PostgreSQL.Exec("INSERT INTO sensor (id, name, description, unit, sensor_groups, device_id) VALUES ($1, $2, $3, $4, $5, $6);",
		sensor.ID, sensor.Name, sensor.Description, sensor.Unit, sensor.Groups, sensor.DeviceID)

	return err
}

func DeleteSensorById(id string) error {
	_, err := db.PostgreSQL.Exec("DELETE FROM sensor WHERE id=$1;", id)

	if err != nil {
		return err
	}

	return DeleteMeasurementsBySensorId(id)
}

func DeleteSensorsByDeviceId(deviceId string) error {
	entries, err := db.PostgreSQL.Query("SELECT id FROM sensor WHERE device_id=$1;", deviceId)

	if err != nil {
		return err
	}

	defer entries.Close()

	for entries.Next() {
		var sensorId string

		err := entries.Scan(&sensorId)

		if err != nil {
			return err
		}

		err = DeleteSensorById(sensorId)

		if err != nil {
			return err
		}
	}

	return entries.Err()
}

func UpdateSensorById(id string, data []byte) error {
	entry := db.PostgreSQL.QueryRow("SELECT * FROM sensor WHERE id = $1;", id)

	sensor := Sensor{}

	err := entry.Scan(&sensor.ID, &sensor.Name, &sensor.Description, &sensor.Unit, &sensor.Groups, &sensor.DeviceID)

	if err != nil {
		return err
	}

	json.Unmarshal(data, &sensor)

	if sensor.Name != "" {
		_, err = db.PostgreSQL.Exec("UPDATE sensor SET name=$1 WHERE id=$2;", sensor.Name, id)

		if err != nil {
			return err
		}
	}

	if sensor.Description != "" {
		_, err = db.PostgreSQL.Exec("UPDATE sensor SET description=$1 WHERE id=$2;", sensor.Description, id)

		if err != nil {
			return err
		}
	}

	if sensor.Unit != "" {
		_, err = db.PostgreSQL.Exec("UPDATE sensor SET unit=$1 WHERE id=$2;", sensor.Unit, id)

		if err != nil {
			return err
		}
	}

	return err
}

func extractSensor(entries *sql.Rows) (Sensor, error) {
	sensor := Sensor{}

	err := entries.Scan(&sensor.ID, &sensor.Name, &sensor.Description, &sensor.Unit, &sensor.Groups, &sensor.DeviceID)

	if err != nil {
		return sensor, err
	}

	var measurements []Measurement

	err = GetAllMeasurementsBySensorId(sensor.ID, &measurements)

	if err != nil {
		return sensor, err
	}

	sensor.Measurements = measurements

	return sensor, nil
}

func validateDeviceID(deviceId string) error {
	if check, err := GetDeviceById(deviceId); err != nil || len(check) == 0 {
		if len(check) == 0 {
			return ValidationError{
				Msg: "invalid device ID specified: " + deviceId,
			}
		}
		return err
	}

	return nil
}
