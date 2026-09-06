package sdl

// [X11EventHook] defines a callback to be used with [SetX11EventHook].
//
// [X11EventHook]: https://wiki.libsdl.org/SDL3/SDL_X11EventHook
type X11EventHook uintptr

// [SetX11EventHook] sets a callback for every X11 event.
//
// [SetX11EventHook]: https://wiki.libsdl.org/SDL3/SDL_SetX11EventHook
// func SetX11EventHook(callback X11EventHook, userdata unsafe.Pointer)  {
//	sdlSetX11EventHook(callback, userdata)
// }

// [SetLinuxThreadPriority] sets the UNIX nice value for a thread.
//
// [SetLinuxThreadPriority]: https://wiki.libsdl.org/SDL3/SDL_SetLinuxThreadPriority
// func SetLinuxThreadPriority(threadID int64, priority int32) bool {
//	return sdlSetLinuxThreadPriority(threadID, priority)
// }

// [SetLinuxThreadPriorityAndPolicy] sets the priority (not nice level) and scheduling policy for a thread.
//
// [SetLinuxThreadPriorityAndPolicy]: https://wiki.libsdl.org/SDL3/SDL_SetLinuxThreadPriorityAndPolicy
// func SetLinuxThreadPriorityAndPolicy(threadID int64, sdlPriority int32, schedPolicy int32) bool {
//	return sdlSetLinuxThreadPriorityAndPolicy(threadID, sdlPriority, schedPolicy)
// }

// [IsTablet] queries if the current device is a tablet.
//
// [IsTablet]: https://wiki.libsdl.org/SDL3/SDL_IsTablet
// func IsTablet() bool {
//	return sdlIsTablet()
// }

// [IsTV] queries if the current device is a TV.
//
// [IsTV]: https://wiki.libsdl.org/SDL3/SDL_IsTV
// func IsTV() bool {
//	return sdlIsTV()
// }

// [Sandbox] defines the application sandbox environment.
//
// [Sandbox]: https://wiki.libsdl.org/SDL3/SDL_Sandbox
type Sandbox uint32

const (
	SandboxNone Sandbox = iota
	SandboxUnknownContainer
	SandboxFlatpak
	SandboxSnap
	SandboxMacOS
)

// [GetSandbox] gets the application sandbox environment, if any.
//
// [GetSandbox]: https://wiki.libsdl.org/SDL3/SDL_GetSandbox
// func GetSandbox() Sandbox {
//	return sdlGetSandbox()
// }

// [OnApplicationWillTerminate] lets iOS apps with external event handling report onApplicationWillTerminate.
//
// [OnApplicationWillTerminate]: https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillTerminate
// func OnApplicationWillTerminate() {
//	sdlOnApplicationWillTerminate()
// }

// [OnApplicationDidReceiveMemoryWarning] lets iOS apps with external event handling report onApplicationDidReceiveMemoryWarning.
//
// [OnApplicationDidReceiveMemoryWarning]: https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidReceiveMemoryWarning
// func OnApplicationDidReceiveMemoryWarning() {
//	sdlOnApplicationDidReceiveMemoryWarning()
// }

// [OnApplicationWillEnterBackground] lets iOS apps with external event handling report onApplicationWillResignActive.
//
// [OnApplicationWillEnterBackground]: https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterBackground
// func OnApplicationWillEnterBackground() {
//	sdlOnApplicationWillEnterBackground()
// }

// [OnApplicationDidEnterBackground] lets iOS apps with external event handling report onApplicationDidEnterBackground.
//
// [OnApplicationDidEnterBackground]: https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterBackground
// func OnApplicationDidEnterBackground() {
//	sdlOnApplicationDidEnterBackground()
// }

// [OnApplicationWillEnterForeground] lets iOS apps with external event handling report onApplicationWillEnterForeground.
//
// [OnApplicationWillEnterForeground]: https://wiki.libsdl.org/SDL3/SDL_OnApplicationWillEnterForeground
// func OnApplicationWillEnterForeground() {
//	sdlOnApplicationWillEnterForeground()
// }

// [OnApplicationDidEnterForeground] lets iOS apps with external event handling report onApplicationDidBecomeActive.
//
// [OnApplicationDidEnterForeground]: https://wiki.libsdl.org/SDL3/SDL_OnApplicationDidEnterForeground
// func OnApplicationDidEnterForeground() {
//	sdlOnApplicationDidEnterForeground()
// }
