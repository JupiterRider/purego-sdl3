package sdl

// [ThreadPriority] is a structure specifying the SDL thread prioritys.
//
// [ThreadPriority]: https://wiki.libsdl.org/SDL3/SDL_ThreadPriority
type ThreadPriority uint32

const (
	ThreadPriorityLow ThreadPriority = iota
	ThreadPriorityNormal
	ThreadPriorityHigh
	ThreadPriorityTimeCritical
)

// [ThreadState] is a structure specifying the SDL thread states.
//
// [ThreadState]: https://wiki.libsdl.org/SDL3/SDL_ThreadState
type ThreadState uint32

const (
	ThreadUnknown  ThreadState = iota // The thread is not valid.
	ThreadAlive                       // The thread is currently running.
	ThreadDetached                    // The thread is detached and can't be waited on.
	ThreadComplete                    // The thread has finished and should be cleaned up with [WaitThread].
)

// func CleanupTLS()  {
//	sdlCleanupTLS()
// }

// func CreateThreadRuntime(fn ThreadFunction, name string, data unsafe.Pointer, pfnBeginThread FunctionPointer, pfnEndThread FunctionPointer) *Thread {
//	return sdlCreateThreadRuntime(fn, name, data, pfnBeginThread, pfnEndThread)
// }

// func CreateThreadWithPropertiesRuntime(props PropertiesID, pfnBeginThread FunctionPointer, pfnEndThread FunctionPointer) *Thread {
//	return sdlCreateThreadWithPropertiesRuntime(props, pfnBeginThread, pfnEndThread)
// }

// func DetachThread(thread *Thread)  {
//	sdlDetachThread(thread)
// }

// func GetCurrentThreadID() ThreadID {
//	return sdlGetCurrentThreadID()
// }

// func GetThreadID(thread *Thread) ThreadID {
//	return sdlGetThreadID(thread)
// }

// func GetThreadName(thread *Thread) string {
//	return sdlGetThreadName(thread)
// }

// func GetThreadState(thread *Thread) ThreadState {
//	return sdlGetThreadState(thread)
// }

// func GetTLS(id *TLSID) unsafe.Pointer {
//	return sdlGetTLS(id)
// }

// func SetCurrentThreadPriority(priority ThreadPriority) bool {
//	return sdlSetCurrentThreadPriority(priority)
// }

// func SetTLS(id *TLSID, value unsafe.Pointer, destructor TLSDestructorCallback) bool {
//	return sdlSetTLS(id, value, destructor)
// }

// func WaitThread(thread *Thread, status *int32)  {
//	sdlWaitThread(thread, status)
// }
