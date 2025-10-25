package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

// [HintPriority] is an enumeration of hint priorities.
//
// [HintPriority]: https://wiki.libsdl.org/SDL3/SDL_HintPriority
type HintPriority uint32

const (
	HintDefault HintPriority = iota
	HintNormal
	HintOverride
)

const (
	HintAllowAltTabWhileGrabbed            = "SDL_ALLOW_ALT_TAB_WHILE_GRABBED"
	HintAndroidAllowRecreateActivity       = "SDL_ANDROID_ALLOW_RECREATE_ACTIVITY"
	HintAndroidBlockOnPause                = "SDL_ANDROID_BLOCK_ON_PAUSE"
	HintAndroidLowLatencyAudio             = "SDL_ANDROID_LOW_LATENCY_AUDIO"
	HintAndroidTrapBackButton              = "SDL_ANDROID_TRAP_BACK_BUTTON"
	HintAppID                              = "SDL_APP_ID"
	HintAppName                            = "SDL_APP_NAME"
	HintAppleTVControllerUiEvents          = "SDL_APPLE_TV_CONTROLLER_UI_EVENTS"
	HintAppleTVRemoteAllowRotation         = "SDL_APPLE_TV_REMOTE_ALLOW_ROTATION"
	HintAudioAlsaDefaultDevice             = "SDL_AUDIO_ALSA_DEFAULT_DEVICE"
	HintAudioAlsaDefaultPlaybackDevice     = "SDL_AUDIO_ALSA_DEFAULT_PLAYBACK_DEVICE"
	HintAudioAlsaDefaultRecordingDevice    = "SDL_AUDIO_ALSA_DEFAULT_RECORDING_DEVICE"
	HintAudioCategory                      = "SDL_AUDIO_CATEGORY"
	HintAudioChannels                      = "SDL_AUDIO_CHANNELS"
	HintAudioDeviceAppIconName             = "SDL_AUDIO_DEVICE_APP_ICON_NAME"
	HintAudioDeviceSampleFrames            = "SDL_AUDIO_DEVICE_SAMPLE_FRAMES"
	HintAudioDeviceStreamName              = "SDL_AUDIO_DEVICE_STREAM_NAME"
	HintAudioDeviceStreamRole              = "SDL_AUDIO_DEVICE_STREAM_ROLE"
	HintAudioDiskInputFile                 = "SDL_AUDIO_DISK_INPUT_FILE"
	HintAudioDiskOutputFile                = "SDL_AUDIO_DISK_OUTPUT_FILE"
	HintAudioDiskTimescale                 = "SDL_AUDIO_DISK_TIMESCALE"
	HintAudioDriver                        = "SDL_AUDIO_DRIVER"
	HintAudioDummyTimescale                = "SDL_AUDIO_DUMMY_TIMESCALE"
	HintAudioFormat                        = "SDL_AUDIO_FORMAT"
	HintAudioFrequency                     = "SDL_AUDIO_FREQUENCY"
	HintAudioIncludeMonitors               = "SDL_AUDIO_INCLUDE_MONITORS"
	HintAutoUpdateJoysticks                = "SDL_AUTO_UPDATE_JOYSTICKS"
	HintAutoUpdateSensors                  = "SDL_AUTO_UPDATE_SENSORS"
	HintBmpSaveLegacyFormat                = "SDL_BMP_SAVE_LEGACY_FORMAT"
	HintCameraDriver                       = "SDL_CAMERA_DRIVER"
	HintCpuFeatureMask                     = "SDL_CPU_FEATURE_MASK"
	HintJoystickDirectInput                = "SDL_JOYSTICK_DIRECTINPUT"
	HintFileDialogDriver                   = "SDL_FILE_DIALOG_DRIVER"
	HintDisplayUsableBounds                = "SDL_DISPLAY_USABLE_BOUNDS"
	HintEmscriptenAsyncify                 = "SDL_EMSCRIPTEN_ASYNCIFY"
	HintEmscriptenCanvasSelector           = "SDL_EMSCRIPTEN_CANVAS_SELECTOR"
	HintEmscriptenKeyboardElement          = "SDL_EMSCRIPTEN_KEYBOARD_ELEMENT"
	HintEnableScreenKeyboard               = "SDL_ENABLE_SCREEN_KEYBOARD"
	HintEvdevDevices                       = "SDL_EVDEV_DEVICES"
	HintEventLogging                       = "SDL_EVENT_LOGGING"
	HintForceRaiseWindow                   = "SDL_FORCE_RAISEWINDOW"
	HintFramebufferAcceleration            = "SDL_FRAMEBUFFER_ACCELERATION"
	HintGamecontrollerConfig               = "SDL_GAMECONTROLLERCONFIG"
	HintGamecontrollerConfigFile           = "SDL_GAMECONTROLLERCONFIG_FILE"
	HintGamecontrollerType                 = "SDL_GAMECONTROLLERTYPE"
	HintGamecontrollerIgnoreDevices        = "SDL_GAMECONTROLLER_IGNORE_DEVICES"
	HintGamecontrollerIgnoreDevicesExcept  = "SDL_GAMECONTROLLER_IGNORE_DEVICES_EXCEPT"
	HintGamecontrollerSensorFusion         = "SDL_GAMECONTROLLER_SENSOR_FUSION"
	HintGDKTextInputDefaultText            = "SDL_GDK_TEXTINPUT_DEFAULT_TEXT"
	HintGDKTextInputDescription            = "SDL_GDK_TEXTINPUT_DESCRIPTION"
	HintGDKTextInputMaxLength              = "SDL_GDK_TEXTINPUT_MAX_LENGTH"
	HintGDKTextInputScope                  = "SDL_GDK_TEXTINPUT_SCOPE"
	HintGDKTextInputTitle                  = "SDL_GDK_TEXTINPUT_TITLE"
	HintHidApiLibUSB                       = "SDL_HIDAPI_LIBUSB"
	HintHidApiLibUSBWhitelist              = "SDL_HIDAPI_LIBUSB_WHITELIST"
	HintHidApiUdev                         = "SDL_HIDAPI_UDEV"
	HintGPUDriver                          = "SDL_GPU_DRIVER"
	HintHidApiEnumerateOnlyControllers     = "SDL_HIDAPI_ENUMERATE_ONLY_CONTROLLERS"
	HintHidApiIgnoreDevices                = "SDL_HIDAPI_IGNORE_DEVICES"
	HintImeImplementedUi                   = "SDL_IME_IMPLEMENTED_UI"
	HintIOSHideHomeIndicator               = "SDL_IOS_HIDE_HOME_INDICATOR"
	HintJoystickAllowBackgroundEvents      = "SDL_JOYSTICK_ALLOW_BACKGROUND_EVENTS"
	HintJoystickArcadeStickDevices         = "SDL_JOYSTICK_ARCADESTICK_DEVICES"
	HintJoystickArcadeStickDevicesExcluded = "SDL_JOYSTICK_ARCADESTICK_DEVICES_EXCLUDED"
	HintJoystickBlacklistDevices           = "SDL_JOYSTICK_BLACKLIST_DEVICES"
	HintJoystickBlacklistDevicesExcluded   = "SDL_JOYSTICK_BLACKLIST_DEVICES_EXCLUDED"
	HintJoystickDevice                     = "SDL_JOYSTICK_DEVICE"
	HintJoystickEnhancedReports            = "SDL_JOYSTICK_ENHANCED_REPORTS"
	HintJoystickFlightStickDevices         = "SDL_JOYSTICK_FLIGHTSTICK_DEVICES"
	HintJoystickFlightStickDevicesExcluded = "SDL_JOYSTICK_FLIGHTSTICK_DEVICES_EXCLUDED"
	HintJoystickGameInput                  = "SDL_JOYSTICK_GAMEINPUT"
	HintJoystickGameCubeDevices            = "SDL_JOYSTICK_GAMECUBE_DEVICES"
	HintJoystickGameCubeDevicesExcluded    = "SDL_JOYSTICK_GAMECUBE_DEVICES_EXCLUDED"
	HintJoystickHidApi                     = "SDL_JOYSTICK_HIDAPI"
	HintJoystickHidApiCombineJoyCons       = "SDL_JOYSTICK_HIDAPI_COMBINE_JOY_CONS"
	HintJoystickHidApiGameCube             = "SDL_JOYSTICK_HIDAPI_GAMECUBE"
	HintJoystickHidApiGameCubeRumbleBrake  = "SDL_JOYSTICK_HIDAPI_GAMECUBE_RUMBLE_BRAKE"
	HintJoystickHidApiJoyCons              = "SDL_JOYSTICK_HIDAPI_JOY_CONS"
	HintJoystickHidApiJoyConHomeLED        = "SDL_JOYSTICK_HIDAPI_JOYCON_HOME_LED"
	HintJoystickHidApiLuna                 = "SDL_JOYSTICK_HIDAPI_LUNA"
	HintJoystickHidApiNintendoClassic      = "SDL_JOYSTICK_HIDAPI_NINTENDO_CLASSIC"
	HintJoystickHidApiPS3                  = "SDL_JOYSTICK_HIDAPI_PS3"
	HintJoystickHidApiPS3SixaxisDriver     = "SDL_JOYSTICK_HIDAPI_PS3_SIXAXIS_DRIVER"
	HintJoystickHidApiPS4                  = "SDL_JOYSTICK_HIDAPI_PS4"
	HintJoystickHidApiPS4ReportInterval    = "SDL_JOYSTICK_HIDAPI_PS4_REPORT_INTERVAL"
	HintJoystickHidApiPS5                  = "SDL_JOYSTICK_HIDAPI_PS5"
	HintJoystickHidApiPS5PlayerLED         = "SDL_JOYSTICK_HIDAPI_PS5_PLAYER_LED"
	HintJoystickHidApiShield               = "SDL_JOYSTICK_HIDAPI_SHIELD"
	HintJoystickHidApiStadia               = "SDL_JOYSTICK_HIDAPI_STADIA"
	HintJoystickHidApiSteam                = "SDL_JOYSTICK_HIDAPI_STEAM"
	HintJoystickHidApiSteamHomeLED         = "SDL_JOYSTICK_HIDAPI_STEAM_HOME_LED"
	HintJoystickHidApiSteamDeck            = "SDL_JOYSTICK_HIDAPI_STEAMDECK"
	HintJoystickHidApiSteamHori            = "SDL_JOYSTICK_HIDAPI_STEAM_HORI"
	HintJoystickHidApiSwitch               = "SDL_JOYSTICK_HIDAPI_SWITCH"
	HintJoystickHidApiSwitchHomeLED        = "SDL_JOYSTICK_HIDAPI_SWITCH_HOME_LED"
	HintJoystickHidApiSwitchPlayerLED      = "SDL_JOYSTICK_HIDAPI_SWITCH_PLAYER_LED"
	HintJoystickHidApiVerticalJoyCons      = "SDL_JOYSTICK_HIDAPI_VERTICAL_JOY_CONS"
	HintJoystickHidApiWii                  = "SDL_JOYSTICK_HIDAPI_WII"
	HintJoystickHidApiWiiPlayerLED         = "SDL_JOYSTICK_HIDAPI_WII_PLAYER_LED"
	HintJoystickHidApiXbox                 = "SDL_JOYSTICK_HIDAPI_XBOX"
	HintJoystickHidApiXbox360              = "SDL_JOYSTICK_HIDAPI_XBOX_360"
	HintJoystickHidApiXbox360PlayerLED     = "SDL_JOYSTICK_HIDAPI_XBOX_360_PLAYER_LED"
	HintJoystickHidApiXbox360Wireless      = "SDL_JOYSTICK_HIDAPI_XBOX_360_WIRELESS"
	HintJoystickHidApiXboxOne              = "SDL_JOYSTICK_HIDAPI_XBOX_ONE"
	HintJoystickHidApiXboxOneHomeLED       = "SDL_JOYSTICK_HIDAPI_XBOX_ONE_HOME_LED"
	HintJoystickIOKit                      = "SDL_JOYSTICK_IOKIT"
	HintJoystickLinuxClassic               = "SDL_JOYSTICK_LINUX_CLASSIC"
	HintJoystickLinuxDeadzones             = "SDL_JOYSTICK_LINUX_DEADZONES"
	HintJoystickLinuxDigitalHats           = "SDL_JOYSTICK_LINUX_DIGITAL_HATS"
	HintJoystickLinuxHatDeadzones          = "SDL_JOYSTICK_LINUX_HAT_DEADZONES"
	HintJoystickMfi                        = "SDL_JOYSTICK_MFI"
	HintJoystickRawinput                   = "SDL_JOYSTICK_RAWINPUT"
	HintJoystickRawinputCorrelateXinput    = "SDL_JOYSTICK_RAWINPUT_CORRELATE_XINPUT"
	HintJoystickRogChakram                 = "SDL_JOYSTICK_ROG_CHAKRAM"
	HintJoystickThread                     = "SDL_JOYSTICK_THREAD"
	HintJoystickThrottleDevices            = "SDL_JOYSTICK_THROTTLE_DEVICES"
	HintJoystickThrottleDevicesExcluded    = "SDL_JOYSTICK_THROTTLE_DEVICES_EXCLUDED"
	HintJoystickWgi                        = "SDL_JOYSTICK_WGI"
	HintJoystickWheelDevices               = "SDL_JOYSTICK_WHEEL_DEVICES"
	HintJoystickWheelDevicesExcluded       = "SDL_JOYSTICK_WHEEL_DEVICES_EXCLUDED"
	HintJoystickZeroCenteredDevices        = "SDL_JOYSTICK_ZERO_CENTERED_DEVICES"
	HintKeycodeOptions                     = "SDL_KEYCODE_OPTIONS"
	HintKmsDrmDeviceIndex                  = "SDL_KMSDRM_DEVICE_INDEX"
	HintKmsDrmRequireDrmMaster             = "SDL_KMSDRM_REQUIRE_DRM_MASTER"
	HintLogging                            = "SDL_LOGGING"
	HintMacBackgroundApp                   = "SDL_MAC_BACKGROUND_APP"
	HintMacCtrlClickEmulateRightClick      = "SDL_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK"
	HintMacOpenGLAsyncDispatch             = "SDL_MAC_OpenGL_ASYNC_DISPATCH"
	HintMacOptionAsAlt                     = "SDL_MAC_OPTION_AS_ALT"
	HintMacScrollMomentum                  = "SDL_MAC_SCROLL_MOMENTUM"
	HintMainCallbackRate                   = "SDL_MAIN_CALLBACK_RATE"
	HintMouseAutoCapture                   = "SDL_MOUSE_AUTO_CAPTURE"
	HintMouseDoubleClickRadius             = "SDL_MOUSE_DOUBLE_CLICK_RADIUS"
	HintMouseDoubleClickTime               = "SDL_MOUSE_DOUBLE_CLICK_TIME"
	HintMouseDefaultSystemCursor           = "SDL_MOUSE_DEFAULT_SYSTEM_CURSOR"
	HintMouseEmulateWarpWithRelative       = "SDL_MOUSE_EMULATE_WARP_WITH_RELATIVE"
	HintMouseFocusClickThrough             = "SDL_MOUSE_FOCUS_CLICKTHROUGH"
	HintMouseNormalSpeedScale              = "SDL_MOUSE_NORMAL_SPEED_SCALE"
	HintMouseRelativeModeCenter            = "SDL_MOUSE_RELATIVE_MODE_CENTER"
	HintMouseRelativeSpeedScale            = "SDL_MOUSE_RELATIVE_SPEED_SCALE"
	HintMouseRelativeSystemScale           = "SDL_MOUSE_RELATIVE_SYSTEM_SCALE"
	HintMouseRelativeWarpMotion            = "SDL_MOUSE_RELATIVE_WARP_MOTION"
	HintMouseRelativeCursorVisible         = "SDL_MOUSE_RELATIVE_CURSOR_VISIBLE"
	HintMouseTouchEvents                   = "SDL_MOUSE_TOUCH_EVENTS"
	HintMuteConsoleKeyboard                = "SDL_MUTE_CONSOLE_KEYBOARD"
	HintNoSignalHandlers                   = "SDL_NO_SIGNAL_HANDLERS"
	HintOpenGLLibrary                      = "SDL_OpenGL_LIBRARY"
	HintEGLLibrary                         = "SDL_EGL_LIBRARY"
	HintOpenGLESDriver                     = "SDL_OpenGL_ES_DRIVER"
	HintOpenVRLibrary                      = "SDL_OPENVR_LIBRARY"
	HintOrientations                       = "SDL_ORIENTATIONS"
	HintPollSentinel                       = "SDL_POLL_SENTINEL"
	HintPreferredLocales                   = "SDL_PREFERRED_LOCALES"
	HintQuitOnLastWindowClose              = "SDL_QUIT_ON_LAST_WINDOW_CLOSE"
	HintRenderDirect3DThreadSafe           = "SDL_RENDER_DIRECT3D_THREADSAFE"
	HintRenderDirect3D11Debug              = "SDL_RENDER_DIRECT3D11_DEBUG"
	HintRenderVulkanDebug                  = "SDL_RENDER_VULKAN_DEBUG"
	HintRenderGPUDebug                     = "SDL_RENDER_GPU_DEBUG"
	HintRenderGPULowPower                  = "SDL_RENDER_GPU_LOW_POWER"
	HintRenderDriver                       = "SDL_RENDER_DRIVER"
	HintRenderLineMethod                   = "SDL_RENDER_LINE_METHOD"
	HintRenderMetalPreferLowPowerDevice    = "SDL_RENDER_METAL_PREFER_LOW_POWER_DEVICE"
	HintRenderVSync                        = "SDL_RENDER_VSYNC"
	HintReturnKeyHidesIme                  = "SDL_RETURN_KEY_HIDES_IME"
	HintRogGamepadMice                     = "SDL_ROG_GAMEPAD_MICE"
	HintRogGamepadMiceExcluded             = "SDL_ROG_GAMEPAD_MICE_EXCLUDED"
	HintRpiVideoLayer                      = "SDL_RPI_VIDEO_LAYER"
	HintScreensaverInhibitActivityName     = "SDL_SCREENSAVER_INHIBIT_ACTIVITY_NAME"
	HintShutdownDBusOnQuit                 = "SDL_SHUTDOWN_DBUS_ON_QUIT"
	HintStorageTitleDriver                 = "SDL_STORAGE_TITLE_DRIVER"
	HintStorageUserDriver                  = "SDL_STORAGE_USER_DRIVER"
	HintThreadForceRealtimeTimeCritical    = "SDL_THREAD_FORCE_REALTIME_TIME_CRITICAL"
	HintThreadPriorityPolicy               = "SDL_THREAD_PRIORITY_POLICY"
	HintTimerResolution                    = "SDL_TIMER_RESOLUTION"
	HintTouchMouseEvents                   = "SDL_TOUCH_MOUSE_EVENTS"
	HintTrackpadIsTouchOnly                = "SDL_TRACKPAD_IS_TOUCH_ONLY"
	HintTVRemoteAsJoystick                 = "SDL_TV_REMOTE_AS_JOYSTICK"
	HintVideoAllowScreensaver              = "SDL_VIDEO_ALLOW_SCREENSAVER"
	HintVideoDisplayPriority               = "SDL_VIDEO_DISPLAY_PRIORITY"
	HintVideoDoubleBuffer                  = "SDL_VIDEO_DOUBLE_BUFFER"
	HintVideoDriver                        = "SDL_VIDEO_DRIVER"
	HintVideoDummySaveFrames               = "SDL_VIDEO_DUMMY_SAVE_FRAMES"
	HintVideoEGLAllowGetDisplayFallback    = "SDL_VIDEO_EGL_ALLOW_GETDISPLAY_FALLBACK"
	HintVideoForceEGL                      = "SDL_VIDEO_FORCE_EGL"
	HintVideoMacFullscreenSpaces           = "SDL_VIDEO_MAC_FULLSCREEN_SPACES"
	HintVideoMacFullscreenMenuVisibility   = "SDL_VIDEO_MAC_FULLSCREEN_MENU_VISIBILITY"
	HintVideoMinimizeOnFocusLoss           = "SDL_VIDEO_MINIMIZE_ON_FOCUS_LOSS"
	HintVideoOffscreenSaveFrames           = "SDL_VIDEO_OFFSCREEN_SAVE_FRAMES"
	HintVideoSyncWindowOperations          = "SDL_VIDEO_SYNC_WINDOW_OPERATIONS"
	HintVideoWaylandAllowLibdecor          = "SDL_VIDEO_WAYLAND_ALLOW_LIBDECOR"
	HintVideoWaylandModeEmulation          = "SDL_VIDEO_WAYLAND_MODE_EMULATION"
	HintVideoWaylandModeScaling            = "SDL_VIDEO_WAYLAND_MODE_SCALING"
	HintVideoWaylandPreferLibdecor         = "SDL_VIDEO_WAYLAND_PREFER_LIBDECOR"
	HintVideoWaylandScaleToDisplay         = "SDL_VIDEO_WAYLAND_SCALE_TO_DISPLAY"
	HintVideoWinD3DCompiler                = "SDL_VIDEO_WIN_D3DCOMPILER"
	HintVideoX11NetWmBypassCompositor      = "SDL_VIDEO_X11_NET_WM_BYPASS_COMPOSITOR"
	HintVideoX11NetWmPing                  = "SDL_VIDEO_X11_NET_WM_PING"
	HintVideoX11NoDirectColor              = "SDL_VIDEO_X11_NODIRECTCOLOR"
	HintVideoX11ScalingFactor              = "SDL_VIDEO_X11_SCALING_FACTOR"
	HintVideoX11VisualID                   = "SDL_VIDEO_X11_VISUALID"
	HintVideoX11WindowVisualID             = "SDL_VIDEO_X11_WINDOW_VISUALID"
	HintVideoX11Xrandr                     = "SDL_VIDEO_X11_XRANDR"
	HintVitaEnableBackTouch                = "SDL_VITA_ENABLE_BACK_TOUCH"
	HintVitaEnableFrontTouch               = "SDL_VITA_ENABLE_FRONT_TOUCH"
	HintVitaModulePath                     = "SDL_VITA_MODULE_PATH"
	HintVitaPVRInit                        = "SDL_VITA_PVR_INIT"
	HintVitaResolution                     = "SDL_VITA_RESOLUTION"
	HintVitaPVROpenGL                      = "SDL_VITA_PVR_OpenGL"
	HintVitaTouchMouseDevice               = "SDL_VITA_TOUCH_MOUSE_DEVICE"
	HintVulkanDisplay                      = "SDL_VULKAN_DISPLAY"
	HintVulkanLibrary                      = "SDL_VULKAN_LIBRARY"
	HintWaveFactChunk                      = "SDL_WAVE_FACT_CHUNK"
	HintWaveChunkLimit                     = "SDL_WAVE_CHUNK_LIMIT"
	HintWaveRiffChunkSize                  = "SDL_WAVE_RIFF_CHUNK_SIZE"
	HintWaveTruncation                     = "SDL_WAVE_TRUNCATION"
	HintWindowActivateWhenRaised           = "SDL_WINDOW_ACTIVATE_WHEN_RAISED"
	HintWindowActivateWhenShown            = "SDL_WINDOW_ACTIVATE_WHEN_SHOWN"
	HintWindowAllowTopmost                 = "SDL_WINDOW_ALLOW_TOPMOST"
	HintWindowFrameUsableWhileCursorHidden = "SDL_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN"
	HintWindowsCloseOnAltF4                = "SDL_WINDOWS_CLOSE_ON_ALT_F4"
	HintWindowsEnableMenuMnemonics         = "SDL_WINDOWS_ENABLE_MENU_MNEMONICS"
	HintWindowsEnableMessageloop           = "SDL_WINDOWS_ENABLE_MESSAGELOOP"
	HintWindowsGameInput                   = "SDL_WINDOWS_GAMEINPUT"
	HintWindowsRawKeyboard                 = "SDL_WINDOWS_RAW_KEYBOARD"
	HintWindowsForceSemaphoreKernel        = "SDL_WINDOWS_FORCE_SEMAPHORE_KERNEL"
	HintWindowsIntresourceIcon             = "SDL_WINDOWS_INTRESOURCE_ICON"
	HintWindowsIntresourceIconSmall        = "SDL_WINDOWS_INTRESOURCE_ICON_SMALL"
	HintWindowsUseD3D9Ex                   = "SDL_WINDOWS_USE_D3D9EX"
	HintWindowsEraseBackgroundMode         = "SDL_WINDOWS_ERASE_BACKGROUND_MODE"
	HintX11ForceOverrideRedirect           = "SDL_X11_FORCE_OVERRIDE_REDIRECT"
	HintX11WindowType                      = "SDL_X11_WINDOW_TYPE"
	HintX11XcbLibrary                      = "SDL_X11_XCB_LIBRARY"
	HintXinputEnabled                      = "SDL_XINPUT_ENABLED"
	HintAssert                             = "SDL_ASSERT"
	HintPenMouseEvents                     = "SDL_PEN_MOUSE_EVENTS"
	HintPenTouchEvents                     = "SDL_PEN_TOUCH_EVENTS"
)

// [HintCallback] is a callback used to send notifications of hint value changes.
//
// [HintCallback]: https://wiki.libsdl.org/SDL3/SDL_HintCallback
type HintCallback uintptr

func NewHintCallback(callback func(userdata unsafe.Pointer, name, oldValue, newValue string)) HintCallback {
	cb := purego.NewCallback(func(userdata unsafe.Pointer, name, oldValue, newValue *byte) uintptr {
		callback(userdata, convert.ToString(name), convert.ToString(oldValue), convert.ToString(newValue))
		return 0
	})

	return HintCallback(cb)
}

// [SetHint] sets a hint with normal priority.
//
// [SetHint]: https://wiki.libsdl.org/SDL3/SDL_SetHint
func SetHint(name, value string) bool {
	return sdlSetHint(name, value)
}

// [AddHintCallback] adds a function to watch a particular hint.
//
// [AddHintCallback]: https://wiki.libsdl.org/SDL3/SDL_AddHintCallback
func AddHintCallback(name string, callback HintCallback, userdata unsafe.Pointer) bool {
	return sdlAddHintCallback(name, callback, userdata)
}

// [GetHint] gets the value of a hint.
//
// [GetHint]: https://wiki.libsdl.org/SDL3/SDL_GetHint
func GetHint(name string) string {
	return sdlGetHint(name)
}

// [GetHintBoolean] gets the boolean value of a hint variable.
//
// [GetHintBoolean]: https://wiki.libsdl.org/SDL3/SDL_GetHintBoolean
func GetHintBoolean(name string, defaultValue bool) bool {
	return sdlGetHintBoolean(name, defaultValue)
}

// [RemoveHintCallback] removes a function watching a particular hint.
//
// [RemoveHintCallback]: https://wiki.libsdl.org/SDL3/SDL_RemoveHintCallback
func RemoveHintCallback(name string, callback HintCallback, userdata unsafe.Pointer) {
	sdlRemoveHintCallback(name, callback, userdata)
}

// [ResetHint] resets a hint to the default value.
//
// [ResetHint]: https://wiki.libsdl.org/SDL3/SDL_ResetHint
func ResetHint(name string) bool {
	return sdlResetHint(name)
}

// [ResetHints] resets all hints to the default values.
//
// [ResetHints]: https://wiki.libsdl.org/SDL3/SDL_ResetHints
func ResetHints() {
	sdlResetHints()
}

// [SetHintWithPriority] sets a hint with a specific priority.
//
// [SetHintWithPriority]: https://wiki.libsdl.org/SDL3/SDL_SetHintWithPriority
func SetHintWithPriority(name string, value string, priority HintPriority) bool {
	return sdlSetHintWithPriority(name, value, priority)
}
