package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

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

// func HasClipboardText() bool {
//	return sdlHasClipboardText()
// }

// func SetPrimarySelectionText(text string) bool {
//	return sdlSetPrimarySelectionText(text)
// }

// func GetPrimarySelectionText() string {
//	return sdlGetPrimarySelectionText()
// }

// func HasPrimarySelectionText() bool {
//	return sdlHasPrimarySelectionText()
// }

type ClipboardDataCallback uintptr

// func SetClipboardData(callback ClipboardDataCallback, cleanup ClipboardCleanupCallback, userdata unsafe.Pointer, mime_types **byte, num_mime_types uint64) bool {
//	return sdlSetClipboardData(callback, cleanup, userdata, mime_types, num_mime_types)
// }

// func ClearClipboardData() bool {
//	return sdlClearClipboardData()
// }

// func GetClipboardData(mime_type string, size *uint64) unsafe.Pointer {
//	return sdlGetClipboardData(mime_type, size)
// }

// func HasClipboardData(mime_type string) bool {
//	return sdlHasClipboardData(mime_type)
// }

// func GetClipboardMimeTypes(num_mime_types *uint64) **byte {
//	return sdlGetClipboardMimeTypes(num_mime_types)
// }
