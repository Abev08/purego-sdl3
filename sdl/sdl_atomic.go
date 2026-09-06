package sdl

// [SpinLock] defines an atomic spinlock.
//
// [SpinLock]: https://wiki.libsdl.org/SDL3/SDL_SpinLock
type SpinLock int32

// [TryLockSpinlock] trys to lock a spin lock by setting it to a non-zero value.
//
// [TryLockSpinlock]: https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock
// func TryLockSpinlock(lock *SpinLock) bool {
//	return sdlTryLockSpinlock(lock)
// }

// [LockSpinlock] locks a spin lock by setting it to a non-zero value.
//
// [LockSpinlock]: https://wiki.libsdl.org/SDL3/SDL_LockSpinlock
// func LockSpinlock(lock *SpinLock) {
//	sdlLockSpinlock(lock)
// }

// [UnlockSpinlock] unlocks a spin lock by setting it to 0.
//
// [UnlockSpinlock]: https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock
// func UnlockSpinlock(lock *SpinLock) {
//	sdlUnlockSpinlock(lock)
// }

// [MemoryBarrierReleaseFunction] inserts a memory release barrier (function version).
//
// [MemoryBarrierReleaseFunction]: https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierReleaseFunction
// func MemoryBarrierReleaseFunction() {
//	sdlMemoryBarrierReleaseFunction()
// }

// [MemoryBarrierAcquireFunction] inserts a memory acquire barrier (function version).
//
// [MemoryBarrierAcquireFunction]: https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierAcquireFunction
// func MemoryBarrierAcquireFunction()  {
//	sdlMemoryBarrierAcquireFunction()
// }

// [AtomicInt] defines a type representing an atomic integer value.
//
// [AtomicInt]: https://wiki.libsdl.org/SDL3/SDL_AtomicInt
type AtomicInt struct {
	Value int32
}

// [CompareAndSwapAtomicInt] sets an atomic variable to a new value if it is currently an old value.
//
// [CompareAndSwapAtomicInt]: https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt
// func CompareAndSwapAtomicInt(a *AtomicInt, oldval int32, newval int32) bool {
//	return sdlCompareAndSwapAtomicInt(a, oldval, newval)
// }

// [SetAtomicInt] sets an atomic variable to a value.
//
// [SetAtomicInt]: https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt
// func SetAtomicInt(a *AtomicInt, v int32) int32 {
//	return sdlSetAtomicInt(a, v)
// }

// [GetAtomicInt] gets the value of an atomic variable.
//
// [GetAtomicInt]: https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt
// func GetAtomicInt(a *AtomicInt) int32 {
//	return sdlGetAtomicInt(a)
// }

// [AddAtomicInt] adds to an atomic variable.
//
// [AddAtomicInt]: https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt
// func AddAtomicInt(a *AtomicInt, v int32) int32 {
//	return sdlAddAtomicInt(a, v)
// }

// [AtomicU32] defines a type representing an atomic unsigned 32-bit value.
//
// [AtomicU32]: https://wiki.libsdl.org/SDL3/SDL_AtomicU32
type AtomicU32 struct {
	Value uint32
}

// [CompareAndSwapAtomicU32] sets an atomic variable to a new value if it is currently an old value.
//
// [CompareAndSwapAtomicU32]: https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32
// func CompareAndSwapAtomicU32(a *AtomicU32, oldval uint32, newval uint32) bool {
//	return sdlCompareAndSwapAtomicU32(a, oldval, newval)
// }

// [SetAtomicU32] sets an atomic variable to a value.
//
// [SetAtomicU32]: https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32
// func SetAtomicU32(a *AtomicU32, v uint32) uint32 {
//	return sdlSetAtomicU32(a, v)
// }

// [GetAtomicU32] gets the value of an atomic variable.
//
// [GetAtomicU32]: https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32
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

// [CompareAndSwapAtomicPointer] sets a pointer to a new value if it is currently an old value.
//
// [CompareAndSwapAtomicPointer]: https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicPointer
// func CompareAndSwapAtomicPointer(a *unsafe.Pointer, oldval unsafe.Pointer, newval unsafe.Pointer) bool {
//	return sdlCompareAndSwapAtomicPointer(a, oldval, newval)
// }

// [SetAtomicPointer] sets a pointer to a value atomically.
//
// [SetAtomicPointer]: https://wiki.libsdl.org/SDL3/SDL_SetAtomicPointer
// func SetAtomicPointer(a *unsafe.Pointer, v unsafe.Pointer) unsafe.Pointer {
//	return sdlSetAtomicPointer(a, v)
// }

// [GetAtomicPointer] gets the value of a pointer atomically.
//
// [GetAtomicPointer]: https://wiki.libsdl.org/SDL3/SDL_GetAtomicPointer
// func GetAtomicPointer(a *unsafe.Pointer) unsafe.Pointer {
//	return sdlGetAtomicPointer(a)
// }
