package server

import (
	"github.com/seeis/sensor-app/backend/logger"
)

func StartServerManager() chan<- Action {
	acts := make(chan Action)

	go func() {
		for act := range acts {
			switch act.Type {
			case GetAllDevicesType:
				jsonResult, err := GetAllDevices()

				act.ReplyChan <- Result{
					JSON: jsonResult,
					Err:  err,
				}
			case GetDeviceByIdType:
				jsonResult, err := GetDeviceById(act.ID)

				act.ReplyChan <- Result{
					JSON: jsonResult,
					Err:  err,
				}
			case CreateNewDeviceType:
				err := CreateNewDevice(act.Data)

				act.ReplyChan <- Result{
					Err: err,
				}
			case DeleteDeviceByIdType:
				err := DeleteDeviceById(act.ID)

				act.ReplyChan <- Result{
					Err: err,
				}
			case UpdateDeviceByIdType:
				err := UpdateDeviceById(act.ID, act.Data)

				act.ReplyChan <- Result{
					Err: err,
				}
			case GetAllSensorsType:
				jsonResult, err := GetAllSensors()

				act.ReplyChan <- Result{
					JSON: jsonResult,
					Err:  err,
				}
			case GetSensorByIdType:
				jsonResult, err := GetSensorById(act.ID)

				act.ReplyChan <- Result{
					JSON: jsonResult,
					Err:  err,
				}
			case CreateNewSensorType:
				err := CreateNewSensor(act.Data)

				act.ReplyChan <- Result{
					Err: err,
				}
			case DeleteSensorByIdType:
				err := DeleteSensorById(act.ID)

				act.ReplyChan <- Result{
					Err: err,
				}
			case UpdateSensorByIdType:
				err := UpdateSensorById(act.ID, act.Data)

				act.ReplyChan <- Result{
					Err: err,
				}
			case GetAllMeasurementsType:
				jsonResult, err := GetAllMeasurements()

				act.ReplyChan <- Result{
					JSON: jsonResult,
					Err:  err,
				}
			case CreateNewMeasurementType:
				err := CreateNewMeasurement(act.Data)

				act.ReplyChan <- Result{
					Err: err,
				}
			case GetSensorAverageValueType:
				jsonResult, err := GetSensorAverageValue(act.Data)

				act.ReplyChan <- Result{
					JSON: jsonResult,
					Err:  err,
				}
			default:
				logger.Fatal.Fatal("Unknown action type: ", act.Type)
			}
		}
	}()
	return acts
}
