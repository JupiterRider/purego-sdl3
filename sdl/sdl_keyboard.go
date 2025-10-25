package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

const (
	PropTextInputTypeNumber             = "SDL.textinput.type"
	PropTextInputCapitalizationNumber   = "SDL.textinput.capitalization"
	PropTextInputAutocorrectBoolean     = "SDL.textinput.autocorrect"
	PropTextInputMultilineBoolean       = "SDL.textinput.multiline"
	PropTextInputAndroidInputTypeNumber = "SDL.textinput.android.inputtype"
)

// [Capitalization] defines the auto capitalization type.
//
// [Capitalization]: https://wiki.libsdl.org/SDL3/SDL_Capitalization
type Capitalization uint32

const (
	CapitalizeNone      Capitalization = iota // No auto-capitalization will be done.
	CapitalizeSentences                       // The first letter of sentences will be capitalized.
	CapitalizeWords                           // The first letter of words will be capitalized.
	CapitalizeLetters                         // All letters will be capitalized.
)

// [TextInputType] defines the text input type.
//
// [TextInputType]: https://wiki.libsdl.org/SDL3/SDL_TextInputType
type TextInputType uint32

const (
	TextInputTypeText                  TextInputType = iota // The input is text.
	TextInputTypeTextName                                   // The input is a person's name.
	TextInputTypeTextEmail                                  // The input is an e-mail address.
	TextInputTypeTextUsername                               // The input is a username.
	TextInputTypeTextPasswordHidden                         // The input is a secure password that is hidden.
	TextInputTypeTextPasswordVisible                        // The input is a secure password that is visible.
	TextInputTypeNumber                                     // The input is a number.
	TextInputTypeNumberPasswordHidden                       // The input is a secure PIN that is hidden.
	TextInputTypeNumberPasswordVisible                      // The input is a secure PIN that is visible.
)

// [KeyboardID] is a unique ID for a keyboard for the time it is connected to the system, and is never reused for the lifetime of the application.
//
// [KeyboardID]: https://wiki.libsdl.org/SDL3/SDL_KeyboardID
type KeyboardID uint32

// [ClearComposition] dismisses the composition window/IME without disabling the subsystem.
//
// [ClearComposition]: https://wiki.libsdl.org/SDL3/SDL_ClearComposition
func ClearComposition(window *Window) bool {
	return sdlClearComposition(window)
}

// [GetKeyboardFocus] queries the window which currently has keyboard focus.
//
// [GetKeyboardFocus]: https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus
func GetKeyboardFocus() *Window {
	return sdlGetKeyboardFocus()
}

// [GetKeyboardNameForID] gets the name of a keyboard.
//
// [GetKeyboardNameForID]: https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID
func GetKeyboardNameForID(instanceId KeyboardID) string {
	return sdlGetKeyboardNameForID(instanceId)
}

// [GetKeyboards] gets a list of currently connected keyboards.
//
// [GetKeyboards]: https://wiki.libsdl.org/SDL3/SDL_GetKeyboards
func GetKeyboards() []KeyboardID {
	var count int32
	keyboards := sdlGetKeyboards(&count)
	defer Free(unsafe.Pointer(keyboards))
	return mem.Copy(keyboards, count)
}

// [GetKeyboardState] gets a snapshot of the current state of the keyboard.
//
// [GetKeyboardState]: https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState
func GetKeyboardState() []bool {
	var numkeys int32
	state := sdlGetKeyboardState(&numkeys)
	return unsafe.Slice(state, numkeys)
}

// [GetKeyFromName] gets a key code from a human-readable name.
//
// [GetKeyFromName]: https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName
func GetKeyFromName(name string) Keycode {
	return sdlGetKeyFromName(name)
}

// [GetKeyFromScancode] gets the key code corresponding to the given scancode according to the current keyboard layout.
//
// [GetKeyFromScancode]: https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode
func GetKeyFromScancode(scancode Scancode, modstate Keymod, keyEvent bool) Keycode {
	return sdlGetKeyFromScancode(scancode, modstate, keyEvent)
}

// [GetKeyName] gets a human-readable name for a key.
//
// [GetKeyName]: https://wiki.libsdl.org/SDL3/SDL_GetKeyName
func GetKeyName(key Keycode) string {
	return sdlGetKeyName(key)
}

// [GetModState] returns an OR'd combination of the modifier keys for the keyboard.
//
// [GetModState]: https://wiki.libsdl.org/SDL3/SDL_GetModState
func GetModState() Keymod {
	return sdlGetModState()
}

// [GetScancodeFromKey] gets the scancode corresponding to the given key code according to the current keyboard layout.
//
// [GetScancodeFromKey]: https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey
func GetScancodeFromKey(key Keycode, modstate *Keymod) Scancode {
	return sdlGetScancodeFromKey(key, modstate)
}

// [GetScancodeFromName] gets a scancode from a human-readable name.
//
// [GetScancodeFromName]: https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName
func GetScancodeFromName(name string) Scancode {
	return sdlGetScancodeFromName(name)
}

// [GetScancodeName] gets a human-readable name for a scancode.
//
// [GetScancodeName]: https://wiki.libsdl.org/SDL3/SDL_GetScancodeName
func GetScancodeName(scancode Scancode) string {
	return sdlGetScancodeName(scancode)
}

// [GetTextInputArea] gets the area used to type Unicode text input.
//
// [GetTextInputArea]: https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea
func GetTextInputArea(window *Window, rect *Rect, cursor *int32) bool {
	return sdlGetTextInputArea(window, rect, cursor)
}

// [HasKeyboard] returns whether a keyboard is currently connected.
//
// [HasKeyboard]: https://wiki.libsdl.org/SDL3/SDL_HasKeyboard
func HasKeyboard() bool {
	return sdlHasKeyboard()
}

// [HasScreenKeyboardSupport] checks whether the platform has screen keyboard support.
//
// [HasScreenKeyboardSupport]: https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport
func HasScreenKeyboardSupport() bool {
	return sdlHasScreenKeyboardSupport()
}

// [ResetKeyboard] clears the state of the keyboard.
//
// [ResetKeyboard]: https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard
func ResetKeyboard() {
	sdlResetKeyboard()
}

// [ScreenKeyboardShown] checks whether the screen keyboard is shown for given window.
//
// [ScreenKeyboardShown]: https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown
func ScreenKeyboardShown(window *Window) bool {
	return sdlScreenKeyboardShown(window)
}

// [SetModState] sets the current key modifier state for the keyboard.
//
// [SetModState]: https://wiki.libsdl.org/SDL3/SDL_SetModState
func SetModState(modstate Keymod) {
	sdlSetModState(modstate)
}

// [SetScancodeName] sets a human-readable name for a scancode.
//
// [SetScancodeName]: https://wiki.libsdl.org/SDL3/SDL_SetScancodeName
func SetScancodeName(scancode Scancode, name string) bool {
	return sdlSetScancodeName(scancode, name)
}

// [SetTextInputArea] sets the area used to type Unicode text input.
//
// [SetTextInputArea]: https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea
func SetTextInputArea(window *Window, rect *Rect, cursor int32) bool {
	return sdlSetTextInputArea(window, rect, cursor)
}

// [StartTextInput] starts accepting Unicode text input events in a window.
//
// [StartTextInput]: https://wiki.libsdl.org/SDL3/SDL_StartTextInput
func StartTextInput(window *Window) bool {
	return sdlStartTextInput(window)
}

// [StartTextInputWithProperties] starts accepting Unicode text input events in a window, with properties describing the input.
//
// [StartTextInputWithProperties]: https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties
func StartTextInputWithProperties(window *Window, props PropertiesID) bool {
	return sdlStartTextInputWithProperties(window, props)
}

// [StopTextInput] stops receiving any text input events in a window.
//
// [StopTextInput]: https://wiki.libsdl.org/SDL3/SDL_StopTextInput
func StopTextInput(window *Window) bool {
	return sdlStopTextInput(window)
}

// [TextInputActive] checks whether or not Unicode text input events are enabled for a window.
//
// [TextInputActive]: https://wiki.libsdl.org/SDL3/SDL_TextInputActive
func TextInputActive(window *Window) bool {
	return sdlTextInputActive(window)
}
