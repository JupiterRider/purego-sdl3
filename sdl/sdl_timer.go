package sdl

import "github.com/ebitengine/purego"

// func AddTimer(interval uint32, callback TimerCallback, userdata unsafe.Pointer) TimerID {
//	return sdlAddTimer(interval, callback, userdata)
// }

// func AddTimerNS(interval uint64, callback NSTimerCallback, userdata unsafe.Pointer) TimerID {
//	return sdlAddTimerNS(interval, callback, userdata)
// }

// Delay wait a specified number of milliseconds before returning.
// func Delay(ms uint32)  {
//	sdlDelay(ms)
// }

// DelayNS waits a specified number of nanoseconds before returning.
func DelayNS(ns uint64) {
	sdlDelayNS(ns)
}

// DelayPrecise wait a specified number of nanoseconds before returning.
// func DelayPrecise(ns uint64)  {
//	sdlDelayPrecise(ns)
// }

// GetPerformanceCounter returns the current value of the high resolution counter.
func GetPerformanceCounter() uint64 {
	ret, _, _ := purego.SyscallN(sdlGetPerformanceCounter)
	return uint64(ret)
}

// GetPerformanceFrequency returns the count per second of the high resolution counter.
func GetPerformanceFrequency() uint64 {
	ret, _, _ := purego.SyscallN(sdlGetPerformanceFrequency)
	return uint64(ret)
}

// GetTicks returns the number of milliseconds that have elapsed since the SDL
// library initialization.
func GetTicks() uint64 {
	ret, _, _ := purego.SyscallN(sdlGetTicks)
	return uint64(ret)
}

// GetTicksNS returns the number of nanoseconds since SDL library initialization.
func GetTicksNS() uint64 {
	ret, _, _ := purego.SyscallN(sdlGetTicksNS)
	return uint64(ret)
}

// RemoveTimer remove a timer created with SDL_AddTimer().
// func RemoveTimer(id TimerID) bool {
//	return sdlRemoveTimer(id)
// }
