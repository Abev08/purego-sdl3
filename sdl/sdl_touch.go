package sdl

import "math"

// [TouchID] specifies a unique ID for a touch device.
//
// [TouchID]: https://wiki.libsdl.org/SDL3/SDL_TouchID
type TouchID uint64

// [FingerID] specifies a unique ID for a single finger on a touch device.
//
// [FingerID]: https://wiki.libsdl.org/SDL3/SDL_FingerID
type FingerID uint64

// [TouchDeviceType] specifies an enum that describes the type of a touch device.
//
// [TouchDeviceType]: https://wiki.libsdl.org/SDL3/SDL_TouchDeviceType
type TouchDeviceType int32

const (
	TouchDeviceInvalid          TouchDeviceType = iota - 1
	TouchDeviceDirect                           // Touch screen with window-relative coordinates.
	TouchDeviceIndirectAbsolute                 // Trackpad with absolute device coordinates.
	TouchDeviceIndirectRelative                 // Trackpad with screen cursor-relative coordinates.
)

// [Finger] specifies data about a single finger in a multitouch event.
//
// [Finger]: https://wiki.libsdl.org/SDL3/SDL_Finger
type Finger struct {
	ID       FingerID // The finger ID.
	X        float32  // The x-axis location of the touch event, normalized (0...1).
	Y        float32  // The y-axis location of the touch event, normalized (0...1).
	Pressure float32  // The quantity of pressure applied, normalized (0...1).
}

const (
	TouchMouseID = MouseID(math.MaxUint32) // The [MouseID] for mouse events simulated with touch input.
	MouseTouchID = TouchID(math.MaxUint64) // The [TouchID] for touch events simulated with mouse input.
)

// [GetTouchDevices] gets a list of registered touch devices.
//
// [GetTouchDevices]: https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices
// func GetTouchDevices(count *int32) *TouchID {
//	return sdlGetTouchDevices(count)
// }

// [GetTouchDeviceName] gets the touch device name as reported from the driver.
//
// [GetTouchDeviceName]: https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName
// func GetTouchDeviceName(touchID TouchID) string {
//	return sdlGetTouchDeviceName(touchID)
// }

// [GetTouchDeviceType] gets the type of the given touch device.
//
// [GetTouchDeviceType]: https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType
// func GetTouchDeviceType(touchID TouchID) TouchDeviceType {
//	return sdlGetTouchDeviceType(touchID)
// }

// [GetTouchFingers] gets a list of active fingers for a given touch device.
//
// [GetTouchFingers]: https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers
// func GetTouchFingers(touchID TouchID, count *int32) **Finger {
//	return sdlGetTouchFingers(touchID, count)
// }
