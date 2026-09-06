package sdl

const (
	PropProcessCreateArgsPointer            = "SDL.process.create.args"
	PropProcessCreateEnvironmentPointer     = "SDL.process.create.environment"
	PropProcessCreateWorkingDirectoryString = "SDL.process.create.working_directory"
	PropProcessCreateStdinNumber            = "SDL.process.create.stdin_option"
	PropProcessCreateStdinPointer           = "SDL.process.create.stdin_source"
	PropProcessCreateStdoutNumber           = "SDL.process.create.stdout_option"
	PropProcessCreateStdoutPointer          = "SDL.process.create.stdout_source"
	PropProcessCreateStderrNumber           = "SDL.process.create.stderr_option"
	PropProcessCreateStderrPointer          = "SDL.process.create.stderr_source"
	PropProcessCreateStderrToStdoutBoolean  = "SDL.process.create.stderr_to_stdout"
	PropProcessCreateBackgroundBoolean      = "SDL.process.create.background"
	PropProcessCreateCmdlineString          = "SDL.process.create.cmdline"
	PropProcessPidNumber                    = "SDL.process.pid"
	PropProcessStdinPointer                 = "SDL.process.stdin"
	PropProcessStdoutPointer                = "SDL.process.stdout"
	PropProcessStderrPointer                = "SDL.process.stderr"
	PropProcessBackgroundBoolean            = "SDL.process.background"
)

// [Process] is an opaque handle representing a system process.
//
// [Process]: https://wiki.libsdl.org/SDL3/SDL_Process
type Process struct{}

// [CreateProcess] creates a new process.
//
// [CreateProcess]: https://wiki.libsdl.org/SDL3/SDL_CreateProcess
// func CreateProcess(args **byte, pipe_stdio bool) *Process {
//	return sdlCreateProcess(args, pipe_stdio)
// }

// [ProcessIO] defines the description of where standard I/O should be directed when creating a process.
//
// [ProcessIO]: https://wiki.libsdl.org/SDL3/SDL_ProcessIO
type ProcessIO uint32

const (
	ProcessStdioInherited ProcessIO = iota // The I/O stream is inherited from the application.
	ProcessStdioNull                       // The I/O stream is ignored.
	ProcessStdioApp                        // The I/O stream is connected to a new [IOStream] that the application can read or write.
	ProcessStdioRedirect                   // The I/O stream is redirected to an existing [IOStream].
)

// [CreateProcessWithProperties] creates a new process with the specified properties.
//
// [CreateProcessWithProperties]: https://wiki.libsdl.org/SDL3/SDL_CreateProcessWithProperties
// func CreateProcessWithProperties(props PropertiesID) *Process {
//	return sdlCreateProcessWithProperties(props)
// }

// [GetProcessProperties] gets the properties associated with a process.
//
// [GetProcessProperties]: https://wiki.libsdl.org/SDL3/SDL_GetProcessProperties
// func GetProcessProperties(process *Process) PropertiesID {
//	return sdlGetProcessProperties(process)
// }

// [ReadProcess] reads all the output from a process.
//
// [ReadProcess]: https://wiki.libsdl.org/SDL3/SDL_ReadProcess
// func ReadProcess(process *Process, datasize *uint64, exitcode *int32) unsafe.Pointer {
//	return sdlReadProcess(process, datasize, exitcode)
// }

// [GetProcessInput] gets the [IOStream] associated with process standard input.
//
// [GetProcessInput]: https://wiki.libsdl.org/SDL3/SDL_GetProcessInput
// func GetProcessInput(process *Process) *IOStream {
//	return sdlGetProcessInput(process)
// }

// [GetProcessOutput] gets the [IOStream] associated with process standard output.
//
// [GetProcessOutput]: https://wiki.libsdl.org/SDL3/SDL_GetProcessOutput
// func GetProcessOutput(process *Process) *IOStream {
//	return sdlGetProcessOutput(process)
// }

// [KillProcess] stops a process.
//
// [KillProcess]: https://wiki.libsdl.org/SDL3/SDL_KillProcess
// func KillProcess(process *Process, force bool) bool {
//	return sdlKillProcess(process, force)
// }

// [WaitProcess] waits for a process to finish.
//
// [WaitProcess]: https://wiki.libsdl.org/SDL3/SDL_WaitProcess
// func WaitProcess(process *Process, block bool, exitcode *int32) bool {
//	return sdlWaitProcess(process, block, exitcode)
// }

// [DestroyProcess] destroys a previously created process object.
//
// [DestroyProcess]: https://wiki.libsdl.org/SDL3/SDL_DestroyProcess
// func DestroyProcess(process *Process)  {
//	sdlDestroyProcess(process)
// }
