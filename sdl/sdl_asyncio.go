package sdl

// [AsyncIO] defines the asynchronous I/O operation structure.
//
// [AsyncIO]: https://wiki.libsdl.org/SDL3/SDL_AsyncIO
type AsyncIO struct{}

// [AsyncIOTaskType] defines the types of asynchronous I/O tasks.
//
// [AsyncIOTaskType]: https://wiki.libsdl.org/SDL3/SDL_AsyncIOTaskType
type AsyncIOTaskType uint32

const (
	AsyncIOTaskRead  AsyncIOTaskType = iota // A read operation.
	AsyncIOTaskWrite                        // A write operation.
	AsyncIOTaskClose                        // A close operation.
)

// [AsyncIOResult] defines the possible outcomes of an asynchronous I/O task.
//
// [AsyncIOResult]: https://wiki.libsdl.org/SDL3/SDL_AsyncIOResult
type AsyncIOResult uint32

const (
	AsyncIOComplete AsyncIOResult = iota // Request was completed without error.
	AsyncIOFailure                       // Request failed for some reason, check [GetError]!
	AsyncIOCanceled                      // Request was canceled before completing.
)

// [AsyncIOOutcome] defines information about a completed asynchronous I/O request.
//
// [AsyncIOOutcome]: https://wiki.libsdl.org/SDL3/SDL_AsyncIOOutcome
type AsyncIOOutcome struct {
	Asyncio          *AsyncIO        // What generated this task. This pointer will be invalid if it was closed!
	Type             AsyncIOTaskType // What sort of task was this? Read, write, etc?
	Result           AsyncIOResult   // The result of the work (success, failure, cancellation).
	Buffer           uintptr         // Buffer where data was read/written.
	Offset           uint64          // Offset in the SDL_AsyncIO where data was read/written.
	BytesRequested   uint64          // Number of bytes the task was to read/write.
	BytesTransferred uint64          // Actual number of bytes that were read/written.
	Userdata         uintptr         // Pointer provided by the app when starting the task
}

// [AsyncIOQueue] is a queue of completed asynchronous I/O tasks.
//
// [AsyncIOQueue]: https://wiki.libsdl.org/SDL3/SDL_AsyncIOQueue
type AsyncIOQueue struct{}

// func AsyncIOFromFile(file string, mode string) *AsyncIO {
//	return sdlAsyncIOFromFile(file, mode)
// }

// func GetAsyncIOSize(asyncio *AsyncIO) int64 {
//	return sdlGetAsyncIOSize(asyncio)
// }

// func ReadAsyncIO(asyncio *AsyncIO, ptr unsafe.Pointer, offset uint64, size uint64, queue *AsyncIOQueue, userdata unsafe.Pointer) bool {
//	return sdlReadAsyncIO(asyncio, ptr, offset, size, queue, userdata)
// }

// func WriteAsyncIO(asyncio *AsyncIO, ptr unsafe.Pointer, offset uint64, size uint64, queue *AsyncIOQueue, userdata unsafe.Pointer) bool {
//	return sdlWriteAsyncIO(asyncio, ptr, offset, size, queue, userdata)
// }

// func CloseAsyncIO(asyncio *AsyncIO, flush bool, queue *AsyncIOQueue, userdata unsafe.Pointer) bool {
//	return sdlCloseAsyncIO(asyncio, flush, queue, userdata)
// }

// func CreateAsyncIOQueue() *AsyncIOQueue {
//	return sdlCreateAsyncIOQueue()
// }

// func DestroyAsyncIOQueue(queue *AsyncIOQueue)  {
//	sdlDestroyAsyncIOQueue(queue)
// }

// func GetAsyncIOResult(queue *AsyncIOQueue, outcome *AsyncIOOutcome) bool {
//	return sdlGetAsyncIOResult(queue, outcome)
// }

// func WaitAsyncIOResult(queue *AsyncIOQueue, outcome *AsyncIOOutcome, timeoutMS int32) bool {
//	return sdlWaitAsyncIOResult(queue, outcome, timeoutMS)
// }

// func SignalAsyncIOQueue(queue *AsyncIOQueue)  {
//	sdlSignalAsyncIOQueue(queue)
// }

// func LoadFileAsync(file string, queue *AsyncIOQueue, userdata unsafe.Pointer) bool {
//	return sdlLoadFileAsync(file, queue, userdata)
// }
