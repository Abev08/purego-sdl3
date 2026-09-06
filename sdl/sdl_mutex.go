package sdl

// [Mutex] defines a means to serialize access to a resource between threads.
//
// [Mutex]: https://wiki.libsdl.org/SDL3/SDL_Mutex
type Mutex struct{}

// [CreateMutex] creates a new mutex.
//
// [CreateMutex]: https://wiki.libsdl.org/SDL3/SDL_CreateMutex
// func CreateMutex() *Mutex {
//	return sdlCreateMutex()
// }

// [LockMutex] locks the mutex.
//
// [LockMutex]: https://wiki.libsdl.org/SDL3/SDL_LockMutex
// func LockMutex(mutex *Mutex)  {
//	sdlLockMutex(mutex)
// }

// [TryLockMutex] tries to lock a mutex without blocking.
//
// [TryLockMutex]: https://wiki.libsdl.org/SDL3/SDL_TryLockMutex
// func TryLockMutex(mutex *Mutex) bool {
//	return sdlTryLockMutex(mutex)
// }

// [UnlockMutex] unlocks the mutex.
//
// [UnlockMutex]: https://wiki.libsdl.org/SDL3/SDL_UnlockMutex
// func UnlockMutex(mutex *Mutex)  {
//	sdlUnlockMutex(mutex)
// }

// [DestroyMutex] destroys a mutex created with [CreateMutex].
//
// [DestroyMutex]: https://wiki.libsdl.org/SDL3/SDL_DestroyMutex
// func DestroyMutex(mutex *Mutex)  {
//	sdlDestroyMutex(mutex)
// }

// [RWLock] is a mutex that allows read-only threads to run in parallel.
//
// [RWLock]: https://wiki.libsdl.org/SDL3/SDL_RWLock
type RWLock struct{}

// [CreateRWLock] creates a new read/write lock.
//
// [CreateRWLock]: https://wiki.libsdl.org/SDL3/SDL_CreateRWLock
// func CreateRWLock() *RWLock {
//	return sdlCreateRWLock()
// }

// [LockRWLockForReading] locks the read/write lock for _read only_ operations.
//
// [LockRWLockForReading]: https://wiki.libsdl.org/SDL3/SDL_LockRWLockForReading
// func LockRWLockForReading(rwlock *RWLock)  {
//	sdlLockRWLockForReading(rwlock)
// }

// [LockRWLockForWriting] locks the read/write lock for _write_ operations.
//
// [LockRWLockForWriting]: https://wiki.libsdl.org/SDL3/SDL_LockRWLockForWriting
// func LockRWLockForWriting(rwlock *RWLock)  {
//	sdlLockRWLockForWriting(rwlock)
// }

// [TryLockRWLockForReading] tries to lock a read/write lock _for reading_ without blocking.
//
// [TryLockRWLockForReading]: https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForReading
// func TryLockRWLockForReading(rwlock *RWLock) bool {
//	return sdlTryLockRWLockForReading(rwlock)
// }

// [TryLockRWLockForWriting] tries to lock a read/write lock _for writing_ without blocking.
//
// [TryLockRWLockForWriting]: https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForWriting
// func TryLockRWLockForWriting(rwlock *RWLock) bool {
//	return sdlTryLockRWLockForWriting(rwlock)
// }

// [UnlockRWLock] unlocks the read/write lock.
//
// [UnlockRWLock]: https://wiki.libsdl.org/SDL3/SDL_UnlockRWLock
// func UnlockRWLock(rwlock *RWLock)  {
//	sdlUnlockRWLock(rwlock)
// }

// [DestroyRWLock] destroys a read/write lock created with [CreateRWLock].
//
// [DestroyRWLock]: https://wiki.libsdl.org/SDL3/SDL_DestroyRWLock
// func DestroyRWLock(rwlock *RWLock)  {
//	sdlDestroyRWLock(rwlock)
// }

// [Semaphore] defines a means to manage access to a resource, by count, between threads.
//
// [Semaphore]: https://wiki.libsdl.org/SDL3/SDL_Semaphore
type Semaphore struct{}

// [CreateSemaphore] creates a semaphore.
//
// [CreateSemaphore]: https://wiki.libsdl.org/SDL3/SDL_CreateSemaphore
// func CreateSemaphore(initial_value uint32) *Semaphore {
//	return sdlCreateSemaphore(initial_value)
// }

// [DestroySemaphore] destroys a semaphore.
//
// [DestroySemaphore]: https://wiki.libsdl.org/SDL3/SDL_DestroySemaphore
// func DestroySemaphore(sem *Semaphore)  {
//	sdlDestroySemaphore(sem)
// }

// [WaitSemaphore] waits until a semaphore has a positive value and then decrements it.
//
// [WaitSemaphore]: https://wiki.libsdl.org/SDL3/SDL_WaitSemaphore
// func WaitSemaphore(sem *Semaphore)  {
//	sdlWaitSemaphore(sem)
// }

// [TryWaitSemaphore] sees if a semaphore has a positive value and decrement it if it does.
//
// [TryWaitSemaphore]: https://wiki.libsdl.org/SDL3/SDL_TryWaitSemaphore
// func TryWaitSemaphore(sem *Semaphore) bool {
//	return sdlTryWaitSemaphore(sem)
// }

// [WaitSemaphoreTimeout] waits until a semaphore has a positive value and then decrements it.
//
// [WaitSemaphoreTimeout]: https://wiki.libsdl.org/SDL3/SDL_WaitSemaphoreTimeout
// func WaitSemaphoreTimeout(sem *Semaphore, timeoutMS int32) bool {
//	return sdlWaitSemaphoreTimeout(sem, timeoutMS)
// }

// [SignalSemaphore] atomically increments a semaphore's value and wake waiting threads.
//
// [SignalSemaphore]: https://wiki.libsdl.org/SDL3/SDL_SignalSemaphore
// func SignalSemaphore(sem *Semaphore)  {
//	sdlSignalSemaphore(sem)
// }

// [GetSemaphoreValue] gets the current value of a semaphore.
//
// [GetSemaphoreValue]: https://wiki.libsdl.org/SDL3/SDL_GetSemaphoreValue
// func GetSemaphoreValue(sem *Semaphore) uint32 {
//	return sdlGetSemaphoreValue(sem)
// }

// [Condition] defines a means to block multiple threads until a condition is satisfied.
//
// [Condition]: https://wiki.libsdl.org/SDL3/SDL_Condition
type Condition struct{}

// [CreateCondition] creates a condition variable.
//
// [CreateCondition]: https://wiki.libsdl.org/SDL3/SDL_CreateCondition
// func CreateCondition() *Condition {
//	return sdlCreateCondition()
// }

// [DestroyCondition] destroys a condition variable.
//
// [DestroyCondition]: https://wiki.libsdl.org/SDL3/SDL_DestroyCondition
// func DestroyCondition(cond *Condition)  {
//	sdlDestroyCondition(cond)
// }

// [SignalCondition] restarts one of the threads that are waiting on the condition variable.
//
// [SignalCondition]: https://wiki.libsdl.org/SDL3/SDL_SignalCondition
// func SignalCondition(cond *Condition)  {
//	sdlSignalCondition(cond)
// }

// [BroadcastCondition] restarts all threads that are waiting on the condition variable.
//
// [BroadcastCondition]: https://wiki.libsdl.org/SDL3/SDL_BroadcastCondition
// func BroadcastCondition(cond *Condition)  {
//	sdlBroadcastCondition(cond)
// }

// [WaitCondition] waits until a condition variable is signaled.
//
// [WaitCondition]: https://wiki.libsdl.org/SDL3/SDL_WaitCondition
// func WaitCondition(cond *Condition, mutex *Mutex)  {
//	sdlWaitCondition(cond, mutex)
// }

// [WaitConditionTimeout] waits until a condition variable is signaled or a certain time has passed.
//
// [WaitConditionTimeout]: https://wiki.libsdl.org/SDL3/SDL_WaitConditionTimeout
// func WaitConditionTimeout(cond *Condition, mutex *Mutex, timeoutMS int32) bool {
//	return sdlWaitConditionTimeout(cond, mutex, timeoutMS)
// }

// [InitStatus] specifies the current status of an [InitState] structure.
//
// [InitStatus]: https://wiki.libsdl.org/SDL3/SDL_InitStatus
type InitStatus uint32

const (
	InitStatusUninitialized InitStatus = iota
	InitStatusInitializing
	InitStatusInitialized
	InitStatusUninitializing
)

// [InitState] is a structure used for thread-safe initialization and shutdown.
//
// [InitState]: https://wiki.libsdl.org/SDL3/SDL_InitState
type InitState struct {
	Status   AtomicInt
	Thread   ThreadID
	reserved uintptr
}

// [ShouldInit] returns whether initialization should be done.
//
// [ShouldInit]: https://wiki.libsdl.org/SDL3/SDL_ShouldInit
// func ShouldInit(state *InitState) bool {
//	return sdlShouldInit(state)
// }

// [ShouldQuit] returns whether cleanup should be done.
//
// [ShouldQuit]: https://wiki.libsdl.org/SDL3/SDL_ShouldQuit
// func ShouldQuit(state *InitState) bool {
//	return sdlShouldQuit(state)
// }

// [SetInitialized] finishes an initialization state transition.
//
// [SetInitialized]: https://wiki.libsdl.org/SDL3/SDL_SetInitialized
// func SetInitialized(state *InitState, initialized bool)  {
//	sdlSetInitialized(state, initialized)
// }
