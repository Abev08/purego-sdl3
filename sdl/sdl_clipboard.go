package sdl

import (
	"unsafe"

	"github.com/abev08/purego-sdl3/internal/convert"
)

// [SetClipboardText] puts UTF-8 text into the clipboard.
//
// [SetClipboardText]: https://wiki.libsdl.org/SDL3/SDL_SetClipboardText
// func SetClipboardText(text string) bool {
//	return sdlSetClipboardText(text)
// }

// [GetClipboardText] gets UTF-8 text from the clipboard.
//
// [GetClipboardText]: https://wiki.libsdl.org/SDL3/SDL_GetClipboardText
func GetClipboardText() string {
	ret := sdlGetClipboardText()
	defer Free(unsafe.Pointer(ret))
	return convert.ToString(ret)
}

// [HasClipboardText] queries whether the clipboard exists and contains a non-empty text string.
//
// [HasClipboardText]: https://wiki.libsdl.org/SDL3/SDL_HasClipboardText
// func HasClipboardText() bool {
//	return sdlHasClipboardText()
// }

// [SetPrimarySelectionText] puts UTF-8 text into the primary selection.
//
// [SetPrimarySelectionText]: https://wiki.libsdl.org/SDL3/SDL_SetPrimarySelectionText
// func SetPrimarySelectionText(text string) bool {
//	return sdlSetPrimarySelectionText(text)
// }

// [GetPrimarySelectionText] gets UTF-8 text from the primary selection.
//
// [GetPrimarySelectionText]: https://wiki.libsdl.org/SDL3/SDL_GetPrimarySelectionText
// func GetPrimarySelectionText() string {
//	return sdlGetPrimarySelectionText()
// }

// [HasPrimarySelectionText] queries whether the primary selection exists and contains a non-empty text.
//
// [HasPrimarySelectionText]: https://wiki.libsdl.org/SDL3/SDL_HasPrimarySelectionText
// func HasPrimarySelectionText() bool {
//	return sdlHasPrimarySelectionText()
// }

// [ClipboardDataCallback] defines callback function that will be called when data for the specified mime-type.
//
// [ClipboardDataCallback]: https://wiki.libsdl.org/SDL3/SDL_ClipboardDataCallback
type ClipboardDataCallback uintptr

// [ClipboardCleanupCallback] defines callback function that will be called when the clipboard is cleared, or when new data is set.
//
// [ClipboardCleanupCallback]: https://wiki.libsdl.org/SDL3/SDL_ClipboardCleanupCallback
type ClipboardCleanupCallback uintptr

// [SetClipboardData] offers clipboard data to the OS.
//
// [SetClipboardData]: https://wiki.libsdl.org/SDL3/SDL_SetClipboardData
// func SetClipboardData(callback ClipboardDataCallback, cleanup ClipboardCleanupCallback, userdata unsafe.Pointer, mime_types **byte, num_mime_types uint64) bool {
//	return sdlSetClipboardData(callback, cleanup, userdata, mime_types, num_mime_types)
// }

// [ClearClipboardData] clears the clipboard data.
//
// [ClearClipboardData]: https://wiki.libsdl.org/SDL3/SDL_ClearClipboardData
// func ClearClipboardData() bool {
//	return sdlClearClipboardData()
// }

// [GetClipboardData] gets the data from the clipboard for a given mime type.
//
// [GetClipboardData]: https://wiki.libsdl.org/SDL3/SDL_GetClipboardData
// func GetClipboardData(mime_type string, size *uint64) unsafe.Pointer {
//	return sdlGetClipboardData(mime_type, size)
// }

// [HasClipboardData] queries whether there is data in the clipboard for the provided mime type.
//
// [HasClipboardData]: https://wiki.libsdl.org/SDL3/SDL_HasClipboardData
// func HasClipboardData(mime_type string) bool {
//	return sdlHasClipboardData(mime_type)
// }

// [GetClipboardMimeTypes] retrieves the list of mime types available in the clipboard.
//
// [GetClipboardMimeTypes]: https://wiki.libsdl.org/SDL3/SDL_GetClipboardMimeTypes
// func GetClipboardMimeTypes(num_mime_types *uint64) **byte {
//	return sdlGetClipboardMimeTypes(num_mime_types)
// }
