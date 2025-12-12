package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

// [MouseWheelDirection] defines the scroll direction types for the Scroll event.
//
// [MouseWheelDirection]: https://wiki.libsdl.org/SDL3/SDL_MouseWheelDirection
type MouseWheelDirection uint32

const (
	MouseWheelNormal  MouseWheelDirection = iota // The scroll direction is normal.
	MouseWheelFlipped                            // The scroll direction is flipped / natural.
)

// [SystemCursor] defines the cursor types for [CreateSystemCursor].
//
// [SystemCursor]: https://wiki.libsdl.org/SDL3/SDL_SystemCursor
type SystemCursor uint32

const (
	SystemCursorDefault    SystemCursor = iota // Default cursor. Usually an arrow.
	SystemCursorText                           // Text selection. Usually an I-beam.
	SystemCursorWait                           // Wait. Usually an hourglass or watch or spinning ball.
	SystemCursorCrosshair                      // Crosshair.
	SystemCursorProgress                       // Program is busy but still interactive. Usually it's WAIT with an arrow.
	SystemCursorNWSEResize                     // Double arrow pointing northwest and southeast.
	SystemCursorNESWResize                     // Double arrow pointing northeast and southwest.
	SystemCursorEWResize                       // Double arrow pointing west and east.
	SystemCursorNSResize                       // Double arrow pointing north and south.
	SystemCursorMove                           // Four pointed arrow pointing north, south, east, and west.
	SystemCursorNotAllowed                     // Not permitted. Usually a slashed circle or crossbones.
	SystemCursorPointer                        // Pointer that indicates a link. Usually a pointing hand.
	SystemCursorNWResize                       // Window resize top-left. This may be a single arrow or a double arrow like NWSE_RESIZE.
	SystemCursorNResize                        // Window resize top. May be NS_RESIZE.
	SystemCursorNEResize                       // Window resize top-right. May be NESW_RESIZE.
	SystemCursorEResize                        // Window resize right. May be EW_RESIZE.
	SystemCursorSEResize                       // Window resize bottom-right. May be NWSE_RESIZE.
	SystemCursorSResize                        // Window resize bottom. May be NS_RESIZE.
	SystemCursorSWResize                       // Window resize bottom-left. May be NESW_RESIZE.
	SystemCursorWResize                        // Window resize left. May be EW_RESIZE.
	SystemCursorCount
)

// [MouseID] is a unique ID for a mouse for the time it is connected to the system, and is never reused for the lifetime of the application.
//
// [MouseID]: https://wiki.libsdl.org/SDL3/SDL_MouseID
type MouseID uint32

// [MouseButtonFlags] is a bitmask of pressed mouse buttons, as reported by [GetMouseState], etc.
//
// [MouseButtonFlags]: https://wiki.libsdl.org/SDL3/SDL_MouseButtonFlags
type MouseButtonFlags uint32

const (
	ButtonLeft MouseButtonFlags = iota + 1
	ButtonMiddle
	ButtonRight
	ButtonX1
	ButtonX2

	ButtonLMask  MouseButtonFlags = 1 << (ButtonLeft - 1)
	ButtonMMask  MouseButtonFlags = 1 << (ButtonMiddle - 1)
	ButtonRMask  MouseButtonFlags = 1 << (ButtonRight - 1)
	ButtonX1Mask MouseButtonFlags = 1 << (ButtonX1 - 1)
	ButtonX2Mask MouseButtonFlags = 1 << (ButtonX2 - 1)
)

// [Cursor] is a structure specifying the structure used to identify an SDL cursors.
//
// [Cursor]: https://wiki.libsdl.org/SDL3/SDL_Cursor
type Cursor struct{}

// [CaptureMouse] captures the mouse and to track input outside an SDL window.
//
// [CaptureMouse]: https://wiki.libsdl.org/SDL3/SDL_CaptureMouse
func CaptureMouse(enabled bool) bool {
	return sdlCaptureMouse(enabled)
}

// [CreateColorCursor] creates a color cursor.
//
// [CreateColorCursor]: https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor
func CreateColorCursor(surface *Surface, hotX int32, hotY int32) *Cursor {
	return sdlCreateColorCursor(surface, hotX, hotY)
}

// [CreateCursor] creates a cursor using the specified bitmap data and mask (in MSB format).
//
// [CreateCursor]: https://wiki.libsdl.org/SDL3/SDL_CreateCursor
func CreateCursor(data *uint8, mask *uint8, w int32, h int32, hotX int32, hotY int32) *Cursor {
	return sdlCreateCursor(data, mask, w, h, hotX, hotY)
}

// [CreateSystemCursor] creates a system cursor.
//
// [CreateSystemCursor]: https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor
func CreateSystemCursor(id SystemCursor) *Cursor {
	return sdlCreateSystemCursor(id)
}

// [CursorVisible] returns true if the cursor is being shown, or false if the cursor is hidden.
//
// [CursorVisible]: https://wiki.libsdl.org/SDL3/SDL_CursorVisible
func CursorVisible() bool {
	return sdlCursorVisible()
}

// [DestroyCursor] frees a previously-created cursor.
//
// [DestroyCursor]: https://wiki.libsdl.org/SDL3/SDL_DestroyCursor
func DestroyCursor(cursor *Cursor) {
	sdlDestroyCursor(cursor)
}

// [GetCursor] gets the active cursor.
//
// [GetCursor]: https://wiki.libsdl.org/SDL3/SDL_GetCursor
func GetCursor() *Cursor {
	return sdlGetCursor()
}

// [GetDefaultCursor] gets the default cursor.
//
// [GetDefaultCursor]: https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor
func GetDefaultCursor() *Cursor {
	return sdlGetDefaultCursor()
}

// [GetGlobalMouseState] queries the platform for the asynchronous mouse button state and the desktop-relative platform-cursor position.
//
// [GetGlobalMouseState]: https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState
func GetGlobalMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetGlobalMouseState(x, y)
}

// [GetMice] returns a list of currently connected mice or nil on failure.
//
// [GetMice]: https://wiki.libsdl.org/SDL3/SDL_GetMice
func GetMice() []MouseID {
	var count int32
	mice := sdlGetMice(&count)
	defer Free(unsafe.Pointer(mice))
	return mem.Copy(mice, count)
}

// [GetMouseFocus] gets the window which currently has mouse focus.
//
// [GetMouseFocus]: https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus
func GetMouseFocus() *Window {
	return sdlGetMouseFocus()
}

// [GetMouseNameForID] returns the name of the selected mouse.
//
// On failure or if the mouse doesn't have a name, this function returns "".
//
// [GetMouseNameForID]: https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID
func GetMouseNameForID(instanceId MouseID) string {
	return sdlGetMouseNameForID(instanceId)
}

// [GetMouseState] queries SDL's cache for the synchronous mouse button state and the window-relative SDL-cursor position.
//
// [GetMouseState]: https://wiki.libsdl.org/SDL3/SDL_GetMouseState
func GetMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetMouseState(x, y)
}

// [GetRelativeMouseState] queries SDL's cache for the synchronous mouse button state and accumulated mouse delta since last call.
//
// [GetRelativeMouseState]: https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState
func GetRelativeMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetRelativeMouseState(x, y)
}

// [GetWindowRelativeMouseMode] queries whether relative mouse mode is enabled for a window.
//
// [GetWindowRelativeMouseMode]: https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode
func GetWindowRelativeMouseMode(window *Window) bool {
	return sdlGetWindowRelativeMouseMode(window)
}

// [HasMouse] returns whether a mouse is currently connected.
//
// [HasMouse]: https://wiki.libsdl.org/SDL3/SDL_HasMouse
func HasMouse() bool {
	return sdlHasMouse()
}

// [HideCursor] hides the cursor.
//
// [HideCursor]: https://wiki.libsdl.org/SDL3/SDL_HideCursor
func HideCursor() bool {
	return sdlHideCursor()
}

// [SetCursor] sets the active cursor.
//
// [SetCursor]: https://wiki.libsdl.org/SDL3/SDL_SetCursor
func SetCursor(cursor *Cursor) bool {
	return sdlSetCursor(cursor)
}

// [SetWindowRelativeMouseMode] sets relative mouse mode for a window.
//
// [SetWindowRelativeMouseMode]: https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode
func SetWindowRelativeMouseMode(window *Window, enabled bool) bool {
	return sdlSetWindowRelativeMouseMode(window, enabled)
}

// [ShowCursor] shows the cursor.
//
// [ShowCursor]: https://wiki.libsdl.org/SDL3/SDL_ShowCursor
func ShowCursor() bool {
	return sdlShowCursor()
}

// [WarpMouseGlobal] moves the mouse to the given position in global screen space.
//
// [WarpMouseGlobal]: https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal
func WarpMouseGlobal(x float32, y float32) bool {
	return sdlWarpMouseGlobal(x, y)
}

// [WarpMouseInWindow] moves the mouse cursor to the given position within the window.
//
// [WarpMouseInWindow]: https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow
func WarpMouseInWindow(window *Window, x float32, y float32) {
	sdlWarpMouseInWindow(window, x, y)
}
