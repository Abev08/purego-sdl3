package img

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
)

var (
	imgVersion                  func() int32
	imgLoadTypedIO              func(*sdl.IOStream, bool, string) *sdl.Surface
	imgLoad                     func(string) *sdl.Surface
	imgLoadIO                   func(*sdl.IOStream, bool) *sdl.Surface
	imgLoadTexture              func(*sdl.Renderer, string) *sdl.Texture
	imgLoadTextureIO            func(*sdl.Renderer, *sdl.IOStream, bool) *sdl.Texture
	imgLoadTextureTypedIO       func(*sdl.Renderer, *sdl.IOStream, bool, string) *sdl.Texture
	imgIsAVIF                   func(*sdl.IOStream) bool
	imgIsICO                    func(*sdl.IOStream) bool
	imgIsCUR                    func(*sdl.IOStream) bool
	imgIsBMP                    func(*sdl.IOStream) bool
	imgIsGIF                    func(*sdl.IOStream) bool
	imgIsJPG                    func(*sdl.IOStream) bool
	imgIsJXL                    func(*sdl.IOStream) bool
	imgIsLBM                    func(*sdl.IOStream) bool
	imgIsPCX                    func(*sdl.IOStream) bool
	imgIsPNG                    func(*sdl.IOStream) bool
	imgIsPNM                    func(*sdl.IOStream) bool
	imgIsSVG                    func(*sdl.IOStream) bool
	imgIsQOI                    func(*sdl.IOStream) bool
	imgIsTIF                    func(*sdl.IOStream) bool
	imgIsXCF                    func(*sdl.IOStream) bool
	imgIsXPM                    func(*sdl.IOStream) bool
	imgIsXV                     func(*sdl.IOStream) bool
	imgIsWEBP                   func(*sdl.IOStream) bool
	imgLoadAVIFIO               func(*sdl.IOStream) *sdl.Surface
	imgLoadICOIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadCURIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadBMPIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadGIFIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadJPGIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadJXLIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadLBMIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadPCXIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadPNGIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadPNMIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadSVGIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadQOIIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadTGAIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadTIFIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadXCFIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadXPMIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadXVIO                 func(*sdl.IOStream) *sdl.Surface
	imgLoadWEBPIO               func(*sdl.IOStream) *sdl.Surface
	imgLoadSizedSVGIO           func(*sdl.IOStream, int32, int32) *sdl.Surface
	imgReadXPMFromArray         func(**byte) *sdl.Surface
	imgReadXPMFromArrayToRGB888 func(**byte) *sdl.Surface
	imgSaveAVIF                 func(*sdl.Surface, string, int32) bool
	imgSaveAVIFIO               func(*sdl.Surface, *sdl.IOStream, bool, int32) bool
	imgSavePNG                  func(*sdl.Surface, string) bool
	imgSavePNGIO                func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveJPG                  func(*sdl.Surface, string, int32) bool
	imgSaveJPGIO                func(*sdl.Surface, *sdl.IOStream, bool, int32) bool
	imgLoadAnimation            func(string) *Animation
	imgLoadAnimationIO          func(*sdl.IOStream, bool) *Animation
	imgLoadAnimationTypedIO     func(*sdl.IOStream, bool, string) *Animation
	imgFreeAnimation            func(*Animation)
	imgLoadGIFAnimationIO       func(*sdl.IOStream) *Animation
	imgLoadWEBPAnimationIO      func(*sdl.IOStream) *Animation

	imgLoadGPUTexture                       func(*sdl.GPUDevice, *sdl.GPUCopyPass, string, *int32, *int32) *sdl.GPUTexture
	imgLoadGPUTextureIO                     func(*sdl.GPUDevice, *sdl.GPUCopyPass, *sdl.IOStream, bool, *int32, *int32) *sdl.GPUTexture
	imgLoadGPUTextureTypedIO                func(*sdl.GPUDevice, *sdl.GPUCopyPass, *sdl.IOStream, bool, string, *int32, *int32) *sdl.GPUTexture
	imgGetClipboardImage                    func() *sdl.Surface
	imgIsANI                                func(*sdl.IOStream) bool
	imgSave                                 func(*sdl.Surface, string) bool
	imgSaveTypedIO                          func(*sdl.Surface, *sdl.IOStream, bool, string) bool
	imgSaveBMP                              func(*sdl.Surface, string) bool
	imgSaveBMPIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveCUR                              func(*sdl.Surface, string) bool
	imgSaveCURIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveGIF                              func(*sdl.Surface, string) bool
	imgSaveGIFIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveICO                              func(*sdl.Surface, string) bool
	imgSaveICOIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveTGA                              func(*sdl.Surface, string) bool
	imgSaveTGAIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveWEBP                             func(*sdl.Surface, string, float32) bool
	imgSaveWEBPIO                           func(*sdl.Surface, *sdl.IOStream, bool, float32) bool
	imgLoadANIAnimationIO                   func(*sdl.IOStream) *Animation
	imgLoadAPNGAnimationIO                  func(*sdl.IOStream) *Animation
	imgLoadAVIFAnimationIO                  func(*sdl.IOStream) *Animation
	imgSaveAnimation                        func(*Animation, string) bool
	imgSaveAnimationTypedIO                 func(*Animation, *sdl.IOStream, bool, string) bool
	imgSaveANIAnimationIO                   func(*Animation, *sdl.IOStream, bool) bool
	imgSaveAPNGAnimationIO                  func(*Animation, *sdl.IOStream, bool) bool
	imgSaveAVIFAnimationIO                  func(*Animation, *sdl.IOStream, bool, int32) bool
	imgSaveGIFAnimationIO                   func(*Animation, *sdl.IOStream, bool) bool
	imgSaveWEBPAnimationIO                  func(*Animation, *sdl.IOStream, bool, int32) bool
	imgCreateAnimatedCursor                 func(*Animation, int32, int32) *sdl.Cursor
	imgCreateAnimationEncoder               func(string) *AnimationEncoder
	imgCreateAnimationEncoderIO             func(*sdl.IOStream, bool, string) *AnimationEncoder
	imgCreateAnimationEncoderWithProperties func(sdl.PropertiesID) *AnimationEncoder
	imgAddAnimationEncoderFrame             func(*AnimationEncoder, *sdl.Surface, uint64) bool
	imgCloseAnimationEncoder                func(*AnimationEncoder) bool
	imgCreateAnimationDecoder               func(string) *AnimationDecoder
	imgCreateAnimationDecoderIO             func(*sdl.IOStream, bool, string) *AnimationDecoder
	imgCreateAnimationDecoderWithProperties func(sdl.PropertiesID) *AnimationDecoder
	imgGetAnimationDecoderProperties        func(*AnimationDecoder) sdl.PropertiesID
	imgGetAnimationDecoderFrame             func(*AnimationDecoder, **sdl.Surface, *uint64) bool
	imgGetAnimationDecoderStatus            func(*AnimationDecoder) AnimationDecoderStatus
	imgResetAnimationDecoder                func(*AnimationDecoder) bool
	imgCloseAnimationDecoder                func(*AnimationDecoder) bool
)
