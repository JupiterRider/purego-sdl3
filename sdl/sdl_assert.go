package sdl

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

// func GetAssertionHandler(puserdata *unsafe.Pointer) AssertionHandler {
//	return sdlGetAssertionHandler(puserdata)
// }

// func GetAssertionReport() *AssertData {
//	return sdlGetAssertionReport()
// }

// func GetDefaultAssertionHandler() AssertionHandler {
//	return sdlGetDefaultAssertionHandler()
// }

// func ReportAssertion(data *AssertData, func string, file string, line int32) AssertState {
//	return sdlReportAssertion(data, func, file, line)
// }

// func ResetAssertionReport()  {
//	sdlResetAssertionReport()
// }

// func SetAssertionHandler(handler AssertionHandler, userdata unsafe.Pointer)  {
//	sdlSetAssertionHandler(handler, userdata)
// }
