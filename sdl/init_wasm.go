//go:build js && wasm

package sdl

import (
	"runtime"
	"syscall/js"
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

var bridge js.Value

func init() {
	runtime.LockOSThread()

	bridge = js.Global().Get("sdlBridge")
	if bridge.IsUndefined() {
		panic("sdlBridge not initialized")
	}

	sdlGetVersion = func() int32 { return int32(bridge.Call("SDL_GetVersion").Int()) }
	sdlGetRevision = func() string { return bridge.Call("SDL_GetRevision").String() }
	sdlInit = func(flags InitFlags) bool { return bridge.Call("SDL_Init", uint32(flags)).Int() != 0 }
	sdlQuit = func() { bridge.Call("SDL_Quit") }
	sdlQuitSubSystem = func(flags InitFlags) { bridge.Call("SDL_QuitSubSystem", flags) }
	sdlCreateWindow = func(title string, w, h int32, flags WindowFlags) *Window {
		return (*Window)(unsafe.Pointer(uintptr(bridge.Call("SDL_CreateWindow", title, w, h, uint32(flags)).Int())))
	}
	sdlDestroyWindow = func(window *Window) { bridge.Call("SDL_DestroyWindow", uintptr(unsafe.Pointer(window))) }
	sdlGetWindowTitle = func(window *Window) string {
		return bridge.Call("SDL_GetWindowTitle", uintptr(unsafe.Pointer(window))).String()
	}
	sdlGetError = func() string { return bridge.Call("SDL_GetError").String() }
	sdlCreateRenderer = func(window *Window, b *byte) *Renderer {
		return (*Renderer)(unsafe.Pointer(uintptr(bridge.Call("SDL_CreateRenderer", uintptr(unsafe.Pointer(window)), convert.ToString(b)).Int())))
	}
	sdlDestroyRenderer = func(renderer *Renderer) { bridge.Call("SDL_DestroyRenderer", uintptr(unsafe.Pointer(renderer))) }
	sdlSetRenderDrawColor = func(renderer *Renderer, r, g, b, a uint8) bool {
		return bridge.Call("SDL_SetRenderDrawColor", uintptr(unsafe.Pointer(renderer)), r, g, b, a).Int() != 0
	}
	sdlRenderClear = func(renderer *Renderer) bool {
		return bridge.Call("SDL_RenderClear", uintptr(unsafe.Pointer(renderer))).Int() != 0
	}
	sdlRenderPresent = func(renderer *Renderer) bool {
		return bridge.Call("SDL_RenderPresent", uintptr(unsafe.Pointer(renderer))).Int() != 0
	}
	sdlRenderLine = func(renderer *Renderer, x1, y1, x2, y2 float32) bool {
		return bridge.Call("SDL_RenderLine", uintptr(unsafe.Pointer(renderer)), x1, y1, x2, y2).Int() != 0
	}
	sdlSetRenderVSync = func(renderer *Renderer, vsync int32) bool {
		return bridge.Call("SDL_SetRenderVSync", uintptr(unsafe.Pointer(renderer)), vsync).Int() != 0
	}
	sdlRenderFillRect = func(renderer *Renderer, rect *FRect) bool {
		return bridge.Call("SDL_RenderFillRect", uintptr(unsafe.Pointer(renderer)), rect.X, rect.Y, rect.W, rect.H).Int() != 0
	}
	sdlPollEvent = func(event *Event) bool {
		ret := bridge.Call("SDL_PollEvent", uintptr(unsafe.Pointer(event)))
		js.CopyBytesToGo(unsafe.Slice((*byte)(unsafe.Pointer(event)), 128), ret.Index(1))
		return ret.Index(0).Int() != 0
	}
	sdlRenderDebugText = func(renderer *Renderer, x, y float32, text string) bool {
		return bridge.Call("SDL_RenderDebugText", uintptr(unsafe.Pointer(renderer)), x, y, text).Int() != 0
	}
}
