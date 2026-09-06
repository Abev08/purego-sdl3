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

// [GetAudioPlaybackDevices] gets a list of currently-connected audio playback devices.
//
// [GetAudioPlaybackDevices]: https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices
// func GetAudioPlaybackDevices(count *int32) *AudioDeviceID {
//	return sdlGetAudioPlaybackDevices(count)
// }

// [GetAudioRecordingDevices] gets a list of currently-connected audio recording devices.
//
// [GetAudioRecordingDevices]: https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices
// func GetAudioRecordingDevices(count *int32) *AudioDeviceID {
//	return sdlGetAudioRecordingDevices(count)
// }

// [GetAudioDeviceName] gets the human-readable name of a specific audio device.
//
// [GetAudioDeviceName]: https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName
// func GetAudioDeviceName(devid AudioDeviceID) string {
//	return sdlGetAudioDeviceName(devid)
// }

// [GetAudioDeviceFormat] gets the current audio format of a specific audio device.
//
// [GetAudioDeviceFormat]: https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat
// func GetAudioDeviceFormat(devid AudioDeviceID, spec *AudioSpec, sample_frames *int32) bool {
//	return sdlGetAudioDeviceFormat(devid, spec, sample_frames)
// }

// [GetAudioDeviceChannelMap] gets the current channel map of an audio device.
//
// [GetAudioDeviceChannelMap]: https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap
// func GetAudioDeviceChannelMap(devid AudioDeviceID, count *int32) *int32 {
//	return sdlGetAudioDeviceChannelMap(devid, count)
// }

// [OpenAudioDevice] opens a specific audio device.
//
// [OpenAudioDevice]: https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice
// func OpenAudioDevice(devid AudioDeviceID, spec *AudioSpec) AudioDeviceID {
//	return sdlOpenAudioDevice(devid, spec)
// }

// [IsAudioDevicePhysical] determines if an audio device is physical (instead of logical).
//
// [IsAudioDevicePhysical]: https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical
// func IsAudioDevicePhysical(devid AudioDeviceID) bool {
//	return sdlIsAudioDevicePhysical(devid)
// }

// [IsAudioDevicePlayback] determines if an audio device is a playback device (instead of recording).
//
// [IsAudioDevicePlayback]: https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback
// func IsAudioDevicePlayback(devid AudioDeviceID) bool {
//	return sdlIsAudioDevicePlayback(devid)
// }

// [PauseAudioDevice] use this function to pause audio playback on a specified device.
//
// [PauseAudioDevice]: https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice
// func PauseAudioDevice(dev AudioDeviceID) bool {
//	return sdlPauseAudioDevice(dev)
// }

// [ResumeAudioDevice] use this function to unpause audio playback on a specified device.
//
// [ResumeAudioDevice]: https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice
// func ResumeAudioDevice(dev AudioDeviceID) bool {
//	return sdlResumeAudioDevice(dev)
// }

// [AudioDevicePaused] use this function to query if an audio device is paused.
//
// [AudioDevicePaused]: https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused
// func AudioDevicePaused(dev AudioDeviceID) bool {
//	return sdlAudioDevicePaused(dev)
// }

// [GetAudioDeviceGain] gets the gain of an audio device.
//
// [GetAudioDeviceGain]: https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain
// func GetAudioDeviceGain(devid AudioDeviceID) float32 {
//	return sdlGetAudioDeviceGain(devid)
// }

// [SetAudioDeviceGain] changes the gain of an audio device.
//
// [SetAudioDeviceGain]: https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain
// func SetAudioDeviceGain(devid AudioDeviceID, gain float32) bool {
//	return sdlSetAudioDeviceGain(devid, gain)
// }

// [CloseAudioDevice] closes a previously-opened audio device.
//
// [CloseAudioDevice]: https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice
// func CloseAudioDevice(devid AudioDeviceID)  {
//	sdlCloseAudioDevice(devid)
// }

// [BindAudioStreams] binds a list of audio streams to an audio device.
//
// [BindAudioStreams]: https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams
// func BindAudioStreams(devid AudioDeviceID, streams **AudioStream, num_streams int32) bool {
//	return sdlBindAudioStreams(devid, streams, num_streams)
// }

// [BindAudioStream] binds a single audio stream to an audio device.
//
// [BindAudioStream]: https://wiki.libsdl.org/SDL3/SDL_BindAudioStream
// func BindAudioStream(devid AudioDeviceID, stream *AudioStream) bool {
//	return sdlBindAudioStream(devid, stream)
// }

// [UnbindAudioStreams] unbinds a list of audio streams from their audio devices.
//
// [UnbindAudioStreams]: https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams
// func UnbindAudioStreams(streams **AudioStream, num_streams int32)  {
//	sdlUnbindAudioStreams(streams, num_streams)
// }

// [UnbindAudioStream] unbinds a single audio stream from its audio device.
//
// [UnbindAudioStream]: https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream
// func UnbindAudioStream(stream *AudioStream)  {
//	sdlUnbindAudioStream(stream)
// }

// [GetAudioStreamDevice] queries an audio stream for its currently-bound device.
//
// [GetAudioStreamDevice]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice
func GetAudioStreamDevice(stream *AudioStream) AudioDeviceID {
	return sdlGetAudioStreamDevice(stream)
}

// [CreateAudioStream] creates a new audio stream.
//
// [CreateAudioStream]: https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream
// func CreateAudioStream(src_spec *AudioSpec, dst_spec *AudioSpec) *AudioStream {
//	return sdlCreateAudioStream(src_spec, dst_spec)
// }

// [GetAudioStreamProperties] gets the properties associated with an audio stream.
//
// [GetAudioStreamProperties]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties
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

// [GetAudioStreamInputChannelMap] gets the current input channel map of an audio stream.
//
// [GetAudioStreamInputChannelMap]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap
// func GetAudioStreamInputChannelMap(stream *AudioStream, count *int32) *int32 {
//	return sdlGetAudioStreamInputChannelMap(stream, count)
// }

// [GetAudioStreamOutputChannelMap] gets the current output channel map of an audio stream.
//
// [GetAudioStreamOutputChannelMap]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap
// func GetAudioStreamOutputChannelMap(stream *AudioStream, count *int32) *int32 {
//	return sdlGetAudioStreamOutputChannelMap(stream, count)
// }

// [SetAudioStreamInputChannelMap] sets the current input channel map of an audio stream.
//
// [SetAudioStreamInputChannelMap]: https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap
// func SetAudioStreamInputChannelMap(stream *AudioStream, chmap *int32, count int32) bool {
//	return sdlSetAudioStreamInputChannelMap(stream, chmap, count)
// }

// [SetAudioStreamOutputChannelMap] sets the current output channel map of an audio stream.
//
// [SetAudioStreamOutputChannelMap]: https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap
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

// [GetAudioStreamData] gets converted/resampled data from the stream.
//
// [GetAudioStreamData]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData
// func GetAudioStreamData(stream *AudioStream, buf unsafe.Pointer, len int32) int32 {
//	return sdlGetAudioStreamData(stream, buf, len)
// }

// [GetAudioStreamAvailable] gets the number of converted/resampled bytes available.
//
// [GetAudioStreamAvailable]: https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable
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

// [LockAudioStream] locks an audio stream for serialized access.
//
// [LockAudioStream]: https://wiki.libsdl.org/SDL3/SDL_LockAudioStream
// func LockAudioStream(stream *AudioStream) bool {
//	return sdlLockAudioStream(stream)
// }

// [UnlockAudioStream] unlocks an audio stream for serialized access.
//
// [UnlockAudioStream]: https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream
// func UnlockAudioStream(stream *AudioStream) bool {
//	return sdlUnlockAudioStream(stream)
// }

// [SetAudioStreamGetCallback] sets a callback that runs when data is requested from an audio stream.
//
// [SetAudioStreamGetCallback]: https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback
// func SetAudioStreamGetCallback(stream *AudioStream, callback AudioStreamCallback, userdata unsafe.Pointer) bool {
//	return sdlSetAudioStreamGetCallback(stream, callback, userdata)
// }

// [SetAudioStreamPutCallback] sets a callback that runs when data is added to an audio stream.
//
// [SetAudioStreamPutCallback]: https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback
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

// [AudioPostmixCallback] defines a callback that fires when data is about to be fed to an audio device.
//
// [AudioPostmixCallback]: https://wiki.libsdl.org/SDL3/SDL_AudioPostmixCallback
type AudioPostmixCallback uintptr

// [SetAudioPostmixCallback] sets a callback that fires when data is about to be fed to an audio device.
//
// [SetAudioPostmixCallback]: https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback
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

// [MixAudio] mixs audio data in a specified format.
//
// [MixAudio]: https://wiki.libsdl.org/SDL3/SDL_MixAudio
// func MixAudio(dst *uint8, src *uint8, format AudioFormat, len uint32, volume float32) bool {
//	return sdlMixAudio(dst, src, format, len, volume)
// }

// [ConvertAudioSamples] converts some audio data of one format to another format.
//
// [ConvertAudioSamples]: https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples
// func ConvertAudioSamples(src_spec *AudioSpec, src_data *uint8, src_len int32, dst_spec *AudioSpec, dst_data **uint8, dst_len *int32) bool {
//	return sdlConvertAudioSamples(src_spec, src_data, src_len, dst_spec, dst_data, dst_len)
// }

// [GetAudioFormatName] gets the human readable name of an audio format.
//
// [GetAudioFormatName]: https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName
// func GetAudioFormatName(format AudioFormat) string {
//	return sdlGetAudioFormatName(format)
// }

// [GetSilenceValueForFormat] gets the appropriate memset value for silencing an audio format.
//
// [GetSilenceValueForFormat]: https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat
// func GetSilenceValueForFormat(format AudioFormat) int32 {
//	return sdlGetSilenceValueForFormat(format)
// }
