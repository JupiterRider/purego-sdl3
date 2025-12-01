package main

import (
	"fmt"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
		panic(sdl.GetError())
	}

	// When keyboard grab is enabled, SDL will continue to handle Alt+Tab when the window is full-screen to ensure the user is not trapped in your application.
	// Uncomment line below to disable that.
	// sdl.SetHint(sdl.HintAllowAltTabWhileGrabbed, "0")

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("Window modes", 960, 540, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	fmt.Println("Available display IDs:", sdl.GetDisplays())

	videoDriversCount := sdl.GetNumVideoDrivers()
	fmt.Print("Available video drivers: ")
	for i := int32(0); i < videoDriversCount; i++ {
		fmt.Print(sdl.GetVideoDriver(i))
		if i < videoDriversCount-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()

	fmt.Printf("Video driver in use: %s\r\n", sdl.GetCurrentVideoDriver())

	// Get available fullscreen display modes
	fullscreenDisplayModes := sdl.GetFullscreenDisplayModes(sdl.GetDisplayForWindow(window))
	fullscreenDisplayModeIdx, fullscreenDisplayModesCount := -1, 0
	if fullscreenDisplayModes != nil {
		fullscreenDisplayModesCount = len(fullscreenDisplayModes)
	}

	fullscreen, onTop, bordered, focusable, resizable := false, false, true, true, true
	opacity := sdl.GetWindowOpacity(window)

	running := true
	for running {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventWindowCloseRequested:
				running = false
			case sdl.EventKeyUp:
				e := event.Key()
				switch e.Key {
				case sdl.KeycodeK:
					sdl.SetWindowKeyboardGrab(window, !sdl.GetWindowKeyboardGrab(window))
				case sdl.KeycodeM:
					sdl.SetWindowMouseGrab(window, !sdl.GetWindowMouseGrab(window))
				case sdl.KeycodeF:
					fallthrough
				case sdl.KeycodeF11:
					fullscreen = !fullscreen
					sdl.SetWindowFullscreen(window, fullscreen)
				case sdl.KeycodeT:
					onTop = !onTop
					sdl.SetWindowAlwaysOnTop(window, onTop)
				case sdl.KeycodeB:
					bordered = !bordered
					sdl.SetWindowBordered(window, bordered)
				case sdl.KeycodeU:
					focusable = !focusable
					sdl.SetWindowFocusable(window, focusable)
				case sdl.KeycodeR:
					resizable = !resizable
					sdl.SetWindowResizable(window, resizable)
				case sdl.KeycodeUp:
					if fullscreenDisplayModesCount > 0 && fullscreenDisplayModesCount-fullscreenDisplayModeIdx > 1 {
						fullscreenDisplayModeIdx++
						sdl.SetWindowFullscreenMode(window, fullscreenDisplayModes[fullscreenDisplayModeIdx])
					}
				case sdl.KeycodeDown:
					if fullscreenDisplayModesCount > 0 && fullscreenDisplayModeIdx > 0 {
						fullscreenDisplayModeIdx--
						sdl.SetWindowFullscreenMode(window, fullscreenDisplayModes[fullscreenDisplayModeIdx])
					}
				case sdl.KeycodeLeft:
					if opacity > 0 {
						opacity -= 0.1
						if opacity < 0 {
							opacity = 0
						}
						if !sdl.SetWindowOpacity(window, opacity) {
							fmt.Println(sdl.GetError())
						}
					}
				case sdl.KeycodeRight:
					if opacity < 1 {
						opacity += 0.1
						if opacity > 1 {
							opacity = 1
						}
						sdl.SetWindowOpacity(window, opacity)
					}
				}
			}
		}

		sdl.SetRenderDrawColor(renderer, 30, 30, 30, 255)
		sdl.RenderClear(renderer)
		sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)

		sdl.RenderDebugText(renderer, 10, 10,
			fmt.Sprintf("%s [K] Keyboard grabbed, when enabled captures system keyboard shortcuts (Alt+Tab, Windows/Super key, etc.)",
				Tif(sdl.GetWindowKeyboardGrab(window), "+", " ")))

		sdl.RenderDebugText(renderer, 10, 30,
			fmt.Sprintf("%s [M] Mouse grabbed, when enabled mouse cursor is restricted to the window",
				Tif(sdl.GetWindowMouseGrab(window), "+", " ")))

		sdl.RenderDebugText(renderer, 10, 50, fmt.Sprintf("%s [F / F11] Window fullscreen", Tif(fullscreen, "+", " ")))

		sdl.RenderDebugText(renderer, 10, 70, fmt.Sprintf("%s [T] Window on top of the others", Tif(onTop, "+", " ")))

		sdl.RenderDebugText(renderer, 10, 90, fmt.Sprintf("%s [B] Window bordered", Tif(bordered, "+", " ")))

		sdl.RenderDebugText(renderer, 10, 110, fmt.Sprintf("%s [U] Window focusable", Tif(focusable, "+", " ")))

		sdl.RenderDebugText(renderer, 10, 130, fmt.Sprintf("%s [R] Window resizable", Tif(resizable, "+", " ")))

		mode := sdl.GetWindowFullscreenMode(window)
		if mode == nil {
			sdl.RenderDebugText(renderer, 10, 150, "  [Up / Down arrows] Window fullscreen mode: borderless fullscreen desktop resolution")
		} else {
			sdl.RenderDebugText(renderer, 10, 150, fmt.Sprintf("  [Up / Down arrows] Window fullscreen mode: %dx%d px %.2f Hz", mode.W, mode.H, mode.RefreshRate))
		}

		sdl.RenderDebugText(renderer, 10, 170, fmt.Sprintf("  [Left / Right arrows] Window opacity: %.0f%%", opacity*100))

		sdl.RenderPresent(renderer)
	}
}

// Ternary if statement
func Tif[T any](condition bool, vTrue, vFalse T) T {
	if condition {
		return vTrue
	}
	return vFalse
}
