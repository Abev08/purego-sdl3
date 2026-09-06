package sdl

// [GetTicks] returns the number of milliseconds that have elapsed since the SDL library initialization.
//
// [GetTicks]: https://wiki.libsdl.org/SDL3/SDL_GetTicks
func GetTicks() uint64 {
	return sdlGetTicks()
}

// [GetTicksNS] returns the number of nanoseconds since SDL library initialization.
//
// [GetTicksNS]: https://wiki.libsdl.org/SDL3/SDL_GetTicksNS
func GetTicksNS() uint64 {
	return sdlGetTicksNS()
}

// [GetPerformanceCounter] returns the current value of the high resolution counter.
//
// [GetPerformanceCounter]: https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter
func GetPerformanceCounter() uint64 {
	return sdlGetPerformanceCounter()
}

// [GetPerformanceFrequency] returns the count per second of the high resolution counter.
//
// [GetPerformanceFrequency]: https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency
func GetPerformanceFrequency() uint64 {
	return sdlGetPerformanceFrequency()
}

// [Delay] waits a specified number of milliseconds before returning.
//
// [Delay]: https://wiki.libsdl.org/SDL3/SDL_Delay
// func Delay(ms uint32) {
//	sdlDelay(ms)
// }

// [DelayNS] waits a specified number of nanoseconds before returning.
//
// [DelayNS]: https://wiki.libsdl.org/SDL3/SDL_DelayNS
func DelayNS(ns uint64) {
	sdlDelayNS(ns)
}

// [DelayPrecise] waits a specified number of nanoseconds before returning.
//
// [DelayPrecise]: https://wiki.libsdl.org/SDL3/SDL_DelayPrecise
// func DelayPrecise(ns uint64) {
//	sdlDelayPrecise(ns)
// }

// [TimerID] specifies definition of the timer ID type.
//
// [TimerID]: https://wiki.libsdl.org/SDL3/SDL_TimerID
type TimerID uint32

// [TimerCallback] specifies function prototype for the millisecond timer callback function.
//
// [TimerCallback]: https://wiki.libsdl.org/SDL3/SDL_TimerCallback
type TimerCallback uintptr

// [AddTimer] calls a callback function at a future time.
//
// [AddTimer]: https://wiki.libsdl.org/SDL3/SDL_AddTimer
// func AddTimer(interval uint32, callback TimerCallback, userdata unsafe.Pointer) TimerID {
//	return sdlAddTimer(interval, callback, userdata)
// }

// [NSTimerCallback] specifies function prototype for the nanosecond timer callback function.
//
// [NSTimerCallback]: https://wiki.libsdl.org/SDL3/SDL_NSTimerCallback
type NSTimerCallback uintptr

// [AddTimerNS] calls a callback function at a future time.
//
// [AddTimerNS]: https://wiki.libsdl.org/SDL3/SDL_AddTimerNS
// func AddTimerNS(interval uint64, callback NSTimerCallback, userdata unsafe.Pointer) TimerID {
//	return sdlAddTimerNS(interval, callback, userdata)
// }

// [RemoveTimer] removes a timer created with [AddTimer].
//
// [RemoveTimer]: https://wiki.libsdl.org/SDL3/SDL_RemoveTimer
// func RemoveTimer(id TimerID) bool {
//	return sdlRemoveTimer(id)
// }
