package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// [AudioFormat] defines the audio format.
//
// [AudioFormat]: https://wiki.libsdl.org/SDL3/SDL_AudioFormat
type AudioFormat uint32

const (
	AudioUnknown AudioFormat = 0x0000 // Unspecified audio format.
	AudioU8      AudioFormat = 0x0008 // Unsigned 8-bit samples.
	AudioS8      AudioFormat = 0x8008 // Signed 8-bit samples.
	AudioS16Le   AudioFormat = 0x8010 // Signed 16-bit samples.
	AudioS16Be   AudioFormat = 0x9010 // Signed 16-bit samples, big-endian byte order.
	AudioS32Le   AudioFormat = 0x8020 // 32-bit integer samples.
	AudioS32Be   AudioFormat = 0x9020 // 32-bit integer samples, big-endian byte order.
	AudioF32Le   AudioFormat = 0x8120 // 32-bit floating point samples.
	AudioF32Be   AudioFormat = 0x9120 // 32-bit floating, big-endian byte order.
	AudioS16     AudioFormat = AudioS16Le
	AudioS32     AudioFormat = AudioS32Le
	AudioF32     AudioFormat = AudioF32Le
)

// [AudioDeviceID] defines the SDL Audio Device instance IDs.
//
// [AudioDeviceID]: https://wiki.libsdl.org/SDL3/SDL_AudioDeviceID
type AudioDeviceID uint32

const (
	AudioDeviceDefaultPlayback  AudioDeviceID = 0xFFFFFFFF
	AudioDeviceDefaultRecording AudioDeviceID = 0xFFFFFFFE
)

// [AudioSpec] defines the format specifier for audio data.
//
// [AudioSpec]: https://wiki.libsdl.org/SDL3/SDL_AudioSpec
type AudioSpec struct {
	Format   AudioFormat // Audio data format.
	Channels int32       // Number of channels: 1 mono, 2 stereo, etc.
	Freq     int32       // Sample rate: sample frames per second.
}

// [AudioStream] is a structure specifying the opaque handle that represents an audio streams.
//
// [AudioStream]: https://wiki.libsdl.org/SDL3/SDL_AudioStream
type AudioStream struct{}

// [AudioStreamCallback] is a callback that fires when data passes through an [AudioStream].
//
// [AudioStreamCallback]: https://wiki.libsdl.org/SDL3/SDL_AudioStreamCallback
type AudioStreamCallback uintptr

func NewAudioStreamCallback(callback func(userdata unsafe.Pointer, stream *AudioStream, additionalAmount, totalAmount int32)) AudioStreamCallback {
	cb := purego.NewCallback(func(userdata unsafe.Pointer, stream *AudioStream, additionalAmount, totalAmount int32) uintptr {
		callback(userdata, stream, additionalAmount, totalAmount)
		return 0
	})

	return AudioStreamCallback(cb)
}

// func AudioDevicePaused(dev AudioDeviceID) bool {
//	return sdlAudioDevicePaused(dev)
// }

// [AudioStreamDevicePaused] queries if an audio device associated with a stream is paused.
//
// [AudioStreamDevicePaused]: https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused
func AudioStreamDevicePaused(stream *AudioStream) bool {
	ret, _, _ := purego.SyscallN(sdlAudioStreamDevicePaused, uintptr(unsafe.Pointer(stream)))
	return byte(ret) != 0
}

// func BindAudioStream(devid AudioDeviceID, stream *AudioStream) bool {
//	return sdlBindAudioStream(devid, stream)
// }

// func BindAudioStreams(devid AudioDeviceID, streams **AudioStream, num_streams int32) bool {
//	return sdlBindAudioStreams(devid, streams, num_streams)
// }

// [ClearAudioStream] clears any pending data in the stream.
//
// [ClearAudioStream]: https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream
func ClearAudioStream(stream *AudioStream) bool {
	ret, _, _ := purego.SyscallN(sdlClearAudioStream, uintptr(unsafe.Pointer(stream)))
	return byte(ret) != 0
}

// func CloseAudioDevice(devid AudioDeviceID)  {
//	sdlCloseAudioDevice(devid)
// }

// func ConvertAudioSamples(src_spec *AudioSpec, src_data *uint8, src_len int32, dst_spec *AudioSpec, dst_data **uint8, dst_len *int32) bool {
//	return sdlConvertAudioSamples(src_spec, src_data, src_len, dst_spec, dst_data, dst_len)
// }

// func CreateAudioStream(src_spec *AudioSpec, dst_spec *AudioSpec) *AudioStream {
//	return sdlCreateAudioStream(src_spec, dst_spec)
// }

// [DestroyAudioStream] frees an audio stream.
//
// [DestroyAudioStream]: https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream
func DestroyAudioStream(stream *AudioStream) {
	sdlDestroyAudioStream(stream)
}

// [FlushAudioStream] tells the stream that you're done sending data,
// and anything being buffered should be converted/resampled and made available immediately.
//
// [FlushAudioStream]: https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream
func FlushAudioStream(stream *AudioStream) bool {
	ret, _, _ := purego.SyscallN(sdlFlushAudioStream, uintptr(unsafe.Pointer(stream)))
	return byte(ret) != 0
}

// func GetAudioDeviceChannelMap(devid AudioDeviceID, count *int32) *int32 {
//	return sdlGetAudioDeviceChannelMap(devid, count)
// }

// func GetAudioDeviceFormat(devid AudioDeviceID, spec *AudioSpec, sample_frames *int32) bool {
//	return sdlGetAudioDeviceFormat(devid, spec, sample_frames)
// }

// func GetAudioDeviceGain(devid AudioDeviceID) float32 {
//	return sdlGetAudioDeviceGain(devid)
// }

// func GetAudioDeviceName(devid AudioDeviceID) string {
//	return sdlGetAudioDeviceName(devid)
// }

// [GetAudioDriver] gets the name of a built in audio driver.
//
// [GetAudioDriver]: https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver
func GetAudioDriver(index int32) string {
	return sdlGetAudioDriver(index)
}

// func GetAudioFormatName(format AudioFormat) string {
//	return sdlGetAudioFormatName(format)
// }

// func GetAudioPlaybackDevices(count *int32) *AudioDeviceID {
//	return sdlGetAudioPlaybackDevices(count)
// }

// func GetAudioRecordingDevices(count *int32) *AudioDeviceID {
//	return sdlGetAudioRecordingDevices(count)
// }

// func GetAudioStreamAvailable(stream *AudioStream) int32 {
//	return sdlGetAudioStreamAvailable(stream)
// }

// func GetAudioStreamData(stream *AudioStream, buf unsafe.Pointer, len int32) int32 {
//	return sdlGetAudioStreamData(stream, buf, len)
// }

// func GetAudioStreamDevice(stream *AudioStream) AudioDeviceID {
//	return sdlGetAudioStreamDevice(stream)
// }

// func GetAudioStreamFormat(stream *AudioStream, src_spec *AudioSpec, dst_spec *AudioSpec) bool {
//	return sdlGetAudioStreamFormat(stream, src_spec, dst_spec)
// }

// func GetAudioStreamFrequencyRatio(stream *AudioStream) float32 {
//	return sdlGetAudioStreamFrequencyRatio(stream)
// }

// func GetAudioStreamGain(stream *AudioStream) float32 {
//	return sdlGetAudioStreamGain(stream)
// }

// func GetAudioStreamInputChannelMap(stream *AudioStream, count *int32) *int32 {
//	return sdlGetAudioStreamInputChannelMap(stream, count)
// }

// func GetAudioStreamOutputChannelMap(stream *AudioStream, count *int32) *int32 {
//	return sdlGetAudioStreamOutputChannelMap(stream, count)
// }

// func GetAudioStreamProperties(stream *AudioStream) PropertiesID {
//	return sdlGetAudioStreamProperties(stream)
// }

// [GetAudioStreamQueued] gets the number of bytes currently queued.
//
// [GetAudioStreamQueued]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued
func GetAudioStreamQueued(stream *AudioStream) int32 {
	ret, _, _ := purego.SyscallN(sdlGetAudioStreamQueued, uintptr(unsafe.Pointer(stream)))
	return int32(ret)
}

// [GetCurrentAudioDriver] gets the name of the current audio driver.
//
// [GetCurrentAudioDriver]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver
func GetCurrentAudioDriver() string {
	return sdlGetCurrentAudioDriver()
}

// [GetNumAudioDrivers] gets the number of built-in audio drivers.
//
// [GetNumAudioDrivers]: https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers
func GetNumAudioDrivers() int32 {
	return sdlGetNumAudioDrivers()
}

// func GetSilenceValueForFormat(format AudioFormat) int32 {
//	return sdlGetSilenceValueForFormat(format)
// }

// func IsAudioDevicePhysical(devid AudioDeviceID) bool {
//	return sdlIsAudioDevicePhysical(devid)
// }

// func IsAudioDevicePlayback(devid AudioDeviceID) bool {
//	return sdlIsAudioDevicePlayback(devid)
// }

// func LoadWAV(path string, spec *AudioSpec, audio_buf **uint8, audio_len *uint32) bool {
//	return sdlLoadWAV(path, spec, audio_buf, audio_len)
// }

// [LoadWAVIO] loads the audio data of a WAVE file into memory and returns true on success.
// The data returned in audioBuf should be disposed with [Free] when it is no longer needed.
//
// [LoadWAVIO]: https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO
func LoadWAVIO(src *IOStream, closeio bool, spec *AudioSpec, audioBuf **uint8, audioLen *uint32) bool {
	return sdlLoadWAVIO(src, closeio, spec, audioBuf, audioLen)
}

// func LockAudioStream(stream *AudioStream) bool {
//	return sdlLockAudioStream(stream)
// }

// func MixAudio(dst *uint8, src *uint8, format AudioFormat, len uint32, volume float32) bool {
//	return sdlMixAudio(dst, src, format, len, volume)
// }

// func OpenAudioDevice(devid AudioDeviceID, spec *AudioSpec) AudioDeviceID {
//	return sdlOpenAudioDevice(devid, spec)
// }

// [OpenAudioDeviceStream] returns an audio stream on success, ready to use, or nil on failure.
// When done with this stream, call [DestroyAudioStream] to free resources and close the device.
//
// [OpenAudioDeviceStream]: https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream
func OpenAudioDeviceStream(devid AudioDeviceID, spec *AudioSpec, callback AudioStreamCallback, userdata unsafe.Pointer) *AudioStream {
	return sdlOpenAudioDeviceStream(devid, spec, callback, userdata)
}

// func PauseAudioDevice(dev AudioDeviceID) bool {
//	return sdlPauseAudioDevice(dev)
// }

// [PauseAudioStreamDevice] pauses audio playback on the audio device associated with an audio stream.
//
// [PauseAudioStreamDevice]: https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice
func PauseAudioStreamDevice(stream *AudioStream) bool {
	ret, _, _ := purego.SyscallN(sdlPauseAudioStreamDevice, uintptr(unsafe.Pointer(stream)))
	return byte(ret) != 0
}

// [PutAudioStreamData] adds data to the stream.
//
// [PutAudioStreamData]: https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func PutAudioStreamData(stream *AudioStream, buf *uint8, len int32) bool {
	ret, _, _ := purego.SyscallN(sdlPutAudioStreamData, uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(buf)), uintptr(len))
	return byte(ret) != 0
}

// func ResumeAudioDevice(dev AudioDeviceID) bool {
//	return sdlResumeAudioDevice(dev)
// }

// [ResumeAudioStreamDevice] unpauses audio playback on the audio device associated with an audio stream.
//
// [ResumeAudioStreamDevice]: https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice
func ResumeAudioStreamDevice(stream *AudioStream) bool {
	ret, _, _ := purego.SyscallN(sdlResumeAudioStreamDevice, uintptr(unsafe.Pointer(stream)))
	return byte(ret) != 0
}

// func SetAudioDeviceGain(devid AudioDeviceID, gain float32) bool {
//	return sdlSetAudioDeviceGain(devid, gain)
// }

// func SetAudioPostmixCallback(devid AudioDeviceID, callback AudioPostmixCallback, userdata unsafe.Pointer) bool {
//	return sdlSetAudioPostmixCallback(devid, callback, userdata)
// }

// func SetAudioStreamFormat(stream *AudioStream, src_spec *AudioSpec, dst_spec *AudioSpec) bool {
//	return sdlSetAudioStreamFormat(stream, src_spec, dst_spec)
// }

// func SetAudioStreamFrequencyRatio(stream *AudioStream, ratio float32) bool {
//	return sdlSetAudioStreamFrequencyRatio(stream, ratio)
// }

// func SetAudioStreamGain(stream *AudioStream, gain float32) bool {
//	return sdlSetAudioStreamGain(stream, gain)
// }

// func SetAudioStreamGetCallback(stream *AudioStream, callback AudioStreamCallback, userdata unsafe.Pointer) bool {
//	return sdlSetAudioStreamGetCallback(stream, callback, userdata)
// }

// func SetAudioStreamInputChannelMap(stream *AudioStream, chmap *int32, count int32) bool {
//	return sdlSetAudioStreamInputChannelMap(stream, chmap, count)
// }

// func SetAudioStreamOutputChannelMap(stream *AudioStream, chmap *int32, count int32) bool {
//	return sdlSetAudioStreamOutputChannelMap(stream, chmap, count)
// }

// func SetAudioStreamPutCallback(stream *AudioStream, callback AudioStreamCallback, userdata unsafe.Pointer) bool {
//	return sdlSetAudioStreamPutCallback(stream, callback, userdata)
// }

// func UnbindAudioStream(stream *AudioStream)  {
//	sdlUnbindAudioStream(stream)
// }

// func UnbindAudioStreams(streams **AudioStream, num_streams int32)  {
//	sdlUnbindAudioStreams(streams, num_streams)
// }

// func UnlockAudioStream(stream *AudioStream) bool {
//	return sdlUnlockAudioStream(stream)
// }
