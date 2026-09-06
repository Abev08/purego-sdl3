package sdl

import "unsafe"

// [MainFunc] is the prototype for the application's main() function.
//
// [MainFunc]: https://wiki.libsdl.org/SDL3/SDL_main_func
type MainFunc func(argc int32, argv **byte) int32

// [Main] ans app-supplied function for program entry.
//
// [Main]: https://wiki.libsdl.org/SDL3/SDL_main
// func main(argc int32, argv **byte) int32 {
//	return sdlmain(argc, argv)
// }

// [SetMainReady] circumvents failure of [Init] when not using [main] as an entry.
//
// [SetMainReady]: https://wiki.libsdl.org/SDL3/SDL_SetMainReady
// func SetMainReady()  {
//	sdlSetMainReady()
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

// [EnterAppMainCallbacks] is an entry point for SDL's use in SDL_MAIN_USE_CALLBACKS.
//
// Calling this function blocks the execution on PC, but it doesn't block on WASM.
//
// [EnterAppMainCallbacks]: https://wiki.libsdl.org/SDL3/SDL_EnterAppMainCallbacks
func EnterAppMainCallbacks(argc int32, argv **byte, appInit AppInitFunc, appIter AppIterateFunc, appEvent AppEventFunc, appQuit AppQuitFunc) int32 {
	return sdlEnterAppMainCallbacks(argc, argv, appInit, appIter, appEvent, appQuit)
}

// [GDKSuspendComplete] callbacks from the application to let the suspend continue.
//
// [GDKSuspendComplete]: https://wiki.libsdl.org/SDL3/SDL_GDKSuspendComplete
// func GDKSuspendComplete()  {
//	sdlGDKSuspendComplete()
// }
