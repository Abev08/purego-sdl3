package sdl

const (
	PropThreadCreateEntryFunctionPointer = "SDL.thread.create.entry_function"
	PropThreadCreateNameString           = "SDL.thread.create.name"
	PropThreadCreateUserdataPointer      = "SDL.thread.create.userdata"
	PropThreadCreateStacksizeNumber      = "SDL.thread.create.stacksize"
)

// [Thread] is the SDL thread object.
//
// [Thread]: https://wiki.libsdl.org/SDL3/SDL_Thread
type Thread struct{}

// [ThreadID] specifies a unique numeric ID that identifies a thread.
//
// [ThreadID]: https://wiki.libsdl.org/SDL3/SDL_ThreadID
type ThreadID uint64

// [TLSID] specifies thread local storage ID.
//
// [TLSID]: https://wiki.libsdl.org/SDL3/SDL_TLSID
type TLSID AtomicInt

// [ThreadPriority] specifies the SDL thread priority.
//
// [ThreadPriority]: https://wiki.libsdl.org/SDL3/SDL_ThreadPriority
type ThreadPriority uint32

const (
	ThreadPriorityLow ThreadPriority = iota
	ThreadPriorityNormal
	ThreadPriorityHigh
	ThreadPriorityTimeCritical
)

// [ThreadState] specifies the SDL thread state.
//
// [ThreadState]: https://wiki.libsdl.org/SDL3/SDL_ThreadState
type ThreadState uint32

const (
	ThreadUnknown  ThreadState = iota // The thread is not valid.
	ThreadAlive                       // The thread is currently running.
	ThreadDetached                    // The thread is detached and can't be waited on.
	ThreadComplete                    // The thread has finished and should be cleaned up with [WaitThread].
)

// [ThreadFunction] specifies the function passed to [CreateThread] as the new thread's entry point.
//
// [ThreadFunction]: https://wiki.libsdl.org/SDL3/SDL_ThreadFunction
type ThreadFunction uintptr

// [CreateThreadRuntime] defines actual entry point for [CreateThread].
//
// [CreateThreadRuntime]: https://wiki.libsdl.org/SDL3/SDL_CreateThreadRuntime
// func CreateThreadRuntime(fn ThreadFunction, name string, data unsafe.Pointer, pfnBeginThread FunctionPointer, pfnEndThread FunctionPointer) *Thread {
//	return sdlCreateThreadRuntime(fn, name, data, pfnBeginThread, pfnEndThread)
// }

// [CreateThreadWithPropertiesRuntime] defines actual entry point for [CreateThreadWithProperties].
//
// [CreateThreadWithPropertiesRuntime]: https://wiki.libsdl.org/SDL3/SDL_CreateThreadWithPropertiesRuntime
// func CreateThreadWithPropertiesRuntime(props PropertiesID, pfnBeginThread FunctionPointer, pfnEndThread FunctionPointer) *Thread {
//	return sdlCreateThreadWithPropertiesRuntime(props, pfnBeginThread, pfnEndThread)
// }

// [GetThreadName] gets the thread name as it was specified in [CreateThread].
//
// [GetThreadName]: https://wiki.libsdl.org/SDL3/SDL_GetThreadName
// func GetThreadName(thread *Thread) string {
//	return sdlGetThreadName(thread)
// }

// [GetCurrentThreadID] gets the thread identifier for the current thread.
//
// [GetCurrentThreadID]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentThreadID
// func GetCurrentThreadID() ThreadID {
//	return sdlGetCurrentThreadID()
// }

// [GetThreadID] gets the thread identifier for the specified thread.
//
// [GetThreadID]: https://wiki.libsdl.org/SDL3/SDL_GetThreadID
// func GetThreadID(thread *Thread) ThreadID {
//	return sdlGetThreadID(thread)
// }

// [SetCurrentThreadPriority] sets the priority for the current thread.
//
// [SetCurrentThreadPriority]: https://wiki.libsdl.org/SDL3/SDL_SetCurrentThreadPriority
// func SetCurrentThreadPriority(priority ThreadPriority) bool {
//	return sdlSetCurrentThreadPriority(priority)
// }

// [WaitThread] waits for a thread to finish.
//
// [WaitThread]: https://wiki.libsdl.org/SDL3/SDL_WaitThread
// func WaitThread(thread *Thread, status *int32) {
//	sdlWaitThread(thread, status)
// }

// [GetThreadState] gets the current state of a thread.
//
// [GetThreadState]: https://wiki.libsdl.org/SDL3/SDL_GetThreadState
// func GetThreadState(thread *Thread) ThreadState {
//	return sdlGetThreadState(thread)
// }

// [DetachThread] lets a thread clean up on exit without intervention.
//
// [DetachThread]: https://wiki.libsdl.org/SDL3/SDL_DetachThread
// func DetachThread(thread *Thread) {
//	sdlDetachThread(thread)
// }

// [GetTLS] gets the current thread's value associated with a thread local storage ID.
//
// [GetTLS]: https://wiki.libsdl.org/SDL3/SDL_GetTLS
// func GetTLS(id *TLSID) unsafe.Pointer {
//	return sdlGetTLS(id)
// }

// [TLSDestructorCallback] defines the callback used to cleanup data passed to [SetTLS].
//
// [TLSDestructorCallback]: https://wiki.libsdl.org/SDL3/SDL_TLSDestructorCallback
type TLSDestructorCallback uintptr

// [SetTLS] sets the current thread's value associated with a thread local storage ID.
//
// [SetTLS]: https://wiki.libsdl.org/SDL3/SDL_SetTLS
// func SetTLS(id *TLSID, value unsafe.Pointer, destructor TLSDestructorCallback) bool {
//	return sdlSetTLS(id, value, destructor)
// }

// [CleanupTLS] cleanups all TLS data for this thread.
//
// [CleanupTLS]: https://wiki.libsdl.org/SDL3/SDL_CleanupTLS
// func CleanupTLS()  {
//	sdlCleanupTLS()
// }
