//go:build js && wasm

package ttf

import (
	"runtime"
	"syscall/js"
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

var bridge js.Value

func init() {
	runtime.LockOSThread()

	bridge = js.Global().Get("sdlBridge")
	if bridge.IsUndefined() {
		panic("sdlBridge not initialized")
	}
	memoryBufferPtr := bridge.Call("getMemoryBufferPtr").Int() // JS memory buffer
	memoryBufferView := bridge.Call("getMemoryBufferView")     // Uint8Array of JS memory buffer
	_, _ = memoryBufferPtr, memoryBufferView

	// purego.RegisterLibFunc(&ttfAddFallbackFont, lib, "TTF_AddFallbackFont")
	// purego.RegisterLibFunc(&ttfAppendTextString, lib, "TTF_AppendTextString")
	// purego.RegisterLibFunc(&ttfClearFallbackFonts, lib, "TTF_ClearFallbackFonts")
	ttfCloseFont = func(font *Font) { bridge.Call("TTF_CloseFont", unsafe.Pointer(font)) }
	// purego.RegisterLibFunc(&ttfCopyFont, lib, "TTF_CopyFont")
	// purego.RegisterLibFunc(&ttfCreateGPUTextEngine, lib, "TTF_CreateGPUTextEngine")
	// purego.RegisterLibFunc(&ttfCreateGPUTextEngineWithProperties, lib, "TTF_CreateGPUTextEngineWithProperties")
	ttfCreateRendererTextEngine = func(renderer *sdl.Renderer) *TextEngine {
		res := bridge.Call("TTF_CreateRendererTextEngine", unsafe.Pointer(renderer)).Int()
		if res == 0 {
			return nil
		}
		return (*TextEngine)(unsafe.Pointer(uintptr(res)))
	}
	// purego.RegisterLibFunc(&ttfCreateRendererTextEngineWithProperties, lib, "TTF_CreateRendererTextEngineWithProperties")
	// purego.RegisterLibFunc(&ttfCreateSurfaceTextEngine, lib, "TTF_CreateSurfaceTextEngine")
	ttfCreateText = func(engine *TextEngine, font *Font, text string, length uint64) *Text {
		res := bridge.Call("TTF_CreateText", unsafe.Pointer(engine), unsafe.Pointer(font), text, length).Int()
		if res == 0 {
			return nil
		}
		return (*Text)(unsafe.Pointer(uintptr(res)))
	}
	// purego.RegisterLibFunc(&ttfDeleteTextString, lib, "TTF_DeleteTextString")
	// purego.RegisterLibFunc(&ttfDestroyGPUTextEngine, lib, "TTF_DestroyGPUTextEngine")
	ttfDestroyRendererTextEngine = func(engine *TextEngine) { bridge.Call("TTF_DestroyRendererTextEngine", unsafe.Pointer(engine)) }
	// purego.RegisterLibFunc(&ttfDestroySurfaceTextEngine, lib, "TTF_DestroySurfaceTextEngine")
	ttfDestroyText = func(text *Text) { bridge.Call("TTF_DestroyText", unsafe.Pointer(text)) }
	ttfDrawRendererText = func(text *Text, x float32, y float32) bool {
		return bridge.Call("TTF_DrawRendererText", unsafe.Pointer(text), x, y).Int() != 0
	}
	// purego.RegisterLibFunc(&ttfDrawSurfaceText, lib, "TTF_DrawSurfaceText")
	// purego.RegisterLibFunc(&ttfFontHasGlyph, lib, "TTF_FontHasGlyph")
	// purego.RegisterLibFunc(&ttfFontIsFixedWidth, lib, "TTF_FontIsFixedWidth")
	// purego.RegisterLibFunc(&ttfFontIsScalable, lib, "TTF_FontIsScalable")
	ttfGetFontAscent = func(font *Font) int32 { return int32(bridge.Call("TTF_GetFontAscent", unsafe.Pointer(font)).Int()) }
	ttfGetFontDescent = func(font *Font) int32 { return int32(bridge.Call("TTF_GetFontDescent", unsafe.Pointer(font)).Int()) }
	ttfGetFontDirection = func(font *Font) Direction {
		return Direction(bridge.Call("TTF_GetFontDirection", unsafe.Pointer(font)).Int())
	}
	ttfGetFontDPI = func(font *Font, hdpi *int32, vdpi *int32) bool {
		res := bridge.Call("TTF_GetFontDPI", unsafe.Pointer(font))
		if hdpi != nil {
			*hdpi = int32(res.Index(1).Int())
		}
		if vdpi != nil {
			*vdpi = int32(res.Index(2).Int())
		}
		return res.Index(0).Int() != 0
	}
	ttfGetFontFamilyName = func(font *Font) string { return bridge.Call("TTF_GetFontFamilyName", unsafe.Pointer(font)).String() }
	ttfGetFontGeneration = func(font *Font) uint32 {
		return uint32(bridge.Call("TTF_GetFontGeneration", unsafe.Pointer(font)).Int())
	}
	ttfGetFontHeight = func(font *Font) int32 { return int32(bridge.Call("TTF_GetFontHeight", unsafe.Pointer(font)).Int()) }
	ttfGetFontHinting = func(font *Font) HintingFlags {
		return HintingFlags(bridge.Call("TTF_GetFontHinting", unsafe.Pointer(font)).Int())
	}
	ttfGetFontKerning = func(font *Font) bool { return bridge.Call("TTF_GetFontKerning", unsafe.Pointer(font)).Int() != 0 }
	ttfGetFontLineSkip = func(font *Font) int32 { return int32(bridge.Call("TTF_GetFontLineSkip", unsafe.Pointer(font)).Int()) }
	ttfGetFontOutline = func(font *Font) int32 { return int32(bridge.Call("TTF_GetFontOutline", unsafe.Pointer(font)).Int()) }
	// purego.RegisterLibFunc(&ttfGetFontProperties, lib, "TTF_GetFontProperties")
	// purego.RegisterLibFunc(&ttfGetFontScript, lib, "TTF_GetFontScript")
	// purego.RegisterLibFunc(&ttfGetFontSDF, lib, "TTF_GetFontSDF")
	// purego.RegisterLibFunc(&ttfGetFontSize, lib, "TTF_GetFontSize")
	// purego.RegisterLibFunc(&ttfGetFontStyle, lib, "TTF_GetFontStyle")
	// purego.RegisterLibFunc(&ttfGetFontStyleName, lib, "TTF_GetFontStyleName")
	// purego.RegisterLibFunc(&ttfGetFontWrapAlignment, lib, "TTF_GetFontWrapAlignment")
	// purego.RegisterLibFunc(&ttfGetFreeTypeVersion, lib, "TTF_GetFreeTypeVersion")
	// purego.RegisterLibFunc(&ttfGetGlyphImage, lib, "TTF_GetGlyphImage")
	// purego.RegisterLibFunc(&ttfGetGlyphImageForIndex, lib, "TTF_GetGlyphImageForIndex")
	// purego.RegisterLibFunc(&ttfGetGlyphKerning, lib, "TTF_GetGlyphKerning")
	// purego.RegisterLibFunc(&ttfGetGlyphMetrics, lib, "TTF_GetGlyphMetrics")
	// purego.RegisterLibFunc(&ttfGetGlyphScript, lib, "TTF_GetGlyphScript")
	// purego.RegisterLibFunc(&ttfGetGPUTextDrawData, lib, "TTF_GetGPUTextDrawData")
	// purego.RegisterLibFunc(&ttfGetGPUTextEngineWinding, lib, "TTF_GetGPUTextEngineWinding")
	// purego.RegisterLibFunc(&ttfGetHarfBuzzVersion, lib, "TTF_GetHarfBuzzVersion")
	// purego.RegisterLibFunc(&ttfGetNextTextSubString, lib, "TTF_GetNextTextSubString")
	// purego.RegisterLibFunc(&ttfGetNumFontFaces, lib, "TTF_GetNumFontFaces")
	// purego.RegisterLibFunc(&ttfGetPreviousTextSubString, lib, "TTF_GetPreviousTextSubString")
	// purego.RegisterLibFunc(&ttfGetStringSize, lib, "TTF_GetStringSize")
	// purego.RegisterLibFunc(&ttfGetStringSizeWrapped, lib, "TTF_GetStringSizeWrapped")
	// purego.RegisterLibFunc(&ttfGetTextColor, lib, "TTF_GetTextColor")
	// purego.RegisterLibFunc(&ttfGetTextColorFloat, lib, "TTF_GetTextColorFloat")
	// purego.RegisterLibFunc(&ttfGetTextDirection, lib, "TTF_GetTextDirection")
	// purego.RegisterLibFunc(&ttfGetTextEngine, lib, "TTF_GetTextEngine")
	// purego.RegisterLibFunc(&ttfGetTextFont, lib, "TTF_GetTextFont")
	// purego.RegisterLibFunc(&ttfGetTextPosition, lib, "TTF_GetTextPosition")
	// purego.RegisterLibFunc(&ttfGetTextProperties, lib, "TTF_GetTextProperties")
	// purego.RegisterLibFunc(&ttfGetTextScript, lib, "TTF_GetTextScript")
	// purego.RegisterLibFunc(&ttfGetTextSize, lib, "TTF_GetTextSize")
	// purego.RegisterLibFunc(&ttfGetTextSubString, lib, "TTF_GetTextSubString")
	// purego.RegisterLibFunc(&ttfGetTextSubStringForLine, lib, "TTF_GetTextSubStringForLine")
	// purego.RegisterLibFunc(&ttfGetTextSubStringForPoint, lib, "TTF_GetTextSubStringForPoint")
	// purego.RegisterLibFunc(&ttfGetTextSubStringsForRange, lib, "TTF_GetTextSubStringsForRange")
	// purego.RegisterLibFunc(&ttfGetTextWrapWidth, lib, "TTF_GetTextWrapWidth")
	ttfInit = func() bool { return bridge.Call("TTF_Init").Int() != 0 }
	// purego.RegisterLibFunc(&ttfInsertTextString, lib, "TTF_InsertTextString")
	// purego.RegisterLibFunc(&ttfMeasureString, lib, "TTF_MeasureString")
	// purego.RegisterLibFunc(&ttfOpenFont, lib, "TTF_OpenFont")
	ttfOpenFontIO = func(src *sdl.IOStream, closeio bool, ptsize float32) *Font {
		res := bridge.Call("TTF_OpenFontIO", unsafe.Pointer(src), closeio, ptsize).Int()
		if res == 0 {
			return nil
		}
		return (*Font)(unsafe.Pointer(uintptr(res)))
	}
	// purego.RegisterLibFunc(&ttfOpenFontWithProperties, lib, "TTF_OpenFontWithProperties")
	ttfQuit = func() { bridge.Call("TTF_Quit") }
	// purego.RegisterLibFunc(&ttfRemoveFallbackFont, lib, "TTF_RemoveFallbackFont")
	ttfRenderGlyphBlended = func(font *Font, ch rune, fg uintptr) *sdl.Surface {
		res := bridge.Call("TTF_RenderGlyph_Blended", unsafe.Pointer(font), uint32(ch), fg).Int()
		if res == 0 {
			return nil
		}
		sur := sdl.Surface{}
		surBytes := unsafe.Slice((*byte)(unsafe.Pointer(&sur)), unsafe.Sizeof(sdl.Surface{}))
		memoryView := bridge.Call("copyBytes", res, unsafe.Sizeof(sdl.Surface{}))
		js.CopyBytesToGo(surBytes, memoryView)
		sdl.StructToSDLPointer[unsafe.Pointer(&sur)] = res
		return &sur
	}
	// purego.RegisterLibFunc(&ttfRenderGlyphLCD, lib, "TTF_RenderGlyph_LCD")
	// purego.RegisterLibFunc(&ttfRenderGlyphShaded, lib, "TTF_RenderGlyph_Shaded")
	// purego.RegisterLibFunc(&ttfRenderGlyphSolid, lib, "TTF_RenderGlyph_Solid")
	ttfRenderTextBlended = func(font *Font, text string, length uint64, fg uintptr) *sdl.Surface {
		res := bridge.Call("TTF_RenderText_Blended", unsafe.Pointer(font), text, length, fg).Int()
		if res == 0 {
			return nil
		}
		sur := sdl.Surface{}
		surBytes := unsafe.Slice((*byte)(unsafe.Pointer(&sur)), unsafe.Sizeof(sdl.Surface{}))
		memoryView := bridge.Call("copyBytes", res, unsafe.Sizeof(sdl.Surface{}))
		js.CopyBytesToGo(surBytes, memoryView)
		sdl.StructToSDLPointer[unsafe.Pointer(&sur)] = res
		return &sur
	}
	ttfRenderTextBlendedWrapped = func(font *Font, text string, length uint64, fg uintptr, wrapWidth int32) *sdl.Surface {
		res := bridge.Call("TTF_RenderText_Blended_Wrapped", unsafe.Pointer(font), text, length, fg, wrapWidth).Int()
		if res == 0 {
			return nil
		}
		sur := sdl.Surface{}
		surBytes := unsafe.Slice((*byte)(unsafe.Pointer(&sur)), unsafe.Sizeof(sdl.Surface{}))
		memoryView := bridge.Call("copyBytes", res, unsafe.Sizeof(sdl.Surface{}))
		js.CopyBytesToGo(surBytes, memoryView)
		sdl.StructToSDLPointer[unsafe.Pointer(&sur)] = res
		return &sur
	}
	// purego.RegisterLibFunc(&ttfRenderTextLCD, lib, "TTF_RenderText_LCD")
	// purego.RegisterLibFunc(&ttfRenderTextLCDWrapped, lib, "TTF_RenderText_LCD_Wrapped")
	// purego.RegisterLibFunc(&ttfRenderTextShaded, lib, "TTF_RenderText_Shaded")
	// purego.RegisterLibFunc(&ttfRenderTextShadedWrapped, lib, "TTF_RenderText_Shaded_Wrapped")
	ttfRenderTextSolid = func(font *Font, text string, length uint64, fg uintptr) *sdl.Surface {
		res := bridge.Call("TTF_RenderText_Solid", unsafe.Pointer(font), text, length, fg).Int()
		if res == 0 {
			return nil
		}
		sur := sdl.Surface{}
		surBytes := unsafe.Slice((*byte)(unsafe.Pointer(&sur)), unsafe.Sizeof(sdl.Surface{}))
		memoryView := bridge.Call("copyBytes", res, unsafe.Sizeof(sdl.Surface{}))
		js.CopyBytesToGo(surBytes, memoryView)
		sdl.StructToSDLPointer[unsafe.Pointer(&sur)] = res
		return &sur
	}
	// purego.RegisterLibFunc(&ttfRenderTextSolidWrapped, lib, "TTF_RenderText_Solid_Wrapped")
	// purego.RegisterLibFunc(&ttfSetFontDirection, lib, "TTF_SetFontDirection")
	// purego.RegisterLibFunc(&ttfSetFontHinting, lib, "TTF_SetFontHinting")
	// purego.RegisterLibFunc(&ttfSetFontKerning, lib, "TTF_SetFontKerning")
	// purego.RegisterLibFunc(&ttfSetFontLanguage, lib, "TTF_SetFontLanguage")
	// purego.RegisterLibFunc(&ttfSetFontLineSkip, lib, "TTF_SetFontLineSkip")
	// purego.RegisterLibFunc(&ttfSetFontOutline, lib, "TTF_SetFontOutline")
	// purego.RegisterLibFunc(&ttfSetFontScript, lib, "TTF_SetFontScript")
	// purego.RegisterLibFunc(&ttfSetFontSDF, lib, "TTF_SetFontSDF")
	// purego.RegisterLibFunc(&ttfSetFontSize, lib, "TTF_SetFontSize")
	// purego.RegisterLibFunc(&ttfSetFontSizeDPI, lib, "TTF_SetFontSizeDPI")
	// purego.RegisterLibFunc(&ttfSetFontStyle, lib, "TTF_SetFontStyle")
	// purego.RegisterLibFunc(&ttfSetFontWrapAlignment, lib, "TTF_SetFontWrapAlignment")
	// purego.RegisterLibFunc(&ttfSetGPUTextEngineWinding, lib, "TTF_SetGPUTextEngineWinding")
	ttfSetTextColor = func(text *Text, r uint8, g uint8, b uint8, a uint8) bool {
		return bridge.Call("TTF_SetTextColor", unsafe.Pointer(text), r, g, b, a).Int() != 0
	}
	// purego.RegisterLibFunc(&ttfSetTextColorFloat, lib, "TTF_SetTextColorFloat")
	// purego.RegisterLibFunc(&ttfSetTextDirection, lib, "TTF_SetTextDirection")
	// purego.RegisterLibFunc(&ttfSetTextEngine, lib, "TTF_SetTextEngine")
	// purego.RegisterLibFunc(&ttfSetTextFont, lib, "TTF_SetTextFont")
	// purego.RegisterLibFunc(&ttfSetTextPosition, lib, "TTF_SetTextPosition")
	// purego.RegisterLibFunc(&ttfSetTextScript, lib, "TTF_SetTextScript")
	// purego.RegisterLibFunc(&ttfSetTextString, lib, "TTF_SetTextString")
	// purego.RegisterLibFunc(&ttfSetTextWrapWhitespaceVisible, lib, "TTF_SetTextWrapWhitespaceVisible")
	// purego.RegisterLibFunc(&ttfSetTextWrapWidth, lib, "TTF_SetTextWrapWidth")
	// purego.RegisterLibFunc(&ttfStringToTag, lib, "TTF_StringToTag")
	// purego.RegisterLibFunc(&ttfTagToString, lib, "TTF_TagToString")
	// purego.RegisterLibFunc(&ttfTextWrapWhitespaceVisible, lib, "TTF_TextWrapWhitespaceVisible")
	// purego.RegisterLibFunc(&ttfUpdateText, lib, "TTF_UpdateText")
	ttfVersion = func() int32 { return int32(bridge.Call("TTF_Version").Int()) }
	// purego.RegisterLibFunc(&ttfWasInit, lib, "TTF_WasInit")

	// Load functions available since 3.2.2
	versionMajor, versionMinor, versionPatch := Version()
	if versionMajor >= 3 && versionMinor >= 2 && versionPatch >= 2 {
		ttfGetFontWeight = func(font *Font) int32 { return int32(bridge.Call("TTF_GetFontWeight", unsafe.Pointer(font)).Int()) }
	}
}
