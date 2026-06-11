//go:build js && wasm

package img

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

	// purego.RegisterLibFunc(&imgFreeAnimation, lib, "IMG_FreeAnimation")
	imgIsAVIF = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isAVIF", unsafe.Pointer(src)).Int() != 0 }
	imgIsBMP = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isBMP", unsafe.Pointer(src)).Int() != 0 }
	imgIsCUR = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isCUR", unsafe.Pointer(src)).Int() != 0 }
	imgIsGIF = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isGIF", unsafe.Pointer(src)).Int() != 0 }
	imgIsICO = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isICO", unsafe.Pointer(src)).Int() != 0 }
	imgIsJPG = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isJPG", unsafe.Pointer(src)).Int() != 0 }
	imgIsJXL = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isJXL", unsafe.Pointer(src)).Int() != 0 }
	imgIsLBM = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isLBM", unsafe.Pointer(src)).Int() != 0 }
	imgIsPCX = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isPCX", unsafe.Pointer(src)).Int() != 0 }
	imgIsPNG = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isPNG", unsafe.Pointer(src)).Int() != 0 }
	imgIsPNM = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isPNM", unsafe.Pointer(src)).Int() != 0 }
	imgIsQOI = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isQOI", unsafe.Pointer(src)).Int() != 0 }
	imgIsSVG = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isSVG", unsafe.Pointer(src)).Int() != 0 }
	imgIsTIF = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isTIF", unsafe.Pointer(src)).Int() != 0 }
	imgIsWEBP = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isWEBP", unsafe.Pointer(src)).Int() != 0 }
	imgIsXCF = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isXCF", unsafe.Pointer(src)).Int() != 0 }
	imgIsXPM = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isXPM", unsafe.Pointer(src)).Int() != 0 }
	imgIsXV = func(src *sdl.IOStream) bool { return bridge.Call("IMG_isXV", unsafe.Pointer(src)).Int() != 0 }
	// purego.RegisterLibFunc(&imgLoad, lib, "IMG_Load")
	imgLoadIO = func(src *sdl.IOStream, closeio bool) *sdl.Surface {
		res := bridge.Call("IMG_Load_IO", unsafe.Pointer(src), closeio).Int()
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
	// purego.RegisterLibFunc(&imgLoadAnimation, lib, "IMG_LoadAnimation")
	// purego.RegisterLibFunc(&imgLoadAnimationIO, lib, "IMG_LoadAnimation_IO")
	// purego.RegisterLibFunc(&imgLoadAnimationTypedIO, lib, "IMG_LoadAnimationTyped_IO")
	// purego.RegisterLibFunc(&imgLoadAVIFIO, lib, "IMG_LoadAVIF_IO")
	// purego.RegisterLibFunc(&imgLoadBMPIO, lib, "IMG_LoadBMP_IO")
	// purego.RegisterLibFunc(&imgLoadCURIO, lib, "IMG_LoadCUR_IO")
	// purego.RegisterLibFunc(&imgLoadGIFIO, lib, "IMG_LoadGIF_IO")
	// purego.RegisterLibFunc(&imgLoadGIFAnimationIO, lib, "IMG_LoadGIFAnimation_IO")
	// purego.RegisterLibFunc(&imgLoadICOIO, lib, "IMG_LoadICO_IO")
	// purego.RegisterLibFunc(&imgLoadJPGIO, lib, "IMG_LoadJPG_IO")
	// purego.RegisterLibFunc(&imgLoadJXLIO, lib, "IMG_LoadJXL_IO")
	// purego.RegisterLibFunc(&imgLoadLBMIO, lib, "IMG_LoadLBM_IO")
	// purego.RegisterLibFunc(&imgLoadPCXIO, lib, "IMG_LoadPCX_IO")
	// purego.RegisterLibFunc(&imgLoadPNGIO, lib, "IMG_LoadPNG_IO")
	// purego.RegisterLibFunc(&imgLoadPNMIO, lib, "IMG_LoadPNM_IO")
	// purego.RegisterLibFunc(&imgLoadQOIIO, lib, "IMG_LoadQOI_IO")
	// purego.RegisterLibFunc(&imgLoadSizedSVGIO, lib, "IMG_LoadSizedSVG_IO")
	// purego.RegisterLibFunc(&imgLoadSVGIO, lib, "IMG_LoadSVG_IO")
	// purego.RegisterLibFunc(&imgLoadTexture, lib, "IMG_LoadTexture")
	imgLoadTextureIO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) *sdl.Texture {
		res := bridge.Call("IMG_LoadTexture_IO", unsafe.Pointer(renderer), unsafe.Pointer(src), closeio).Int()
		if res == 0 {
			return nil
		}
		tex := sdl.Texture{}
		texBytes := unsafe.Slice((*byte)(unsafe.Pointer(&tex)), unsafe.Sizeof(sdl.Texture{}))
		memoryView := bridge.Call("copyBytes", res, unsafe.Sizeof(sdl.Texture{}))
		js.CopyBytesToGo(texBytes, memoryView)
		sdl.StructToSDLPointer[unsafe.Pointer(&tex)] = res
		return &tex
	}
	// purego.RegisterLibFunc(&imgLoadTextureTypedIO, lib, "IMG_LoadTextureTyped_IO")
	// purego.RegisterLibFunc(&imgLoadTGAIO, lib, "IMG_LoadTGA_IO")
	// purego.RegisterLibFunc(&imgLoadTIFIO, lib, "IMG_LoadTIF_IO")
	// purego.RegisterLibFunc(&imgLoadTypedIO, lib, "IMG_LoadTyped_IO")
	// purego.RegisterLibFunc(&imgLoadWEBPIO, lib, "IMG_LoadWEBP_IO")
	// purego.RegisterLibFunc(&imgLoadWEBPAnimationIO, lib, "IMG_LoadWEBPAnimation_IO")
	// purego.RegisterLibFunc(&imgLoadXCFIO, lib, "IMG_LoadXCF_IO")
	// purego.RegisterLibFunc(&imgLoadXPMIO, lib, "IMG_LoadXPM_IO")
	// purego.RegisterLibFunc(&imgLoadXVIO, lib, "IMG_LoadXV_IO")
	// purego.RegisterLibFunc(&imgReadXPMFromArray, lib, "IMG_ReadXPMFromArray")
	// purego.RegisterLibFunc(&imgReadXPMFromArrayToRGB888, lib, "IMG_ReadXPMFromArrayToRGB888")
	// purego.RegisterLibFunc(&imgSaveAVIF, lib, "IMG_SaveAVIF")
	// purego.RegisterLibFunc(&imgSaveAVIFIO, lib, "IMG_SaveAVIF_IO")
	// purego.RegisterLibFunc(&imgSaveJPG, lib, "IMG_SaveJPG")
	// purego.RegisterLibFunc(&imgSaveJPGIO, lib, "IMG_SaveJPG_IO")
	// purego.RegisterLibFunc(&imgSavePNG, lib, "IMG_SavePNG")
	// purego.RegisterLibFunc(&imgSavePNGIO, lib, "IMG_SavePNG_IO")
	imgVersion = func() int32 { return int32(bridge.Call("IMG_Version").Int()) }

	// // Load functions available since 3.4.0
	// versionMajor, versionMinor, _ := Version()
	// if versionMajor >= 3 && versionMinor >= 4 {
	// 	purego.RegisterLibFunc(&imgLoadGPUTexture, lib, "IMG_LoadGPUTexture")
	// 	purego.RegisterLibFunc(&imgLoadGPUTextureIO, lib, "IMG_LoadGPUTexture_IO")
	// 	purego.RegisterLibFunc(&imgLoadGPUTextureTypedIO, lib, "IMG_LoadGPUTextureTyped_IO")
	// 	purego.RegisterLibFunc(&imgGetClipboardImage, lib, "IMG_GetClipboardImage")
	// 	purego.RegisterLibFunc(&imgIsANI, lib, "IMG_isANI")
	// 	purego.RegisterLibFunc(&imgSave, lib, "IMG_Save")
	// 	purego.RegisterLibFunc(&imgSaveTypedIO, lib, "IMG_SaveTyped_IO")
	// 	purego.RegisterLibFunc(&imgSaveBMP, lib, "IMG_SaveBMP")
	// 	purego.RegisterLibFunc(&imgSaveBMPIO, lib, "IMG_SaveBMP_IO")
	// 	purego.RegisterLibFunc(&imgSaveCUR, lib, "IMG_SaveCUR")
	// 	purego.RegisterLibFunc(&imgSaveCURIO, lib, "IMG_SaveCUR_IO")
	// 	purego.RegisterLibFunc(&imgSaveGIF, lib, "IMG_SaveGIF")
	// 	purego.RegisterLibFunc(&imgSaveGIFIO, lib, "IMG_SaveGIF_IO")
	// 	purego.RegisterLibFunc(&imgSaveICO, lib, "IMG_SaveICO")
	// 	purego.RegisterLibFunc(&imgSaveICOIO, lib, "IMG_SaveICO_IO")
	// 	purego.RegisterLibFunc(&imgSaveTGA, lib, "IMG_SaveTGA")
	// 	purego.RegisterLibFunc(&imgSaveTGAIO, lib, "IMG_SaveTGA_IO")
	// 	purego.RegisterLibFunc(&imgSaveWEBP, lib, "IMG_SaveWEBP")
	// 	purego.RegisterLibFunc(&imgSaveWEBPIO, lib, "IMG_SaveWEBP_IO")
	// 	purego.RegisterLibFunc(&imgLoadANIAnimationIO, lib, "IMG_LoadANIAnimation_IO")
	// 	purego.RegisterLibFunc(&imgLoadAPNGAnimationIO, lib, "IMG_LoadAPNGAnimation_IO")
	// 	purego.RegisterLibFunc(&imgLoadAVIFAnimationIO, lib, "IMG_LoadAVIFAnimation_IO")
	// 	purego.RegisterLibFunc(&imgSaveAnimation, lib, "IMG_SaveAnimation")
	// 	purego.RegisterLibFunc(&imgSaveAnimationTypedIO, lib, "IMG_SaveAnimationTyped_IO")
	// 	purego.RegisterLibFunc(&imgSaveANIAnimationIO, lib, "IMG_SaveANIAnimation_IO")
	// 	purego.RegisterLibFunc(&imgSaveAPNGAnimationIO, lib, "IMG_SaveAPNGAnimation_IO")
	// 	purego.RegisterLibFunc(&imgSaveAVIFAnimationIO, lib, "IMG_SaveAVIFAnimation_IO")
	// 	purego.RegisterLibFunc(&imgSaveGIFAnimationIO, lib, "IMG_SaveGIFAnimation_IO")
	// 	purego.RegisterLibFunc(&imgSaveWEBPAnimationIO, lib, "IMG_SaveWEBPAnimation_IO")
	// 	purego.RegisterLibFunc(&imgCreateAnimatedCursor, lib, "IMG_CreateAnimatedCursor")
	// 	purego.RegisterLibFunc(&imgCreateAnimationEncoder, lib, "IMG_CreateAnimationEncoder")
	// 	purego.RegisterLibFunc(&imgCreateAnimationEncoderIO, lib, "IMG_CreateAnimationEncoder_IO")
	// 	purego.RegisterLibFunc(&imgCreateAnimationEncoderWithProperties, lib, "IMG_CreateAnimationEncoderWithProperties")
	// 	purego.RegisterLibFunc(&imgAddAnimationEncoderFrame, lib, "IMG_AddAnimationEncoderFrame")
	// 	purego.RegisterLibFunc(&imgCloseAnimationEncoder, lib, "IMG_CloseAnimationEncoder")
	// 	purego.RegisterLibFunc(&imgCreateAnimationDecoder, lib, "IMG_CreateAnimationDecoder")
	// 	purego.RegisterLibFunc(&imgCreateAnimationDecoderIO, lib, "IMG_CreateAnimationDecoder_IO")
	// 	purego.RegisterLibFunc(&imgCreateAnimationDecoderWithProperties, lib, "IMG_CreateAnimationDecoderWithProperties")
	// 	purego.RegisterLibFunc(&imgGetAnimationDecoderProperties, lib, "IMG_GetAnimationDecoderProperties")
	// 	purego.RegisterLibFunc(&imgGetAnimationDecoderFrame, lib, "IMG_GetAnimationDecoderFrame")
	// 	purego.RegisterLibFunc(&imgGetAnimationDecoderStatus, lib, "IMG_GetAnimationDecoderStatus")
	// 	purego.RegisterLibFunc(&imgResetAnimationDecoder, lib, "IMG_ResetAnimationDecoder")
	// 	purego.RegisterLibFunc(&imgCloseAnimationDecoder, lib, "IMG_CloseAnimationDecoder")
	// }
}
