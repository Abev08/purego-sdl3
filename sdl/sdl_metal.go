package sdl

// [MetalView] defines a handle to a CAMetalLayer-backed NSView (macOS) or UIView (iOS/tvOS).
//
// [MetalView]: https://wiki.libsdl.org/SDL3/SDL_MetalView
type MetalView uintptr

// [MetalCreateView] creates a CAMetalLayer-backed NSView/UIView and attach it to the specified.
//
// [MetalCreateView]: https://wiki.libsdl.org/SDL3/SDL_Metal_CreateView
// func MetalCreateView(window *Window) MetalView {
//	return sdlMetal_CreateView(window)
// }

// [MetalDestroyView] destroys an existing [MetalView] object.
//
// [MetalDestroyView]: https://wiki.libsdl.org/SDL3/SDL_Metal_DestroyView
// func MetalDestroyView(view MetalView) {
//	sdlMetal_DestroyView(view)
// }

// [MetalGetLayer] gets a pointer to the backing CAMetalLayer for the given view.
//
// [MetalGetLayer]: https://wiki.libsdl.org/SDL3/SDL_Metal_GetLayer
// func MetalGetLayer(view MetalView) unsafe.Pointer {
//	return sdlMetal_GetLayer(view)
// }
