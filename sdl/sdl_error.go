package sdl

import "fmt"

// [GetError] retrieves a message about the last error that occurred on the current thread.
//
// [GetError]: https://wiki.libsdl.org/SDL3/SDL_GetError
func GetError() string {
	return sdlGetError()
}

// [ClearError] clears any previous error message for this thread.
//
// [ClearError]: https://wiki.libsdl.org/SDL3/SDL_ClearError
func ClearError() bool {
	return sdlClearError()
}

// func OutOfMemory() bool {
//	return sdlOutOfMemory()
// }

// [SetError] sets the SDL error message for the current thread.
//
// [SetError]: https://wiki.libsdl.org/SDL3/SDL_SetError
func SetError(format string, a ...any) bool {
	return sdlSetError(fmt.Sprintf(format, a...))
}

// func SetErrorV(fmt string, ap va_list) bool {
//	return sdlSetErrorV(fmt, ap)
// }

// [InvalidParamError] sets the SDL error message, standardizes error reporting on unsupported operations.
//
// [InvalidParamError]: https://wiki.libsdl.org/SDL3/SDL_InvalidParamError
func InvalidParamError(param string) bool {
	return SetError("Parameter '%s' is invalid", param)
}
