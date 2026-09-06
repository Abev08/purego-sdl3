package sdl

import "unsafe"

const (
	PropAppMetadataNameString       = "SDL.app.metadata.name"
	PropAppMetadataVersionString    = "SDL.app.metadata.version"
	PropAppMetadataIdentifierString = "SDL.app.metadata.identifier"
	PropAppMetadataCreatorString    = "SDL.app.metadata.creator"
	PropAppMetadataCopyrightString  = "SDL.app.metadata.copyright"
	PropAppMetadataUrlString        = "SDL.app.metadata.url"
	PropAppMetadataTypeString       = "SDL.app.metadata.type"
)

// [InitFlags] defines the initialization flags for [Init] and/or [InitSubSystem].
//
// [InitFlags]: https://wiki.libsdl.org/SDL3/SDL_InitFlags
type InitFlags uint32

const (
	InitAudio    InitFlags = 0x00000010 // [INIT_AUDIO] implies [INIT_EVENTS].
	InitVideo    InitFlags = 0x00000020 // [INIT_VIDEO] implies [INIT_EVENTS], should be initialized on the main thread.
	InitJoystick InitFlags = 0x00000200 // [INIT_JOYSTICK] implies [INIT_EVENTS].
	InitHaptic   InitFlags = 0x00001000
	InitGamepad  InitFlags = 0x00002000 // [INIT_GAMEPAD] implies [INIT_JOYSTICK].
	InitEvents   InitFlags = 0x00004000
	InitSensor   InitFlags = 0x00008000 // [INIT_SENSOR] implies [INIT_EVENTS].
	InitCamera   InitFlags = 0x00010000 // [INIT_CAMERA] implies [INIT_EVENTS].
)

// [AppResult] defines the return values for optional main callbacks.
//
// [AppResult]: https://wiki.libsdl.org/SDL3/SDL_AppResult
type AppResult uint32

const (
	AppContinue AppResult = iota // Value that requests that the app continue from the main callbacks.
	AppSuccess                   // Value that requests termination with success from the main callbacks.
	AppFailure                   // Value that requests termination with error from the main callbacks.
)

// [AppInitFunc] defines function pointer typedef for [AppInit].
//
// [AppInitFunc]: https://wiki.libsdl.org/SDL3/SDL_AppInit_func
type AppInitFunc func(appState *unsafe.Pointer, argc int32, argv **byte) AppResult

// [AppIterateFunc] defines function pointer typedef for [AppIterate].
//
// [AppIterateFunc]: https://wiki.libsdl.org/SDL3/SDL_AppIterate_func
type AppIterateFunc func(appState unsafe.Pointer) AppResult

// [AppEventFunc] defines function pointer typedef for [AppEvent].
//
// [AppEventFunc]: https://wiki.libsdl.org/SDL3/SDL_AppEvent_func
type AppEventFunc func(appState unsafe.Pointer, event *Event) AppResult

// [AppQuitFunc] defines function pointer typedef for [AppQuit].
//
// [AppQuitFunc]: https://wiki.libsdl.org/SDL3/SDL_AppQuit_func
type AppQuitFunc func(appState unsafe.Pointer, result int32)

// [Init] initializes the SDL library.
//
// [Init]: https://wiki.libsdl.org/SDL3/SDL_Init
func Init(flags InitFlags) bool {
	return sdlInit(flags)
}

// [InitSubSystem] is a compatibility function to initialize the SDL library.
//
// This function and [Init] are interchangeable.
//
// [InitSubSystem]: https://wiki.libsdl.org/SDL3/SDL_InitSubSystem
func InitSubSystem(flags InitFlags) bool {
	return sdlInit(flags)
}

// [QuitSubSystem] shuts down specific SDL subsystems.
//
// You still need to call [Quit] even if you close all open subsystems with this function.
//
// [QuitSubSystem]: https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem
func QuitSubSystem(flags InitFlags) {
	sdlQuitSubSystem(flags)
}

// [WasInit] gets a mask of the specified subsystems which are currently initialized.
//
// [WasInit]: https://wiki.libsdl.org/SDL3/SDL_WasInit
// func WasInit(flags InitFlags) InitFlags {
//	return sdlWasInit(flags)
// }

// [Quit] cleans up all initialized subsystems.
//
// [Quit]: https://wiki.libsdl.org/SDL3/SDL_Quit
func Quit() {
	sdlQuit()
}

// [IsMainThread] returns whether this is the main thread.
//
// [IsMainThread]: https://wiki.libsdl.org/SDL3/SDL_IsMainThread
func IsMainThread() bool {
	return sdlIsMainThread()
}

// [MainThreadCallback] defines callback run on the main thread.
//
// [MainThreadCallback]: https://wiki.libsdl.org/SDL3/SDL_MainThreadCallback
type MainThreadCallback uintptr

// [RunOnMainThread] calls a function on the main thread during event processing.
//
// [RunOnMainThread]: https://wiki.libsdl.org/SDL3/SDL_RunOnMainThread
// func RunOnMainThread(callback MainThreadCallback, userdata unsafe.Pointer, wait_complete bool) bool {
//	return sdlRunOnMainThread(callback, userdata, wait_complete)
// }

// [SetAppMetadata] specifies basic metadata about your app.
//
// [SetAppMetadata]: https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata
// func SetAppMetadata(appname string, appversion string, appidentifier string) bool {
//	return sdlSetAppMetadata(appname, appversion, appidentifier)
// }

// [SetAppMetadataProperty] specifies metadata about your app through a set of properties.
//
// [SetAppMetadataProperty]: https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty
// func SetAppMetadataProperty(name string, value string) bool {
//	return sdlSetAppMetadataProperty(name, value)
// }

// [GetAppMetadataProperty] gets metadata about your app.
//
// [GetAppMetadataProperty]: https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty
func GetAppMetadataProperty(name string) string {
	return sdlGetAppMetadataProperty(name)
}
