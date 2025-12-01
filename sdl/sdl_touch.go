package sdl

// [TouchDeviceType] is an enum that describes the type of a touch device.
//
// [TouchDeviceType]: https://wiki.libsdl.org/SDL3/SDL_TouchDeviceType
type TouchDeviceType int32

const (
	TouchDeviceInvalid          TouchDeviceType = iota - 1
	TouchDeviceDirect                           // Touch screen with window-relative coordinates.
	TouchDeviceIndirectAbsolute                 // Trackpad with absolute device coordinates.
	TouchDeviceIndirectRelative                 // Trackpad with screen cursor-relative coordinates.
)

// [TouchID] is a unique ID for a touch device.
//
// [TouchID]: https://wiki.libsdl.org/SDL3/SDL_TouchID
type TouchID uint64

// [FingerID] is a unique ID for a single finger on a touch device.
//
// [FingerID]: https://wiki.libsdl.org/SDL3/SDL_FingerID
type FingerID uint64

// func GetTouchDeviceName(touchID TouchID) string {
//	return sdlGetTouchDeviceName(touchID)
// }

// func GetTouchDevices(count *int32) *TouchID {
//	return sdlGetTouchDevices(count)
// }

// func GetTouchDeviceType(touchID TouchID) TouchDeviceType {
//	return sdlGetTouchDeviceType(touchID)
// }

// func GetTouchFingers(touchID TouchID, count *int32) **Finger {
//	return sdlGetTouchFingers(touchID, count)
// }
