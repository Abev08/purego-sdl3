package sdl

// [GUID] is a 128-bit identifier for an input device that identifies that device across runs of SDL programs on the same platform.
//
// [GUID]: https://wiki.libsdl.org/SDL3/SDL_GUID
type GUID struct {
	Data [16]uint8
}

// [GUIDToString] gets an ASCII string representation for a given [GUID].
//
// [GUIDToString]: https://wiki.libsdl.org/SDL3/SDL_GUIDToString
// func GUIDToString(guid GUID, pszGUID string, cbGUID int32)  {
//	sdlGUIDToString(guid, pszGUID, cbGUID)
// }

// [StringToGUID] converts a GUID string into a [GUID] structure.
//
// [StringToGUID]: https://wiki.libsdl.org/SDL3/SDL_StringToGUID
// func StringToGUID(pchGUID string) GUID {
//	return sdlStringToGUID(pchGUID)
// }
