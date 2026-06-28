package sdl

import "unsafe"

// [MainFunc] is the prototype for the application's main() function.
//
// [MainFunc]: https://wiki.libsdl.org/SDL3/SDL_main_func
type MainFunc func(argc int32, argv **byte) int32

// [AppInitFunc] is a function pointer typedef for [AppInit].
//
// [AppInitFunc]: https://wiki.libsdl.org/SDL3/SDL_AppInit_func
type AppInitFunc func(appState *unsafe.Pointer, argc int32, argv **byte) AppResult

// [AppIterateFunc] is a function pointer typedef for [AppIterate].
//
// [AppIterateFunc]: https://wiki.libsdl.org/SDL3/SDL_AppIterate_func
type AppIterateFunc func(appState unsafe.Pointer) AppResult

// [AppEventFunc] is a function pointer typedef for [AppEvent].
//
// [AppEventFunc]: https://wiki.libsdl.org/SDL3/SDL_AppEvent_func
type AppEventFunc func(appState unsafe.Pointer, event *Event) AppResult

// [AppQuitFunc] is a function pointer typedef for [AppQuit].
//
// [AppQuitFunc]: https://wiki.libsdl.org/SDL3/SDL_AppQuit_func
type AppQuitFunc func(appState unsafe.Pointer, result int32)

// [EnterAppMainCallbacks] is an entry point for SDL's use in SDL_MAIN_USE_CALLBACKS.
//
// Calling this function blocks the execution on PC, but it doesn't block on WASM.
//
// [EnterAppMainCallbacks]: https://wiki.libsdl.org/SDL3/SDL_EnterAppMainCallbacks
func EnterAppMainCallbacks(argc int32, argv **byte, appInit AppInitFunc, appIter AppIterateFunc, appEvent AppEventFunc, appQuit AppQuitFunc) int32 {
	return sdlEnterAppMainCallbacks(argc, argv, appInit, appIter, appEvent, appQuit)
}

// func GDKSuspendComplete()  {
//	sdlGDKSuspendComplete()
// }

// func main(argc int32, argv **byte) int32 {
//	return sdlmain(argc, argv)
// }

// [RunApp] initializes and launches an SDL application, by doing platform-specific initialization before calling your
// mainFunction and cleanups after it returns, if that is needed for a specific platform, otherwise it just calls
// mainFunction.
//
// Calling this function blocks the execution.
//
// [RunApp]: https://wiki.libsdl.org/SDL3/SDL_RunApp
func RunApp(argc int32, argv **byte, mainFunction MainFunc, reserved unsafe.Pointer) int32 {
	return sdlRunApp(argc, argv, mainFunction, reserved)
}

// func SetMainReady()  {
//	sdlSetMainReady()
// }
