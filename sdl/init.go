package sdl

import (
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/shared"
)

var (
	// sdlabs                                   func(int32) int32
	// sdlacos                                  func(float64) float64
	// sdlacosf                                 func(float32) float32
	sdlAcquireCameraFrame         func(*Camera, *uint64) *Surface
	sdlAcquireGPUCommandBuffer    func(*GPUDevice) *GPUCommandBuffer
	sdlAcquireGPUSwapchainTexture func(*GPUCommandBuffer, *Window, **GPUTexture, *uint32, *uint32) bool
	// sdlAddAtomicInt                          func(*AtomicInt, int32) int32
	sdlAddEventWatch func(EventFilter, unsafe.Pointer) bool
	// sdlAddGamepadMapping                     func(string) int32
	// sdlAddGamepadMappingsFromFile            func(string) int32
	// sdlAddGamepadMappingsFromIO              func(*IOStream, bool) int32
	sdlAddHintCallback          func(string, HintCallback, unsafe.Pointer) bool
	sdlAddSurfaceAlternateImage func(*Surface, *Surface) bool
	// sdlAddTimer                              func(uint32, TimerCallback, unsafe.Pointer) TimerID
	// sdlAddTimerNS                            func(uint64, NSTimerCallback, unsafe.Pointer) TimerID
	sdlAddVulkanRenderSemaphores func(*Renderer, uint32, int64, int64) bool
	// sdlaligned_alloc                         func(uint64, uint64) unsafe.Pointer
	// sdlaligned_free                          func(unsafe.Pointer)
	// sdlasin                                  func(float64) float64
	// sdlasinf                                 func(float32) float32
	// sdlasprintf                              func(**byte, string) int32
	// sdlAsyncIOFromFile                       func(string, string) *AsyncIO
	// sdlatan                                  func(float64) float64
	// sdlatan2                                 func(float64, float64) float64
	// sdlatan2f                                func(float32, float32) float32
	// sdlatanf                                 func(float32) float32
	// sdlatof                                  func(string) float64
	// sdlatoi                                  func(string) int32
	// sdlAttachVirtualJoystick                 func(*VirtualJoystickDesc) JoystickID
	// sdlAudioDevicePaused                     func(AudioDeviceID) bool
	sdlAudioStreamDevicePaused uintptr
	// sdlBeginGPUComputePass                   func(*GPUCommandBuffer, *GPUStorageTextureReadWriteBinding, uint32, *GPUStorageBufferReadWriteBinding, uint32) *GPUComputePass
	sdlBeginGPUCopyPass   func(*GPUCommandBuffer) *GPUCopyPass
	sdlBeginGPURenderPass func(*GPUCommandBuffer, *GPUColorTargetInfo, uint32, *GPUDepthStencilTargetInfo) *GPURenderPass
	// sdlBindAudioStream                       func(AudioDeviceID, *AudioStream) bool
	// sdlBindAudioStreams                      func(AudioDeviceID, **AudioStream, int32) bool
	// sdlBindGPUComputePipeline                func(*GPUComputePass, *GPUComputePipeline)
	// sdlBindGPUComputeSamplers                func(*GPUComputePass, uint32, *GPUTextureSamplerBinding, uint32)
	// sdlBindGPUComputeStorageBuffers          func(*GPUComputePass, uint32, **GPUBuffer, uint32)
	// sdlBindGPUComputeStorageTextures         func(*GPUComputePass, uint32, **GPUTexture, uint32)
	sdlBindGPUFragmentSamplers func(*GPURenderPass, uint32, *GPUTextureSamplerBinding, uint32)
	// sdlBindGPUFragmentStorageBuffers         func(*GPURenderPass, uint32, **GPUBuffer, uint32)
	// sdlBindGPUFragmentStorageTextures        func(*GPURenderPass, uint32, **GPUTexture, uint32)
	sdlBindGPUGraphicsPipeline func(*GPURenderPass, *GPUGraphicsPipeline)
	sdlBindGPUIndexBuffer      func(*GPURenderPass, *GPUBufferBinding, GPUIndexElementSize)
	sdlBindGPUVertexBuffers    func(*GPURenderPass, uint32, *GPUBufferBinding, uint32)
	// sdlBindGPUVertexSamplers                 func(*GPURenderPass, uint32, *GPUTextureSamplerBinding, uint32)
	// sdlBindGPUVertexStorageBuffers           func(*GPURenderPass, uint32, **GPUBuffer, uint32)
	// sdlBindGPUVertexStorageTextures          func(*GPURenderPass, uint32, **GPUTexture, uint32)
	// sdlBlitGPUTexture                        func(*GPUCommandBuffer, *GPUBlitInfo)
	sdlBlitSurface func(*Surface, *Rect, *Surface, *Rect) bool
	// sdlBlitSurface9Grid                      func(*Surface, *Rect, int32, int32, int32, int32, float32, ScaleMode, *Surface, *Rect) bool
	// sdlBlitSurfaceScaled                     func(*Surface, *Rect, *Surface, *Rect, ScaleMode) bool
	// sdlBlitSurfaceTiled                      func(*Surface, *Rect, *Surface, *Rect) bool
	// sdlBlitSurfaceTiledWithScale             func(*Surface, *Rect, float32, ScaleMode, *Surface, *Rect) bool
	// sdlBlitSurfaceUnchecked                  func(*Surface, *Rect, *Surface, *Rect) bool
	// sdlBlitSurfaceUncheckedScaled            func(*Surface, *Rect, *Surface, *Rect, ScaleMode) bool
	// sdlBroadcastCondition                    func(*Condition)
	// sdlbsearch                               func(unsafe.Pointer, unsafe.Pointer, uint64, uint64, CompareCallback) unsafe.Pointer
	// sdlbsearch_r                             func(unsafe.Pointer, unsafe.Pointer, uint64, uint64, CompareCallback_r, unsafe.Pointer) unsafe.Pointer
	// sdlCalculateGPUTextureFormatSize         func(GPUTextureFormat, uint32, uint32, uint32) uint32
	// sdlcalloc                                func(uint64, uint64) unsafe.Pointer
	// sdlCancelGPUCommandBuffer                func(*GPUCommandBuffer) bool
	sdlCaptureMouse func(bool) bool
	// sdlceil                                  func(float64) float64
	// sdlceilf                                 func(float32) float32
	sdlClaimWindowForGPUDevice func(*GPUDevice, *Window) bool
	// sdlCleanupTLS                            func()
	sdlClearAudioStream uintptr
	// sdlClearClipboardData                    func() bool
	sdlClearComposition func(*Window) bool
	sdlClearError       func() bool
	sdlClearProperty    func(PropertiesID, string) bool
	// sdlClearSurface                          func(*Surface, float32, float32, float32, float32) bool
	// sdlClickTrayEntry                        func(*TrayEntry)
	// sdlCloseAsyncIO                          func(*AsyncIO, bool, *AsyncIOQueue, unsafe.Pointer) bool
	// sdlCloseAudioDevice                      func(AudioDeviceID)
	sdlCloseCamera  func(*Camera)
	sdlCloseGamepad func(*Gamepad)
	// sdlCloseHaptic                           func(*Haptic)
	sdlCloseIO       func(*IOStream) bool
	sdlCloseJoystick func(*Joystick)
	// sdlCloseSensor                           func(*Sensor)
	// sdlCloseStorage                          func(*Storage) bool
	// sdlCompareAndSwapAtomicInt               func(*AtomicInt, int32, int32) bool
	// sdlCompareAndSwapAtomicPointer           func(*unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	// sdlCompareAndSwapAtomicU32               func(*AtomicU32, uint32, uint32) bool
	// sdlComposeCustomBlendMode                func(BlendFactor, BlendFactor, BlendOperation, BlendFactor, BlendFactor, BlendOperation) BlendMode
	// sdlConvertAudioSamples                   func(*AudioSpec, *uint8, int32, *AudioSpec, **uint8, *int32) bool
	sdlConvertEventToRenderCoordinates func(*Renderer, *Event) bool
	sdlConvertPixels                   func(int32, int32, PixelFormat, unsafe.Pointer, int32, PixelFormat, unsafe.Pointer, int32) bool
	sdlConvertPixelsAndColorspace      func(int32, int32, PixelFormat, Colorspace, PropertiesID, unsafe.Pointer, int32, PixelFormat, Colorspace, PropertiesID, unsafe.Pointer, int32) bool
	sdlConvertSurface                  func(*Surface, PixelFormat) *Surface
	sdlConvertSurfaceAndColorspace     func(*Surface, PixelFormat, *Palette, Colorspace, PropertiesID) *Surface
	// sdlCopyFile                              func(string, string) bool
	// sdlCopyGPUBufferToBuffer                 func(*GPUCopyPass, *GPUBufferLocation, *GPUBufferLocation, uint32, bool)
	// sdlCopyGPUTextureToTexture               func(*GPUCopyPass, *GPUTextureLocation, *GPUTextureLocation, uint32, uint32, uint32, bool)
	sdlCopyProperties func(PropertiesID, PropertiesID) bool
	// sdlcopysign                              func(float64, float64) float64
	// sdlcopysignf                             func(float32, float32) float32
	// sdlCopyStorageFile                       func(*Storage, string, string) bool
	// sdlcos                                   func(float64) float64
	// sdlcosf                                  func(float32) float32
	// sdlcrc16                                 func(uint16, unsafe.Pointer, uint64) uint16
	// sdlcrc32                                 func(uint32, unsafe.Pointer, uint64) uint32
	// sdlCreateAsyncIOQueue                    func() *AsyncIOQueue
	// sdlCreateAudioStream                     func(*AudioSpec, *AudioSpec) *AudioStream
	sdlCreateColorCursor func(*Surface, int32, int32) *Cursor
	// sdlCreateCondition                       func() *Condition
	sdlCreateCursor func(*uint8, *uint8, int32, int32, int32, int32) *Cursor
	// sdlCreateDirectory                       func(string) bool
	// sdlCreateEnvironment                     func(bool) *Environment
	sdlCreateGPUBuffer func(*GPUDevice, *GPUBufferCreateInfo) *GPUBuffer
	// sdlCreateGPUComputePipeline              func(*GPUDevice, *GPUComputePipelineCreateInfo) *GPUComputePipeline
	sdlCreateGPUDevice func(GPUShaderFormat, bool, *byte) *GPUDevice
	// sdlCreateGPUDeviceWithProperties         func(PropertiesID) *GPUDevice
	sdlCreateGPUGraphicsPipeline func(*GPUDevice, *GPUGraphicsPipelineCreateInfo) *GPUGraphicsPipeline
	sdlCreateGPUSampler          func(*GPUDevice, *GPUSamplerCreateInfo) *GPUSampler
	sdlCreateGPUShader           func(*GPUDevice, *GPUShaderCreateInfo) *GPUShader
	sdlCreateGPUTexture          func(*GPUDevice, *GPUTextureCreateInfo) *GPUTexture
	sdlCreateGPUTransferBuffer   func(*GPUDevice, *GPUTransferBufferCreateInfo) *GPUTransferBuffer
	// sdlCreateHapticEffect                    func(*Haptic, *HapticEffect) int32
	// sdlCreateMutex                           func() *Mutex
	sdlCreatePalette func(int32) *Palette
	// sdlCreatePopupWindow                     func(*Window, int32, int32, int32, int32, WindowFlags) *Window
	// sdlCreateProcess                         func(**byte, bool) *Process
	// sdlCreateProcessWithProperties           func(PropertiesID) *Process
	sdlCreateProperties             func() PropertiesID
	sdlCreateRenderer               func(*Window, *byte) *Renderer
	sdlCreateRendererWithProperties func(PropertiesID) *Renderer
	// sdlCreateRWLock                          func() *RWLock
	// sdlCreateSemaphore                       func(uint32) *Semaphore
	sdlCreateSoftwareRenderer func(*Surface) *Renderer
	// sdlCreateStorageDirectory                func(*Storage, string) bool
	sdlCreateSurface               func(int32, int32, PixelFormat) *Surface
	sdlCreateSurfaceFrom           func(int32, int32, PixelFormat, unsafe.Pointer, int32) *Surface
	sdlCreateSurfacePalette        func(*Surface) *Palette
	sdlCreateSystemCursor          func(SystemCursor) *Cursor
	sdlCreateTexture               func(*Renderer, PixelFormat, TextureAccess, int32, int32) *Texture
	sdlCreateTextureFromSurface    func(*Renderer, *Surface) *Texture
	sdlCreateTextureWithProperties func(*Renderer, PropertiesID) *Texture
	// sdlCreateThreadRuntime                   func(ThreadFunction, string, unsafe.Pointer, FunctionPointer, FunctionPointer) *Thread
	// sdlCreateThreadWithPropertiesRuntime     func(PropertiesID, FunctionPointer, FunctionPointer) *Thread
	// sdlCreateTray                            func(*Surface, string) *Tray
	// sdlCreateTrayMenu                        func(*Tray) *TrayMenu
	// sdlCreateTraySubmenu                     func(*TrayEntry) *TrayMenu
	sdlCreateWindow            func(string, int32, int32, WindowFlags) *Window
	sdlCreateWindowAndRenderer func(string, int32, int32, WindowFlags, **Window, **Renderer) bool
	// sdlCreateWindowWithProperties            func(PropertiesID) *Window
	sdlCursorVisible func() bool
	// sdlDateTimeToTime                        func(*DateTime, *Time) bool
	// sdlDelay                                 func(uint32)
	sdlDelayNS func(uint64)
	// sdlDelayPrecise                          func(uint64)
	// sdlDestroyAsyncIOQueue                   func(*AsyncIOQueue)
	sdlDestroyAudioStream func(*AudioStream)
	// sdlDestroyCondition                      func(*Condition)
	sdlDestroyCursor func(*Cursor)
	// sdlDestroyEnvironment                    func(*Environment)
	sdlDestroyGPUDevice func(*GPUDevice)
	// sdlDestroyHapticEffect                   func(*Haptic, int32)
	// sdlDestroyMutex                          func(*Mutex)
	sdlDestroyPalette func(*Palette)
	// sdlDestroyProcess                        func(*Process)
	sdlDestroyProperties func(PropertiesID)
	sdlDestroyRenderer   func(*Renderer)
	// sdlDestroyRWLock                         func(*RWLock)
	// sdlDestroySemaphore                      func(*Semaphore)
	sdlDestroySurface func(*Surface)
	sdlDestroyTexture func(*Texture)
	// sdlDestroyTray                           func(*Tray)
	sdlDestroyWindow func(*Window)
	// sdlDestroyWindowSurface                  func(*Window) bool
	// sdlDetachThread                          func(*Thread)
	// sdlDetachVirtualJoystick                 func(JoystickID) bool
	// sdlDisableScreenSaver                    func() bool
	// sdlDispatchGPUCompute                    func(*GPUComputePass, uint32, uint32, uint32)
	// sdlDispatchGPUComputeIndirect            func(*GPUComputePass, *GPUBuffer, uint32)
	// sdlDownloadFromGPUBuffer                 func(*GPUCopyPass, *GPUBufferRegion, *GPUTransferBufferLocation)
	// sdlDownloadFromGPUTexture                func(*GPUCopyPass, *GPUTextureRegion, *GPUTextureTransferInfo)
	sdlDrawGPUIndexedPrimitives func(*GPURenderPass, uint32, uint32, uint32, int32, uint32)
	// sdlDrawGPUIndexedPrimitivesIndirect      func(*GPURenderPass, *GPUBuffer, uint32, uint32)
	sdlDrawGPUPrimitives func(*GPURenderPass, uint32, uint32, uint32, uint32)
	// sdlDrawGPUPrimitivesIndirect             func(*GPURenderPass, *GPUBuffer, uint32, uint32)
	sdlDuplicateSurface func(*Surface) *Surface
	// sdlEGL_GetCurrentConfig                  func() EGLConfig
	// sdlEGL_GetCurrentDisplay                 func() EGLDisplay
	// sdlEGL_GetProcAddress                    func(string) FunctionPointer
	// sdlEGL_GetWindowSurface                  func(*Window) EGLSurface
	// sdlEGL_SetAttributeCallbacks             func(EGLAttribArrayCallback, EGLIntArrayCallback, EGLIntArrayCallback, unsafe.Pointer)
	// sdlEnableScreenSaver                     func() bool
	// sdlEndGPUComputePass                     func(*GPUComputePass)
	sdlEndGPUCopyPass   func(*GPUCopyPass)
	sdlEndGPURenderPass func(*GPURenderPass)
	// sdlEnterAppMainCallbacks                 func(int32, **byte, AppInit_func, AppIterate_func, AppEvent_func, AppQuit_func) int32
	// sdlEnumerateDirectory                    func(string, EnumerateDirectoryCallback, unsafe.Pointer) bool
	sdlEnumerateProperties func(PropertiesID, EnumeratePropertiesCallback, unsafe.Pointer) bool
	// sdlEnumerateStorageDirectory             func(*Storage, string, EnumerateDirectoryCallback, unsafe.Pointer) bool
	sdlEventEnabled func(EventType) bool
	// sdlexp                                   func(float64) float64
	// sdlexpf                                  func(float32) float32
	// sdlfabs                                  func(float64) float64
	// sdlfabsf                                 func(float32) float32
	sdlFillSurfaceRect func(*Surface, *Rect, uint32) bool
	// sdlFillSurfaceRects                      func(*Surface, *Rect, int32, uint32) bool
	sdlFilterEvents func(EventFilter, unsafe.Pointer)
	sdlFlashWindow  func(*Window, FlashOperation) bool
	sdlFlipSurface  func(*Surface, FlipMode) bool
	// sdlfloor                                 func(float64) float64
	// sdlfloorf                                func(float32) float32
	sdlFlushAudioStream uintptr
	sdlFlushEvent       func(EventType)
	sdlFlushEvents      func(EventType, EventType)
	// sdlFlushIO                               func(*IOStream) bool
	sdlFlushRenderer uintptr
	// sdlfmod                                  func(float64, float64) float64
	// sdlfmodf                                 func(float32, float32) float32
	sdlfree uintptr
	// sdlGamepadConnected                      func(*Gamepad) bool
	// sdlGamepadEventsEnabled                  func() bool
	// sdlGamepadHasAxis                        func(*Gamepad, GamepadAxis) bool
	// sdlGamepadHasButton                      func(*Gamepad, GamepadButton) bool
	// sdlGamepadHasSensor                      func(*Gamepad, SensorType) bool
	// sdlGamepadSensorEnabled                  func(*Gamepad, SensorType) bool
	// sdlGDKSuspendComplete                    func()
	// sdlGenerateMipmapsForGPUTexture          func(*GPUCommandBuffer, *GPUTexture)
	sdlGetAppMetadataProperty func(string) string
	// sdlGetAssertionHandler                   func(*unsafe.Pointer) AssertionHandler
	// sdlGetAssertionReport                    func() *AssertData
	// sdlGetAsyncIOResult                      func(*AsyncIOQueue, *AsyncIOOutcome) bool
	// sdlGetAsyncIOSize                        func(*AsyncIO) int64
	// sdlGetAtomicInt                          func(*AtomicInt) int32
	// sdlGetAtomicPointer                      func(*unsafe.Pointer) unsafe.Pointer
	// sdlGetAtomicU32                          func(*AtomicU32) uint32
	// sdlGetAudioDeviceChannelMap              func(AudioDeviceID, *int32) *int32
	// sdlGetAudioDeviceFormat                  func(AudioDeviceID, *AudioSpec, *int32) bool
	// sdlGetAudioDeviceGain                    func(AudioDeviceID) float32
	// sdlGetAudioDeviceName                    func(AudioDeviceID) string
	sdlGetAudioDriver func(int32) string
	// sdlGetAudioFormatName                    func(AudioFormat) string
	// sdlGetAudioPlaybackDevices               func(*int32) *AudioDeviceID
	// sdlGetAudioRecordingDevices              func(*int32) *AudioDeviceID
	// sdlGetAudioStreamAvailable               func(*AudioStream) int32
	// sdlGetAudioStreamData                    func(*AudioStream, unsafe.Pointer, int32) int32
	// sdlGetAudioStreamDevice                  func(*AudioStream) AudioDeviceID
	// sdlGetAudioStreamFormat                  func(*AudioStream, *AudioSpec, *AudioSpec) bool
	// sdlGetAudioStreamFrequencyRatio          func(*AudioStream) float32
	// sdlGetAudioStreamGain                    func(*AudioStream) float32
	// sdlGetAudioStreamInputChannelMap         func(*AudioStream, *int32) *int32
	// sdlGetAudioStreamOutputChannelMap        func(*AudioStream, *int32) *int32
	// sdlGetAudioStreamProperties              func(*AudioStream) PropertiesID
	sdlGetAudioStreamQueued      uintptr
	sdlGetBasePath               func() string
	sdlGetBooleanProperty        func(PropertiesID, string, bool) bool
	sdlGetCameraDriver           func(int32) string
	sdlGetCameraFormat           func(*Camera, *CameraSpec) bool
	sdlGetCameraID               func(*Camera) CameraID
	sdlGetCameraName             func(CameraID) string
	sdlGetCameraPermissionState  func(*Camera) int32
	sdlGetCameraPosition         func(CameraID) CameraPosition
	sdlGetCameraProperties       func(*Camera) PropertiesID
	sdlGetCameras                func(*int32) *CameraID
	sdlGetCameraSupportedFormats func(CameraID, *int32) **CameraSpec
	// sdlGetClipboardData                      func(string, *uint64) unsafe.Pointer
	// sdlGetClipboardMimeTypes                 func(*uint64) **byte
	sdlGetClipboardText                func() *byte
	sdlGetClosestFullscreenDisplayMode func(DisplayID, int32, int32, float32, bool, *DisplayMode) bool
	// sdlGetCPUCacheLineSize                   func() int32
	sdlGetCurrentAudioDriver  func() string
	sdlGetCurrentCameraDriver func() string
	// sdlGetCurrentDirectory                   func() string
	sdlGetCurrentDisplayMode func(DisplayID) *DisplayMode
	// sdlGetCurrentDisplayOrientation          func(DisplayID) DisplayOrientation
	sdlGetCurrentRenderOutputSize func(*Renderer, *int32, *int32) bool
	// sdlGetCurrentThreadID                    func() ThreadID
	// sdlGetCurrentTime                        func(*Time) bool
	sdlGetCurrentVideoDriver func() string
	sdlGetCursor             func() *Cursor
	// sdlGetDateTimeLocalePreferences          func(*DateFormat, *TimeFormat) bool
	// sdlGetDayOfWeek                          func(int32, int32, int32) int32
	// sdlGetDayOfYear                          func(int32, int32, int32) int32
	// sdlGetDaysInMonth                        func(int32, int32) int32
	// sdlGetDefaultAssertionHandler            func() AssertionHandler
	sdlGetDefaultCursor func() *Cursor
	// sdlGetDefaultLogOutputFunction           func() LogOutputFunction
	// sdlGetDesktopDisplayMode                 func(DisplayID) *DisplayMode
	// sdlGetDisplayBounds                      func(DisplayID, *Rect) bool
	sdlGetDisplayContentScale func(DisplayID) float32
	// sdlGetDisplayForPoint                    func(*Point) DisplayID
	// sdlGetDisplayForRect                     func(*Rect) DisplayID
	sdlGetDisplayForWindow func(*Window) DisplayID
	sdlGetDisplayName      func(DisplayID) string
	// sdlGetDisplayProperties                  func(DisplayID) PropertiesID
	sdlGetDisplays func(*int32) *DisplayID
	// sdlGetDisplayUsableBounds                func(DisplayID, *Rect) bool
	// sdlgetenv                                func(string) string
	// sdlgetenv_unsafe                         func(string) string
	// sdlGetEnvironment                        func() *Environment
	// sdlGetEnvironmentVariable                func(*Environment, string) string
	// sdlGetEnvironmentVariables               func(*Environment) **byte
	sdlGetError                  func() string
	sdlGetEventFilter            func(*EventFilter, *unsafe.Pointer) bool
	sdlGetFloatProperty          func(PropertiesID, string, float32) float32
	sdlGetFullscreenDisplayModes func(DisplayID, *int32) **DisplayMode
	// sdlGetGamepadAppleSFSymbolsNameForAxis   func(*Gamepad, GamepadAxis) string
	// sdlGetGamepadAppleSFSymbolsNameForButton func(*Gamepad, GamepadButton) string
	// sdlGetGamepadAxis                        func(*Gamepad, GamepadAxis) int16
	// sdlGetGamepadAxisFromString              func(string) GamepadAxis
	sdlGetGamepadBindings func(*Gamepad, *int32) **GamepadBinding
	// sdlGetGamepadButton                      func(*Gamepad, GamepadButton) bool
	// sdlGetGamepadButtonFromString            func(string) GamepadButton
	// sdlGetGamepadButtonLabel                 func(*Gamepad, GamepadButton) GamepadButtonLabel
	// sdlGetGamepadButtonLabelForType          func(GamepadType, GamepadButton) GamepadButtonLabel
	// sdlGetGamepadConnectionState             func(*Gamepad) JoystickConnectionState
	// sdlGetGamepadFirmwareVersion             func(*Gamepad) uint16
	sdlGetGamepadFromID func(JoystickID) *Gamepad
	// sdlGetGamepadFromPlayerIndex             func(int32) *Gamepad
	// sdlGetGamepadGUIDForID                   func(JoystickID) GUID
	// sdlGetGamepadID                          func(*Gamepad) JoystickID
	// sdlGetGamepadJoystick                    func(*Gamepad) *Joystick
	// sdlGetGamepadMapping                     func(*Gamepad) string
	// sdlGetGamepadMappingForGUID              func(GUID) string
	// sdlGetGamepadMappingForID                func(JoystickID) string
	// sdlGetGamepadMappings                    func(*int32) **byte
	sdlGetGamepadName      func(*Gamepad) string
	sdlGetGamepadNameForID func(JoystickID) string
	// sdlGetGamepadPath                        func(*Gamepad) string
	// sdlGetGamepadPathForID                   func(JoystickID) string
	// sdlGetGamepadPlayerIndex                 func(*Gamepad) int32
	// sdlGetGamepadPlayerIndexForID            func(JoystickID) int32
	// sdlGetGamepadPowerInfo                   func(*Gamepad, *int32) PowerState
	// sdlGetGamepadProduct                     func(*Gamepad) uint16
	// sdlGetGamepadProductForID                func(JoystickID) uint16
	// sdlGetGamepadProductVersion              func(*Gamepad) uint16
	// sdlGetGamepadProductVersionForID         func(JoystickID) uint16
	// sdlGetGamepadProperties                  func(*Gamepad) PropertiesID
	sdlGetGamepads func(*int32) *JoystickID
	// sdlGetGamepadSensorData                  func(*Gamepad, SensorType, *float32, int32) bool
	// sdlGetGamepadSensorDataRate              func(*Gamepad, SensorType) float32
	sdlGetGamepadSerial func(*Gamepad) string
	// sdlGetGamepadSteamHandle                 func(*Gamepad) uint64
	// sdlGetGamepadStringForAxis               func(GamepadAxis) string
	sdlGetGamepadStringForButton func(GamepadButton) string
	sdlGetGamepadStringForType   func(GamepadType) string
	// sdlGetGamepadTouchpadFinger              func(*Gamepad, int32, int32, *bool, *float32, *float32, *float32) bool
	sdlGetGamepadType func(*Gamepad) GamepadType
	// sdlGetGamepadTypeForID                   func(JoystickID) GamepadType
	// sdlGetGamepadTypeFromString              func(string) GamepadType
	// sdlGetGamepadVendor                      func(*Gamepad) uint16
	// sdlGetGamepadVendorForID                 func(JoystickID) uint16
	sdlGetGlobalMouseState          func(*float32, *float32) MouseButtonFlags
	sdlGetGlobalProperties          func() PropertiesID
	sdlGetGPUDeviceDriver           func(*GPUDevice) string
	sdlGetGPUDriver                 func(int32) string
	sdlGetGPUShaderFormats          func(*GPUDevice) GPUShaderFormat
	sdlGetGPUSwapchainTextureFormat func(*GPUDevice, *Window) GPUTextureFormat
	// sdlGetGrabbedWindow                      func() *Window
	// sdlGetHapticEffectStatus                 func(*Haptic, int32) bool
	// sdlGetHapticFeatures                     func(*Haptic) uint32
	// sdlGetHapticFromID                       func(HapticID) *Haptic
	// sdlGetHapticID                           func(*Haptic) HapticID
	// sdlGetHapticName                         func(*Haptic) string
	// sdlGetHapticNameForID                    func(HapticID) string
	// sdlGetHaptics                            func(*int32) *HapticID
	sdlGetHint        func(string) string
	sdlGetHintBoolean func(string, bool) bool
	// sdlGetIOProperties                       func(*IOStream) PropertiesID
	// sdlGetIOSize                             func(*IOStream) int64
	// sdlGetIOStatus                           func(*IOStream) IOStatus
	sdlGetJoystickAxis             func(*Joystick, int32) int16
	sdlGetJoystickAxisInitialState func(*Joystick, int32, *int16) bool
	sdlGetJoystickBall             func(*Joystick, int32, *int32, *int32) bool
	sdlGetJoystickButton           func(*Joystick, int32) bool
	sdlGetJoystickConnectionState  func(*Joystick) JoystickConnectionState
	sdlGetJoystickFirmwareVersion  func(*Joystick) uint16
	sdlGetJoystickFromID           func(JoystickID) *Joystick
	sdlGetJoystickFromPlayerIndex  func(int32) *Joystick
	// sdlGetJoystickGUID                func(*Joystick) GUID
	// sdlGetJoystickGUIDForID           func(JoystickID) GUID
	// sdlGetJoystickGUIDInfo            func(GUID, *uint16, *uint16, *uint16, *uint16)
	sdlGetJoystickHat                 func(*Joystick, int32) uint8
	sdlGetJoystickID                  func(*Joystick) JoystickID
	sdlGetJoystickName                func(*Joystick) string
	sdlGetJoystickNameForID           func(JoystickID) string
	sdlGetJoystickPath                func(*Joystick) string
	sdlGetJoystickPathForID           func(JoystickID) string
	sdlGetJoystickPlayerIndex         func(*Joystick) int32
	sdlGetJoystickPlayerIndexForID    func(JoystickID) int32
	sdlGetJoystickPowerInfo           func(*Joystick, *int32) PowerState
	sdlGetJoystickProduct             func(*Joystick) uint16
	sdlGetJoystickProductForID        func(JoystickID) uint16
	sdlGetJoystickProductVersion      func(*Joystick) uint16
	sdlGetJoystickProductVersionForID func(JoystickID) uint16
	sdlGetJoystickProperties          func(*Joystick) PropertiesID
	sdlGetJoysticks                   func(*int32) *JoystickID
	sdlGetJoystickSerial              func(*Joystick) string
	sdlGetJoystickType                func(*Joystick) JoystickType
	sdlGetJoystickTypeForID           func(JoystickID) JoystickType
	sdlGetJoystickVendor              func(*Joystick) uint16
	sdlGetJoystickVendorForID         func(JoystickID) uint16
	sdlGetKeyboardFocus               func() *Window
	sdlGetKeyboardNameForID           func(KeyboardID) string
	sdlGetKeyboards                   func(*int32) *KeyboardID
	sdlGetKeyboardState               func(*int32) *bool
	sdlGetKeyFromName                 func(string) Keycode
	sdlGetKeyFromScancode             func(Scancode, Keymod, bool) Keycode
	sdlGetKeyName                     func(Keycode) string
	// sdlGetLogOutputFunction                  func(*LogOutputFunction, *unsafe.Pointer)
	// sdlGetLogPriority                        func(int32) LogPriority
	// sdlGetMasksForPixelFormat                func(PixelFormat, *int32, *uint32, *uint32, *uint32, *uint32) bool
	// sdlGetMaxHapticEffects                   func(*Haptic) int32
	// sdlGetMaxHapticEffectsPlaying            func(*Haptic) int32
	// sdlGetMemoryFunctions                    func(*malloc_func, *calloc_func, *realloc_func, *free_func)
	sdlGetMice           func(*int32) *MouseID
	sdlGetModState       func() Keymod
	sdlGetMouseFocus     func() *Window
	sdlGetMouseNameForID func(MouseID) string
	sdlGetMouseState     func(*float32, *float32) MouseButtonFlags
	// sdlGetNaturalDisplayOrientation          func(DisplayID) DisplayOrientation
	// sdlGetNumAllocations                     func() int32
	sdlGetNumAudioDrivers  func() int32
	sdlGetNumberProperty   func(PropertiesID, string, int64) int64
	sdlGetNumCameraDrivers func() int32
	// sdlGetNumGamepadTouchpadFingers          func(*Gamepad, int32) int32
	// sdlGetNumGamepadTouchpads                func(*Gamepad) int32
	sdlGetNumGPUDrivers func() int32
	// sdlGetNumHapticAxes                      func(*Haptic) int32
	sdlGetNumJoystickAxes    func(*Joystick) int32
	sdlGetNumJoystickBalls   func(*Joystick) int32
	sdlGetNumJoystickButtons func(*Joystick) int32
	sdlGetNumJoystickHats    func(*Joystick) int32
	// sdlGetNumLogicalCPUCores                 func() int32
	sdlGetNumRenderDrivers func() int32
	sdlGetNumVideoDrivers  func() int32
	// sdlGetOriginalMemoryFunctions            func(*malloc_func, *calloc_func, *realloc_func, *free_func)
	// sdlGetPathInfo                           func(string, *PathInfo) bool
	sdlGetPerformanceCounter   uintptr
	sdlGetPerformanceFrequency uintptr
	sdlGetPixelFormatDetails   func(PixelFormat) *PixelFormatDetails
	// sdlGetPixelFormatForMasks                func(int32, uint32, uint32, uint32, uint32) PixelFormat
	// sdlGetPixelFormatName                    func(PixelFormat) string
	// sdlGetPlatform                           func() string
	sdlGetPointerProperty  func(PropertiesID, string, unsafe.Pointer) unsafe.Pointer
	sdlGetPowerInfo        func(*int32, *int32) PowerState
	sdlGetPreferredLocales func(*int32) **struct{ language, country *byte }
	// sdlGetPrefPath                           func(string, string) string
	sdlGetPrimaryDisplay func() DisplayID
	// sdlGetPrimarySelectionText               func() string
	// sdlGetProcessInput                       func(*Process) *IOStream
	// sdlGetProcessOutput                      func(*Process) *IOStream
	// sdlGetProcessProperties                  func(*Process) PropertiesID
	sdlGetPropertyType func(PropertiesID, string) PropertyType
	// sdlGetRealGamepadType                    func(*Gamepad) GamepadType
	// sdlGetRealGamepadTypeForID               func(JoystickID) GamepadType
	sdlGetRectAndLineIntersection       func(*Rect, *int32, *int32, *int32, *int32) bool
	sdlGetRectAndLineIntersectionFloat  func(*FRect, *float32, *float32, *float32, *float32) bool
	sdlGetRectEnclosingPoints           func(*Point, int32, *Rect, *Rect) bool
	sdlGetRectEnclosingPointsFloat      func(*FPoint, int32, *FRect, *FRect) bool
	sdlGetRectIntersection              func(*Rect, *Rect, *Rect) bool
	sdlGetRectIntersectionFloat         func(*FRect, *FRect, *FRect) bool
	sdlGetRectUnion                     func(*Rect, *Rect, *Rect) bool
	sdlGetRectUnionFloat                func(*FRect, *FRect, *FRect) bool
	sdlGetRelativeMouseState            func(*float32, *float32) MouseButtonFlags
	sdlGetRenderClipRect                func(*Renderer, *Rect) bool
	sdlGetRenderColorScale              func(*Renderer, *float32) bool
	sdlGetRenderDrawBlendMode           func(*Renderer, *BlendMode) bool
	sdlGetRenderDrawColor               func(*Renderer, *uint8, *uint8, *uint8, *uint8) bool
	sdlGetRenderDrawColorFloat          func(*Renderer, *float32, *float32, *float32, *float32) bool
	sdlGetRenderDriver                  func(int32) string
	sdlGetRenderer                      func(*Window) *Renderer
	sdlGetRendererFromTexture           func(*Texture) *Renderer
	sdlGetRendererName                  func(*Renderer) string
	sdlGetRendererProperties            func(*Renderer) PropertiesID
	sdlGetRenderLogicalPresentation     func(*Renderer, *int32, *int32, *RendererLogicalPresentation) bool
	sdlGetRenderLogicalPresentationRect func(*Renderer, *FRect) bool
	sdlGetRenderMetalCommandEncoder     func(*Renderer) unsafe.Pointer
	sdlGetRenderMetalLayer              func(*Renderer) unsafe.Pointer
	sdlGetRenderOutputSize              func(*Renderer, *int32, *int32) bool
	sdlGetRenderSafeArea                func(*Renderer, *Rect) bool
	sdlGetRenderScale                   func(*Renderer, *float32, *float32) bool
	sdlGetRenderTarget                  func(*Renderer) *Texture
	sdlGetRenderViewport                func(*Renderer, *Rect) bool
	sdlGetRenderVSync                   func(*Renderer, *int32) bool
	sdlGetRenderWindow                  func(*Renderer) *Window
	sdlGetRevision                      func() string
	// sdlGetRGB                                func(uint32, *PixelFormatDetails, *Palette, *uint8, *uint8, *uint8)
	// sdlGetRGBA                               func(uint32, *PixelFormatDetails, *Palette, *uint8, *uint8, *uint8, *uint8)
	// sdlGetSandbox                            func() Sandbox
	sdlGetScancodeFromKey  func(Keycode, *Keymod) Scancode
	sdlGetScancodeFromName func(string) Scancode
	sdlGetScancodeName     func(Scancode) string
	// sdlGetSemaphoreValue                     func(*Semaphore) uint32
	// sdlGetSensorData                         func(*Sensor, *float32, int32) bool
	// sdlGetSensorFromID                       func(SensorID) *Sensor
	// sdlGetSensorID                           func(*Sensor) SensorID
	// sdlGetSensorName                         func(*Sensor) string
	// sdlGetSensorNameForID                    func(SensorID) string
	// sdlGetSensorNonPortableType              func(*Sensor) int32
	// sdlGetSensorNonPortableTypeForID         func(SensorID) int32
	// sdlGetSensorProperties                   func(*Sensor) PropertiesID
	// sdlGetSensors                            func(*int32) *SensorID
	// sdlGetSensorType                         func(*Sensor) SensorType
	// sdlGetSensorTypeForID                    func(SensorID) SensorType
	// sdlGetSilenceValueForFormat              func(AudioFormat) int32
	// sdlGetSIMDAlignment                      func() uint64
	// sdlGetStorageFileSize                    func(*Storage, string, *uint64) bool
	// sdlGetStoragePathInfo                    func(*Storage, string, *PathInfo) bool
	// sdlGetStorageSpaceRemaining              func(*Storage) uint64
	sdlGetStringProperty    func(PropertiesID, string, string) string
	sdlGetSurfaceAlphaMod   func(*Surface, *uint8) bool
	sdlGetSurfaceBlendMode  func(*Surface, *BlendMode) bool
	sdlGetSurfaceClipRect   func(*Surface, *Rect) bool
	sdlGetSurfaceColorKey   func(*Surface, *uint32) bool
	sdlGetSurfaceColorMod   func(*Surface, *uint8, *uint8, *uint8) bool
	sdlGetSurfaceColorspace func(*Surface) Colorspace
	sdlGetSurfaceImages     func(*Surface, *int32) **Surface
	sdlGetSurfacePalette    func(*Surface) *Palette
	sdlGetSurfaceProperties func(*Surface) PropertiesID
	// sdlGetSystemRAM                          func() int32
	// sdlGetSystemTheme                        func() SystemTheme
	sdlGetTextInputArea        func(*Window, *Rect, *int32) bool
	sdlGetTextureAlphaMod      func(*Texture, *uint8) bool
	sdlGetTextureAlphaModFloat func(*Texture, *float32) bool
	sdlGetTextureBlendMode     func(*Texture, *BlendMode) bool
	sdlGetTextureColorMod      func(*Texture, *uint8, *uint8, *uint8) bool
	sdlGetTextureColorModFloat func(*Texture, *float32, *float32, *float32) bool
	sdlGetTextureProperties    func(*Texture) PropertiesID
	sdlGetTextureScaleMode     func(*Texture, *ScaleMode) bool
	sdlGetTextureSize          func(*Texture, *float32, *float32) bool
	// sdlGetThreadID                           func(*Thread) ThreadID
	// sdlGetThreadName                         func(*Thread) string
	// sdlGetThreadState                        func(*Thread) ThreadState
	sdlGetTicks   uintptr
	sdlGetTicksNS uintptr
	// sdlGetTLS                                func(*TLSID) unsafe.Pointer
	// sdlGetTouchDeviceName                    func(TouchID) string
	// sdlGetTouchDevices                       func(*int32) *TouchID
	// sdlGetTouchDeviceType                    func(TouchID) TouchDeviceType
	// sdlGetTouchFingers                       func(TouchID, *int32) **Finger
	// sdlGetTrayEntries                        func(*TrayMenu, *int32) **TrayEntry
	// sdlGetTrayEntryChecked                   func(*TrayEntry) bool
	// sdlGetTrayEntryEnabled                   func(*TrayEntry) bool
	// sdlGetTrayEntryLabel                     func(*TrayEntry) string
	// sdlGetTrayEntryParent                    func(*TrayEntry) *TrayMenu
	// sdlGetTrayMenu                           func(*Tray) *TrayMenu
	// sdlGetTrayMenuParentEntry                func(*TrayMenu) *TrayEntry
	// sdlGetTrayMenuParentTray                 func(*TrayMenu) *Tray
	// sdlGetTraySubmenu                        func(*TrayEntry) *TrayMenu
	// sdlGetUserFolder                         func(Folder) string
	sdlGetVersion     func() int32
	sdlGetVideoDriver func(int32) string
	// sdlGetWindowAspectRatio                  func(*Window, *float32, *float32) bool
	// sdlGetWindowBordersSize                  func(*Window, *int32, *int32, *int32, *int32) bool
	sdlGetWindowDisplayScale func(*Window) float32
	sdlGetWindowFlags        func(*Window) WindowFlags
	sdlGetWindowFromEvent    func(*Event) *Window
	// sdlGetWindowFromID                       func(WindowID) *Window
	sdlGetWindowFullscreenMode func(*Window) *DisplayMode
	// sdlGetWindowICCProfile                   func(*Window, *uint64) unsafe.Pointer
	sdlGetWindowID           func(*Window) WindowID
	sdlGetWindowKeyboardGrab func(*Window) bool
	// sdlGetWindowMaximumSize                  func(*Window, *int32, *int32) bool
	// sdlGetWindowMinimumSize                  func(*Window, *int32, *int32) bool
	sdlGetWindowMouseGrab func(*Window) bool
	// sdlGetWindowMouseRect                    func(*Window) *Rect
	sdlGetWindowOpacity func(*Window) float32
	// sdlGetWindowParent                       func(*Window) *Window
	sdlGetWindowPixelDensity func(*Window) float32
	// sdlGetWindowPixelFormat                  func(*Window) PixelFormat
	sdlGetWindowPosition func(*Window, *int32, *int32) bool
	// sdlGetWindowProperties                   func(*Window) PropertiesID
	sdlGetWindowRelativeMouseMode func(*Window) bool
	// sdlGetWindows                            func(*int32) **Window
	// sdlGetWindowSafeArea                     func(*Window, *Rect) bool
	sdlGetWindowSize         func(*Window, *int32, *int32) bool
	sdlGetWindowSizeInPixels func(*Window, *int32, *int32) bool
	sdlGetWindowSurface      func(*Window) *Surface
	// sdlGetWindowSurfaceVSync                 func(*Window, *int32) bool
	sdlGetWindowTitle   func(*Window) string
	sdlGLCreateContext  func(*Window) GLContext
	sdlGLDestroyContext func(GLContext) bool
	// sdlGL_ExtensionSupported                 func(string) bool
	sdlGLGetAttribute      func(GLAttr, *int32) bool
	sdlGLGetCurrentContext func() GLContext
	sdlGLGetCurrentWindow  func() *Window
	// sdlGL_GetProcAddress                     func(string) FunctionPointer
	// sdlGL_GetSwapInterval                    func(*int32) bool
	// sdlGL_LoadLibrary                        func(string) bool
	// sdlGL_MakeCurrent                        func(*Window, GLContext) bool
	// sdlGL_ResetAttributes                    func()
	sdlGLSetAttribute    func(GLAttr, int32) bool
	sdlGLSetSwapInterval uintptr
	sdlGLSwapWindow      uintptr
	// sdlGL_UnloadLibrary                      func()
	// sdlGlobDirectory                         func(string, string, GlobFlags, *int32) **byte
	// sdlGlobStorageDirectory                  func(*Storage, string, string, GlobFlags, *int32) **byte
	// sdlGPUSupportsProperties                 func(PropertiesID) bool
	// sdlGPUSupportsShaderFormats              func(GPUShaderFormat, string) bool
	// sdlGPUTextureFormatTexelBlockSize        func(GPUTextureFormat) uint32
	// sdlGPUTextureSupportsFormat              func(*GPUDevice, GPUTextureFormat, GPUTextureType, GPUTextureUsageFlags) bool
	// sdlGPUTextureSupportsSampleCount         func(*GPUDevice, GPUTextureFormat, GPUSampleCount) bool
	// sdlGUIDToString                          func(GUID, string, int32)
	// sdlHapticEffectSupported                 func(*Haptic, *HapticEffect) bool
	// sdlHapticRumbleSupported                 func(*Haptic) bool
	// sdlHasAltiVec                            func() bool
	// sdlHasARMSIMD                            func() bool
	// sdlHasAVX                                func() bool
	// sdlHasAVX2                               func() bool
	// sdlHasAVX512F                            func() bool
	// sdlHasClipboardData                      func(string) bool
	// sdlHasClipboardText                      func() bool
	sdlHasEvent  func(EventType) bool
	sdlHasEvents func(EventType, EventType) bool
	// sdlHasExactlyOneBitSet32                 func(uint32) bool
	// sdlHasGamepad                            func() bool
	sdlHasJoystick func() bool
	sdlHasKeyboard func() bool
	// sdlHasLASX                               func() bool
	// sdlHasLSX                                func() bool
	// sdlHasMMX                                func() bool
	sdlHasMouse func() bool
	// sdlHasNEON                               func() bool
	// sdlHasPrimarySelectionText               func() bool
	sdlHasProperty              func(PropertiesID, string) bool
	sdlHasRectIntersection      func(*Rect, *Rect) bool
	sdlHasRectIntersectionFloat func(*FRect, *FRect) bool
	sdlHasScreenKeyboardSupport func() bool
	// sdlHasSSE                                func() bool
	// sdlHasSSE2                               func() bool
	// sdlHasSSE3                               func() bool
	// sdlHasSSE41                              func() bool
	// sdlHasSSE42                              func() bool
	// sdlhid_ble_scan                          func(bool)
	// sdlhid_close                             func(*hid_device) int32
	// sdlhid_device_change_count               func() uint32
	// sdlhid_enumerate                         func(uint16, uint16) *hid_device_info
	// sdlhid_exit                              func() int32
	// sdlhid_free_enumeration                  func(*hid_device_info)
	// sdlhid_get_device_info                   func(*hid_device) *hid_device_info
	// sdlhid_get_feature_report                func(*hid_device, *uint8, uint64) int32
	// sdlhid_get_indexed_string                func(*hid_device, int32, *wchar_t, uint64) int32
	// sdlhid_get_input_report                  func(*hid_device, *uint8, uint64) int32
	// sdlhid_get_manufacturer_string           func(*hid_device, *wchar_t, uint64) int32
	// sdlhid_get_product_string                func(*hid_device, *wchar_t, uint64) int32
	// sdlhid_get_report_descriptor             func(*hid_device, *uint8, uint64) int32
	// sdlhid_get_serial_number_string          func(*hid_device, *wchar_t, uint64) int32
	// sdlhid_init                              func() int32
	// sdlhid_open                              func(uint16, uint16, *wchar_t) *hid_device
	// sdlhid_open_path                         func(string) *hid_device
	// sdlhid_read                              func(*hid_device, *uint8, uint64) int32
	// sdlhid_read_timeout                      func(*hid_device, *uint8, uint64, int32) int32
	// sdlhid_send_feature_report               func(*hid_device, *uint8, uint64) int32
	// sdlhid_set_nonblocking                   func(*hid_device, int32) int32
	// sdlhid_write                             func(*hid_device, *uint8, uint64) int32
	sdlHideCursor func() bool
	sdlHideWindow func(*Window) bool
	// sdliconv                                 func(iconv_t, **byte, *uint64, **byte, *uint64) uint64
	// sdliconv_close                           func(iconv_t) int32
	// sdliconv_open                            func(string, string) iconv_t
	// sdliconv_string                          func(string, string, string, uint64) string
	sdlInit func(InitFlags) bool
	// sdlInitHapticRumble                      func(*Haptic) bool
	// sdlInitSubSystem                         func(InitFlags) bool
	// sdlInsertGPUDebugLabel                   func(*GPUCommandBuffer, string)
	// sdlInsertTrayEntryAt                     func(*TrayMenu, int32, string, TrayEntryFlags) *TrayEntry
	sdlIOFromConstMem func([]byte, int) *IOStream
	// sdlIOFromDynamicMem                      func() *IOStream
	sdlIOFromFile func(string, string) *IOStream
	sdlIOFromMem  func([]byte, int) *IOStream
	// sdlIOprintf                              func(*IOStream, string) uint64
	// sdlIOvprintf                             func(*IOStream, string, va_list) uint64
	// sdlisalnum                               func(int32) int32
	// sdlisalpha                               func(int32) int32
	// sdlIsAudioDevicePhysical                 func(AudioDeviceID) bool
	// sdlIsAudioDevicePlayback                 func(AudioDeviceID) bool
	// sdlisblank                               func(int32) int32
	// sdliscntrl                               func(int32) int32
	// sdlisdigit                               func(int32) int32
	// sdlIsGamepad                             func(JoystickID) bool
	// sdlisgraph                               func(int32) int32
	// sdlisinf                                 func(float64) int32
	// sdlisinff                                func(float32) int32
	// sdlIsJoystickHaptic                      func(*Joystick) bool
	// sdlIsJoystickVirtual func(JoystickID) bool
	// sdlislower                               func(int32) int32
	sdlIsMainThread func() bool
	// sdlIsMouseHaptic                         func() bool
	// sdlisnan                                 func(float64) int32
	// sdlisnanf                                func(float32) int32
	// sdlisprint                               func(int32) int32
	// sdlispunct                               func(int32) int32
	// sdlisspace                               func(int32) int32
	// sdlIsTablet                              func() bool
	// sdlIsTV                                  func() bool
	// sdlisupper                               func(int32) int32
	// sdlisxdigit                              func(int32) int32
	// sdlitoa                                  func(int32, string, int32) string
	sdlJoystickConnected     func(*Joystick) bool
	sdlJoystickEventsEnabled func() bool
	// sdlKillProcess                           func(*Process, bool) bool
	// sdllltoa                                 func(int64, string, int32) string
	sdlLoadBMP   func(string) *Surface
	sdlLoadBMPIO func(*IOStream, bool) *Surface
	sdlLoadFile  func(string, *uint64) unsafe.Pointer
	// sdlLoadFile_IO                           func(*IOStream, *uint64, bool) unsafe.Pointer
	// sdlLoadFileAsync                         func(string, *AsyncIOQueue, unsafe.Pointer) bool
	// sdlLoadFunction                          func(*SharedObject, string) FunctionPointer
	// sdlLoadObject                            func(string) *SharedObject
	// sdlLoadWAV                               func(string, *AudioSpec, **uint8, *uint32) bool
	sdlLoadWAVIO func(*IOStream, bool, *AudioSpec, **uint8, *uint32) bool
	// sdlLockAudioStream                       func(*AudioStream) bool
	sdlLockJoysticks func()
	// sdlLockMutex                             func(*Mutex)
	sdlLockProperties func(PropertiesID) bool
	// sdlLockRWLockForReading                  func(*RWLock)
	// sdlLockRWLockForWriting                  func(*RWLock)
	// sdlLockSpinlock                          func(*SpinLock)
	sdlLockSurface          func(*Surface) bool
	sdlLockTexture          func(*Texture, *Rect, *unsafe.Pointer, *int32) bool
	sdlLockTextureToSurface func(*Texture, *Rect, **Surface) bool
	sdlLog                  func(string)
	// sdllog                                   func(float64) float64
	// sdllog10                                 func(float64) float64
	// sdllog10f                                func(float32) float32
	// sdlLogCritical                           func(int32, string)
	// sdlLogDebug                              func(int32, string)
	sdlLogError func(LogCategory, string)
	// sdllogf                                  func(float32) float32
	// sdlLogInfo                               func(int32, string)
	sdlLogMessage func(LogCategory, LogPriority, string)
	// sdlLogMessageV                           func(int32, LogPriority, string, va_list)
	// sdlLogTrace                              func(int32, string)
	// sdlLogVerbose                            func(int32, string)
	// sdlLogWarn                               func(int32, string)
	// sdllround                                func(float64) int64
	// sdllroundf                               func(float32) int64
	// sdlltoa                                  func(int64, string, int32) string
	// sdlmain                                  func(int32, **byte) int32
	// sdlmalloc                                func(uint64) unsafe.Pointer
	sdlMapGPUTransferBuffer func(*GPUDevice, *GPUTransferBuffer, bool) unsafe.Pointer
	sdlMapRGB               func(*PixelFormatDetails, *Palette, uint8, uint8, uint8) uint32
	// sdlMapRGBA                               func(*PixelFormatDetails, *Palette, uint8, uint8, uint8, uint8) uint32
	sdlMapSurfaceRGB func(*Surface, uint8, uint8, uint8) uint32
	// sdlMapSurfaceRGBA                        func(*Surface, uint8, uint8, uint8, uint8) uint32
	// sdlMaximizeWindow                        func(*Window) bool
	// sdlmemcmp                                func(unsafe.Pointer, unsafe.Pointer, uint64) int32
	// sdlmemcpy                                func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	// sdlmemmove                               func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	// sdlMemoryBarrierAcquireFunction          func()
	// sdlMemoryBarrierReleaseFunction          func()
	// sdlmemset                                func(unsafe.Pointer, int32, uint64) unsafe.Pointer
	// sdlmemset4                               func(unsafe.Pointer, uint32, uint64) unsafe.Pointer
	// sdlMetal_CreateView                      func(*Window) MetalView
	// sdlMetal_DestroyView                     func(MetalView)
	// sdlMetal_GetLayer                        func(MetalView) unsafe.Pointer
	// sdlMinimizeWindow                        func(*Window) bool
	// sdlMixAudio                              func(*uint8, *uint8, AudioFormat, uint32, float32) bool
	// sdlmodf                                  func(float64, *double) float64
	// sdlmodff                                 func(float32, *float32) float32
	// sdlMostSignificantBitIndex32             func(uint32) int32
	// sdlmurmur3_32                            func(unsafe.Pointer, uint64, uint32) uint32
	// sdlOnApplicationDidEnterBackground       func()
	// sdlOnApplicationDidEnterForeground       func()
	// sdlOnApplicationDidReceiveMemoryWarning  func()
	// sdlOnApplicationWillEnterBackground      func()
	// sdlOnApplicationWillEnterForeground      func()
	// sdlOnApplicationWillTerminate            func()
	// sdlOpenAudioDevice                       func(AudioDeviceID, *AudioSpec) AudioDeviceID
	sdlOpenAudioDeviceStream func(AudioDeviceID, *AudioSpec, AudioStreamCallback, unsafe.Pointer) *AudioStream
	sdlOpenCamera            func(CameraID, *CameraSpec) *Camera
	// sdlOpenFileStorage                       func(string) *Storage
	sdlOpenGamepad func(JoystickID) *Gamepad
	// sdlOpenHaptic                            func(HapticID) *Haptic
	// sdlOpenHapticFromJoystick                func(*Joystick) *Haptic
	// sdlOpenHapticFromMouse                   func() *Haptic
	// sdlOpenIO                                func(*IOStreamInterface, unsafe.Pointer) *IOStream
	sdlOpenJoystick func(JoystickID) *Joystick
	// sdlOpenSensor                            func(SensorID) *Sensor
	// sdlOpenStorage                           func(*StorageInterface, unsafe.Pointer) *Storage
	// sdlOpenTitleStorage                      func(string, PropertiesID) *Storage
	sdlOpenURL func(string) bool
	// sdlOpenUserStorage                       func(string, string, PropertiesID) *Storage
	// sdlOutOfMemory                           func() bool
	// sdlPauseAudioDevice                      func(AudioDeviceID) bool
	sdlPauseAudioStreamDevice uintptr
	// sdlPauseHaptic                           func(*Haptic) bool
	sdlPeepEvents func(*Event, int32, EventAction, EventType, EventType) int32
	// sdlPlayHapticRumble                      func(*Haptic, float32, uint32) bool
	sdlPollEvent uintptr
	// sdlPopGPUDebugGroup                      func(*GPUCommandBuffer)
	// sdlpow                                   func(float64, float64) float64
	// sdlpowf                                  func(float32, float32) float32
	// sdlPremultiplyAlpha                      func(int32, int32, PixelFormat, unsafe.Pointer, int32, PixelFormat, unsafe.Pointer, int32, bool) bool
	// sdlPremultiplySurfaceAlpha               func(*Surface, bool) bool
	sdlPumpEvents func()
	sdlPushEvent  func(*Event) bool
	// sdlPushGPUComputeUniformData             func(*GPUCommandBuffer, uint32, unsafe.Pointer, uint32)
	// sdlPushGPUDebugGroup                     func(*GPUCommandBuffer, string)
	sdlPushGPUFragmentUniformData func(*GPUCommandBuffer, uint32, unsafe.Pointer, uint32)
	sdlPushGPUVertexUniformData   func(*GPUCommandBuffer, uint32, unsafe.Pointer, uint32)
	sdlPutAudioStreamData         uintptr
	// sdlqsort                                 func(unsafe.Pointer, uint64, uint64, CompareCallback)
	// sdlqsort_r                               func(unsafe.Pointer, uint64, uint64, CompareCallback_r, unsafe.Pointer)
	// sdlQueryGPUFence                         func(*GPUDevice, *GPUFence) bool
	sdlQuit          func()
	sdlQuitSubSystem func(InitFlags)
	sdlRaiseWindow   func(*Window) bool
	// sdlrand                                  func(int32) int32
	// sdlrand_bits                             func() uint32
	// sdlrand_bits_r                           func(*uint64) uint32
	// sdlrand_r                                func(*uint64, int32) int32
	// sdlrandf                                 func() float32
	// sdlrandf_r                               func(*uint64) float32
	// sdlReadAsyncIO                           func(*AsyncIO, unsafe.Pointer, uint64, uint64, *AsyncIOQueue, unsafe.Pointer) bool
	// sdlReadIO                                func(*IOStream, unsafe.Pointer, uint64) uint64
	// sdlReadProcess                           func(*Process, *uint64, *int32) unsafe.Pointer
	// sdlReadS16BE                             func(*IOStream, *int16) bool
	// sdlReadS16LE                             func(*IOStream, *int16) bool
	// sdlReadS32BE                             func(*IOStream, *int32) bool
	// sdlReadS32LE                             func(*IOStream, *int32) bool
	// sdlReadS64BE                             func(*IOStream, *int64) bool
	// sdlReadS64LE                             func(*IOStream, *int64) bool
	// sdlReadS8                                func(*IOStream, *int8) bool
	// sdlReadStorageFile                       func(*Storage, string, unsafe.Pointer, uint64) bool
	// sdlReadSurfacePixel                      func(*Surface, int32, int32, *uint8, *uint8, *uint8, *uint8) bool
	// sdlReadSurfacePixelFloat                 func(*Surface, int32, int32, *float32, *float32, *float32, *float32) bool
	// sdlReadU16BE                             func(*IOStream, *uint16) bool
	// sdlReadU16LE                             func(*IOStream, *uint16) bool
	// sdlReadU32BE                             func(*IOStream, *uint32) bool
	// sdlReadU32LE                             func(*IOStream, *uint32) bool
	// sdlReadU64BE                             func(*IOStream, *uint64) bool
	// sdlReadU64LE                             func(*IOStream, *uint64) bool
	// sdlReadU8                                func(*IOStream, *uint8) bool
	// sdlrealloc                               func(unsafe.Pointer, uint64) unsafe.Pointer
	sdlRegisterEvents     func(int32) uint32
	sdlReleaseCameraFrame func(*Camera, *Surface)
	sdlReleaseGPUBuffer   func(*GPUDevice, *GPUBuffer)
	// sdlReleaseGPUComputePipeline             func(*GPUDevice, *GPUComputePipeline)
	// sdlReleaseGPUFence                       func(*GPUDevice, *GPUFence)
	sdlReleaseGPUGraphicsPipeline func(*GPUDevice, *GPUGraphicsPipeline)
	sdlReleaseGPUSampler          func(*GPUDevice, *GPUSampler)
	sdlReleaseGPUShader           func(*GPUDevice, *GPUShader)
	sdlReleaseGPUTexture          func(*GPUDevice, *GPUTexture)
	sdlReleaseGPUTransferBuffer   func(*GPUDevice, *GPUTransferBuffer)
	sdlReleaseWindowFromGPUDevice func(*GPUDevice, *Window)
	// sdlReloadGamepadMappings                 func() bool
	sdlRemoveEventWatch   func(EventFilter, unsafe.Pointer)
	sdlRemoveHintCallback func(string, HintCallback, unsafe.Pointer)
	// sdlRemovePath                            func(string) bool
	// sdlRemoveStoragePath                     func(*Storage, string) bool
	sdlRemoveSurfaceAlternateImages func(*Surface)
	// sdlRemoveTimer                           func(TimerID) bool
	// sdlRemoveTrayEntry                       func(*TrayEntry)
	// sdlRenamePath                            func(string, string) bool
	// sdlRenameStoragePath                     func(*Storage, string, string) bool
	sdlRenderClear                 uintptr
	sdlRenderClipEnabled           func(*Renderer) bool
	sdlRenderCoordinatesFromWindow func(*Renderer, float32, float32, *float32, *float32) bool
	sdlRenderCoordinatesToWindow   func(*Renderer, float32, float32, *float32, *float32) bool
	sdlRenderDebugText             func(*Renderer, float32, float32, string) bool
	sdlRenderDebugTextFormat       func(*Renderer, float32, float32, string) bool
	sdlRenderFillRect              uintptr
	sdlRenderFillRects             uintptr
	sdlRenderGeometry              uintptr
	sdlRenderGeometryRaw           uintptr
	sdlRenderLine                  func(*Renderer, float32, float32, float32, float32) bool
	sdlRenderLines                 uintptr
	sdlRenderPoint                 func(*Renderer, float32, float32) bool
	sdlRenderPoints                uintptr
	sdlRenderPresent               uintptr
	sdlRenderReadPixels            func(*Renderer, *Rect) *Surface
	sdlRenderRect                  uintptr
	sdlRenderRects                 uintptr
	sdlRenderTexture               uintptr
	sdlRenderTexture9Grid          func(*Renderer, *Texture, *FRect, float32, float32, float32, float32, float32, *FRect) bool
	sdlRenderTextureAffine         uintptr
	sdlRenderTextureRotated        func(*Renderer, *Texture, *FRect, *FRect, float64, *FPoint, FlipMode) bool
	sdlRenderTextureTiled          func(*Renderer, *Texture, *FRect, float32, *FRect) bool
	sdlRenderViewportSet           func(*Renderer) bool
	// sdlReportAssertion                       func(*AssertData, string, string, int32) AssertState
	// sdlResetAssertionReport                  func()
	sdlResetHint          func(string) bool
	sdlResetHints         func()
	sdlResetKeyboard      func()
	sdlResetLogPriorities func()
	sdlRestoreWindow      func(*Window) bool
	// sdlResumeAudioDevice                     func(AudioDeviceID) bool
	sdlResumeAudioStreamDevice uintptr
	// sdlResumeHaptic                          func(*Haptic) bool
	// sdlround                                 func(float64) float64
	// sdlroundf                                func(float32) float32
	// sdlRumbleGamepad                         func(*Gamepad, uint16, uint16, uint32) bool
	// sdlRumbleGamepadTriggers                 func(*Gamepad, uint16, uint16, uint32) bool
	sdlRumbleJoystick         func(*Joystick, uint16, uint16, uint32) bool
	sdlRumbleJoystickTriggers func(*Joystick, uint16, uint16, uint32) bool
	// sdlRunApp                                func(int32, **byte, main_func, unsafe.Pointer) int32
	// sdlRunHapticEffect                       func(*Haptic, int32, uint32) bool
	// sdlRunOnMainThread                       func(MainThreadCallback, unsafe.Pointer, bool) bool
	sdlSaveBMP   func(*Surface, string) bool
	sdlSaveBMPIO func(*Surface, *IOStream, bool) bool
	// sdlSaveFile                              func(string, unsafe.Pointer, uint64) bool
	// sdlSaveFile_IO                           func(*IOStream, unsafe.Pointer, uint64, bool) bool
	// sdlscalbn                                func(float64, int32) float64
	// sdlscalbnf                               func(float32, int32) float32
	sdlScaleSurface        func(*Surface, int32, int32, ScaleMode) *Surface
	sdlScreenKeyboardShown func(*Window) bool
	// sdlScreenSaverEnabled                    func() bool
	// sdlSeekIO                                func(*IOStream, int64, IOWhence) int64
	// sdlSendGamepadEffect                     func(*Gamepad, unsafe.Pointer, int32) bool
	sdlSendJoystickEffect func(*Joystick, unsafe.Pointer, int32) bool
	// sdlSendJoystickVirtualSensorData func(*Joystick, SensorType, uint64, *float32, int32) bool
	// sdlSetAppMetadata                        func(string, string, string) bool
	// sdlSetAppMetadataProperty                func(string, string) bool
	// sdlSetAssertionHandler                   func(AssertionHandler, unsafe.Pointer)
	// sdlSetAtomicInt                          func(*AtomicInt, int32) int32
	// sdlSetAtomicPointer                      func(*unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	// sdlSetAtomicU32                          func(*AtomicU32, uint32) uint32
	// sdlSetAudioDeviceGain                    func(AudioDeviceID, float32) bool
	// sdlSetAudioPostmixCallback               func(AudioDeviceID, AudioPostmixCallback, unsafe.Pointer) bool
	// sdlSetAudioStreamFormat                  func(*AudioStream, *AudioSpec, *AudioSpec) bool
	// sdlSetAudioStreamFrequencyRatio          func(*AudioStream, float32) bool
	// sdlSetAudioStreamGain                    func(*AudioStream, float32) bool
	// sdlSetAudioStreamGetCallback             func(*AudioStream, AudioStreamCallback, unsafe.Pointer) bool
	// sdlSetAudioStreamInputChannelMap         func(*AudioStream, *int32, int32) bool
	// sdlSetAudioStreamOutputChannelMap        func(*AudioStream, *int32, int32) bool
	// sdlSetAudioStreamPutCallback             func(*AudioStream, AudioStreamCallback, unsafe.Pointer) bool
	sdlSetBooleanProperty func(PropertiesID, string, bool) bool
	// sdlSetClipboardData                      func(ClipboardDataCallback, ClipboardCleanupCallback, unsafe.Pointer, **byte, uint64) bool
	// sdlSetClipboardText                      func(string) bool
	// sdlSetCurrentThreadPriority              func(ThreadPriority) bool
	sdlSetCursor func(*Cursor) bool
	// sdlsetenv_unsafe                         func(string, string, int32) int32
	// sdlSetEnvironmentVariable                func(*Environment, string, string, bool) bool
	sdlSetError func(string) bool
	// sdlSetErrorV                             func(string, va_list) bool
	sdlSetEventEnabled  func(EventType, bool)
	sdlSetEventFilter   func(EventFilter, unsafe.Pointer)
	sdlSetFloatProperty func(PropertiesID, string, float32) bool
	// sdlSetGamepadEventsEnabled               func(bool)
	// sdlSetGamepadLED                         func(*Gamepad, uint8, uint8, uint8) bool
	// sdlSetGamepadMapping                     func(JoystickID, string) bool
	// sdlSetGamepadPlayerIndex                 func(*Gamepad, int32) bool
	// sdlSetGamepadSensorEnabled               func(*Gamepad, SensorType, bool) bool
	// sdlSetGPUAllowedFramesInFlight           func(*GPUDevice, uint32) bool
	// sdlSetGPUBlendConstants                  func(*GPURenderPass, FColor)
	sdlSetGPUBufferName func(*GPUDevice, *GPUBuffer, string)
	sdlSetGPUScissor    func(*GPURenderPass, *Rect)
	// sdlSetGPUStencilReference                func(*GPURenderPass, uint8)
	sdlSetGPUSwapchainParameters func(*GPUDevice, *Window, GPUSwapchainComposition, GPUPresentMode) bool
	// sdlSetGPUTextureName                     func(*GPUDevice, *GPUTexture, string)
	sdlSetGPUViewport func(*GPURenderPass, *GPUViewport)
	// sdlSetHapticAutocenter                   func(*Haptic, int32) bool
	// sdlSetHapticGain                         func(*Haptic, int32) bool
	sdlSetHint             func(string, string) bool
	sdlSetHintWithPriority func(string, string, HintPriority) bool
	// sdlSetInitialized                        func(*InitState, bool)
	sdlSetJoystickEventsEnabled func(bool)
	sdlSetJoystickLED           func(*Joystick, uint8, uint8, uint8) bool
	sdlSetJoystickPlayerIndex   func(*Joystick, int32) bool
	// sdlSetJoystickVirtualAxis                func(*Joystick, int32, int16) bool
	// sdlSetJoystickVirtualBall                func(*Joystick, int32, int16, int16) bool
	// sdlSetJoystickVirtualButton              func(*Joystick, int32, bool) bool
	// sdlSetJoystickVirtualHat                 func(*Joystick, int32, uint8) bool
	// sdlSetJoystickVirtualTouchpad            func(*Joystick, int32, int32, bool, float32, float32, float32) bool
	// sdlSetLinuxThreadPriority                func(int64, int32) bool
	// sdlSetLinuxThreadPriorityAndPolicy       func(int64, int32, int32) bool
	// sdlSetLogOutputFunction                  func(LogOutputFunction, unsafe.Pointer)
	sdlSetLogPriorities func(LogPriority)
	sdlSetLogPriority   func(int32, LogPriority)
	// sdlSetLogPriorityPrefix                  func(LogPriority, string) bool
	// sdlSetMainReady                          func()
	// sdlSetMemoryFunctions                    func(malloc_func, calloc_func, realloc_func, free_func) bool
	sdlSetModState                   func(Keymod)
	sdlSetNumberProperty             func(PropertiesID, string, int64) bool
	sdlSetPaletteColors              func(*Palette, *Color, int32, int32) bool
	sdlSetPointerProperty            func(PropertiesID, string, unsafe.Pointer) bool
	sdlSetPointerPropertyWithCleanup func(PropertiesID, string, unsafe.Pointer, CleanupPropertyCallback, unsafe.Pointer) bool
	// sdlSetPrimarySelectionText               func(string) bool
	sdlSetRenderClipRect            func(*Renderer, *Rect) bool
	sdlSetRenderColorScale          func(*Renderer, float32) bool
	sdlSetRenderDrawBlendMode       func(*Renderer, BlendMode) bool
	sdlSetRenderDrawColor           uintptr
	sdlSetRenderDrawColorFloat      func(*Renderer, float32, float32, float32, float32) bool
	sdlSetRenderLogicalPresentation func(*Renderer, int32, int32, RendererLogicalPresentation) bool
	sdlSetRenderScale               func(*Renderer, float32, float32) bool
	sdlSetRenderTarget              func(*Renderer, *Texture) bool
	sdlSetRenderViewport            func(*Renderer, *Rect) bool
	sdlSetRenderVSync               func(*Renderer, int32) bool
	sdlSetScancodeName              func(Scancode, string) bool
	sdlSetStringProperty            func(PropertiesID, string, string) bool
	sdlSetSurfaceAlphaMod           func(*Surface, uint8) bool
	sdlSetSurfaceBlendMode          func(*Surface, BlendMode) bool
	sdlSetSurfaceClipRect           func(*Surface, *Rect) bool
	sdlSetSurfaceColorKey           func(*Surface, bool, uint32) bool
	sdlSetSurfaceColorMod           func(*Surface, uint8, uint8, uint8) bool
	sdlSetSurfaceColorspace         func(*Surface, Colorspace) bool
	sdlSetSurfacePalette            func(*Surface, *Palette) bool
	sdlSetSurfaceRLE                func(*Surface, bool) bool
	sdlSetTextInputArea             func(*Window, *Rect, int32) bool
	sdlSetTextureAlphaMod           uintptr
	sdlSetTextureAlphaModFloat      func(*Texture, float32) bool
	sdlSetTextureBlendMode          func(*Texture, BlendMode) bool
	sdlSetTextureColorMod           uintptr
	sdlSetTextureColorModFloat      func(*Texture, float32, float32, float32) bool
	sdlSetTextureScaleMode          func(*Texture, ScaleMode) bool
	// sdlSetTLS                                func(*TLSID, unsafe.Pointer, TLSDestructorCallback) bool
	// sdlSetTrayEntryCallback                  func(*TrayEntry, TrayCallback, unsafe.Pointer)
	// sdlSetTrayEntryChecked                   func(*TrayEntry, bool)
	// sdlSetTrayEntryEnabled                   func(*TrayEntry, bool)
	// sdlSetTrayEntryLabel                     func(*TrayEntry, string)
	// sdlSetTrayIcon                           func(*Tray, *Surface)
	// sdlSetTrayTooltip                        func(*Tray, string)
	sdlSetWindowAlwaysOnTop func(*Window, bool) bool
	// sdlSetWindowAspectRatio                  func(*Window, float32, float32) bool
	sdlSetWindowBordered       func(*Window, bool) bool
	sdlSetWindowFocusable      func(*Window, bool) bool
	sdlSetWindowFullscreen     func(*Window, bool) bool
	sdlSetWindowFullscreenMode func(*Window, *DisplayMode) bool
	sdlSetWindowHitTest        func(*Window, func(*Window, *Point, unsafe.Pointer) uintptr, unsafe.Pointer) bool
	sdlSetWindowIcon           func(*Window, *Surface) bool
	sdlSetWindowKeyboardGrab   func(*Window, bool) bool
	// sdlSetWindowMaximumSize                  func(*Window, int32, int32) bool
	// sdlSetWindowMinimumSize                  func(*Window, int32, int32) bool
	// sdlSetWindowModal                        func(*Window, bool) bool
	sdlSetWindowMouseGrab func(*Window, bool) bool
	// sdlSetWindowMouseRect                    func(*Window, *Rect) bool
	sdlSetWindowOpacity func(*Window, float32) bool
	// sdlSetWindowParent                       func(*Window, *Window) bool
	sdlSetWindowPosition          func(*Window, int32, int32) bool
	sdlSetWindowRelativeMouseMode func(*Window, bool) bool
	sdlSetWindowResizable         func(*Window, bool) bool
	// sdlSetWindowShape                        func(*Window, *Surface) bool
	sdlSetWindowSize func(*Window, int32, int32) bool
	// sdlSetWindowSurfaceVSync                 func(*Window, int32) bool
	sdlSetWindowTitle func(*Window, string) bool
	// sdlSetX11EventHook                       func(X11EventHook, unsafe.Pointer)
	// sdlShouldInit                            func(*InitState) bool
	// sdlShouldQuit                            func(*InitState) bool
	sdlShowCursor                   func() bool
	sdlShowFileDialogWithProperties func(FileDialogType, DialogFileCallback, unsafe.Pointer, PropertiesID)
	sdlShowMessageBox               func(*MessageBoxData, *int32) bool
	sdlShowOpenFileDialog           func(DialogFileCallback, unsafe.Pointer, *Window, []DialogFileFilter, int32, *byte, bool)
	sdlShowOpenFolderDialog         func(DialogFileCallback, unsafe.Pointer, *Window, string, bool)
	sdlShowSaveFileDialog           func(DialogFileCallback, unsafe.Pointer, *Window, []DialogFileFilter, int32, string)
	sdlShowSimpleMessageBox         func(MessageBoxFlags, string, string, *Window) bool
	sdlShowWindow                   func(*Window) bool
	// sdlShowWindowSystemMenu                  func(*Window, int32, int32) bool
	// sdlSignalAsyncIOQueue                    func(*AsyncIOQueue)
	// sdlSignalCondition                       func(*Condition)
	// sdlSignalSemaphore                       func(*Semaphore)
	// sdlsin                                   func(float64) float64
	// sdlsinf                                  func(float32) float32
	// sdlsize_add_check_overflow               func(uint64, uint64, *uint64) bool
	// sdlsize_add_check_overflow_builtin       func(uint64, uint64, *uint64) bool
	// sdlsize_mul_check_overflow               func(uint64, uint64, *uint64) bool
	// sdlsize_mul_check_overflow_builtin       func(uint64, uint64, *uint64) bool
	// sdlsnprintf                              func(string, uint64, string) int32
	// sdlsqrt                                  func(float64) float64
	// sdlsqrtf                                 func(float32) float32
	// sdlsrand                                 func(uint64)
	// sdlsscanf                                func(string, string) int32
	sdlStartTextInput               func(*Window) bool
	sdlStartTextInputWithProperties func(*Window, PropertiesID) bool
	// sdlStepBackUTF8                          func(string, **byte) uint32
	// sdlStepUTF8                              func(**byte, *uint64) uint32
	// sdlStopHapticEffect                      func(*Haptic, int32) bool
	// sdlStopHapticEffects                     func(*Haptic) bool
	// sdlStopHapticRumble                      func(*Haptic) bool
	sdlStopTextInput func(*Window) bool
	// sdlStorageReady                          func(*Storage) bool
	// sdlstrcasecmp                            func(string, string) int32
	// sdlstrcasestr                            func(string, string) string
	// sdlstrchr                                func(string, int32) string
	// sdlstrcmp                                func(string, string) int32
	// sdlstrdup                                func(string) string
	// sdlStringToGUID                          func(string) GUID
	// sdlstrlcat                               func(string, string, uint64) uint64
	// sdlstrlcpy                               func(string, string, uint64) uint64
	// sdlstrlen                                func(string) uint64
	// sdlstrlwr                                func(string) string
	// sdlstrncasecmp                           func(string, string, uint64) int32
	// sdlstrncmp                               func(string, string, uint64) int32
	// sdlstrndup                               func(string, uint64) string
	// sdlstrnlen                               func(string, uint64) uint64
	// sdlstrnstr                               func(string, string, uint64) string
	// sdlstrpbrk                               func(string, string) string
	// sdlstrrchr                               func(string, int32) string
	// sdlstrrev                                func(string) string
	sdlstrstr func(string, string) string
	// sdlstrtod                                func(string, **byte) float64
	// sdlstrtok_r                              func(string, string, **byte) string
	// sdlstrtol                                func(string, **byte, int32) int64
	// sdlstrtoll                               func(string, **byte, int32) int64
	// sdlstrtoul                               func(string, **byte, int32) uint64
	// sdlstrtoull                              func(string, **byte, int32) uint64
	// sdlstrupr                                func(string) string
	sdlSubmitGPUCommandBuffer func(*GPUCommandBuffer) bool
	// sdlSubmitGPUCommandBufferAndAcquireFence func(*GPUCommandBuffer) *GPUFence
	sdlSurfaceHasAlternateImages func(*Surface) bool
	sdlSurfaceHasColorKey        func(*Surface) bool
	sdlSurfaceHasRLE             func(*Surface) bool
	// sdlSwapFloat                             func(float32) float32
	// sdlswprintf                              func(*wchar_t, uint64, *wchar_t) int32
	sdlSyncWindow func(*Window) bool
	// sdltan                                   func(float64) float64
	// sdltanf                                  func(float32) float32
	// sdlTellIO                                func(*IOStream) int64
	sdlTextInputActive func(*Window) bool
	// sdlTimeFromWindows                       func(uint32, uint32) Time
	// sdlTimeToDateTime                        func(Time, *DateTime, bool) bool
	// sdlTimeToWindows                         func(Time, *uint32, *uint32)
	// sdltolower                               func(int32) int32
	// sdltoupper                               func(int32) int32
	// sdltrunc                                 func(float64) float64
	// sdltruncf                                func(float32) float32
	// sdlTryLockMutex                          func(*Mutex) bool
	// sdlTryLockRWLockForReading               func(*RWLock) bool
	// sdlTryLockRWLockForWriting               func(*RWLock) bool
	// sdlTryLockSpinlock                       func(*SpinLock) bool
	// sdlTryWaitSemaphore                      func(*Semaphore) bool
	// sdlUCS4ToUTF8                            func(uint32, string) string
	// sdluitoa                                 func(uint32, string, int32) string
	// sdlulltoa                                func(uint64, string, int32) string
	// sdlultoa                                 func(uint64, string, int32) string
	// sdlUnbindAudioStream                     func(*AudioStream)
	// sdlUnbindAudioStreams                    func(**AudioStream, int32)
	// sdlUnloadObject                          func(*SharedObject)
	// sdlUnlockAudioStream                     func(*AudioStream) bool
	sdlUnlockJoysticks func()
	// sdlUnlockMutex                           func(*Mutex)
	sdlUnlockProperties func(PropertiesID)
	// sdlUnlockRWLock                          func(*RWLock)
	// sdlUnlockSpinlock                        func(*SpinLock)
	sdlUnlockSurface          func(*Surface)
	sdlUnlockTexture          func(*Texture)
	sdlUnmapGPUTransferBuffer func(*GPUDevice, *GPUTransferBuffer)
	// sdlunsetenv_unsafe                       func(string) int32
	// sdlUnsetEnvironmentVariable              func(*Environment, string) bool
	// sdlUpdateGamepads                        func()
	// sdlUpdateHapticEffect                    func(*Haptic, int32, *HapticEffect) bool
	sdlUpdateJoysticks func()
	sdlUpdateNVTexture uintptr
	// sdlUpdateSensors                         func()
	sdlUpdateTexture       uintptr
	sdlUpdateWindowSurface func(*Window) bool
	// sdlUpdateWindowSurfaceRects              func(*Window, *Rect, int32) bool
	sdlUpdateYUVTexture   uintptr
	sdlUploadToGPUBuffer  func(*GPUCopyPass, *GPUTransferBufferLocation, *GPUBufferRegion, bool)
	sdlUploadToGPUTexture func(*GPUCopyPass, *GPUTextureTransferInfo, *GPUTextureRegion, bool)
	// sdlutf8strlcpy                           func(string, string, uint64) uint64
	// sdlutf8strlen                            func(string) uint64
	// sdlutf8strnlen                           func(string, uint64) uint64
	// sdlvasprintf                             func(**byte, string, va_list) int32
	// sdlvsnprintf                             func(string, uint64, string, va_list) int32
	// sdlvsscanf                               func(string, string, va_list) int32
	// sdlvswprintf                             func(*wchar_t, uint64, *wchar_t, va_list) int32
	sdlWaitAndAcquireGPUSwapchainTexture func(*GPUCommandBuffer, *Window, **GPUTexture, *uint32, *uint32) bool
	// sdlWaitAsyncIOResult                     func(*AsyncIOQueue, *AsyncIOOutcome, int32) bool
	// sdlWaitCondition                         func(*Condition, *Mutex)
	// sdlWaitConditionTimeout                  func(*Condition, *Mutex, int32) bool
	sdlWaitEvent        func(*Event) bool
	sdlWaitEventTimeout func(*Event, int32) bool
	// sdlWaitForGPUFences                      func(*GPUDevice, bool, **GPUFence, uint32) bool
	// sdlWaitForGPUIdle                        func(*GPUDevice) bool
	// sdlWaitForGPUSwapchain                   func(*GPUDevice, *Window) bool
	// sdlWaitProcess                           func(*Process, bool, *int32) bool
	// sdlWaitSemaphore                         func(*Semaphore)
	// sdlWaitSemaphoreTimeout                  func(*Semaphore, int32) bool
	// sdlWaitThread                            func(*Thread, *int32)
	sdlWarpMouseGlobal   func(float32, float32) bool
	sdlWarpMouseInWindow func(*Window, float32, float32)
	// sdlWasInit                               func(InitFlags) InitFlags
	// sdlwcscasecmp                            func(*wchar_t, *wchar_t) int32
	// sdlwcscmp                                func(*wchar_t, *wchar_t) int32
	// sdlwcsdup                                func(*wchar_t) *wchar_t
	// sdlwcslcat                               func(*wchar_t, *wchar_t, uint64) uint64
	// sdlwcslcpy                               func(*wchar_t, *wchar_t, uint64) uint64
	// sdlwcslen                                func(*wchar_t) uint64
	// sdlwcsncasecmp                           func(*wchar_t, *wchar_t, uint64) int32
	// sdlwcsncmp                               func(*wchar_t, *wchar_t, uint64) int32
	// sdlwcsnlen                               func(*wchar_t, uint64) uint64
	// sdlwcsnstr                               func(*wchar_t, *wchar_t, uint64) *wchar_t
	// sdlwcsstr                                func(*wchar_t, *wchar_t) *wchar_t
	// sdlwcstol                                func(*wchar_t, **wchar_t, int32) int64
	// sdlWindowHasSurface                      func(*Window) bool
	sdlWindowSupportsGPUPresentMode func(*GPUDevice, *Window, GPUPresentMode) bool
	// sdlWindowSupportsGPUSwapchainComposition func(*GPUDevice, *Window, GPUSwapchainComposition) bool
	// sdlWriteAsyncIO                          func(*AsyncIO, unsafe.Pointer, uint64, uint64, *AsyncIOQueue, unsafe.Pointer) bool
	// sdlWriteIO                               func(*IOStream, unsafe.Pointer, uint64) uint64
	// sdlWriteS16BE                            func(*IOStream, int16) bool
	// sdlWriteS16LE                            func(*IOStream, int16) bool
	// sdlWriteS32BE                            func(*IOStream, int32) bool
	// sdlWriteS32LE                            func(*IOStream, int32) bool
	// sdlWriteS64BE                            func(*IOStream, int64) bool
	// sdlWriteS64LE                            func(*IOStream, int64) bool
	// sdlWriteS8                               func(*IOStream, int8) bool
	// sdlWriteStorageFile                      func(*Storage, string, unsafe.Pointer, uint64) bool
	// sdlWriteSurfacePixel                     func(*Surface, int32, int32, uint8, uint8, uint8, uint8) bool
	// sdlWriteSurfacePixelFloat                func(*Surface, int32, int32, float32, float32, float32, float32) bool
	// sdlWriteU16BE                            func(*IOStream, uint16) bool
	// sdlWriteU16LE                            func(*IOStream, uint16) bool
	// sdlWriteU32BE                            func(*IOStream, uint32) bool
	// sdlWriteU32LE                            func(*IOStream, uint32) bool
	// sdlWriteU64BE                            func(*IOStream, uint64) bool
	// sdlWriteU64LE                            func(*IOStream, uint64) bool
	// sdlWriteU8 func(*IOStream, uint8) bool
)

func init() {
	runtime.LockOSThread()

	var filename string
	switch runtime.GOOS {
	case "linux", "freebsd":
		filename = "libSDL3.so.0"
	case "windows":
		filename = "SDL3.dll"
	case "darwin":
		filename = "libSDL3.dylib"
	}

	lib, err := shared.Load(filename)
	if err != nil {
		panic(err)
	}

	// purego.RegisterLibFunc(&sdlabs, lib, "SDL_abs")
	// purego.RegisterLibFunc(&sdlacos, lib, "SDL_acos")
	// purego.RegisterLibFunc(&sdlacosf, lib, "SDL_acosf")
	purego.RegisterLibFunc(&sdlAcquireCameraFrame, lib, "SDL_AcquireCameraFrame")
	purego.RegisterLibFunc(&sdlAcquireGPUCommandBuffer, lib, "SDL_AcquireGPUCommandBuffer")
	purego.RegisterLibFunc(&sdlAcquireGPUSwapchainTexture, lib, "SDL_AcquireGPUSwapchainTexture")
	// purego.RegisterLibFunc(&sdlAddAtomicInt, lib, "SDL_AddAtomicInt")
	purego.RegisterLibFunc(&sdlAddEventWatch, lib, "SDL_AddEventWatch")
	// purego.RegisterLibFunc(&sdlAddGamepadMapping, lib, "SDL_AddGamepadMapping")
	// purego.RegisterLibFunc(&sdlAddGamepadMappingsFromFile, lib, "SDL_AddGamepadMappingsFromFile")
	// purego.RegisterLibFunc(&sdlAddGamepadMappingsFromIO, lib, "SDL_AddGamepadMappingsFromIO")
	purego.RegisterLibFunc(&sdlAddHintCallback, lib, "SDL_AddHintCallback")
	purego.RegisterLibFunc(&sdlAddSurfaceAlternateImage, lib, "SDL_AddSurfaceAlternateImage")
	// purego.RegisterLibFunc(&sdlAddTimer, lib, "SDL_AddTimer")
	// purego.RegisterLibFunc(&sdlAddTimerNS, lib, "SDL_AddTimerNS")
	purego.RegisterLibFunc(&sdlAddVulkanRenderSemaphores, lib, "SDL_AddVulkanRenderSemaphores")
	// purego.RegisterLibFunc(&sdlaligned_alloc, lib, "SDL_aligned_alloc")
	// purego.RegisterLibFunc(&sdlaligned_free, lib, "SDL_aligned_free")
	// purego.RegisterLibFunc(&sdlasin, lib, "SDL_asin")
	// purego.RegisterLibFunc(&sdlasinf, lib, "SDL_asinf")
	// purego.RegisterLibFunc(&sdlasprintf, lib, "SDL_asprintf")
	// purego.RegisterLibFunc(&sdlAsyncIOFromFile, lib, "SDL_AsyncIOFromFile")
	// purego.RegisterLibFunc(&sdlatan, lib, "SDL_atan")
	// purego.RegisterLibFunc(&sdlatan2, lib, "SDL_atan2")
	// purego.RegisterLibFunc(&sdlatan2f, lib, "SDL_atan2f")
	// purego.RegisterLibFunc(&sdlatanf, lib, "SDL_atanf")
	// purego.RegisterLibFunc(&sdlatof, lib, "SDL_atof")
	// purego.RegisterLibFunc(&sdlatoi, lib, "SDL_atoi")
	// purego.RegisterLibFunc(&sdlAttachVirtualJoystick, lib, "SDL_AttachVirtualJoystick")
	// purego.RegisterLibFunc(&sdlAudioDevicePaused, lib, "SDL_AudioDevicePaused")
	sdlAudioStreamDevicePaused = shared.Get(lib, "SDL_AudioStreamDevicePaused")
	// purego.RegisterLibFunc(&sdlBeginGPUComputePass, lib, "SDL_BeginGPUComputePass")
	purego.RegisterLibFunc(&sdlBeginGPUCopyPass, lib, "SDL_BeginGPUCopyPass")
	purego.RegisterLibFunc(&sdlBeginGPURenderPass, lib, "SDL_BeginGPURenderPass")
	// purego.RegisterLibFunc(&sdlBindAudioStream, lib, "SDL_BindAudioStream")
	// purego.RegisterLibFunc(&sdlBindAudioStreams, lib, "SDL_BindAudioStreams")
	// purego.RegisterLibFunc(&sdlBindGPUComputePipeline, lib, "SDL_BindGPUComputePipeline")
	// purego.RegisterLibFunc(&sdlBindGPUComputeSamplers, lib, "SDL_BindGPUComputeSamplers")
	// purego.RegisterLibFunc(&sdlBindGPUComputeStorageBuffers, lib, "SDL_BindGPUComputeStorageBuffers")
	// purego.RegisterLibFunc(&sdlBindGPUComputeStorageTextures, lib, "SDL_BindGPUComputeStorageTextures")
	purego.RegisterLibFunc(&sdlBindGPUFragmentSamplers, lib, "SDL_BindGPUFragmentSamplers")
	// purego.RegisterLibFunc(&sdlBindGPUFragmentStorageBuffers, lib, "SDL_BindGPUFragmentStorageBuffers")
	// purego.RegisterLibFunc(&sdlBindGPUFragmentStorageTextures, lib, "SDL_BindGPUFragmentStorageTextures")
	purego.RegisterLibFunc(&sdlBindGPUGraphicsPipeline, lib, "SDL_BindGPUGraphicsPipeline")
	purego.RegisterLibFunc(&sdlBindGPUIndexBuffer, lib, "SDL_BindGPUIndexBuffer")
	purego.RegisterLibFunc(&sdlBindGPUVertexBuffers, lib, "SDL_BindGPUVertexBuffers")
	// purego.RegisterLibFunc(&sdlBindGPUVertexSamplers, lib, "SDL_BindGPUVertexSamplers")
	// purego.RegisterLibFunc(&sdlBindGPUVertexStorageBuffers, lib, "SDL_BindGPUVertexStorageBuffers")
	// purego.RegisterLibFunc(&sdlBindGPUVertexStorageTextures, lib, "SDL_BindGPUVertexStorageTextures")
	// purego.RegisterLibFunc(&sdlBlitGPUTexture, lib, "SDL_BlitGPUTexture")
	purego.RegisterLibFunc(&sdlBlitSurface, lib, "SDL_BlitSurface")
	// purego.RegisterLibFunc(&sdlBlitSurface9Grid, lib, "SDL_BlitSurface9Grid")
	// purego.RegisterLibFunc(&sdlBlitSurfaceScaled, lib, "SDL_BlitSurfaceScaled")
	// purego.RegisterLibFunc(&sdlBlitSurfaceTiled, lib, "SDL_BlitSurfaceTiled")
	// purego.RegisterLibFunc(&sdlBlitSurfaceTiledWithScale, lib, "SDL_BlitSurfaceTiledWithScale")
	// purego.RegisterLibFunc(&sdlBlitSurfaceUnchecked, lib, "SDL_BlitSurfaceUnchecked")
	// purego.RegisterLibFunc(&sdlBlitSurfaceUncheckedScaled, lib, "SDL_BlitSurfaceUncheckedScaled")
	// purego.RegisterLibFunc(&sdlBroadcastCondition, lib, "SDL_BroadcastCondition")
	// purego.RegisterLibFunc(&sdlbsearch, lib, "SDL_bsearch")
	// purego.RegisterLibFunc(&sdlbsearch_r, lib, "SDL_bsearch_r")
	// purego.RegisterLibFunc(&sdlCalculateGPUTextureFormatSize, lib, "SDL_CalculateGPUTextureFormatSize")
	// purego.RegisterLibFunc(&sdlcalloc, lib, "SDL_calloc")
	// purego.RegisterLibFunc(&sdlCancelGPUCommandBuffer, lib, "SDL_CancelGPUCommandBuffer")
	purego.RegisterLibFunc(&sdlCaptureMouse, lib, "SDL_CaptureMouse")
	// purego.RegisterLibFunc(&sdlceil, lib, "SDL_ceil")
	// purego.RegisterLibFunc(&sdlceilf, lib, "SDL_ceilf")
	purego.RegisterLibFunc(&sdlClaimWindowForGPUDevice, lib, "SDL_ClaimWindowForGPUDevice")
	// purego.RegisterLibFunc(&sdlCleanupTLS, lib, "SDL_CleanupTLS")
	sdlClearAudioStream = shared.Get(lib, "SDL_ClearAudioStream")
	// purego.RegisterLibFunc(&sdlClearClipboardData, lib, "SDL_ClearClipboardData")
	purego.RegisterLibFunc(&sdlClearComposition, lib, "SDL_ClearComposition")
	purego.RegisterLibFunc(&sdlClearError, lib, "SDL_ClearError")
	purego.RegisterLibFunc(&sdlClearProperty, lib, "SDL_ClearProperty")
	// purego.RegisterLibFunc(&sdlClearSurface, lib, "SDL_ClearSurface")
	// purego.RegisterLibFunc(&sdlClickTrayEntry, lib, "SDL_ClickTrayEntry")
	// purego.RegisterLibFunc(&sdlCloseAsyncIO, lib, "SDL_CloseAsyncIO")
	// purego.RegisterLibFunc(&sdlCloseAudioDevice, lib, "SDL_CloseAudioDevice")
	purego.RegisterLibFunc(&sdlCloseCamera, lib, "SDL_CloseCamera")
	purego.RegisterLibFunc(&sdlCloseGamepad, lib, "SDL_CloseGamepad")
	// purego.RegisterLibFunc(&sdlCloseHaptic, lib, "SDL_CloseHaptic")
	purego.RegisterLibFunc(&sdlCloseIO, lib, "SDL_CloseIO")
	purego.RegisterLibFunc(&sdlCloseJoystick, lib, "SDL_CloseJoystick")
	// purego.RegisterLibFunc(&sdlCloseSensor, lib, "SDL_CloseSensor")
	// purego.RegisterLibFunc(&sdlCloseStorage, lib, "SDL_CloseStorage")
	// purego.RegisterLibFunc(&sdlCompareAndSwapAtomicInt, lib, "SDL_CompareAndSwapAtomicInt")
	// purego.RegisterLibFunc(&sdlCompareAndSwapAtomicPointer, lib, "SDL_CompareAndSwapAtomicPointer")
	// purego.RegisterLibFunc(&sdlCompareAndSwapAtomicU32, lib, "SDL_CompareAndSwapAtomicU32")
	// purego.RegisterLibFunc(&sdlComposeCustomBlendMode, lib, "SDL_ComposeCustomBlendMode")
	// purego.RegisterLibFunc(&sdlConvertAudioSamples, lib, "SDL_ConvertAudioSamples")
	purego.RegisterLibFunc(&sdlConvertEventToRenderCoordinates, lib, "SDL_ConvertEventToRenderCoordinates")
	purego.RegisterLibFunc(&sdlConvertPixels, lib, "SDL_ConvertPixels")
	purego.RegisterLibFunc(&sdlConvertPixelsAndColorspace, lib, "SDL_ConvertPixelsAndColorspace")
	purego.RegisterLibFunc(&sdlConvertSurface, lib, "SDL_ConvertSurface")
	purego.RegisterLibFunc(&sdlConvertSurfaceAndColorspace, lib, "SDL_ConvertSurfaceAndColorspace")
	// purego.RegisterLibFunc(&sdlCopyFile, lib, "SDL_CopyFile")
	// purego.RegisterLibFunc(&sdlCopyGPUBufferToBuffer, lib, "SDL_CopyGPUBufferToBuffer")
	// purego.RegisterLibFunc(&sdlCopyGPUTextureToTexture, lib, "SDL_CopyGPUTextureToTexture")
	purego.RegisterLibFunc(&sdlCopyProperties, lib, "SDL_CopyProperties")
	// purego.RegisterLibFunc(&sdlcopysign, lib, "SDL_copysign")
	// purego.RegisterLibFunc(&sdlcopysignf, lib, "SDL_copysignf")
	// purego.RegisterLibFunc(&sdlCopyStorageFile, lib, "SDL_CopyStorageFile")
	// purego.RegisterLibFunc(&sdlcos, lib, "SDL_cos")
	// purego.RegisterLibFunc(&sdlcosf, lib, "SDL_cosf")
	// purego.RegisterLibFunc(&sdlcrc16, lib, "SDL_crc16")
	// purego.RegisterLibFunc(&sdlcrc32, lib, "SDL_crc32")
	// purego.RegisterLibFunc(&sdlCreateAsyncIOQueue, lib, "SDL_CreateAsyncIOQueue")
	// purego.RegisterLibFunc(&sdlCreateAudioStream, lib, "SDL_CreateAudioStream")
	purego.RegisterLibFunc(&sdlCreateColorCursor, lib, "SDL_CreateColorCursor")
	// purego.RegisterLibFunc(&sdlCreateCondition, lib, "SDL_CreateCondition")
	purego.RegisterLibFunc(&sdlCreateCursor, lib, "SDL_CreateCursor")
	// purego.RegisterLibFunc(&sdlCreateDirectory, lib, "SDL_CreateDirectory")
	// purego.RegisterLibFunc(&sdlCreateEnvironment, lib, "SDL_CreateEnvironment")
	purego.RegisterLibFunc(&sdlCreateGPUBuffer, lib, "SDL_CreateGPUBuffer")
	// purego.RegisterLibFunc(&sdlCreateGPUComputePipeline, lib, "SDL_CreateGPUComputePipeline")
	purego.RegisterLibFunc(&sdlCreateGPUDevice, lib, "SDL_CreateGPUDevice")
	// purego.RegisterLibFunc(&sdlCreateGPUDeviceWithProperties, lib, "SDL_CreateGPUDeviceWithProperties")
	purego.RegisterLibFunc(&sdlCreateGPUGraphicsPipeline, lib, "SDL_CreateGPUGraphicsPipeline")
	purego.RegisterLibFunc(&sdlCreateGPUSampler, lib, "SDL_CreateGPUSampler")
	purego.RegisterLibFunc(&sdlCreateGPUShader, lib, "SDL_CreateGPUShader")
	purego.RegisterLibFunc(&sdlCreateGPUTexture, lib, "SDL_CreateGPUTexture")
	purego.RegisterLibFunc(&sdlCreateGPUTransferBuffer, lib, "SDL_CreateGPUTransferBuffer")
	// purego.RegisterLibFunc(&sdlCreateHapticEffect, lib, "SDL_CreateHapticEffect")
	// purego.RegisterLibFunc(&sdlCreateMutex, lib, "SDL_CreateMutex")
	purego.RegisterLibFunc(&sdlCreatePalette, lib, "SDL_CreatePalette")
	// purego.RegisterLibFunc(&sdlCreatePopupWindow, lib, "SDL_CreatePopupWindow")
	// purego.RegisterLibFunc(&sdlCreateProcess, lib, "SDL_CreateProcess")
	// purego.RegisterLibFunc(&sdlCreateProcessWithProperties, lib, "SDL_CreateProcessWithProperties")
	purego.RegisterLibFunc(&sdlCreateProperties, lib, "SDL_CreateProperties")
	purego.RegisterLibFunc(&sdlCreateRenderer, lib, "SDL_CreateRenderer")
	purego.RegisterLibFunc(&sdlCreateRendererWithProperties, lib, "SDL_CreateRendererWithProperties")
	// purego.RegisterLibFunc(&sdlCreateRWLock, lib, "SDL_CreateRWLock")
	// purego.RegisterLibFunc(&sdlCreateSemaphore, lib, "SDL_CreateSemaphore")
	purego.RegisterLibFunc(&sdlCreateSoftwareRenderer, lib, "SDL_CreateSoftwareRenderer")
	// purego.RegisterLibFunc(&sdlCreateStorageDirectory, lib, "SDL_CreateStorageDirectory")
	purego.RegisterLibFunc(&sdlCreateSurface, lib, "SDL_CreateSurface")
	purego.RegisterLibFunc(&sdlCreateSurfaceFrom, lib, "SDL_CreateSurfaceFrom")
	purego.RegisterLibFunc(&sdlCreateSurfacePalette, lib, "SDL_CreateSurfacePalette")
	purego.RegisterLibFunc(&sdlCreateSystemCursor, lib, "SDL_CreateSystemCursor")
	purego.RegisterLibFunc(&sdlCreateTexture, lib, "SDL_CreateTexture")
	purego.RegisterLibFunc(&sdlCreateTextureFromSurface, lib, "SDL_CreateTextureFromSurface")
	purego.RegisterLibFunc(&sdlCreateTextureWithProperties, lib, "SDL_CreateTextureWithProperties")
	// purego.RegisterLibFunc(&sdlCreateThreadRuntime, lib, "SDL_CreateThreadRuntime")
	// purego.RegisterLibFunc(&sdlCreateThreadWithPropertiesRuntime, lib, "SDL_CreateThreadWithPropertiesRuntime")
	// purego.RegisterLibFunc(&sdlCreateTray, lib, "SDL_CreateTray")
	// purego.RegisterLibFunc(&sdlCreateTrayMenu, lib, "SDL_CreateTrayMenu")
	// purego.RegisterLibFunc(&sdlCreateTraySubmenu, lib, "SDL_CreateTraySubmenu")
	purego.RegisterLibFunc(&sdlCreateWindow, lib, "SDL_CreateWindow")
	purego.RegisterLibFunc(&sdlCreateWindowAndRenderer, lib, "SDL_CreateWindowAndRenderer")
	// purego.RegisterLibFunc(&sdlCreateWindowWithProperties, lib, "SDL_CreateWindowWithProperties")
	purego.RegisterLibFunc(&sdlCursorVisible, lib, "SDL_CursorVisible")
	// purego.RegisterLibFunc(&sdlDateTimeToTime, lib, "SDL_DateTimeToTime")
	// purego.RegisterLibFunc(&sdlDelay, lib, "SDL_Delay")
	purego.RegisterLibFunc(&sdlDelayNS, lib, "SDL_DelayNS")
	// purego.RegisterLibFunc(&sdlDelayPrecise, lib, "SDL_DelayPrecise")
	// purego.RegisterLibFunc(&sdlDestroyAsyncIOQueue, lib, "SDL_DestroyAsyncIOQueue")
	purego.RegisterLibFunc(&sdlDestroyAudioStream, lib, "SDL_DestroyAudioStream")
	// purego.RegisterLibFunc(&sdlDestroyCondition, lib, "SDL_DestroyCondition")
	purego.RegisterLibFunc(&sdlDestroyCursor, lib, "SDL_DestroyCursor")
	// purego.RegisterLibFunc(&sdlDestroyEnvironment, lib, "SDL_DestroyEnvironment")
	purego.RegisterLibFunc(&sdlDestroyGPUDevice, lib, "SDL_DestroyGPUDevice")
	// purego.RegisterLibFunc(&sdlDestroyHapticEffect, lib, "SDL_DestroyHapticEffect")
	// purego.RegisterLibFunc(&sdlDestroyMutex, lib, "SDL_DestroyMutex")
	purego.RegisterLibFunc(&sdlDestroyPalette, lib, "SDL_DestroyPalette")
	// purego.RegisterLibFunc(&sdlDestroyProcess, lib, "SDL_DestroyProcess")
	purego.RegisterLibFunc(&sdlDestroyProperties, lib, "SDL_DestroyProperties")
	purego.RegisterLibFunc(&sdlDestroyRenderer, lib, "SDL_DestroyRenderer")
	// purego.RegisterLibFunc(&sdlDestroyRWLock, lib, "SDL_DestroyRWLock")
	// purego.RegisterLibFunc(&sdlDestroySemaphore, lib, "SDL_DestroySemaphore")
	purego.RegisterLibFunc(&sdlDestroySurface, lib, "SDL_DestroySurface")
	purego.RegisterLibFunc(&sdlDestroyTexture, lib, "SDL_DestroyTexture")
	// purego.RegisterLibFunc(&sdlDestroyTray, lib, "SDL_DestroyTray")
	purego.RegisterLibFunc(&sdlDestroyWindow, lib, "SDL_DestroyWindow")
	// purego.RegisterLibFunc(&sdlDestroyWindowSurface, lib, "SDL_DestroyWindowSurface")
	// purego.RegisterLibFunc(&sdlDetachThread, lib, "SDL_DetachThread")
	// purego.RegisterLibFunc(&sdlDetachVirtualJoystick, lib, "SDL_DetachVirtualJoystick")
	// purego.RegisterLibFunc(&sdlDisableScreenSaver, lib, "SDL_DisableScreenSaver")
	// purego.RegisterLibFunc(&sdlDispatchGPUCompute, lib, "SDL_DispatchGPUCompute")
	// purego.RegisterLibFunc(&sdlDispatchGPUComputeIndirect, lib, "SDL_DispatchGPUComputeIndirect")
	// purego.RegisterLibFunc(&sdlDownloadFromGPUBuffer, lib, "SDL_DownloadFromGPUBuffer")
	// purego.RegisterLibFunc(&sdlDownloadFromGPUTexture, lib, "SDL_DownloadFromGPUTexture")
	purego.RegisterLibFunc(&sdlDrawGPUIndexedPrimitives, lib, "SDL_DrawGPUIndexedPrimitives")
	// purego.RegisterLibFunc(&sdlDrawGPUIndexedPrimitivesIndirect, lib, "SDL_DrawGPUIndexedPrimitivesIndirect")
	purego.RegisterLibFunc(&sdlDrawGPUPrimitives, lib, "SDL_DrawGPUPrimitives")
	// purego.RegisterLibFunc(&sdlDrawGPUPrimitivesIndirect, lib, "SDL_DrawGPUPrimitivesIndirect")
	purego.RegisterLibFunc(&sdlDuplicateSurface, lib, "SDL_DuplicateSurface")
	// purego.RegisterLibFunc(&sdlEGL_GetCurrentConfig, lib, "SDL_EGL_GetCurrentConfig")
	// purego.RegisterLibFunc(&sdlEGL_GetCurrentDisplay, lib, "SDL_EGL_GetCurrentDisplay")
	// purego.RegisterLibFunc(&sdlEGL_GetProcAddress, lib, "SDL_EGL_GetProcAddress")
	// purego.RegisterLibFunc(&sdlEGL_GetWindowSurface, lib, "SDL_EGL_GetWindowSurface")
	// purego.RegisterLibFunc(&sdlEGL_SetAttributeCallbacks, lib, "SDL_EGL_SetAttributeCallbacks")
	// purego.RegisterLibFunc(&sdlEnableScreenSaver, lib, "SDL_EnableScreenSaver")
	// purego.RegisterLibFunc(&sdlEndGPUComputePass, lib, "SDL_EndGPUComputePass")
	purego.RegisterLibFunc(&sdlEndGPUCopyPass, lib, "SDL_EndGPUCopyPass")
	purego.RegisterLibFunc(&sdlEndGPURenderPass, lib, "SDL_EndGPURenderPass")
	// purego.RegisterLibFunc(&sdlEnterAppMainCallbacks, lib, "SDL_EnterAppMainCallbacks")
	// purego.RegisterLibFunc(&sdlEnumerateDirectory, lib, "SDL_EnumerateDirectory")
	purego.RegisterLibFunc(&sdlEnumerateProperties, lib, "SDL_EnumerateProperties")
	// purego.RegisterLibFunc(&sdlEnumerateStorageDirectory, lib, "SDL_EnumerateStorageDirectory")
	purego.RegisterLibFunc(&sdlEventEnabled, lib, "SDL_EventEnabled")
	// purego.RegisterLibFunc(&sdlexp, lib, "SDL_exp")
	// purego.RegisterLibFunc(&sdlexpf, lib, "SDL_expf")
	// purego.RegisterLibFunc(&sdlfabs, lib, "SDL_fabs")
	// purego.RegisterLibFunc(&sdlfabsf, lib, "SDL_fabsf")
	purego.RegisterLibFunc(&sdlFillSurfaceRect, lib, "SDL_FillSurfaceRect")
	// purego.RegisterLibFunc(&sdlFillSurfaceRects, lib, "SDL_FillSurfaceRects")
	purego.RegisterLibFunc(&sdlFilterEvents, lib, "SDL_FilterEvents")
	purego.RegisterLibFunc(&sdlFlashWindow, lib, "SDL_FlashWindow")
	purego.RegisterLibFunc(&sdlFlipSurface, lib, "SDL_FlipSurface")
	// purego.RegisterLibFunc(&sdlfloor, lib, "SDL_floor")
	// purego.RegisterLibFunc(&sdlfloorf, lib, "SDL_floorf")
	sdlFlushAudioStream = shared.Get(lib, "SDL_FlushAudioStream")
	purego.RegisterLibFunc(&sdlFlushEvent, lib, "SDL_FlushEvent")
	purego.RegisterLibFunc(&sdlFlushEvents, lib, "SDL_FlushEvents")
	// purego.RegisterLibFunc(&sdlFlushIO, lib, "SDL_FlushIO")
	sdlFlushRenderer = shared.Get(lib, "SDL_FlushRenderer")
	// purego.RegisterLibFunc(&sdlfmod, lib, "SDL_fmod")
	// purego.RegisterLibFunc(&sdlfmodf, lib, "SDL_fmodf")
	sdlfree = shared.Get(lib, "SDL_free")
	// purego.RegisterLibFunc(&sdlGamepadConnected, lib, "SDL_GamepadConnected")
	// purego.RegisterLibFunc(&sdlGamepadEventsEnabled, lib, "SDL_GamepadEventsEnabled")
	// purego.RegisterLibFunc(&sdlGamepadHasAxis, lib, "SDL_GamepadHasAxis")
	// purego.RegisterLibFunc(&sdlGamepadHasButton, lib, "SDL_GamepadHasButton")
	// purego.RegisterLibFunc(&sdlGamepadHasSensor, lib, "SDL_GamepadHasSensor")
	// purego.RegisterLibFunc(&sdlGamepadSensorEnabled, lib, "SDL_GamepadSensorEnabled")
	// purego.RegisterLibFunc(&sdlGDKSuspendComplete, lib, "SDL_GDKSuspendComplete")
	// purego.RegisterLibFunc(&sdlGenerateMipmapsForGPUTexture, lib, "SDL_GenerateMipmapsForGPUTexture")
	purego.RegisterLibFunc(&sdlGetAppMetadataProperty, lib, "SDL_GetAppMetadataProperty")
	// purego.RegisterLibFunc(&sdlGetAssertionHandler, lib, "SDL_GetAssertionHandler")
	// purego.RegisterLibFunc(&sdlGetAssertionReport, lib, "SDL_GetAssertionReport")
	// purego.RegisterLibFunc(&sdlGetAsyncIOResult, lib, "SDL_GetAsyncIOResult")
	// purego.RegisterLibFunc(&sdlGetAsyncIOSize, lib, "SDL_GetAsyncIOSize")
	// purego.RegisterLibFunc(&sdlGetAtomicInt, lib, "SDL_GetAtomicInt")
	// purego.RegisterLibFunc(&sdlGetAtomicPointer, lib, "SDL_GetAtomicPointer")
	// purego.RegisterLibFunc(&sdlGetAtomicU32, lib, "SDL_GetAtomicU32")
	// purego.RegisterLibFunc(&sdlGetAudioDeviceChannelMap, lib, "SDL_GetAudioDeviceChannelMap")
	// purego.RegisterLibFunc(&sdlGetAudioDeviceFormat, lib, "SDL_GetAudioDeviceFormat")
	// purego.RegisterLibFunc(&sdlGetAudioDeviceGain, lib, "SDL_GetAudioDeviceGain")
	// purego.RegisterLibFunc(&sdlGetAudioDeviceName, lib, "SDL_GetAudioDeviceName")
	purego.RegisterLibFunc(&sdlGetAudioDriver, lib, "SDL_GetAudioDriver")
	// purego.RegisterLibFunc(&sdlGetAudioFormatName, lib, "SDL_GetAudioFormatName")
	// purego.RegisterLibFunc(&sdlGetAudioPlaybackDevices, lib, "SDL_GetAudioPlaybackDevices")
	// purego.RegisterLibFunc(&sdlGetAudioRecordingDevices, lib, "SDL_GetAudioRecordingDevices")
	// purego.RegisterLibFunc(&sdlGetAudioStreamAvailable, lib, "SDL_GetAudioStreamAvailable")
	// purego.RegisterLibFunc(&sdlGetAudioStreamData, lib, "SDL_GetAudioStreamData")
	// purego.RegisterLibFunc(&sdlGetAudioStreamDevice, lib, "SDL_GetAudioStreamDevice")
	// purego.RegisterLibFunc(&sdlGetAudioStreamFormat, lib, "SDL_GetAudioStreamFormat")
	// purego.RegisterLibFunc(&sdlGetAudioStreamFrequencyRatio, lib, "SDL_GetAudioStreamFrequencyRatio")
	// purego.RegisterLibFunc(&sdlGetAudioStreamGain, lib, "SDL_GetAudioStreamGain")
	// purego.RegisterLibFunc(&sdlGetAudioStreamInputChannelMap, lib, "SDL_GetAudioStreamInputChannelMap")
	// purego.RegisterLibFunc(&sdlGetAudioStreamOutputChannelMap, lib, "SDL_GetAudioStreamOutputChannelMap")
	// purego.RegisterLibFunc(&sdlGetAudioStreamProperties, lib, "SDL_GetAudioStreamProperties")
	sdlGetAudioStreamQueued = shared.Get(lib, "SDL_GetAudioStreamQueued")
	purego.RegisterLibFunc(&sdlGetBasePath, lib, "SDL_GetBasePath")
	purego.RegisterLibFunc(&sdlGetBooleanProperty, lib, "SDL_GetBooleanProperty")
	purego.RegisterLibFunc(&sdlGetCameraDriver, lib, "SDL_GetCameraDriver")
	purego.RegisterLibFunc(&sdlGetCameraFormat, lib, "SDL_GetCameraFormat")
	purego.RegisterLibFunc(&sdlGetCameraID, lib, "SDL_GetCameraID")
	purego.RegisterLibFunc(&sdlGetCameraName, lib, "SDL_GetCameraName")
	purego.RegisterLibFunc(&sdlGetCameraPermissionState, lib, "SDL_GetCameraPermissionState")
	purego.RegisterLibFunc(&sdlGetCameraPosition, lib, "SDL_GetCameraPosition")
	purego.RegisterLibFunc(&sdlGetCameraProperties, lib, "SDL_GetCameraProperties")
	purego.RegisterLibFunc(&sdlGetCameras, lib, "SDL_GetCameras")
	purego.RegisterLibFunc(&sdlGetCameraSupportedFormats, lib, "SDL_GetCameraSupportedFormats")
	// purego.RegisterLibFunc(&sdlGetClipboardData, lib, "SDL_GetClipboardData")
	// purego.RegisterLibFunc(&sdlGetClipboardMimeTypes, lib, "SDL_GetClipboardMimeTypes")
	purego.RegisterLibFunc(&sdlGetClipboardText, lib, "SDL_GetClipboardText")
	purego.RegisterLibFunc(&sdlGetClosestFullscreenDisplayMode, lib, "SDL_GetClosestFullscreenDisplayMode")
	// purego.RegisterLibFunc(&sdlGetCPUCacheLineSize, lib, "SDL_GetCPUCacheLineSize")
	purego.RegisterLibFunc(&sdlGetCurrentAudioDriver, lib, "SDL_GetCurrentAudioDriver")
	purego.RegisterLibFunc(&sdlGetCurrentCameraDriver, lib, "SDL_GetCurrentCameraDriver")
	// purego.RegisterLibFunc(&sdlGetCurrentDirectory, lib, "SDL_GetCurrentDirectory")
	purego.RegisterLibFunc(&sdlGetCurrentDisplayMode, lib, "SDL_GetCurrentDisplayMode")
	// purego.RegisterLibFunc(&sdlGetCurrentDisplayOrientation, lib, "SDL_GetCurrentDisplayOrientation")
	purego.RegisterLibFunc(&sdlGetCurrentRenderOutputSize, lib, "SDL_GetCurrentRenderOutputSize")
	// purego.RegisterLibFunc(&sdlGetCurrentThreadID, lib, "SDL_GetCurrentThreadID")
	// purego.RegisterLibFunc(&sdlGetCurrentTime, lib, "SDL_GetCurrentTime")
	purego.RegisterLibFunc(&sdlGetCurrentVideoDriver, lib, "SDL_GetCurrentVideoDriver")
	purego.RegisterLibFunc(&sdlGetCursor, lib, "SDL_GetCursor")
	// purego.RegisterLibFunc(&sdlGetDateTimeLocalePreferences, lib, "SDL_GetDateTimeLocalePreferences")
	// purego.RegisterLibFunc(&sdlGetDayOfWeek, lib, "SDL_GetDayOfWeek")
	// purego.RegisterLibFunc(&sdlGetDayOfYear, lib, "SDL_GetDayOfYear")
	// purego.RegisterLibFunc(&sdlGetDaysInMonth, lib, "SDL_GetDaysInMonth")
	// purego.RegisterLibFunc(&sdlGetDefaultAssertionHandler, lib, "SDL_GetDefaultAssertionHandler")
	purego.RegisterLibFunc(&sdlGetDefaultCursor, lib, "SDL_GetDefaultCursor")
	// purego.RegisterLibFunc(&sdlGetDefaultLogOutputFunction, lib, "SDL_GetDefaultLogOutputFunction")
	// purego.RegisterLibFunc(&sdlGetDesktopDisplayMode, lib, "SDL_GetDesktopDisplayMode")
	// purego.RegisterLibFunc(&sdlGetDisplayBounds, lib, "SDL_GetDisplayBounds")
	purego.RegisterLibFunc(&sdlGetDisplayContentScale, lib, "SDL_GetDisplayContentScale")
	// purego.RegisterLibFunc(&sdlGetDisplayForPoint, lib, "SDL_GetDisplayForPoint")
	// purego.RegisterLibFunc(&sdlGetDisplayForRect, lib, "SDL_GetDisplayForRect")
	purego.RegisterLibFunc(&sdlGetDisplayForWindow, lib, "SDL_GetDisplayForWindow")
	purego.RegisterLibFunc(&sdlGetDisplayName, lib, "SDL_GetDisplayName")
	// purego.RegisterLibFunc(&sdlGetDisplayProperties, lib, "SDL_GetDisplayProperties")
	purego.RegisterLibFunc(&sdlGetDisplays, lib, "SDL_GetDisplays")
	// purego.RegisterLibFunc(&sdlGetDisplayUsableBounds, lib, "SDL_GetDisplayUsableBounds")
	// purego.RegisterLibFunc(&sdlgetenv, lib, "SDL_getenv")
	// purego.RegisterLibFunc(&sdlgetenv_unsafe, lib, "SDL_getenv_unsafe")
	// purego.RegisterLibFunc(&sdlGetEnvironment, lib, "SDL_GetEnvironment")
	// purego.RegisterLibFunc(&sdlGetEnvironmentVariable, lib, "SDL_GetEnvironmentVariable")
	// purego.RegisterLibFunc(&sdlGetEnvironmentVariables, lib, "SDL_GetEnvironmentVariables")
	purego.RegisterLibFunc(&sdlGetError, lib, "SDL_GetError")
	purego.RegisterLibFunc(&sdlGetEventFilter, lib, "SDL_GetEventFilter")
	purego.RegisterLibFunc(&sdlGetFloatProperty, lib, "SDL_GetFloatProperty")
	purego.RegisterLibFunc(&sdlGetFullscreenDisplayModes, lib, "SDL_GetFullscreenDisplayModes")
	// purego.RegisterLibFunc(&sdlGetGamepadAppleSFSymbolsNameForAxis, lib, "SDL_GetGamepadAppleSFSymbolsNameForAxis")
	// purego.RegisterLibFunc(&sdlGetGamepadAppleSFSymbolsNameForButton, lib, "SDL_GetGamepadAppleSFSymbolsNameForButton")
	// purego.RegisterLibFunc(&sdlGetGamepadAxis, lib, "SDL_GetGamepadAxis")
	// purego.RegisterLibFunc(&sdlGetGamepadAxisFromString, lib, "SDL_GetGamepadAxisFromString")
	purego.RegisterLibFunc(&sdlGetGamepadBindings, lib, "SDL_GetGamepadBindings")
	// purego.RegisterLibFunc(&sdlGetGamepadButton, lib, "SDL_GetGamepadButton")
	// purego.RegisterLibFunc(&sdlGetGamepadButtonFromString, lib, "SDL_GetGamepadButtonFromString")
	// purego.RegisterLibFunc(&sdlGetGamepadButtonLabel, lib, "SDL_GetGamepadButtonLabel")
	// purego.RegisterLibFunc(&sdlGetGamepadButtonLabelForType, lib, "SDL_GetGamepadButtonLabelForType")
	// purego.RegisterLibFunc(&sdlGetGamepadConnectionState, lib, "SDL_GetGamepadConnectionState")
	// purego.RegisterLibFunc(&sdlGetGamepadFirmwareVersion, lib, "SDL_GetGamepadFirmwareVersion")
	purego.RegisterLibFunc(&sdlGetGamepadFromID, lib, "SDL_GetGamepadFromID")
	// purego.RegisterLibFunc(&sdlGetGamepadFromPlayerIndex, lib, "SDL_GetGamepadFromPlayerIndex")
	// purego.RegisterLibFunc(&sdlGetGamepadGUIDForID, lib, "SDL_GetGamepadGUIDForID")
	// purego.RegisterLibFunc(&sdlGetGamepadID, lib, "SDL_GetGamepadID")
	// purego.RegisterLibFunc(&sdlGetGamepadJoystick, lib, "SDL_GetGamepadJoystick")
	// purego.RegisterLibFunc(&sdlGetGamepadMapping, lib, "SDL_GetGamepadMapping")
	// purego.RegisterLibFunc(&sdlGetGamepadMappingForGUID, lib, "SDL_GetGamepadMappingForGUID")
	// purego.RegisterLibFunc(&sdlGetGamepadMappingForID, lib, "SDL_GetGamepadMappingForID")
	// purego.RegisterLibFunc(&sdlGetGamepadMappings, lib, "SDL_GetGamepadMappings")
	purego.RegisterLibFunc(&sdlGetGamepadName, lib, "SDL_GetGamepadName")
	purego.RegisterLibFunc(&sdlGetGamepadNameForID, lib, "SDL_GetGamepadNameForID")
	// purego.RegisterLibFunc(&sdlGetGamepadPath, lib, "SDL_GetGamepadPath")
	// purego.RegisterLibFunc(&sdlGetGamepadPathForID, lib, "SDL_GetGamepadPathForID")
	// purego.RegisterLibFunc(&sdlGetGamepadPlayerIndex, lib, "SDL_GetGamepadPlayerIndex")
	// purego.RegisterLibFunc(&sdlGetGamepadPlayerIndexForID, lib, "SDL_GetGamepadPlayerIndexForID")
	// purego.RegisterLibFunc(&sdlGetGamepadPowerInfo, lib, "SDL_GetGamepadPowerInfo")
	// purego.RegisterLibFunc(&sdlGetGamepadProduct, lib, "SDL_GetGamepadProduct")
	// purego.RegisterLibFunc(&sdlGetGamepadProductForID, lib, "SDL_GetGamepadProductForID")
	// purego.RegisterLibFunc(&sdlGetGamepadProductVersion, lib, "SDL_GetGamepadProductVersion")
	// purego.RegisterLibFunc(&sdlGetGamepadProductVersionForID, lib, "SDL_GetGamepadProductVersionForID")
	// purego.RegisterLibFunc(&sdlGetGamepadProperties, lib, "SDL_GetGamepadProperties")
	purego.RegisterLibFunc(&sdlGetGamepads, lib, "SDL_GetGamepads")
	// purego.RegisterLibFunc(&sdlGetGamepadSensorData, lib, "SDL_GetGamepadSensorData")
	// purego.RegisterLibFunc(&sdlGetGamepadSensorDataRate, lib, "SDL_GetGamepadSensorDataRate")
	purego.RegisterLibFunc(&sdlGetGamepadSerial, lib, "SDL_GetGamepadSerial")
	// purego.RegisterLibFunc(&sdlGetGamepadSteamHandle, lib, "SDL_GetGamepadSteamHandle")
	// purego.RegisterLibFunc(&sdlGetGamepadStringForAxis, lib, "SDL_GetGamepadStringForAxis")
	purego.RegisterLibFunc(&sdlGetGamepadStringForButton, lib, "SDL_GetGamepadStringForButton")
	purego.RegisterLibFunc(&sdlGetGamepadStringForType, lib, "SDL_GetGamepadStringForType")
	// purego.RegisterLibFunc(&sdlGetGamepadTouchpadFinger, lib, "SDL_GetGamepadTouchpadFinger")
	purego.RegisterLibFunc(&sdlGetGamepadType, lib, "SDL_GetGamepadType")
	// purego.RegisterLibFunc(&sdlGetGamepadTypeForID, lib, "SDL_GetGamepadTypeForID")
	// purego.RegisterLibFunc(&sdlGetGamepadTypeFromString, lib, "SDL_GetGamepadTypeFromString")
	// purego.RegisterLibFunc(&sdlGetGamepadVendor, lib, "SDL_GetGamepadVendor")
	// purego.RegisterLibFunc(&sdlGetGamepadVendorForID, lib, "SDL_GetGamepadVendorForID")
	purego.RegisterLibFunc(&sdlGetGlobalMouseState, lib, "SDL_GetGlobalMouseState")
	purego.RegisterLibFunc(&sdlGetGlobalProperties, lib, "SDL_GetGlobalProperties")
	purego.RegisterLibFunc(&sdlGetGPUDeviceDriver, lib, "SDL_GetGPUDeviceDriver")
	purego.RegisterLibFunc(&sdlGetGPUDriver, lib, "SDL_GetGPUDriver")
	purego.RegisterLibFunc(&sdlGetGPUShaderFormats, lib, "SDL_GetGPUShaderFormats")
	purego.RegisterLibFunc(&sdlGetGPUSwapchainTextureFormat, lib, "SDL_GetGPUSwapchainTextureFormat")
	// purego.RegisterLibFunc(&sdlGetGrabbedWindow, lib, "SDL_GetGrabbedWindow")
	// purego.RegisterLibFunc(&sdlGetHapticEffectStatus, lib, "SDL_GetHapticEffectStatus")
	// purego.RegisterLibFunc(&sdlGetHapticFeatures, lib, "SDL_GetHapticFeatures")
	// purego.RegisterLibFunc(&sdlGetHapticFromID, lib, "SDL_GetHapticFromID")
	// purego.RegisterLibFunc(&sdlGetHapticID, lib, "SDL_GetHapticID")
	// purego.RegisterLibFunc(&sdlGetHapticName, lib, "SDL_GetHapticName")
	// purego.RegisterLibFunc(&sdlGetHapticNameForID, lib, "SDL_GetHapticNameForID")
	// purego.RegisterLibFunc(&sdlGetHaptics, lib, "SDL_GetHaptics")
	purego.RegisterLibFunc(&sdlGetHint, lib, "SDL_GetHint")
	purego.RegisterLibFunc(&sdlGetHintBoolean, lib, "SDL_GetHintBoolean")
	// purego.RegisterLibFunc(&sdlGetIOProperties, lib, "SDL_GetIOProperties")
	// purego.RegisterLibFunc(&sdlGetIOSize, lib, "SDL_GetIOSize")
	// purego.RegisterLibFunc(&sdlGetIOStatus, lib, "SDL_GetIOStatus")
	purego.RegisterLibFunc(&sdlGetJoystickAxis, lib, "SDL_GetJoystickAxis")
	purego.RegisterLibFunc(&sdlGetJoystickAxisInitialState, lib, "SDL_GetJoystickAxisInitialState")
	purego.RegisterLibFunc(&sdlGetJoystickBall, lib, "SDL_GetJoystickBall")
	purego.RegisterLibFunc(&sdlGetJoystickButton, lib, "SDL_GetJoystickButton")
	purego.RegisterLibFunc(&sdlGetJoystickConnectionState, lib, "SDL_GetJoystickConnectionState")
	purego.RegisterLibFunc(&sdlGetJoystickFirmwareVersion, lib, "SDL_GetJoystickFirmwareVersion")
	purego.RegisterLibFunc(&sdlGetJoystickFromID, lib, "SDL_GetJoystickFromID")
	purego.RegisterLibFunc(&sdlGetJoystickFromPlayerIndex, lib, "SDL_GetJoystickFromPlayerIndex")
	// purego.RegisterLibFunc(&sdlGetJoystickGUID, lib, "SDL_GetJoystickGUID")
	// purego.RegisterLibFunc(&sdlGetJoystickGUIDForID, lib, "SDL_GetJoystickGUIDForID")
	// purego.RegisterLibFunc(&sdlGetJoystickGUIDInfo, lib, "SDL_GetJoystickGUIDInfo")
	purego.RegisterLibFunc(&sdlGetJoystickHat, lib, "SDL_GetJoystickHat")
	purego.RegisterLibFunc(&sdlGetJoystickID, lib, "SDL_GetJoystickID")
	purego.RegisterLibFunc(&sdlGetJoystickName, lib, "SDL_GetJoystickName")
	purego.RegisterLibFunc(&sdlGetJoystickNameForID, lib, "SDL_GetJoystickNameForID")
	purego.RegisterLibFunc(&sdlGetJoystickPath, lib, "SDL_GetJoystickPath")
	purego.RegisterLibFunc(&sdlGetJoystickPathForID, lib, "SDL_GetJoystickPathForID")
	purego.RegisterLibFunc(&sdlGetJoystickPlayerIndex, lib, "SDL_GetJoystickPlayerIndex")
	purego.RegisterLibFunc(&sdlGetJoystickPlayerIndexForID, lib, "SDL_GetJoystickPlayerIndexForID")
	purego.RegisterLibFunc(&sdlGetJoystickPowerInfo, lib, "SDL_GetJoystickPowerInfo")
	purego.RegisterLibFunc(&sdlGetJoystickProduct, lib, "SDL_GetJoystickProduct")
	purego.RegisterLibFunc(&sdlGetJoystickProductForID, lib, "SDL_GetJoystickProductForID")
	purego.RegisterLibFunc(&sdlGetJoystickProductVersion, lib, "SDL_GetJoystickProductVersion")
	purego.RegisterLibFunc(&sdlGetJoystickProductVersionForID, lib, "SDL_GetJoystickProductVersionForID")
	purego.RegisterLibFunc(&sdlGetJoystickProperties, lib, "SDL_GetJoystickProperties")
	purego.RegisterLibFunc(&sdlGetJoysticks, lib, "SDL_GetJoysticks")
	purego.RegisterLibFunc(&sdlGetJoystickSerial, lib, "SDL_GetJoystickSerial")
	purego.RegisterLibFunc(&sdlGetJoystickType, lib, "SDL_GetJoystickType")
	purego.RegisterLibFunc(&sdlGetJoystickTypeForID, lib, "SDL_GetJoystickTypeForID")
	purego.RegisterLibFunc(&sdlGetJoystickVendor, lib, "SDL_GetJoystickVendor")
	purego.RegisterLibFunc(&sdlGetJoystickVendorForID, lib, "SDL_GetJoystickVendorForID")
	purego.RegisterLibFunc(&sdlGetKeyboardFocus, lib, "SDL_GetKeyboardFocus")
	purego.RegisterLibFunc(&sdlGetKeyboardNameForID, lib, "SDL_GetKeyboardNameForID")
	purego.RegisterLibFunc(&sdlGetKeyboards, lib, "SDL_GetKeyboards")
	purego.RegisterLibFunc(&sdlGetKeyboardState, lib, "SDL_GetKeyboardState")
	purego.RegisterLibFunc(&sdlGetKeyFromName, lib, "SDL_GetKeyFromName")
	purego.RegisterLibFunc(&sdlGetKeyFromScancode, lib, "SDL_GetKeyFromScancode")
	purego.RegisterLibFunc(&sdlGetKeyName, lib, "SDL_GetKeyName")
	// purego.RegisterLibFunc(&sdlGetLogOutputFunction, lib, "SDL_GetLogOutputFunction")
	// purego.RegisterLibFunc(&sdlGetLogPriority, lib, "SDL_GetLogPriority")
	// purego.RegisterLibFunc(&sdlGetMasksForPixelFormat, lib, "SDL_GetMasksForPixelFormat")
	// purego.RegisterLibFunc(&sdlGetMaxHapticEffects, lib, "SDL_GetMaxHapticEffects")
	// purego.RegisterLibFunc(&sdlGetMaxHapticEffectsPlaying, lib, "SDL_GetMaxHapticEffectsPlaying")
	// purego.RegisterLibFunc(&sdlGetMemoryFunctions, lib, "SDL_GetMemoryFunctions")
	purego.RegisterLibFunc(&sdlGetMice, lib, "SDL_GetMice")
	purego.RegisterLibFunc(&sdlGetModState, lib, "SDL_GetModState")
	purego.RegisterLibFunc(&sdlGetMouseFocus, lib, "SDL_GetMouseFocus")
	purego.RegisterLibFunc(&sdlGetMouseNameForID, lib, "SDL_GetMouseNameForID")
	purego.RegisterLibFunc(&sdlGetMouseState, lib, "SDL_GetMouseState")
	// purego.RegisterLibFunc(&sdlGetNaturalDisplayOrientation, lib, "SDL_GetNaturalDisplayOrientation")
	// purego.RegisterLibFunc(&sdlGetNumAllocations, lib, "SDL_GetNumAllocations")
	purego.RegisterLibFunc(&sdlGetNumAudioDrivers, lib, "SDL_GetNumAudioDrivers")
	purego.RegisterLibFunc(&sdlGetNumberProperty, lib, "SDL_GetNumberProperty")
	purego.RegisterLibFunc(&sdlGetNumCameraDrivers, lib, "SDL_GetNumCameraDrivers")
	// purego.RegisterLibFunc(&sdlGetNumGamepadTouchpadFingers, lib, "SDL_GetNumGamepadTouchpadFingers")
	// purego.RegisterLibFunc(&sdlGetNumGamepadTouchpads, lib, "SDL_GetNumGamepadTouchpads")
	purego.RegisterLibFunc(&sdlGetNumGPUDrivers, lib, "SDL_GetNumGPUDrivers")
	// purego.RegisterLibFunc(&sdlGetNumHapticAxes, lib, "SDL_GetNumHapticAxes")
	purego.RegisterLibFunc(&sdlGetNumJoystickAxes, lib, "SDL_GetNumJoystickAxes")
	purego.RegisterLibFunc(&sdlGetNumJoystickBalls, lib, "SDL_GetNumJoystickBalls")
	purego.RegisterLibFunc(&sdlGetNumJoystickButtons, lib, "SDL_GetNumJoystickButtons")
	purego.RegisterLibFunc(&sdlGetNumJoystickHats, lib, "SDL_GetNumJoystickHats")
	// purego.RegisterLibFunc(&sdlGetNumLogicalCPUCores, lib, "SDL_GetNumLogicalCPUCores")
	purego.RegisterLibFunc(&sdlGetNumRenderDrivers, lib, "SDL_GetNumRenderDrivers")
	purego.RegisterLibFunc(&sdlGetNumVideoDrivers, lib, "SDL_GetNumVideoDrivers")
	// purego.RegisterLibFunc(&sdlGetOriginalMemoryFunctions, lib, "SDL_GetOriginalMemoryFunctions")
	// purego.RegisterLibFunc(&sdlGetPathInfo, lib, "SDL_GetPathInfo")
	sdlGetPerformanceCounter = shared.Get(lib, "SDL_GetPerformanceCounter")
	sdlGetPerformanceFrequency = shared.Get(lib, "SDL_GetPerformanceFrequency")
	purego.RegisterLibFunc(&sdlGetPixelFormatDetails, lib, "SDL_GetPixelFormatDetails")
	// purego.RegisterLibFunc(&sdlGetPixelFormatForMasks, lib, "SDL_GetPixelFormatForMasks")
	// purego.RegisterLibFunc(&sdlGetPixelFormatName, lib, "SDL_GetPixelFormatName")
	// purego.RegisterLibFunc(&sdlGetPlatform, lib, "SDL_GetPlatform")
	purego.RegisterLibFunc(&sdlGetPointerProperty, lib, "SDL_GetPointerProperty")
	purego.RegisterLibFunc(&sdlGetPowerInfo, lib, "SDL_GetPowerInfo")
	purego.RegisterLibFunc(&sdlGetPreferredLocales, lib, "SDL_GetPreferredLocales")
	// purego.RegisterLibFunc(&sdlGetPrefPath, lib, "SDL_GetPrefPath")
	purego.RegisterLibFunc(&sdlGetPrimaryDisplay, lib, "SDL_GetPrimaryDisplay")
	// purego.RegisterLibFunc(&sdlGetPrimarySelectionText, lib, "SDL_GetPrimarySelectionText")
	// purego.RegisterLibFunc(&sdlGetProcessInput, lib, "SDL_GetProcessInput")
	// purego.RegisterLibFunc(&sdlGetProcessOutput, lib, "SDL_GetProcessOutput")
	// purego.RegisterLibFunc(&sdlGetProcessProperties, lib, "SDL_GetProcessProperties")
	purego.RegisterLibFunc(&sdlGetPropertyType, lib, "SDL_GetPropertyType")
	// purego.RegisterLibFunc(&sdlGetRealGamepadType, lib, "SDL_GetRealGamepadType")
	// purego.RegisterLibFunc(&sdlGetRealGamepadTypeForID, lib, "SDL_GetRealGamepadTypeForID")
	purego.RegisterLibFunc(&sdlGetRectAndLineIntersection, lib, "SDL_GetRectAndLineIntersection")
	purego.RegisterLibFunc(&sdlGetRectAndLineIntersectionFloat, lib, "SDL_GetRectAndLineIntersectionFloat")
	purego.RegisterLibFunc(&sdlGetRectEnclosingPoints, lib, "SDL_GetRectEnclosingPoints")
	purego.RegisterLibFunc(&sdlGetRectEnclosingPointsFloat, lib, "SDL_GetRectEnclosingPointsFloat")
	purego.RegisterLibFunc(&sdlGetRectIntersection, lib, "SDL_GetRectIntersection")
	purego.RegisterLibFunc(&sdlGetRectIntersectionFloat, lib, "SDL_GetRectIntersectionFloat")
	purego.RegisterLibFunc(&sdlGetRectUnion, lib, "SDL_GetRectUnion")
	purego.RegisterLibFunc(&sdlGetRectUnionFloat, lib, "SDL_GetRectUnionFloat")
	purego.RegisterLibFunc(&sdlGetRelativeMouseState, lib, "SDL_GetRelativeMouseState")
	purego.RegisterLibFunc(&sdlGetRenderClipRect, lib, "SDL_GetRenderClipRect")
	purego.RegisterLibFunc(&sdlGetRenderColorScale, lib, "SDL_GetRenderColorScale")
	purego.RegisterLibFunc(&sdlGetRenderDrawBlendMode, lib, "SDL_GetRenderDrawBlendMode")
	purego.RegisterLibFunc(&sdlGetRenderDrawColor, lib, "SDL_GetRenderDrawColor")
	purego.RegisterLibFunc(&sdlGetRenderDrawColorFloat, lib, "SDL_GetRenderDrawColorFloat")
	purego.RegisterLibFunc(&sdlGetRenderDriver, lib, "SDL_GetRenderDriver")
	purego.RegisterLibFunc(&sdlGetRenderer, lib, "SDL_GetRenderer")
	purego.RegisterLibFunc(&sdlGetRendererFromTexture, lib, "SDL_GetRendererFromTexture")
	purego.RegisterLibFunc(&sdlGetRendererName, lib, "SDL_GetRendererName")
	purego.RegisterLibFunc(&sdlGetRendererProperties, lib, "SDL_GetRendererProperties")
	purego.RegisterLibFunc(&sdlGetRenderLogicalPresentation, lib, "SDL_GetRenderLogicalPresentation")
	purego.RegisterLibFunc(&sdlGetRenderLogicalPresentationRect, lib, "SDL_GetRenderLogicalPresentationRect")
	purego.RegisterLibFunc(&sdlGetRenderMetalCommandEncoder, lib, "SDL_GetRenderMetalCommandEncoder")
	purego.RegisterLibFunc(&sdlGetRenderMetalLayer, lib, "SDL_GetRenderMetalLayer")
	purego.RegisterLibFunc(&sdlGetRenderOutputSize, lib, "SDL_GetRenderOutputSize")
	purego.RegisterLibFunc(&sdlGetRenderSafeArea, lib, "SDL_GetRenderSafeArea")
	purego.RegisterLibFunc(&sdlGetRenderScale, lib, "SDL_GetRenderScale")
	purego.RegisterLibFunc(&sdlGetRenderTarget, lib, "SDL_GetRenderTarget")
	purego.RegisterLibFunc(&sdlGetRenderViewport, lib, "SDL_GetRenderViewport")
	purego.RegisterLibFunc(&sdlGetRenderVSync, lib, "SDL_GetRenderVSync")
	purego.RegisterLibFunc(&sdlGetRenderWindow, lib, "SDL_GetRenderWindow")
	purego.RegisterLibFunc(&sdlGetRevision, lib, "SDL_GetRevision")
	// purego.RegisterLibFunc(&sdlGetRGB, lib, "SDL_GetRGB")
	// purego.RegisterLibFunc(&sdlGetRGBA, lib, "SDL_GetRGBA")
	// purego.RegisterLibFunc(&sdlGetSandbox, lib, "SDL_GetSandbox")
	purego.RegisterLibFunc(&sdlGetScancodeFromKey, lib, "SDL_GetScancodeFromKey")
	purego.RegisterLibFunc(&sdlGetScancodeFromName, lib, "SDL_GetScancodeFromName")
	purego.RegisterLibFunc(&sdlGetScancodeName, lib, "SDL_GetScancodeName")
	// purego.RegisterLibFunc(&sdlGetSemaphoreValue, lib, "SDL_GetSemaphoreValue")
	// purego.RegisterLibFunc(&sdlGetSensorData, lib, "SDL_GetSensorData")
	// purego.RegisterLibFunc(&sdlGetSensorFromID, lib, "SDL_GetSensorFromID")
	// purego.RegisterLibFunc(&sdlGetSensorID, lib, "SDL_GetSensorID")
	// purego.RegisterLibFunc(&sdlGetSensorName, lib, "SDL_GetSensorName")
	// purego.RegisterLibFunc(&sdlGetSensorNameForID, lib, "SDL_GetSensorNameForID")
	// purego.RegisterLibFunc(&sdlGetSensorNonPortableType, lib, "SDL_GetSensorNonPortableType")
	// purego.RegisterLibFunc(&sdlGetSensorNonPortableTypeForID, lib, "SDL_GetSensorNonPortableTypeForID")
	// purego.RegisterLibFunc(&sdlGetSensorProperties, lib, "SDL_GetSensorProperties")
	// purego.RegisterLibFunc(&sdlGetSensors, lib, "SDL_GetSensors")
	// purego.RegisterLibFunc(&sdlGetSensorType, lib, "SDL_GetSensorType")
	// purego.RegisterLibFunc(&sdlGetSensorTypeForID, lib, "SDL_GetSensorTypeForID")
	// purego.RegisterLibFunc(&sdlGetSilenceValueForFormat, lib, "SDL_GetSilenceValueForFormat")
	// purego.RegisterLibFunc(&sdlGetSIMDAlignment, lib, "SDL_GetSIMDAlignment")
	// purego.RegisterLibFunc(&sdlGetStorageFileSize, lib, "SDL_GetStorageFileSize")
	// purego.RegisterLibFunc(&sdlGetStoragePathInfo, lib, "SDL_GetStoragePathInfo")
	// purego.RegisterLibFunc(&sdlGetStorageSpaceRemaining, lib, "SDL_GetStorageSpaceRemaining")
	purego.RegisterLibFunc(&sdlGetStringProperty, lib, "SDL_GetStringProperty")
	purego.RegisterLibFunc(&sdlGetSurfaceAlphaMod, lib, "SDL_GetSurfaceAlphaMod")
	purego.RegisterLibFunc(&sdlGetSurfaceBlendMode, lib, "SDL_GetSurfaceBlendMode")
	purego.RegisterLibFunc(&sdlGetSurfaceClipRect, lib, "SDL_GetSurfaceClipRect")
	purego.RegisterLibFunc(&sdlGetSurfaceColorKey, lib, "SDL_GetSurfaceColorKey")
	purego.RegisterLibFunc(&sdlGetSurfaceColorMod, lib, "SDL_GetSurfaceColorMod")
	purego.RegisterLibFunc(&sdlGetSurfaceColorspace, lib, "SDL_GetSurfaceColorspace")
	purego.RegisterLibFunc(&sdlGetSurfaceImages, lib, "SDL_GetSurfaceImages")
	purego.RegisterLibFunc(&sdlGetSurfacePalette, lib, "SDL_GetSurfacePalette")
	purego.RegisterLibFunc(&sdlGetSurfaceProperties, lib, "SDL_GetSurfaceProperties")
	// purego.RegisterLibFunc(&sdlGetSystemRAM, lib, "SDL_GetSystemRAM")
	// purego.RegisterLibFunc(&sdlGetSystemTheme, lib, "SDL_GetSystemTheme")
	purego.RegisterLibFunc(&sdlGetTextInputArea, lib, "SDL_GetTextInputArea")
	purego.RegisterLibFunc(&sdlGetTextureAlphaMod, lib, "SDL_GetTextureAlphaMod")
	purego.RegisterLibFunc(&sdlGetTextureAlphaModFloat, lib, "SDL_GetTextureAlphaModFloat")
	purego.RegisterLibFunc(&sdlGetTextureBlendMode, lib, "SDL_GetTextureBlendMode")
	purego.RegisterLibFunc(&sdlGetTextureColorMod, lib, "SDL_GetTextureColorMod")
	purego.RegisterLibFunc(&sdlGetTextureColorModFloat, lib, "SDL_GetTextureColorModFloat")
	purego.RegisterLibFunc(&sdlGetTextureProperties, lib, "SDL_GetTextureProperties")
	purego.RegisterLibFunc(&sdlGetTextureScaleMode, lib, "SDL_GetTextureScaleMode")
	purego.RegisterLibFunc(&sdlGetTextureSize, lib, "SDL_GetTextureSize")
	// purego.RegisterLibFunc(&sdlGetThreadID, lib, "SDL_GetThreadID")
	// purego.RegisterLibFunc(&sdlGetThreadName, lib, "SDL_GetThreadName")
	// purego.RegisterLibFunc(&sdlGetThreadState, lib, "SDL_GetThreadState")
	sdlGetTicks = shared.Get(lib, "SDL_GetTicks")
	sdlGetTicksNS = shared.Get(lib, "SDL_GetTicksNS")
	// purego.RegisterLibFunc(&sdlGetTLS, lib, "SDL_GetTLS")
	// purego.RegisterLibFunc(&sdlGetTouchDeviceName, lib, "SDL_GetTouchDeviceName")
	// purego.RegisterLibFunc(&sdlGetTouchDevices, lib, "SDL_GetTouchDevices")
	// purego.RegisterLibFunc(&sdlGetTouchDeviceType, lib, "SDL_GetTouchDeviceType")
	// purego.RegisterLibFunc(&sdlGetTouchFingers, lib, "SDL_GetTouchFingers")
	// purego.RegisterLibFunc(&sdlGetTrayEntries, lib, "SDL_GetTrayEntries")
	// purego.RegisterLibFunc(&sdlGetTrayEntryChecked, lib, "SDL_GetTrayEntryChecked")
	// purego.RegisterLibFunc(&sdlGetTrayEntryEnabled, lib, "SDL_GetTrayEntryEnabled")
	// purego.RegisterLibFunc(&sdlGetTrayEntryLabel, lib, "SDL_GetTrayEntryLabel")
	// purego.RegisterLibFunc(&sdlGetTrayEntryParent, lib, "SDL_GetTrayEntryParent")
	// purego.RegisterLibFunc(&sdlGetTrayMenu, lib, "SDL_GetTrayMenu")
	// purego.RegisterLibFunc(&sdlGetTrayMenuParentEntry, lib, "SDL_GetTrayMenuParentEntry")
	// purego.RegisterLibFunc(&sdlGetTrayMenuParentTray, lib, "SDL_GetTrayMenuParentTray")
	// purego.RegisterLibFunc(&sdlGetTraySubmenu, lib, "SDL_GetTraySubmenu")
	// purego.RegisterLibFunc(&sdlGetUserFolder, lib, "SDL_GetUserFolder")
	purego.RegisterLibFunc(&sdlGetVersion, lib, "SDL_GetVersion")
	purego.RegisterLibFunc(&sdlGetVideoDriver, lib, "SDL_GetVideoDriver")
	// purego.RegisterLibFunc(&sdlGetWindowAspectRatio, lib, "SDL_GetWindowAspectRatio")
	// purego.RegisterLibFunc(&sdlGetWindowBordersSize, lib, "SDL_GetWindowBordersSize")
	purego.RegisterLibFunc(&sdlGetWindowDisplayScale, lib, "SDL_GetWindowDisplayScale")
	purego.RegisterLibFunc(&sdlGetWindowFlags, lib, "SDL_GetWindowFlags")
	purego.RegisterLibFunc(&sdlGetWindowFromEvent, lib, "SDL_GetWindowFromEvent")
	// purego.RegisterLibFunc(&sdlGetWindowFromID, lib, "SDL_GetWindowFromID")
	purego.RegisterLibFunc(&sdlGetWindowFullscreenMode, lib, "SDL_GetWindowFullscreenMode")
	// purego.RegisterLibFunc(&sdlGetWindowICCProfile, lib, "SDL_GetWindowICCProfile")
	purego.RegisterLibFunc(&sdlGetWindowID, lib, "SDL_GetWindowID")
	purego.RegisterLibFunc(&sdlGetWindowKeyboardGrab, lib, "SDL_GetWindowKeyboardGrab")
	// purego.RegisterLibFunc(&sdlGetWindowMaximumSize, lib, "SDL_GetWindowMaximumSize")
	// purego.RegisterLibFunc(&sdlGetWindowMinimumSize, lib, "SDL_GetWindowMinimumSize")
	purego.RegisterLibFunc(&sdlGetWindowMouseGrab, lib, "SDL_GetWindowMouseGrab")
	// purego.RegisterLibFunc(&sdlGetWindowMouseRect, lib, "SDL_GetWindowMouseRect")
	purego.RegisterLibFunc(&sdlGetWindowOpacity, lib, "SDL_GetWindowOpacity")
	// purego.RegisterLibFunc(&sdlGetWindowParent, lib, "SDL_GetWindowParent")
	purego.RegisterLibFunc(&sdlGetWindowPixelDensity, lib, "SDL_GetWindowPixelDensity")
	// purego.RegisterLibFunc(&sdlGetWindowPixelFormat, lib, "SDL_GetWindowPixelFormat")
	purego.RegisterLibFunc(&sdlGetWindowPosition, lib, "SDL_GetWindowPosition")
	// purego.RegisterLibFunc(&sdlGetWindowProperties, lib, "SDL_GetWindowProperties")
	purego.RegisterLibFunc(&sdlGetWindowRelativeMouseMode, lib, "SDL_GetWindowRelativeMouseMode")
	// purego.RegisterLibFunc(&sdlGetWindows, lib, "SDL_GetWindows")
	// purego.RegisterLibFunc(&sdlGetWindowSafeArea, lib, "SDL_GetWindowSafeArea")
	purego.RegisterLibFunc(&sdlGetWindowSize, lib, "SDL_GetWindowSize")
	purego.RegisterLibFunc(&sdlGetWindowSizeInPixels, lib, "SDL_GetWindowSizeInPixels")
	purego.RegisterLibFunc(&sdlGetWindowSurface, lib, "SDL_GetWindowSurface")
	// purego.RegisterLibFunc(&sdlGetWindowSurfaceVSync, lib, "SDL_GetWindowSurfaceVSync")
	purego.RegisterLibFunc(&sdlGetWindowTitle, lib, "SDL_GetWindowTitle")
	purego.RegisterLibFunc(&sdlGLCreateContext, lib, "SDL_GL_CreateContext")
	purego.RegisterLibFunc(&sdlGLDestroyContext, lib, "SDL_GL_DestroyContext")
	// purego.RegisterLibFunc(&sdlGL_ExtensionSupported, lib, "SDL_GL_ExtensionSupported")
	purego.RegisterLibFunc(&sdlGLGetAttribute, lib, "SDL_GL_GetAttribute")
	purego.RegisterLibFunc(&sdlGLGetCurrentContext, lib, "SDL_GL_GetCurrentContext")
	purego.RegisterLibFunc(&sdlGLGetCurrentWindow, lib, "SDL_GL_GetCurrentWindow")
	// purego.RegisterLibFunc(&sdlGL_GetProcAddress, lib, "SDL_GL_GetProcAddress")
	// purego.RegisterLibFunc(&sdlGL_GetSwapInterval, lib, "SDL_GL_GetSwapInterval")
	// purego.RegisterLibFunc(&sdlGL_LoadLibrary, lib, "SDL_GL_LoadLibrary")
	// purego.RegisterLibFunc(&sdlGL_MakeCurrent, lib, "SDL_GL_MakeCurrent")
	// purego.RegisterLibFunc(&sdlGL_ResetAttributes, lib, "SDL_GL_ResetAttributes")
	purego.RegisterLibFunc(&sdlGLSetAttribute, lib, "SDL_GL_SetAttribute")
	sdlGLSetSwapInterval = shared.Get(lib, "SDL_GL_SetSwapInterval")
	sdlGLSwapWindow = shared.Get(lib, "SDL_GL_SwapWindow")
	// purego.RegisterLibFunc(&sdlGL_UnloadLibrary, lib, "SDL_GL_UnloadLibrary")
	// purego.RegisterLibFunc(&sdlGlobDirectory, lib, "SDL_GlobDirectory")
	// purego.RegisterLibFunc(&sdlGlobStorageDirectory, lib, "SDL_GlobStorageDirectory")
	// purego.RegisterLibFunc(&sdlGPUSupportsProperties, lib, "SDL_GPUSupportsProperties")
	// purego.RegisterLibFunc(&sdlGPUSupportsShaderFormats, lib, "SDL_GPUSupportsShaderFormats")
	// purego.RegisterLibFunc(&sdlGPUTextureFormatTexelBlockSize, lib, "SDL_GPUTextureFormatTexelBlockSize")
	// purego.RegisterLibFunc(&sdlGPUTextureSupportsFormat, lib, "SDL_GPUTextureSupportsFormat")
	// purego.RegisterLibFunc(&sdlGPUTextureSupportsSampleCount, lib, "SDL_GPUTextureSupportsSampleCount")
	// purego.RegisterLibFunc(&sdlGUIDToString, lib, "SDL_GUIDToString")
	// purego.RegisterLibFunc(&sdlHapticEffectSupported, lib, "SDL_HapticEffectSupported")
	// purego.RegisterLibFunc(&sdlHapticRumbleSupported, lib, "SDL_HapticRumbleSupported")
	// purego.RegisterLibFunc(&sdlHasAltiVec, lib, "SDL_HasAltiVec")
	// purego.RegisterLibFunc(&sdlHasARMSIMD, lib, "SDL_HasARMSIMD")
	// purego.RegisterLibFunc(&sdlHasAVX, lib, "SDL_HasAVX")
	// purego.RegisterLibFunc(&sdlHasAVX2, lib, "SDL_HasAVX2")
	// purego.RegisterLibFunc(&sdlHasAVX512F, lib, "SDL_HasAVX512F")
	// purego.RegisterLibFunc(&sdlHasClipboardData, lib, "SDL_HasClipboardData")
	// purego.RegisterLibFunc(&sdlHasClipboardText, lib, "SDL_HasClipboardText")
	purego.RegisterLibFunc(&sdlHasEvent, lib, "SDL_HasEvent")
	purego.RegisterLibFunc(&sdlHasEvents, lib, "SDL_HasEvents")
	// purego.RegisterLibFunc(&sdlHasExactlyOneBitSet32, lib, "SDL_HasExactlyOneBitSet32")
	// purego.RegisterLibFunc(&sdlHasGamepad, lib, "SDL_HasGamepad")
	purego.RegisterLibFunc(&sdlHasJoystick, lib, "SDL_HasJoystick")
	purego.RegisterLibFunc(&sdlHasKeyboard, lib, "SDL_HasKeyboard")
	// purego.RegisterLibFunc(&sdlHasLASX, lib, "SDL_HasLASX")
	// purego.RegisterLibFunc(&sdlHasLSX, lib, "SDL_HasLSX")
	// purego.RegisterLibFunc(&sdlHasMMX, lib, "SDL_HasMMX")
	purego.RegisterLibFunc(&sdlHasMouse, lib, "SDL_HasMouse")
	// purego.RegisterLibFunc(&sdlHasNEON, lib, "SDL_HasNEON")
	// purego.RegisterLibFunc(&sdlHasPrimarySelectionText, lib, "SDL_HasPrimarySelectionText")
	purego.RegisterLibFunc(&sdlHasProperty, lib, "SDL_HasProperty")
	purego.RegisterLibFunc(&sdlHasRectIntersection, lib, "SDL_HasRectIntersection")
	purego.RegisterLibFunc(&sdlHasRectIntersectionFloat, lib, "SDL_HasRectIntersectionFloat")
	purego.RegisterLibFunc(&sdlHasScreenKeyboardSupport, lib, "SDL_HasScreenKeyboardSupport")
	// purego.RegisterLibFunc(&sdlHasSSE, lib, "SDL_HasSSE")
	// purego.RegisterLibFunc(&sdlHasSSE2, lib, "SDL_HasSSE2")
	// purego.RegisterLibFunc(&sdlHasSSE3, lib, "SDL_HasSSE3")
	// purego.RegisterLibFunc(&sdlHasSSE41, lib, "SDL_HasSSE41")
	// purego.RegisterLibFunc(&sdlHasSSE42, lib, "SDL_HasSSE42")
	// purego.RegisterLibFunc(&sdlhid_ble_scan, lib, "SDL_hid_ble_scan")
	// purego.RegisterLibFunc(&sdlhid_close, lib, "SDL_hid_close")
	// purego.RegisterLibFunc(&sdlhid_device_change_count, lib, "SDL_hid_device_change_count")
	// purego.RegisterLibFunc(&sdlhid_enumerate, lib, "SDL_hid_enumerate")
	// purego.RegisterLibFunc(&sdlhid_exit, lib, "SDL_hid_exit")
	// purego.RegisterLibFunc(&sdlhid_free_enumeration, lib, "SDL_hid_free_enumeration")
	// purego.RegisterLibFunc(&sdlhid_get_device_info, lib, "SDL_hid_get_device_info")
	// purego.RegisterLibFunc(&sdlhid_get_feature_report, lib, "SDL_hid_get_feature_report")
	// purego.RegisterLibFunc(&sdlhid_get_indexed_string, lib, "SDL_hid_get_indexed_string")
	// purego.RegisterLibFunc(&sdlhid_get_input_report, lib, "SDL_hid_get_input_report")
	// purego.RegisterLibFunc(&sdlhid_get_manufacturer_string, lib, "SDL_hid_get_manufacturer_string")
	// purego.RegisterLibFunc(&sdlhid_get_product_string, lib, "SDL_hid_get_product_string")
	// purego.RegisterLibFunc(&sdlhid_get_report_descriptor, lib, "SDL_hid_get_report_descriptor")
	// purego.RegisterLibFunc(&sdlhid_get_serial_number_string, lib, "SDL_hid_get_serial_number_string")
	// purego.RegisterLibFunc(&sdlhid_init, lib, "SDL_hid_init")
	// purego.RegisterLibFunc(&sdlhid_open, lib, "SDL_hid_open")
	// purego.RegisterLibFunc(&sdlhid_open_path, lib, "SDL_hid_open_path")
	// purego.RegisterLibFunc(&sdlhid_read, lib, "SDL_hid_read")
	// purego.RegisterLibFunc(&sdlhid_read_timeout, lib, "SDL_hid_read_timeout")
	// purego.RegisterLibFunc(&sdlhid_send_feature_report, lib, "SDL_hid_send_feature_report")
	// purego.RegisterLibFunc(&sdlhid_set_nonblocking, lib, "SDL_hid_set_nonblocking")
	// purego.RegisterLibFunc(&sdlhid_write, lib, "SDL_hid_write")
	purego.RegisterLibFunc(&sdlHideCursor, lib, "SDL_HideCursor")
	purego.RegisterLibFunc(&sdlHideWindow, lib, "SDL_HideWindow")
	// purego.RegisterLibFunc(&sdliconv, lib, "SDL_iconv")
	// purego.RegisterLibFunc(&sdliconv_close, lib, "SDL_iconv_close")
	// purego.RegisterLibFunc(&sdliconv_open, lib, "SDL_iconv_open")
	// purego.RegisterLibFunc(&sdliconv_string, lib, "SDL_iconv_string")
	purego.RegisterLibFunc(&sdlInit, lib, "SDL_Init")
	// purego.RegisterLibFunc(&sdlInitHapticRumble, lib, "SDL_InitHapticRumble")
	// purego.RegisterLibFunc(&sdlInitSubSystem, lib, "SDL_InitSubSystem")
	// purego.RegisterLibFunc(&sdlInsertGPUDebugLabel, lib, "SDL_InsertGPUDebugLabel")
	// purego.RegisterLibFunc(&sdlInsertTrayEntryAt, lib, "SDL_InsertTrayEntryAt")
	purego.RegisterLibFunc(&sdlIOFromConstMem, lib, "SDL_IOFromConstMem")
	// purego.RegisterLibFunc(&sdlIOFromDynamicMem, lib, "SDL_IOFromDynamicMem")
	purego.RegisterLibFunc(&sdlIOFromFile, lib, "SDL_IOFromFile")
	purego.RegisterLibFunc(&sdlIOFromMem, lib, "SDL_IOFromMem")
	// purego.RegisterLibFunc(&sdlIOprintf, lib, "SDL_IOprintf")
	// purego.RegisterLibFunc(&sdlIOvprintf, lib, "SDL_IOvprintf")
	// purego.RegisterLibFunc(&sdlisalnum, lib, "SDL_isalnum")
	// purego.RegisterLibFunc(&sdlisalpha, lib, "SDL_isalpha")
	// purego.RegisterLibFunc(&sdlIsAudioDevicePhysical, lib, "SDL_IsAudioDevicePhysical")
	// purego.RegisterLibFunc(&sdlIsAudioDevicePlayback, lib, "SDL_IsAudioDevicePlayback")
	// purego.RegisterLibFunc(&sdlisblank, lib, "SDL_isblank")
	// purego.RegisterLibFunc(&sdliscntrl, lib, "SDL_iscntrl")
	// purego.RegisterLibFunc(&sdlisdigit, lib, "SDL_isdigit")
	// purego.RegisterLibFunc(&sdlIsGamepad, lib, "SDL_IsGamepad")
	// purego.RegisterLibFunc(&sdlisgraph, lib, "SDL_isgraph")
	// purego.RegisterLibFunc(&sdlisinf, lib, "SDL_isinf")
	// purego.RegisterLibFunc(&sdlisinff, lib, "SDL_isinff")
	// purego.RegisterLibFunc(&sdlIsJoystickHaptic, lib, "SDL_IsJoystickHaptic")
	// purego.RegisterLibFunc(&sdlIsJoystickVirtual, lib, "SDL_IsJoystickVirtual")
	// purego.RegisterLibFunc(&sdlislower, lib, "SDL_islower")
	purego.RegisterLibFunc(&sdlIsMainThread, lib, "SDL_IsMainThread")
	// purego.RegisterLibFunc(&sdlIsMouseHaptic, lib, "SDL_IsMouseHaptic")
	// purego.RegisterLibFunc(&sdlisnan, lib, "SDL_isnan")
	// purego.RegisterLibFunc(&sdlisnanf, lib, "SDL_isnanf")
	// purego.RegisterLibFunc(&sdlisprint, lib, "SDL_isprint")
	// purego.RegisterLibFunc(&sdlispunct, lib, "SDL_ispunct")
	// purego.RegisterLibFunc(&sdlisspace, lib, "SDL_isspace")
	// purego.RegisterLibFunc(&sdlIsTablet, lib, "SDL_IsTablet")
	// purego.RegisterLibFunc(&sdlIsTV, lib, "SDL_IsTV")
	// purego.RegisterLibFunc(&sdlisupper, lib, "SDL_isupper")
	// purego.RegisterLibFunc(&sdlisxdigit, lib, "SDL_isxdigit")
	// purego.RegisterLibFunc(&sdlitoa, lib, "SDL_itoa")
	purego.RegisterLibFunc(&sdlJoystickConnected, lib, "SDL_JoystickConnected")
	purego.RegisterLibFunc(&sdlJoystickEventsEnabled, lib, "SDL_JoystickEventsEnabled")
	// purego.RegisterLibFunc(&sdlKillProcess, lib, "SDL_KillProcess")
	// purego.RegisterLibFunc(&sdllltoa, lib, "SDL_lltoa")
	purego.RegisterLibFunc(&sdlLoadBMP, lib, "SDL_LoadBMP")
	purego.RegisterLibFunc(&sdlLoadBMPIO, lib, "SDL_LoadBMP_IO")
	purego.RegisterLibFunc(&sdlLoadFile, lib, "SDL_LoadFile")
	// purego.RegisterLibFunc(&sdlLoadFile_IO, lib, "SDL_LoadFile_IO")
	// purego.RegisterLibFunc(&sdlLoadFileAsync, lib, "SDL_LoadFileAsync")
	// purego.RegisterLibFunc(&sdlLoadFunction, lib, "SDL_LoadFunction")
	// purego.RegisterLibFunc(&sdlLoadObject, lib, "SDL_LoadObject")
	// purego.RegisterLibFunc(&sdlLoadWAV, lib, "SDL_LoadWAV")
	purego.RegisterLibFunc(&sdlLoadWAVIO, lib, "SDL_LoadWAV_IO")
	// purego.RegisterLibFunc(&sdlLockAudioStream, lib, "SDL_LockAudioStream")
	purego.RegisterLibFunc(&sdlLockJoysticks, lib, "SDL_LockJoysticks")
	// purego.RegisterLibFunc(&sdlLockMutex, lib, "SDL_LockMutex")
	purego.RegisterLibFunc(&sdlLockProperties, lib, "SDL_LockProperties")
	// purego.RegisterLibFunc(&sdlLockRWLockForReading, lib, "SDL_LockRWLockForReading")
	// purego.RegisterLibFunc(&sdlLockRWLockForWriting, lib, "SDL_LockRWLockForWriting")
	// purego.RegisterLibFunc(&sdlLockSpinlock, lib, "SDL_LockSpinlock")
	purego.RegisterLibFunc(&sdlLockSurface, lib, "SDL_LockSurface")
	purego.RegisterLibFunc(&sdlLockTexture, lib, "SDL_LockTexture")
	purego.RegisterLibFunc(&sdlLockTextureToSurface, lib, "SDL_LockTextureToSurface")
	purego.RegisterLibFunc(&sdlLog, lib, "SDL_Log")
	// purego.RegisterLibFunc(&sdllog, lib, "SDL_log")
	// purego.RegisterLibFunc(&sdllog10, lib, "SDL_log10")
	// purego.RegisterLibFunc(&sdllog10f, lib, "SDL_log10f")
	// purego.RegisterLibFunc(&sdlLogCritical, lib, "SDL_LogCritical")
	// purego.RegisterLibFunc(&sdlLogDebug, lib, "SDL_LogDebug")
	purego.RegisterLibFunc(&sdlLogError, lib, "SDL_LogError")
	// purego.RegisterLibFunc(&sdllogf, lib, "SDL_logf")
	// purego.RegisterLibFunc(&sdlLogInfo, lib, "SDL_LogInfo")
	purego.RegisterLibFunc(&sdlLogMessage, lib, "SDL_LogMessage")
	// purego.RegisterLibFunc(&sdlLogMessageV, lib, "SDL_LogMessageV")
	// purego.RegisterLibFunc(&sdlLogTrace, lib, "SDL_LogTrace")
	// purego.RegisterLibFunc(&sdlLogVerbose, lib, "SDL_LogVerbose")
	// purego.RegisterLibFunc(&sdlLogWarn, lib, "SDL_LogWarn")
	// purego.RegisterLibFunc(&sdllround, lib, "SDL_lround")
	// purego.RegisterLibFunc(&sdllroundf, lib, "SDL_lroundf")
	// purego.RegisterLibFunc(&sdlltoa, lib, "SDL_ltoa")
	// purego.RegisterLibFunc(&sdlmain, lib, "SDL_main")
	// purego.RegisterLibFunc(&sdlmalloc, lib, "SDL_malloc")
	purego.RegisterLibFunc(&sdlMapGPUTransferBuffer, lib, "SDL_MapGPUTransferBuffer")
	purego.RegisterLibFunc(&sdlMapRGB, lib, "SDL_MapRGB")
	// purego.RegisterLibFunc(&sdlMapRGBA, lib, "SDL_MapRGBA")
	purego.RegisterLibFunc(&sdlMapSurfaceRGB, lib, "SDL_MapSurfaceRGB")
	// purego.RegisterLibFunc(&sdlMapSurfaceRGBA, lib, "SDL_MapSurfaceRGBA")
	// purego.RegisterLibFunc(&sdlMaximizeWindow, lib, "SDL_MaximizeWindow")
	// purego.RegisterLibFunc(&sdlmemcmp, lib, "SDL_memcmp")
	// purego.RegisterLibFunc(&sdlmemcpy, lib, "SDL_memcpy")
	// purego.RegisterLibFunc(&sdlmemmove, lib, "SDL_memmove")
	// purego.RegisterLibFunc(&sdlMemoryBarrierAcquireFunction, lib, "SDL_MemoryBarrierAcquireFunction")
	// purego.RegisterLibFunc(&sdlMemoryBarrierReleaseFunction, lib, "SDL_MemoryBarrierReleaseFunction")
	// purego.RegisterLibFunc(&sdlmemset, lib, "SDL_memset")
	// purego.RegisterLibFunc(&sdlmemset4, lib, "SDL_memset4")
	// purego.RegisterLibFunc(&sdlMetal_CreateView, lib, "SDL_Metal_CreateView")
	// purego.RegisterLibFunc(&sdlMetal_DestroyView, lib, "SDL_Metal_DestroyView")
	// purego.RegisterLibFunc(&sdlMetal_GetLayer, lib, "SDL_Metal_GetLayer")
	// purego.RegisterLibFunc(&sdlMinimizeWindow, lib, "SDL_MinimizeWindow")
	// purego.RegisterLibFunc(&sdlMixAudio, lib, "SDL_MixAudio")
	// purego.RegisterLibFunc(&sdlmodf, lib, "SDL_modf")
	// purego.RegisterLibFunc(&sdlmodff, lib, "SDL_modff")
	// purego.RegisterLibFunc(&sdlMostSignificantBitIndex32, lib, "SDL_MostSignificantBitIndex32")
	// purego.RegisterLibFunc(&sdlmurmur3_32, lib, "SDL_murmur3_32")
	// purego.RegisterLibFunc(&sdlOnApplicationDidEnterBackground, lib, "SDL_OnApplicationDidEnterBackground")
	// purego.RegisterLibFunc(&sdlOnApplicationDidEnterForeground, lib, "SDL_OnApplicationDidEnterForeground")
	// purego.RegisterLibFunc(&sdlOnApplicationDidReceiveMemoryWarning, lib, "SDL_OnApplicationDidReceiveMemoryWarning")
	// purego.RegisterLibFunc(&sdlOnApplicationWillEnterBackground, lib, "SDL_OnApplicationWillEnterBackground")
	// purego.RegisterLibFunc(&sdlOnApplicationWillEnterForeground, lib, "SDL_OnApplicationWillEnterForeground")
	// purego.RegisterLibFunc(&sdlOnApplicationWillTerminate, lib, "SDL_OnApplicationWillTerminate")
	// purego.RegisterLibFunc(&sdlOpenAudioDevice, lib, "SDL_OpenAudioDevice")
	purego.RegisterLibFunc(&sdlOpenAudioDeviceStream, lib, "SDL_OpenAudioDeviceStream")
	purego.RegisterLibFunc(&sdlOpenCamera, lib, "SDL_OpenCamera")
	// purego.RegisterLibFunc(&sdlOpenFileStorage, lib, "SDL_OpenFileStorage")
	purego.RegisterLibFunc(&sdlOpenGamepad, lib, "SDL_OpenGamepad")
	// purego.RegisterLibFunc(&sdlOpenHaptic, lib, "SDL_OpenHaptic")
	// purego.RegisterLibFunc(&sdlOpenHapticFromJoystick, lib, "SDL_OpenHapticFromJoystick")
	// purego.RegisterLibFunc(&sdlOpenHapticFromMouse, lib, "SDL_OpenHapticFromMouse")
	// purego.RegisterLibFunc(&sdlOpenIO, lib, "SDL_OpenIO")
	purego.RegisterLibFunc(&sdlOpenJoystick, lib, "SDL_OpenJoystick")
	// purego.RegisterLibFunc(&sdlOpenSensor, lib, "SDL_OpenSensor")
	// purego.RegisterLibFunc(&sdlOpenStorage, lib, "SDL_OpenStorage")
	// purego.RegisterLibFunc(&sdlOpenTitleStorage, lib, "SDL_OpenTitleStorage")
	purego.RegisterLibFunc(&sdlOpenURL, lib, "SDL_OpenURL")
	// purego.RegisterLibFunc(&sdlOpenUserStorage, lib, "SDL_OpenUserStorage")
	// purego.RegisterLibFunc(&sdlOutOfMemory, lib, "SDL_OutOfMemory")
	// purego.RegisterLibFunc(&sdlPauseAudioDevice, lib, "SDL_PauseAudioDevice")
	sdlPauseAudioStreamDevice = shared.Get(lib, "SDL_PauseAudioStreamDevice")
	// purego.RegisterLibFunc(&sdlPauseHaptic, lib, "SDL_PauseHaptic")
	purego.RegisterLibFunc(&sdlPeepEvents, lib, "SDL_PeepEvents")
	// purego.RegisterLibFunc(&sdlPlayHapticRumble, lib, "SDL_PlayHapticRumble")
	sdlPollEvent = shared.Get(lib, "SDL_PollEvent")
	// purego.RegisterLibFunc(&sdlPopGPUDebugGroup, lib, "SDL_PopGPUDebugGroup")
	// purego.RegisterLibFunc(&sdlpow, lib, "SDL_pow")
	// purego.RegisterLibFunc(&sdlpowf, lib, "SDL_powf")
	// purego.RegisterLibFunc(&sdlPremultiplyAlpha, lib, "SDL_PremultiplyAlpha")
	// purego.RegisterLibFunc(&sdlPremultiplySurfaceAlpha, lib, "SDL_PremultiplySurfaceAlpha")
	purego.RegisterLibFunc(&sdlPumpEvents, lib, "SDL_PumpEvents")
	purego.RegisterLibFunc(&sdlPushEvent, lib, "SDL_PushEvent")
	// purego.RegisterLibFunc(&sdlPushGPUComputeUniformData, lib, "SDL_PushGPUComputeUniformData")
	// purego.RegisterLibFunc(&sdlPushGPUDebugGroup, lib, "SDL_PushGPUDebugGroup")
	purego.RegisterLibFunc(&sdlPushGPUFragmentUniformData, lib, "SDL_PushGPUFragmentUniformData")
	purego.RegisterLibFunc(&sdlPushGPUVertexUniformData, lib, "SDL_PushGPUVertexUniformData")
	sdlPutAudioStreamData = shared.Get(lib, "SDL_PutAudioStreamData")
	// purego.RegisterLibFunc(&sdlqsort, lib, "SDL_qsort")
	// purego.RegisterLibFunc(&sdlqsort_r, lib, "SDL_qsort_r")
	// purego.RegisterLibFunc(&sdlQueryGPUFence, lib, "SDL_QueryGPUFence")
	purego.RegisterLibFunc(&sdlQuit, lib, "SDL_Quit")
	purego.RegisterLibFunc(&sdlQuitSubSystem, lib, "SDL_QuitSubSystem")
	purego.RegisterLibFunc(&sdlRaiseWindow, lib, "SDL_RaiseWindow")
	// purego.RegisterLibFunc(&sdlrand, lib, "SDL_rand")
	// purego.RegisterLibFunc(&sdlrand_bits, lib, "SDL_rand_bits")
	// purego.RegisterLibFunc(&sdlrand_bits_r, lib, "SDL_rand_bits_r")
	// purego.RegisterLibFunc(&sdlrand_r, lib, "SDL_rand_r")
	// purego.RegisterLibFunc(&sdlrandf, lib, "SDL_randf")
	// purego.RegisterLibFunc(&sdlrandf_r, lib, "SDL_randf_r")
	// purego.RegisterLibFunc(&sdlReadAsyncIO, lib, "SDL_ReadAsyncIO")
	// purego.RegisterLibFunc(&sdlReadIO, lib, "SDL_ReadIO")
	// purego.RegisterLibFunc(&sdlReadProcess, lib, "SDL_ReadProcess")
	// purego.RegisterLibFunc(&sdlReadS16BE, lib, "SDL_ReadS16BE")
	// purego.RegisterLibFunc(&sdlReadS16LE, lib, "SDL_ReadS16LE")
	// purego.RegisterLibFunc(&sdlReadS32BE, lib, "SDL_ReadS32BE")
	// purego.RegisterLibFunc(&sdlReadS32LE, lib, "SDL_ReadS32LE")
	// purego.RegisterLibFunc(&sdlReadS64BE, lib, "SDL_ReadS64BE")
	// purego.RegisterLibFunc(&sdlReadS64LE, lib, "SDL_ReadS64LE")
	// purego.RegisterLibFunc(&sdlReadS8, lib, "SDL_ReadS8")
	// purego.RegisterLibFunc(&sdlReadStorageFile, lib, "SDL_ReadStorageFile")
	// purego.RegisterLibFunc(&sdlReadSurfacePixel, lib, "SDL_ReadSurfacePixel")
	// purego.RegisterLibFunc(&sdlReadSurfacePixelFloat, lib, "SDL_ReadSurfacePixelFloat")
	// purego.RegisterLibFunc(&sdlReadU16BE, lib, "SDL_ReadU16BE")
	// purego.RegisterLibFunc(&sdlReadU16LE, lib, "SDL_ReadU16LE")
	// purego.RegisterLibFunc(&sdlReadU32BE, lib, "SDL_ReadU32BE")
	// purego.RegisterLibFunc(&sdlReadU32LE, lib, "SDL_ReadU32LE")
	// purego.RegisterLibFunc(&sdlReadU64BE, lib, "SDL_ReadU64BE")
	// purego.RegisterLibFunc(&sdlReadU64LE, lib, "SDL_ReadU64LE")
	// purego.RegisterLibFunc(&sdlReadU8, lib, "SDL_ReadU8")
	// purego.RegisterLibFunc(&sdlrealloc, lib, "SDL_realloc")
	purego.RegisterLibFunc(&sdlRegisterEvents, lib, "SDL_RegisterEvents")
	purego.RegisterLibFunc(&sdlReleaseCameraFrame, lib, "SDL_ReleaseCameraFrame")
	purego.RegisterLibFunc(&sdlReleaseGPUBuffer, lib, "SDL_ReleaseGPUBuffer")
	// purego.RegisterLibFunc(&sdlReleaseGPUComputePipeline, lib, "SDL_ReleaseGPUComputePipeline")
	// purego.RegisterLibFunc(&sdlReleaseGPUFence, lib, "SDL_ReleaseGPUFence")
	purego.RegisterLibFunc(&sdlReleaseGPUGraphicsPipeline, lib, "SDL_ReleaseGPUGraphicsPipeline")
	purego.RegisterLibFunc(&sdlReleaseGPUSampler, lib, "SDL_ReleaseGPUSampler")
	purego.RegisterLibFunc(&sdlReleaseGPUShader, lib, "SDL_ReleaseGPUShader")
	purego.RegisterLibFunc(&sdlReleaseGPUTexture, lib, "SDL_ReleaseGPUTexture")
	purego.RegisterLibFunc(&sdlReleaseGPUTransferBuffer, lib, "SDL_ReleaseGPUTransferBuffer")
	purego.RegisterLibFunc(&sdlReleaseWindowFromGPUDevice, lib, "SDL_ReleaseWindowFromGPUDevice")
	// purego.RegisterLibFunc(&sdlReloadGamepadMappings, lib, "SDL_ReloadGamepadMappings")
	purego.RegisterLibFunc(&sdlRemoveEventWatch, lib, "SDL_RemoveEventWatch")
	purego.RegisterLibFunc(&sdlRemoveHintCallback, lib, "SDL_RemoveHintCallback")
	// purego.RegisterLibFunc(&sdlRemovePath, lib, "SDL_RemovePath")
	// purego.RegisterLibFunc(&sdlRemoveStoragePath, lib, "SDL_RemoveStoragePath")
	purego.RegisterLibFunc(&sdlRemoveSurfaceAlternateImages, lib, "SDL_RemoveSurfaceAlternateImages")
	// purego.RegisterLibFunc(&sdlRemoveTimer, lib, "SDL_RemoveTimer")
	// purego.RegisterLibFunc(&sdlRemoveTrayEntry, lib, "SDL_RemoveTrayEntry")
	// purego.RegisterLibFunc(&sdlRenamePath, lib, "SDL_RenamePath")
	// purego.RegisterLibFunc(&sdlRenameStoragePath, lib, "SDL_RenameStoragePath")
	sdlRenderClear = shared.Get(lib, "SDL_RenderClear")
	purego.RegisterLibFunc(&sdlRenderClipEnabled, lib, "SDL_RenderClipEnabled")
	purego.RegisterLibFunc(&sdlRenderCoordinatesFromWindow, lib, "SDL_RenderCoordinatesFromWindow")
	purego.RegisterLibFunc(&sdlRenderCoordinatesToWindow, lib, "SDL_RenderCoordinatesToWindow")
	purego.RegisterLibFunc(&sdlRenderDebugText, lib, "SDL_RenderDebugText")
	purego.RegisterLibFunc(&sdlRenderDebugTextFormat, lib, "SDL_RenderDebugTextFormat")
	sdlRenderFillRect = shared.Get(lib, "SDL_RenderFillRect")
	sdlRenderFillRects = shared.Get(lib, "SDL_RenderFillRects")
	sdlRenderGeometry = shared.Get(lib, "SDL_RenderGeometry")
	sdlRenderGeometryRaw = shared.Get(lib, "SDL_RenderGeometryRaw")
	purego.RegisterLibFunc(&sdlRenderLine, lib, "SDL_RenderLine")
	sdlRenderLines = shared.Get(lib, "SDL_RenderLines")
	purego.RegisterLibFunc(&sdlRenderPoint, lib, "SDL_RenderPoint")
	sdlRenderPoints = shared.Get(lib, "SDL_RenderPoints")
	sdlRenderPresent = shared.Get(lib, "SDL_RenderPresent")
	purego.RegisterLibFunc(&sdlRenderReadPixels, lib, "SDL_RenderReadPixels")
	sdlRenderRect = shared.Get(lib, "SDL_RenderRect")
	sdlRenderRects = shared.Get(lib, "SDL_RenderRects")
	sdlRenderTexture = shared.Get(lib, "SDL_RenderTexture")
	purego.RegisterLibFunc(&sdlRenderTexture9Grid, lib, "SDL_RenderTexture9Grid")
	sdlRenderTextureAffine = shared.Get(lib, "SDL_RenderTextureAffine")
	purego.RegisterLibFunc(&sdlRenderTextureRotated, lib, "SDL_RenderTextureRotated")
	purego.RegisterLibFunc(&sdlRenderTextureTiled, lib, "SDL_RenderTextureTiled")
	purego.RegisterLibFunc(&sdlRenderViewportSet, lib, "SDL_RenderViewportSet")
	// purego.RegisterLibFunc(&sdlReportAssertion, lib, "SDL_ReportAssertion")
	// purego.RegisterLibFunc(&sdlResetAssertionReport, lib, "SDL_ResetAssertionReport")
	purego.RegisterLibFunc(&sdlResetHint, lib, "SDL_ResetHint")
	purego.RegisterLibFunc(&sdlResetHints, lib, "SDL_ResetHints")
	purego.RegisterLibFunc(&sdlResetKeyboard, lib, "SDL_ResetKeyboard")
	purego.RegisterLibFunc(&sdlResetLogPriorities, lib, "SDL_ResetLogPriorities")
	purego.RegisterLibFunc(&sdlRestoreWindow, lib, "SDL_RestoreWindow")
	// purego.RegisterLibFunc(&sdlResumeAudioDevice, lib, "SDL_ResumeAudioDevice")
	sdlResumeAudioStreamDevice = shared.Get(lib, "SDL_ResumeAudioStreamDevice")
	// purego.RegisterLibFunc(&sdlResumeHaptic, lib, "SDL_ResumeHaptic")
	// purego.RegisterLibFunc(&sdlround, lib, "SDL_round")
	// purego.RegisterLibFunc(&sdlroundf, lib, "SDL_roundf")
	// purego.RegisterLibFunc(&sdlRumbleGamepad, lib, "SDL_RumbleGamepad")
	// purego.RegisterLibFunc(&sdlRumbleGamepadTriggers, lib, "SDL_RumbleGamepadTriggers")
	purego.RegisterLibFunc(&sdlRumbleJoystick, lib, "SDL_RumbleJoystick")
	purego.RegisterLibFunc(&sdlRumbleJoystickTriggers, lib, "SDL_RumbleJoystickTriggers")
	// purego.RegisterLibFunc(&sdlRunApp, lib, "SDL_RunApp")
	// purego.RegisterLibFunc(&sdlRunHapticEffect, lib, "SDL_RunHapticEffect")
	// purego.RegisterLibFunc(&sdlRunOnMainThread, lib, "SDL_RunOnMainThread")
	purego.RegisterLibFunc(&sdlSaveBMP, lib, "SDL_SaveBMP")
	purego.RegisterLibFunc(&sdlSaveBMPIO, lib, "SDL_SaveBMP_IO")
	// purego.RegisterLibFunc(&sdlSaveFile, lib, "SDL_SaveFile")
	// purego.RegisterLibFunc(&sdlSaveFile_IO, lib, "SDL_SaveFile_IO")
	// purego.RegisterLibFunc(&sdlscalbn, lib, "SDL_scalbn")
	// purego.RegisterLibFunc(&sdlscalbnf, lib, "SDL_scalbnf")
	purego.RegisterLibFunc(&sdlScaleSurface, lib, "SDL_ScaleSurface")
	purego.RegisterLibFunc(&sdlScreenKeyboardShown, lib, "SDL_ScreenKeyboardShown")
	// purego.RegisterLibFunc(&sdlScreenSaverEnabled, lib, "SDL_ScreenSaverEnabled")
	// purego.RegisterLibFunc(&sdlSeekIO, lib, "SDL_SeekIO")
	// purego.RegisterLibFunc(&sdlSendGamepadEffect, lib, "SDL_SendGamepadEffect")
	purego.RegisterLibFunc(&sdlSendJoystickEffect, lib, "SDL_SendJoystickEffect")
	// purego.RegisterLibFunc(&sdlSendJoystickVirtualSensorData, lib, "SDL_SendJoystickVirtualSensorData")
	// purego.RegisterLibFunc(&sdlSetAppMetadata, lib, "SDL_SetAppMetadata")
	// purego.RegisterLibFunc(&sdlSetAppMetadataProperty, lib, "SDL_SetAppMetadataProperty")
	// purego.RegisterLibFunc(&sdlSetAssertionHandler, lib, "SDL_SetAssertionHandler")
	// purego.RegisterLibFunc(&sdlSetAtomicInt, lib, "SDL_SetAtomicInt")
	// purego.RegisterLibFunc(&sdlSetAtomicPointer, lib, "SDL_SetAtomicPointer")
	// purego.RegisterLibFunc(&sdlSetAtomicU32, lib, "SDL_SetAtomicU32")
	// purego.RegisterLibFunc(&sdlSetAudioDeviceGain, lib, "SDL_SetAudioDeviceGain")
	// purego.RegisterLibFunc(&sdlSetAudioPostmixCallback, lib, "SDL_SetAudioPostmixCallback")
	// purego.RegisterLibFunc(&sdlSetAudioStreamFormat, lib, "SDL_SetAudioStreamFormat")
	// purego.RegisterLibFunc(&sdlSetAudioStreamFrequencyRatio, lib, "SDL_SetAudioStreamFrequencyRatio")
	// purego.RegisterLibFunc(&sdlSetAudioStreamGain, lib, "SDL_SetAudioStreamGain")
	// purego.RegisterLibFunc(&sdlSetAudioStreamGetCallback, lib, "SDL_SetAudioStreamGetCallback")
	// purego.RegisterLibFunc(&sdlSetAudioStreamInputChannelMap, lib, "SDL_SetAudioStreamInputChannelMap")
	// purego.RegisterLibFunc(&sdlSetAudioStreamOutputChannelMap, lib, "SDL_SetAudioStreamOutputChannelMap")
	// purego.RegisterLibFunc(&sdlSetAudioStreamPutCallback, lib, "SDL_SetAudioStreamPutCallback")
	purego.RegisterLibFunc(&sdlSetBooleanProperty, lib, "SDL_SetBooleanProperty")
	// purego.RegisterLibFunc(&sdlSetClipboardData, lib, "SDL_SetClipboardData")
	// purego.RegisterLibFunc(&sdlSetClipboardText, lib, "SDL_SetClipboardText")
	// purego.RegisterLibFunc(&sdlSetCurrentThreadPriority, lib, "SDL_SetCurrentThreadPriority")
	purego.RegisterLibFunc(&sdlSetCursor, lib, "SDL_SetCursor")
	// purego.RegisterLibFunc(&sdlsetenv_unsafe, lib, "SDL_setenv_unsafe")
	// purego.RegisterLibFunc(&sdlSetEnvironmentVariable, lib, "SDL_SetEnvironmentVariable")
	purego.RegisterLibFunc(&sdlSetError, lib, "SDL_SetError")
	// purego.RegisterLibFunc(&sdlSetErrorV, lib, "SDL_SetErrorV")
	purego.RegisterLibFunc(&sdlSetEventEnabled, lib, "SDL_SetEventEnabled")
	purego.RegisterLibFunc(&sdlSetEventFilter, lib, "SDL_SetEventFilter")
	purego.RegisterLibFunc(&sdlSetFloatProperty, lib, "SDL_SetFloatProperty")
	// purego.RegisterLibFunc(&sdlSetGamepadEventsEnabled, lib, "SDL_SetGamepadEventsEnabled")
	// purego.RegisterLibFunc(&sdlSetGamepadLED, lib, "SDL_SetGamepadLED")
	// purego.RegisterLibFunc(&sdlSetGamepadMapping, lib, "SDL_SetGamepadMapping")
	// purego.RegisterLibFunc(&sdlSetGamepadPlayerIndex, lib, "SDL_SetGamepadPlayerIndex")
	// purego.RegisterLibFunc(&sdlSetGamepadSensorEnabled, lib, "SDL_SetGamepadSensorEnabled")
	// purego.RegisterLibFunc(&sdlSetGPUAllowedFramesInFlight, lib, "SDL_SetGPUAllowedFramesInFlight")
	// purego.RegisterLibFunc(&sdlSetGPUBlendConstants, lib, "SDL_SetGPUBlendConstants")
	purego.RegisterLibFunc(&sdlSetGPUBufferName, lib, "SDL_SetGPUBufferName")
	purego.RegisterLibFunc(&sdlSetGPUScissor, lib, "SDL_SetGPUScissor")
	// purego.RegisterLibFunc(&sdlSetGPUStencilReference, lib, "SDL_SetGPUStencilReference")
	purego.RegisterLibFunc(&sdlSetGPUSwapchainParameters, lib, "SDL_SetGPUSwapchainParameters")
	// purego.RegisterLibFunc(&sdlSetGPUTextureName, lib, "SDL_SetGPUTextureName")
	purego.RegisterLibFunc(&sdlSetGPUViewport, lib, "SDL_SetGPUViewport")
	// purego.RegisterLibFunc(&sdlSetHapticAutocenter, lib, "SDL_SetHapticAutocenter")
	// purego.RegisterLibFunc(&sdlSetHapticGain, lib, "SDL_SetHapticGain")
	purego.RegisterLibFunc(&sdlSetHint, lib, "SDL_SetHint")
	purego.RegisterLibFunc(&sdlSetHintWithPriority, lib, "SDL_SetHintWithPriority")
	// purego.RegisterLibFunc(&sdlSetInitialized, lib, "SDL_SetInitialized")
	purego.RegisterLibFunc(&sdlSetJoystickEventsEnabled, lib, "SDL_SetJoystickEventsEnabled")
	purego.RegisterLibFunc(&sdlSetJoystickLED, lib, "SDL_SetJoystickLED")
	purego.RegisterLibFunc(&sdlSetJoystickPlayerIndex, lib, "SDL_SetJoystickPlayerIndex")
	// purego.RegisterLibFunc(&sdlSetJoystickVirtualAxis, lib, "SDL_SetJoystickVirtualAxis")
	// purego.RegisterLibFunc(&sdlSetJoystickVirtualBall, lib, "SDL_SetJoystickVirtualBall")
	// purego.RegisterLibFunc(&sdlSetJoystickVirtualButton, lib, "SDL_SetJoystickVirtualButton")
	// purego.RegisterLibFunc(&sdlSetJoystickVirtualHat, lib, "SDL_SetJoystickVirtualHat")
	// purego.RegisterLibFunc(&sdlSetJoystickVirtualTouchpad, lib, "SDL_SetJoystickVirtualTouchpad")
	// purego.RegisterLibFunc(&sdlSetLinuxThreadPriority, lib, "SDL_SetLinuxThreadPriority")
	// purego.RegisterLibFunc(&sdlSetLinuxThreadPriorityAndPolicy, lib, "SDL_SetLinuxThreadPriorityAndPolicy")
	// purego.RegisterLibFunc(&sdlSetLogOutputFunction, lib, "SDL_SetLogOutputFunction")
	purego.RegisterLibFunc(&sdlSetLogPriorities, lib, "SDL_SetLogPriorities")
	purego.RegisterLibFunc(&sdlSetLogPriority, lib, "SDL_SetLogPriority")
	// purego.RegisterLibFunc(&sdlSetLogPriorityPrefix, lib, "SDL_SetLogPriorityPrefix")
	// purego.RegisterLibFunc(&sdlSetMainReady, lib, "SDL_SetMainReady")
	// purego.RegisterLibFunc(&sdlSetMemoryFunctions, lib, "SDL_SetMemoryFunctions")
	purego.RegisterLibFunc(&sdlSetModState, lib, "SDL_SetModState")
	purego.RegisterLibFunc(&sdlSetNumberProperty, lib, "SDL_SetNumberProperty")
	purego.RegisterLibFunc(&sdlSetPaletteColors, lib, "SDL_SetPaletteColors")
	purego.RegisterLibFunc(&sdlSetPointerProperty, lib, "SDL_SetPointerProperty")
	purego.RegisterLibFunc(&sdlSetPointerPropertyWithCleanup, lib, "SDL_SetPointerPropertyWithCleanup")
	// purego.RegisterLibFunc(&sdlSetPrimarySelectionText, lib, "SDL_SetPrimarySelectionText")
	purego.RegisterLibFunc(&sdlSetRenderClipRect, lib, "SDL_SetRenderClipRect")
	purego.RegisterLibFunc(&sdlSetRenderColorScale, lib, "SDL_SetRenderColorScale")
	purego.RegisterLibFunc(&sdlSetRenderDrawBlendMode, lib, "SDL_SetRenderDrawBlendMode")
	sdlSetRenderDrawColor = shared.Get(lib, "SDL_SetRenderDrawColor")
	purego.RegisterLibFunc(&sdlSetRenderDrawColorFloat, lib, "SDL_SetRenderDrawColorFloat")
	purego.RegisterLibFunc(&sdlSetRenderLogicalPresentation, lib, "SDL_SetRenderLogicalPresentation")
	purego.RegisterLibFunc(&sdlSetRenderScale, lib, "SDL_SetRenderScale")
	purego.RegisterLibFunc(&sdlSetRenderTarget, lib, "SDL_SetRenderTarget")
	purego.RegisterLibFunc(&sdlSetRenderViewport, lib, "SDL_SetRenderViewport")
	purego.RegisterLibFunc(&sdlSetRenderVSync, lib, "SDL_SetRenderVSync")
	purego.RegisterLibFunc(&sdlSetScancodeName, lib, "SDL_SetScancodeName")
	purego.RegisterLibFunc(&sdlSetStringProperty, lib, "SDL_SetStringProperty")
	purego.RegisterLibFunc(&sdlSetSurfaceAlphaMod, lib, "SDL_SetSurfaceAlphaMod")
	purego.RegisterLibFunc(&sdlSetSurfaceBlendMode, lib, "SDL_SetSurfaceBlendMode")
	purego.RegisterLibFunc(&sdlSetSurfaceClipRect, lib, "SDL_SetSurfaceClipRect")
	purego.RegisterLibFunc(&sdlSetSurfaceColorKey, lib, "SDL_SetSurfaceColorKey")
	purego.RegisterLibFunc(&sdlSetSurfaceColorMod, lib, "SDL_SetSurfaceColorMod")
	purego.RegisterLibFunc(&sdlSetSurfaceColorspace, lib, "SDL_SetSurfaceColorspace")
	purego.RegisterLibFunc(&sdlSetSurfacePalette, lib, "SDL_SetSurfacePalette")
	purego.RegisterLibFunc(&sdlSetSurfaceRLE, lib, "SDL_SetSurfaceRLE")
	purego.RegisterLibFunc(&sdlSetTextInputArea, lib, "SDL_SetTextInputArea")
	sdlSetTextureAlphaMod = shared.Get(lib, "SDL_SetTextureAlphaMod")
	purego.RegisterLibFunc(&sdlSetTextureAlphaModFloat, lib, "SDL_SetTextureAlphaModFloat")
	purego.RegisterLibFunc(&sdlSetTextureBlendMode, lib, "SDL_SetTextureBlendMode")
	sdlSetTextureColorMod = shared.Get(lib, "SDL_SetTextureColorMod")
	purego.RegisterLibFunc(&sdlSetTextureColorModFloat, lib, "SDL_SetTextureColorModFloat")
	purego.RegisterLibFunc(&sdlSetTextureScaleMode, lib, "SDL_SetTextureScaleMode")
	// purego.RegisterLibFunc(&sdlSetTLS, lib, "SDL_SetTLS")
	// purego.RegisterLibFunc(&sdlSetTrayEntryCallback, lib, "SDL_SetTrayEntryCallback")
	// purego.RegisterLibFunc(&sdlSetTrayEntryChecked, lib, "SDL_SetTrayEntryChecked")
	// purego.RegisterLibFunc(&sdlSetTrayEntryEnabled, lib, "SDL_SetTrayEntryEnabled")
	// purego.RegisterLibFunc(&sdlSetTrayEntryLabel, lib, "SDL_SetTrayEntryLabel")
	// purego.RegisterLibFunc(&sdlSetTrayIcon, lib, "SDL_SetTrayIcon")
	// purego.RegisterLibFunc(&sdlSetTrayTooltip, lib, "SDL_SetTrayTooltip")
	purego.RegisterLibFunc(&sdlSetWindowAlwaysOnTop, lib, "SDL_SetWindowAlwaysOnTop")
	// purego.RegisterLibFunc(&sdlSetWindowAspectRatio, lib, "SDL_SetWindowAspectRatio")
	purego.RegisterLibFunc(&sdlSetWindowBordered, lib, "SDL_SetWindowBordered")
	purego.RegisterLibFunc(&sdlSetWindowFocusable, lib, "SDL_SetWindowFocusable")
	purego.RegisterLibFunc(&sdlSetWindowFullscreen, lib, "SDL_SetWindowFullscreen")
	purego.RegisterLibFunc(&sdlSetWindowFullscreenMode, lib, "SDL_SetWindowFullscreenMode")
	purego.RegisterLibFunc(&sdlSetWindowHitTest, lib, "SDL_SetWindowHitTest")
	purego.RegisterLibFunc(&sdlSetWindowIcon, lib, "SDL_SetWindowIcon")
	purego.RegisterLibFunc(&sdlSetWindowKeyboardGrab, lib, "SDL_SetWindowKeyboardGrab")
	// purego.RegisterLibFunc(&sdlSetWindowMaximumSize, lib, "SDL_SetWindowMaximumSize")
	// purego.RegisterLibFunc(&sdlSetWindowMinimumSize, lib, "SDL_SetWindowMinimumSize")
	// purego.RegisterLibFunc(&sdlSetWindowModal, lib, "SDL_SetWindowModal")
	purego.RegisterLibFunc(&sdlSetWindowMouseGrab, lib, "SDL_SetWindowMouseGrab")
	// purego.RegisterLibFunc(&sdlSetWindowMouseRect, lib, "SDL_SetWindowMouseRect")
	purego.RegisterLibFunc(&sdlSetWindowOpacity, lib, "SDL_SetWindowOpacity")
	// purego.RegisterLibFunc(&sdlSetWindowParent, lib, "SDL_SetWindowParent")
	purego.RegisterLibFunc(&sdlSetWindowPosition, lib, "SDL_SetWindowPosition")
	purego.RegisterLibFunc(&sdlSetWindowRelativeMouseMode, lib, "SDL_SetWindowRelativeMouseMode")
	purego.RegisterLibFunc(&sdlSetWindowResizable, lib, "SDL_SetWindowResizable")
	// purego.RegisterLibFunc(&sdlSetWindowShape, lib, "SDL_SetWindowShape")
	purego.RegisterLibFunc(&sdlSetWindowSize, lib, "SDL_SetWindowSize")
	// purego.RegisterLibFunc(&sdlSetWindowSurfaceVSync, lib, "SDL_SetWindowSurfaceVSync")
	purego.RegisterLibFunc(&sdlSetWindowTitle, lib, "SDL_SetWindowTitle")
	// purego.RegisterLibFunc(&sdlSetX11EventHook, lib, "SDL_SetX11EventHook")
	// purego.RegisterLibFunc(&sdlShouldInit, lib, "SDL_ShouldInit")
	// purego.RegisterLibFunc(&sdlShouldQuit, lib, "SDL_ShouldQuit")
	purego.RegisterLibFunc(&sdlShowCursor, lib, "SDL_ShowCursor")
	purego.RegisterLibFunc(&sdlShowFileDialogWithProperties, lib, "SDL_ShowFileDialogWithProperties")
	purego.RegisterLibFunc(&sdlShowMessageBox, lib, "SDL_ShowMessageBox")
	purego.RegisterLibFunc(&sdlShowOpenFileDialog, lib, "SDL_ShowOpenFileDialog")
	purego.RegisterLibFunc(&sdlShowOpenFolderDialog, lib, "SDL_ShowOpenFolderDialog")
	purego.RegisterLibFunc(&sdlShowSaveFileDialog, lib, "SDL_ShowSaveFileDialog")
	purego.RegisterLibFunc(&sdlShowSimpleMessageBox, lib, "SDL_ShowSimpleMessageBox")
	purego.RegisterLibFunc(&sdlShowWindow, lib, "SDL_ShowWindow")
	// purego.RegisterLibFunc(&sdlShowWindowSystemMenu, lib, "SDL_ShowWindowSystemMenu")
	// purego.RegisterLibFunc(&sdlSignalAsyncIOQueue, lib, "SDL_SignalAsyncIOQueue")
	// purego.RegisterLibFunc(&sdlSignalCondition, lib, "SDL_SignalCondition")
	// purego.RegisterLibFunc(&sdlSignalSemaphore, lib, "SDL_SignalSemaphore")
	// purego.RegisterLibFunc(&sdlsin, lib, "SDL_sin")
	// purego.RegisterLibFunc(&sdlsinf, lib, "SDL_sinf")
	// purego.RegisterLibFunc(&sdlsize_add_check_overflow, lib, "SDL_size_add_check_overflow")
	// purego.RegisterLibFunc(&sdlsize_add_check_overflow_builtin, lib, "SDL_size_add_check_overflow_builtin")
	// purego.RegisterLibFunc(&sdlsize_mul_check_overflow, lib, "SDL_size_mul_check_overflow")
	// purego.RegisterLibFunc(&sdlsize_mul_check_overflow_builtin, lib, "SDL_size_mul_check_overflow_builtin")
	// purego.RegisterLibFunc(&sdlsnprintf, lib, "SDL_snprintf")
	// purego.RegisterLibFunc(&sdlsqrt, lib, "SDL_sqrt")
	// purego.RegisterLibFunc(&sdlsqrtf, lib, "SDL_sqrtf")
	// purego.RegisterLibFunc(&sdlsrand, lib, "SDL_srand")
	// purego.RegisterLibFunc(&sdlsscanf, lib, "SDL_sscanf")
	purego.RegisterLibFunc(&sdlStartTextInput, lib, "SDL_StartTextInput")
	purego.RegisterLibFunc(&sdlStartTextInputWithProperties, lib, "SDL_StartTextInputWithProperties")
	// purego.RegisterLibFunc(&sdlStepBackUTF8, lib, "SDL_StepBackUTF8")
	// purego.RegisterLibFunc(&sdlStepUTF8, lib, "SDL_StepUTF8")
	// purego.RegisterLibFunc(&sdlStopHapticEffect, lib, "SDL_StopHapticEffect")
	// purego.RegisterLibFunc(&sdlStopHapticEffects, lib, "SDL_StopHapticEffects")
	// purego.RegisterLibFunc(&sdlStopHapticRumble, lib, "SDL_StopHapticRumble")
	purego.RegisterLibFunc(&sdlStopTextInput, lib, "SDL_StopTextInput")
	// purego.RegisterLibFunc(&sdlStorageReady, lib, "SDL_StorageReady")
	// purego.RegisterLibFunc(&sdlstrcasecmp, lib, "SDL_strcasecmp")
	// purego.RegisterLibFunc(&sdlstrcasestr, lib, "SDL_strcasestr")
	// purego.RegisterLibFunc(&sdlstrchr, lib, "SDL_strchr")
	// purego.RegisterLibFunc(&sdlstrcmp, lib, "SDL_strcmp")
	// purego.RegisterLibFunc(&sdlstrdup, lib, "SDL_strdup")
	// purego.RegisterLibFunc(&sdlStringToGUID, lib, "SDL_StringToGUID")
	// purego.RegisterLibFunc(&sdlstrlcat, lib, "SDL_strlcat")
	// purego.RegisterLibFunc(&sdlstrlcpy, lib, "SDL_strlcpy")
	// purego.RegisterLibFunc(&sdlstrlen, lib, "SDL_strlen")
	// purego.RegisterLibFunc(&sdlstrlwr, lib, "SDL_strlwr")
	// purego.RegisterLibFunc(&sdlstrncasecmp, lib, "SDL_strncasecmp")
	// purego.RegisterLibFunc(&sdlstrncmp, lib, "SDL_strncmp")
	// purego.RegisterLibFunc(&sdlstrndup, lib, "SDL_strndup")
	// purego.RegisterLibFunc(&sdlstrnlen, lib, "SDL_strnlen")
	// purego.RegisterLibFunc(&sdlstrnstr, lib, "SDL_strnstr")
	// purego.RegisterLibFunc(&sdlstrpbrk, lib, "SDL_strpbrk")
	// purego.RegisterLibFunc(&sdlstrrchr, lib, "SDL_strrchr")
	// purego.RegisterLibFunc(&sdlstrrev, lib, "SDL_strrev")
	purego.RegisterLibFunc(&sdlstrstr, lib, "SDL_strstr")
	// purego.RegisterLibFunc(&sdlstrtod, lib, "SDL_strtod")
	// purego.RegisterLibFunc(&sdlstrtok_r, lib, "SDL_strtok_r")
	// purego.RegisterLibFunc(&sdlstrtol, lib, "SDL_strtol")
	// purego.RegisterLibFunc(&sdlstrtoll, lib, "SDL_strtoll")
	// purego.RegisterLibFunc(&sdlstrtoul, lib, "SDL_strtoul")
	// purego.RegisterLibFunc(&sdlstrtoull, lib, "SDL_strtoull")
	// purego.RegisterLibFunc(&sdlstrupr, lib, "SDL_strupr")
	purego.RegisterLibFunc(&sdlSubmitGPUCommandBuffer, lib, "SDL_SubmitGPUCommandBuffer")
	// purego.RegisterLibFunc(&sdlSubmitGPUCommandBufferAndAcquireFence, lib, "SDL_SubmitGPUCommandBufferAndAcquireFence")
	purego.RegisterLibFunc(&sdlSurfaceHasAlternateImages, lib, "SDL_SurfaceHasAlternateImages")
	purego.RegisterLibFunc(&sdlSurfaceHasColorKey, lib, "SDL_SurfaceHasColorKey")
	purego.RegisterLibFunc(&sdlSurfaceHasRLE, lib, "SDL_SurfaceHasRLE")
	// purego.RegisterLibFunc(&sdlSwapFloat, lib, "SDL_SwapFloat")
	// purego.RegisterLibFunc(&sdlswprintf, lib, "SDL_swprintf")
	purego.RegisterLibFunc(&sdlSyncWindow, lib, "SDL_SyncWindow")
	// purego.RegisterLibFunc(&sdltan, lib, "SDL_tan")
	// purego.RegisterLibFunc(&sdltanf, lib, "SDL_tanf")
	// purego.RegisterLibFunc(&sdlTellIO, lib, "SDL_TellIO")
	purego.RegisterLibFunc(&sdlTextInputActive, lib, "SDL_TextInputActive")
	// purego.RegisterLibFunc(&sdlTimeFromWindows, lib, "SDL_TimeFromWindows")
	// purego.RegisterLibFunc(&sdlTimeToDateTime, lib, "SDL_TimeToDateTime")
	// purego.RegisterLibFunc(&sdlTimeToWindows, lib, "SDL_TimeToWindows")
	// purego.RegisterLibFunc(&sdltolower, lib, "SDL_tolower")
	// purego.RegisterLibFunc(&sdltoupper, lib, "SDL_toupper")
	// purego.RegisterLibFunc(&sdltrunc, lib, "SDL_trunc")
	// purego.RegisterLibFunc(&sdltruncf, lib, "SDL_truncf")
	// purego.RegisterLibFunc(&sdlTryLockMutex, lib, "SDL_TryLockMutex")
	// purego.RegisterLibFunc(&sdlTryLockRWLockForReading, lib, "SDL_TryLockRWLockForReading")
	// purego.RegisterLibFunc(&sdlTryLockRWLockForWriting, lib, "SDL_TryLockRWLockForWriting")
	// purego.RegisterLibFunc(&sdlTryLockSpinlock, lib, "SDL_TryLockSpinlock")
	// purego.RegisterLibFunc(&sdlTryWaitSemaphore, lib, "SDL_TryWaitSemaphore")
	// purego.RegisterLibFunc(&sdlUCS4ToUTF8, lib, "SDL_UCS4ToUTF8")
	// purego.RegisterLibFunc(&sdluitoa, lib, "SDL_uitoa")
	// purego.RegisterLibFunc(&sdlulltoa, lib, "SDL_ulltoa")
	// purego.RegisterLibFunc(&sdlultoa, lib, "SDL_ultoa")
	// purego.RegisterLibFunc(&sdlUnbindAudioStream, lib, "SDL_UnbindAudioStream")
	// purego.RegisterLibFunc(&sdlUnbindAudioStreams, lib, "SDL_UnbindAudioStreams")
	// purego.RegisterLibFunc(&sdlUnloadObject, lib, "SDL_UnloadObject")
	// purego.RegisterLibFunc(&sdlUnlockAudioStream, lib, "SDL_UnlockAudioStream")
	purego.RegisterLibFunc(&sdlUnlockJoysticks, lib, "SDL_UnlockJoysticks")
	// purego.RegisterLibFunc(&sdlUnlockMutex, lib, "SDL_UnlockMutex")
	purego.RegisterLibFunc(&sdlUnlockProperties, lib, "SDL_UnlockProperties")
	// purego.RegisterLibFunc(&sdlUnlockRWLock, lib, "SDL_UnlockRWLock")
	// purego.RegisterLibFunc(&sdlUnlockSpinlock, lib, "SDL_UnlockSpinlock")
	purego.RegisterLibFunc(&sdlUnlockSurface, lib, "SDL_UnlockSurface")
	purego.RegisterLibFunc(&sdlUnlockTexture, lib, "SDL_UnlockTexture")
	purego.RegisterLibFunc(&sdlUnmapGPUTransferBuffer, lib, "SDL_UnmapGPUTransferBuffer")
	// purego.RegisterLibFunc(&sdlunsetenv_unsafe, lib, "SDL_unsetenv_unsafe")
	// purego.RegisterLibFunc(&sdlUnsetEnvironmentVariable, lib, "SDL_UnsetEnvironmentVariable")
	// purego.RegisterLibFunc(&sdlUpdateGamepads, lib, "SDL_UpdateGamepads")
	// purego.RegisterLibFunc(&sdlUpdateHapticEffect, lib, "SDL_UpdateHapticEffect")
	purego.RegisterLibFunc(&sdlUpdateJoysticks, lib, "SDL_UpdateJoysticks")
	sdlUpdateNVTexture = shared.Get(lib, "SDL_UpdateNVTexture")
	// purego.RegisterLibFunc(&sdlUpdateSensors, lib, "SDL_UpdateSensors")
	sdlUpdateTexture = shared.Get(lib, "SDL_UpdateTexture")
	purego.RegisterLibFunc(&sdlUpdateWindowSurface, lib, "SDL_UpdateWindowSurface")
	// purego.RegisterLibFunc(&sdlUpdateWindowSurfaceRects, lib, "SDL_UpdateWindowSurfaceRects")
	sdlUpdateYUVTexture = shared.Get(lib, "SDL_UpdateYUVTexture")
	purego.RegisterLibFunc(&sdlUploadToGPUBuffer, lib, "SDL_UploadToGPUBuffer")
	purego.RegisterLibFunc(&sdlUploadToGPUTexture, lib, "SDL_UploadToGPUTexture")
	// purego.RegisterLibFunc(&sdlutf8strlcpy, lib, "SDL_utf8strlcpy")
	// purego.RegisterLibFunc(&sdlutf8strlen, lib, "SDL_utf8strlen")
	// purego.RegisterLibFunc(&sdlutf8strnlen, lib, "SDL_utf8strnlen")
	// purego.RegisterLibFunc(&sdlvasprintf, lib, "SDL_vasprintf")
	// purego.RegisterLibFunc(&sdlvsnprintf, lib, "SDL_vsnprintf")
	// purego.RegisterLibFunc(&sdlvsscanf, lib, "SDL_vsscanf")
	// purego.RegisterLibFunc(&sdlvswprintf, lib, "SDL_vswprintf")
	purego.RegisterLibFunc(&sdlWaitAndAcquireGPUSwapchainTexture, lib, "SDL_WaitAndAcquireGPUSwapchainTexture")
	// purego.RegisterLibFunc(&sdlWaitAsyncIOResult, lib, "SDL_WaitAsyncIOResult")
	// purego.RegisterLibFunc(&sdlWaitCondition, lib, "SDL_WaitCondition")
	// purego.RegisterLibFunc(&sdlWaitConditionTimeout, lib, "SDL_WaitConditionTimeout")
	purego.RegisterLibFunc(&sdlWaitEvent, lib, "SDL_WaitEvent")
	purego.RegisterLibFunc(&sdlWaitEventTimeout, lib, "SDL_WaitEventTimeout")
	// purego.RegisterLibFunc(&sdlWaitForGPUFences, lib, "SDL_WaitForGPUFences")
	// purego.RegisterLibFunc(&sdlWaitForGPUIdle, lib, "SDL_WaitForGPUIdle")
	// purego.RegisterLibFunc(&sdlWaitForGPUSwapchain, lib, "SDL_WaitForGPUSwapchain")
	// purego.RegisterLibFunc(&sdlWaitProcess, lib, "SDL_WaitProcess")
	// purego.RegisterLibFunc(&sdlWaitSemaphore, lib, "SDL_WaitSemaphore")
	// purego.RegisterLibFunc(&sdlWaitSemaphoreTimeout, lib, "SDL_WaitSemaphoreTimeout")
	// purego.RegisterLibFunc(&sdlWaitThread, lib, "SDL_WaitThread")
	purego.RegisterLibFunc(&sdlWarpMouseGlobal, lib, "SDL_WarpMouseGlobal")
	purego.RegisterLibFunc(&sdlWarpMouseInWindow, lib, "SDL_WarpMouseInWindow")
	// purego.RegisterLibFunc(&sdlWasInit, lib, "SDL_WasInit")
	// purego.RegisterLibFunc(&sdlwcscasecmp, lib, "SDL_wcscasecmp")
	// purego.RegisterLibFunc(&sdlwcscmp, lib, "SDL_wcscmp")
	// purego.RegisterLibFunc(&sdlwcsdup, lib, "SDL_wcsdup")
	// purego.RegisterLibFunc(&sdlwcslcat, lib, "SDL_wcslcat")
	// purego.RegisterLibFunc(&sdlwcslcpy, lib, "SDL_wcslcpy")
	// purego.RegisterLibFunc(&sdlwcslen, lib, "SDL_wcslen")
	// purego.RegisterLibFunc(&sdlwcsncasecmp, lib, "SDL_wcsncasecmp")
	// purego.RegisterLibFunc(&sdlwcsncmp, lib, "SDL_wcsncmp")
	// purego.RegisterLibFunc(&sdlwcsnlen, lib, "SDL_wcsnlen")
	// purego.RegisterLibFunc(&sdlwcsnstr, lib, "SDL_wcsnstr")
	// purego.RegisterLibFunc(&sdlwcsstr, lib, "SDL_wcsstr")
	// purego.RegisterLibFunc(&sdlwcstol, lib, "SDL_wcstol")
	// purego.RegisterLibFunc(&sdlWindowHasSurface, lib, "SDL_WindowHasSurface")
	purego.RegisterLibFunc(&sdlWindowSupportsGPUPresentMode, lib, "SDL_WindowSupportsGPUPresentMode")
	// purego.RegisterLibFunc(&sdlWindowSupportsGPUSwapchainComposition, lib, "SDL_WindowSupportsGPUSwapchainComposition")
	// purego.RegisterLibFunc(&sdlWriteAsyncIO, lib, "SDL_WriteAsyncIO")
	// purego.RegisterLibFunc(&sdlWriteIO, lib, "SDL_WriteIO")
	// purego.RegisterLibFunc(&sdlWriteS16BE, lib, "SDL_WriteS16BE")
	// purego.RegisterLibFunc(&sdlWriteS16LE, lib, "SDL_WriteS16LE")
	// purego.RegisterLibFunc(&sdlWriteS32BE, lib, "SDL_WriteS32BE")
	// purego.RegisterLibFunc(&sdlWriteS32LE, lib, "SDL_WriteS32LE")
	// purego.RegisterLibFunc(&sdlWriteS64BE, lib, "SDL_WriteS64BE")
	// purego.RegisterLibFunc(&sdlWriteS64LE, lib, "SDL_WriteS64LE")
	// purego.RegisterLibFunc(&sdlWriteS8, lib, "SDL_WriteS8")
	// purego.RegisterLibFunc(&sdlWriteStorageFile, lib, "SDL_WriteStorageFile")
	// purego.RegisterLibFunc(&sdlWriteSurfacePixel, lib, "SDL_WriteSurfacePixel")
	// purego.RegisterLibFunc(&sdlWriteSurfacePixelFloat, lib, "SDL_WriteSurfacePixelFloat")
	// purego.RegisterLibFunc(&sdlWriteU16BE, lib, "SDL_WriteU16BE")
	// purego.RegisterLibFunc(&sdlWriteU16LE, lib, "SDL_WriteU16LE")
	// purego.RegisterLibFunc(&sdlWriteU32BE, lib, "SDL_WriteU32BE")
	// purego.RegisterLibFunc(&sdlWriteU32LE, lib, "SDL_WriteU32LE")
	// purego.RegisterLibFunc(&sdlWriteU64BE, lib, "SDL_WriteU64BE")
	// purego.RegisterLibFunc(&sdlWriteU64LE, lib, "SDL_WriteU64LE")
	// purego.RegisterLibFunc(&sdlWriteU8, lib, "SDL_WriteU8")
}
