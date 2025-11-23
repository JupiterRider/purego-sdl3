package sdl

// [InitStatus] is a structure specifying the current status of an [InitState] structures.
//
// [InitStatus]: https://wiki.libsdl.org/SDL3/SDL_InitStatus
type InitStatus uint32

const (
	InitStatusUninitialized InitStatus = iota
	InitStatusInitializing
	InitStatusInitialized
	InitStatusUninitializing
)

// func BroadcastCondition(cond *Condition)  {
//	sdlBroadcastCondition(cond)
// }

// func CreateCondition() *Condition {
//	return sdlCreateCondition()
// }

// func CreateMutex() *Mutex {
//	return sdlCreateMutex()
// }

// func CreateRWLock() *RWLock {
//	return sdlCreateRWLock()
// }

// func CreateSemaphore(initial_value uint32) *Semaphore {
//	return sdlCreateSemaphore(initial_value)
// }

// func DestroyCondition(cond *Condition)  {
//	sdlDestroyCondition(cond)
// }

// func DestroyMutex(mutex *Mutex)  {
//	sdlDestroyMutex(mutex)
// }

// func DestroyRWLock(rwlock *RWLock)  {
//	sdlDestroyRWLock(rwlock)
// }

// func DestroySemaphore(sem *Semaphore)  {
//	sdlDestroySemaphore(sem)
// }

// func GetSemaphoreValue(sem *Semaphore) uint32 {
//	return sdlGetSemaphoreValue(sem)
// }

// func LockMutex(mutex *Mutex)  {
//	sdlLockMutex(mutex)
// }

// func LockRWLockForReading(rwlock *RWLock)  {
//	sdlLockRWLockForReading(rwlock)
// }

// func LockRWLockForWriting(rwlock *RWLock)  {
//	sdlLockRWLockForWriting(rwlock)
// }

// func SetInitialized(state *InitState, initialized bool)  {
//	sdlSetInitialized(state, initialized)
// }

// func ShouldInit(state *InitState) bool {
//	return sdlShouldInit(state)
// }

// func ShouldQuit(state *InitState) bool {
//	return sdlShouldQuit(state)
// }

// func SignalCondition(cond *Condition)  {
//	sdlSignalCondition(cond)
// }

// func SignalSemaphore(sem *Semaphore)  {
//	sdlSignalSemaphore(sem)
// }

// func TryLockMutex(mutex *Mutex) bool {
//	return sdlTryLockMutex(mutex)
// }

// func TryLockRWLockForReading(rwlock *RWLock) bool {
//	return sdlTryLockRWLockForReading(rwlock)
// }

// func TryLockRWLockForWriting(rwlock *RWLock) bool {
//	return sdlTryLockRWLockForWriting(rwlock)
// }

// func TryWaitSemaphore(sem *Semaphore) bool {
//	return sdlTryWaitSemaphore(sem)
// }

// func UnlockMutex(mutex *Mutex)  {
//	sdlUnlockMutex(mutex)
// }

// func UnlockRWLock(rwlock *RWLock)  {
//	sdlUnlockRWLock(rwlock)
// }

// func WaitCondition(cond *Condition, mutex *Mutex)  {
//	sdlWaitCondition(cond, mutex)
// }

// func WaitConditionTimeout(cond *Condition, mutex *Mutex, timeoutMS int32) bool {
//	return sdlWaitConditionTimeout(cond, mutex, timeoutMS)
// }

// func WaitSemaphore(sem *Semaphore)  {
//	sdlWaitSemaphore(sem)
// }

// func WaitSemaphoreTimeout(sem *Semaphore, timeoutMS int32) bool {
//	return sdlWaitSemaphoreTimeout(sem, timeoutMS)
// }
