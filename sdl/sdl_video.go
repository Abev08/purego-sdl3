package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

type DisplayID uint32

type WindowID uint32

type HitTestResult uint32

type HitTest func(window *Window, point *Point, data unsafe.Pointer) HitTestResult

const (
	HitTestNormal HitTestResult = iota
	HitTestDraggable
	HitTestResizeTopLeft
	HitTestResizeTop
	HitTestResizeTopRight
	HitTestResizeRight
	HitTestResizeBottomRight
	HitTestResizeBottom
	HitTestResizeBottomLeft
	HitTestResizeLeft
)

type GLAttr uint32

const (
	GLRedSize GLAttr = iota
	GLGreenSize
	GLBlueSize
	GLAlphaSize
	GLBufferSize
	GLDoubleBuffer
	GLDepthSize
	GLStencilSize
	GLAccumRedSize
	GLAccumGreenSize
	GLAccumBlueSize
	GLAccumAlphaSize
	GLStereo
	GLMultisampleBuffers
	GLMultisampleSamples
	GLAcceleratedVisual
	GLRetainedBacking
	GLContextMajorVersion
	GLContextMinorVersion
	GLContextFlags
	GLContextProfileMask
	GLShareWithCurrentContext
	GLFramebufferSRGBCapable
	GLContextReleaseBehavior
	GLContextResetNotification
	GLContextNoError
	GLFloatBuffers
	GLEGLPlatform
)

type FlashOperation uint32

const (
	FlashCancel FlashOperation = iota
	FlashBriefly
	FlashUntilFocused
)

type DisplayModeData struct{}

type DisplayMode struct {
	DisplayID              DisplayID
	Format                 PixelFormat
	W                      int32
	H                      int32
	PixelDensity           float32
	RefreshRate            float32
	RefreshRateNumerator   int32
	RefreshRateDenominator int32
	Internal               *DisplayModeData // private
}

type DisplayOrientation uint32

const (
	OrientationUnknown DisplayOrientation = iota
	OrientationLandscape
	OrientationLandscapeFlipped
	OrientationPortrait
	OrientationPortraitFlipped
)

type SystemTheme uint32

const (
	SystemThemeUnknown SystemTheme = iota
	SystemThemeLight
	SystemThemeDark
)

type Window struct{}

type WindowFlags uint64

const (
	WindowFullscreen        WindowFlags = 0x0000000000000001
	WindowOpenGL            WindowFlags = 0x0000000000000002
	WindowOccluded          WindowFlags = 0x0000000000000004
	WindowHidden            WindowFlags = 0x0000000000000008
	WindowBorderless        WindowFlags = 0x0000000000000010
	WindowResizable         WindowFlags = 0x0000000000000020
	WindowMinimized         WindowFlags = 0x0000000000000040
	WindowMaximized         WindowFlags = 0x0000000000000080
	WindowMouseGrabbed      WindowFlags = 0x0000000000000100
	WindowInputFocus        WindowFlags = 0x0000000000000200
	WindowMouseFocus        WindowFlags = 0x0000000000000400
	WindowExternal          WindowFlags = 0x0000000000000800
	WindowModal             WindowFlags = 0x0000000000001000
	WindowHighPixelDensity  WindowFlags = 0x0000000000002000
	WindowMouseCapture      WindowFlags = 0x0000000000004000
	WindowMouseRelativeMode WindowFlags = 0x0000000000008000
	WindowAlwaysOnTop       WindowFlags = 0x0000000000010000
	WindowUtility           WindowFlags = 0x0000000000020000
	WindowTooltip           WindowFlags = 0x0000000000040000
	WindowPopupMenu         WindowFlags = 0x0000000000080000
	WindowKeyboardGrabbed   WindowFlags = 0x0000000000100000
	WindowVulkan            WindowFlags = 0x0000000010000000
	WindowMetal             WindowFlags = 0x0000000020000000
	WindowTransparent       WindowFlags = 0x0000000040000000
	WindowNotFocusable      WindowFlags = 0x0000000080000000
)

const (
	WindowPosCenteredMask = 0x2FFF0000
	WindowPosCentered     = WindowPosCenteredMask
)

func WindowPosCenteredDisplay(displayID DisplayID) uint32 {
	return WindowPosCenteredMask | uint32(displayID)
}

// DestroyWindow destroys a window.
func DestroyWindow(window *Window) {
	sdlDestroyWindow(window)
}

// func CreatePopupWindow(parent *Window, offset_x int32, offset_y int32, w int32, h int32, flags WindowFlags) *Window {
//	return sdlCreatePopupWindow(parent, offset_x, offset_y, w, h, flags)
// }

func CreateWindow(title string, w int32, h int32, flags WindowFlags) *Window {
	return sdlCreateWindow(title, w, h, flags)
}

// func CreateWindowWithProperties(props PropertiesID) *Window {
//	return sdlCreateWindowWithProperties(props)
// }

// func DestroyWindowSurface(window *Window) bool {
//	return sdlDestroyWindowSurface(window)
// }

// func DisableScreenSaver() bool {
//	return sdlDisableScreenSaver()
// }

// func EGL_GetCurrentConfig() EGLConfig {
//	return sdlEGL_GetCurrentConfig()
// }

// func EGL_GetCurrentDisplay() EGLDisplay {
//	return sdlEGL_GetCurrentDisplay()
// }

// func EGL_GetProcAddress(proc string) FunctionPointer {
//	return sdlEGL_GetProcAddress(proc)
// }

// func EGL_GetWindowSurface(window *Window) EGLSurface {
//	return sdlEGL_GetWindowSurface(window)
// }

// func EGL_SetAttributeCallbacks(platformAttribCallback EGLAttribArrayCallback, surfaceAttribCallback EGLIntArrayCallback, contextAttribCallback EGLIntArrayCallback, userdata unsafe.Pointer)  {
//	sdlEGL_SetAttributeCallbacks(platformAttribCallback, surfaceAttribCallback, contextAttribCallback, userdata)
// }

// func EnableScreenSaver() bool {
//	return sdlEnableScreenSaver()
// }

func FlashWindow(window *Window, operation FlashOperation) bool {
	return sdlFlashWindow(window, operation)
}

func GetClosestFullscreenDisplayMode(displayID DisplayID, w int32, h int32, refreshRate float32, includeHighDensityModes bool, closest *DisplayMode) bool {
	return sdlGetClosestFullscreenDisplayMode(displayID, w, h, refreshRate, includeHighDensityModes, closest)
}

func GetCurrentDisplayMode(displayID DisplayID) *DisplayMode {
	return sdlGetCurrentDisplayMode(displayID)
}

// func GetCurrentDisplayOrientation(displayID DisplayID) DisplayOrientation {
//	return sdlGetCurrentDisplayOrientation(displayID)
// }

func GetCurrentVideoDriver() string {
	return sdlGetCurrentVideoDriver()
}

// func GetDesktopDisplayMode(displayID DisplayID) *DisplayMode {
//	return sdlGetDesktopDisplayMode(displayID)
// }

// func GetDisplayBounds(displayID DisplayID, rect *Rect) bool {
//	return sdlGetDisplayBounds(displayID, rect)
// }

func GetDisplayContentScale(displayID DisplayID) float32 {
	return sdlGetDisplayContentScale(displayID)
}

// func GetDisplayForPoint(point *Point) DisplayID {
//	return sdlGetDisplayForPoint(point)
// }

// func GetDisplayForRect(rect *Rect) DisplayID {
//	return sdlGetDisplayForRect(rect)
// }

func GetDisplayForWindow(window *Window) DisplayID {
	return sdlGetDisplayForWindow(window)
}

func GetDisplayName(displayID DisplayID) string {
	return sdlGetDisplayName(displayID)
}

// func GetDisplayProperties(displayID DisplayID) PropertiesID {
//	return sdlGetDisplayProperties(displayID)
// }

// GetDisplays gets a list of currently connected displays.
func GetDisplays() []DisplayID {
	var count int32
	displays := sdlGetDisplays(&count)
	defer Free(unsafe.Pointer(displays))
	return mem.Copy(displays, count)
}

// func GetDisplayUsableBounds(displayID DisplayID, rect *Rect) bool {
//	return sdlGetDisplayUsableBounds(displayID, rect)
// }

// GetFullscreenDisplayModes gets a list of fullscreen display modes available on a display, or nil on error.
func GetFullscreenDisplayModes(displayID DisplayID) []*DisplayMode {
	var count int32
	displayModes := sdlGetFullscreenDisplayModes(displayID, &count)
	defer Free(unsafe.Pointer(displayModes))
	return mem.DeepCopy(displayModes, count)
}

// func GetGrabbedWindow() *Window {
//	return sdlGetGrabbedWindow()
// }

// func GetNaturalDisplayOrientation(displayID DisplayID) DisplayOrientation {
//	return sdlGetNaturalDisplayOrientation(displayID)
// }

// GetNumVideoDrivers gets the number of video drivers compiled into SDL.
func GetNumVideoDrivers() int32 {
	return sdlGetNumVideoDrivers()
}

func GetPrimaryDisplay() DisplayID {
	return sdlGetPrimaryDisplay()
}

// func GetSystemTheme() SystemTheme {
//	return sdlGetSystemTheme()
// }

// GetVideoDriver gets the name of a built in video driver.
func GetVideoDriver(index int32) string {
	return sdlGetVideoDriver(index)
}

// func GetWindowAspectRatio(window *Window, min_aspect *float32, max_aspect *float32) bool {
//	return sdlGetWindowAspectRatio(window, min_aspect, max_aspect)
// }

// func GetWindowBordersSize(window *Window, top *int32, left *int32, bottom *int32, right *int32) bool {
//	return sdlGetWindowBordersSize(window, top, left, bottom, right)
// }

func GetWindowDisplayScale(window *Window) float32 {
	return sdlGetWindowDisplayScale(window)
}

// func GetWindowFlags(window *Window) WindowFlags {
//	return sdlGetWindowFlags(window)
// }

// func GetWindowFromID(id WindowID) *Window {
//	return sdlGetWindowFromID(id)
// }

// GetWindowFullscreenMode queries the display mode to use when a window is visible at fullscreen.
func GetWindowFullscreenMode(window *Window) *DisplayMode {
	return sdlGetWindowFullscreenMode(window)
}

// func GetWindowICCProfile(window *Window, size *uint64) unsafe.Pointer {
//	return sdlGetWindowICCProfile(window, size)
// }

// GetWindowID returns the ID of the window on success or 0 on failure.
func GetWindowID(window *Window) WindowID {
	return sdlGetWindowID(window)
}

// GetWindowKeyboardGrab returns true if keyboard is grabbed, and false otherwise.
func GetWindowKeyboardGrab(window *Window) bool {
	return sdlGetWindowKeyboardGrab(window)
}

// func GetWindowMaximumSize(window *Window, w *int32, h *int32) bool {
//	return sdlGetWindowMaximumSize(window, w, h)
// }

// func GetWindowMinimumSize(window *Window, w *int32, h *int32) bool {
//	return sdlGetWindowMinimumSize(window, w, h)
// }

// GetWindowMouseGrab returns true if mouse is grabbed, and false otherwise.
func GetWindowMouseGrab(window *Window) bool {
	return sdlGetWindowMouseGrab(window)
}

// func GetWindowMouseRect(window *Window) *Rect {
//	return sdlGetWindowMouseRect(window)
// }

// GetWindowOpacity gets the opacity of a window.
func GetWindowOpacity(window *Window) float32 {
	return sdlGetWindowOpacity(window)
}

// func GetWindowParent(window *Window) *Window {
//	return sdlGetWindowParent(window)
// }

func GetWindowPixelDensity(window *Window) float32 {
	return sdlGetWindowPixelDensity(window)
}

// func GetWindowPixelFormat(window *Window) PixelFormat {
//	return sdlGetWindowPixelFormat(window)
// }

func GetWindowPosition(window *Window, x *int32, y *int32) bool {
	return sdlGetWindowPosition(window, x, y)
}

// func GetWindowProperties(window *Window) PropertiesID {
//	return sdlGetWindowProperties(window)
// }

// func GetWindows(count *int32) **Window {
//	return sdlGetWindows(count)
// }

// func GetWindowSafeArea(window *Window, rect *Rect) bool {
//	return sdlGetWindowSafeArea(window, rect)
// }

func GetWindowSize(window *Window, w *int32, h *int32) bool {
	return sdlGetWindowSize(window, w, h)
}

func GetWindowSizeInPixels(window *Window, w *int32, h *int32) bool {
	return sdlGetWindowSizeInPixels(window, w, h)
}

// GetWindowSurface gets the SDL surface associated with the window.
func GetWindowSurface(window *Window) *Surface {
	return sdlGetWindowSurface(window)
}

// func GetWindowSurfaceVSync(window *Window, vsync *int32) bool {
//	return sdlGetWindowSurfaceVSync(window, vsync)
// }

func GetWindowTitle(window *Window) string {
	return sdlGetWindowTitle(window)
}

// func GL_CreateContext(window *Window) GLContext {
//	return sdlGL_CreateContext(window)
// }

// func GL_DestroyContext(context GLContext) bool {
//	return sdlGL_DestroyContext(context)
// }

// func GL_ExtensionSupported(extension string) bool {
//	return sdlGL_ExtensionSupported(extension)
// }

// func GL_GetAttribute(attr GLAttr, value *int32) bool {
//	return sdlGL_GetAttribute(attr, value)
// }

// func GL_GetCurrentContext() GLContext {
//	return sdlGL_GetCurrentContext()
// }

// func GL_GetCurrentWindow() *Window {
//	return sdlGL_GetCurrentWindow()
// }

// func GL_GetProcAddress(proc string) FunctionPointer {
//	return sdlGL_GetProcAddress(proc)
// }

// func GL_GetSwapInterval(interval *int32) bool {
//	return sdlGL_GetSwapInterval(interval)
// }

// func GL_LoadLibrary(path string) bool {
//	return sdlGL_LoadLibrary(path)
// }

// func GL_MakeCurrent(window *Window, context GLContext) bool {
//	return sdlGL_MakeCurrent(window, context)
// }

// func GL_ResetAttributes()  {
//	sdlGL_ResetAttributes()
// }

// func GL_SetAttribute(attr GLAttr, value int32) bool {
//	return sdlGL_SetAttribute(attr, value)
// }

// func GL_SetSwapInterval(interval int32) bool {
//	return sdlGL_SetSwapInterval(interval)
// }

// func GL_SwapWindow(window *Window) bool {
//	return sdlGL_SwapWindow(window)
// }

// func GL_UnloadLibrary()  {
//	sdlGL_UnloadLibrary()
// }

// HideWindow hides a window.
func HideWindow(window *Window) bool {
	return sdlHideWindow(window)
}

// func MaximizeWindow(window *Window) bool {
//	return sdlMaximizeWindow(window)
// }

// func MinimizeWindow(window *Window) bool {
//	return sdlMinimizeWindow(window)
// }

// RaiseWindow requests that a window be raised above other windows and gain the input focus.
func RaiseWindow(window *Window) bool {
	return sdlRaiseWindow(window)
}

func RestoreWindow(window *Window) bool {
	return sdlRestoreWindow(window)
}

// func ScreenSaverEnabled() bool {
//	return sdlScreenSaverEnabled()
// }

// SetWindowAlwaysOnTop sets the window to always be above the others.
func SetWindowAlwaysOnTop(window *Window, onTop bool) bool {
	return sdlSetWindowAlwaysOnTop(window, onTop)
}

// func SetWindowAspectRatio(window *Window, min_aspect float32, max_aspect float32) bool {
//	return sdlSetWindowAspectRatio(window, min_aspect, max_aspect)
// }

// SetWindowBordered sets the border state of a window.
func SetWindowBordered(window *Window, bordered bool) bool {
	return sdlSetWindowBordered(window, bordered)
}

// SetWindowFocusable sets whether the window may have input focus.
func SetWindowFocusable(window *Window, focusable bool) bool {
	return sdlSetWindowFocusable(window, focusable)
}

// SetWindowFullscreen requests that the window's fullscreen state be changed.
func SetWindowFullscreen(window *Window, fullscreen bool) bool {
	return sdlSetWindowFullscreen(window, fullscreen)
}

// SetWindowFullscreenMode sets the display mode to use when a window is visible and fullscreen.
func SetWindowFullscreenMode(window *Window, mode *DisplayMode) bool {
	return sdlSetWindowFullscreenMode(window, mode)
}

func SetWindowHitTest(window *Window, callback HitTest, callbackData unsafe.Pointer) bool {
	wrapper := func(win *Window, point *Point, data unsafe.Pointer) uintptr {
		return uintptr(callback(win, point, data))
	}

	return sdlSetWindowHitTest(window, wrapper, callbackData)
}

func SetWindowIcon(window *Window, icon *Surface) bool {
	return sdlSetWindowIcon(window, icon)
}

// SetWindowKeyboardGrab enables capture of system keyboard shortcuts like Alt+Tab or the Meta/Super key.
// Note that not all system keyboard shortcuts can be captured by applications (one example is Ctrl+Alt+Del on Windows).
func SetWindowKeyboardGrab(window *Window, grabbed bool) bool {
	return sdlSetWindowKeyboardGrab(window, grabbed)
}

// func SetWindowMaximumSize(window *Window, max_w int32, max_h int32) bool {
//	return sdlSetWindowMaximumSize(window, max_w, max_h)
// }

// func SetWindowMinimumSize(window *Window, min_w int32, min_h int32) bool {
//	return sdlSetWindowMinimumSize(window, min_w, min_h)
// }

// func SetWindowModal(window *Window, modal bool) bool {
//	return sdlSetWindowModal(window, modal)
// }

// SetWindowMouseGrab enables restriction of the mouse cursor to the window.
func SetWindowMouseGrab(window *Window, grabbed bool) bool {
	return sdlSetWindowMouseGrab(window, grabbed)
}

// func SetWindowMouseRect(window *Window, rect *Rect) bool {
//	return sdlSetWindowMouseRect(window, rect)
// }

// SetWindowOpacity sets the opacity for a window.
func SetWindowOpacity(window *Window, opacity float32) bool {
	return sdlSetWindowOpacity(window, opacity)
}

// func SetWindowParent(window *Window, parent *Window) bool {
//	return sdlSetWindowParent(window, parent)
// }

func SetWindowPosition(window *Window, x int32, y int32) bool {
	return sdlSetWindowPosition(window, x, y)
}

// SetWindowResizable sets the user-resizable state of a window.
func SetWindowResizable(window *Window, resizable bool) bool {
	return sdlSetWindowResizable(window, resizable)
}

// func SetWindowShape(window *Window, shape *Surface) bool {
//	return sdlSetWindowShape(window, shape)
// }

func SetWindowSize(window *Window, w int32, h int32) bool {
	return sdlSetWindowSize(window, w, h)
}

// func SetWindowSurfaceVSync(window *Window, vsync int32) bool {
//	return sdlSetWindowSurfaceVSync(window, vsync)
// }

func SetWindowTitle(window *Window, title string) bool {
	return sdlSetWindowTitle(window, title)
}

// ShowWindow shows a window.
func ShowWindow(window *Window) bool {
	return sdlShowWindow(window)
}

// func ShowWindowSystemMenu(window *Window, x int32, y int32) bool {
//	return sdlShowWindowSystemMenu(window, x, y)
// }

func SyncWindow(window *Window) bool {
	return sdlSyncWindow(window)
}

// UpdateWindowSurface copies the window surface to the screen.
func UpdateWindowSurface(window *Window) bool {
	return sdlUpdateWindowSurface(window)
}

// func UpdateWindowSurfaceRects(window *Window, rects *Rect, numrects int32) bool {
//	return sdlUpdateWindowSurfaceRects(window, rects, numrects)
// }

// func WindowHasSurface(window *Window) bool {
//	return sdlWindowHasSurface(window)
// }
