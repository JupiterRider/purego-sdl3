package sdl

type MouseID uint32

type MouseButtonFlags uint32

const (
	ButtonLeft   MouseButtonFlags = 1
	ButtonMiddle MouseButtonFlags = 2
	ButtonRight  MouseButtonFlags = 3
	ButtonX1     MouseButtonFlags = 4
	ButtonX2     MouseButtonFlags = 5

	ButtonLmask  = 1 << (ButtonLeft - 1)
	ButtonMmask  = 1 << (ButtonMiddle - 1)
	ButtonRmask  = 1 << (ButtonRight - 1)
	ButtonX1mask = 1 << (ButtonX1 - 1)
	ButtonX2mask = 1 << (ButtonX2 - 1)
)

// func CaptureMouse(enabled bool) bool {
//	return sdlCaptureMouse(enabled)
// }

// func CreateColorCursor(surface *Surface, hot_x int32, hot_y int32) *Cursor {
//	return sdlCreateColorCursor(surface, hot_x, hot_y)
// }

// func CreateCursor(data *uint8, mask *uint8, w int32, h int32, hot_x int32, hot_y int32) *Cursor {
//	return sdlCreateCursor(data, mask, w, h, hot_x, hot_y)
// }

// func CreateSystemCursor(id SystemCursor) *Cursor {
//	return sdlCreateSystemCursor(id)
// }

// func CursorVisible() bool {
//	return sdlCursorVisible()
// }

// func DestroyCursor(cursor *Cursor)  {
//	sdlDestroyCursor(cursor)
// }

// func GetCursor() *Cursor {
//	return sdlGetCursor()
// }

// func GetDefaultCursor() *Cursor {
//	return sdlGetDefaultCursor()
// }

// func GetGlobalMouseState(x *float32, y *float32) MouseButtonFlags {
//	return sdlGetGlobalMouseState(x, y)
// }

// func GetMice(count *int32) *MouseID {
//	return sdlGetMice(count)
// }

// func GetMouseFocus() *Window {
//	return sdlGetMouseFocus()
// }

// func GetMouseNameForID(instance_id MouseID) string {
//	return sdlGetMouseNameForID(instance_id)
// }

// GetMouseState queries SDL's cache for the synchronous mouse button state and the window-relative SDL-cursor position.
func GetMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetMouseState(x, y)
}

// func GetRelativeMouseState(x *float32, y *float32) MouseButtonFlags {
//	return sdlGetRelativeMouseState(x, y)
// }

// func GetWindowRelativeMouseMode(window *Window) bool {
//	return sdlGetWindowRelativeMouseMode(window)
// }

// func HasMouse() bool {
//	return sdlHasMouse()
// }

// HideCursor hides the cursor.
func HideCursor() bool {
	return sdlHideCursor()
}

// func SetCursor(cursor *Cursor) bool {
//	return sdlSetCursor(cursor)
// }

// func SetWindowRelativeMouseMode(window *Window, enabled bool) bool {
//	return sdlSetWindowRelativeMouseMode(window, enabled)
// }

// ShowCursor shows the cursor.
func ShowCursor() bool {
	return sdlShowCursor()
}

// func WarpMouseGlobal(x float32, y float32) bool {
//	return sdlWarpMouseGlobal(x, y)
// }

// WarpMouseInWindow moves the mouse cursor to the given position within the window.
func WarpMouseInWindow(window *Window, x float32, y float32) {
	sdlWarpMouseInWindow(window, x, y)
}

// WarpMouseInWindow moves the mouse cursor to the given position within the window.
func (window *Window) WarpMouseInWindow(x float32, y float32) {
	WarpMouseInWindow(window, x, y)
}
