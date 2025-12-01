package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

// [MessageBoxColorType] is an enumeration of indices inside the colors array of [MessageBoxColorScheme].
//
// [MessageBoxColorType]: https://wiki.libsdl.org/SDL3/SDL_MessageBoxColorType
type MessageBoxColorType uint32

const (
	MessageBoxColorBackground MessageBoxColorType = iota
	MessageBoxColorText
	MessageBoxColorButtonBorder
	MessageBoxColorButtonBackground
	MessageBoxColorButtonSelected
	MessageBoxColorCount
)

// [MessageBoxFlags] defines the message box flags.
//
// [MessageBoxFlags]: https://wiki.libsdl.org/SDL3/SDL_MessageBoxFlags
type MessageBoxFlags uint32

const (
	MessageBoxError              MessageBoxFlags = 0x00000010 // Error dialog.
	MessageBoxWarning            MessageBoxFlags = 0x00000020 // Warning dialog.
	MessageBoxInformation        MessageBoxFlags = 0x00000040 // Informational dialog.
	MessageBoxButtonsLeftToRight MessageBoxFlags = 0x00000080 // Buttons placed left to right.
	MessageBoxButtonsRightToLeft MessageBoxFlags = 0x00000100 // Buttons placed right to left.
)

// [MessageBoxButtonFlags] defines the [MessageBoxButtonData] flags.
//
// [MessageBoxButtonFlags]: https://wiki.libsdl.org/SDL3/SDL_MessageBoxButtonFlags
type MessageBoxButtonFlags uint32

const (
	MessageBoxButtonReturnKeyDefault MessageBoxButtonFlags = 0x00000001
	MessageBoxButtonEscapeKeyDefault MessageBoxButtonFlags = 0x00000002
)

// [MessageBoxButtonData] defines the individual button data.
//
// [MessageBoxButtonData]: https://wiki.libsdl.org/SDL3/SDL_MessageBoxButtonData
type MessageBoxButtonData struct {
	Flags    MessageBoxButtonFlags
	ButtonID int32 // User defined button id (value returned via [ShowMessageBox]).
	text     *byte // The UTF-8 button text.
}

// Text gets the UTF-8 button text.
func (m *MessageBoxButtonData) Text() string {
	return convert.ToString(m.text)
}

// SetText sets the text of a button.
func (m *MessageBoxButtonData) SetText(s string) {
	m.text = convert.ToBytePtr(s)
}

// [MessageBoxColor] defines the rgb value used in a message box color scheme.
//
// [MessageBoxColor]: https://wiki.libsdl.org/SDL3/SDL_MessageBoxColor
type MessageBoxColor struct {
	R, G, B uint8
}

// [MessageBoxColorScheme] is a set of colors to use for message box dialogs.
//
// [MessageBoxColorScheme]: https://wiki.libsdl.org/SDL3/SDL_MessageBoxColorScheme
type MessageBoxColorScheme struct {
	Colors [MessageBoxColorCount]MessageBoxColor
}

// [MessageBoxData] defines the messagebox structure containing title, text, window, etc.
//
// [MessageBoxData]: https://wiki.libsdl.org/SDL3/SDL_MessageBoxData
type MessageBoxData struct {
	Flags       MessageBoxFlags
	Window      *Window // Parent window, can be nil.
	title       *byte   // UTF-8 title.
	message     *byte   // UTF-8 message text.
	numbuttons  int32
	buttons     *MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme // [MessageBoxColorScheme], can be nil to use system settings.
}

// Title gets the UTF-8 title.
func (m *MessageBoxData) Title() string {
	return convert.ToString(m.title)
}

// SetTitle sets the title of a MessageBox.
func (m *MessageBoxData) SetTitle(s string) {
	m.title = convert.ToBytePtr(s)
}

// Message gets the UTF-8 message text.
func (m *MessageBoxData) Message() string {
	return convert.ToString(m.message)
}

// SetMessage sets the message text of a MessageBox.
func (m *MessageBoxData) SetMessage(s string) {
	m.message = convert.ToBytePtr(s)
}

// Buttons gets MessageBox buttons.
func (m *MessageBoxData) Buttons() []MessageBoxButtonData {
	return unsafe.Slice(m.buttons, m.numbuttons)
}

// SetButtons sets MessageBox buttons.
func (m *MessageBoxData) SetButtons(buttons ...MessageBoxButtonData) {
	m.numbuttons = int32(len(buttons))
	if m.numbuttons == 0 {
		m.buttons = nil
		return
	}
	m.buttons = &buttons[0]
}

// [ShowMessageBox] creates a modal message box.
//
// [ShowMessageBox]: https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox
func ShowMessageBox(messageboxdata *MessageBoxData, buttonid *int32) bool {
	return sdlShowMessageBox(messageboxdata, buttonid)
}

// [ShowSimpleMessageBox] displays a simple modal message box.
//
// [ShowSimpleMessageBox]: https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox
func ShowSimpleMessageBox(flags MessageBoxFlags, title string, message string, window *Window) bool {
	return sdlShowSimpleMessageBox(flags, title, message, window)
}
