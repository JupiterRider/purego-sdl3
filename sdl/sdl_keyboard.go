package sdl

type KeyboardID uint32

// func ClearComposition(window *Window) bool {
//	return sdlClearComposition(window)
// }

// func GetKeyboardFocus() *Window {
//	return sdlGetKeyboardFocus()
// }

// func GetKeyboardNameForID(instance_id KeyboardID) string {
//	return sdlGetKeyboardNameForID(instance_id)
// }

// func GetKeyboards(count *int32) *KeyboardID {
//	return sdlGetKeyboards(count)
// }

// func GetKeyboardState(numkeys *int32) *bool {
//	return sdlGetKeyboardState(numkeys)
// }

// func GetKeyFromName(name string) Keycode {
//	return sdlGetKeyFromName(name)
// }

// func GetKeyFromScancode(scancode Scancode, modstate Keymod, key_event bool) Keycode {
//	return sdlGetKeyFromScancode(scancode, modstate, key_event)
// }

// func GetKeyName(key Keycode) string {
//	return sdlGetKeyName(key)
// }

// GetModState returns an OR'd combination of the modifier keys for the keyboard.
func GetModState() Keymod {
	return sdlGetModState()
}

// func GetScancodeFromKey(key Keycode, modstate *Keymod) Scancode {
//	return sdlGetScancodeFromKey(key, modstate)
// }

// func GetScancodeFromName(name string) Scancode {
//	return sdlGetScancodeFromName(name)
// }

// func GetScancodeName(scancode Scancode) string {
//	return sdlGetScancodeName(scancode)
// }

// func GetTextInputArea(window *Window, rect *Rect, cursor *int32) bool {
//	return sdlGetTextInputArea(window, rect, cursor)
// }

// func HasKeyboard() bool {
//	return sdlHasKeyboard()
// }

// func HasScreenKeyboardSupport() bool {
//	return sdlHasScreenKeyboardSupport()
// }

// func ResetKeyboard()  {
//	sdlResetKeyboard()
// }

// func ScreenKeyboardShown(window *Window) bool {
//	return sdlScreenKeyboardShown(window)
// }

// func SetModState(modstate Keymod)  {
//	sdlSetModState(modstate)
// }

// func SetScancodeName(scancode Scancode, name string) bool {
//	return sdlSetScancodeName(scancode, name)
// }

// func SetTextInputArea(window *Window, rect *Rect, cursor int32) bool {
//	return sdlSetTextInputArea(window, rect, cursor)
// }

// StartTextInput starts accepting Unicode text input events in a window.
func StartTextInput(window *Window) bool {
	return sdlStartTextInput(window)
}

// StartTextInput starts accepting Unicode text input events in a window.
func (window *Window) StartTextInput() bool {
	return StartTextInput(window)
}

// func StartTextInputWithProperties(window *Window, props PropertiesID) bool {
//	return sdlStartTextInputWithProperties(window, props)
// }

// StopTextInput stops receiving any text input events in a window.
func StopTextInput(window *Window) bool {
	return sdlStopTextInput(window)
}

// StopTextInput stops receiving any text input events in a window.
func (window *Window) StopTextInput() bool {
	return StopTextInput(window)
}

// func TextInputActive(window *Window) bool {
//	return sdlTextInputActive(window)
// }
