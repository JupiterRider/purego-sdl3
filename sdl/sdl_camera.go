package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

// [CameraPosition] is a structure specifying the position of camera in relation to system devices.
//
// [CameraPosition]: https://wiki.libsdl.org/SDL3/SDL_CameraPosition
type CameraPosition uint32

const (
	CameraPositionUnknown CameraPosition = iota
	CameraPositionFrontFacing
	CameraPositionBackFacing
)

// [CameraID] is a unique ID for a camera device for the time it is connected to the system, and is never reused for the lifetime of the application.
//
// [CameraID]: https://wiki.libsdl.org/SDL3/SDL_CameraID
type CameraID uint32

// [Camera] is a structure specifying the opaque structure used to identify an opened SDL cameras.
//
// [Camera]: https://wiki.libsdl.org/SDL3/SDL_Camera
type Camera struct{}

// [CameraSpec] is a structure specifying the details of an output format for a camera devices.
//
// [CameraSpec]: https://wiki.libsdl.org/SDL3/SDL_CameraSpec
type CameraSpec struct {
	Format               PixelFormat // Frame format.
	Colorspace           Colorspace  // Frame colorspace.
	Width                int32       // Frame width.
	Height               int32       // Frame height.
	FramerateNumerator   int32       // Frame rate numerator ((num / denom) == FPS, (denom / num) == duration in seconds).
	FramerateDenominator int32       // Frame rate denominator ((num / denom) == FPS, (denom / num) == duration in seconds).
}

// [AcquireCameraFrame] acquires a frame.
//
// [AcquireCameraFrame]: https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame
func AcquireCameraFrame(camera *Camera, timestampNS *uint64) *Surface {
	return sdlAcquireCameraFrame(camera, timestampNS)
}

// [CloseCamera] shuts down camera processing and close the camera device.
//
// [CloseCamera]: https://wiki.libsdl.org/SDL3/SDL_CloseCamera
func CloseCamera(camera *Camera) {
	sdlCloseCamera(camera)
}

// [GetCameraDriver] returns the name of the camera driver at the requested index, or "" if an invalid index was specified.
//
// [GetCameraDriver]: https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver
func GetCameraDriver(index int32) string {
	return sdlGetCameraDriver(index)
}

// [GetCameraFormat] gets the spec that a camera is using when generating images.
//
// [GetCameraFormat]: https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat
func GetCameraFormat(camera *Camera, spec *CameraSpec) bool {
	return sdlGetCameraFormat(camera, spec)
}

// [GetCameraID] gets the instance ID of an opened camera.
//
// [GetCameraID]: https://wiki.libsdl.org/SDL3/SDL_GetCameraID
func GetCameraID(camera *Camera) CameraID {
	return sdlGetCameraID(camera)
}

// [GetCameraName] returns a human-readable device name or "" on failure.
//
// [GetCameraName]: https://wiki.libsdl.org/SDL3/SDL_GetCameraName
func GetCameraName(instanceId CameraID) string {
	return sdlGetCameraName(instanceId)
}

// [GetCameraPermissionState] queries if camera access has been approved by the user.
//
// [GetCameraPermissionState]: https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState
func GetCameraPermissionState(camera *Camera) int32 {
	return sdlGetCameraPermissionState(camera)
}

// [GetCameraPosition] gets the position of the camera in relation to the system.
//
// [GetCameraPosition]: https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition
func GetCameraPosition(instanceId CameraID) CameraPosition {
	return sdlGetCameraPosition(instanceId)
}

// [GetCameraProperties] gets the properties associated with an opened camera.
//
// [GetCameraProperties]: https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties
func GetCameraProperties(camera *Camera) PropertiesID {
	return sdlGetCameraProperties(camera)
}

// [GetCameras] gets a list of currently connected camera devices.
//
// [GetCameras]: https://wiki.libsdl.org/SDL3/SDL_GetCameras
func GetCameras() []CameraID {
	var count int32
	cameras := sdlGetCameras(&count)
	if cameras == nil {
		return nil
	}
	defer Free(unsafe.Pointer(cameras))
	return mem.Copy(cameras, count)
}

// [GetCameraSupportedFormats] gets the list of native formats/sizes a camera supports.
//
// [GetCameraSupportedFormats]: https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats
func GetCameraSupportedFormats(instanceId CameraID) []*CameraSpec {
	var count int32
	formats := sdlGetCameraSupportedFormats(instanceId, &count)
	if formats == nil {
		return nil
	}
	defer Free(unsafe.Pointer(formats))
	return mem.DeepCopy(formats, count)
}

// [GetCurrentCameraDriver] returns the name of the current camera driver or "" if no driver has been initialized.
//
// [GetCurrentCameraDriver]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver
func GetCurrentCameraDriver() string {
	return sdlGetCurrentCameraDriver()
}

// [GetNumCameraDrivers] gets the number of built-in camera drivers.
//
// [GetNumCameraDrivers]: https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers
func GetNumCameraDrivers() int32 {
	return sdlGetNumCameraDrivers()
}

// [OpenCamera] opens a video recording device (a "camera").
//
// [OpenCamera]: https://wiki.libsdl.org/SDL3/SDL_OpenCamera
func OpenCamera(instanceId CameraID, spec *CameraSpec) *Camera {
	return sdlOpenCamera(instanceId, spec)
}

// [ReleaseCameraFrame] releases a frame of video acquired from a camera.
//
// [ReleaseCameraFrame]: https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame
func ReleaseCameraFrame(camera *Camera, frame *Surface) {
	sdlReleaseCameraFrame(camera, frame)
}
