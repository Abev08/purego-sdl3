package sdl

// [Sensor] is a opaque structure used to identify an opened SDL sensor.
//
// [Sensor]: https://wiki.libsdl.org/SDL3/SDL_Sensor
type Sensor struct{}

// [SensorID] defines a unique ID for a sensor for the time it is connected to the system, and is never reused for the lifetime of the application.
//
// [SensorID]: https://wiki.libsdl.org/SDL3/SDL_SensorID
type SensorID uint32

// [SensorType] specifies different sensors defined by SDL.
//
// [SensorType]: https://wiki.libsdl.org/SDL3/SDL_SensorType
type SensorType int32

const (
	SensorInvalid SensorType = iota - 1 // Returned for an invalid sensor.
	SensorUnknown                       // Unknown sensor type.
	SensorAccel                         // Accelerometer.
	SensorGyro                          // Gyroscope.
	SensorAccelL                        // Accelerometer for left Joy-Con controller and Wii nunchuk.
	SensorGyroL                         // Gyroscope for left Joy-Con controller.
	SensorAccelR                        // Accelerometer for right Joy-Con controller.
	SensorGyroR                         // Gyroscope for right Joy-Con controller.
	SensorCount
)

// [GetSensors] gets a list of currently connected sensors.
//
// [GetSensors]: https://wiki.libsdl.org/SDL3/SDL_GetSensors
// func GetSensors(count *int32) *SensorID {
//	return sdlGetSensors(count)
// }

// [GetSensorNameForID] gets the implementation dependent name of a sensor.
//
// [GetSensorNameForID]: https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID
// func GetSensorNameForID(instanceID SensorID) string {
//	return sdlGetSensorNameForID(instanceID)
// }

// [GetSensorTypeForID] gets the type of a sensor.
//
// [GetSensorTypeForID]: https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID
// func GetSensorTypeForID(instanceID SensorID) SensorType {
//	return sdlGetSensorTypeForID(instanceID)
// }

// [GetSensorNonPortableTypeForID] gets the platform dependent type of a sensor.
//
// [GetSensorNonPortableTypeForID]: https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID
// func GetSensorNonPortableTypeForID(instanceID SensorID) int32 {
//	return sdlGetSensorNonPortableTypeForID(instanceID)
// }

// [OpenSensor] opens a sensor for use.
//
// [OpenSensor]: https://wiki.libsdl.org/SDL3/SDL_OpenSensor
// func OpenSensor(instanceID SensorID) *Sensor {
//	return sdlOpenSensor(instanceID)
// }

// [GetSensorFromID] returns the [Sensor] associated with an instance ID.
//
// [GetSensorFromID]: https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID
// func GetSensorFromID(instanceID SensorID) *Sensor {
//	return sdlGetSensorFromID(instanceID)
// }

// [GetSensorProperties] gets the properties associated with a sensor.
//
// [GetSensorProperties]: https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties
// func GetSensorProperties(sensor *Sensor) PropertiesID {
//	return sdlGetSensorProperties(sensor)
// }

// [GetSensorName] gets the implementation dependent name of a sensor.
//
// [GetSensorName]: https://wiki.libsdl.org/SDL3/SDL_GetSensorName
// func GetSensorName(sensor *Sensor) string {
//	return sdlGetSensorName(sensor)
// }

// [GetSensorType] gets the type of a sensor.
//
// [GetSensorType]: https://wiki.libsdl.org/SDL3/SDL_GetSensorType
// func GetSensorType(sensor *Sensor) SensorType {
//	return sdlGetSensorType(sensor)
// }

// [GetSensorNonPortableType] gets the platform dependent type of a sensor.
//
// [GetSensorNonPortableType]: https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType
// func GetSensorNonPortableType(sensor *Sensor) int32 {
//	return sdlGetSensorNonPortableType(sensor)
// }

// [GetSensorID] gets the instance ID of a sensor.
//
// [GetSensorID]: https://wiki.libsdl.org/SDL3/SDL_GetSensorID
// func GetSensorID(sensor *Sensor) SensorID {
//	return sdlGetSensorID(sensor)
// }

// [GetSensorData] gets the current state of an opened sensor.
//
// [GetSensorData]: https://wiki.libsdl.org/SDL3/SDL_GetSensorData
// func GetSensorData(sensor *Sensor, data *float32, numValues int32) bool {
//	return sdlGetSensorData(sensor, data, numValues)
// }

// [CloseSensor] closes a sensor previously opened with [OpenSensor].
//
// [CloseSensor]: https://wiki.libsdl.org/SDL3/SDL_CloseSensor
// func CloseSensor(sensor *Sensor) {
//	sdlCloseSensor(sensor)
// }

// [UpdateSensors] updates the current state of the open sensors.
//
// [UpdateSensors]: https://wiki.libsdl.org/SDL3/SDL_UpdateSensors
// func UpdateSensors() {
//	sdlUpdateSensors()
// }
