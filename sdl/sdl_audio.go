package sdl

import (
	"unsafe"
)

const PropAudiostreamAutoCleanupBoolean = "SDL.audiostream.auto_cleanup"

const (
	AudioMaskBitsize   uint32 = 0xFF    // Mask of bits in an [AudioFormat] that contains the format bit size.
	AudioMaskFloat     uint32 = 1 << 8  // Mask of bits in an [AudioFormat] that contain the floating point flag.
	AudioMaskBigEndian uint32 = 1 << 12 // Mask of bits in an [AudioFormat] that contain the bigendian flag.
	AudioMaskSigned    uint32 = 1 << 15 // Mask of bits in an [AudioFormat] that contain the signed data flag.
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

// [AudioStreamDataCompleteCallback] is a callback that fires for completed [PutAudioStreamDataNoCopy] data.
//
// Available since SDL 3.4.0.
//
// [AudioStreamDataCompleteCallback]: https://wiki.libsdl.org/SDL3/SDL_AudioStreamDataCompleteCallback
type AudioStreamDataCompleteCallback uintptr

// [GetNumAudioDrivers] gets the number of built-in audio drivers.
//
// [GetNumAudioDrivers]: https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers
func GetNumAudioDrivers() int32 {
	return sdlGetNumAudioDrivers()
}

// [GetAudioDriver] gets the name of a built in audio driver.
//
// [GetAudioDriver]: https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver
func GetAudioDriver(index int32) string {
	return sdlGetAudioDriver(index)
}

// [GetCurrentAudioDriver] gets the name of the current audio driver.
//
// [GetCurrentAudioDriver]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver
func GetCurrentAudioDriver() string {
	return sdlGetCurrentAudioDriver()
}

// func GetAudioPlaybackDevices(count *int32) *AudioDeviceID {
//	return sdlGetAudioPlaybackDevices(count)
// }

// func GetAudioRecordingDevices(count *int32) *AudioDeviceID {
//	return sdlGetAudioRecordingDevices(count)
// }

// func GetAudioDeviceName(devid AudioDeviceID) string {
//	return sdlGetAudioDeviceName(devid)
// }

// func GetAudioDeviceFormat(devid AudioDeviceID, spec *AudioSpec, sample_frames *int32) bool {
//	return sdlGetAudioDeviceFormat(devid, spec, sample_frames)
// }

// func GetAudioDeviceChannelMap(devid AudioDeviceID, count *int32) *int32 {
//	return sdlGetAudioDeviceChannelMap(devid, count)
// }

// func OpenAudioDevice(devid AudioDeviceID, spec *AudioSpec) AudioDeviceID {
//	return sdlOpenAudioDevice(devid, spec)
// }

// func IsAudioDevicePhysical(devid AudioDeviceID) bool {
//	return sdlIsAudioDevicePhysical(devid)
// }

// func IsAudioDevicePlayback(devid AudioDeviceID) bool {
//	return sdlIsAudioDevicePlayback(devid)
// }

// func PauseAudioDevice(dev AudioDeviceID) bool {
//	return sdlPauseAudioDevice(dev)
// }

// func ResumeAudioDevice(dev AudioDeviceID) bool {
//	return sdlResumeAudioDevice(dev)
// }

// func AudioDevicePaused(dev AudioDeviceID) bool {
//	return sdlAudioDevicePaused(dev)
// }

// func GetAudioDeviceGain(devid AudioDeviceID) float32 {
//	return sdlGetAudioDeviceGain(devid)
// }

// func SetAudioDeviceGain(devid AudioDeviceID, gain float32) bool {
//	return sdlSetAudioDeviceGain(devid, gain)
// }

// func CloseAudioDevice(devid AudioDeviceID)  {
//	sdlCloseAudioDevice(devid)
// }

// func BindAudioStreams(devid AudioDeviceID, streams **AudioStream, num_streams int32) bool {
//	return sdlBindAudioStreams(devid, streams, num_streams)
// }

// func BindAudioStream(devid AudioDeviceID, stream *AudioStream) bool {
//	return sdlBindAudioStream(devid, stream)
// }

// func UnbindAudioStreams(streams **AudioStream, num_streams int32)  {
//	sdlUnbindAudioStreams(streams, num_streams)
// }

// func UnbindAudioStream(stream *AudioStream)  {
//	sdlUnbindAudioStream(stream)
// }

// [GetAudioStreamDevice] queries an audio stream for its currently-bound device.
//
// [GetAudioStreamDevice]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice
func GetAudioStreamDevice(stream *AudioStream) AudioDeviceID {
	return sdlGetAudioStreamDevice(stream)
}

// func CreateAudioStream(src_spec *AudioSpec, dst_spec *AudioSpec) *AudioStream {
//	return sdlCreateAudioStream(src_spec, dst_spec)
// }

// func GetAudioStreamProperties(stream *AudioStream) PropertiesID {
//	return sdlGetAudioStreamProperties(stream)
// }

// [GetAudioStreamFormat] queries the current format of an audio stream.
//
// [GetAudioStreamFormat]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat
func GetAudioStreamFormat(stream *AudioStream, srcSpec *AudioSpec, dstSpec *AudioSpec) bool {
	return sdlGetAudioStreamFormat(stream, srcSpec, dstSpec)
}

// [SetAudioStreamFormat] changes the input and output formats of an audio stream.
//
// [SetAudioStreamFormat]: https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat
func SetAudioStreamFormat(stream *AudioStream, srcSpec *AudioSpec, dstSpec *AudioSpec) bool {
	return sdlSetAudioStreamFormat(stream, srcSpec, dstSpec)
}

// [GetAudioStreamFrequencyRatio] gets the frequency ratio of an audio stream.
//
// [GetAudioStreamFrequencyRatio]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio
func GetAudioStreamFrequencyRatio(stream *AudioStream) float32 {
	return sdlGetAudioStreamFrequencyRatio(stream)
}

// [SetAudioStreamFrequencyRatio] changes the frequency ratio of an audio stream.
//
// [SetAudioStreamFrequencyRatio]: https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio
func SetAudioStreamFrequencyRatio(stream *AudioStream, ratio float32) bool {
	return sdlSetAudioStreamFrequencyRatio(stream, ratio)
}

// [GetAudioStreamGain] gets the gain of an audio stream.
//
// [GetAudioStreamGain]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain
func GetAudioStreamGain(stream *AudioStream) float32 {
	return sdlGetAudioStreamGain(stream)
}

// [SetAudioStreamGain] changes the gain of an audio stream.
//
// [SetAudioStreamGain]: https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain
func SetAudioStreamGain(stream *AudioStream, gain float32) bool {
	return sdlSetAudioStreamGain(stream, gain)
}

// func GetAudioStreamInputChannelMap(stream *AudioStream, count *int32) *int32 {
//	return sdlGetAudioStreamInputChannelMap(stream, count)
// }

// func GetAudioStreamOutputChannelMap(stream *AudioStream, count *int32) *int32 {
//	return sdlGetAudioStreamOutputChannelMap(stream, count)
// }

// func SetAudioStreamInputChannelMap(stream *AudioStream, chmap *int32, count int32) bool {
//	return sdlSetAudioStreamInputChannelMap(stream, chmap, count)
// }

// func SetAudioStreamOutputChannelMap(stream *AudioStream, chmap *int32, count int32) bool {
//	return sdlSetAudioStreamOutputChannelMap(stream, chmap, count)
// }

// [PutAudioStreamData] adds data to the stream.
//
// [PutAudioStreamData]: https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData
func PutAudioStreamData(stream *AudioStream, buf *uint8, len int32) bool {
	return sdlPutAudioStreamData(stream, buf, len)
}

// [PutAudioStreamDataNoCopy] adds external data to an audio stream without copying it.
//
// Available since SDL 3.4.0.
//
// [PutAudioStreamDataNoCopy]: https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamDataNoCopy
// func PutAudioStreamDataNoCopy(stream *AudioStream, buf *uint8, len int32, callback AudioStreamDataCompleteCallback, userdata uintptr) bool {
// 	ret, _, _ := purego.SyscallN(sdlPutAudioStreamDataNoCopy, uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(buf)), uintptr(len), uintptr(callback), userdata)
// 	return byte(ret) != 0
// }

// [PutAudioStreamPlanarData] adds data to the stream with each channel in a separate array.
//
// Available since SDL 3.4.0.
//
// [PutAudioStreamPlanarData]: https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamPlanarData
// func PutAudioStreamPlanarData(stream *AudioStream, channelBuffers *[][]uint8, numChannels, numSamples int32) bool {
// 	ret, _, _ := purego.SyscallN(sdlPutAudioStreamPlanarData, uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(channelBuffers)), uintptr(numChannels), uintptr(numSamples))
// 	return byte(ret) != 0
// }

// func GetAudioStreamData(stream *AudioStream, buf unsafe.Pointer, len int32) int32 {
//	return sdlGetAudioStreamData(stream, buf, len)
// }

// func GetAudioStreamAvailable(stream *AudioStream) int32 {
//	return sdlGetAudioStreamAvailable(stream)
// }

// [GetAudioStreamQueued] gets the number of bytes currently queued.
//
// [GetAudioStreamQueued]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued
func GetAudioStreamQueued(stream *AudioStream) int32 {
	return sdlGetAudioStreamQueued(stream)
}

// [FlushAudioStream] tells the stream that you're done sending data,
// and anything being buffered should be converted/resampled and made available immediately.
//
// [FlushAudioStream]: https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream
func FlushAudioStream(stream *AudioStream) bool {
	return sdlFlushAudioStream(stream)
}

// [ClearAudioStream] clears any pending data in the stream.
//
// [ClearAudioStream]: https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream
func ClearAudioStream(stream *AudioStream) bool {
	return sdlClearAudioStream(stream)
}

// [PauseAudioStreamDevice] pauses audio playback on the audio device associated with an audio stream.
//
// [PauseAudioStreamDevice]: https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice
func PauseAudioStreamDevice(stream *AudioStream) bool {
	return sdlPauseAudioStreamDevice(stream)
}

// [ResumeAudioStreamDevice] unpauses audio playback on the audio device associated with an audio stream.
//
// [ResumeAudioStreamDevice]: https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice
func ResumeAudioStreamDevice(stream *AudioStream) bool {
	return sdlResumeAudioStreamDevice(stream)
}

// [AudioStreamDevicePaused] queries if an audio device associated with a stream is paused.
//
// [AudioStreamDevicePaused]: https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused
func AudioStreamDevicePaused(stream *AudioStream) bool {
	return sdlAudioStreamDevicePaused(stream)
}

// func LockAudioStream(stream *AudioStream) bool {
//	return sdlLockAudioStream(stream)
// }

// func UnlockAudioStream(stream *AudioStream) bool {
//	return sdlUnlockAudioStream(stream)
// }

// func SetAudioStreamGetCallback(stream *AudioStream, callback AudioStreamCallback, userdata unsafe.Pointer) bool {
//	return sdlSetAudioStreamGetCallback(stream, callback, userdata)
// }

// func SetAudioStreamPutCallback(stream *AudioStream, callback AudioStreamCallback, userdata unsafe.Pointer) bool {
//	return sdlSetAudioStreamPutCallback(stream, callback, userdata)
// }

// [DestroyAudioStream] frees an audio stream.
//
// [DestroyAudioStream]: https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream
func DestroyAudioStream(stream *AudioStream) {
	sdlDestroyAudioStream(stream)
}

// [OpenAudioDeviceStream] returns an audio stream on success, ready to use, or nil on failure.
// When done with this stream, call [DestroyAudioStream] to free resources and close the device.
//
// [OpenAudioDeviceStream]: https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream
func OpenAudioDeviceStream(devid AudioDeviceID, spec *AudioSpec, callback AudioStreamCallback, userdata unsafe.Pointer) *AudioStream {
	return sdlOpenAudioDeviceStream(devid, spec, callback, userdata)
}

// func SetAudioPostmixCallback(devid AudioDeviceID, callback AudioPostmixCallback, userdata unsafe.Pointer) bool {
//	return sdlSetAudioPostmixCallback(devid, callback, userdata)
// }

// [LoadWAVIO] loads the audio data of a WAVE file into memory and returns true on success.
// The data returned in audioBuf should be disposed with [Free] when it is no longer needed.
//
// [LoadWAVIO]: https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO
func LoadWAVIO(src *IOStream, closeio bool, spec *AudioSpec, audioBuf **uint8, audioLen *uint32) bool {
	return sdlLoadWAVIO(src, closeio, spec, audioBuf, audioLen)
}

// [LoadWAV] loads a WAV from a file path.
// The data returned in audioBuf should be disposed with [Free] when it is no longer needed.
//
// [LoadWAV]: https://wiki.libsdl.org/SDL3/SDL_LoadWAV
func LoadWAV(path string, spec *AudioSpec, audioBuf **uint8, audioLen *uint32) bool {
	return sdlLoadWAV(path, spec, audioBuf, audioLen)
}

// func MixAudio(dst *uint8, src *uint8, format AudioFormat, len uint32, volume float32) bool {
//	return sdlMixAudio(dst, src, format, len, volume)
// }

// func ConvertAudioSamples(src_spec *AudioSpec, src_data *uint8, src_len int32, dst_spec *AudioSpec, dst_data **uint8, dst_len *int32) bool {
//	return sdlConvertAudioSamples(src_spec, src_data, src_len, dst_spec, dst_data, dst_len)
// }

// func GetAudioFormatName(format AudioFormat) string {
//	return sdlGetAudioFormatName(format)
// }

// func GetSilenceValueForFormat(format AudioFormat) int32 {
//	return sdlGetSilenceValueForFormat(format)
// }
