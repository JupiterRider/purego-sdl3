package sdl

// [AppResult] defines the return values for optional main callbacks.
//
// [AppResult]: https://wiki.libsdl.org/SDL3/SDL_AppResult
type AppResult uint32

const (
	AppContinue AppResult = iota // Value that requests that the app continue from the main callbacks.
	AppSuccess                   // Value that requests termination with success from the main callbacks.
	AppFailure                   // Value that requests termination with error from the main callbacks.
)

// [InitFlags] defines the initialization flags for [Init] and/or [InitSubSystem].
//
// [InitFlags]: https://wiki.libsdl.org/SDL3/SDL_InitFlags
type InitFlags uint32

const (
	InitAudio    InitFlags = 0x00000010 // [INIT_AUDIO] implies [INIT_EVENTS].
	InitVideo    InitFlags = 0x00000020 // [INIT_VIDEO] implies [INIT_EVENTS], should be initialized on the main thread.
	InitJoystick InitFlags = 0x00000200 // [INIT_JOYSTICK] implies [INIT_EVENTS].
	InitHaptic   InitFlags = 0x00001000
	InitGamepad  InitFlags = 0x00002000 // [INIT_GAMEPAD] implies [INIT_JOYSTICK].
	InitEvents   InitFlags = 0x00004000
	InitSensor   InitFlags = 0x00008000 // [INIT_SENSOR] implies [INIT_EVENTS].
	InitCamera   InitFlags = 0x00010000 // [INIT_CAMERA] implies [INIT_EVENTS].
)

// [Init] initializes the SDL library.
//
// [Init]: https://wiki.libsdl.org/SDL3/SDL_Init
func Init(flags InitFlags) bool {
	return sdlInit(flags)
}

// [InitSubSystem] is a compatibility function to initialize the SDL library.
//
// This function and [Init] are interchangeable.
//
// [InitSubSystem]: https://wiki.libsdl.org/SDL3/SDL_InitSubSystem
func InitSubSystem(flags InitFlags) bool {
	return sdlInit(flags)
}

// [Quit] cleans up all initialized subsystems.
//
// [Quit]: https://wiki.libsdl.org/SDL3/SDL_Quit
func Quit() {
	sdlQuit()
}

// [QuitSubSystem] shuts down specific SDL subsystems.
//
// You still need to call [Quit] even if you close all open subsystems with this function.
//
// [QuitSubSystem]: https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem
func QuitSubSystem(flags InitFlags) {
	sdlQuitSubSystem(flags)
}

// [GetAppMetadataProperty] gets metadata about your app.
//
// [GetAppMetadataProperty]: https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty
func GetAppMetadataProperty(name string) string {
	return sdlGetAppMetadataProperty(name)
}

// [IsMainThread] returns whether this is the main thread.
//
// [IsMainThread]: https://wiki.libsdl.org/SDL3/SDL_IsMainThread
func IsMainThread() bool {
	return sdlIsMainThread()
}

// func RunOnMainThread(callback MainThreadCallback, userdata unsafe.Pointer, wait_complete bool) bool {
//	return sdlRunOnMainThread(callback, userdata, wait_complete)
// }

// func SetAppMetadata(appname string, appversion string, appidentifier string) bool {
//	return sdlSetAppMetadata(appname, appversion, appidentifier)
// }

// func SetAppMetadataProperty(name string, value string) bool {
//	return sdlSetAppMetadataProperty(name, value)
// }

// func WasInit(flags InitFlags) InitFlags {
//	return sdlWasInit(flags)
// }
