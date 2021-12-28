package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/seeis/sensor-app/backend/constants"
	"github.com/seeis/sensor-app/backend/db"
)

// AverageRequest structure: used for parsing request data
type AverageRequest struct {
	DeviceID  string `json:"deviceId" yaml:"deviceId"`
	SensorID  string `json:"sensorId" yaml:"sensorId"`
	StartTime string `json:"startTime" yaml:"startTime"`
	EndTime   string `json:"endTime" yaml:"endTime"`
}

// AverageResponse structure: used for parsing to JSON
type AverageResponse struct {
	Value string `json:"sensorAverageValue" yaml:"sensorAverageValue"`
}

func GetAllMeasurements() ([]byte, error) {
	query := `import "json" from(bucket: "` + constants.BucketName + `") ` +
		`|> range(start: -30m) ` +
		`|> filter(fn: (r) => r._measurement == "cpu") ` +
		`|> map(fn: (r) => ({r with jsonResult: string(v: json.encode(v: ` +
		`{"measuredAt": r._time, "value": r._value, "sensorId": r.sensorId, "deviceId": r.deviceId})), }),)`

	queryAPI := db.InfluxClient.QueryAPI(constants.OrgName)

	entries, err := queryAPI.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer

	for entries.Next() {
		buffer.WriteString(fmt.Sprint(entries.Record().ValueByKey("jsonResult")))
	}

	if entries.Err() != nil {
		return nil, entries.Err()
	}

	return buffer.Bytes(), nil
}

func GetAllMeasurementsBySensorId(sensorId string, measurements *[]Measurement) error {
	query := `from(bucket: "` + constants.BucketName + `") ` +
		`|> range(start: -30m) ` +
		`|> filter(fn: (r) => r._measurement == "cpu" and r.sensorId == "` + sensorId + `") ` +
		`|> map(fn:(r) => ({ _time: r._time, _value: r._value, sensorId: r.sensorId, deviceId: r.deviceId }))`

	queryAPI := db.InfluxClient.QueryAPI(constants.OrgName)

	entries, err := queryAPI.Query(context.Background(), query)

	if err != nil {
		return err
	}

	for entries.Next() {
		*measurements = append(*measurements, Measurement{
			MeasuredAt: entries.Record().Time(),
			Value:      fmt.Sprint(entries.Record().Value()),
			SensorID:   fmt.Sprint(entries.Record().ValueByKey("sensorId")),
			DeviceID:   fmt.Sprint(entries.Record().ValueByKey("deviceId")),
		})
	}

	return entries.Err()
}

func CreateNewMeasurement(data []byte) error {
	var measurement Measurement

	json.Unmarshal(data, &measurement)

	err := validateSensorAndDeviceIDs(measurement.DeviceID, measurement.SensorID)

	if err != nil {
		return err
	}

	p := influxdb2.NewPointWithMeasurement("cpu").
		AddTag("type", "temperature").
		AddTag("sensorId", measurement.SensorID).
		AddTag("deviceId", measurement.DeviceID).
		AddField("value", measurement.Value).
		SetTime(measurement.MeasuredAt)

	writeAPI := db.InfluxClient.WriteAPI(constants.OrgName, constants.BucketName)
	writeAPI.WritePoint(p)

	return nil
}

func GetSensorAverageValue(data []byte) ([]byte, error) {
	var averageRequest AverageRequest

	json.Unmarshal(data, &averageRequest)

	err := validateSensorAndDeviceIDs(averageRequest.DeviceID, averageRequest.SensorID)

	if err != nil {
		return nil, err
	}

	query := `from(bucket: "` + constants.BucketName + `") ` +
		`|> range(start: ` + averageRequest.StartTime + `, stop: ` + averageRequest.EndTime + `) ` +
		`|> filter(fn: (r) => ` +
		`r._measurement == "` + constants.TypeOfMeasurement + `" and ` +
		`r.deviceId == "` + averageRequest.DeviceID + `" and ` +
		`r.sensorId == "` + averageRequest.SensorID + `") ` +
		`|> map(fn:(r) => ({ r with _value: float(v: r._value) }))` +
		`|> mean()`

	queryAPI := db.InfluxClient.QueryAPI(constants.OrgName)

	entries, err := queryAPI.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var averageResponse AverageResponse

	for entries.Next() {
		averageResponse.Value = fmt.Sprint(entries.Record().Value())
	}

	if entries.Err() != nil {
		return nil, entries.Err()
	}

	return json.Marshal(averageResponse)
}

func DeleteMeasurementsBySensorId(sensorId string) error {
	orgDomain, err := db.InfluxClient.OrganizationsAPI().FindOrganizationByName(context.Background(), constants.OrgName)

	if err != nil {
		return err
	}

	bucketDomain, err := db.InfluxClient.BucketsAPI().FindBucketByName(context.Background(), constants.BucketName)

	if err != nil {
		return err
	}

	return db.InfluxClient.DeleteAPI().Delete(context.Background(), orgDomain, bucketDomain,
		time.Now().AddDate(0, 0, -3),
		time.Now(),
		`sensorId="`+sensorId+`"`)
}

func validateSensorAndDeviceIDs(deviceId, sensorId string) error {
	if check, err := GetDeviceById(deviceId); err != nil || len(check) == 0 {
		if len(check) == 0 {
			return ValidationError{
				Msg: "invalid device ID specified: " + deviceId,
			}
		}
		return err
	}

	if check, err := GetSensorById(sensorId); err != nil || len(check) == 0 {
		if len(check) == 0 {
			return ValidationError{
				Msg: "invalid sensor ID specified: " + sensorId,
			}
		}
		return err
	}

	return nil
}
