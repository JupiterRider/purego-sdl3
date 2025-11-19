package sdl

import "fmt"

// [LogPriority] is a structure specifying the predefined log priorities.
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

// [LogCategory] is a structure specifying the predefined log categories.
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

// func GetDefaultLogOutputFunction() LogOutputFunction {
//	return sdlGetDefaultLogOutputFunction()
// }

// func GetLogOutputFunction(callback *LogOutputFunction, userdata *unsafe.Pointer)  {
//	sdlGetLogOutputFunction(callback, userdata)
// }

// func GetLogPriority(category int32) LogPriority {
//	return sdlGetLogPriority(category)
// }

// [Log] logs a message with [LOG_CATEGORY_APPLICATION] and [LOG_PRIORITY_INFO].
//
// [Log]: https://wiki.libsdl.org/SDL3/SDL_Log
func Log(format string, a ...any) {
	sdlLog(fmt.Sprintf(format, a...))
}

// func LogCritical(category int32, fmt string)  {
//	sdlLogCritical(category, fmt)
// }

// func LogDebug(category int32, fmt string)  {
//	sdlLogDebug(category, fmt)
// }

// [LogError] logs a message with [LOG_PRIORITY_ERROR].
//
// [LogError]: https://wiki.libsdl.org/SDL3/SDL_LogError
func LogError(category LogCategory, format string, a ...any) {
	sdlLogError(category, fmt.Sprintf(format, a...))
}

// func LogInfo(category int32, fmt string)  {
//	sdlLogInfo(category, fmt)
// }

// [LogMessage] logs a message with the specified category and priority.
//
// [LogMessage]: https://wiki.libsdl.org/SDL3/SDL_LogMessage
func LogMessage(category LogCategory, priority LogPriority, format string, a ...any) {
	sdlLogMessage(category, priority, fmt.Sprintf(format, a...))
}

// func LogMessageV(category int32, priority LogPriority, fmt string, ap va_list)  {
//	sdlLogMessageV(category, priority, fmt, ap)
// }

// func LogTrace(category int32, fmt string)  {
//	sdlLogTrace(category, fmt)
// }

// func LogVerbose(category int32, fmt string)  {
//	sdlLogVerbose(category, fmt)
// }

// func LogWarn(category int32, fmt string)  {
//	sdlLogWarn(category, fmt)
// }

// [ResetLogPriorities] resets all priorities to default.
//
// [ResetLogPriorities]: https://wiki.libsdl.org/SDL3/SDL_ResetLogPriorities
func ResetLogPriorities() {
	sdlResetLogPriorities()
}

// func SetLogOutputFunction(callback LogOutputFunction, userdata unsafe.Pointer)  {
//	sdlSetLogOutputFunction(callback, userdata)
// }

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

// func SetLogPriorityPrefix(priority LogPriority, prefix string) bool {
//	return sdlSetLogPriorityPrefix(priority, prefix)
// }
