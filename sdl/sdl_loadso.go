package sdl

// [SharedObject] is an opaque datatype that represents a loaded shared object.
//
// [SharedObject]: https://wiki.libsdl.org/SDL3/SDL_SharedObject
type SharedObject struct{}

// [LoadObject] dynamically loads a shared object.
//
// [LoadObject]: https://wiki.libsdl.org/SDL3/SDL_LoadObject
// func LoadObject(sofile *byte) *SharedObject {
// 	return sdlLoadObject(sofile)
// }

// [LoadFunction] looks up the address of the named function in a shared object.
//
// [LoadFunction]: https://wiki.libsdl.org/SDL3/SDL_LoadFunction
// func LoadFunction(handle *SharedObject, name *byte) FunctionPointer {
// 	return sdlLoadFunction(handle, name)
// }

// [UnloadObject] unloads a shared object from memory.
//
// [UnloadObject]: https://wiki.libsdl.org/SDL3/SDL_UnloadObject
// func UnloadObject(handle *SharedObject) {
// 	sdlUnloadObject(handle)
// }
