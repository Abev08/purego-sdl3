package sdl

import (
	"fmt"
	"unsafe"
)

// [LogCategory] specifies the predefined log categories.
//
// [LogCategory]: https://wiki.libsdl.org/SDL3/SDL_LogCategory
type LogCategory uint32

const (
	LogCategoryApplication LogCategory = iota
	LogCategoryError
	LogCategoryAssert
	LogCategorySystem
	LogCategoryAudio
	LogCategoryVideo
	LogCategoryRender
	LogCategoryInput
	LogCategoryTest
	LogCategoryGPU
	LogCategoryReserved2
	LogCategoryReserved3
	LogCategoryReserved4
	LogCategoryReserved5
	LogCategoryReserved6
	LogCategoryReserved7
	LogCategoryReserved8
	LogCategoryReserved9
	LogCategoryReserved10
	LogCategoryCustom
)

// [LogPriority] specifies the predefined log priorities.
//
// [LogPriority]: https://wiki.libsdl.org/SDL3/SDL_LogPriority
type LogPriority uint32

const (
	LogPriorityInvalid LogPriority = iota
	LogPriorityTrace
	LogPriorityVerbose
	LogPriorityDebug
	LogPriorityInfo
	LogPriorityWarn
	LogPriorityError
	LogPriorityCritical
	LogPriorityCount
)

// [SetLogPriorities] sets the priority of all log categories.
//
// [SetLogPriorities]: https://wiki.libsdl.org/SDL3/SDL_SetLogPriorities
func SetLogPriorities(priority LogPriority) {
	sdlSetLogPriorities(priority)
}

// [SetLogPriority] sets the priority of a particular log category.
//
// [SetLogPriority]: https://wiki.libsdl.org/SDL3/SDL_SetLogPriority
func SetLogPriority(category int32, priority LogPriority) {
	sdlSetLogPriority(category, priority)
}

// [GetLogPriority] gets the priority of a particular log category.
//
// [GetLogPriority]: https://wiki.libsdl.org/SDL3/SDL_GetLogPriority
// func GetLogPriority(category int32) LogPriority {
//	return sdlGetLogPriority(category)
// }

// [ResetLogPriorities] resets all priorities to default.
//
// [ResetLogPriorities]: https://wiki.libsdl.org/SDL3/SDL_ResetLogPriorities
func ResetLogPriorities() {
	sdlResetLogPriorities()
}

// [SetLogPriorityPrefix] sets the text prepended to log messages of a given priority.
//
// [SetLogPriorityPrefix]: https://wiki.libsdl.org/SDL3/SDL_SetLogPriorityPrefix
// func SetLogPriorityPrefix(priority LogPriority, prefix string) bool {
//	return sdlSetLogPriorityPrefix(priority, prefix)
// }

// [Log] logs a message with [LOG_CATEGORY_APPLICATION] and [LOG_PRIORITY_INFO].
//
// [Log]: https://wiki.libsdl.org/SDL3/SDL_Log
func Log(format string, a ...any) {
	sdlLog(fmt.Sprintf(format, a...))
}

// [LogTrace] logs a message with [LogPriorityTrace].
//
// [LogTrace]: https://wiki.libsdl.org/SDL3/SDL_LogTrace
// func LogTrace(category int32, fmt string) {
//	sdlLogTrace(category, fmt)
// }

// [LogVerbose] logs a message with [LogPriorityVerbose].
//
// [LogVerbose]: https://wiki.libsdl.org/SDL3/SDL_LogVerbose
// func LogVerbose(category int32, fmt string) {
//	sdlLogVerbose(category, fmt)
// }

// [LogDebug] logs a message with [LogPriorityDebug].
//
// [LogDebug]: https://wiki.libsdl.org/SDL3/SDL_LogDebug
// func LogDebug(category int32, fmt string) {
//	sdlLogDebug(category, fmt)
// }

// [LogInfo] logs a message with [LogPriorityInfo].
//
// [LogInfo]: https://wiki.libsdl.org/SDL3/SDL_LogInfo
// func LogInfo(category int32, fmt string) {
//	sdlLogInfo(category, fmt)
// }

// [LogWarn] logs a message with [LogPriorityWarn].
//
// [LogWarn]: https://wiki.libsdl.org/SDL3/SDL_LogWarn
// func LogWarn(category int32, fmt string) {
//	sdlLogWarn(category, fmt)
// }

// [LogError] logs a message with [LogPriorityError].
//
// [LogError]: https://wiki.libsdl.org/SDL3/SDL_LogError
func LogError(category LogCategory, format string, a ...any) {
	sdlLogError(category, fmt.Sprintf(format, a...))
}

// [LogCritical] logs a message with [LogPriorityCritical].
//
// [LogCritical]: https://wiki.libsdl.org/SDL3/SDL_LogCritical
// func LogCritical(category int32, fmt string) {
//	sdlLogCritical(category, fmt)
// }

// [LogMessage] logs a message with the specified category and priority.
//
// [LogMessage]: https://wiki.libsdl.org/SDL3/SDL_LogMessage
func LogMessage(category LogCategory, priority LogPriority, format string, a ...any) {
	sdlLogMessage(category, priority, fmt.Sprintf(format, a...))
}

// [LogMessageV] logs a message with the specified category and priority.
//
// [LogMessageV]: https://wiki.libsdl.org/SDL3/SDL_LogMessageV
// func LogMessageV(category int32, priority LogPriority, fmt string, ap va_list) {
//	sdlLogMessageV(category, priority, fmt, ap)
// }

// [LogOutputFunction] defines the prototype for the log output callback function.
//
// [LogOutputFunction]: https://wiki.libsdl.org/SDL3/SDL_LogOutputFunction
type LogOutputFunction uintptr

// [GetDefaultLogOutputFunction] gets the default log output function.
//
// [GetDefaultLogOutputFunction]: https://wiki.libsdl.org/SDL3/SDL_GetDefaultLogOutputFunction
// func GetDefaultLogOutputFunction() LogOutputFunction {
//	return sdlGetDefaultLogOutputFunction()
// }

// [GetLogOutputFunction] gets the current log output function.
//
// [GetLogOutputFunction]: https://wiki.libsdl.org/SDL3/SDL_GetLogOutputFunction
// func GetLogOutputFunction(callback *LogOutputFunction, userdata *unsafe.Pointer) {
//	sdlGetLogOutputFunction(callback, userdata)
// }

// [SetLogOutputFunction] replaces the default log output function with one of your own.
//
// [SetLogOutputFunction]: https://wiki.libsdl.org/SDL3/SDL_SetLogOutputFunction
func SetLogOutputFunction(callback LogOutputFunction, userdata unsafe.Pointer) {
	sdlSetLogOutputFunction(callback, userdata)
}
