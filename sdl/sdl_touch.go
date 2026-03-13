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

// [Finger] is a structure storing data about a single finger in a multitouch event.
//
// [Finger]: https://wiki.libsdl.org/SDL3/SDL_Finger
type Finger struct {
	ID       FingerID // The finger ID.
	X        float32  // The x-axis location of the touch event, normalized (0...1).
	Y        float32  // The y-axis location of the touch event, normalized (0...1).
	Pressure float32  // The quantity of pressure applied, normalized (0...1).
}

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
