package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

type GamepadBindingType uint32

const (
	GamepadBindTypeNone GamepadBindingType = iota
	GamepadBindTypeButton
	GamepadBindTypeAxis
	GamepadBindTypeHat
)

type GamepadAxis int32

const (
	GamepadAxisInvalid GamepadAxis = iota - 1
	GamepadAxisLeftX
	GamepadAxisLeftY
	GamepadAxisRightX
	GamepadAxisRightY
	GamepadAxisLeftTrigger
	GamepadAxisRightTrigger
	GamepadAxisCount
)

type GamepadButtonLabel uint32

const (
	GamepadButtonLabelUnknown GamepadButtonLabel = iota
	GamepadButtonLabelA
	GamepadButtonLabelB
	GamepadButtonLabelX
	GamepadButtonLabelY
	GamepadButtonLabelCross
	GamepadButtonLabelCircle
	GamepadButtonLabelSquare
	GamepadButtonLabelTriangle
)

type GamepadButton int32

const (
	GamepadButtonInvalid GamepadButton = iota - 1
	GamepadButtonSouth
	GamepadButtonEast
	GamepadButtonWest
	GamepadButtonNorth
	GamepadButtonBack
	GamepadButtonGuide
	GamepadButtonStart
	GamepadButtonLeftStick
	GamepadButtonRightStick
	GamepadButtonLeftShoulder
	GamepadButtonRightShoulder
	GamepadButtonDpadUp
	GamepadButtonDpadDown
	GamepadButtonDpadLeft
	GamepadButtonDpadRight
	GamepadButtonMisc1
	GamepadButtonRightPaddle1
	GamepadButtonLeftPaddle1
	GamepadButtonRightPaddle2
	GamepadButtonLeftPaddle2
	GamepadButtonTouchpad
	GamepadButtonMisc2
	GamepadButtonMisc3
	GamepadButtonMisc4
	GamepadButtonMisc5
	GamepadButtonMisc6
	GamepadButtonCount
)

type GamepadType uint32

const (
	GamepadTypeUnknown GamepadType = iota
	GamepadTypeStandard
	GamepadTypeXbox360
	GamepadTypeXboxOne
	GamepadTypePS3
	GamepadTypePS4
	GamepadTypePS5
	GamepadTypeNintendoSwitchPro
	GamepadTypeNintendoSwitchJoyConLeft
	GamepadTypeNintendoSwitchJoyConRight
	GamepadTypeNintendoSwitchJoyConPair
	GamepadTypeCount
)

type Gamepad struct{}

type GamepadBinding struct {
	InputType  GamepadBindingType
	input      [12]byte
	OutputType GamepadBindingType
	output     [12]byte
}

func (g *GamepadBinding) InputButton() int32 {
	return *(*int32)(unsafe.Pointer(&g.input))
}

func (g *GamepadBinding) InputAxis() struct{ Axis, AxisMin, AxisMax int32 } {
	return *(*struct{ Axis, AxisMin, AxisMax int32 })(unsafe.Pointer(&g.input))
}

func (g *GamepadBinding) InputHat() struct{ Hat, HatMask int32 } {
	return *(*struct{ Hat, HatMask int32 })(unsafe.Pointer(&g.input))
}

func (g *GamepadBinding) OutputButton() GamepadButton {
	return *(*GamepadButton)(unsafe.Pointer(&g.output))
}

func (g *GamepadBinding) OutputAxis() struct {
	Axis             GamepadAxis
	AxisMin, AxisMax int32
} {
	return *(*struct {
		Axis             GamepadAxis
		AxisMin, AxisMax int32
	})(unsafe.Pointer(&g.output))
}

// func AddGamepadMapping(mapping string) int32 {
//	return sdlAddGamepadMapping(mapping)
// }

// func AddGamepadMappingsFromFile(file string) int32 {
//	return sdlAddGamepadMappingsFromFile(file)
// }

// func AddGamepadMappingsFromIO(src *IOStream, closeio bool) int32 {
//	return sdlAddGamepadMappingsFromIO(src, closeio)
// }

func CloseGamepad(gamepad *Gamepad) {
	sdlCloseGamepad(gamepad)
}

// func GamepadConnected(gamepad *Gamepad) bool {
//	return sdlGamepadConnected(gamepad)
// }

// func GamepadEventsEnabled() bool {
//	return sdlGamepadEventsEnabled()
// }

// func GamepadHasAxis(gamepad *Gamepad, axis GamepadAxis) bool {
//	return sdlGamepadHasAxis(gamepad, axis)
// }

// func GamepadHasButton(gamepad *Gamepad, button GamepadButton) bool {
//	return sdlGamepadHasButton(gamepad, button)
// }

// func GamepadHasSensor(gamepad *Gamepad, type SensorType) bool {
//	return sdlGamepadHasSensor(gamepad, type)
// }

// func GamepadSensorEnabled(gamepad *Gamepad, type SensorType) bool {
//	return sdlGamepadSensorEnabled(gamepad, type)
// }

// func GetGamepadAppleSFSymbolsNameForAxis(gamepad *Gamepad, axis GamepadAxis) string {
//	return sdlGetGamepadAppleSFSymbolsNameForAxis(gamepad, axis)
// }

// func GetGamepadAppleSFSymbolsNameForButton(gamepad *Gamepad, button GamepadButton) string {
//	return sdlGetGamepadAppleSFSymbolsNameForButton(gamepad, button)
// }

// func GetGamepadAxis(gamepad *Gamepad, axis GamepadAxis) int16 {
//	return sdlGetGamepadAxis(gamepad, axis)
// }

// func GetGamepadAxisFromString(str string) GamepadAxis {
//	return sdlGetGamepadAxisFromString(str)
// }

// GetGamepadBindings returns the SDL joystick layer bindings for a gamepad or nil on failure.
//
// The returned slice must not be freed with e.g. [Free].
func GetGamepadBindings(gamepad *Gamepad) []*GamepadBinding {
	var count int32
	bindings := sdlGetGamepadBindings(gamepad, &count)
	if bindings == nil {
		return nil
	}
	defer Free(unsafe.Pointer(bindings))
	return mem.DeepCopy(bindings, count)
}

// func GetGamepadButton(gamepad *Gamepad, button GamepadButton) bool {
//	return sdlGetGamepadButton(gamepad, button)
// }

// func GetGamepadButtonFromString(str string) GamepadButton {
//	return sdlGetGamepadButtonFromString(str)
// }

// func GetGamepadButtonLabel(gamepad *Gamepad, button GamepadButton) GamepadButtonLabel {
//	return sdlGetGamepadButtonLabel(gamepad, button)
// }

// func GetGamepadButtonLabelForType(type GamepadType, button GamepadButton) GamepadButtonLabel {
//	return sdlGetGamepadButtonLabelForType(type, button)
// }

// func GetGamepadConnectionState(gamepad *Gamepad) JoystickConnectionState {
//	return sdlGetGamepadConnectionState(gamepad)
// }

// func GetGamepadFirmwareVersion(gamepad *Gamepad) uint16 {
//	return sdlGetGamepadFirmwareVersion(gamepad)
// }

func GetGamepadFromID(instanceId JoystickID) *Gamepad {
	return sdlGetGamepadFromID(instanceId)
}

// func GetGamepadFromPlayerIndex(player_index int32) *Gamepad {
//	return sdlGetGamepadFromPlayerIndex(player_index)
// }

// func GetGamepadGUIDForID(instance_id JoystickID) GUID {
//	return sdlGetGamepadGUIDForID(instance_id)
// }

// func GetGamepadID(gamepad *Gamepad) JoystickID {
//	return sdlGetGamepadID(gamepad)
// }

// func GetGamepadJoystick(gamepad *Gamepad) *Joystick {
//	return sdlGetGamepadJoystick(gamepad)
// }

// func GetGamepadMapping(gamepad *Gamepad) string {
//	return sdlGetGamepadMapping(gamepad)
// }

// func GetGamepadMappingForGUID(guid GUID) string {
//	return sdlGetGamepadMappingForGUID(guid)
// }

// func GetGamepadMappingForID(instance_id JoystickID) string {
//	return sdlGetGamepadMappingForID(instance_id)
// }

// func GetGamepadMappings(count *int32) **byte {
//	return sdlGetGamepadMappings(count)
// }

func GetGamepadName(gamepad *Gamepad) string {
	return sdlGetGamepadName(gamepad)
}

func GetGamepadNameForID(instanceId JoystickID) string {
	return sdlGetGamepadNameForID(instanceId)
}

// func GetGamepadPath(gamepad *Gamepad) string {
//	return sdlGetGamepadPath(gamepad)
// }

// func GetGamepadPathForID(instance_id JoystickID) string {
//	return sdlGetGamepadPathForID(instance_id)
// }

// func GetGamepadPlayerIndex(gamepad *Gamepad) int32 {
//	return sdlGetGamepadPlayerIndex(gamepad)
// }

// func GetGamepadPlayerIndexForID(instance_id JoystickID) int32 {
//	return sdlGetGamepadPlayerIndexForID(instance_id)
// }

// func GetGamepadPowerInfo(gamepad *Gamepad, percent *int32) PowerState {
//	return sdlGetGamepadPowerInfo(gamepad, percent)
// }

// func GetGamepadProduct(gamepad *Gamepad) uint16 {
//	return sdlGetGamepadProduct(gamepad)
// }

// func GetGamepadProductForID(instance_id JoystickID) uint16 {
//	return sdlGetGamepadProductForID(instance_id)
// }

// func GetGamepadProductVersion(gamepad *Gamepad) uint16 {
//	return sdlGetGamepadProductVersion(gamepad)
// }

// func GetGamepadProductVersionForID(instance_id JoystickID) uint16 {
//	return sdlGetGamepadProductVersionForID(instance_id)
// }

// func GetGamepadProperties(gamepad *Gamepad) PropertiesID {
//	return sdlGetGamepadProperties(gamepad)
// }

// GetGamepads returns a list of currently connected gamepads or nil on failure.
//
// The returned slice must not be freed with e.g. [Free].
func GetGamepads() []JoystickID {
	var count int32
	gamepads := sdlGetGamepads(&count)
	defer Free(unsafe.Pointer(gamepads))
	return mem.Copy(gamepads, count)
}

// func GetGamepadSensorData(gamepad *Gamepad, type SensorType, data *float32, num_values int32) bool {
//	return sdlGetGamepadSensorData(gamepad, type, data, num_values)
// }

// func GetGamepadSensorDataRate(gamepad *Gamepad, type SensorType) float32 {
//	return sdlGetGamepadSensorDataRate(gamepad, type)
// }

// GetGamepadSerial returns the serial number of an opened gamepad, or "" if unavailable.
func GetGamepadSerial(gamepad *Gamepad) string {
	return sdlGetGamepadSerial(gamepad)
}

// func GetGamepadSteamHandle(gamepad *Gamepad) uint64 {
//	return sdlGetGamepadSteamHandle(gamepad)
// }

// func GetGamepadStringForAxis(axis GamepadAxis) string {
//	return sdlGetGamepadStringForAxis(axis)
// }

// GetGamepadStringForButton returns the name for the given button.
func GetGamepadStringForButton(button GamepadButton) string {
	return sdlGetGamepadStringForButton(button)
}

func GetGamepadStringForType(gamepadType GamepadType) string {
	return sdlGetGamepadStringForType(gamepadType)
}

// func GetGamepadTouchpadFinger(gamepad *Gamepad, touchpad int32, finger int32, down *bool, x *float32, y *float32, pressure *float32) bool {
//	return sdlGetGamepadTouchpadFinger(gamepad, touchpad, finger, down, x, y, pressure)
// }

func GetGamepadType(gamepad *Gamepad) GamepadType {
	return sdlGetGamepadType(gamepad)
}

// func GetGamepadTypeForID(instance_id JoystickID) GamepadType {
//	return sdlGetGamepadTypeForID(instance_id)
// }

// func GetGamepadTypeFromString(str string) GamepadType {
//	return sdlGetGamepadTypeFromString(str)
// }

// func GetGamepadVendor(gamepad *Gamepad) uint16 {
//	return sdlGetGamepadVendor(gamepad)
// }

// func GetGamepadVendorForID(instance_id JoystickID) uint16 {
//	return sdlGetGamepadVendorForID(instance_id)
// }

// func GetNumGamepadTouchpadFingers(gamepad *Gamepad, touchpad int32) int32 {
//	return sdlGetNumGamepadTouchpadFingers(gamepad, touchpad)
// }

// func GetNumGamepadTouchpads(gamepad *Gamepad) int32 {
//	return sdlGetNumGamepadTouchpads(gamepad)
// }

// func GetRealGamepadType(gamepad *Gamepad) GamepadType {
//	return sdlGetRealGamepadType(gamepad)
// }

// func GetRealGamepadTypeForID(instance_id JoystickID) GamepadType {
//	return sdlGetRealGamepadTypeForID(instance_id)
// }

// func HasGamepad() bool {
//	return sdlHasGamepad()
// }

// func IsGamepad(instance_id JoystickID) bool {
//	return sdlIsGamepad(instance_id)
// }

func OpenGamepad(instanceId JoystickID) *Gamepad {
	return sdlOpenGamepad(instanceId)
}

// func ReloadGamepadMappings() bool {
//	return sdlReloadGamepadMappings()
// }

// func RumbleGamepad(gamepad *Gamepad, low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) bool {
//	return sdlRumbleGamepad(gamepad, low_frequency_rumble, high_frequency_rumble, duration_ms)
// }

// func RumbleGamepadTriggers(gamepad *Gamepad, left_rumble uint16, right_rumble uint16, duration_ms uint32) bool {
//	return sdlRumbleGamepadTriggers(gamepad, left_rumble, right_rumble, duration_ms)
// }

// func SendGamepadEffect(gamepad *Gamepad, data unsafe.Pointer, size int32) bool {
//	return sdlSendGamepadEffect(gamepad, data, size)
// }

// func SetGamepadEventsEnabled(enabled bool)  {
//	sdlSetGamepadEventsEnabled(enabled)
// }

// func SetGamepadLED(gamepad *Gamepad, red uint8, green uint8, blue uint8) bool {
//	return sdlSetGamepadLED(gamepad, red, green, blue)
// }

// func SetGamepadMapping(instance_id JoystickID, mapping string) bool {
//	return sdlSetGamepadMapping(instance_id, mapping)
// }

// func SetGamepadPlayerIndex(gamepad *Gamepad, player_index int32) bool {
//	return sdlSetGamepadPlayerIndex(gamepad, player_index)
// }

// func SetGamepadSensorEnabled(gamepad *Gamepad, type SensorType, enabled bool) bool {
//	return sdlSetGamepadSensorEnabled(gamepad, type, enabled)
// }

// func UpdateGamepads()  {
//	sdlUpdateGamepads()
// }
