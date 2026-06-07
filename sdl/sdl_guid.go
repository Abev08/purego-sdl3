package sdl

// [GUID] is a 128-bit identifier for an input device that identifies that device across runs of SDL programs on the same platform.
//
// [GUID]: https://wiki.libsdl.org/SDL3/SDL_GUID
type GUID struct {
	Data [16]uint8
}

// func GUIDToString(guid GUID, pszGUID string, cbGUID int32)  {
//	sdlGUIDToString(guid, pszGUID, cbGUID)
// }

// func StringToGUID(pchGUID string) GUID {
//	return sdlStringToGUID(pchGUID)
// }
