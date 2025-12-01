package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

// [Joystick] is a structure specifying the joystick structure used to identify an SDL joysticks.
//
// [Joystick]: https://wiki.libsdl.org/SDL3/SDL_Joystick
type Joystick struct{}

// [JoystickID] is a unique ID for a joystick for the time it is connected to the system, and is never reused for the lifetime of the application.
//
// [JoystickID]: https://wiki.libsdl.org/SDL3/SDL_JoystickID
type JoystickID uint32

// [JoystickType] is an enum of some common joystick types.
//
// [JoystickType]: https://wiki.libsdl.org/SDL3/SDL_JoystickType
type JoystickType uint16

const (
	JoystickTypeUnknown JoystickType = iota
	JoystickTypeGamepad
	JoystickTypeWheel
	JoystickTypeArcadeStick
	JoystickTypeFlightStick
	JoystickTypeDancePad
	JoystickTypeGuitar
	JoystickTypeDrumKit
	JoystickTypeArcadePad
	JoystickTypeThrottle
	JoystickTypeCount
)

// [JoystickConnectionState] defines the possible connection states for a joystick device.
//
// [JoystickConnectionState]: https://wiki.libsdl.org/SDL3/SDL_JoystickConnectionState
type JoystickConnectionState int32

const (
	JoystickConnectionInvalid JoystickConnectionState = iota - 1
	JoystickConnectionUnknown
	JoystickConnectionWired
	JoystickConnectionWireless
)

// [LockJoysticks] locks for atomic access to the joystick API.
//
// [LockJoysticks]: https://wiki.libsdl.org/SDL3/SDL_LockJoysticks
func LockJoysticks() {
	sdlLockJoysticks()
}

// [UnlockJoysticks] unlocks for atomic access to the joystick API.
//
// [UnlockJoysticks]: https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks
func UnlockJoysticks() {
	sdlUnlockJoysticks()
}

// [HasJoystick] returns whether a joystick is currently connected.
//
// [HasJoystick]: https://wiki.libsdl.org/SDL3/SDL_HasJoystick
func HasJoystick() bool {
	return sdlHasJoystick()
}

// [GetJoysticks] gets a list of currently connected joysticks.
//
// [GetJoysticks]: https://wiki.libsdl.org/SDL3/SDL_GetJoysticks
func GetJoysticks() []JoystickID {
	var count int32
	joysticks := sdlGetJoysticks(&count)
	if joysticks == nil {
		return nil
	}
	defer Free(unsafe.Pointer(joysticks))
	return mem.Copy(joysticks, count)
}

// [GetJoystickNameForID] gets the implementation dependent name of a joystick.
//
// [GetJoystickNameForID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID
func GetJoystickNameForID(instanceId JoystickID) string {
	return sdlGetJoystickNameForID(instanceId)
}

// [GetJoystickPathForID] gets the implementation dependent path of a joystick.
//
// [GetJoystickPathForID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID
func GetJoystickPathForID(instanceId JoystickID) string {
	return sdlGetJoystickPathForID(instanceId)
}

// [GetJoystickPlayerIndexForID] gets the player index of a joystick.
//
// [GetJoystickPlayerIndexForID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID
func GetJoystickPlayerIndexForID(instanceId JoystickID) int32 {
	return sdlGetJoystickPlayerIndexForID(instanceId)
}

// func GetJoystickGUIDForID(instanceId JoystickID) GUID {
// 	return sdlGetJoystickGUIDForID(instanceId)
// }

// [GetJoystickVendorForID] gets the USB vendor ID of a joystick, if available.
//
// [GetJoystickVendorForID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID
func GetJoystickVendorForID(instanceId JoystickID) uint16 {
	return sdlGetJoystickVendorForID(instanceId)
}

// [GetJoystickProductForID] gets the USB product ID of a joystick, if available.
//
// [GetJoystickProductForID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID
func GetJoystickProductForID(instanceId JoystickID) uint16 {
	return sdlGetJoystickProductForID(instanceId)
}

// [GetJoystickProductVersionForID] gets the product version of a joystick, if available.
//
// [GetJoystickProductVersionForID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID
func GetJoystickProductVersionForID(instanceId JoystickID) uint16 {
	return sdlGetJoystickProductVersionForID(instanceId)
}

// [GetJoystickTypeForID] gets the type of a joystick, if available.
//
// [GetJoystickTypeForID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID
func GetJoystickTypeForID(instanceId JoystickID) JoystickType {
	return sdlGetJoystickTypeForID(instanceId)
}

// [OpenJoystick] opens a joystick for use.
//
// [OpenJoystick]: https://wiki.libsdl.org/SDL3/SDL_OpenJoystick
func OpenJoystick(instanceId JoystickID) *Joystick {
	return sdlOpenJoystick(instanceId)
}

// [GetJoystickFromID] gets the [Joystick] associated with an instance ID, if it has been opened.
//
// [GetJoystickFromID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID
func GetJoystickFromID(instanceId JoystickID) *Joystick {
	return sdlGetJoystickFromID(instanceId)
}

// [GetJoystickFromPlayerIndex] gets the [Joystick] associated with a player index.
//
// [GetJoystickFromPlayerIndex]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex
func GetJoystickFromPlayerIndex(playerIndex int32) *Joystick {
	return sdlGetJoystickFromPlayerIndex(playerIndex)
}

// type VirtualJoystickTouchpadDesc struct {
// 	NFingers uint16
// 	Padding1 uint16
// 	Padding2 uint16
// 	Padding3 uint16
// }

// type VirtualJoystickSensorDesc struct {
// 	Type SensorType
// 	Rate float32
// }
//
// type VirtualJoystickDesc struct{}

// func AttachVirtualJoystick(desc *VirtualJoystickDesc) JoystickID {
// 	return sdlAttachVirtualJoystick(desc)
// }

// func DetachVirtualJoystick(instance_id JoystickID) bool {
// 	return sdlDetachVirtualJoystick(instance_id)
// }

// func IsJoystickVirtual(instance_id JoystickID) bool {
// 	return sdlIsJoystickVirtual(instance_id)
// }

// func SetJoystickVirtualAxis(joystick *Joystick, axis int32, value int16) bool {
// 	return sdlSetJoystickVirtualAxis(joystick, axis, value)
// }

// func SetJoystickVirtualBall(joystick *Joystick, ball int32, xrel int16, yrel int16) bool {
// 	return sdlSetJoystickVirtualBall(joystick, ball, xrel, yrel)
// }

// func SetJoystickVirtualButton(joystick *Joystick, button int32, down bool) bool {
// 	return sdlSetJoystickVirtualButton(joystick, button, down)
// }

// func SetJoystickVirtualHat(joystick *Joystick, hat int32, value uint8) bool {
// 	return sdlSetJoystickVirtualHat(joystick, hat, value)
// }

// func SetJoystickVirtualTouchpad(joystick *Joystick, touchpad int32, finger int32, down bool, x float32, y float32, pressure float32) bool {
// 	return sdlSetJoystickVirtualTouchpad(joystick, touchpad, finger, down, x, y, pressure)
// }

// func SendJoystickVirtualSensorData(joystick *Joystick, sensorType SensorType, sensor_timestamp uint64, data *float32, num_values int32) bool {
// 	return sdlSendJoystickVirtualSensorData(joystick, sensorType, sensor_timestamp, data, num_values)
// }

// [GetJoystickProperties] gets the properties associated with a joystick.
//
// [GetJoystickProperties]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties
func GetJoystickProperties(joystick *Joystick) PropertiesID {
	return sdlGetJoystickProperties(joystick)
}

const (
	PropJoystickCapMonoLEDBoolean       = "SDL.joystick.cap.mono_led"
	PropJoystickCapRGBLEDBoolean        = "SDL.joystick.cap.rgb_led"
	PropJoystickCapPlayerLEDBoolean     = "SDL.joystick.cap.player_led"
	PropJoystickCapRumbleBoolean        = "SDL.joystick.cap.rumble"
	PropJoystickCapTriggerRumbleBoolean = "SDL.joystick.cap.trigger_rumble"
)

// [GetJoystickName] gets the implementation dependent name of a joystick.
//
// [GetJoystickName]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickName
func GetJoystickName(joystick *Joystick) string {
	return sdlGetJoystickName(joystick)
}

// [GetJoystickPath] gets the implementation dependent path of a joystick.
//
// [GetJoystickPath]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath
func GetJoystickPath(joystick *Joystick) string {
	return sdlGetJoystickPath(joystick)
}

// [GetJoystickPlayerIndex] gets the player index of an opened joystick.
//
// [GetJoystickPlayerIndex]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex
func GetJoystickPlayerIndex(joystick *Joystick) int32 {
	return sdlGetJoystickPlayerIndex(joystick)
}

// [SetJoystickPlayerIndex] sets the player index of an opened joystick.
//
// [SetJoystickPlayerIndex]: https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex
func SetJoystickPlayerIndex(joystick *Joystick, playerIndex int32) bool {
	return sdlSetJoystickPlayerIndex(joystick, playerIndex)
}

// func GetJoystickGUID(joystick *Joystick) GUID {
// 	return sdlGetJoystickGUID(joystick)
// }

// [GetJoystickVendor] gets the USB vendor ID of an opened joystick, if available.
//
// [GetJoystickVendor]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor
func GetJoystickVendor(joystick *Joystick) uint16 {
	return sdlGetJoystickVendor(joystick)
}

// [GetJoystickProduct] gets the USB product ID of an opened joystick, if available.
//
// [GetJoystickProduct]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct
func GetJoystickProduct(joystick *Joystick) uint16 {
	return sdlGetJoystickProduct(joystick)
}

// [GetJoystickProductVersion] gets the product version of an opened joystick, if available.
//
// [GetJoystickProductVersion]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion
func GetJoystickProductVersion(joystick *Joystick) uint16 {
	return sdlGetJoystickProductVersion(joystick)
}

// [GetJoystickFirmwareVersion] gets the firmware version of an opened joystick, if available.
//
// [GetJoystickFirmwareVersion]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion
func GetJoystickFirmwareVersion(joystick *Joystick) uint16 {
	return sdlGetJoystickFirmwareVersion(joystick)
}

// [GetJoystickSerial] gets the serial number of an opened joystick, if available.
//
// [GetJoystickSerial]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial
func GetJoystickSerial(joystick *Joystick) string {
	return sdlGetJoystickSerial(joystick)
}

// [GetJoystickType] gets the type of an opened joystick.
//
// [GetJoystickType]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickType
func GetJoystickType(joystick *Joystick) JoystickType {
	return sdlGetJoystickType(joystick)
}

// func GetJoystickGUIDInfo(guid GUID, vendor *uint16, product *uint16, version *uint16, crc16 *uint16) {
// 	sdlGetJoystickGUIDInfo(guid, vendor, product, version, crc16)
// }

// [JoystickConnected] gets the status of a specified joystick.
//
// [JoystickConnected]: https://wiki.libsdl.org/SDL3/SDL_JoystickConnected
func JoystickConnected(joystick *Joystick) bool {
	return sdlJoystickConnected(joystick)
}

// [GetJoystickID] gets the instance ID of an opened joystick.
//
// [GetJoystickID]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickID
func GetJoystickID(joystick *Joystick) JoystickID {
	return sdlGetJoystickID(joystick)
}

// [GetNumJoystickAxes] gets the number of general axis controls on a joystick.
//
// [GetNumJoystickAxes]: https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes
func GetNumJoystickAxes(joystick *Joystick) int32 {
	return sdlGetNumJoystickAxes(joystick)
}

// [GetNumJoystickBalls] gets the number of trackballs on a joystick.
//
// [GetNumJoystickBalls]: https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls
func GetNumJoystickBalls(joystick *Joystick) int32 {
	return sdlGetNumJoystickBalls(joystick)
}

// [GetNumJoystickHats] gets the number of POV hats on a joystick.
//
// [GetNumJoystickHats]: https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats
func GetNumJoystickHats(joystick *Joystick) int32 {
	return sdlGetNumJoystickHats(joystick)
}

// [GetNumJoystickButtons] gets the number of buttons on a joystick.
//
// [GetNumJoystickButtons]: https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons
func GetNumJoystickButtons(joystick *Joystick) int32 {
	return sdlGetNumJoystickButtons(joystick)
}

// [SetJoystickEventsEnabled] sets the state of joystick event processing.
//
// [SetJoystickEventsEnabled]: https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled
func SetJoystickEventsEnabled(enabled bool) {
	sdlSetJoystickEventsEnabled(enabled)
}

// [JoystickEventsEnabled] queries the state of joystick event processing.
//
// [JoystickEventsEnabled]: https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled
func JoystickEventsEnabled() bool {
	return sdlJoystickEventsEnabled()
}

// [UpdateJoysticks] updates the current state of the open joysticks.
//
// [UpdateJoysticks]: https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks
func UpdateJoysticks() {
	sdlUpdateJoysticks()
}

// [GetJoystickAxis] gets the current state of an axis control on a joystick.
//
// [GetJoystickAxis]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis
func GetJoystickAxis(joystick *Joystick, axis int32) int16 {
	return sdlGetJoystickAxis(joystick, axis)
}

// [GetJoystickAxisInitialState] gets the initial state of an axis control on a joystick.
//
// [GetJoystickAxisInitialState]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState
func GetJoystickAxisInitialState(joystick *Joystick, axis int32, state *int16) bool {
	return sdlGetJoystickAxisInitialState(joystick, axis, state)
}

// [GetJoystickBall] gets the ball axis change since the last poll.
//
// [GetJoystickBall]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall
func GetJoystickBall(joystick *Joystick, ball int32, dx *int32, dy *int32) bool {
	return sdlGetJoystickBall(joystick, ball, dx, dy)
}

// [GetJoystickHat] gets the current state of a POV hat on a joystick.
//
// [GetJoystickHat]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat
func GetJoystickHat(joystick *Joystick, hat int32) uint8 {
	return sdlGetJoystickHat(joystick, hat)
}

const (
	HatCentered  = 0x00
	HatUp        = 0x01
	HatRight     = 0x02
	HatDown      = 0x04
	HatLeft      = 0x08
	HatRightUp   = HatRight | HatUp
	HatRightDown = HatRight | HatDown
	HatLeftUp    = HatLeft | HatUp
	HatLeftDown  = HatLeft | HatDown
)

// [GetJoystickButton] gets the current state of a button on a joystick.
//
// [GetJoystickButton]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton
func GetJoystickButton(joystick *Joystick, button int32) bool {
	return sdlGetJoystickButton(joystick, button)
}

// [RumbleJoystick] starts a rumble effect.
//
// [RumbleJoystick]: https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick
func RumbleJoystick(joystick *Joystick, lowFrequencyRumble uint16, highFrequencyRumble uint16, durationMs uint32) bool {
	return sdlRumbleJoystick(joystick, lowFrequencyRumble, highFrequencyRumble, durationMs)
}

// [RumbleJoystickTriggers] starts a rumble effect in the joystick's triggers.
//
// [RumbleJoystickTriggers]: https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers
func RumbleJoystickTriggers(joystick *Joystick, leftRumble uint16, rightRumble uint16, durationMs uint32) bool {
	return sdlRumbleJoystickTriggers(joystick, leftRumble, rightRumble, durationMs)
}

// [SetJoystickLED] updates a joystick's LED color.
//
// [SetJoystickLED]: https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED
func SetJoystickLED(joystick *Joystick, red uint8, green uint8, blue uint8) bool {
	return sdlSetJoystickLED(joystick, red, green, blue)
}

// [SendJoystickEffect] sends a joystick specific effect packet.
//
// [SendJoystickEffect]: https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect
func SendJoystickEffect(joystick *Joystick, data unsafe.Pointer, size int32) bool {
	return sdlSendJoystickEffect(joystick, data, size)
}

// [CloseJoystick] closes a joystick previously opened with [OpenJoystick].
//
// [CloseJoystick]: https://wiki.libsdl.org/SDL3/SDL_CloseJoystick
func CloseJoystick(joystick *Joystick) {
	sdlCloseJoystick(joystick)
}

// [GetJoystickConnectionState] gets the connection state of a joystick.
//
// [GetJoystickConnectionState]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState
func GetJoystickConnectionState(joystick *Joystick) JoystickConnectionState {
	return sdlGetJoystickConnectionState(joystick)
}

// [GetJoystickPowerInfo] gets the battery state of a joystick.
//
// [GetJoystickPowerInfo]: https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo
func GetJoystickPowerInfo(joystick *Joystick, percent *int32) PowerState {
	return sdlGetJoystickPowerInfo(joystick, percent)
}
