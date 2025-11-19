package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

// [EventAction] is a structure specifying the type of action to request from [PeepEvents].
//
// [EventAction]: https://wiki.libsdl.org/SDL3/SDL_EventAction
type EventAction uint32

const (
	AddEvent  EventAction = iota // Add events to the back of the queue.
	PeekEvent                    // Check but don't remove events from the queue front.
	GetEvent                     // Retrieve/remove events from the front of the queue.
)

// [EventType] is a structure specifying the types of events that can be delivered.
//
// [EventType]: https://wiki.libsdl.org/SDL3/SDL_EventType
type EventType uint32

const (
	EventFirst                      EventType = 0x0   // Unused (do not remove).
	EventQuit                       EventType = 0x100 // User-requested quit.
	EventTerminating                EventType = 0x101 // The application is being terminated by the OS. This event must be handled in a callback set with [AddEventWatch]. Called on iOS in applicationWillTerminate(). Called on Android in onDestroy().
	EventLowMemory                  EventType = 0x102 // The application is low on memory, free memory if possible. This event must be handled in a callback set with [AddEventWatch]. Called on iOS in applicationDidReceiveMemoryWarning(). Called on Android in onTrimMemory().
	EventWillEnterBackground        EventType = 0x103 // The application is about to enter the background. This event must be handled in a callback set with [AddEventWatch]. Called on iOS in applicationWillResignActive(). Called on Android in onPause().
	EventDidEnterBackground         EventType = 0x104 // The application did enter the background and may not get CPU for some time. This event must be handled in a callback set with [AddEventWatch]. Called on iOS in applicationDidEnterBackground(). Called on Android in onPause().
	EventWillEnterForeground        EventType = 0x105 // The application is about to enter the foreground. This event must be handled in a callback set with [AddEventWatch]. Called on iOS in applicationWillEnterForeground(). Called on Android in onResume().
	EventDidEnterForeground         EventType = 0x106 // The application is now interactive. This event must be handled in a callback set with [AddEventWatch]. Called on iOS in applicationDidBecomeActive(). Called on Android in onResume().
	EventLocaleChanged              EventType = 0x107 // The user's locale preferences have changed.
	EventSystemThemeChanged         EventType = 0x108 // The system theme changed.
	EventDisplayOrientation         EventType = 0x151 // Display orientation has changed to data1.
	EventDisplayFirst               EventType = 0x151
	EventDisplayAdded               EventType = 0x152 // Display has been added to the system.
	EventDisplayRemoved             EventType = 0x153 // Display has been removed from the system.
	EventDisplayMoved               EventType = 0x154 // Display has changed position.
	EventDisplayDesktopModeChanged  EventType = 0x155 // Display has changed desktop mode.
	EventDisplayCurrentModeChanged  EventType = 0x156 // Display has changed current mode.
	EventDisplayContentScaleChanged EventType = 0x157 // Display has changed content scale.
	EventDisplayLast                EventType = 0x157
	EventWindowShown                EventType = 0x202 // Window has been shown.
	EventWindowFirst                EventType = 0x202
	EventWindowHidden               EventType = 0x203 // Window has been hidden.
	EventWindowExposed              EventType = 0x204 // Window has been exposed and should be redrawn, and can be redrawn directly from event watchers for this event. data1 is 1 for live-resize expose events, 0 otherwise.
	EventWindowMoved                EventType = 0x205 // Window has been moved to data1, data2.
	EventWindowResized              EventType = 0x206 // Window has been resized to data1xdata2.
	EventWindowPixelSizeChanged     EventType = 0x207 // The pixel size of the window has changed to data1xdata2.
	EventWindowMetalViewResized     EventType = 0x208 // The pixel size of a Metal view associated with the window has changed.
	EventWindowMinimized            EventType = 0x209 // Window has been minimized.
	EventWindowMaximized            EventType = 0x20A // Window has been maximized.
	EventWindowRestored             EventType = 0x20B // Window has been restored to normal size and position.
	EventWindowMouseEnter           EventType = 0x20C // Window has gained mouse focus.
	EventWindowMouseLeave           EventType = 0x20D // Window has lost mouse focus.
	EventWindowFocusGained          EventType = 0x20E // Window has gained keyboard focus.
	EventWindowFocusLost            EventType = 0x20F // Window has lost keyboard focus.
	EventWindowCloseRequested       EventType = 0x210 // The window manager requests that the window be closed.
	EventWindowHitTest              EventType = 0x211 // Window had a hit test that wasn't [HITTEST_NORMAL].
	EventWindowICCProfChanged       EventType = 0x212 // The ICC profile of the window's display has changed.
	EventWindowDisplayChanged       EventType = 0x213 // Window has been moved to display data1.
	EventWindowDisplayScaleChanged  EventType = 0x214 // Window display scale has been changed.
	EventWindowSafeAreaChanged      EventType = 0x215 // The window safe area has been changed.
	EventWindowOccluded             EventType = 0x216 // The window has been occluded.
	EventWindowEnterFullscreen      EventType = 0x217 // The window has entered fullscreen mode.
	EventWindowLeaveFullscreen      EventType = 0x218 // The window has left fullscreen mode.
	EventWindowDestroyed            EventType = 0x219 // The window with the associated ID is being or has been destroyed. If this message is being handled in an event watcher, the window handle is still valid and can still be used to retrieve any properties associated with the window. Otherwise, the handle has already been destroyed and all resources associated with it are invalid.
	EventWindowHDRStateChanged      EventType = 0x21A // Window HDR properties have changed.
	EventWindowLast                 EventType = 0x21A
	EventKeyDown                    EventType = 0x300 // Key pressed.
	EventKeyUp                      EventType = 0x301 // Key released.
	EventTextEditing                EventType = 0x302 // Keyboard text editing (composition).
	EventTextInput                  EventType = 0x303 // Keyboard text input.
	EventKeymapChanged              EventType = 0x304 // Keymap changed due to a system event such as an input language or keyboard layout change.
	EventKeyboardAdded              EventType = 0x305 // A new keyboard has been inserted into the system.
	EventKeyboardRemoved            EventType = 0x306 // A keyboard has been removed.
	EventTextEditingCandidates      EventType = 0x307 // Keyboard text editing candidates.
	EventMouseMotion                EventType = 0x400 // Mouse moved.
	EventMouseButtonDown            EventType = 0x401 // Mouse button pressed.
	EventMouseButtonUp              EventType = 0x402 // Mouse button released.
	EventMouseWheel                 EventType = 0x403 // Mouse wheel motion.
	EventMouseAdded                 EventType = 0x404 // A new mouse has been inserted into the system.
	EventMouseRemoved               EventType = 0x405 // A mouse has been removed.
	EventJoystickAxisMotion         EventType = 0x600 // Joystick axis motion.
	EventJoystickBallMotion         EventType = 0x601 // Joystick trackball motion.
	EventJoystickHatMotion          EventType = 0x602 // Joystick hat position change.
	EventJoystickButtonDown         EventType = 0x603 // Joystick button pressed.
	EventJoystickButtonUp           EventType = 0x604 // Joystick button released.
	EventJoystickAdded              EventType = 0x605 // A new joystick has been inserted into the system.
	EventJoystickRemoved            EventType = 0x606 // An opened joystick has been removed.
	EventJoystickBatteryUpdated     EventType = 0x607 // Joystick battery level change.
	EventJoystickUpdateComplete     EventType = 0x608 // Joystick update is complete.
	EventGamepadAxisMotion          EventType = 0x650 // Gamepad axis motion.
	EventGamepadButtonDown          EventType = 0x651 // Gamepad button pressed.
	EventGamepadButtonUp            EventType = 0x652 // Gamepad button released.
	EventGamepadAdded               EventType = 0x653 // A new gamepad has been inserted into the system.
	EventGamepadRemoved             EventType = 0x654 // A gamepad has been removed.
	EventGamepadRemapped            EventType = 0x655 // The gamepad mapping was updated.
	EventGamepadTouchpadDown        EventType = 0x656 // Gamepad touchpad was touched.
	EventGamepadTouchpadMotion      EventType = 0x657 // Gamepad touchpad finger was moved.
	EventGamepadTouchpadUp          EventType = 0x658 // Gamepad touchpad finger was lifted.
	EventGamepadSensorUpdate        EventType = 0x659 // Gamepad sensor was updated.
	EventGamepadUpdateComplete      EventType = 0x65A // Gamepad update is complete.
	EventGamepadSteamHandleUpdated  EventType = 0x65B // Gamepad Steam handle has changed.
	EventFingerDown                 EventType = 0x700
	EventFingerUp                   EventType = 0x701
	EventFingerMotion               EventType = 0x702
	EventFingerCanceled             EventType = 0x703
	EventClipboardUpdate            EventType = 0x900  // The clipboard changed.
	EventDropFile                   EventType = 0x1000 // The system requests a file open.
	EventDropText                   EventType = 0x1001 // Text/plain drag-and-drop event.
	EventDropBegin                  EventType = 0x1002 // A new set of drops is beginning (nil filename).
	EventDropComplete               EventType = 0x1003 // Current set of drops is now complete (nil filename).
	EventDropPosition               EventType = 0x1004 // Position while moving over the window.
	EventAudioDeviceAdded           EventType = 0x1100 // A new audio device is available.
	EventAudioDeviceRemoved         EventType = 0x1101 // An audio device has been removed.
	EventAudioDeviceFormatChanged   EventType = 0x1102 // An audio device's format has been changed by the system.
	EventSensorUpdate               EventType = 0x1200 // A sensor was updated.
	EventPenProximityIn             EventType = 0x1300 // Pressure-sensitive pen has become available.
	EventPenProximityOut            EventType = 0x1301 // Pressure-sensitive pen has become unavailable.
	EventPenDown                    EventType = 0x1302 // Pressure-sensitive pen touched drawing surface.
	EventPenUp                      EventType = 0x1303 // Pressure-sensitive pen stopped touching drawing surface.
	EventPenButtonDown              EventType = 0x1304 // Pressure-sensitive pen button pressed.
	EventPenButtonUp                EventType = 0x1305 // Pressure-sensitive pen button released.
	EventPenMotion                  EventType = 0x1306 // Pressure-sensitive pen is moving on the tablet.
	EventPenAxis                    EventType = 0x1307 // Pressure-sensitive pen angle/pressure/etc changed.
	EventCameraDeviceAdded          EventType = 0x1400 // A new camera device is available.
	EventCameraDeviceRemoved        EventType = 0x1401 // A camera device has been removed.
	EventCameraDeviceApproved       EventType = 0x1402 // A camera device has been approved for use by the user.
	EventCameraDeviceDenied         EventType = 0x1403 // A camera device has been denied for use by the user.
	EventRenderTargetsReset         EventType = 0x2000 // The render targets have been reset and their contents need to be updated.
	EventRenderDeviceReset          EventType = 0x2001 // The device has been reset and all textures need to be recreated.
	EventRenderDeviceLost           EventType = 0x2002 // The device has been lost and can't be recovered.
	EventPrivate0                   EventType = 0x4000
	EventPrivate1                   EventType = 0x4001
	EventPrivate2                   EventType = 0x4002
	EventPrivate3                   EventType = 0x4003
	EventPollSentinel               EventType = 0x7F00     // Signals the end of an event poll cycle.
	EventUser                       EventType = 0x8000     // Events [EVENT_USER] through [EVENT_LAST] are for your use, and should be allocated with [RegisterEvents].
	EventLast                       EventType = 0xFFFF     // Events [EVENT_USER] through [EVENT_LAST] are for your use, and should be allocated with [RegisterEvents]. This last event is only for bounding internal arrays.
	EventEnumPadding                EventType = 0x7FFFFFFF // This just makes sure the enum is the size of Uint32.
)

// [Event] is a structure specifying the structure for all events in SDLs.
//
// [Event]: https://wiki.libsdl.org/SDL3/SDL_Event
type Event [128]byte

// [Type] returns type of the event.
//
// [Type]: https://wiki.libsdl.org/SDL3/SDL_EventType
func (e *Event) Type() EventType {
	return *(*EventType)(unsafe.Pointer(e))
}

// [Common] returns common fields shared by every event.
//
// [Common]: https://wiki.libsdl.org/SDL3/SDL_CommonEvent
func (e *Event) Common() CommonEvent {
	return *(*CommonEvent)(unsafe.Pointer(e))
}

// [Display] returns display state change event data (event.display.*)
//
// [Display]: https://wiki.libsdl.org/SDL3/SDL_DisplayEvent
func (e *Event) Display() DisplayEvent {
	return *(*DisplayEvent)(unsafe.Pointer(e))
}

// [Window] returns window state change event data (event.window.*)
//
// [Window]: https://wiki.libsdl.org/SDL3/SDL_WindowEvent
func (e *Event) Window() WindowEvent {
	return *(*WindowEvent)(unsafe.Pointer(e))
}

// [KDevice] returns keyboard device event data (event.kdevice.*).
//
// [KDevice]: https://wiki.libsdl.org/SDL3/SDL_KeyboardDeviceEvent
func (e *Event) KDevice() KeyboardDeviceEvent {
	return *(*KeyboardDeviceEvent)(unsafe.Pointer(e))
}

// [Key] returns keyboard button event data (event.key.*).
//
// [Key]: https://wiki.libsdl.org/SDL3/SDL_KeyboardEvent
func (e *Event) Key() KeyboardEvent {
	return *(*KeyboardEvent)(unsafe.Pointer(e))
}

// [Edit] returns keyboard text editing event data (event.edit.*).
//
// [Edit]: https://wiki.libsdl.org/SDL3/SDL_TextEditingEvent
func (e *Event) Edit() TextEditingEvent {
	return *(*TextEditingEvent)(unsafe.Pointer(e))
}

// [EditCandidates] returns keyboard IME candidates event data (event.edit_candidates.*).
//
// [EditCandidates]: https://wiki.libsdl.org/SDL3/SDL_TextEditingCandidatesEvent
func (e *Event) EditCandidates() TextEditingCandidatesEvent {
	return *(*TextEditingCandidatesEvent)(unsafe.Pointer(e))
}

// [Text] returns keyboard text input event data (event.text.*).
//
// [Text]: https://wiki.libsdl.org/SDL3/SDL_TextInputEvent
func (e *Event) Text() TextInputEvent {
	return *(*TextInputEvent)(unsafe.Pointer(e))
}

// [MDevice] returns mouse device event data (event.mdevice.*).
//
// [MDevice]: https://wiki.libsdl.org/SDL3/SDL_MouseDeviceEvent
func (e *Event) MDevice() MouseDeviceEvent {
	return *(*MouseDeviceEvent)(unsafe.Pointer(e))
}

// [Motion] returns mouse motion event data (event.motion.*).
//
// [Motion]: https://wiki.libsdl.org/SDL3/SDL_MouseMotionEvent
func (e *Event) Motion() MouseMotionEvent {
	return *(*MouseMotionEvent)(unsafe.Pointer(e))
}

// [Button] returns mouse button event data (event.button.*).
//
// [Button]: https://wiki.libsdl.org/SDL3/SDL_MouseButtonEvent
func (e *Event) Button() MouseButtonEvent {
	return *(*MouseButtonEvent)(unsafe.Pointer(e))
}

// [Wheel] returns mouse wheel event data (event.wheel.*).
//
// [Wheel]: https://wiki.libsdl.org/SDL3/SDL_MouseWheelEvent
func (e *Event) Wheel() MouseWheelEvent {
	return *(*MouseWheelEvent)(unsafe.Pointer(e))
}

// [JDevice] returns joystick device event data (event.jdevice.*).
//
// [JDevice]: https://wiki.libsdl.org/SDL3/SDL_JoyDeviceEvent
func (e *Event) JDevice() JoyDeviceEvent {
	return *(*JoyDeviceEvent)(unsafe.Pointer(e))
}

// [JAxis] returns joystick axis motion event data (event.jaxis.*).
//
// [JAxis]: https://wiki.libsdl.org/SDL3/SDL_JoyAxisEvent
func (e *Event) JAxis() JoyAxisEvent {
	return *(*JoyAxisEvent)(unsafe.Pointer(e))
}

// [JBall] returns joystick trackball motion event data (event.jball.*).
//
// [JBall]: https://wiki.libsdl.org/SDL3/SDL_JoyBallEvent
func (e *Event) JBall() JoyBallEvent {
	return *(*JoyBallEvent)(unsafe.Pointer(e))
}

// [JHat] returns joystick hat position change event data (event.jhat.*).
//
// [JHat]: https://wiki.libsdl.org/SDL3/SDL_JoyHatEvent
func (e *Event) JHat() JoyHatEvent {
	return *(*JoyHatEvent)(unsafe.Pointer(e))
}

// [JButton] returns joystick button event data (event.jbutton.*).
//
// [JButton]: https://wiki.libsdl.org/SDL3/SDL_JoyButtonEvent
func (e *Event) JButton() JoyButtonEvent {
	return *(*JoyButtonEvent)(unsafe.Pointer(e))
}

// [JBattery] returns joystick battery level change event data (event.jbattery.*).
//
// [JBattery]: https://wiki.libsdl.org/SDL3/SDL_JoyBatteryEvent
func (e *Event) JBattery() JoyBatteryEvent {
	return *(*JoyBatteryEvent)(unsafe.Pointer(e))
}

// [GDevice] returns gamepad device event data (event.gdevice.*).
//
// [GDevice]: https://wiki.libsdl.org/SDL3/SDL_GamepadDeviceEvent
func (e *Event) GDevice() GamepadDeviceEvent {
	return *(*GamepadDeviceEvent)(unsafe.Pointer(e))
}

// [GAxis] returns gamepad axis motion event data (event.gaxis.*).
//
// [GAxis]: https://wiki.libsdl.org/SDL3/SDL_GamepadAxisEvent
func (e *Event) GAxis() GamepadAxisEvent {
	return *(*GamepadAxisEvent)(unsafe.Pointer(e))
}

// [GButton] returns gamepad button event data (event.gbutton.*).
//
// [GButton]: https://wiki.libsdl.org/SDL3/SDL_GamepadButtonEvent
func (e *Event) GButton() GamepadButtonEvent {
	return *(*GamepadButtonEvent)(unsafe.Pointer(e))
}

// [GTouchpad] returns gamepad touchpad event data (event.gtouchpad.*).
//
// [GTouchpad]: https://wiki.libsdl.org/SDL3/SDL_GamepadTouchpadEvent
func (e *Event) GTouchpad() GamepadTouchpadEvent {
	return *(*GamepadTouchpadEvent)(unsafe.Pointer(e))
}

// [GSensor] returns gamepad sensor event data (event.gsensor.*).
//
// [GSensor]: https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEvent
func (e *Event) GSensor() GamepadSensorEvent {
	return *(*GamepadSensorEvent)(unsafe.Pointer(e))
}

// [ADevice] returns audio device event data (event.adevice.*).
//
// [ADevice]: https://wiki.libsdl.org/SDL3/SDL_AudioDeviceEvent
func (e *Event) ADevice() AudioDeviceEvent {
	return *(*AudioDeviceEvent)(unsafe.Pointer(e))
}

// [CDevice] returns camera device event data (event.cdevice.*).
//
// [CDevice]: https://wiki.libsdl.org/SDL3/SDL_CameraDeviceEvent
func (e *Event) CDevice() CameraDeviceEvent {
	return *(*CameraDeviceEvent)(unsafe.Pointer(e))
}

// [Sensor] returns sensor event data (event.sensor.*).
//
// [Sensor]: https://wiki.libsdl.org/SDL3/SDL_SensorEvent
func (e *Event) Sensor() SensorEvent {
	return *(*SensorEvent)(unsafe.Pointer(e))
}

// [Quit] returns quit event data.
//
// [Quit]: https://wiki.libsdl.org/SDL3/SDL_QuitEvent
func (e *Event) Quit() QuitEvent {
	return *(*QuitEvent)(unsafe.Pointer(e))
}

// [User] returns user-defined event data (event.user.*).
//
// [User]: https://wiki.libsdl.org/SDL3/SDL_UserEvent
func (e *Event) User() UserEvent {
	return *(*UserEvent)(unsafe.Pointer(e))
}

// [TFinger] returns touch finger event data (event.tfinger.*).
//
// [TFinger]: https://wiki.libsdl.org/SDL3/SDL_TouchFingerEvent
func (e *Event) TFinger() TouchFingerEvent {
	return *(*TouchFingerEvent)(unsafe.Pointer(e))
}

// [PProximity] returns pressure-sensitive pen proximity event data (event.pproximity.*).
//
// [PProximity]: https://wiki.libsdl.org/SDL3/SDL_PenProximityEvent
func (e *Event) PProximity() PenProximityEvent {
	return *(*PenProximityEvent)(unsafe.Pointer(e))
}

// [PTouch] returns pressure-sensitive pen touched event data (event.ptouch.*).
//
// [PTouch]: https://wiki.libsdl.org/SDL3/SDL_PenTouchEvent
func (e *Event) PTouch() PenTouchEvent {
	return *(*PenTouchEvent)(unsafe.Pointer(e))
}

// [PMotion] returns pressure-sensitive pen motion event data (event.pmotion.*).
//
// [PMotion]: https://wiki.libsdl.org/SDL3/SDL_PenMotionEvent
func (e *Event) PMotion() PenMotionEvent {
	return *(*PenMotionEvent)(unsafe.Pointer(e))
}

// [PButton] returns pressure-sensitive pen button event data (event.pbutton.*).
//
// [PButton]: https://wiki.libsdl.org/SDL3/SDL_PenButtonEvent
func (e *Event) PButton() PenButtonEvent {
	return *(*PenButtonEvent)(unsafe.Pointer(e))
}

// [PAxis] returns pressure-sensitive pen pressure / angle event data (event.paxis.*).
//
// [PAxis]: https://wiki.libsdl.org/SDL3/SDL_PenAxisEvent
func (e *Event) PAxis() PenAxisEvent {
	return *(*PenAxisEvent)(unsafe.Pointer(e))
}

// [Render] returns renderer event data (event.render.*).
//
// [Render]: https://wiki.libsdl.org/SDL3/SDL_RenderEvent
func (e *Event) Render() RenderEvent {
	return *(*RenderEvent)(unsafe.Pointer(e))
}

// [Drop] returns event data related to dropped text or a file open request by the system (event.drop.*).
//
// [Drop]: https://wiki.libsdl.org/SDL3/SDL_DropEvent
func (e *Event) Drop() DropEvent {
	return *(*DropEvent)(unsafe.Pointer(e))
}

// [Clipboard] returns clipboard contents have changed event data (event.clipboard.*).
//
// [Clipboard]: https://wiki.libsdl.org/SDL3/SDL_ClipboardEvent
func (e *Event) Clipboard() ClipboardEvent {
	return *(*ClipboardEvent)(unsafe.Pointer(e))
}

// [CommonEvent] defines the fields shared by every event.
//
// [CommonEvent]: https://wiki.libsdl.org/SDL3/SDL_CommonEvent
type CommonEvent struct {
	Type      EventType // Event type, shared with all events, uint32 to cover user events which are not in the [EventType] enumeration.
	reserved  uint32
	Timestamp uint64 // Timestamp in nanoseconds, populated using [GetTicksNS].
}

// [DisplayEvent] defines the display state change event data (event.display.*).
//
// [DisplayEvent]: https://wiki.libsdl.org/SDL3/SDL_DisplayEvent
type DisplayEvent struct {
	CommonEvent
	DisplayID DisplayID // The associated display.
	Data1     int32     // Event dependent data.
	Data2     int32     // Event dependent data.
}

// [WindowEvent] defines the window state change event data (event.window.*).
//
// [WindowEvent]: https://wiki.libsdl.org/SDL3/SDL_WindowEvent
type WindowEvent struct {
	CommonEvent
	WindowID WindowID // The associated window.
	Data1    int32    // Event dependent data.
	Data2    int32    // Event dependent data.
}

// [KeyboardDeviceEvent] defines the keyboard device event structure (event.kdevice.*).
//
// [KeyboardDeviceEvent]: https://wiki.libsdl.org/SDL3/SDL_KeyboardDeviceEvent
type KeyboardDeviceEvent struct {
	CommonEvent
	Which KeyboardID // The keyboard instance id.
}

// [KeyboardEvent] defines the keyboard button event structure (event.key.*).
//
// [KeyboardEvent]: https://wiki.libsdl.org/SDL3/SDL_KeyboardEvent
type KeyboardEvent struct {
	CommonEvent
	WindowID WindowID   // The window with keyboard focus, if any.
	Which    KeyboardID // The keyboard instance id, or 0 if unknown or virtual.
	Scancode Scancode   // SDL physical key code.
	Key      Keycode    // SDL virtual key code.
	Mod      Keymod     // Current key modifiers.
	Raw      uint16     // The platform dependent scancode for this event.
	Down     bool       // True if the key is pressed.
	Repeat   bool       // True if this is a key repeat.
}

// [TextEditingEvent] defines the keyboard text editing event structure (event.edit.*).
//
// [TextEditingEvent]: https://wiki.libsdl.org/SDL3/SDL_TextEditingEvent
type TextEditingEvent struct {
	CommonEvent
	WindowID WindowID // The window with keyboard focus, if any.
	text     *byte    // The editing text.
	Start    int32    // The start cursor of selected editing text, or -1 if not set.
	Length   int32    // The length of selected editing text, or -1 if not set.
}

// Text returns the editing text.
func (t *TextEditingEvent) Text() string {
	return convert.ToString(t.text)
}

// [TextEditingCandidatesEvent] defines the keyboard IME candidates event structure (event.edit_candidates.*).
//
// [TextEditingCandidatesEvent]: https://wiki.libsdl.org/SDL3/SDL_TextEditingCandidatesEvent
type TextEditingCandidatesEvent struct {
	CommonEvent
	WindowID          WindowID // The window with keyboard focus, if any.
	candidates        **byte   // The list of candidates, or nil if there are no candidates available.
	numCandidates     int32    // The number of strings in [candidates].
	SelectedCandidate int32    // The index of the selected candidate, or -1 if no candidate is selected.
	Horizontal        bool     // true if the list is horizontal, false if it's vertical.
	_                 uint8    // Padding 1.
	_                 uint8    // Padding 2.
	_                 uint8    // Padding 3.
}

// Candidates returns the list of candidates,
// or empty if there are no candidates available.
func (t *TextEditingCandidatesEvent) Candidates() []string {
	if t.candidates == nil || t.numCandidates <= 0 {
		return []string{}
	}
	candidates := make([]string, t.numCandidates)
	for i, v := range unsafe.Slice(t.candidates, t.numCandidates) {
		candidates[i] = convert.ToString(v)
	}
	return candidates
}

// [TextInputEvent] defines the keyboard text input event structure (event.text.*).
//
// [TextInputEvent]: https://wiki.libsdl.org/SDL3/SDL_TextInputEvent
type TextInputEvent struct {
	CommonEvent
	WindowID WindowID // The window with keyboard focus, if any.
	text     *byte    // The input text, UTF-8 encoded.
}

// Text returns the input text, UTF-8 encoded.
func (t *TextInputEvent) Text() string {
	return convert.ToString(t.text)
}

// [MouseDeviceEvent] defines the mouse device event structure (event.mdevice.*).
//
// [MouseDeviceEvent]: https://wiki.libsdl.org/SDL3/SDL_MouseDeviceEvent
type MouseDeviceEvent struct {
	CommonEvent
	Which MouseID // The mouse instance id.
}

// [MouseMotionEvent] defines the mouse motion event structure (event.motion.*).
//
// [MouseMotionEvent]: https://wiki.libsdl.org/SDL3/SDL_MouseMotionEvent
type MouseMotionEvent struct {
	CommonEvent
	WindowID WindowID         // The window with mouse focus, if any.
	Which    MouseID          // The mouse instance id in relative mode, [TOUCH_MOUSEID] for touch events, or 0.
	State    MouseButtonFlags // The current button state.
	X        float32          // X coordinate, relative to window.
	Y        float32          // Y coordinate, relative to window.
	Xrel     float32          // The relative motion in the X direction.
	Yrel     float32          // The relative motion in the Y direction.
}

// [MouseButtonEvent] defines the mouse button event structure (event.button.*).
//
// [MouseButtonEvent]: https://wiki.libsdl.org/SDL3/SDL_MouseButtonEvent
type MouseButtonEvent struct {
	CommonEvent
	WindowID WindowID // The window with mouse focus, if any.
	Which    MouseID  // The mouse instance id in relative mode, [TOUCH_MOUSEID] for touch events, or 0.
	Button   uint8    // The mouse button index.
	Down     bool     // True if the button is pressed.
	Clicks   uint8    // 1 for single-click, 2 for double-click, etc.
	_        uint8    // Padding
	X        float32  // X coordinate, relative to window.
	Y        float32  // Y coordinate, relative to window.
}

// [MouseWheelEvent] defines the mouse wheel event structure (event.wheel.*).
//
// [MouseWheelEvent]: https://wiki.libsdl.org/SDL3/SDL_MouseWheelEvent
type MouseWheelEvent struct {
	CommonEvent
	WindowID  WindowID            // The window with mouse focus, if any.
	Which     MouseID             // The mouse instance id in relative mode or 0.
	X         float32             // The amount scrolled horizontally, positive to the right and negative to the left.
	Y         float32             // The amount scrolled vertically, positive away from the user and negative toward the user.
	Direction MouseWheelDirection // Set to one of the [MOUSEWHEEL_*] defines. When FLIPPED the values in X and Y will be opposite. Multiply by -1 to change them back.
	MouseX    float32             // X coordinate, relative to window
	MouseY    float32             // Y coordinate, relative to window
	IntegerX  int32               // The amount scrolled horizontally, accumulated to whole scroll "ticks" (added in 3.2.12).
	IntegerY  int32               // The amount scrolled vertically, accumulated to whole scroll "ticks" (added in 3.2.12).
}

// [JoyAxisEvent] defines the joystick axis motion event structure (event.jaxis.*).
//
// [JoyAxisEvent]: https://wiki.libsdl.org/SDL3/SDL_JoyAxisEvent
type JoyAxisEvent struct {
	CommonEvent
	Which JoystickID // The joystick instance id.
	Axis  uint8      // The joystick axis index.
	_     uint8      // Padding 1.
	_     uint8      // Padding 2.
	_     uint8      // Padding 3.
	Value int16      // The axis value (range: -32768 to 32767).
	_     uint16     // Padding 4.
}

// [JoyBallEvent] defines the joystick trackball motion event structure (event.jball.*).
//
// [JoyBallEvent]: https://wiki.libsdl.org/SDL3/SDL_JoyBallEvent
type JoyBallEvent struct {
	CommonEvent
	Which JoystickID // The joystick instance id.
	Ball  uint8      // The joystick trackball index.
	_     uint8      // Padding 1.
	_     uint8      // Padding 2.
	_     uint8      // Padding 3.
	Xrel  int16      // The relative motion in the X direction.
	Yrel  int16      // The relative motion in the Y direction.
}

// [JoyHatEvent] defines the joystick hat position change event structure (event.jhat.*).
//
// [JoyHatEvent]: https://wiki.libsdl.org/SDL3/SDL_JoyHatEvent
type JoyHatEvent struct {
	CommonEvent
	Which JoystickID // The joystick instance id.
	Hat   uint8      // The joystick hat index.
	// The hat position value.
	//  * [HAT_LEFTUP] [HAT_UP] [HAT_RIGHTUP]
	//  * [HAT_LEFT] [HAT_CENTERED] [HAT_RIGHT]
	//  * [HAT_LEFTDOWN] [HAT_DOWN] [HAT_RIGHTDOWN]
	//
	//  * Note that zero means the POV is centered.
	Value uint8
	_     uint8 // Padding 1.
	_     uint8 // Padding 2.
}

// [JoyButtonEvent] defines the joystick button event structure (event.jbutton.*).
//
// [JoyButtonEvent]: https://wiki.libsdl.org/SDL3/SDL_JoyButtonEvent
type JoyButtonEvent struct {
	CommonEvent
	Which  JoystickID // The joystick instance id.
	Button uint8      // The joystick button index.
	Down   bool       // True if the button is pressed.
	_      uint8      // Padding 1.
	_      uint8      // Padding 2.
}

// [JoyDeviceEvent] defines the joystick device event structure (event.jdevice.*).
//
// [JoyDeviceEvent]: https://wiki.libsdl.org/SDL3/SDL_JoyDeviceEvent
type JoyDeviceEvent struct {
	CommonEvent
	Which JoystickID // The joystick instance id.
}

// [JoyBatteryEvent] defines the joystick battery level change event structure (event.jbattery.*).
//
// [JoyBatteryEvent]: https://wiki.libsdl.org/SDL3/SDL_JoyBatteryEvent
type JoyBatteryEvent struct {
	CommonEvent
	Which   JoystickID // The joystick instance id.
	State   PowerState // The joystick battery state.
	Percent int32      // The joystick battery percent charge remaining.
}

// [GamepadAxisEvent] defines the gamepad axis motion event structure (event.gaxis.*).
//
// [GamepadAxisEvent]: https://wiki.libsdl.org/SDL3/SDL_GamepadAxisEvent
type GamepadAxisEvent struct {
	CommonEvent
	Which JoystickID // The joystick instance id.
	Axis  uint8      // The gamepad axis ([GamepadAxis]).
	_     uint8      // Padding 1.
	_     uint8      // Padding 2.
	_     uint8      // Padding 3.
	Value int16      // The axis value (range: -32768 to 32767).
	_     uint16     // Padding 4.
}

// [GamepadButtonEvent] defines the gamepad button event structure (event.gbutton.*).
//
// [GamepadButtonEvent]: https://wiki.libsdl.org/SDL3/SDL_GamepadButtonEvent
type GamepadButtonEvent struct {
	CommonEvent
	Which  JoystickID // The joystick instance id.
	Button uint8      // The gamepad button ([GamepadButton]).
	Down   bool       // true if the button is pressed.
	_      uint8      // Padding 1.
	_      uint8      // Padding 2.
}

// [GamepadDeviceEvent] defines the gamepad device event structure (event.gdevice.*).
//
// [GamepadDeviceEvent]: https://wiki.libsdl.org/SDL3/SDL_GamepadDeviceEvent
type GamepadDeviceEvent struct {
	CommonEvent
	Which JoystickID // The joystick instance id.
}

// [GamepadTouchpadEvent] defines the gamepad touchpad event structure (event.gtouchpad.*).
//
// [GamepadTouchpadEvent]: https://wiki.libsdl.org/SDL3/SDL_GamepadTouchpadEvent
type GamepadTouchpadEvent struct {
	CommonEvent
	Which    JoystickID // The joystick instance id.
	Touchpad int32      // The index of the touchpad.
	Finger   int32      // The index of the finger on the touchpad.
	X        float32    // Normalized in the range 0...1 with 0 being on the left.
	Y        float32    // Normalized in the range 0...1 with 0 being at the top.
	Pressure float32    // Normalized in the range 0...1.
}

// [GamepadSensorEvent] defines the gamepad sensor event structure (event.gsensor.*).
//
// [GamepadSensorEvent]: https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEvent
type GamepadSensorEvent struct {
	CommonEvent
	Which           JoystickID // The joystick instance id.
	Sensor          int32      // The type of the sensor, one of the values of [SensorType].
	Data            [3]float32 // Up to 3 values from the sensor, as defined in SDL_sensor.h.
	SensorTimestamp uint64     // The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock.
}

// [AudioDeviceEvent] defines the audio device event structure (event.adevice.*).
//
// [AudioDeviceEvent]: https://wiki.libsdl.org/SDL3/SDL_AudioDeviceEvent
type AudioDeviceEvent struct {
	CommonEvent
	Which     AudioDeviceID // [AudioDeviceID] for the device being added or removed or changing.
	Recording bool          // False if a playback device, true if a recording device.
	_         uint8         // Padding 1.
	_         uint8         // Padding 2.
	_         uint8         // Padding 3.
}

// [CameraDeviceEvent] defines the camera device event structure (event.cdevice.*).
//
// [CameraDeviceEvent]: https://wiki.libsdl.org/SDL3/SDL_CameraDeviceEvent
type CameraDeviceEvent struct {
	CommonEvent
	Which CameraID // [CameraID] for the device being added or removed or changing.
}

// [RenderEvent] defines the renderer event structure (event.render.*).
//
// [RenderEvent]: https://wiki.libsdl.org/SDL3/SDL_RenderEvent
type RenderEvent struct {
	CommonEvent
	WindowID WindowID // The window containing the renderer in question.
}

// [TouchFingerEvent] defines the touch finger event structure (event.tfinger.*).
//
// [TouchFingerEvent]: https://wiki.libsdl.org/SDL3/SDL_TouchFingerEvent
type TouchFingerEvent struct {
	CommonEvent
	TouchID  TouchID // The touch device id.
	FingerID FingerID
	X        float32  // Normalized in the range 0...1.
	Y        float32  // Normalized in the range 0...1.
	Dx       float32  // Normalized in the range -1...1.
	Dy       float32  // Normalized in the range -1...1.
	Pressure float32  // Normalized in the range 0...1.
	WindowID WindowID // The window underneath the finger, if any.
}

// [PenProximityEvent] defines the pressure-sensitive pen proximity event structure (event.pproximity.*).
//
// [PenProximityEvent]: https://wiki.libsdl.org/SDL3/SDL_PenProximityEvent
type PenProximityEvent struct {
	CommonEvent
	WindowID WindowID // The window with pen focus, if any.
	Which    PenID    // The pen instance id.
}

// [PenMotionEvent] defines the pressure-sensitive pen motion event structure (event.pmotion.*).
//
// [PenMotionEvent]: https://wiki.libsdl.org/SDL3/SDL_PenMotionEvent
type PenMotionEvent struct {
	CommonEvent
	WindowID WindowID      // The window with pen focus, if any.
	Which    PenID         // The pen instance id.
	PenState PenInputFlags // Complete pen input state at time of event.
	X        float32       // X coordinate, relative to window.
	Y        float32       // Y coordinate, relative to window.
}

// [PenTouchEvent] defines the pressure-sensitive pen touched event structure (event.ptouch.*).
//
// [PenTouchEvent]: https://wiki.libsdl.org/SDL3/SDL_PenTouchEvent
type PenTouchEvent struct {
	CommonEvent
	WindowID WindowID      // The window with pen focus, if any.
	Which    PenID         // The pen instance id.
	PenState PenInputFlags // Complete pen input state at time of event.
	X        float32       // X coordinate, relative to window.
	Y        float32       // Y coordinate, relative to window.
	Eraser   bool          // True if eraser end is used (not all pens support this).
	Down     bool          // True if the pen is touching or false if the pen is lifted off.
}

// [PenButtonEvent] defines the pressure-sensitive pen button event structure (event.pbutton.*).
//
// [PenButtonEvent]: https://wiki.libsdl.org/SDL3/SDL_PenButtonEvent
type PenButtonEvent struct {
	CommonEvent
	WindowID WindowID      // The window with mouse focus, if any.
	Which    PenID         // The pen instance id.
	PenState PenInputFlags // Complete pen input state at time of event.
	X        float32       // X coordinate, relative to window.
	Y        float32       // Y coordinate, relative to window.
	Button   uint8         // The pen button index (first button is 1).
	Down     bool          // True if the button is pressed.
}

// [PenAxisEvent] defines the pressure-sensitive pen pressure / angle event structure (event.paxis.*).
//
// [PenAxisEvent]: https://wiki.libsdl.org/SDL3/SDL_PenAxisEvent
type PenAxisEvent struct {
	CommonEvent
	WindowID WindowID      // The window with pen focus, if any.
	Which    PenID         // The pen instance id.
	PenState PenInputFlags // Complete pen input state at time of event.
	X        float32       // X coordinate, relative to window.
	Y        float32       // Y coordinate, relative to window.
	Axis     PenAxis       // Axis that has changed.
	Value    float32       // New value of axis.
}

// [DropEvent] is an event used to drop text or request a file open by the system (event.drop.*).
//
// [DropEvent]: https://wiki.libsdl.org/SDL3/SDL_DropEvent
type DropEvent struct {
	CommonEvent
	WindowID WindowID // The window that was dropped on, if any.
	X        float32  // X coordinate, relative to window (not on begin).
	Y        float32  // Y coordinate, relative to window (not on begin).
	source   *byte    // The source app that sent this drop event, or nil if that isn't available.
	data     *byte    // The text for [EVENT_DROP_TEXT] and the file name for [EVENT_DROP_FILE], nil for other events.
}

// Source returns the source app that sent this drop event, or nil if that isn't available.
func (d *DropEvent) Source() string {
	return convert.ToString(d.source)
}

// Data returns the text for [EVENT_DROP_TEXT] and the file name for [EVENT_DROP_FILE], nil for other events.
func (d *DropEvent) Data() string {
	return convert.ToString(d.data)
}

// [ClipboardEvent] is an event triggered when the clipboard contents have changed (event.clipboard.*).
//
// [ClipboardEvent]: https://wiki.libsdl.org/SDL3/SDL_ClipboardEvent
type ClipboardEvent struct {
	CommonEvent
	Owner        bool   // Are we owning the clipboard (internal update).
	NumMimeTypes int32  // Number of mime types.
	mimeTypes    **byte // Current mime types.
}

// MimeTypes returns current mime types.
func (c *ClipboardEvent) MimeTypes() []string {
	mimeTypes := make([]string, c.NumMimeTypes)
	for i, v := range unsafe.Slice(c.mimeTypes, c.NumMimeTypes) {
		mimeTypes[i] = convert.ToString(v)
	}

	return mimeTypes
}

// [SensorEvent] defines the sensor event structure (event.sensor.*).
//
// [SensorEvent]: https://wiki.libsdl.org/SDL3/SDL_SensorEvent
type SensorEvent struct {
	CommonEvent
	Which           SensorID   // The instance ID of the sensor.
	Data            [6]float32 // Up to 6 values from the sensor - additional values can be queried using [GetSensorData].
	SensorTimestamp uint64     // The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock.
}

// [QuitEvent] is a structure specifying the "quit requested" events.
//
// [QuitEvent]: https://wiki.libsdl.org/SDL3/SDL_QuitEvent
type QuitEvent struct {
	CommonEvent
}

// [UserEvent] is a user-defined event type (event.user.*).
//
// [UserEvent]: https://wiki.libsdl.org/SDL3/SDL_UserEvent
type UserEvent struct {
	CommonEvent
	WindowID WindowID       // The associated window if any.
	Code     int32          // User defined event code.
	Data1    unsafe.Pointer // User defined data pointer.
	Data2    unsafe.Pointer // User defined data pointer.
}

// [EventFilter] is a C function pointer used for callbacks that watch the event queue. Use [NewEventFilter] for creation.
//
// [EventFilter]: https://wiki.libsdl.org/SDL3/SDL_EventFilter
type EventFilter uintptr

// NewEventFilter converts the Go function to a C function pointer.
func NewEventFilter(filter func(userdata unsafe.Pointer, event *Event) bool) EventFilter {
	// workaround to avoid panic "expected function with one uintptr-sized result" on Windows
	cb := purego.NewCallback(func(userdata unsafe.Pointer, event *Event) uintptr {
		if filter(userdata, event) {
			return 1
		}
		return 0
	})

	return EventFilter(cb)
}

// [PollEvent] polls for currently pending events.
//
// [PollEvent]: https://wiki.libsdl.org/SDL3/SDL_PollEvent
func PollEvent(event *Event) bool {
	ret, _, _ := purego.SyscallN(sdlPollEvent, uintptr(unsafe.Pointer(event)))
	return byte(ret) != 0
}

// [AddEventWatch] adds a callback to be triggered when an event is added to the event queue.
//
// [AddEventWatch]: https://wiki.libsdl.org/SDL3/SDL_AddEventWatch
func AddEventWatch(filter EventFilter, userdata unsafe.Pointer) bool {
	return sdlAddEventWatch(filter, userdata)
}

// [EventEnabled] returns true if the event is being processed, false otherwise.
//
// [EventEnabled]: https://wiki.libsdl.org/SDL3/SDL_EventEnabled
func EventEnabled(eventType EventType) bool {
	return sdlEventEnabled(eventType)
}

// [FilterEvents] runs a specific filter function on the current event queue, removing any events for which the filter returns false.
//
// [FilterEvents]: https://wiki.libsdl.org/SDL3/SDL_FilterEvents
func FilterEvents(filter EventFilter, userdata unsafe.Pointer) {
	sdlFilterEvents(filter, userdata)
}

// [FlushEvent] clears events of a specific type from the event queue.
//
// [FlushEvent]: https://wiki.libsdl.org/SDL3/SDL_FlushEvent
func FlushEvent(eventType EventType) {
	sdlFlushEvent(eventType)
}

// [FlushEvents] clears events of a range of types from the event queue.
//
// [FlushEvents]: https://wiki.libsdl.org/SDL3/SDL_FlushEvents
func FlushEvents(minType, maxType EventType) {
	sdlFlushEvents(minType, maxType)
}

// [GetEventFilter] queries the current event filter.
//
// [GetEventFilter]: https://wiki.libsdl.org/SDL3/SDL_GetEventFilter
func GetEventFilter(filter *EventFilter, userdata *unsafe.Pointer) bool {
	return sdlGetEventFilter(filter, userdata)
}

// [GetWindowFromEvent] returns the associated window with an event or nil if there is none.
//
// [GetWindowFromEvent]: https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent
func GetWindowFromEvent(event *Event) *Window {
	return sdlGetWindowFromEvent(event)
}

// [HasEvent] checks for the existence of a certain event type in the event queue.
//
// [HasEvent]: https://wiki.libsdl.org/SDL3/SDL_HasEvent
func HasEvent(eventType EventType) bool {
	return sdlHasEvent(eventType)
}

// [HasEvents] checks for the existence of certain event types in the event queue.
//
// Returns true if events with type >= minType and <= maxType are present, or false if not.
//
// [HasEvents]: https://wiki.libsdl.org/SDL3/SDL_HasEvents
func HasEvents(minType, maxType EventType) bool {
	return sdlHasEvents(minType, maxType)
}

// [PeepEvents] checks the event queue for messages and optionally returns them.
//
// Example:
//
//	sdl.PumpEvents()
//	var events [2]sdl.Event
//	sdl.PeepEvents(&events[0], 2, sdl.PeekEvent, sdl.EventFirst, sdl.EventLast)
//
// [PeepEvents]: https://wiki.libsdl.org/SDL3/SDL_PeepEvents
func PeepEvents(events *Event, numevents int32, action EventAction, minType, maxType EventType) int32 {
	return sdlPeepEvents(events, numevents, action, minType, maxType)
}

// [PumpEvents] updates the event queue and internal input device state.
//
// [PumpEvents]: https://wiki.libsdl.org/SDL3/SDL_PumpEvents
func PumpEvents() {
	sdlPumpEvents()
}

// [PushEvent] adds an event to the event queue.
//
// [PushEvent]: https://wiki.libsdl.org/SDL3/SDL_PushEvent
func PushEvent(event *Event) bool {
	return sdlPushEvent(event)
}

// [RegisterEvents] allocates a set of user-defined events, and return the beginning event number for that set of events.
//
// [RegisterEvents]: https://wiki.libsdl.org/SDL3/SDL_RegisterEvents
func RegisterEvents(numevents int32) uint32 {
	return sdlRegisterEvents(numevents)
}

// [RemoveEventWatch] removes an event watch callback added with [AddEventWatch].
//
// This function takes the same input as [AddEventWatch] to identify and delete the corresponding callback.
//
// [RemoveEventWatch]: https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch
func RemoveEventWatch(filter EventFilter, userdata unsafe.Pointer) {
	sdlRemoveEventWatch(filter, userdata)
}

// [SetEventEnabled] sets the state of processing events by type.
//
// [SetEventEnabled]: https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled
func SetEventEnabled(eventType EventType, enabled bool) {
	sdlSetEventEnabled(eventType, enabled)
}

// [SetEventFilter] sets up a filter to process all events before they are added to the internal event queue.
//
// [SetEventFilter]: https://wiki.libsdl.org/SDL3/SDL_SetEventFilter
func SetEventFilter(filter EventFilter, userdata unsafe.Pointer) {
	sdlSetEventFilter(filter, userdata)
}

// [WaitEvent] waits indefinitely for the next available event.
//
// [WaitEvent]: https://wiki.libsdl.org/SDL3/SDL_WaitEvent
func WaitEvent(event *Event) bool {
	return sdlWaitEvent(event)
}

// [WaitEventTimeout] waits until the specified timeout (in milliseconds) for the next available event.
//
// [WaitEventTimeout]: https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout
func WaitEventTimeout(event *Event, timeoutMS int32) bool {
	return sdlWaitEventTimeout(event, timeoutMS)
}
