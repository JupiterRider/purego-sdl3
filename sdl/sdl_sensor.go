package sdl

// [SensorType] is a structure specifying the different sensors defined by SDLs.
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
)

// [SensorID] is a unique ID for a sensor for the time it is connected to the system, and is never reused for the lifetime of the application.
//
// [SensorID]: https://wiki.libsdl.org/SDL3/SDL_SensorID
type SensorID uint32

// func CloseSensor(sensor *Sensor)  {
//	sdlCloseSensor(sensor)
// }

// func GetSensorData(sensor *Sensor, data *float32, num_values int32) bool {
//	return sdlGetSensorData(sensor, data, num_values)
// }

// func GetSensorFromID(instance_id SensorID) *Sensor {
//	return sdlGetSensorFromID(instance_id)
// }

// func GetSensorID(sensor *Sensor) SensorID {
//	return sdlGetSensorID(sensor)
// }

// func GetSensorName(sensor *Sensor) string {
//	return sdlGetSensorName(sensor)
// }

// func GetSensorNameForID(instance_id SensorID) string {
//	return sdlGetSensorNameForID(instance_id)
// }

// func GetSensorNonPortableType(sensor *Sensor) int32 {
//	return sdlGetSensorNonPortableType(sensor)
// }

// func GetSensorNonPortableTypeForID(instance_id SensorID) int32 {
//	return sdlGetSensorNonPortableTypeForID(instance_id)
// }

// func GetSensorProperties(sensor *Sensor) PropertiesID {
//	return sdlGetSensorProperties(sensor)
// }

// func GetSensors(count *int32) *SensorID {
//	return sdlGetSensors(count)
// }

// func GetSensorType(sensor *Sensor) SensorType {
//	return sdlGetSensorType(sensor)
// }

// func GetSensorTypeForID(instance_id SensorID) SensorType {
//	return sdlGetSensorTypeForID(instance_id)
// }

// func OpenSensor(instance_id SensorID) *Sensor {
//	return sdlOpenSensor(instance_id)
// }

// func UpdateSensors()  {
//	sdlUpdateSensors()
// }
