package sdl

import (
	"unsafe"
)

var (
	// sdlabs                                   func(int32) int32
	// sdlacos                                  func(float64) float64
	// sdlacosf                                 func(float32) float32
	sdlAcquireCameraFrame         func(*Camera, *uint64) *Surface
	sdlAcquireGPUCommandBuffer    func(*GPUDevice) *GPUCommandBuffer
	sdlAcquireGPUSwapchainTexture func(*GPUCommandBuffer, *Window, **GPUTexture, *uint32, *uint32) bool
	// sdlAddAtomicInt                          func(*AtomicInt, int32) int32
	// sdlAddAtomicU32  												func(*AtomicU32, int32) uint32
	sdlAddEventWatch func(EventFilter, unsafe.Pointer) bool
	// sdlAddGamepadMapping                     func(string) int32
	// sdlAddGamepadMappingsFromFile            func(string) int32
	// sdlAddGamepadMappingsFromIO              func(*IOStream, bool) int32
	sdlAddHintCallback          func(string, HintCallback, unsafe.Pointer) bool
	sdlAddSurfaceAlternateImage func(*Surface, *Surface) bool
	// sdlAddTimer                              func(uint32, TimerCallback, unsafe.Pointer) TimerID
	// sdlAddTimerNS                            func(uint64, NSTimerCallback, unsafe.Pointer) TimerID
	sdlAddVulkanRenderSemaphores func(*Renderer, uint32, int64, int64) bool
	// sdlaligned_alloc                         func(uint64, uint64) unsafe.Pointer
	// sdlaligned_free                          func(unsafe.Pointer)
	// sdlasin                                  func(float64) float64
	// sdlasinf                                 func(float32) float32
	// sdlasprintf                              func(**byte, string) int32
	// sdlAsyncIOFromFile                       func(string, string) *AsyncIO
	// sdlatan                                  func(float64) float64
	// sdlatan2                                 func(float64, float64) float64
	// sdlatan2f                                func(float32, float32) float32
	// sdlatanf                                 func(float32) float32
	// sdlatof                                  func(string) float64
	// sdlatoi                                  func(string) int32
	// sdlAttachVirtualJoystick                 func(*VirtualJoystickDesc) JoystickID
	// sdlAudioDevicePaused                     func(AudioDeviceID) bool
	sdlAudioStreamDevicePaused func(stream *AudioStream) bool
	// sdlBeginGPUComputePass                   func(*GPUCommandBuffer, *GPUStorageTextureReadWriteBinding, uint32, *GPUStorageBufferReadWriteBinding, uint32) *GPUComputePass
	sdlBeginGPUCopyPass   func(*GPUCommandBuffer) *GPUCopyPass
	sdlBeginGPURenderPass func(*GPUCommandBuffer, *GPUColorTargetInfo, uint32, *GPUDepthStencilTargetInfo) *GPURenderPass
	// sdlBindAudioStream                       func(AudioDeviceID, *AudioStream) bool
	// sdlBindAudioStreams                      func(AudioDeviceID, **AudioStream, int32) bool
	// sdlBindGPUComputePipeline                func(*GPUComputePass, *GPUComputePipeline)
	// sdlBindGPUComputeSamplers                func(*GPUComputePass, uint32, *GPUTextureSamplerBinding, uint32)
	// sdlBindGPUComputeStorageBuffers          func(*GPUComputePass, uint32, **GPUBuffer, uint32)
	// sdlBindGPUComputeStorageTextures         func(*GPUComputePass, uint32, **GPUTexture, uint32)
	sdlBindGPUFragmentSamplers       func(*GPURenderPass, uint32, *GPUTextureSamplerBinding, uint32)
	sdlBindGPUFragmentStorageBuffers func(*GPURenderPass, uint32, **GPUBuffer, uint32)
	// sdlBindGPUFragmentStorageTextures        func(*GPURenderPass, uint32, **GPUTexture, uint32)
	sdlBindGPUGraphicsPipeline func(*GPURenderPass, *GPUGraphicsPipeline)
	sdlBindGPUIndexBuffer      func(*GPURenderPass, *GPUBufferBinding, GPUIndexElementSize)
	sdlBindGPUVertexBuffers    func(*GPURenderPass, uint32, *GPUBufferBinding, uint32)
	// sdlBindGPUVertexSamplers                 func(*GPURenderPass, uint32, *GPUTextureSamplerBinding, uint32)
	sdlBindGPUVertexStorageBuffers func(*GPURenderPass, uint32, **GPUBuffer, uint32)
	// sdlBindGPUVertexStorageTextures          func(*GPURenderPass, uint32, **GPUTexture, uint32)
	// sdlBlitGPUTexture                        func(*GPUCommandBuffer, *GPUBlitInfo)
	sdlBlitSurface func(*Surface, *Rect, *Surface, *Rect) bool
	// sdlBlitSurface9Grid                      func(*Surface, *Rect, int32, int32, int32, int32, float32, ScaleMode, *Surface, *Rect) bool
	// sdlBlitSurfaceScaled                     func(*Surface, *Rect, *Surface, *Rect, ScaleMode) bool
	// sdlBlitSurfaceTiled                      func(*Surface, *Rect, *Surface, *Rect) bool
	// sdlBlitSurfaceTiledWithScale             func(*Surface, *Rect, float32, ScaleMode, *Surface, *Rect) bool
	// sdlBlitSurfaceUnchecked                  func(*Surface, *Rect, *Surface, *Rect) bool
	// sdlBlitSurfaceUncheckedScaled            func(*Surface, *Rect, *Surface, *Rect, ScaleMode) bool
	// sdlBroadcastCondition                    func(*Condition)
	// sdlbsearch                               func(unsafe.Pointer, unsafe.Pointer, uint64, uint64, CompareCallback) unsafe.Pointer
	// sdlbsearch_r                             func(unsafe.Pointer, unsafe.Pointer, uint64, uint64, CompareCallback_r, unsafe.Pointer) unsafe.Pointer
	// sdlCalculateGPUTextureFormatSize         func(GPUTextureFormat, uint32, uint32, uint32) uint32
	// sdlcalloc                                func(uint64, uint64) unsafe.Pointer
	// sdlCancelGPUCommandBuffer                func(*GPUCommandBuffer) bool
	sdlCaptureMouse func(bool) bool
	// sdlceil                                  func(float64) float64
	// sdlceilf                                 func(float32) float32
	sdlClaimWindowForGPUDevice func(*GPUDevice, *Window) bool
	// sdlCleanupTLS                            func()
	sdlClearAudioStream func(stream *AudioStream) bool
	// sdlClearClipboardData                    func() bool
	sdlClearComposition func(*Window) bool
	sdlClearError       func() bool
	sdlClearProperty    func(PropertiesID, string) bool
	// sdlClearSurface                          func(*Surface, float32, float32, float32, float32) bool
	// sdlClickTrayEntry                        func(*TrayEntry)
	// sdlCloseAsyncIO                          func(*AsyncIO, bool, *AsyncIOQueue, unsafe.Pointer) bool
	// sdlCloseAudioDevice                      func(AudioDeviceID)
	sdlCloseCamera  func(*Camera)
	sdlCloseGamepad func(*Gamepad)
	// sdlCloseHaptic                           func(*Haptic)
	sdlCloseIO       func(*IOStream) bool
	sdlCloseJoystick func(*Joystick)
	// sdlCloseSensor                           func(*Sensor)
	// sdlCloseStorage                          func(*Storage) bool
	// sdlCompareAndSwapAtomicInt               func(*AtomicInt, int32, int32) bool
	// sdlCompareAndSwapAtomicPointer           func(*unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	// sdlCompareAndSwapAtomicU32               func(*AtomicU32, uint32, uint32) bool
	// sdlComposeCustomBlendMode                func(BlendFactor, BlendFactor, BlendOperation, BlendFactor, BlendFactor, BlendOperation) BlendMode
	// sdlConvertAudioSamples                   func(*AudioSpec, *uint8, int32, *AudioSpec, **uint8, *int32) bool
	sdlConvertEventToRenderCoordinates func(*Renderer, *Event) bool
	sdlConvertPixels                   func(int32, int32, PixelFormat, unsafe.Pointer, int32, PixelFormat, unsafe.Pointer, int32) bool
	sdlConvertPixelsAndColorspace      func(int32, int32, PixelFormat, Colorspace, PropertiesID, unsafe.Pointer, int32, PixelFormat, Colorspace, PropertiesID, unsafe.Pointer, int32) bool
	sdlConvertSurface                  func(*Surface, PixelFormat) *Surface
	sdlConvertSurfaceAndColorspace     func(*Surface, PixelFormat, *Palette, Colorspace, PropertiesID) *Surface
	// sdlCopyFile                              func(string, string) bool
	// sdlCopyGPUBufferToBuffer                 func(*GPUCopyPass, *GPUBufferLocation, *GPUBufferLocation, uint32, bool)
	// sdlCopyGPUTextureToTexture               func(*GPUCopyPass, *GPUTextureLocation, *GPUTextureLocation, uint32, uint32, uint32, bool)
	sdlCopyProperties func(PropertiesID, PropertiesID) bool
	// sdlcopysign                              func(float64, float64) float64
	// sdlcopysignf                             func(float32, float32) float32
	// sdlCopyStorageFile                       func(*Storage, string, string) bool
	// sdlcos                                   func(float64) float64
	// sdlcosf                                  func(float32) float32
	// sdlcrc16                                 func(uint16, unsafe.Pointer, uint64) uint16
	// sdlcrc32                                 func(uint32, unsafe.Pointer, uint64) uint32
	sdlCreateAnimatedCursor func(*CursorFrameInfo, int32, int32, int32) *Cursor
	// sdlCreateAsyncIOQueue                    func() *AsyncIOQueue
	// sdlCreateAudioStream                     func(*AudioSpec, *AudioSpec) *AudioStream
	sdlCreateColorCursor func(*Surface, int32, int32) *Cursor
	// sdlCreateCondition                       func() *Condition
	sdlCreateCursor func(*uint8, *uint8, int32, int32, int32, int32) *Cursor
	// sdlCreateDirectory                       func(string) bool
	// sdlCreateEnvironment                     func(bool) *Environment
	sdlCreateGPUBuffer func(*GPUDevice, *GPUBufferCreateInfo) *GPUBuffer
	// sdlCreateGPUComputePipeline              func(*GPUDevice, *GPUComputePipelineCreateInfo) *GPUComputePipeline
	sdlCreateGPUDevice func(GPUShaderFormat, bool, *byte) *GPUDevice
	// sdlCreateGPUDeviceWithProperties         func(PropertiesID) *GPUDevice
	sdlCreateGPUGraphicsPipeline func(*GPUDevice, *GPUGraphicsPipelineCreateInfo) *GPUGraphicsPipeline
	sdlCreateGPURenderer         func(*GPUDevice, *Window) *Renderer
	// sdlCreateGPURenderState      func(*Renderer, *GPURenderStateCreateInfo) *GPURenderState
	sdlCreateGPUSampler        func(*GPUDevice, *GPUSamplerCreateInfo) *GPUSampler
	sdlCreateGPUShader         func(*GPUDevice, *GPUShaderCreateInfo) *GPUShader
	sdlCreateGPUTexture        func(*GPUDevice, *GPUTextureCreateInfo) *GPUTexture
	sdlCreateGPUTransferBuffer func(*GPUDevice, *GPUTransferBufferCreateInfo) *GPUTransferBuffer
	// sdlCreateHapticEffect                    func(*Haptic, *HapticEffect) int32
	// sdlCreateMutex                           func() *Mutex
	sdlCreatePalette func(int32) *Palette
	// sdlCreatePopupWindow                     func(*Window, int32, int32, int32, int32, WindowFlags) *Window
	// sdlCreateProcess                         func(**byte, bool) *Process
	// sdlCreateProcessWithProperties           func(PropertiesID) *Process
	sdlCreateProperties             func() PropertiesID
	sdlCreateRenderer               func(*Window, *byte) *Renderer
	sdlCreateRendererWithProperties func(PropertiesID) *Renderer
	// sdlCreateRWLock                          func() *RWLock
	// sdlCreateSemaphore                       func(uint32) *Semaphore
	sdlCreateSoftwareRenderer func(*Surface) *Renderer
	// sdlCreateStorageDirectory                func(*Storage, string) bool
	sdlCreateSurface               func(int32, int32, PixelFormat) *Surface
	sdlCreateSurfaceFrom           func(int32, int32, PixelFormat, unsafe.Pointer, int32) *Surface
	sdlCreateSurfacePalette        func(*Surface) *Palette
	sdlCreateSystemCursor          func(SystemCursor) *Cursor
	sdlCreateTexture               func(*Renderer, PixelFormat, TextureAccess, int32, int32) *Texture
	sdlCreateTextureFromSurface    func(*Renderer, *Surface) *Texture
	sdlCreateTextureWithProperties func(*Renderer, PropertiesID) *Texture
	// sdlCreateThreadRuntime                   func(ThreadFunction, string, unsafe.Pointer, FunctionPointer, FunctionPointer) *Thread
	// sdlCreateThreadWithPropertiesRuntime     func(PropertiesID, FunctionPointer, FunctionPointer) *Thread
	// sdlCreateTray                            func(*Surface, string) *Tray
	// sdlCreateTrayMenu                        func(*Tray) *TrayMenu
	// sdlCreateTraySubmenu                     func(*TrayEntry) *TrayMenu
	sdlCreateWindow               func(string, int32, int32, WindowFlags) *Window
	sdlCreateWindowAndRenderer    func(string, int32, int32, WindowFlags, **Window, **Renderer) bool
	sdlCreateWindowWithProperties func(PropertiesID) *Window
	sdlCursorVisible              func() bool
	// sdlDateTimeToTime                        func(*DateTime, *Time) bool
	// sdlDelay                                 func(uint32)
	sdlDelayNS func(uint64)
	// sdlDelayPrecise                          func(uint64)
	// sdlDestroyAsyncIOQueue                   func(*AsyncIOQueue)
	sdlDestroyAudioStream func(*AudioStream)
	// sdlDestroyCondition                      func(*Condition)
	sdlDestroyCursor func(*Cursor)
	// sdlDestroyEnvironment                    func(*Environment)
	sdlDestroyGPUDevice func(*GPUDevice)
	// sdlDestroyGPURenderState func(*GPURenderState)
	// sdlDestroyHapticEffect                   func(*Haptic, int32)
	// sdlDestroyMutex                          func(*Mutex)
	sdlDestroyPalette func(*Palette)
	// sdlDestroyProcess                        func(*Process)
	sdlDestroyProperties func(PropertiesID)
	sdlDestroyRenderer   func(*Renderer)
	// sdlDestroyRWLock                         func(*RWLock)
	// sdlDestroySemaphore                      func(*Semaphore)
	sdlDestroySurface func(*Surface)
	sdlDestroyTexture func(*Texture)
	// sdlDestroyTray                           func(*Tray)
	sdlDestroyWindow        func(*Window)
	sdlDestroyWindowSurface func(*Window) bool
	// sdlDetachThread                          func(*Thread)
	// sdlDetachVirtualJoystick                 func(JoystickID) bool
	sdlDisableScreenSaver func() bool
	// sdlDispatchGPUCompute                    func(*GPUComputePass, uint32, uint32, uint32)
	// sdlDispatchGPUComputeIndirect            func(*GPUComputePass, *GPUBuffer, uint32)
	// sdlDownloadFromGPUBuffer                 func(*GPUCopyPass, *GPUBufferRegion, *GPUTransferBufferLocation)
	// sdlDownloadFromGPUTexture                func(*GPUCopyPass, *GPUTextureRegion, *GPUTextureTransferInfo)
	sdlDrawGPUIndexedPrimitives func(*GPURenderPass, uint32, uint32, uint32, int32, uint32)
	// sdlDrawGPUIndexedPrimitivesIndirect      func(*GPURenderPass, *GPUBuffer, uint32, uint32)
	sdlDrawGPUPrimitives func(*GPURenderPass, uint32, uint32, uint32, uint32)
	// sdlDrawGPUPrimitivesIndirect             func(*GPURenderPass, *GPUBuffer, uint32, uint32)
	sdlDuplicateSurface func(*Surface) *Surface
	// sdlEGL_GetCurrentConfig                  func() EGLConfig
	// sdlEGL_GetCurrentDisplay                 func() EGLDisplay
	// sdlEGL_GetProcAddress                    func(string) FunctionPointer
	// sdlEGL_GetWindowSurface                  func(*Window) EGLSurface
	// sdlEGL_SetAttributeCallbacks             func(EGLAttribArrayCallback, EGLIntArrayCallback, EGLIntArrayCallback, unsafe.Pointer)
	sdlEnableScreenSaver func() bool
	// sdlEndGPUComputePass                     func(*GPUComputePass)
	sdlEndGPUCopyPass        func(*GPUCopyPass)
	sdlEndGPURenderPass      func(*GPURenderPass)
	sdlEnterAppMainCallbacks func(int32, **byte, AppInitFunc, AppIterateFunc, AppEventFunc, AppQuitFunc) int32
	// sdlEnumerateDirectory                    func(string, EnumerateDirectoryCallback, unsafe.Pointer) bool
	sdlEnumerateProperties func(PropertiesID, EnumeratePropertiesCallback, unsafe.Pointer) bool
	// sdlEnumerateStorageDirectory             func(*Storage, string, EnumerateDirectoryCallback, unsafe.Pointer) bool
	sdlEventEnabled func(EventType) bool
	// sdlexp                                   func(float64) float64
	// sdlexpf                                  func(float32) float32
	// sdlfabs                                  func(float64) float64
	// sdlfabsf                                 func(float32) float32
	sdlFillSurfaceRect func(*Surface, *Rect, uint32) bool
	// sdlFillSurfaceRects                      func(*Surface, *Rect, int32, uint32) bool
	sdlFilterEvents func(EventFilter, unsafe.Pointer)
	sdlFlashWindow  func(*Window, FlashOperation) bool
	sdlFlipSurface  func(*Surface, FlipMode) bool
	// sdlfloor                                 func(float64) float64
	// sdlfloorf                                func(float32) float32
	sdlFlushAudioStream func(stream *AudioStream) bool
	sdlFlushEvent       func(EventType)
	sdlFlushEvents      func(EventType, EventType)
	// sdlFlushIO                               func(*IOStream) bool
	sdlFlushRenderer func(renderer *Renderer) bool
	// sdlfmod                                  func(float64, float64) float64
	// sdlfmodf                                 func(float32, float32) float32
	sdlfree func(mem unsafe.Pointer)
	// sdlGamepadConnected                      func(*Gamepad) bool
	// sdlGamepadEventsEnabled                  func() bool
	// sdlGamepadHasAxis                        func(*Gamepad, GamepadAxis) bool
	// sdlGamepadHasButton                      func(*Gamepad, GamepadButton) bool
	// sdlGamepadHasSensor                      func(*Gamepad, SensorType) bool
	// sdlGamepadSensorEnabled                  func(*Gamepad, SensorType) bool
	// sdlGDKSuspendComplete                    func()
	// sdlGenerateMipmapsForGPUTexture          func(*GPUCommandBuffer, *GPUTexture)
	sdlGetAppMetadataProperty func(string) string
	// sdlGetAssertionHandler                   func(*unsafe.Pointer) AssertionHandler
	// sdlGetAssertionReport                    func() *AssertData
	// sdlGetAsyncIOResult                      func(*AsyncIOQueue, *AsyncIOOutcome) bool
	// sdlGetAsyncIOSize                        func(*AsyncIO) int64
	// sdlGetAtomicInt                          func(*AtomicInt) int32
	// sdlGetAtomicPointer                      func(*unsafe.Pointer) unsafe.Pointer
	// sdlGetAtomicU32                          func(*AtomicU32) uint32
	// sdlGetAudioDeviceChannelMap              func(AudioDeviceID, *int32) *int32
	// sdlGetAudioDeviceFormat                  func(AudioDeviceID, *AudioSpec, *int32) bool
	// sdlGetAudioDeviceGain                    func(AudioDeviceID) float32
	// sdlGetAudioDeviceName                    func(AudioDeviceID) string
	sdlGetAudioDriver func(int32) string
	// sdlGetAudioFormatName                    func(AudioFormat) string
	// sdlGetAudioPlaybackDevices               func(*int32) *AudioDeviceID
	// sdlGetAudioRecordingDevices              func(*int32) *AudioDeviceID
	// sdlGetAudioStreamAvailable               func(*AudioStream) int32
	// sdlGetAudioStreamData                    func(*AudioStream, unsafe.Pointer, int32) int32
	sdlGetAudioStreamDevice         func(*AudioStream) AudioDeviceID
	sdlGetAudioStreamFormat         func(*AudioStream, *AudioSpec, *AudioSpec) bool
	sdlGetAudioStreamFrequencyRatio func(*AudioStream) float32
	sdlGetAudioStreamGain           func(*AudioStream) float32
	// sdlGetAudioStreamInputChannelMap         func(*AudioStream, *int32) *int32
	// sdlGetAudioStreamOutputChannelMap        func(*AudioStream, *int32) *int32
	// sdlGetAudioStreamProperties              func(*AudioStream) PropertiesID
	sdlGetAudioStreamQueued      func(stream *AudioStream) int32
	sdlGetBasePath               func() string
	sdlGetBooleanProperty        func(PropertiesID, string, bool) bool
	sdlGetCameraDriver           func(int32) string
	sdlGetCameraFormat           func(*Camera, *CameraSpec) bool
	sdlGetCameraID               func(*Camera) CameraID
	sdlGetCameraName             func(CameraID) string
	sdlGetCameraPermissionState  func(*Camera) int32
	sdlGetCameraPosition         func(CameraID) CameraPosition
	sdlGetCameraProperties       func(*Camera) PropertiesID
	sdlGetCameras                func(*int32) *CameraID
	sdlGetCameraSupportedFormats func(CameraID, *int32) **CameraSpec
	// sdlGetClipboardData                      func(string, *uint64) unsafe.Pointer
	// sdlGetClipboardMimeTypes                 func(*uint64) **byte
	sdlGetClipboardText                func() *byte
	sdlGetClosestFullscreenDisplayMode func(DisplayID, int32, int32, float32, bool, *DisplayMode) bool
	sdlGetCPUCacheLineSize             func() int32
	sdlGetCurrentAudioDriver           func() string
	sdlGetCurrentCameraDriver          func() string
	sdlGetCurrentDirectory             func() *byte
	sdlGetCurrentDisplayMode           func(DisplayID) *DisplayMode
	sdlGetCurrentDisplayOrientation    func(DisplayID) DisplayOrientation
	sdlGetCurrentRenderOutputSize      func(*Renderer, *int32, *int32) bool
	// sdlGetCurrentThreadID                    func() ThreadID
	// sdlGetCurrentTime                        func(*Time) bool
	sdlGetCurrentVideoDriver func() string
	sdlGetCursor             func() *Cursor
	// sdlGetDateTimeLocalePreferences          func(*DateFormat, *TimeFormat) bool
	// sdlGetDayOfWeek                          func(int32, int32, int32) int32
	// sdlGetDayOfYear                          func(int32, int32, int32) int32
	// sdlGetDaysInMonth                        func(int32, int32) int32
	// sdlGetDefaultAssertionHandler            func() AssertionHandler
	sdlGetDefaultCursor func() *Cursor
	// sdlGetDefaultLogOutputFunction           func() LogOutputFunction
	sdlGetDefaultTextureScaleMode func(*Renderer, *ScaleMode) bool
	sdlGetDesktopDisplayMode      func(DisplayID) *DisplayMode
	sdlGetDisplayBounds           func(DisplayID, *Rect) bool
	sdlGetDisplayContentScale     func(DisplayID) float32
	sdlGetDisplayForPoint         func(*Point) DisplayID
	sdlGetDisplayForRect          func(*Rect) DisplayID
	sdlGetDisplayForWindow        func(*Window) DisplayID
	sdlGetDisplayName             func(DisplayID) string
	sdlGetDisplayProperties       func(DisplayID) PropertiesID
	sdlGetDisplays                func(*int32) *DisplayID
	sdlGetDisplayUsableBounds     func(DisplayID, *Rect) bool
	// sdlgetenv                                func(string) string
	// sdlgetenv_unsafe                         func(string) string
	// sdlGetEnvironment                        func() *Environment
	// sdlGetEnvironmentVariable                func(*Environment, string) string
	// sdlGetEnvironmentVariables               func(*Environment) **byte
	sdlGetError func() string
	// sdlGetEventDescription       func(*Event, *byte, int32) int32
	sdlGetEventFilter            func(*EventFilter, *unsafe.Pointer) bool
	sdlGetFloatProperty          func(PropertiesID, string, float32) float32
	sdlGetFullscreenDisplayModes func(DisplayID, *int32) **DisplayMode
	// sdlGetGamepadAppleSFSymbolsNameForAxis   func(*Gamepad, GamepadAxis) string
	// sdlGetGamepadAppleSFSymbolsNameForButton func(*Gamepad, GamepadButton) string
	// sdlGetGamepadAxis                        func(*Gamepad, GamepadAxis) int16
	// sdlGetGamepadAxisFromString              func(string) GamepadAxis
	sdlGetGamepadBindings func(*Gamepad, *int32) **GamepadBinding
	// sdlGetGamepadButton                      func(*Gamepad, GamepadButton) bool
	// sdlGetGamepadButtonFromString            func(string) GamepadButton
	// sdlGetGamepadButtonLabel                 func(*Gamepad, GamepadButton) GamepadButtonLabel
	// sdlGetGamepadButtonLabelForType          func(GamepadType, GamepadButton) GamepadButtonLabel
	// sdlGetGamepadConnectionState             func(*Gamepad) JoystickConnectionState
	// sdlGetGamepadFirmwareVersion             func(*Gamepad) uint16
	sdlGetGamepadFromID func(JoystickID) *Gamepad
	// sdlGetGamepadFromPlayerIndex             func(int32) *Gamepad
	// sdlGetGamepadGUIDForID                   func(JoystickID) GUID
	// sdlGetGamepadID                          func(*Gamepad) JoystickID
	// sdlGetGamepadJoystick                    func(*Gamepad) *Joystick
	// sdlGetGamepadMapping                     func(*Gamepad) string
	// sdlGetGamepadMappingForGUID              func(GUID) string
	// sdlGetGamepadMappingForID                func(JoystickID) string
	// sdlGetGamepadMappings                    func(*int32) **byte
	sdlGetGamepadName      func(*Gamepad) string
	sdlGetGamepadNameForID func(JoystickID) string
	// sdlGetGamepadPath                        func(*Gamepad) string
	// sdlGetGamepadPathForID                   func(JoystickID) string
	// sdlGetGamepadPlayerIndex                 func(*Gamepad) int32
	// sdlGetGamepadPlayerIndexForID            func(JoystickID) int32
	// sdlGetGamepadPowerInfo                   func(*Gamepad, *int32) PowerState
	// sdlGetGamepadProduct                     func(*Gamepad) uint16
	// sdlGetGamepadProductForID                func(JoystickID) uint16
	// sdlGetGamepadProductVersion              func(*Gamepad) uint16
	// sdlGetGamepadProductVersionForID         func(JoystickID) uint16
	// sdlGetGamepadProperties                  func(*Gamepad) PropertiesID
	sdlGetGamepads func(*int32) *JoystickID
	// sdlGetGamepadSensorData                  func(*Gamepad, SensorType, *float32, int32) bool
	// sdlGetGamepadSensorDataRate              func(*Gamepad, SensorType) float32
	sdlGetGamepadSerial func(*Gamepad) string
	// sdlGetGamepadSteamHandle                 func(*Gamepad) uint64
	// sdlGetGamepadStringForAxis               func(GamepadAxis) string
	sdlGetGamepadStringForButton func(GamepadButton) string
	sdlGetGamepadStringForType   func(GamepadType) string
	// sdlGetGamepadTouchpadFinger              func(*Gamepad, int32, int32, *bool, *float32, *float32, *float32) bool
	sdlGetGamepadType func(*Gamepad) GamepadType
	// sdlGetGamepadTypeForID                   func(JoystickID) GamepadType
	// sdlGetGamepadTypeFromString              func(string) GamepadType
	// sdlGetGamepadVendor                      func(*Gamepad) uint16
	// sdlGetGamepadVendorForID                 func(JoystickID) uint16
	sdlGetGlobalMouseState                func(*float32, *float32) MouseButtonFlags
	sdlGetGlobalProperties                func() PropertiesID
	sdlGetGPUDeviceDriver                 func(*GPUDevice) string
	sdlGetGPUDeviceProperties             func(*GPUDevice) PropertiesID
	sdlGetGPUDriver                       func(int32) string
	sdlGetGPURendererDevice               func(*Renderer) *GPUDevice
	sdlGetGPUShaderFormats                func(*GPUDevice) GPUShaderFormat
	sdlGetGPUSwapchainTextureFormat       func(*GPUDevice, *Window) GPUTextureFormat
	sdlGetGPUTextureFormatFromPixelFormat func(PixelFormat) GPUTextureFormat
	sdlGetGrabbedWindow                   func() *Window
	// sdlGetHapticEffectStatus                 func(*Haptic, int32) bool
	// sdlGetHapticFeatures                     func(*Haptic) uint32
	// sdlGetHapticFromID                       func(HapticID) *Haptic
	// sdlGetHapticID                           func(*Haptic) HapticID
	// sdlGetHapticName                         func(*Haptic) string
	// sdlGetHapticNameForID                    func(HapticID) string
	// sdlGetHaptics                            func(*int32) *HapticID
	sdlGetHint        func(string) string
	sdlGetHintBoolean func(string, bool) bool
	// sdlGetIOProperties                       func(*IOStream) PropertiesID
	// sdlGetIOSize                             func(*IOStream) int64
	// sdlGetIOStatus                           func(*IOStream) IOStatus
	sdlGetJoystickAxis             func(*Joystick, int32) int16
	sdlGetJoystickAxisInitialState func(*Joystick, int32, *int16) bool
	sdlGetJoystickBall             func(*Joystick, int32, *int32, *int32) bool
	sdlGetJoystickButton           func(*Joystick, int32) bool
	sdlGetJoystickConnectionState  func(*Joystick) JoystickConnectionState
	sdlGetJoystickFirmwareVersion  func(*Joystick) uint16
	sdlGetJoystickFromID           func(JoystickID) *Joystick
	sdlGetJoystickFromPlayerIndex  func(int32) *Joystick
	// sdlGetJoystickGUID                func(*Joystick) GUID
	// sdlGetJoystickGUIDForID           func(JoystickID) GUID
	// sdlGetJoystickGUIDInfo            func(GUID, *uint16, *uint16, *uint16, *uint16)
	sdlGetJoystickHat                 func(*Joystick, int32) uint8
	sdlGetJoystickID                  func(*Joystick) JoystickID
	sdlGetJoystickName                func(*Joystick) string
	sdlGetJoystickNameForID           func(JoystickID) string
	sdlGetJoystickPath                func(*Joystick) string
	sdlGetJoystickPathForID           func(JoystickID) string
	sdlGetJoystickPlayerIndex         func(*Joystick) int32
	sdlGetJoystickPlayerIndexForID    func(JoystickID) int32
	sdlGetJoystickPowerInfo           func(*Joystick, *int32) PowerState
	sdlGetJoystickProduct             func(*Joystick) uint16
	sdlGetJoystickProductForID        func(JoystickID) uint16
	sdlGetJoystickProductVersion      func(*Joystick) uint16
	sdlGetJoystickProductVersionForID func(JoystickID) uint16
	sdlGetJoystickProperties          func(*Joystick) PropertiesID
	sdlGetJoysticks                   func(*int32) *JoystickID
	sdlGetJoystickSerial              func(*Joystick) string
	sdlGetJoystickType                func(*Joystick) JoystickType
	sdlGetJoystickTypeForID           func(JoystickID) JoystickType
	sdlGetJoystickVendor              func(*Joystick) uint16
	sdlGetJoystickVendorForID         func(JoystickID) uint16
	sdlGetKeyboardFocus               func() *Window
	sdlGetKeyboardNameForID           func(KeyboardID) string
	sdlGetKeyboards                   func(*int32) *KeyboardID
	sdlGetKeyboardState               func(*int32) *bool
	sdlGetKeyFromName                 func(string) Keycode
	sdlGetKeyFromScancode             func(Scancode, Keymod, bool) Keycode
	sdlGetKeyName                     func(Keycode) string
	// sdlGetLogOutputFunction                  func(*LogOutputFunction, *unsafe.Pointer)
	// sdlGetLogPriority                        func(int32) LogPriority
	// sdlGetMasksForPixelFormat                func(PixelFormat, *int32, *uint32, *uint32, *uint32, *uint32) bool
	// sdlGetMaxHapticEffects                   func(*Haptic) int32
	// sdlGetMaxHapticEffectsPlaying            func(*Haptic) int32
	// sdlGetMemoryFunctions                    func(*malloc_func, *calloc_func, *realloc_func, *free_func)
	sdlGetMice                      func(*int32) *MouseID
	sdlGetModState                  func() Keymod
	sdlGetMouseFocus                func() *Window
	sdlGetMouseNameForID            func(MouseID) string
	sdlGetMouseState                func(*float32, *float32) MouseButtonFlags
	sdlGetNaturalDisplayOrientation func(DisplayID) DisplayOrientation
	// sdlGetNumAllocations                     func() int32
	sdlGetNumAudioDrivers  func() int32
	sdlGetNumberProperty   func(PropertiesID, string, int64) int64
	sdlGetNumCameraDrivers func() int32
	// sdlGetNumGamepadTouchpadFingers          func(*Gamepad, int32) int32
	// sdlGetNumGamepadTouchpads                func(*Gamepad) int32
	sdlGetNumGPUDrivers func() int32
	// sdlGetNumHapticAxes                      func(*Haptic) int32
	sdlGetNumJoystickAxes    func(*Joystick) int32
	sdlGetNumJoystickBalls   func(*Joystick) int32
	sdlGetNumJoystickButtons func(*Joystick) int32
	sdlGetNumJoystickHats    func(*Joystick) int32
	sdlGetNumLogicalCPUCores func() int32
	sdlGetNumRenderDrivers   func() int32
	sdlGetNumVideoDrivers    func() int32
	// sdlGetOriginalMemoryFunctions            func(*malloc_func, *calloc_func, *realloc_func, *free_func)
	// sdlGetPathInfo                           func(string, *PathInfo) bool
	// sdlGetPenDeviceType        func(PenID) PenDeviceType
	sdlGetPerformanceCounter   func() uint64
	sdlGetPerformanceFrequency func() uint64
	sdlGetPixelFormatDetails   func(PixelFormat) *PixelFormatDetails
	// sdlGetPixelFormatForMasks                func(int32, uint32, uint32, uint32, uint32) PixelFormat
	sdlGetPixelFormatFromGPUTextureFormat func(GPUTextureFormat) PixelFormat
	// sdlGetPixelFormatName                    func(PixelFormat) string
	// sdlGetPlatform                           func() string
	sdlGetPointerProperty  func(PropertiesID, string, unsafe.Pointer) unsafe.Pointer
	sdlGetPowerInfo        func(*int32, *int32) PowerState
	sdlGetPreferredLocales func(*int32) **struct{ language, country *byte }
	sdlGetPrefPath         func(*byte, *byte) *byte
	sdlGetPrimaryDisplay   func() DisplayID
	// sdlGetPrimarySelectionText               func() string
	// sdlGetProcessInput                       func(*Process) *IOStream
	// sdlGetProcessOutput                      func(*Process) *IOStream
	// sdlGetProcessProperties                  func(*Process) PropertiesID
	sdlGetPropertyType func(PropertiesID, string) PropertyType
	// sdlGetRealGamepadType                    func(*Gamepad) GamepadType
	// sdlGetRealGamepadTypeForID               func(JoystickID) GamepadType
	sdlGetRectAndLineIntersection       func(*Rect, *int32, *int32, *int32, *int32) bool
	sdlGetRectAndLineIntersectionFloat  func(*FRect, *float32, *float32, *float32, *float32) bool
	sdlGetRectEnclosingPoints           func(*Point, int32, *Rect, *Rect) bool
	sdlGetRectEnclosingPointsFloat      func(*FPoint, int32, *FRect, *FRect) bool
	sdlGetRectIntersection              func(*Rect, *Rect, *Rect) bool
	sdlGetRectIntersectionFloat         func(*FRect, *FRect, *FRect) bool
	sdlGetRectUnion                     func(*Rect, *Rect, *Rect) bool
	sdlGetRectUnionFloat                func(*FRect, *FRect, *FRect) bool
	sdlGetRelativeMouseState            func(*float32, *float32) MouseButtonFlags
	sdlGetRenderClipRect                func(*Renderer, *Rect) bool
	sdlGetRenderColorScale              func(*Renderer, *float32) bool
	sdlGetRenderDrawBlendMode           func(*Renderer, *BlendMode) bool
	sdlGetRenderDrawColor               func(*Renderer, *uint8, *uint8, *uint8, *uint8) bool
	sdlGetRenderDrawColorFloat          func(*Renderer, *float32, *float32, *float32, *float32) bool
	sdlGetRenderDriver                  func(int32) string
	sdlGetRenderer                      func(*Window) *Renderer
	sdlGetRendererFromTexture           func(*Texture) *Renderer
	sdlGetRendererName                  func(*Renderer) string
	sdlGetRendererProperties            func(*Renderer) PropertiesID
	sdlGetRenderLogicalPresentation     func(*Renderer, *int32, *int32, *RendererLogicalPresentation) bool
	sdlGetRenderLogicalPresentationRect func(*Renderer, *FRect) bool
	sdlGetRenderMetalCommandEncoder     func(*Renderer) unsafe.Pointer
	sdlGetRenderMetalLayer              func(*Renderer) unsafe.Pointer
	sdlGetRenderOutputSize              func(*Renderer, *int32, *int32) bool
	sdlGetRenderSafeArea                func(*Renderer, *Rect) bool
	sdlGetRenderScale                   func(*Renderer, *float32, *float32) bool
	sdlGetRenderTarget                  func(*Renderer) *Texture
	sdlGetRenderTextureAddressMode      func(*Renderer, *TextureAddressMode, *TextureAddressMode) bool
	sdlGetRenderViewport                func(*Renderer, *Rect) bool
	sdlGetRenderVSync                   func(*Renderer, *int32) bool
	sdlGetRenderWindow                  func(*Renderer) *Window
	sdlGetRevision                      func() string
	// sdlGetRGB                                func(uint32, *PixelFormatDetails, *Palette, *uint8, *uint8, *uint8)
	// sdlGetRGBA                               func(uint32, *PixelFormatDetails, *Palette, *uint8, *uint8, *uint8, *uint8)
	// sdlGetSandbox                            func() Sandbox
	sdlGetScancodeFromKey  func(Keycode, *Keymod) Scancode
	sdlGetScancodeFromName func(string) Scancode
	sdlGetScancodeName     func(Scancode) string
	// sdlGetSemaphoreValue                     func(*Semaphore) uint32
	// sdlGetSensorData                         func(*Sensor, *float32, int32) bool
	// sdlGetSensorFromID                       func(SensorID) *Sensor
	// sdlGetSensorID                           func(*Sensor) SensorID
	// sdlGetSensorName                         func(*Sensor) string
	// sdlGetSensorNameForID                    func(SensorID) string
	// sdlGetSensorNonPortableType              func(*Sensor) int32
	// sdlGetSensorNonPortableTypeForID         func(SensorID) int32
	// sdlGetSensorProperties                   func(*Sensor) PropertiesID
	// sdlGetSensors                            func(*int32) *SensorID
	// sdlGetSensorType                         func(*Sensor) SensorType
	// sdlGetSensorTypeForID                    func(SensorID) SensorType
	// sdlGetSilenceValueForFormat              func(AudioFormat) int32
	sdlGetSIMDAlignment func() uint64
	// sdlGetStorageFileSize                    func(*Storage, string, *uint64) bool
	// sdlGetStoragePathInfo                    func(*Storage, string, *PathInfo) bool
	// sdlGetStorageSpaceRemaining              func(*Storage) uint64
	sdlGetStringProperty       func(PropertiesID, string, string) string
	sdlGetSurfaceAlphaMod      func(*Surface, *uint8) bool
	sdlGetSurfaceBlendMode     func(*Surface, *BlendMode) bool
	sdlGetSurfaceClipRect      func(*Surface, *Rect) bool
	sdlGetSurfaceColorKey      func(*Surface, *uint32) bool
	sdlGetSurfaceColorMod      func(*Surface, *uint8, *uint8, *uint8) bool
	sdlGetSurfaceColorspace    func(*Surface) Colorspace
	sdlGetSurfaceImages        func(*Surface, *int32) **Surface
	sdlGetSurfacePalette       func(*Surface) *Palette
	sdlGetSurfaceProperties    func(*Surface) PropertiesID
	sdlGetSystemPageSize       func() int32
	sdlGetSystemRAM            func() int32
	sdlGetSystemTheme          func() SystemTheme
	sdlGetTextInputArea        func(*Window, *Rect, *int32) bool
	sdlGetTextureAlphaMod      func(*Texture, *uint8) bool
	sdlGetTextureAlphaModFloat func(*Texture, *float32) bool
	sdlGetTextureBlendMode     func(*Texture, *BlendMode) bool
	sdlGetTextureColorMod      func(*Texture, *uint8, *uint8, *uint8) bool
	sdlGetTextureColorModFloat func(*Texture, *float32, *float32, *float32) bool
	sdlGetTexturePalette       func(*Texture) *Palette
	sdlGetTextureProperties    func(*Texture) PropertiesID
	sdlGetTextureScaleMode     func(*Texture, *ScaleMode) bool
	sdlGetTextureSize          func(*Texture, *float32, *float32) bool
	// sdlGetThreadID                           func(*Thread) ThreadID
	// sdlGetThreadName                         func(*Thread) string
	// sdlGetThreadState                        func(*Thread) ThreadState
	sdlGetTicks   func() uint64
	sdlGetTicksNS func() uint64
	// sdlGetTLS                                func(*TLSID) unsafe.Pointer
	// sdlGetTouchDeviceName                    func(TouchID) string
	// sdlGetTouchDevices                       func(*int32) *TouchID
	// sdlGetTouchDeviceType                    func(TouchID) TouchDeviceType
	// sdlGetTouchFingers                       func(TouchID, *int32) **Finger
	// sdlGetTrayEntries                        func(*TrayMenu, *int32) **TrayEntry
	// sdlGetTrayEntryChecked                   func(*TrayEntry) bool
	// sdlGetTrayEntryEnabled                   func(*TrayEntry) bool
	// sdlGetTrayEntryLabel                     func(*TrayEntry) string
	// sdlGetTrayEntryParent                    func(*TrayEntry) *TrayMenu
	// sdlGetTrayMenu                           func(*Tray) *TrayMenu
	// sdlGetTrayMenuParentEntry                func(*TrayMenu) *TrayEntry
	// sdlGetTrayMenuParentTray                 func(*TrayMenu) *Tray
	// sdlGetTraySubmenu                        func(*TrayEntry) *TrayMenu
	// sdlGetUserFolder                         func(Folder) string
	sdlGetVersion              func() int32
	sdlGetVideoDriver          func(int32) string
	sdlGetWindowAspectRatio    func(*Window, *float32, *float32) bool
	sdlGetWindowBordersSize    func(*Window, *int32, *int32, *int32, *int32) bool
	sdlGetWindowDisplayScale   func(*Window) float32
	sdlGetWindowFlags          func(*Window) WindowFlags
	sdlGetWindowFromEvent      func(*Event) *Window
	sdlGetWindowFromID         func(WindowID) *Window
	sdlGetWindowFullscreenMode func(*Window) *DisplayMode
	// sdlGetWindowICCProfile                   func(*Window, *uint64) unsafe.Pointer
	sdlGetWindowID                func(*Window) WindowID
	sdlGetWindowKeyboardGrab      func(*Window) bool
	sdlGetWindowMaximumSize       func(*Window, *int32, *int32) bool
	sdlGetWindowMinimumSize       func(*Window, *int32, *int32) bool
	sdlGetWindowMouseGrab         func(*Window) bool
	sdlGetWindowMouseRect         func(*Window) *Rect
	sdlGetWindowOpacity           func(*Window) float32
	sdlGetWindowParent            func(*Window) *Window
	sdlGetWindowPixelDensity      func(*Window) float32
	sdlGetWindowPixelFormat       func(*Window) PixelFormat
	sdlGetWindowPosition          func(*Window, *int32, *int32) bool
	sdlGetWindowProgressState     func(*Window) ProgressState
	sdlGetWindowProgressValue     func(*Window) float32
	sdlGetWindowProperties        func(*Window) PropertiesID
	sdlGetWindowRelativeMouseMode func(*Window) bool
	sdlGetWindows                 func(*int32) **Window
	sdlGetWindowSafeArea          func(*Window, *Rect) bool
	sdlGetWindowSize              func(*Window, *int32, *int32) bool
	sdlGetWindowSizeInPixels      func(*Window, *int32, *int32) bool
	sdlGetWindowSurface           func(*Window) *Surface
	sdlGetWindowSurfaceVSync      func(*Window, *int32) bool
	sdlGetWindowTitle             func(*Window) string
	sdlGLCreateContext            func(*Window) GLContext
	sdlGLDestroyContext           func(GLContext) bool
	// sdlGL_ExtensionSupported                 func(string) bool
	sdlGLGetAttribute      func(GLAttr, *int32) bool
	sdlGLGetCurrentContext func() GLContext
	sdlGLGetCurrentWindow  func() *Window
	// sdlGL_GetProcAddress                     func(string) FunctionPointer
	// sdlGL_GetSwapInterval                    func(*int32) bool
	// sdlGL_LoadLibrary                        func(string) bool
	// sdlGL_MakeCurrent                        func(*Window, GLContext) bool
	// sdlGL_ResetAttributes                    func()
	sdlGLSetAttribute    func(GLAttr, int32) bool
	sdlGLSetSwapInterval func(interval int32) bool
	sdlGLSwapWindow      func(window *Window) bool
	// sdlGL_UnloadLibrary                      func()
	// sdlGlobDirectory                         func(string, string, GlobFlags, *int32) **byte
	// sdlGlobStorageDirectory                  func(*Storage, string, string, GlobFlags, *int32) **byte
	// sdlGPUSupportsProperties                 func(PropertiesID) bool
	// sdlGPUSupportsShaderFormats              func(GPUShaderFormat, string) bool
	// sdlGPUTextureFormatTexelBlockSize        func(GPUTextureFormat) uint32
	// sdlGPUTextureSupportsFormat              func(*GPUDevice, GPUTextureFormat, GPUTextureType, GPUTextureUsageFlags) bool
	// sdlGPUTextureSupportsSampleCount         func(*GPUDevice, GPUTextureFormat, GPUSampleCount) bool
	// sdlGUIDToString                          func(GUID, string, int32)
	// sdlHapticEffectSupported                 func(*Haptic, *HapticEffect) bool
	// sdlHapticRumbleSupported                 func(*Haptic) bool
	sdlHasAltiVec func() bool
	sdlHasARMSIMD func() bool
	sdlHasAVX     func() bool
	sdlHasAVX2    func() bool
	sdlHasAVX512F func() bool
	// sdlHasClipboardData                      func(string) bool
	// sdlHasClipboardText                      func() bool
	sdlHasEvent  func(EventType) bool
	sdlHasEvents func(EventType, EventType) bool
	// sdlHasExactlyOneBitSet32                 func(uint32) bool
	// sdlHasGamepad                            func() bool
	sdlHasJoystick func() bool
	sdlHasKeyboard func() bool
	sdlHasLASX     func() bool
	sdlHasLSX      func() bool
	sdlHasMMX      func() bool
	sdlHasMouse    func() bool
	sdlHasNEON     func() bool
	// sdlHasPrimarySelectionText               func() bool
	sdlHasProperty              func(PropertiesID, string) bool
	sdlHasRectIntersection      func(*Rect, *Rect) bool
	sdlHasRectIntersectionFloat func(*FRect, *FRect) bool
	sdlHasScreenKeyboardSupport func() bool
	sdlHasSSE                   func() bool
	sdlHasSSE2                  func() bool
	sdlHasSSE3                  func() bool
	sdlHasSSE41                 func() bool
	sdlHasSSE42                 func() bool
	// sdlhid_ble_scan                          func(bool)
	// sdlhid_close                             func(*hid_device) int32
	// sdlhid_device_change_count               func() uint32
	// sdlhid_enumerate                         func(uint16, uint16) *hid_device_info
	// sdlhid_exit                              func() int32
	// sdlhid_free_enumeration                  func(*hid_device_info)
	// sdlhid_get_device_info                   func(*hid_device) *hid_device_info
	// sdlhid_get_feature_report                func(*hid_device, *uint8, uint64) int32
	// sdlhid_get_indexed_string                func(*hid_device, int32, *wchar_t, uint64) int32
	// sdlhid_get_input_report                  func(*hid_device, *uint8, uint64) int32
	// sdlhid_get_manufacturer_string           func(*hid_device, *wchar_t, uint64) int32
	// sdlhid_get_product_string                func(*hid_device, *wchar_t, uint64) int32
	// sdlhid_get_properties                		func(*hid_device) PropertiesID
	// sdlhid_get_report_descriptor             func(*hid_device, *uint8, uint64) int32
	// sdlhid_get_serial_number_string          func(*hid_device, *wchar_t, uint64) int32
	// sdlhid_init                              func() int32
	// sdlhid_open                              func(uint16, uint16, *wchar_t) *hid_device
	// sdlhid_open_path                         func(string) *hid_device
	// sdlhid_read                              func(*hid_device, *uint8, uint64) int32
	// sdlhid_read_timeout                      func(*hid_device, *uint8, uint64, int32) int32
	// sdlhid_send_feature_report               func(*hid_device, *uint8, uint64) int32
	// sdlhid_set_nonblocking                   func(*hid_device, int32) int32
	// sdlhid_write                             func(*hid_device, *uint8, uint64) int32
	sdlHideCursor func() bool
	sdlHideWindow func(*Window) bool
	// sdliconv                                 func(iconv_t, **byte, *uint64, **byte, *uint64) uint64
	// sdliconv_close                           func(iconv_t) int32
	// sdliconv_open                            func(string, string) iconv_t
	// sdliconv_string                          func(string, string, string, uint64) string
	sdlInit func(InitFlags) bool
	// sdlInitHapticRumble                      func(*Haptic) bool
	// sdlInitSubSystem                         func(InitFlags) bool
	// sdlInsertGPUDebugLabel                   func(*GPUCommandBuffer, string)
	// sdlInsertTrayEntryAt                     func(*TrayMenu, int32, string, TrayEntryFlags) *TrayEntry
	sdlIOFromConstMem func([]byte, int) *IOStream
	// sdlIOFromDynamicMem                      func() *IOStream
	sdlIOFromFile func(string, string) *IOStream
	sdlIOFromMem  func([]byte, int) *IOStream
	// sdlIOprintf                              func(*IOStream, string) uint64
	// sdlIOvprintf                             func(*IOStream, string, va_list) uint64
	// sdlisalnum                               func(int32) int32
	// sdlisalpha                               func(int32) int32
	// sdlIsAudioDevicePhysical                 func(AudioDeviceID) bool
	// sdlIsAudioDevicePlayback                 func(AudioDeviceID) bool
	// sdlisblank                               func(int32) int32
	// sdliscntrl                               func(int32) int32
	// sdlisdigit                               func(int32) int32
	// sdlIsGamepad                             func(JoystickID) bool
	// sdlisgraph                               func(int32) int32
	// sdlisinf                                 func(float64) int32
	// sdlisinff                                func(float32) int32
	// sdlIsJoystickHaptic                      func(*Joystick) bool
	// sdlIsJoystickVirtual func(JoystickID) bool
	// sdlislower                               func(int32) int32
	sdlIsMainThread func() bool
	// sdlIsMouseHaptic                         func() bool
	// sdlisnan                                 func(float64) int32
	// sdlisnanf                                func(float32) int32
	// sdlisprint                               func(int32) int32
	// sdlispunct                               func(int32) int32
	// sdlisspace                               func(int32) int32
	// sdlIsTablet                              func() bool
	// sdlIsTV                                  func() bool
	// sdlisupper                               func(int32) int32
	// sdlisxdigit                              func(int32) int32
	// sdlitoa                                  func(int32, string, int32) string
	sdlJoystickConnected     func(*Joystick) bool
	sdlJoystickEventsEnabled func() bool
	// sdlKillProcess                           func(*Process, bool) bool
	// sdllltoa                                 func(int64, string, int32) string
	sdlLoadBMP   func(string) *Surface
	sdlLoadBMPIO func(*IOStream, bool) *Surface
	sdlLoadFile  func(string, *uint64) unsafe.Pointer
	// sdlLoadFile_IO                           func(*IOStream, *uint64, bool) unsafe.Pointer
	// sdlLoadFileAsync                         func(string, *AsyncIOQueue, unsafe.Pointer) bool
	// sdlLoadFunction                          func(*SharedObject, string) FunctionPointer
	// sdlLoadObject                            func(string) *SharedObject
	sdlLoadPNG func(string) *Surface
	// sdlLoadPNGIO   func(*IOStream, bool) *Surface
	sdlLoadSurface func(string) *Surface
	// sdlLoadSurfaceIO func(*IOStream, bool) *Surface
	sdlLoadWAV   func(string, *AudioSpec, **uint8, *uint32) bool
	sdlLoadWAVIO func(*IOStream, bool, *AudioSpec, **uint8, *uint32) bool
	// sdlLockAudioStream                       func(*AudioStream) bool
	sdlLockJoysticks func()
	// sdlLockMutex                             func(*Mutex)
	sdlLockProperties func(PropertiesID) bool
	// sdlLockRWLockForReading                  func(*RWLock)
	// sdlLockRWLockForWriting                  func(*RWLock)
	// sdlLockSpinlock                          func(*SpinLock)
	sdlLockSurface          func(*Surface) bool
	sdlLockTexture          func(*Texture, *Rect, *unsafe.Pointer, *int32) bool
	sdlLockTextureToSurface func(*Texture, *Rect, **Surface) bool
	sdlLog                  func(string)
	// sdllog                                   func(float64) float64
	// sdllog10                                 func(float64) float64
	// sdllog10f                                func(float32) float32
	// sdlLogCritical                           func(int32, string)
	// sdlLogDebug                              func(int32, string)
	sdlLogError func(LogCategory, string)
	// sdllogf                                  func(float32) float32
	// sdlLogInfo                               func(int32, string)
	sdlLogMessage func(LogCategory, LogPriority, string)
	// sdlLogMessageV                           func(int32, LogPriority, string, va_list)
	// sdlLogTrace                              func(int32, string)
	// sdlLogVerbose                            func(int32, string)
	// sdlLogWarn                               func(int32, string)
	// sdllround                                func(float64) int64
	// sdllroundf                               func(float32) int64
	// sdlltoa                                  func(int64, string, int32) string
	// sdlmain                                  func(int32, **byte) int32
	// sdlmalloc                                func(uint64) unsafe.Pointer
	sdlMapGPUTransferBuffer func(*GPUDevice, *GPUTransferBuffer, bool) unsafe.Pointer
	sdlMapRGB               func(*PixelFormatDetails, *Palette, uint8, uint8, uint8) uint32
	// sdlMapRGBA                               func(*PixelFormatDetails, *Palette, uint8, uint8, uint8, uint8) uint32
	sdlMapSurfaceRGB func(*Surface, uint8, uint8, uint8) uint32
	// sdlMapSurfaceRGBA                        func(*Surface, uint8, uint8, uint8, uint8) uint32
	sdlMaximizeWindow func(*Window) bool
	// sdlmemcmp                                func(unsafe.Pointer, unsafe.Pointer, uint64) int32
	// sdlmemcpy                                func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	// sdlmemmove                               func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	// sdlMemoryBarrierAcquireFunction          func()
	// sdlMemoryBarrierReleaseFunction          func()
	// sdlmemset                                func(unsafe.Pointer, int32, uint64) unsafe.Pointer
	// sdlmemset4                               func(unsafe.Pointer, uint32, uint64) unsafe.Pointer
	// sdlMetal_CreateView                      func(*Window) MetalView
	// sdlMetal_DestroyView                     func(MetalView)
	// sdlMetal_GetLayer                        func(MetalView) unsafe.Pointer
	sdlMinimizeWindow func(*Window) bool
	// sdlMixAudio                              func(*uint8, *uint8, AudioFormat, uint32, float32) bool
	// sdlmodf                                  func(float64, *double) float64
	// sdlmodff                                 func(float32, *float32) float32
	// sdlMostSignificantBitIndex32             func(uint32) int32
	// sdlmurmur3_32                            func(unsafe.Pointer, uint64, uint32) uint32
	// sdlOnApplicationDidEnterBackground       func()
	// sdlOnApplicationDidEnterForeground       func()
	// sdlOnApplicationDidReceiveMemoryWarning  func()
	// sdlOnApplicationWillEnterBackground      func()
	// sdlOnApplicationWillEnterForeground      func()
	// sdlOnApplicationWillTerminate            func()
	// sdlOpenAudioDevice                       func(AudioDeviceID, *AudioSpec) AudioDeviceID
	sdlOpenAudioDeviceStream func(AudioDeviceID, *AudioSpec, AudioStreamCallback, unsafe.Pointer) *AudioStream
	sdlOpenCamera            func(CameraID, *CameraSpec) *Camera
	// sdlOpenFileStorage                       func(string) *Storage
	sdlOpenGamepad func(JoystickID) *Gamepad
	// sdlOpenHaptic                            func(HapticID) *Haptic
	// sdlOpenHapticFromJoystick                func(*Joystick) *Haptic
	// sdlOpenHapticFromMouse                   func() *Haptic
	// sdlOpenIO                                func(*IOStreamInterface, unsafe.Pointer) *IOStream
	sdlOpenJoystick func(JoystickID) *Joystick
	// sdlOpenSensor                            func(SensorID) *Sensor
	// sdlOpenStorage                           func(*StorageInterface, unsafe.Pointer) *Storage
	// sdlOpenTitleStorage                      func(string, PropertiesID) *Storage
	sdlOpenURL func(string) bool
	// sdlOpenUserStorage                       func(string, string, PropertiesID) *Storage
	// sdlOutOfMemory                           func() bool
	// sdlPauseAudioDevice                      func(AudioDeviceID) bool
	sdlPauseAudioStreamDevice func(stream *AudioStream) bool
	// sdlPauseHaptic                           func(*Haptic) bool
	sdlPeepEvents func(*Event, int32, EventAction, EventType, EventType) int32
	// sdlPlayHapticRumble                      func(*Haptic, float32, uint32) bool
	sdlPollEvent func(event *Event) bool
	// sdlPopGPUDebugGroup                      func(*GPUCommandBuffer)
	// sdlpow                                   func(float64, float64) float64
	// sdlpowf                                  func(float32, float32) float32
	// sdlPremultiplyAlpha                      func(int32, int32, PixelFormat, unsafe.Pointer, int32, PixelFormat, unsafe.Pointer, int32, bool) bool
	// sdlPremultiplySurfaceAlpha               func(*Surface, bool) bool
	sdlPumpEvents func()
	sdlPushEvent  func(*Event) bool
	// sdlPushGPUComputeUniformData             func(*GPUCommandBuffer, uint32, unsafe.Pointer, uint32)
	// sdlPushGPUDebugGroup                     func(*GPUCommandBuffer, string)
	sdlPushGPUFragmentUniformData func(*GPUCommandBuffer, uint32, unsafe.Pointer, uint32)
	sdlPushGPUVertexUniformData   func(*GPUCommandBuffer, uint32, unsafe.Pointer, uint32)
	sdlPutAudioStreamData         func(stream *AudioStream, buf *uint8, len int32) bool
	// sdlPutAudioStreamDataNoCopy   						uintptr
	// sdlPutAudioStreamPlanarData   						uintptr
	// sdlqsort                                 func(unsafe.Pointer, uint64, uint64, CompareCallback)
	// sdlqsort_r                               func(unsafe.Pointer, uint64, uint64, CompareCallback_r, unsafe.Pointer)
	// sdlQueryGPUFence                         func(*GPUDevice, *GPUFence) bool
	sdlQuit          func()
	sdlQuitSubSystem func(InitFlags)
	sdlRaiseWindow   func(*Window) bool
	// sdlrand                                  func(int32) int32
	// sdlrand_bits                             func() uint32
	// sdlrand_bits_r                           func(*uint64) uint32
	// sdlrand_r                                func(*uint64, int32) int32
	// sdlrandf                                 func() float32
	// sdlrandf_r                               func(*uint64) float32
	// sdlReadAsyncIO                           func(*AsyncIO, unsafe.Pointer, uint64, uint64, *AsyncIOQueue, unsafe.Pointer) bool
	// sdlReadIO                                func(*IOStream, unsafe.Pointer, uint64) uint64
	// sdlReadProcess                           func(*Process, *uint64, *int32) unsafe.Pointer
	// sdlReadS16BE                             func(*IOStream, *int16) bool
	// sdlReadS16LE                             func(*IOStream, *int16) bool
	// sdlReadS32BE                             func(*IOStream, *int32) bool
	// sdlReadS32LE                             func(*IOStream, *int32) bool
	// sdlReadS64BE                             func(*IOStream, *int64) bool
	// sdlReadS64LE                             func(*IOStream, *int64) bool
	// sdlReadS8                                func(*IOStream, *int8) bool
	// sdlReadStorageFile                       func(*Storage, string, unsafe.Pointer, uint64) bool
	// sdlReadSurfacePixel                      func(*Surface, int32, int32, *uint8, *uint8, *uint8, *uint8) bool
	// sdlReadSurfacePixelFloat                 func(*Surface, int32, int32, *float32, *float32, *float32, *float32) bool
	// sdlReadU16BE                             func(*IOStream, *uint16) bool
	// sdlReadU16LE                             func(*IOStream, *uint16) bool
	// sdlReadU32BE                             func(*IOStream, *uint32) bool
	// sdlReadU32LE                             func(*IOStream, *uint32) bool
	// sdlReadU64BE                             func(*IOStream, *uint64) bool
	// sdlReadU64LE                             func(*IOStream, *uint64) bool
	// sdlReadU8                                func(*IOStream, *uint8) bool
	// sdlrealloc                               func(unsafe.Pointer, uint64) unsafe.Pointer
	sdlRegisterEvents     func(int32) uint32
	sdlReleaseCameraFrame func(*Camera, *Surface)
	sdlReleaseGPUBuffer   func(*GPUDevice, *GPUBuffer)
	// sdlReleaseGPUComputePipeline             func(*GPUDevice, *GPUComputePipeline)
	// sdlReleaseGPUFence                       func(*GPUDevice, *GPUFence)
	sdlReleaseGPUGraphicsPipeline func(*GPUDevice, *GPUGraphicsPipeline)
	sdlReleaseGPUSampler          func(*GPUDevice, *GPUSampler)
	sdlReleaseGPUShader           func(*GPUDevice, *GPUShader)
	sdlReleaseGPUTexture          func(*GPUDevice, *GPUTexture)
	sdlReleaseGPUTransferBuffer   func(*GPUDevice, *GPUTransferBuffer)
	sdlReleaseWindowFromGPUDevice func(*GPUDevice, *Window)
	// sdlReloadGamepadMappings                 func() bool
	sdlRemoveEventWatch   func(EventFilter, unsafe.Pointer)
	sdlRemoveHintCallback func(string, HintCallback, unsafe.Pointer)
	// sdlRemovePath                            func(string) bool
	// sdlRemoveStoragePath                     func(*Storage, string) bool
	sdlRemoveSurfaceAlternateImages func(*Surface)
	// sdlRemoveTimer                           func(TimerID) bool
	// sdlRemoveTrayEntry                       func(*TrayEntry)
	// sdlRenamePath                            func(string, string) bool
	// sdlRenameStoragePath                     func(*Storage, string, string) bool
	sdlRenderClear                 func(renderer *Renderer) bool
	sdlRenderClipEnabled           func(*Renderer) bool
	sdlRenderCoordinatesFromWindow func(*Renderer, float32, float32, *float32, *float32) bool
	sdlRenderCoordinatesToWindow   func(*Renderer, float32, float32, *float32, *float32) bool
	sdlRenderDebugText             func(*Renderer, float32, float32, string) bool
	sdlRenderDebugTextFormat       func(*Renderer, float32, float32, string) bool
	sdlRenderFillRect              func(renderer *Renderer, rect *FRect) bool
	sdlRenderFillRects             func(renderer *Renderer, rects []FRect) bool
	sdlRenderGeometry              func(renderer *Renderer, texture *Texture, vertices []Vertex, indices []int32) bool
	sdlRenderGeometryRaw           func(renderer *Renderer, texture *Texture, xy []FPoint, color []FColor, uv []FPoint, indices []int32) bool
	sdlRenderLine                  func(*Renderer, float32, float32, float32, float32) bool
	sdlRenderLines                 func(renderer *Renderer, points []FPoint) bool
	sdlRenderPoint                 func(*Renderer, float32, float32) bool
	sdlRenderPoints                func(renderer *Renderer, points []FPoint) bool
	sdlRenderPresent               func(renderer *Renderer) bool
	sdlRenderReadPixels            func(*Renderer, *Rect) *Surface
	sdlRenderRect                  func(renderer *Renderer, rect *FRect) bool
	sdlRenderRects                 func(renderer *Renderer, rects []FRect) bool
	sdlRenderTexture               func(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect) bool
	sdlRenderTexture9Grid          func(*Renderer, *Texture, *FRect, float32, float32, float32, float32, float32, *FRect) bool
	sdlRenderTexture9GridTiled     func(*Renderer, *Texture, *FRect, float32, float32, float32, float32, float32, *FRect, float32) bool
	sdlRenderTextureAffine         func(renderer *Renderer, texture *Texture, srcrect *FRect, origin *FPoint, right *FPoint, down *FPoint) bool
	sdlRenderTextureRotated        func(*Renderer, *Texture, *FRect, *FRect, float64, *FPoint, FlipMode) bool
	sdlRenderTextureTiled          func(*Renderer, *Texture, *FRect, float32, *FRect) bool
	sdlRenderViewportSet           func(*Renderer) bool
	// sdlReportAssertion                       func(*AssertData, string, string, int32) AssertState
	// sdlResetAssertionReport                  func()
	sdlResetHint          func(string) bool
	sdlResetHints         func()
	sdlResetKeyboard      func()
	sdlResetLogPriorities func()
	sdlRestoreWindow      func(*Window) bool
	// sdlResumeAudioDevice                     func(AudioDeviceID) bool
	sdlResumeAudioStreamDevice func(stream *AudioStream) bool
	// sdlResumeHaptic                          func(*Haptic) bool
	sdlRotateSurface func(*Surface, float32) *Surface
	// sdlround                                 func(float64) float64
	// sdlroundf                                func(float32) float32
	// sdlRumbleGamepad                         func(*Gamepad, uint16, uint16, uint32) bool
	// sdlRumbleGamepadTriggers                 func(*Gamepad, uint16, uint16, uint32) bool
	sdlRumbleJoystick         func(*Joystick, uint16, uint16, uint32) bool
	sdlRumbleJoystickTriggers func(*Joystick, uint16, uint16, uint32) bool
	sdlRunApp                 func(int32, **byte, MainFunc, unsafe.Pointer) int32
	// sdlRunHapticEffect                       func(*Haptic, int32, uint32) bool
	// sdlRunOnMainThread                       func(MainThreadCallback, unsafe.Pointer, bool) bool
	sdlSaveBMP   func(*Surface, string) bool
	sdlSaveBMPIO func(*Surface, *IOStream, bool) bool
	// sdlSaveFile                              func(string, unsafe.Pointer, uint64) bool
	// sdlSaveFile_IO                           func(*IOStream, unsafe.Pointer, uint64, bool) bool
	sdlSavePNG func(*Surface, string) bool
	// sdlSavePNGIO func(*Surface, *IOStream, bool) bool
	// sdlscalbn                                func(float64, int32) float64
	// sdlscalbnf                               func(float32, int32) float32
	sdlScaleSurface        func(*Surface, int32, int32, ScaleMode) *Surface
	sdlScreenKeyboardShown func(*Window) bool
	sdlScreenSaverEnabled  func() bool
	// sdlSeekIO                                func(*IOStream, int64, IOWhence) int64
	// sdlSendGamepadEffect                     func(*Gamepad, unsafe.Pointer, int32) bool
	sdlSendJoystickEffect func(*Joystick, unsafe.Pointer, int32) bool
	// sdlSendJoystickVirtualSensorData func(*Joystick, SensorType, uint64, *float32, int32) bool
	// sdlSetAppMetadata                        func(string, string, string) bool
	// sdlSetAppMetadataProperty                func(string, string) bool
	// sdlSetAssertionHandler                   func(AssertionHandler, unsafe.Pointer)
	// sdlSetAtomicInt                          func(*AtomicInt, int32) int32
	// sdlSetAtomicPointer                      func(*unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	// sdlSetAtomicU32                          func(*AtomicU32, uint32) uint32
	// sdlSetAudioDeviceGain                    func(AudioDeviceID, float32) bool
	// sdlSetAudioPostmixCallback               func(AudioDeviceID, AudioPostmixCallback, unsafe.Pointer) bool
	sdlSetAudioStreamFormat         func(*AudioStream, *AudioSpec, *AudioSpec) bool
	sdlSetAudioStreamFrequencyRatio func(*AudioStream, float32) bool
	sdlSetAudioStreamGain           func(*AudioStream, float32) bool
	// sdlSetAudioStreamGetCallback             func(*AudioStream, AudioStreamCallback, unsafe.Pointer) bool
	// sdlSetAudioStreamInputChannelMap         func(*AudioStream, *int32, int32) bool
	// sdlSetAudioStreamOutputChannelMap        func(*AudioStream, *int32, int32) bool
	// sdlSetAudioStreamPutCallback             func(*AudioStream, AudioStreamCallback, unsafe.Pointer) bool
	sdlSetBooleanProperty func(PropertiesID, string, bool) bool
	// sdlSetClipboardData                      func(ClipboardDataCallback, ClipboardCleanupCallback, unsafe.Pointer, **byte, uint64) bool
	// sdlSetClipboardText                      func(string) bool
	// sdlSetCurrentThreadPriority              func(ThreadPriority) bool
	sdlSetCursor                  func(*Cursor) bool
	sdlSetDefaultTextureScaleMode func(*Renderer, ScaleMode) bool
	// sdlsetenv_unsafe                         func(string, string, int32) int32
	// sdlSetEnvironmentVariable                func(*Environment, string, string, bool) bool
	sdlSetError func(string) bool
	// sdlSetErrorV                             func(string, va_list) bool
	sdlSetEventEnabled  func(EventType, bool)
	sdlSetEventFilter   func(EventFilter, unsafe.Pointer)
	sdlSetFloatProperty func(PropertiesID, string, float32) bool
	// sdlSetGamepadEventsEnabled               func(bool)
	// sdlSetGamepadLED                         func(*Gamepad, uint8, uint8, uint8) bool
	// sdlSetGamepadMapping                     func(JoystickID, string) bool
	// sdlSetGamepadPlayerIndex                 func(*Gamepad, int32) bool
	// sdlSetGamepadSensorEnabled               func(*Gamepad, SensorType, bool) bool
	// sdlSetGPUAllowedFramesInFlight           func(*GPUDevice, uint32) bool
	// sdlSetGPUBlendConstants                  func(*GPURenderPass, FColor)
	sdlSetGPUBufferName func(*GPUDevice, *GPUBuffer, string)
	// sdlSetGPURenderState func(*Renderer, *GPURenderState) bool
	// sdlSetGPURenderStateFragmentUniforms func(*GPURenderState, uint32, uintptr, uint32) bool
	sdlSetGPUScissor func(*GPURenderPass, *Rect)
	// sdlSetGPUStencilReference                func(*GPURenderPass, uint8)
	sdlSetGPUSwapchainParameters func(*GPUDevice, *Window, GPUSwapchainComposition, GPUPresentMode) bool
	// sdlSetGPUTextureName                     func(*GPUDevice, *GPUTexture, string)
	sdlSetGPUViewport func(*GPURenderPass, *GPUViewport)
	// sdlSetHapticAutocenter                   func(*Haptic, int32) bool
	// sdlSetHapticGain                         func(*Haptic, int32) bool
	sdlSetHint             func(string, string) bool
	sdlSetHintWithPriority func(string, string, HintPriority) bool
	// sdlSetInitialized                        func(*InitState, bool)
	sdlSetJoystickEventsEnabled func(bool)
	sdlSetJoystickLED           func(*Joystick, uint8, uint8, uint8) bool
	sdlSetJoystickPlayerIndex   func(*Joystick, int32) bool
	// sdlSetJoystickVirtualAxis                func(*Joystick, int32, int16) bool
	// sdlSetJoystickVirtualBall                func(*Joystick, int32, int16, int16) bool
	// sdlSetJoystickVirtualButton              func(*Joystick, int32, bool) bool
	// sdlSetJoystickVirtualHat                 func(*Joystick, int32, uint8) bool
	// sdlSetJoystickVirtualTouchpad            func(*Joystick, int32, int32, bool, float32, float32, float32) bool
	// sdlSetLinuxThreadPriority                func(int64, int32) bool
	// sdlSetLinuxThreadPriorityAndPolicy       func(int64, int32, int32) bool
	sdlSetLogOutputFunction func(LogOutputFunction, unsafe.Pointer)
	sdlSetLogPriorities     func(LogPriority)
	sdlSetLogPriority       func(int32, LogPriority)
	// sdlSetLogPriorityPrefix                  func(LogPriority, string) bool
	// sdlSetMainReady                          func()
	// sdlSetMemoryFunctions                    func(malloc_func, calloc_func, realloc_func, free_func) bool
	sdlSetModState                   func(Keymod)
	sdlSetNumberProperty             func(PropertiesID, string, int64) bool
	sdlSetPaletteColors              func(*Palette, *Color, int32, int32) bool
	sdlSetPointerProperty            func(PropertiesID, string, unsafe.Pointer) bool
	sdlSetPointerPropertyWithCleanup func(PropertiesID, string, unsafe.Pointer, CleanupPropertyCallback, unsafe.Pointer) bool
	// sdlSetPrimarySelectionText               func(string) bool
	// sdlSetRelativeMouseTransform    func(MouseMotionTransformCallback, uintptr) bool
	sdlSetRenderClipRect            func(*Renderer, *Rect) bool
	sdlSetRenderColorScale          func(*Renderer, float32) bool
	sdlSetRenderDrawBlendMode       func(*Renderer, BlendMode) bool
	sdlSetRenderDrawColor           func(renderer *Renderer, r, g, b, a uint8) bool
	sdlSetRenderDrawColorFloat      func(*Renderer, float32, float32, float32, float32) bool
	sdlSetRenderLogicalPresentation func(*Renderer, int32, int32, RendererLogicalPresentation) bool
	sdlSetRenderScale               func(*Renderer, float32, float32) bool
	sdlSetRenderTarget              func(*Renderer, *Texture) bool
	sdlSetRenderTextureAddressMode  func(*Renderer, TextureAddressMode, TextureAddressMode) bool
	sdlSetRenderViewport            func(*Renderer, *Rect) bool
	sdlSetRenderVSync               func(*Renderer, int32) bool
	sdlSetScancodeName              func(Scancode, string) bool
	sdlSetStringProperty            func(PropertiesID, string, string) bool
	sdlSetSurfaceAlphaMod           func(*Surface, uint8) bool
	sdlSetSurfaceBlendMode          func(*Surface, BlendMode) bool
	sdlSetSurfaceClipRect           func(*Surface, *Rect) bool
	sdlSetSurfaceColorKey           func(*Surface, bool, uint32) bool
	sdlSetSurfaceColorMod           func(*Surface, uint8, uint8, uint8) bool
	sdlSetSurfaceColorspace         func(*Surface, Colorspace) bool
	sdlSetSurfacePalette            func(*Surface, *Palette) bool
	sdlSetSurfaceRLE                func(*Surface, bool) bool
	sdlSetTextInputArea             func(*Window, *Rect, int32) bool
	sdlSetTextureAlphaMod           func(texture *Texture, alpha uint8) bool
	sdlSetTextureAlphaModFloat      func(*Texture, float32) bool
	sdlSetTextureBlendMode          func(*Texture, BlendMode) bool
	sdlSetTextureColorMod           func(texture *Texture, r uint8, g uint8, b uint8) bool
	sdlSetTextureColorModFloat      func(*Texture, float32, float32, float32) bool
	sdlSetTexturePalette            func(*Texture, *Palette) bool
	sdlSetTextureScaleMode          func(*Texture, ScaleMode) bool
	// sdlSetTLS                                func(*TLSID, unsafe.Pointer, TLSDestructorCallback) bool
	// sdlSetTrayEntryCallback                  func(*TrayEntry, TrayCallback, unsafe.Pointer)
	// sdlSetTrayEntryChecked                   func(*TrayEntry, bool)
	// sdlSetTrayEntryEnabled                   func(*TrayEntry, bool)
	// sdlSetTrayEntryLabel                     func(*TrayEntry, string)
	// sdlSetTrayIcon                           func(*Tray, *Surface)
	// sdlSetTrayTooltip                        func(*Tray, string)
	sdlSetWindowAlwaysOnTop func(*Window, bool) bool
	sdlSetWindowAspectRatio func(*Window, float32, float32) bool
	sdlSetWindowBordered    func(*Window, bool) bool
	// sdlSetWindowFillDocument   func(*Window, bool) bool
	sdlSetWindowFocusable         func(*Window, bool) bool
	sdlSetWindowFullscreen        func(*Window, bool) bool
	sdlSetWindowFullscreenMode    func(*Window, *DisplayMode) bool
	sdlSetWindowHitTest           func(*Window, func(*Window, *Point, unsafe.Pointer) uintptr, unsafe.Pointer) bool
	sdlSetWindowIcon              func(*Window, *Surface) bool
	sdlSetWindowKeyboardGrab      func(*Window, bool) bool
	sdlSetWindowMaximumSize       func(*Window, int32, int32) bool
	sdlSetWindowMinimumSize       func(*Window, int32, int32) bool
	sdlSetWindowModal             func(*Window, bool) bool
	sdlSetWindowMouseGrab         func(*Window, bool) bool
	sdlSetWindowMouseRect         func(*Window, *Rect) bool
	sdlSetWindowOpacity           func(*Window, float32) bool
	sdlSetWindowParent            func(*Window, *Window) bool
	sdlSetWindowPosition          func(*Window, int32, int32) bool
	sdlSetWindowProgressState     func(*Window, ProgressState) bool
	sdlSetWindowProgressValue     func(*Window, float32) bool
	sdlSetWindowRelativeMouseMode func(*Window, bool) bool
	sdlSetWindowResizable         func(*Window, bool) bool
	// sdlSetWindowShape                        func(*Window, *Surface) bool
	sdlSetWindowSize         func(*Window, int32, int32) bool
	sdlSetWindowSurfaceVSync func(*Window, int32) bool
	sdlSetWindowTitle        func(*Window, string) bool
	// sdlSetX11EventHook                       func(X11EventHook, unsafe.Pointer)
	// sdlShouldInit                            func(*InitState) bool
	// sdlShouldQuit                            func(*InitState) bool
	sdlShowCursor                   func() bool
	sdlShowFileDialogWithProperties func(FileDialogType, DialogFileCallback, unsafe.Pointer, PropertiesID)
	sdlShowMessageBox               func(*MessageBoxData, *int32) bool
	sdlShowOpenFileDialog           func(DialogFileCallback, unsafe.Pointer, *Window, []DialogFileFilter, int32, *byte, bool)
	sdlShowOpenFolderDialog         func(DialogFileCallback, unsafe.Pointer, *Window, string, bool)
	sdlShowSaveFileDialog           func(DialogFileCallback, unsafe.Pointer, *Window, []DialogFileFilter, int32, string)
	sdlShowSimpleMessageBox         func(MessageBoxFlags, string, string, *Window) bool
	sdlShowWindow                   func(*Window) bool
	sdlShowWindowSystemMenu         func(*Window, int32, int32) bool
	// sdlSignalAsyncIOQueue                    func(*AsyncIOQueue)
	// sdlSignalCondition                       func(*Condition)
	// sdlSignalSemaphore                       func(*Semaphore)
	// sdlsin                                   func(float64) float64
	// sdlsinf                                  func(float32) float32
	// sdlsize_add_check_overflow               func(uint64, uint64, *uint64) bool
	// sdlsize_add_check_overflow_builtin       func(uint64, uint64, *uint64) bool
	// sdlsize_mul_check_overflow               func(uint64, uint64, *uint64) bool
	// sdlsize_mul_check_overflow_builtin       func(uint64, uint64, *uint64) bool
	// sdlsnprintf                              func(string, uint64, string) int32
	// sdlsqrt                                  func(float64) float64
	// sdlsqrtf                                 func(float32) float32
	// sdlsrand                                 func(uint64)
	// sdlsscanf                                func(string, string) int32
	sdlStartTextInput               func(*Window) bool
	sdlStartTextInputWithProperties func(*Window, PropertiesID) bool
	// sdlStepBackUTF8                          func(string, **byte) uint32
	// sdlStepUTF8                              func(**byte, *uint64) uint32
	// sdlStopHapticEffect                      func(*Haptic, int32) bool
	// sdlStopHapticEffects                     func(*Haptic) bool
	// sdlStopHapticRumble                      func(*Haptic) bool
	sdlStopTextInput func(*Window) bool
	// sdlStorageReady                          func(*Storage) bool
	// sdlstrcasecmp                            func(string, string) int32
	// sdlstrcasestr                            func(string, string) string
	// sdlstrchr                                func(string, int32) string
	// sdlstrcmp                                func(string, string) int32
	// sdlstrdup                                func(string) string
	sdlStretchSurface func(*Surface, *Rect, *Surface, *Rect, ScaleMode) bool
	// sdlStringToGUID                          func(string) GUID
	// sdlstrlcat                               func(string, string, uint64) uint64
	// sdlstrlcpy                               func(string, string, uint64) uint64
	// sdlstrlen                                func(string) uint64
	// sdlstrlwr                                func(string) string
	// sdlstrncasecmp                           func(string, string, uint64) int32
	// sdlstrncmp                               func(string, string, uint64) int32
	// sdlstrndup                               func(string, uint64) string
	// sdlstrnlen                               func(string, uint64) uint64
	// sdlstrnstr                               func(string, string, uint64) string
	// sdlstrpbrk                               func(string, string) string
	// sdlstrrchr                               func(string, int32) string
	// sdlstrrev                                func(string) string
	sdlstrstr func(string, string) string
	// sdlstrtod                                func(string, **byte) float64
	// sdlstrtok_r                              func(string, string, **byte) string
	// sdlstrtol                                func(string, **byte, int32) int64
	// sdlstrtoll                               func(string, **byte, int32) int64
	// sdlstrtoul                               func(string, **byte, int32) uint64
	// sdlstrtoull                              func(string, **byte, int32) uint64
	// sdlstrupr                                func(string) string
	sdlSubmitGPUCommandBuffer func(*GPUCommandBuffer) bool
	// sdlSubmitGPUCommandBufferAndAcquireFence func(*GPUCommandBuffer) *GPUFence
	sdlSurfaceHasAlternateImages func(*Surface) bool
	sdlSurfaceHasColorKey        func(*Surface) bool
	sdlSurfaceHasRLE             func(*Surface) bool
	// sdlSwapFloat                             func(float32) float32
	// sdlswprintf                              func(*wchar_t, uint64, *wchar_t) int32
	sdlSyncWindow func(*Window) bool
	// sdltan                                   func(float64) float64
	// sdltanf                                  func(float32) float32
	// sdlTellIO                                func(*IOStream) int64
	sdlTextInputActive func(*Window) bool
	// sdlTimeFromWindows                       func(uint32, uint32) Time
	// sdlTimeToDateTime                        func(Time, *DateTime, bool) bool
	// sdlTimeToWindows                         func(Time, *uint32, *uint32)
	// sdltolower                               func(int32) int32
	// sdltoupper                               func(int32) int32
	// sdltrunc                                 func(float64) float64
	// sdltruncf                                func(float32) float32
	// sdlTryLockMutex                          func(*Mutex) bool
	// sdlTryLockRWLockForReading               func(*RWLock) bool
	// sdlTryLockRWLockForWriting               func(*RWLock) bool
	// sdlTryLockSpinlock                       func(*SpinLock) bool
	// sdlTryWaitSemaphore                      func(*Semaphore) bool
	// sdlUCS4ToUTF8                            func(uint32, string) string
	// sdluitoa                                 func(uint32, string, int32) string
	// sdlulltoa                                func(uint64, string, int32) string
	// sdlultoa                                 func(uint64, string, int32) string
	// sdlUnbindAudioStream                     func(*AudioStream)
	// sdlUnbindAudioStreams                    func(**AudioStream, int32)
	// sdlUnloadObject                          func(*SharedObject)
	// sdlUnlockAudioStream                     func(*AudioStream) bool
	sdlUnlockJoysticks func()
	// sdlUnlockMutex                           func(*Mutex)
	sdlUnlockProperties func(PropertiesID)
	// sdlUnlockRWLock                          func(*RWLock)
	// sdlUnlockSpinlock                        func(*SpinLock)
	sdlUnlockSurface          func(*Surface)
	sdlUnlockTexture          func(*Texture)
	sdlUnmapGPUTransferBuffer func(*GPUDevice, *GPUTransferBuffer)
	// sdlunsetenv_unsafe                       func(string) int32
	// sdlUnsetEnvironmentVariable              func(*Environment, string) bool
	// sdlUpdateGamepads                        func()
	// sdlUpdateHapticEffect                    func(*Haptic, int32, *HapticEffect) bool
	sdlUpdateJoysticks func()
	sdlUpdateNVTexture func(texture *Texture, rect *Rect, yPlane *uint8, yPitch int32, uvPlane *uint8, uvPitch int32) bool
	// sdlUpdateSensors                         func()
	sdlUpdateTexture       func(texture *Texture, rect *Rect, pixels unsafe.Pointer, pitch int32) bool
	sdlUpdateWindowSurface func(*Window) bool
	// sdlUpdateWindowSurfaceRects              func(*Window, *Rect, int32) bool
	sdlUpdateYUVTexture   func(texture *Texture, rect *Rect, yPlane *uint8, yPitch int32, uPlane *uint8, uPitch int32, vPlane *uint8, vPitch int32) bool
	sdlUploadToGPUBuffer  func(*GPUCopyPass, *GPUTransferBufferLocation, *GPUBufferRegion, bool)
	sdlUploadToGPUTexture func(*GPUCopyPass, *GPUTextureTransferInfo, *GPUTextureRegion, bool)
	// sdlutf8strlcpy                           func(string, string, uint64) uint64
	// sdlutf8strlen                            func(string) uint64
	// sdlutf8strnlen                           func(string, uint64) uint64
	// sdlvasprintf                             func(**byte, string, va_list) int32
	// sdlvsnprintf                             func(string, uint64, string, va_list) int32
	// sdlvsscanf                               func(string, string, va_list) int32
	// sdlvswprintf                             func(*wchar_t, uint64, *wchar_t, va_list) int32
	sdlWaitAndAcquireGPUSwapchainTexture func(*GPUCommandBuffer, *Window, **GPUTexture, *uint32, *uint32) bool
	// sdlWaitAsyncIOResult                     func(*AsyncIOQueue, *AsyncIOOutcome, int32) bool
	// sdlWaitCondition                         func(*Condition, *Mutex)
	// sdlWaitConditionTimeout                  func(*Condition, *Mutex, int32) bool
	sdlWaitEvent        func(*Event) bool
	sdlWaitEventTimeout func(*Event, int32) bool
	// sdlWaitForGPUFences                      func(*GPUDevice, bool, **GPUFence, uint32) bool
	// sdlWaitForGPUIdle                        func(*GPUDevice) bool
	// sdlWaitForGPUSwapchain                   func(*GPUDevice, *Window) bool
	// sdlWaitProcess                           func(*Process, bool, *int32) bool
	// sdlWaitSemaphore                         func(*Semaphore)
	// sdlWaitSemaphoreTimeout                  func(*Semaphore, int32) bool
	// sdlWaitThread                            func(*Thread, *int32)
	sdlWarpMouseGlobal   func(float32, float32) bool
	sdlWarpMouseInWindow func(*Window, float32, float32)
	// sdlWasInit                               func(InitFlags) InitFlags
	// sdlwcscasecmp                            func(*wchar_t, *wchar_t) int32
	// sdlwcscmp                                func(*wchar_t, *wchar_t) int32
	// sdlwcsdup                                func(*wchar_t) *wchar_t
	// sdlwcslcat                               func(*wchar_t, *wchar_t, uint64) uint64
	// sdlwcslcpy                               func(*wchar_t, *wchar_t, uint64) uint64
	// sdlwcslen                                func(*wchar_t) uint64
	// sdlwcsncasecmp                           func(*wchar_t, *wchar_t, uint64) int32
	// sdlwcsncmp                               func(*wchar_t, *wchar_t, uint64) int32
	// sdlwcsnlen                               func(*wchar_t, uint64) uint64
	// sdlwcsnstr                               func(*wchar_t, *wchar_t, uint64) *wchar_t
	// sdlwcsstr                                func(*wchar_t, *wchar_t) *wchar_t
	// sdlwcstol                                func(*wchar_t, **wchar_t, int32) int64
	sdlWindowHasSurface             func(*Window) bool
	sdlWindowSupportsGPUPresentMode func(*GPUDevice, *Window, GPUPresentMode) bool
	// sdlWindowSupportsGPUSwapchainComposition func(*GPUDevice, *Window, GPUSwapchainComposition) bool
	// sdlWriteAsyncIO                          func(*AsyncIO, unsafe.Pointer, uint64, uint64, *AsyncIOQueue, unsafe.Pointer) bool
	// sdlWriteIO                               func(*IOStream, unsafe.Pointer, uint64) uint64
	// sdlWriteS16BE                            func(*IOStream, int16) bool
	// sdlWriteS16LE                            func(*IOStream, int16) bool
	// sdlWriteS32BE                            func(*IOStream, int32) bool
	// sdlWriteS32LE                            func(*IOStream, int32) bool
	// sdlWriteS64BE                            func(*IOStream, int64) bool
	// sdlWriteS64LE                            func(*IOStream, int64) bool
	// sdlWriteS8                               func(*IOStream, int8) bool
	// sdlWriteStorageFile                      func(*Storage, string, unsafe.Pointer, uint64) bool
	// sdlWriteSurfacePixel                     func(*Surface, int32, int32, uint8, uint8, uint8, uint8) bool
	// sdlWriteSurfacePixelFloat                func(*Surface, int32, int32, float32, float32, float32, float32) bool
	// sdlWriteU16BE                            func(*IOStream, uint16) bool
	// sdlWriteU16LE                            func(*IOStream, uint16) bool
	// sdlWriteU32BE                            func(*IOStream, uint32) bool
	// sdlWriteU32LE                            func(*IOStream, uint32) bool
	// sdlWriteU64BE                            func(*IOStream, uint64) bool
	// sdlWriteU64LE                            func(*IOStream, uint64) bool
	// sdlWriteU8 func(*IOStream, uint8) bool

	// NewEventFilter converts the Go function to a C function pointer.
	NewEventFilter                 func(filter func(userdata unsafe.Pointer, event *Event) bool) EventFilter
	NewAudioStreamCallback         func(callback func(userdata unsafe.Pointer, stream *AudioStream, additionalAmount, totalAmount int32)) AudioStreamCallback
	NewDialogFileCallback          func(callback func(userdata unsafe.Pointer, filelist []string, filter int32)) DialogFileCallback
	NewHintCallback                func(callback func(userdata unsafe.Pointer, name, oldValue, newValue string)) HintCallback
	NewLogOutputFunctionCallback   func(callback func(userdata unsafe.Pointer, category LogCategory, priority LogPriority, message string)) LogOutputFunction
	NewCleanupPropertyCallback     func(callback func(userdata, value unsafe.Pointer)) CleanupPropertyCallback
	NewEnumeratePropertiesCallback func(callback func(userdata unsafe.Pointer, props PropertiesID, name string)) EnumeratePropertiesCallback
)
