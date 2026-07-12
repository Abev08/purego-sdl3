package sdl

import "github.com/jupiterrider/purego-sdl3/internal/convert"

// [AssertState] defines the possible outcomes from a triggered assertion.
//
// [AssertState]: https://wiki.libsdl.org/SDL3/SDL_AssertState
type AssertState uint32

const (
	AssertionRetry        AssertState = iota // Retry the assert immediately.
	AssertionBreak                           // Make the debugger trigger a breakpoint.
	AssertionAbort                           // Terminate the program.
	AssertionIgnore                          // Ignore the assert.
	AssertionAlwaysIgnore                    // Ignore the assert from now on.
)

// [AssertData] defines information about an assertion failure.
//
// [AssertData]: https://wiki.libsdl.org/SDL3/SDL_AssertData
type AssertData struct {
	AlwaysIgnore bool        // True if app should always continue when assertion is triggered.
	TriggerCount uint32      // Number of times this assertion has been triggered.
	condition    *byte       // A string of this assert's test code.
	filename     *byte       // The source file where this assert lives.
	Linenum      int32       // The line in `filename` where this assert lives.
	function     *byte       // The name of the function where this assert lives.
	Next         *AssertData // Next item in the linked list.
}

// Condition returns a string of this assert's test code.
func (ad *AssertData) Condition() string {
	return convert.ToString(ad.condition)
}

// Filename returns the source file where this assert lives.
func (ad *AssertData) Filename() string {
	return convert.ToString(ad.filename)
}

// Function returns the name of the function where this assert lives.
func (ad *AssertData) Function() string {
	return convert.ToString(ad.function)
}

// func ReportAssertion(data *AssertData, func string, file string, line int32) AssertState {
//	return sdlReportAssertion(data, func, file, line)
// }

// [AssertionHandler] is a callback that fires when an SDL assertion fails.
//
// [AssertionHandler]: https://wiki.libsdl.org/SDL3/SDL_AssertionHandler
type AssertionHandler AssertState

// func SetAssertionHandler(handler AssertionHandler, userdata unsafe.Pointer)  {
//	sdlSetAssertionHandler(handler, userdata)
// }

// func GetDefaultAssertionHandler() AssertionHandler {
//	return sdlGetDefaultAssertionHandler()
// }

// func GetAssertionHandler(puserdata *unsafe.Pointer) AssertionHandler {
//	return sdlGetAssertionHandler(puserdata)
// }

// func GetAssertionReport() *AssertData {
//	return sdlGetAssertionReport()
// }

// func ResetAssertionReport()  {
//	sdlResetAssertionReport()
// }
