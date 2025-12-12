package sdl

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

// func CreateProcess(args **byte, pipe_stdio bool) *Process {
//	return sdlCreateProcess(args, pipe_stdio)
// }

// func CreateProcessWithProperties(props PropertiesID) *Process {
//	return sdlCreateProcessWithProperties(props)
// }

// func DestroyProcess(process *Process)  {
//	sdlDestroyProcess(process)
// }

// func GetProcessInput(process *Process) *IOStream {
//	return sdlGetProcessInput(process)
// }

// func GetProcessOutput(process *Process) *IOStream {
//	return sdlGetProcessOutput(process)
// }

// func GetProcessProperties(process *Process) PropertiesID {
//	return sdlGetProcessProperties(process)
// }

// func KillProcess(process *Process, force bool) bool {
//	return sdlKillProcess(process, force)
// }

// func ReadProcess(process *Process, datasize *uint64, exitcode *int32) unsafe.Pointer {
//	return sdlReadProcess(process, datasize, exitcode)
// }

// func WaitProcess(process *Process, block bool, exitcode *int32) bool {
//	return sdlWaitProcess(process, block, exitcode)
// }
