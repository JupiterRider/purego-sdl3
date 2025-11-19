package sdl

// [PowerState] is the basic state for the system's power supply return by [GetPowerInfo].
//
// [PowerState]: https://wiki.libsdl.org/SDL3/SDL_PowerState
type PowerState int32

const (
	PowerStateError     PowerState = iota - 1 // Error determining power status.
	PowerStateUnknown                         // Cannot determine power status.
	PowerStateOnBattery                       // Not plugged in, running on the battery.
	PowerStateNoBattery                       // Plugged in, no battery available.
	PowerStateCharging                        // Plugged in, charging battery.
	PowerStateCharged                         // Plugged in, battery charged.
)

// [GetPowerInfo] gets the current power supply details.
//
// [GetPowerInfo]: https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo
func GetPowerInfo(seconds *int32, percent *int32) PowerState {
	return sdlGetPowerInfo(seconds, percent)
}
