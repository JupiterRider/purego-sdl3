package sdl

type Window struct{}

type WindowFlags uint64

const (
	WindowOpenGL    WindowFlags = 0x0000000000000002
	WindowResizable WindowFlags = 0x0000000000000020
)

type WindowID uint32

// DestroyWindow destroys a window.
func DestroyWindow(window *Window) {
	sdlDestroyWindow(window)
}

// Destroy destroys a window.
func (window *Window) Destroy() {
	DestroyWindow(window)
}
