package ttf

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
)

var (
	ttfAddFallbackFont                        func(*Font, *Font) bool
	ttfAppendTextString                       func(*Text, string, uint64) bool
	ttfClearFallbackFonts                     func(*Font)
	ttfCloseFont                              func(*Font)
	ttfCopyFont                               func(*Font) *Font
	ttfCreateGPUTextEngine                    func(*sdl.GPUDevice) *TextEngine
	ttfCreateGPUTextEngineWithProperties      func(sdl.PropertiesID) *TextEngine
	ttfCreateRendererTextEngine               func(*sdl.Renderer) *TextEngine
	ttfCreateRendererTextEngineWithProperties func(sdl.PropertiesID) *TextEngine
	ttfCreateSurfaceTextEngine                func() *TextEngine
	ttfCreateText                             func(*TextEngine, *Font, string, uint64) *Text
	ttfDeleteTextString                       func(*Text, int32, int32) bool
	ttfDestroyGPUTextEngine                   func(*TextEngine)
	ttfDestroyRendererTextEngine              func(*TextEngine)
	ttfDestroySurfaceTextEngine               func(*TextEngine)
	ttfDestroyText                            func(*Text)
	ttfDrawRendererText                       func(*Text, float32, float32) bool
	ttfDrawSurfaceText                        func(*Text, int32, int32, *sdl.Surface) bool
	ttfFontHasGlyph                           func(*Font, rune) bool
	ttfFontIsFixedWidth                       func(*Font) bool
	ttfFontIsScalable                         func(*Font) bool
	ttfGetFontAscent                          func(*Font) int32
	ttfGetFontDescent                         func(*Font) int32
	ttfGetFontDirection                       func(*Font) Direction
	ttfGetFontDPI                             func(*Font, *int32, *int32) bool
	ttfGetFontFamilyName                      func(*Font) string
	ttfGetFontGeneration                      func(*Font) uint32
	ttfGetFontHeight                          func(*Font) int32
	ttfGetFontHinting                         func(*Font) HintingFlags
	ttfGetFontKerning                         func(*Font) bool
	ttfGetFontLineSkip                        func(*Font) int32
	ttfGetFontOutline                         func(*Font) int32
	ttfGetFontProperties                      func(*Font) sdl.PropertiesID
	ttfGetFontScript                          func(*Font) uint32
	ttfGetFontSDF                             func(*Font) bool
	ttfGetFontSize                            func(*Font) float32
	ttfGetFontStyle                           func(*Font) FontStyleFlags
	ttfGetFontStyleName                       func(*Font) string
	ttfGetFontWeight                          func(*Font) int32
	ttfGetFontWrapAlignment                   func(*Font) HorizontalAlignment
	ttfGetFreeTypeVersion                     func(*int32, *int32, *int32)
	ttfGetGlyphImage                          func(*Font, rune, *ImageType) *sdl.Surface
	ttfGetGlyphImageForIndex                  func(*Font, uint32, *ImageType) *sdl.Surface
	ttfGetGlyphKerning                        func(*Font, rune, rune, *int32) bool
	ttfGetGlyphMetrics                        func(*Font, rune, *int32, *int32, *int32, *int32, *int32) bool
	ttfGetGlyphScript                         func(rune) uint32
	ttfGetGPUTextDrawData                     func(*Text) *GPUAtlasDrawSequence
	ttfGetGPUTextEngineWinding                func(*TextEngine) GPUTextEngineWinding
	ttfGetHarfBuzzVersion                     func(*int32, *int32, *int32)
	ttfGetNextTextSubString                   func(*Text, *SubString, *SubString) bool
	ttfGetNumFontFaces                        func(*Font) int32
	ttfGetPreviousTextSubString               func(*Text, *SubString, *SubString) bool
	ttfGetStringSize                          func(*Font, string, uint64, *int32, *int32) bool
	ttfGetStringSizeWrapped                   func(*Font, string, uint64, int32, *int32, *int32) bool
	ttfGetTextColor                           func(*Text, *uint8, *uint8, *uint8, *uint8) bool
	ttfGetTextColorFloat                      func(*Text, *float32, *float32, *float32, *float32) bool
	ttfGetTextDirection                       func(*Text) Direction
	ttfGetTextEngine                          func(*Text) *TextEngine
	ttfGetTextFont                            func(*Text) *Font
	ttfGetTextPosition                        func(*Text, *int32, *int32) bool
	ttfGetTextProperties                      func(*Text) sdl.PropertiesID
	ttfGetTextScript                          func(*Text) uint32
	ttfGetTextSize                            func(*Text, *int32, *int32) bool
	ttfGetTextSubString                       func(*Text, int32, *SubString) bool
	ttfGetTextSubStringForLine                func(*Text, int32, *SubString) bool
	ttfGetTextSubStringForPoint               func(*Text, int32, int32, *SubString) bool
	ttfGetTextSubStringsForRange              func(*Text, int32, int32, *int32)
	ttfGetTextWrapWidth                       func(*Text, *int32) bool
	ttfInit                                   func() bool
	ttfInsertTextString                       func(*Text, int32, string, uint64) bool
	ttfMeasureString                          func(*Font, string, uint64, int32, *int32, *uint64) bool
	ttfOpenFont                               func(string, float32) *Font
	ttfOpenFontIO                             func(*sdl.IOStream, bool, float32) *Font
	ttfOpenFontWithProperties                 func(sdl.PropertiesID) *Font
	ttfQuit                                   func()
	ttfRemoveFallbackFont                     func(*Font, *Font)
	ttfRenderGlyphBlended                     func(*Font, rune, uintptr) *sdl.Surface
	ttfRenderGlyphLCD                         func(*Font, rune, uintptr, uintptr) *sdl.Surface
	ttfRenderGlyphShaded                      func(*Font, rune, uintptr, uintptr) *sdl.Surface
	ttfRenderGlyphSolid                       func(*Font, rune, uintptr) *sdl.Surface
	ttfRenderTextBlended                      func(*Font, string, uint64, uintptr) *sdl.Surface
	ttfRenderTextBlendedWrapped               func(*Font, string, uint64, uintptr, int32) *sdl.Surface
	ttfRenderTextLCD                          func(*Font, string, uint64, uintptr, uintptr) *sdl.Surface
	ttfRenderTextLCDWrapped                   func(*Font, string, uint64, uintptr, uintptr, int32) *sdl.Surface
	ttfRenderTextShaded                       func(*Font, string, uint64, uintptr, uintptr) *sdl.Surface
	ttfRenderTextShadedWrapped                func(*Font, string, uint64, uintptr, uintptr, int32) *sdl.Surface
	ttfRenderTextSolid                        func(*Font, string, uint64, uintptr) *sdl.Surface
	ttfRenderTextSolidWrapped                 func(*Font, string, uint64, uintptr, int32) *sdl.Surface
	ttfSetFontDirection                       func(*Font, Direction) bool
	ttfSetFontHinting                         func(*Font, HintingFlags)
	ttfSetFontKerning                         func(*Font, bool)
	ttfSetFontLanguage                        func(*Font, string) bool
	ttfSetFontLineSkip                        func(*Font, int32)
	ttfSetFontOutline                         func(*Font, int32) bool
	ttfSetFontScript                          func(*Font, uint32) bool
	ttfSetFontSDF                             func(*Font, bool) bool
	ttfSetFontSize                            func(*Font, float32) bool
	ttfSetFontSizeDPI                         func(*Font, float32, int32, int32) bool
	ttfSetFontStyle                           func(*Font, FontStyleFlags)
	ttfSetFontWrapAlignment                   func(*Font, HorizontalAlignment)
	ttfSetGPUTextEngineWinding                func(*TextEngine, GPUTextEngineWinding)
	ttfSetTextColor                           func(*Text, uint8, uint8, uint8, uint8) bool
	ttfSetTextColorFloat                      func(*Text, float32, float32, float32, float32) bool
	ttfSetTextDirection                       func(*Text, Direction) bool
	ttfSetTextEngine                          func(*Text, *TextEngine) bool
	ttfSetTextFont                            func(*Text, *Font) bool
	ttfSetTextPosition                        func(*Text, int32, int32) bool
	ttfSetTextScript                          func(*Text, uint32) bool
	ttfSetTextString                          func(*Text, string, uint64) bool
	ttfSetTextWrapWhitespaceVisible           func(*Text, bool) bool
	ttfSetTextWrapWidth                       func(*Text, int32) bool
	ttfStringToTag                            func(string) uint32
	ttfTagToString                            func(uint32, *byte, uint64)
	ttfTextWrapWhitespaceVisible              func(*Text) bool
	ttfUpdateText                             func(*Text) bool
	ttfVersion                                func() int32
	ttfWasInit                                func() int32
)
