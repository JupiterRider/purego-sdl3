package sdl

// [SpinLock] is an atomic spinlock.
//
// [SpinLock]: https://wiki.libsdl.org/SDL3/SDL_SpinLock
type SpinLock int32

// [AtomicInt] is a structure representing an atomic integer value.
//
// [AtomicInt]: https://wiki.libsdl.org/SDL3/SDL_AtomicInt
type AtomicInt struct {
	Value int32
}

// [AtomicU32] is a structure representing an atomic unsigned 32-bit value.
//
// [AtomicU32]: https://wiki.libsdl.org/SDL3/SDL_AtomicU32
type AtomicU32 struct {
	Value uint32
}

// func AddAtomicInt(a *AtomicInt, v int32) int32 {
//	return sdlAddAtomicInt(a, v)
// }

// func CompareAndSwapAtomicInt(a *AtomicInt, oldval int32, newval int32) bool {
//	return sdlCompareAndSwapAtomicInt(a, oldval, newval)
// }

// func CompareAndSwapAtomicPointer(a *unsafe.Pointer, oldval unsafe.Pointer, newval unsafe.Pointer) bool {
//	return sdlCompareAndSwapAtomicPointer(a, oldval, newval)
// }

// func CompareAndSwapAtomicU32(a *AtomicU32, oldval uint32, newval uint32) bool {
//	return sdlCompareAndSwapAtomicU32(a, oldval, newval)
// }

// func GetAtomicInt(a *AtomicInt) int32 {
//	return sdlGetAtomicInt(a)
// }

// func GetAtomicPointer(a *unsafe.Pointer) unsafe.Pointer {
//	return sdlGetAtomicPointer(a)
// }

// func GetAtomicU32(a *AtomicU32) uint32 {
//	return sdlGetAtomicU32(a)
// }

// [AddAtomicU32] adds to an atomic variable.
//
// Available since SDL 3.4.0.
//
// [AddAtomicU32]: https://wiki.libsdl.org/SDL3/SDL_AddAtomicU32
// func AddAtomicU32(a *AtomicU32, v int32) uint32 {
// 	return sdlAddAtomicU32(a, v)
// }

// func LockSpinlock(lock *SpinLock)  {
//	sdlLockSpinlock(lock)
// }

// func MemoryBarrierAcquireFunction()  {
//	sdlMemoryBarrierAcquireFunction()
// }

// func MemoryBarrierReleaseFunction()  {
//	sdlMemoryBarrierReleaseFunction()
// }

// func SetAtomicInt(a *AtomicInt, v int32) int32 {
//	return sdlSetAtomicInt(a, v)
// }

// func SetAtomicPointer(a *unsafe.Pointer, v unsafe.Pointer) unsafe.Pointer {
//	return sdlSetAtomicPointer(a, v)
// }

// func SetAtomicU32(a *AtomicU32, v uint32) uint32 {
//	return sdlSetAtomicU32(a, v)
// }

// func TryLockSpinlock(lock *SpinLock) bool {
//	return sdlTryLockSpinlock(lock)
// }

// func UnlockSpinlock(lock *SpinLock)  {
//	sdlUnlockSpinlock(lock)
// }
