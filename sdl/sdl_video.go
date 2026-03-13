package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

const (
	PropGlobalVideoWaylandWlDisplayPointer = "SDL.video.wayland.wl_display"

	PropDisplayHdrEnabledBoolean            = "SDL.display.HDR_enabled"
	PropDisplayKmsdrmPanelOrientationNumber = "SDL.display.KMSDRM.panel_orientation"
	PropDisplayWaylandWlOutputPointer       = "SDL.display.wayland.wl_output"
	PropDisplayWindowsHmonitorPointer       = "SDL.display.windows.hmonitor"

	PropWindowCreateAlwaysOnTopBoolean              = "SDL.window.create.always_on_top"
	PropWindowCreateBorderlessBoolean               = "SDL.window.create.borderless"
	PropWindowCreateConstrainPopupBoolean           = "SDL.window.create.constrain_popup"
	PropWindowCreateFocusableBoolean                = "SDL.window.create.focusable"
	PropWindowCreateExternalGraphicsContextBoolean  = "SDL.window.create.external_graphics_context"
	PropWindowCreateFlagsNumber                     = "SDL.window.create.flags"
	PropWindowCreateFullscreenBoolean               = "SDL.window.create.fullscreen"
	PropWindowCreateHeightNumber                    = "SDL.window.create.height"
	PropWindowCreateHiddenBoolean                   = "SDL.window.create.hidden"
	PropWindowCreateHighPixelDensityBoolean         = "SDL.window.create.high_pixel_density"
	PropWindowCreateMaximizedBoolean                = "SDL.window.create.maximized"
	PropWindowCreateMenuBoolean                     = "SDL.window.create.menu"
	PropWindowCreateMetalBoolean                    = "SDL.window.create.metal"
	PropWindowCreateMinimizedBoolean                = "SDL.window.create.minimized"
	PropWindowCreateModalBoolean                    = "SDL.window.create.modal"
	PropWindowCreateMouseGrabbedBoolean             = "SDL.window.create.mouse_grabbed"
	PropWindowCreateOpenglBoolean                   = "SDL.window.create.opengl"
	PropWindowCreateParentPointer                   = "SDL.window.create.parent"
	PropWindowCreateResizableBoolean                = "SDL.window.create.resizable"
	PropWindowCreateTitleString                     = "SDL.window.create.title"
	PropWindowCreateTransparentBoolean              = "SDL.window.create.transparent"
	PropWindowCreateTooltipBoolean                  = "SDL.window.create.tooltip"
	PropWindowCreateUtilityBoolean                  = "SDL.window.create.utility"
	PropWindowCreateVulkanBoolean                   = "SDL.window.create.vulkan"
	PropWindowCreateWidthNumber                     = "SDL.window.create.width"
	PropWindowCreateXNumber                         = "SDL.window.create.x"
	PropWindowCreateYNumber                         = "SDL.window.create.y"
	PropWindowCreateCocoaWindowPointer              = "SDL.window.create.cocoa.window"
	PropWindowCreateCocoaViewPointer                = "SDL.window.create.cocoa.view"
	PropWindowCreateWindowscenePointer              = "SDL.window.create.uikit.windowscene"
	PropWindowCreateWaylandSurfaceRoleCustomBoolean = "SDL.window.create.wayland.surface_role_custom"
	PropWindowCreateWaylandCreateEglWindowBoolean   = "SDL.window.create.wayland.create_egl_window"
	PropWindowCreateWaylandWlSurfacePointer         = "SDL.window.create.wayland.wl_surface"
	PropWindowCreateWin32HwndPointer                = "SDL.window.create.win32.hwnd"
	PropWindowCreateWin32PixelFormatHwndPointer     = "SDL.window.create.win32.pixel_format_hwnd"
	PropWindowCreateX11WindowNumber                 = "SDL.window.create.x11.window"
	PropWindowCreateEmscriptenCanvasIdString        = "SDL.window.create.emscripten.canvas_id"
	PropWindowCreateEmscriptenKeyboardElementString = "SDL.window.create.emscripten.keyboard_element"

	PropWindowShapePointer                         = "SDL.window.shape"
	PropWindowHdrEnabledBoolean                    = "SDL.window.HDR_enabled"
	PropWindowSdrWhiteLevelFloat                   = "SDL.window.SDR_white_level"
	PropWindowHdrHeadroomFloat                     = "SDL.window.HDR_headroom"
	PropWindowAndroidWindowPointer                 = "SDL.window.android.window"
	PropWindowAndroidSurfacePointer                = "SDL.window.android.surface"
	PropWindowUikitWindowPointer                   = "SDL.window.uikit.window"
	PropWindowUikitMetalViewTagNumber              = "SDL.window.uikit.metal_view_tag"
	PropWindowUikitOpenglFramebufferNumber         = "SDL.window.uikit.opengl.framebuffer"
	PropWindowUikitOpenglRenderbufferNumber        = "SDL.window.uikit.opengl.renderbuffer"
	PropWindowUikitOpenglResolveFramebufferNumber  = "SDL.window.uikit.opengl.resolve_framebuffer"
	PropWindowKmsdrmDeviceIndexNumber              = "SDL.window.kmsdrm.dev_index"
	PropWindowKmsdrmDrmFdNumber                    = "SDL.window.kmsdrm.drm_fd"
	PropWindowKmsdrmGbmDevicePointer               = "SDL.window.kmsdrm.gbm_dev"
	PropWindowCocoaWindowPointer                   = "SDL.window.cocoa.window"
	PropWindowCocoaMetalViewTagNumber              = "SDL.window.cocoa.metal_view_tag"
	PropWindowOpenvrOverlayIdNumber                = "SDL.window.openvr.overlay_id"
	PropWindowVivanteDisplayPointer                = "SDL.window.vivante.display"
	PropWindowVivanteWindowPointer                 = "SDL.window.vivante.window"
	PropWindowVivanteSurfacePointer                = "SDL.window.vivante.surface"
	PropWindowWin32HwndPointer                     = "SDL.window.win32.hwnd"
	PropWindowWin32HdcPointer                      = "SDL.window.win32.hdc"
	PropWindowWin32InstancePointer                 = "SDL.window.win32.instance"
	PropWindowWaylandDisplayPointer                = "SDL.window.wayland.display"
	PropWindowWaylandSurfacePointer                = "SDL.window.wayland.surface"
	PropWindowWaylandViewportPointer               = "SDL.window.wayland.viewport"
	PropWindowWaylandEglWindowPointer              = "SDL.window.wayland.egl_window"
	PropWindowWaylandXdgSurfacePointer             = "SDL.window.wayland.xdg_surface"
	PropWindowWaylandXdgToplevelPointer            = "SDL.window.wayland.xdg_toplevel"
	PropWindowWaylandXdgToplevelExportHandleString = "SDL.window.wayland.xdg_toplevel_export_handle"
	PropWindowWaylandXdgPopupPointer               = "SDL.window.wayland.xdg_popup"
	PropWindowWaylandXdgPositionerPointer          = "SDL.window.wayland.xdg_positioner"
	PropWindowX11DisplayPointer                    = "SDL.window.x11.display"
	PropWindowX11ScreenNumber                      = "SDL.window.x11.screen"
	PropWindowX11WindowNumber                      = "SDL.window.x11.window"
	PropWindowEmscriptenCanvasIdString             = "SDL.window.emscripten.canvas_id"
	PropWindowEmscriptenKeyboardElementString      = "SDL.window.emscripten.keyboard_element"
)

const (
	WindowSurfaceVsyncDisabled int32 = 0
	WindowSurfaceVsyncAdaptive int32 = (-1)
)

// [DisplayID] is a unique ID for a display for the time it is connected to the system, and is never reused for the lifetime of the application.
//
// [DisplayID]: https://wiki.libsdl.org/SDL3/SDL_DisplayID
type DisplayID uint32

// [WindowID] is a unique ID for a window.
//
// [WindowID]: https://wiki.libsdl.org/SDL3/SDL_WindowID
type WindowID uint32

// [HitTestResult] possible return values from the [HitTest] callback.
//
// [HitTestResult]: https://wiki.libsdl.org/SDL3/SDL_HitTestResult
type HitTestResult uint32

// [HitTest] callback used for hit-testing.
//
// [HitTest]: https://wiki.libsdl.org/SDL3/SDL_HitTest
type HitTest func(window *Window, point *Point, data unsafe.Pointer) HitTestResult

const (
	HitTestNormal            HitTestResult = iota // Region is normal. No special properties.
	HitTestDraggable                              // Region can drag entire window.
	HitTestResizeTopLeft                          // Region is the resizable top-left corner border.
	HitTestResizeTop                              // Region is the resizable top border.
	HitTestResizeTopRight                         // Region is the resizable top-right corner border.
	HitTestResizeRight                            // Region is the resizable right border.
	HitTestResizeBottomRight                      // Region is the resizable bottom-right corner border.
	HitTestResizeBottom                           // Region is the resizable bottom border.
	HitTestResizeBottomLeft                       // Region is the resizable bottom-left corner border.
	HitTestResizeLeft                             // Region is the resizable left border.
)

// [GLAttr] an enumeration of OpenGL configuration attributes.
//
// [GLAttr]: https://wiki.libsdl.org/SDL3/SDL_GLAttr
type GLAttr uint32

const (
	GLRedSize                  GLAttr = iota // The minimum number of bits for the red channel of the color buffer; defaults to 8.
	GLGreenSize                              // The minimum number of bits for the green channel of the color buffer; defaults to 8.
	GLBlueSize                               // The minimum number of bits for the blue channel of the color buffer; defaults to 8.
	GLAlphaSize                              // The minimum number of bits for the alpha channel of the color buffer; defaults to 8.
	GLBufferSize                             // The minimum number of bits for frame buffer size; defaults to 0.
	GLDoubleBuffer                           // Whether the output is single or double buffered; defaults to double buffering on.
	GLDepthSize                              // The minimum number of bits in the depth buffer; defaults to 16.
	GLStencilSize                            // The minimum number of bits in the stencil buffer; defaults to 0.
	GLAccumRedSize                           // The minimum number of bits for the red channel of the accumulation buffer; defaults to 0.
	GLAccumGreenSize                         // The minimum number of bits for the green channel of the accumulation buffer; defaults to 0.
	GLAccumBlueSize                          // The minimum number of bits for the blue channel of the accumulation buffer; defaults to 0.
	GLAccumAlphaSize                         // The minimum number of bits for the alpha channel of the accumulation buffer; defaults to 0.
	GLStereo                                 // Whether the output is stereo 3D; defaults to off.
	GLMultisampleBuffers                     // The number of buffers used for multisample anti-aliasing; defaults to 0.
	GLMultisampleSamples                     // The number of samples used around the current pixel used for multisample anti-aliasing.
	GLAcceleratedVisual                      // Set to 1 to require hardware acceleration, set to 0 to force software rendering; defaults to allow either.
	GLRetainedBacking                        // Not used (deprecated).
	GLContextMajorVersion                    // OpenGL context major version.
	GLContextMinorVersion                    // OpenGL context minor version.
	GLContextFlags                           // Some combination of 0 or more of elements of the SDL_GLContextFlag enumeration; defaults to 0.
	GLContextProfileMask                     // Type of GL context (Core, Compatibility, ES). See SDL_GLProfile; default value depends on platform.
	GLShareWithCurrentContext                // OpenGL context sharing; defaults to 0.
	GLFramebufferSRGBCapable                 // Requests sRGB capable visual; defaults to 0.
	GLContextReleaseBehavior                 // Sets context the release behavior. See SDL_GLContextReleaseFlag; defaults to FLUSH.
	GLContextResetNotification               // Set context reset notification. See SDL_GLContextResetNotification; defaults to NO_NOTIFICATION.
	GLContextNoError
	GLFloatBuffers
	GLEGLPlatform
)

// [GLProfile] defines possible values to be set for the [GLContextProfileMask] attribute.
//
// [GLProfile]: https://wiki.libsdl.org/SDL3/SDL_GLProfile
type GLProfile uint32

const (
	GLContextProfileCore          GLProfile = 0x0001 // OpenGL Core Profile context.
	GLContextProfileCompatibility GLProfile = 0x0002 // OpenGL Compatibility Profile context.
	GLContextProfileES            GLProfile = 0x0004 // GLX_CONTEXT_ES2_PROFILE_BIT_EXT.
)

// [GLContextFlag] defines possible flags to be set for the [GLContextFlags] attribute.
//
// [GLContextFlag]: https://wiki.libsdl.org/SDL3/SDL_GLContextFlag
type GLContextFlag uint32

const (
	GLContextDebugFlag             GLContextFlag = 0x0001
	GLContextForwardCompatibleFlag GLContextFlag = 0x0002
	GLContextRobustAccessFlag      GLContextFlag = 0x0004
	GLContextResetIsolationFlag    GLContextFlag = 0x0008
)

// [GLContextReleaseFlag] defines possible values to be set for the [GLContextReleaseBehavior] attribute.
//
// [GLContextReleaseFlag]: https://wiki.libsdl.org/SDL3/SDL_GLContextReleaseFlag
type GLContextReleaseFlag uint32

const (
	GLContextReleaseBehaviorNone  GLContextReleaseFlag = 0x0000
	GLContextReleaseBehaviorFlush GLContextReleaseFlag = 0x0001
)

// [GLContextResetNotification_] defines possible values to be set [GLContextResetNotification] attribute.
//
// [GLContextResetNotification_]: https://wiki.libsdl.org/SDL3/SDL_GLContextResetNotification
type GLContextResetNotification_ uint32

const (
	GLContextResetNoNotification GLContextResetNotification_ = 0x0000
	GLContextResetLoseContext    GLContextResetNotification_ = 0x0001
)

// [FlashOperation] window flash operation.
//
// [FlashOperation]: https://wiki.libsdl.org/SDL3/SDL_FlashOperation
type FlashOperation uint32

const (
	FlashCancel       FlashOperation = iota // Cancel any window flash state
	FlashBriefly                            // Flash the window briefly to get attention
	FlashUntilFocused                       // Flash the window until it gets focus
)

// [ProgressState] defines window progress state
//
// [ProgressState]: https://wiki.libsdl.org/SDL3/SDL_ProgressState
type ProgressState int32

const (
	ProgressStateInvalid       ProgressState = iota - 1 // An invalid progress state indicating an error; check [GetError].
	ProgressStateNone                                   // No progress bar is shown.
	ProgressStateIndeterminate                          // The progress bar is shown in a indeterminate state.
	ProgressStateNormal                                 // The progress bar is shown in a normal state.
	ProgressStatePaused                                 // The progress bar is shown in a paused state.
	ProgressStateError                                  // The progress bar is shown in a state indicating the application had an error.
)

// [DisplayModeData] internal display mode data.
//
// [DisplayModeData]: https://wiki.libsdl.org/SDL3/SDL_DisplayModeData
type DisplayModeData struct{}

// [DisplayMode] defines a display mode.
//
// [DisplayMode]: https://wiki.libsdl.org/SDL3/SDL_DisplayMode
type DisplayMode struct {
	DisplayID              DisplayID        // The display this mode is associated with
	Format                 PixelFormat      // Pixel format
	W                      int32            // Width
	H                      int32            // Height
	PixelDensity           float32          // Scale converting size to pixels (e.g. a 1920x1080 mode with 2.0 scale would have 3840x2160 pixels)
	RefreshRate            float32          // Refresh rate (or 0.0f for unspecified)
	RefreshRateNumerator   int32            // Precise refresh rate numerator (or 0 for unspecified)
	RefreshRateDenominator int32            // Precise refresh rate denominator
	Internal               *DisplayModeData // Private
}

// [DisplayOrientation] describes the way a display is rotated.
//
// [DisplayOrientation]: https://wiki.libsdl.org/SDL3/SDL_DisplayOrientation
type DisplayOrientation uint32

const (
	OrientationUnknown          DisplayOrientation = iota // The display orientation can't be determined
	OrientationLandscape                                  // The display is in landscape mode, with the right side up, relative to portrait mode
	OrientationLandscapeFlipped                           // The display is in landscape mode, with the left side up, relative to portrait mode
	OrientationPortrait                                   // The display is in portrait mode
	OrientationPortraitFlipped                            // The display is in portrait mode, upside down
)

// [SystemTheme].
//
// [SystemTheme]: https://wiki.libsdl.org/SDL3/SDL_SystemTheme
type SystemTheme uint32

const (
	SystemThemeUnknown SystemTheme = iota // Unknown system theme
	SystemThemeLight                      // Light colored system theme
	SystemThemeDark                       // Dark colored system theme
)

// [Window] is used as an opaque handle to a window.
//
// [Window]: https://wiki.libsdl.org/SDL3/SDL_Window
type Window struct{}

// [WindowFlags] the flags on a window.
//
// [WindowFlags]: https://wiki.libsdl.org/SDL3/SDL_WindowFlags
type WindowFlags uint64

const (
	WindowFullscreen        WindowFlags = 0x0000000000000001 // Window is in fullscreen mode
	WindowOpenGL            WindowFlags = 0x0000000000000002 // Window usable with OpenGL context
	WindowOccluded          WindowFlags = 0x0000000000000004 // Window is occluded
	WindowHidden            WindowFlags = 0x0000000000000008 // Window is neither mapped onto the desktop nor shown in the taskbar/dock/window list; SDL_ShowWindow() is required for it to become visible
	WindowBorderless        WindowFlags = 0x0000000000000010 // No window decoration
	WindowResizable         WindowFlags = 0x0000000000000020 // Window can be resized
	WindowMinimized         WindowFlags = 0x0000000000000040 // Window is minimized
	WindowMaximized         WindowFlags = 0x0000000000000080 // Window is maximized
	WindowMouseGrabbed      WindowFlags = 0x0000000000000100 // Window has grabbed mouse input
	WindowInputFocus        WindowFlags = 0x0000000000000200 // Window has input focus
	WindowMouseFocus        WindowFlags = 0x0000000000000400 // Window has mouse focus
	WindowExternal          WindowFlags = 0x0000000000000800 // Window not created by SDL
	WindowModal             WindowFlags = 0x0000000000001000 // Window is modal
	WindowHighPixelDensity  WindowFlags = 0x0000000000002000 // Window uses high pixel density back buffer if possible
	WindowMouseCapture      WindowFlags = 0x0000000000004000 // Window has mouse captured (unrelated to MOUSE_GRABBED)
	WindowMouseRelativeMode WindowFlags = 0x0000000000008000 // Window has relative mode enabled
	WindowAlwaysOnTop       WindowFlags = 0x0000000000010000 // Window should always be above others
	WindowUtility           WindowFlags = 0x0000000000020000 // Window should be treated as a utility window, not showing in the task bar and window list
	WindowTooltip           WindowFlags = 0x0000000000040000 // Window should be treated as a tooltip and does not get mouse or keyboard focus, requires a parent window
	WindowPopupMenu         WindowFlags = 0x0000000000080000 // Window should be treated as a popup menu, requires a parent window
	WindowKeyboardGrabbed   WindowFlags = 0x0000000000100000 // Window has grabbed keyboard input
	WindowFillDocument      WindowFlags = 0x0000000000200000 // Window is in fill-document mode (Emscripten only), since SDL 3.4.0
	WindowVulkan            WindowFlags = 0x0000000010000000 // Window usable for Vulkan surface
	WindowMetal             WindowFlags = 0x0000000020000000 // Window usable for Metal view
	WindowTransparent       WindowFlags = 0x0000000040000000 // Window with transparent buffer
	WindowNotFocusable      WindowFlags = 0x0000000080000000 // Window should not be focusable
)

const (
	WindowPosUndefinedMask uint32 = 0x1FFF0000
	WindowPosUndefined            = WindowPosUndefinedMask
	WindowPosCenteredMask  uint32 = 0x2FFF0000
	WindowPosCentered             = WindowPosCenteredMask
)

type GLContextState struct{}

// [GLContext] is an opaque handle to an OpenGL context.
//
// [GLContext]: https://wiki.libsdl.org/SDL3/SDL_GLContext
type GLContext *GLContextState

func WindowPosCenteredDisplay(displayID DisplayID) uint32 {
	return WindowPosCenteredMask | uint32(displayID)
}

// [DestroyWindow] destroys a window.
//
// [DestroyWindow]: https://wiki.libsdl.org/SDL3/SDL_DestroyWindow
func DestroyWindow(window *Window) {
	sdlDestroyWindow(window)
}

// func CreatePopupWindow(parent *Window, offset_x int32, offset_y int32, w int32, h int32, flags WindowFlags) *Window {
//	return sdlCreatePopupWindow(parent, offset_x, offset_y, w, h, flags)
// }

// [CreateWindow] creates a window with the specified dimensions and flags.
//
// [CreateWindow]: https://wiki.libsdl.org/SDL3/SDL_CreateWindow
func CreateWindow(title string, w int32, h int32, flags WindowFlags) *Window {
	return sdlCreateWindow(title, w, h, flags)
}

// [CreateWindowWithProperties] creates a window with the specified properties.
//
// [CreateWindowWithProperties]: https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties
func CreateWindowWithProperties(props PropertiesID) *Window {
	return sdlCreateWindowWithProperties(props)
}

// [DestroyWindowSurface] destroys the surface associated with the window.
//
// [DestroyWindowSurface]: https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface
func DestroyWindowSurface(window *Window) bool {
	return sdlDestroyWindowSurface(window)
}

// [DisableScreenSaver] prevents the screen from being blanked by a screen saver.
//
// [DisableScreenSaver]: https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver
func DisableScreenSaver() bool {
	return sdlDisableScreenSaver()
}

// func EGL_GetCurrentConfig() EGLConfig {
//	return sdlEGL_GetCurrentConfig()
// }

// func EGL_GetCurrentDisplay() EGLDisplay {
//	return sdlEGL_GetCurrentDisplay()
// }

// func EGL_GetProcAddress(proc string) FunctionPointer {
//	return sdlEGL_GetProcAddress(proc)
// }

// func EGL_GetWindowSurface(window *Window) EGLSurface {
//	return sdlEGL_GetWindowSurface(window)
// }

// func EGL_SetAttributeCallbacks(platformAttribCallback EGLAttribArrayCallback, surfaceAttribCallback EGLIntArrayCallback, contextAttribCallback EGLIntArrayCallback, userdata unsafe.Pointer)  {
//	sdlEGL_SetAttributeCallbacks(platformAttribCallback, surfaceAttribCallback, contextAttribCallback, userdata)
// }

// [EnableScreenSaver] allows the screen to be blanked by a screen saver.
//
// [EnableScreenSaver]: https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver
func EnableScreenSaver() bool {
	return sdlEnableScreenSaver()
}

// [FlashWindow] requests a window to demand attention from the user.
//
// [FlashWindow]: https://wiki.libsdl.org/SDL3/SDL_FlashWindow
func FlashWindow(window *Window, operation FlashOperation) bool {
	return sdlFlashWindow(window, operation)
}

// [SetWindowProgressState] sets the state of the progress bar for the given window’s taskbar icon.
//
// Available since SDL 3.4.0.
//
// [SetWindowProgressState]: https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressState
func SetWindowProgressState(window *Window, state ProgressState) bool {
	return sdlSetWindowProgressState(window, state)
}

// [GetWindowProgressState] gets the state of the progress bar for the given window’s taskbar icon.
//
// Available since SDL 3.4.0.
//
// [GetWindowProgressState]: https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressState
func GetWindowProgressState(window *Window) ProgressState {
	return sdlGetWindowProgressState(window)
}

// [SetWindowProgressValue] sets the value of the progress bar for the given window’s taskbar icon.
//
// Available since SDL 3.4.0.
//
// [SetWindowProgressValue]: https://wiki.libsdl.org/SDL3/SDL_SetWindowProgressValue
func SetWindowProgressValue(window *Window, value float32) bool {
	return sdlSetWindowProgressValue(window, value)
}

// [GetWindowProgressValue] gets the value of the progress bar for the given window’s taskbar icon.
//
// Available since SDL 3.4.0.
//
// [GetWindowProgressValue]: https://wiki.libsdl.org/SDL3/SDL_GetWindowProgressValue
func GetWindowProgressValue(window *Window) float32 {
	return sdlGetWindowProgressValue(window)
}

// [GetClosestFullscreenDisplayMode] gets the closest match to the requested display mode.
//
// [GetClosestFullscreenDisplayMode]: https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode
func GetClosestFullscreenDisplayMode(displayID DisplayID, w int32, h int32, refreshRate float32, includeHighDensityModes bool, closest *DisplayMode) bool {
	return sdlGetClosestFullscreenDisplayMode(displayID, w, h, refreshRate, includeHighDensityModes, closest)
}

// [GetCurrentDisplayMode] gets information about the current display mode.
//
// [GetCurrentDisplayMode]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode
func GetCurrentDisplayMode(displayID DisplayID) *DisplayMode {
	return sdlGetCurrentDisplayMode(displayID)
}

// [GetCurrentDisplayOrientation] gets the orientation of a display.
//
// [GetCurrentDisplayOrientation]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation
func GetCurrentDisplayOrientation(displayID DisplayID) DisplayOrientation {
	return sdlGetCurrentDisplayOrientation(displayID)
}

// [GetCurrentVideoDriver] gets the name of the currently initialized video driver.
//
// [GetCurrentVideoDriver]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver
func GetCurrentVideoDriver() string {
	return sdlGetCurrentVideoDriver()
}

// [GetDesktopDisplayMode] gets information about the desktop's display mode.
//
// [GetDesktopDisplayMode]: https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode
func GetDesktopDisplayMode(displayID DisplayID) *DisplayMode {
	return sdlGetDesktopDisplayMode(displayID)
}

// [GetDisplayBounds] gets the desktop area represented by a display.
//
// [GetDisplayBounds]: https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds
func GetDisplayBounds(displayID DisplayID, rect *Rect) bool {
	return sdlGetDisplayBounds(displayID, rect)
}

// [GetDisplayContentScale] gets the content scale of a display.
//
// [GetDisplayContentScale]: https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale
func GetDisplayContentScale(displayID DisplayID) float32 {
	return sdlGetDisplayContentScale(displayID)
}

// [GetDisplayForPoint] gets the display containing a point.
//
// [GetDisplayForPoint]: https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint
func GetDisplayForPoint(point *Point) DisplayID {
	return sdlGetDisplayForPoint(point)
}

// [GetDisplayForRect] gets the display primarily containing a rect.
//
// [GetDisplayForRect]: https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect
func GetDisplayForRect(rect *Rect) DisplayID {
	return sdlGetDisplayForRect(rect)
}

// [GetDisplayForWindow] gets the display associated with a window.
//
// [GetDisplayForWindow]: https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow
func GetDisplayForWindow(window *Window) DisplayID {
	return sdlGetDisplayForWindow(window)
}

// [GetDisplayName] gets the name of a display in UTF-8 encoding.
//
// [GetDisplayName]: https://wiki.libsdl.org/SDL3/SDL_GetDisplayName
func GetDisplayName(displayID DisplayID) string {
	return sdlGetDisplayName(displayID)
}

// [GetDisplayProperties] gets the properties associated with a display.
//
// [GetDisplayProperties]: https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties
func GetDisplayProperties(displayID DisplayID) PropertiesID {
	return sdlGetDisplayProperties(displayID)
}

// [GetDisplays] gets a list of currently connected displays.
//
// [GetDisplays]: https://wiki.libsdl.org/SDL3/SDL_GetDisplays
func GetDisplays() []DisplayID {
	var count int32
	displays := sdlGetDisplays(&count)
	defer Free(unsafe.Pointer(displays))
	return mem.Copy(displays, count)
}

// [GetDisplayUsableBounds] gets the usable desktop area represented by a display, in screen coordinates.
//
// [GetDisplayUsableBounds]: https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds
func GetDisplayUsableBounds(displayID DisplayID, rect *Rect) bool {
	return sdlGetDisplayUsableBounds(displayID, rect)
}

// [GetFullscreenDisplayModes] gets a list of fullscreen display modes available on a display, or nil on error.
//
// [GetFullscreenDisplayModes]: https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes
func GetFullscreenDisplayModes(displayID DisplayID) []*DisplayMode {
	var count int32
	displayModes := sdlGetFullscreenDisplayModes(displayID, &count)
	defer Free(unsafe.Pointer(displayModes))
	return mem.DeepCopy(displayModes, count)
}

// [GetGrabbedWindow] gets the window that currently has an input grab enabled.
//
// [GetGrabbedWindow]: https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow
func GetGrabbedWindow() *Window {
	return sdlGetGrabbedWindow()
}

// [GetNaturalDisplayOrientation] gets the orientation of a display when it is unrotated.
//
// [GetNaturalDisplayOrientation]: https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation
func GetNaturalDisplayOrientation(displayID DisplayID) DisplayOrientation {
	return sdlGetNaturalDisplayOrientation(displayID)
}

// [GetNumVideoDrivers] gets the number of video drivers compiled into SDL.
//
// [GetNumVideoDrivers]: https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers
func GetNumVideoDrivers() int32 {
	return sdlGetNumVideoDrivers()
}

// [GetPrimaryDisplay] returns the primary display.
//
// [GetPrimaryDisplay]: https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay
func GetPrimaryDisplay() DisplayID {
	return sdlGetPrimaryDisplay()
}

// [GetSystemTheme] gets the current system theme.
//
// [GetSystemTheme]: https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme
func GetSystemTheme() SystemTheme {
	return sdlGetSystemTheme()
}

// [GetVideoDriver] gets the name of a built in video driver.
//
// [GetVideoDriver]: https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver
func GetVideoDriver(index int32) string {
	return sdlGetVideoDriver(index)
}

// [GetWindowAspectRatio] gets the aspect ratio of a window's client area.
//
// [GetWindowAspectRatio]: https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio
func GetWindowAspectRatio(window *Window, minAspect, maxAspect *float32) bool {
	return sdlGetWindowAspectRatio(window, minAspect, maxAspect)
}

// [GetWindowBordersSize] gets the size of a window's borders (decorations) around the client area.
//
// [GetWindowBordersSize]: https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize
func GetWindowBordersSize(window *Window, top, left, bottom, right *int32) bool {
	return sdlGetWindowBordersSize(window, top, left, bottom, right)
}

// [GetWindowDisplayScale] gets the content display scale relative to a window's pixel size.
//
// [GetWindowDisplayScale]: https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale
func GetWindowDisplayScale(window *Window) float32 {
	return sdlGetWindowDisplayScale(window)
}

// [GetWindowFlags] returns a mask of the [WindowFlags] associated with window.
//
// Example:
//
//	// true if the window has the "WindowResizable" flag
//	if sdl.GetWindowFlags(window)&sdl.WindowResizable == sdl.WindowResizable {
//		// ...
//	}
//
// [GetWindowFlags]: https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags
func GetWindowFlags(window *Window) WindowFlags {
	return sdlGetWindowFlags(window)
}

// [GetWindowFromID] gets a window from a stored ID.
//
// [GetWindowFromID]: https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID
func GetWindowFromID(id WindowID) *Window {
	return sdlGetWindowFromID(id)
}

// [GetWindowFullscreenMode] queries the display mode to use when a window is visible at fullscreen.
//
// [GetWindowFullscreenMode]: https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode
func GetWindowFullscreenMode(window *Window) *DisplayMode {
	return sdlGetWindowFullscreenMode(window)
}

// func GetWindowICCProfile(window *Window, size *uint64) unsafe.Pointer {
//	return sdlGetWindowICCProfile(window, size)
// }

// [GetWindowID] returns the ID of the window on success or 0 on failure.
//
// [GetWindowID]: https://wiki.libsdl.org/SDL3/SDL_GetWindowID
func GetWindowID(window *Window) WindowID {
	return sdlGetWindowID(window)
}

// [GetWindowKeyboardGrab] returns true if keyboard is grabbed, and false otherwise.
//
// [GetWindowKeyboardGrab]: https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab
func GetWindowKeyboardGrab(window *Window) bool {
	return sdlGetWindowKeyboardGrab(window)
}

// [GetWindowMaximumSize] gets the maximum size of a window's client area.
//
// [GetWindowMaximumSize]: https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize
func GetWindowMaximumSize(window *Window, w, h *int32) bool {
	return sdlGetWindowMaximumSize(window, w, h)
}

// [GetWindowMinimumSize] gets the minimum size of a window's client area.
//
// [GetWindowMinimumSize]: https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize
func GetWindowMinimumSize(window *Window, w, h *int32) bool {
	return sdlGetWindowMinimumSize(window, w, h)
}

// [GetWindowMouseGrab] returns true if mouse is grabbed, and false otherwise.
//
// [GetWindowMouseGrab]: https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab
func GetWindowMouseGrab(window *Window) bool {
	return sdlGetWindowMouseGrab(window)
}

// [GetWindowMouseRect] gets the mouse confinement rectangle of a window.
//
// [GetWindowMouseRect]: https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect
func GetWindowMouseRect(window *Window) *Rect {
	return sdlGetWindowMouseRect(window)
}

// [GetWindowOpacity] gets the opacity of a window.
//
// [GetWindowOpacity]: https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity
func GetWindowOpacity(window *Window) float32 {
	return sdlGetWindowOpacity(window)
}

// [GetWindowParent] gets parent of a window.
//
// [GetWindowParent]: https://wiki.libsdl.org/SDL3/SDL_GetWindowParent
func GetWindowParent(window *Window) *Window {
	return sdlGetWindowParent(window)
}

// [GetWindowPixelDensity] gets the pixel density of a window.
//
// [GetWindowPixelDensity]: https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity
func GetWindowPixelDensity(window *Window) float32 {
	return sdlGetWindowPixelDensity(window)
}

// [GetWindowPixelFormat] gets the pixel format associated with the window.
//
// [GetWindowPixelFormat]: https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat
func GetWindowPixelFormat(window *Window) PixelFormat {
	return sdlGetWindowPixelFormat(window)
}

// [GetWindowPosition] gets the position of a window.
//
// [GetWindowPosition]: https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition
func GetWindowPosition(window *Window, x *int32, y *int32) bool {
	return sdlGetWindowPosition(window, x, y)
}

// [GetWindowProperties] gets the properties associated with a window.
//
// [GetWindowProperties]: https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties
func GetWindowProperties(window *Window) PropertiesID {
	return sdlGetWindowProperties(window)
}

// [GetWindows] gets a list of valid windows.
//
// [GetWindows]: https://wiki.libsdl.org/SDL3/SDL_GetWindows
func GetWindows() []*Window {
	var count int32
	windows := sdlGetWindows(&count)
	defer Free(unsafe.Pointer(windows))
	return mem.Copy(windows, count)
}

// [GetWindowSafeArea] gets the safe area for this window.
//
// [GetWindowSafeArea]: https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea
func GetWindowSafeArea(window *Window, rect *Rect) bool {
	return sdlGetWindowSafeArea(window, rect)
}

// [GetWindowSize] gets the size of a window's client area.
//
// [GetWindowSize]: https://wiki.libsdl.org/SDL3/SDL_GetWindowSize
func GetWindowSize(window *Window, w *int32, h *int32) bool {
	return sdlGetWindowSize(window, w, h)
}

// [GetWindowSizeInPixels] gets the size of a window's client area, in pixels.
//
// [GetWindowSizeInPixels]: https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels
func GetWindowSizeInPixels(window *Window, w *int32, h *int32) bool {
	return sdlGetWindowSizeInPixels(window, w, h)
}

// [GetWindowSurface] gets the SDL surface associated with the window.
//
// [GetWindowSurface]: https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface
func GetWindowSurface(window *Window) *Surface {
	return sdlGetWindowSurface(window)
}

// [GetWindowSurfaceVSync] gets VSync for the window surface.
//
// [GetWindowSurfaceVSync]: https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync
func GetWindowSurfaceVSync(window *Window, vsync *int32) bool {
	return sdlGetWindowSurfaceVSync(window, vsync)
}

// [GetWindowTitle] gets the title of a window.
//
// [GetWindowTitle]: https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle
func GetWindowTitle(window *Window) string {
	return sdlGetWindowTitle(window)
}

// [GLCreateContext] creates an OpenGL context for an OpenGL window, and makes it current.
//
// [GLCreateContext]: https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext
func GLCreateContext(window *Window) GLContext {
	return sdlGLCreateContext(window)
}

// [GLDestroyContext] deletes an OpenGL context.
//
// [GLDestroyContext]: https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext
func GLDestroyContext(context GLContext) bool {
	return sdlGLDestroyContext(context)
}

// func GL_ExtensionSupported(extension string) bool {
//	return sdlGL_ExtensionSupported(extension)
// }

// [GLGetAttribute] gets the actual value for an attribute from the current context.
//
// [GLGetAttribute]: https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute
func GLGetAttribute(attr GLAttr, value *int32) bool {
	return sdlGLGetAttribute(attr, value)
}

// [GLGetCurrentContext] returns the currently active OpenGL context or nil on failure.
//
// [GLGetCurrentContext]: https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext
func GLGetCurrentContext() GLContext {
	return sdlGLGetCurrentContext()
}

// [GLGetCurrentWindow] returns the currently active OpenGL window on success or nil on failure.
//
// [GLGetCurrentWindow]: https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow
func GLGetCurrentWindow() *Window {
	return sdlGLGetCurrentWindow()
}

// func GL_GetProcAddress(proc string) FunctionPointer {
//	return sdlGL_GetProcAddress(proc)
// }

// func GL_GetSwapInterval(interval *int32) bool {
//	return sdlGL_GetSwapInterval(interval)
// }

// func GL_LoadLibrary(path string) bool {
//	return sdlGL_LoadLibrary(path)
// }

// func GL_MakeCurrent(window *Window, context GLContext) bool {
//	return sdlGL_MakeCurrent(window, context)
// }

// func GL_ResetAttributes()  {
//	sdlGL_ResetAttributes()
// }

// [GLSetAttribute] sets the OpenGL attribute attr to value. The requested attributes should be set before creating an OpenGL window.
//
// You should use [GLGetAttribute] to check the values after creating the OpenGL context,
// since the values obtained can differ from the requested ones.
//
// [GLSetAttribute]: https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute
func GLSetAttribute(attr GLAttr, value int32) bool {
	return sdlGLSetAttribute(attr, value)
}

// [GLSetSwapInterval] sets the swap interval for the current OpenGL context.
//
// [GLSetSwapInterval]: https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval
func GLSetSwapInterval(interval int32) bool {
	ret, _, _ := purego.SyscallN(sdlGLSetSwapInterval, uintptr(interval))
	return byte(ret) != 0
}

// [GLSwapWindow] updates a window with OpenGL rendering.
//
// [GLSwapWindow]: https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow
func GLSwapWindow(window *Window) bool {
	ret, _, _ := purego.SyscallN(sdlGLSwapWindow, uintptr(unsafe.Pointer(window)))
	return byte(ret) != 0
}

// func GL_UnloadLibrary()  {
//	sdlGL_UnloadLibrary()
// }

// [HideWindow] hides a window.
//
// [HideWindow]: https://wiki.libsdl.org/SDL3/SDL_HideWindow
func HideWindow(window *Window) bool {
	return sdlHideWindow(window)
}

// [MaximizeWindow] requests that the window be made as large as possible.
//
// [MaximizeWindow]: https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow
func MaximizeWindow(window *Window) bool {
	return sdlMaximizeWindow(window)
}

// [MinimizeWindow] requests that the window be minimized to an iconic representation.
//
// [MinimizeWindow]: https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow
func MinimizeWindow(window *Window) bool {
	return sdlMinimizeWindow(window)
}

// [RaiseWindow] requests that a window be raised above other windows and gain the input focus.
//
// [RaiseWindow]: https://wiki.libsdl.org/SDL3/SDL_RaiseWindow
func RaiseWindow(window *Window) bool {
	return sdlRaiseWindow(window)
}

// [RestoreWindow] requests that the size and position of a minimized or maximized window be restored.
//
// [RestoreWindow]: https://wiki.libsdl.org/SDL3/SDL_RestoreWindow
func RestoreWindow(window *Window) bool {
	return sdlRestoreWindow(window)
}

// [ScreenSaverEnabled] checks whether the screensaver is currently enabled.
//
// [ScreenSaverEnabled]: https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled
func ScreenSaverEnabled() bool {
	return sdlScreenSaverEnabled()
}

// [SetWindowAlwaysOnTop] sets the window to always be above the others.
//
// [SetWindowAlwaysOnTop]: https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop
func SetWindowAlwaysOnTop(window *Window, onTop bool) bool {
	return sdlSetWindowAlwaysOnTop(window, onTop)
}

// [SetWindowFillDocument] sets the window to fill the current document space (Emscripten only).
//
// Available since SDL 3.4.0.
//
// [SetWindowFillDocument]: https://wiki.libsdl.org/SDL3/SDL_SetWindowFillDocument
// func SetWindowFillDocument(window *Window, fill bool) bool {
// 	return sdlSetWindowFillDocument(window, fill)
// }

// [SetWindowAspectRatio] requests that the aspect ratio of a window's client area be set.
//
// [SetWindowAspectRatio]: https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio
func SetWindowAspectRatio(window *Window, minAspect, maxAspect float32) bool {
	return sdlSetWindowAspectRatio(window, minAspect, maxAspect)
}

// [SetWindowBordered] sets the border state of a window.
//
// [SetWindowBordered]: https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered
func SetWindowBordered(window *Window, bordered bool) bool {
	return sdlSetWindowBordered(window, bordered)
}

// [SetWindowFocusable] sets whether the window may have input focus.
//
// [SetWindowFocusable]: https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable
func SetWindowFocusable(window *Window, focusable bool) bool {
	return sdlSetWindowFocusable(window, focusable)
}

// [SetWindowFullscreen] requests that the window's fullscreen state be changed.
//
// [SetWindowFullscreen]: https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen
func SetWindowFullscreen(window *Window, fullscreen bool) bool {
	return sdlSetWindowFullscreen(window, fullscreen)
}

// [SetWindowFullscreenMode] sets the display mode to use when a window is visible and fullscreen.
//
// [SetWindowFullscreenMode]: https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode
func SetWindowFullscreenMode(window *Window, mode *DisplayMode) bool {
	return sdlSetWindowFullscreenMode(window, mode)
}

// [SetWindowHitTest] provide a callback that decides if a window region has special properties.
//
// [SetWindowHitTest]: https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest
func SetWindowHitTest(window *Window, callback HitTest, callbackData unsafe.Pointer) bool {
	wrapper := func(win *Window, point *Point, data unsafe.Pointer) uintptr {
		return uintptr(callback(win, point, data))
	}

	return sdlSetWindowHitTest(window, wrapper, callbackData)
}

// [SetWindowIcon] sets the icon for a window.
//
// [SetWindowIcon]: https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon
func SetWindowIcon(window *Window, icon *Surface) bool {
	return sdlSetWindowIcon(window, icon)
}

// [SetWindowKeyboardGrab] enables capture of system keyboard shortcuts like Alt+Tab or the Meta/Super key.
// Note that not all system keyboard shortcuts can be captured by applications (one example is Ctrl+Alt+Del on Windows).
//
// [SetWindowKeyboardGrab]: https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab
func SetWindowKeyboardGrab(window *Window, grabbed bool) bool {
	return sdlSetWindowKeyboardGrab(window, grabbed)
}

// [SetWindowMaximumSize] sets the maximum size of a window's client area.
//
// [SetWindowMaximumSize]: https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize
func SetWindowMaximumSize(window *Window, maxW, maxH int32) bool {
	return sdlSetWindowMaximumSize(window, maxW, maxH)
}

// [SetWindowMinimumSize] sets the minimum size of a window's client area.
//
// [SetWindowMinimumSize]: https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize
func SetWindowMinimumSize(window *Window, minW, minH int32) bool {
	return sdlSetWindowMinimumSize(window, minW, minH)
}

// [SetWindowModal] toggles the state of the window as modal.
//
// [SetWindowModal]: https://wiki.libsdl.org/SDL3/SDL_SetWindowModal
func SetWindowModal(window *Window, modal bool) bool {
	return sdlSetWindowModal(window, modal)
}

// [SetWindowMouseGrab] enables restriction of the mouse cursor to the window.
//
// [SetWindowMouseGrab]: https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab
func SetWindowMouseGrab(window *Window, grabbed bool) bool {
	return sdlSetWindowMouseGrab(window, grabbed)
}

// [SetWindowMouseRect] confines the cursor to the specified area of a window.
//
// [SetWindowMouseRect]: https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect
func SetWindowMouseRect(window *Window, rect *Rect) bool {
	return sdlSetWindowMouseRect(window, rect)
}

// [SetWindowOpacity] sets the opacity for a window.
//
// [SetWindowOpacity]: https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity
func SetWindowOpacity(window *Window, opacity float32) bool {
	return sdlSetWindowOpacity(window, opacity)
}

// [SetWindowParent] sets the window as a child of a parent window.
//
// [SetWindowParent]: https://wiki.libsdl.org/SDL3/SDL_SetWindowParent
func SetWindowParent(window *Window, parent *Window) bool {
	return sdlSetWindowParent(window, parent)
}

// [SetWindowPosition] requests that the window's position be set.
//
// [SetWindowPosition]: https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition
func SetWindowPosition(window *Window, x int32, y int32) bool {
	return sdlSetWindowPosition(window, x, y)
}

// [SetWindowResizable] sets the user-resizable state of a window.
//
// [SetWindowResizable]: https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable
func SetWindowResizable(window *Window, resizable bool) bool {
	return sdlSetWindowResizable(window, resizable)
}

// func SetWindowShape(window *Window, shape *Surface) bool {
//	return sdlSetWindowShape(window, shape)
// }

// [SetWindowSize] requests that the size of a window's client area be set.
//
// [SetWindowSize]: https://wiki.libsdl.org/SDL3/SDL_SetWindowSize
func SetWindowSize(window *Window, w int32, h int32) bool {
	return sdlSetWindowSize(window, w, h)
}

// [SetWindowSurfaceVSync] toggles VSync for the window surface.
//
// [SetWindowSurfaceVSync]: https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync
func SetWindowSurfaceVSync(window *Window, vsync int32) bool {
	return sdlSetWindowSurfaceVSync(window, vsync)
}

// [SetWindowTitle] requests that the size of a window's client area be set.
//
// [SetWindowTitle]: https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle
func SetWindowTitle(window *Window, title string) bool {
	return sdlSetWindowTitle(window, title)
}

// [ShowWindow] shows a window.
//
// [ShowWindow]: https://wiki.libsdl.org/SDL3/SDL_ShowWindow
func ShowWindow(window *Window) bool {
	return sdlShowWindow(window)
}

// [ShowWindowSystemMenu] displays the system-level window menu.
//
// [ShowWindowSystemMenu]: https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu
func ShowWindowSystemMenu(window *Window, x, y int32) bool {
	return sdlShowWindowSystemMenu(window, x, y)
}

// [SyncWindow] blocks until any pending window state is finalized.
//
// [SyncWindow]: https://wiki.libsdl.org/SDL3/SDL_SyncWindow
func SyncWindow(window *Window) bool {
	return sdlSyncWindow(window)
}

// [UpdateWindowSurface] copies the window surface to the screen.
//
// [UpdateWindowSurface]: https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface
func UpdateWindowSurface(window *Window) bool {
	return sdlUpdateWindowSurface(window)
}

// func UpdateWindowSurfaceRects(window *Window, rects *Rect, numrects int32) bool {
//	return sdlUpdateWindowSurfaceRects(window, rects, numrects)
// }

// [WindowHasSurface] returns whether the window has a surface associated with it.
//
// [WindowHasSurface]: https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface
func WindowHasSurface(window *Window) bool {
	return sdlWindowHasSurface(window)
}
