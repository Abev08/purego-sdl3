package sdl

import "unsafe"

const (
	PropIOStreamWindowsHandlePointer   = "SDL.iostream.windows.handle"
	PropIOStreamStdioFilePointer       = "SDL.iostream.stdio.file"
	PropIOStreamFileDescriptorNumber   = "SDL.iostream.file_descriptor"
	PropIOStreamAndroidAassetPointer   = "SDL.iostream.android.aasset"
	PropIOStreamMemoryPointer          = "SDL.iostream.memory.base"
	PropIOStreamMemorySizeNumber       = "SDL.iostream.memory.size"
	PropIOStreamMemoryFreeFuncPointer  = "SDL.iostream.memory.free"
	PropIOStreamDynamicMemoryPointer   = "SDL.iostream.dynamic.memory"
	PropIOStreamDynamicChunksizeNumber = "SDL.iostream.dynamic.chunksize"
)

// [IOStatus] defines the [IOStream] status, set by a read or write operation.
//
// [IOStatus]: https://wiki.libsdl.org/SDL3/SDL_IOStatus
type IOStatus uint32

const (
	IOStatusReady     IOStatus = iota // Everything is ready (no errors and not EOF).
	IOStatusError                     // Read or write I/O error.
	IOStatusEof                       // End of file.
	IOStatusNotReady                  // Non blocking I/O, not ready.
	IOStatusReadOnly                  // Tried to write a read-only buffer.
	IOStatusWriteOnly                 // Tried to read a write-only buffer.
)

// [IOWhence] defines the possible `whence` values for [IOStream] seeking.
//
// [IOWhence]: https://wiki.libsdl.org/SDL3/SDL_IOWhence
type IOWhence uint32

const (
	IOSeekSet IOWhence = iota // Seek from the beginning of data.
	IOSeekCur                 // Seek relative to current read point.
	IOSeekEnd                 // Seek relative to the end of data.
)

// [IOStreamInterface] defines the function pointers that drive an [IOStream].
//
// [IOStreamInterface]: https://wiki.libsdl.org/SDL3/SDL_IOStreamInterface
type IOStreamInterface struct {
	Version uint32
	Size    *func(userdata uintptr) int64                                // Returns the number of bytes in this [IOStream].
	Seek    *func(userdata uintptr, offset int64, whence IOWhence) int64 // Seeks to `offset` relative to `whence`, one of stdio's whence values: [IOSeekSet], [IOSeekCur], [IOSeekEnd].
	Read    *func(userdata, ptr, size uintptr, status *IOStatus) uintptr // Reads up to `size` bytes from the data stream to the area pointed.
	Write   *func(userdata, ptr, size uintptr, status *IOStatus) uintptr // Writes exactly `size` bytes from the area pointed at by `ptr`.
	Flush   *func(userdata uintptr, status *IOStatus) bool               // If the stream is buffering, make sure the data is written out.
	Close   *func(userdata uintptr) bool                                 // Closes and frees any allocated resources.
}

// [IOStream] specifies the read/write operation structures.
//
// [IOStream]: https://wiki.libsdl.org/SDL3/SDL_IOStream
type IOStream struct{}

// [IOFromFile] returns an [IOStream] for the named file. The mode can be "r" for read-only.
//
// [IOFromFile]: https://wiki.libsdl.org/SDL3/SDL_IOFromFile
func IOFromFile(file string, mode string) *IOStream {
	return sdlIOFromFile(file, mode)
}

// [IOFromMem] use this function to prepare a read-write memory buffer for use with [IOStream].
//
// [IOFromMem]: https://wiki.libsdl.org/SDL3/SDL_IOFromMem
func IOFromMem(mem []byte) *IOStream {
	return sdlIOFromMem(mem, len(mem))
}

// [IOFromConstMem] returns a read-only memory buffer for use with [IOStream] or nil on failure.
//
// [IOFromConstMem]: https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem
func IOFromConstMem(mem []byte) *IOStream {
	return sdlIOFromConstMem(mem, len(mem))
}

// [IOFromDynamicMem] use this function to create an [IOStream] that is backed by dynamically.
//
// [IOFromDynamicMem]: https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem
// func IOFromDynamicMem() *IOStream {
//	return sdlIOFromDynamicMem()
// }

// [OpenIO] creates a custom [IOStream].
//
// [OpenIO]: https://wiki.libsdl.org/SDL3/SDL_OpenIO
// func OpenIO(iface *IOStreamInterface, userdata unsafe.Pointer) *IOStream {
//	return sdlOpenIO(iface, userdata)
// }

// [CloseIO] closes and free an allocated [IOStream] structure.
//
// [CloseIO]: https://wiki.libsdl.org/SDL3/SDL_CloseIO
func CloseIO(context *IOStream) bool {
	return sdlCloseIO(context)
}

// [GetIOProperties] gets the properties associated with an [IOStream].
//
// [GetIOProperties]: https://wiki.libsdl.org/SDL3/SDL_GetIOProperties
// func GetIOProperties(context *IOStream) PropertiesID {
//	return sdlGetIOProperties(context)
// }

// [GetIOStatus] queries the stream status of an [IOStream].
//
// [GetIOStatus]: https://wiki.libsdl.org/SDL3/SDL_GetIOStatus
// func GetIOStatus(context *IOStream) IOStatus {
//	return sdlGetIOStatus(context)
// }

// [GetIOSize] use this function to get the size of the data stream in an [IOStream].
//
// [GetIOSize]: https://wiki.libsdl.org/SDL3/SDL_GetIOSize
// func GetIOSize(context *IOStream) int64 {
//	return sdlGetIOSize(context)
// }

// [SeekIO] seeks within an [IOStream] data stream.
//
// [SeekIO]: https://wiki.libsdl.org/SDL3/SDL_SeekIO
// func SeekIO(context *IOStream, offset int64, whence IOWhence) int64 {
//	return sdlSeekIO(context, offset, whence)
// }

// [TellIO] determines the current read/write offset in an [IOStream] data stream.
//
// [TellIO]: https://wiki.libsdl.org/SDL3/SDL_TellIO
// func TellIO(context *IOStream) int64 {
//	return sdlTellIO(context)
// }

// [ReadIO] reads from a data source.
//
// [ReadIO]: https://wiki.libsdl.org/SDL3/SDL_ReadIO
// func ReadIO(context *IOStream, ptr unsafe.Pointer, size uint64) uint64 {
//	return sdlReadIO(context, ptr, size)
// }

// [WriteIO] writes to an [IOStream] data stream.
//
// [WriteIO]: https://wiki.libsdl.org/SDL3/SDL_WriteIO
// func WriteIO(context *IOStream, ptr unsafe.Pointer, size uint64) uint64 {
//	return sdlWriteIO(context, ptr, size)
// }

// [IOprintf] prints to an [IOStream] data stream.
//
// [IOprintf]: https://wiki.libsdl.org/SDL3/SDL_IOprintf
// func IOprintf(context *IOStream, fmt string) uint64 {
//	return sdlIOprintf(context, fmt)
// }

// [IOvprintf] prints to an [IOStream] data stream.
//
// [IOvprintf]: https://wiki.libsdl.org/SDL3/SDL_IOvprintf
// func IOvprintf(context *IOStream, fmt string, ap va_list) uint64 {
//	return sdlIOvprintf(context, fmt, ap)
// }

// [FlushIO] flushes any buffered data in the stream.
//
// [FlushIO]: https://wiki.libsdl.org/SDL3/SDL_FlushIO
// func FlushIO(context *IOStream) bool {
//	return sdlFlushIO(context)
// }

// [LoadFileIO] loads all the data from an SDL data stream.
//
// [LoadFileIO]: https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO
// func LoadFile_IO(src *IOStream, datasize *uint64, closeio bool) unsafe.Pointer {
//	return sdlLoadFile_IO(src, datasize, closeio)
// }

// [LoadFile] loads all the data from a file path.
//
// [LoadFile]: https://wiki.libsdl.org/SDL3/SDL_LoadFile
func LoadFile(file string, dataSize *uint64) unsafe.Pointer {
	return sdlLoadFile(file, dataSize)
}

// [SaveFileIO] saves all the data into an SDL data stream.
//
// [SaveFileIO]: https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO
// func SaveFile_IO(src *IOStream, data unsafe.Pointer, datasize uint64, closeio bool) bool {
//	return sdlSaveFile_IO(src, data, datasize, closeio)
// }

// [SaveFile] saves all the data into a file path.
//
// [SaveFile]: https://wiki.libsdl.org/SDL3/SDL_SaveFile
// func SaveFile(file string, data unsafe.Pointer, datasize uint64) bool {
//	return sdlSaveFile(file, data, datasize)
// }

// [ReadU8] reads a byte from an [IOStream].
//
// [ReadU8]: https://wiki.libsdl.org/SDL3/SDL_ReadU8
// func ReadU8(src *IOStream, value *uint8) bool {
//	return sdlReadU8(src, value)
// }

// [ReadS8] reads a signed byte from an [IOStream].
//
// [ReadS8]: https://wiki.libsdl.org/SDL3/SDL_ReadS8
// func ReadS8(src *IOStream, value *int8) bool {
//	return sdlReadS8(src, value)
// }

// [ReadU16LE] reads 16 bits of little-endian data from an [IOStream].
//
// [ReadU16LE]: https://wiki.libsdl.org/SDL3/SDL_ReadU16LE
// func ReadU16LE(src *IOStream, value *uint16) bool {
//	return sdlReadU16LE(src, value)
// }

// [ReadS16LE] reads 16 bits of little-endian data from an [IOStream].
//
// [ReadS16LE]: https://wiki.libsdl.org/SDL3/SDL_ReadS16LE
// func ReadS16LE(src *IOStream, value *int16) bool {
//	return sdlReadS16LE(src, value)
// }

// [ReadU16BE] reads 16 bits of big-endian data from an [IOStream].
//
// [ReadU16BE]: https://wiki.libsdl.org/SDL3/SDL_ReadU16BE
// func ReadU16BE(src *IOStream, value *uint16) bool {
//	return sdlReadU16BE(src, value)
// }

// [ReadS16BE] reads 16 bits of big-endian data from an [IOStream].
//
// [ReadS16BE]: https://wiki.libsdl.org/SDL3/SDL_ReadS16BE
// func ReadS16BE(src *IOStream, value *int16) bool {
//	return sdlReadS16BE(src, value)
// }

// [ReadU32LE] reads 32 bits of little-endian data from an [IOStream].
//
// [ReadU32LE]: https://wiki.libsdl.org/SDL3/SDL_ReadU32LE
// func ReadU32LE(src *IOStream, value *uint32) bool {
//	return sdlReadU32LE(src, value)
// }

// [ReadS32LE] reads 32 bits of little-endian data from an [IOStream].
//
// [ReadS32LE]: https://wiki.libsdl.org/SDL3/SDL_ReadS32LE
// func ReadS32LE(src *IOStream, value *int32) bool {
//	return sdlReadS32LE(src, value)
// }

// [ReadU32BE] reads 32 bits of big-endian data from an [IOStream].
//
// [ReadU32BE]: https://wiki.libsdl.org/SDL3/SDL_ReadU32BE
// func ReadU32BE(src *IOStream, value *uint32) bool {
//	return sdlReadU32BE(src, value)
// }

// [ReadS32BE] reads 32 bits of big-endian data from an [IOStream].
//
// [ReadS32BE]: https://wiki.libsdl.org/SDL3/SDL_ReadS32BE
// func ReadS32BE(src *IOStream, value *int32) bool {
//	return sdlReadS32BE(src, value)
// }

// [ReadU64LE] reads 64 bits of little-endian data from an [IOStream].
//
// [ReadU64LE]: https://wiki.libsdl.org/SDL3/SDL_ReadU64LE
// func ReadU64LE(src *IOStream, value *uint64) bool {
//	return sdlReadU64LE(src, value)
// }

// [ReadS64LE] reads 64 bits of little-endian data from an [IOStream].
//
// [ReadS64LE]: https://wiki.libsdl.org/SDL3/SDL_ReadS64LE
// func ReadS64LE(src *IOStream, value *int64) bool {
//	return sdlReadS64LE(src, value)
// }

// [ReadU64BE] reads 64 bits of big-endian data from an [IOStream].
//
// [ReadU64BE]: https://wiki.libsdl.org/SDL3/SDL_ReadU64BE
// func ReadU64BE(src *IOStream, value *uint64) bool {
//	return sdlReadU64BE(src, value)
// }

// [ReadS64BE] reads 64 bits of big-endian data from an [IOStream].
//
// [ReadS64BE]: https://wiki.libsdl.org/SDL3/SDL_ReadS64BE
// func ReadS64BE(src *IOStream, value *int64) bool {
//	return sdlReadS64BE(src, value)
// }

// [WriteU8] writes a byte to an [IOStream].
//
// [WriteU8]: https://wiki.libsdl.org/SDL3/SDL_WriteU8
// func WriteU8(dst *IOStream, value uint8) bool {
//	return sdlWriteU8(dst, value)
// }

// [WriteS8] writes a signed byte to an [IOStream].
//
// [WriteS8]: https://wiki.libsdl.org/SDL3/SDL_WriteS8
// func WriteS8(dst *IOStream, value int8) bool {
//	return sdlWriteS8(dst, value)
// }

// [WriteU16LE] writes 16 bits in native format to an [IOStream] as little-endian data.
//
// [WriteU16LE]: https://wiki.libsdl.org/SDL3/SDL_WriteU16LE
// func WriteU16LE(dst *IOStream, value uint16) bool {
//	return sdlWriteU16LE(dst, value)
// }

// [WriteS16LE] writes 16 bits in native format to an [IOStream] as little-endian data.
//
// [WriteS16LE]: https://wiki.libsdl.org/SDL3/SDL_WriteS16LE
// func WriteS16LE(dst *IOStream, value int16) bool {
//	return sdlWriteS16LE(dst, value)
// }

// [WriteU16BE] writes 16 bits in native format to an [IOStream] as big-endian data.
//
// [WriteU16BE]: https://wiki.libsdl.org/SDL3/SDL_WriteU16BE
// func WriteU16BE(dst *IOStream, value uint16) bool {
//	return sdlWriteU16BE(dst, value)
// }

// [WriteS16BE] writes 16 bits in native format to an [IOStream] as big-endian data.
//
// [WriteS16BE]: https://wiki.libsdl.org/SDL3/SDL_WriteS16BE
// func WriteS16BE(dst *IOStream, value int16) bool {
//	return sdlWriteS16BE(dst, value)
// }

// [WriteU32LE] writes 32 bits in native format to an [IOStream] as little-endian data.
//
// [WriteU32LE]: https://wiki.libsdl.org/SDL3/SDL_WriteU32LE
// func WriteU32LE(dst *IOStream, value uint32) bool {
//	return sdlWriteU32LE(dst, value)
// }

// [WriteS32LE] writes 32 bits in native format to an [IOStream] as little-endian data.
//
// [WriteS32LE]: https://wiki.libsdl.org/SDL3/SDL_WriteS32LE
// func WriteS32LE(dst *IOStream, value int32) bool {
//	return sdlWriteS32LE(dst, value)
// }

// [WriteU32BE] writes 32 bits in native format to an [IOStream] as big-endian data.
//
// [WriteU32BE]: https://wiki.libsdl.org/SDL3/SDL_WriteU32BE
// func WriteU32BE(dst *IOStream, value uint32) bool {
//	return sdlWriteU32BE(dst, value)
// }

// [WriteS32BE] writes 32 bits in native format to an [IOStream] as big-endian data.
//
// [WriteS32BE]: https://wiki.libsdl.org/SDL3/SDL_WriteS32BE
// func WriteS32BE(dst *IOStream, value int32) bool {
//	return sdlWriteS32BE(dst, value)
// }

// [WriteU64LE] writes 64 bits in native format to an [IOStream] as little-endian data.
//
// [WriteU64LE]: https://wiki.libsdl.org/SDL3/SDL_WriteU64LE
// func WriteU64LE(dst *IOStream, value uint64) bool {
//	return sdlWriteU64LE(dst, value)
// }

// [WriteS64LE] writes 64 bits in native format to an [IOStream] as little-endian data.
//
// [WriteS64LE]: https://wiki.libsdl.org/SDL3/SDL_WriteS64LE
// func WriteS64LE(dst *IOStream, value int64) bool {
//	return sdlWriteS64LE(dst, value)
// }

// [WriteU64BE] writes 64 bits in native format to an [IOStream] as big-endian data.
//
// [WriteU64BE]: https://wiki.libsdl.org/SDL3/SDL_WriteU64BE
// func WriteU64BE(dst *IOStream, value uint64) bool {
//	return sdlWriteU64BE(dst, value)
// }

// [WriteS64BE] writes 64 bits in native format to an [IOStream] as big-endian data.
//
// [WriteS64BE]: https://wiki.libsdl.org/SDL3/SDL_WriteS64BE
// func WriteS64BE(dst *IOStream, value int64) bool {
//	return sdlWriteS64BE(dst, value)
// }
