package sdl

// [Sandbox] defines the application sandbox environment.
//
// [Sandbox]: https://wiki.libsdl.org/SDL3/SDL_Sandbox
type Sandbox uint32

const (
	SandboxNone Sandbox = iota
	SandboxUnknownContainer
	SandboxFlatpak
	SandboxSnap
	SandboxMacOS
)

// func GetSandbox() Sandbox {
//	return sdlGetSandbox()
// }

// func IsTablet() bool {
//	return sdlIsTablet()
// }

// func IsTV() bool {
//	return sdlIsTV()
// }

// func OnApplicationDidEnterBackground()  {
//	sdlOnApplicationDidEnterBackground()
// }

// func OnApplicationDidEnterForeground()  {
//	sdlOnApplicationDidEnterForeground()
// }

// func OnApplicationDidReceiveMemoryWarning()  {
//	sdlOnApplicationDidReceiveMemoryWarning()
// }

// func OnApplicationWillEnterBackground()  {
//	sdlOnApplicationWillEnterBackground()
// }

// func OnApplicationWillEnterForeground()  {
//	sdlOnApplicationWillEnterForeground()
// }

// func OnApplicationWillTerminate()  {
//	sdlOnApplicationWillTerminate()
// }

// func SetLinuxThreadPriority(threadID int64, priority int32) bool {
//	return sdlSetLinuxThreadPriority(threadID, priority)
// }

// func SetLinuxThreadPriorityAndPolicy(threadID int64, sdlPriority int32, schedPolicy int32) bool {
//	return sdlSetLinuxThreadPriorityAndPolicy(threadID, sdlPriority, schedPolicy)
// }

// func SetX11EventHook(callback X11EventHook, userdata unsafe.Pointer)  {
//	sdlSetX11EventHook(callback, userdata)
// }
