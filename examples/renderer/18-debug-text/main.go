package main

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
)

const (
	WindowWidth  = 640
	WindowHeight = 480
)

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/renderer/debug-text", WindowWidth, WindowHeight, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	isRun := true

	for isRun {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				isRun = false
			case sdl.EventKeyDown:
				if event.Key().Scancode == sdl.ScancodeEscape {
					isRun = false
				}
			}
		}

		charSize := sdl.DebugTextFontCharacterSize

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		sdl.SetRenderDrawColor(renderer, 255, 255, 255, sdl.AlphaOpaque)
		sdl.RenderDebugText(renderer, 272, 100, "Hello world!")
		sdl.RenderDebugText(renderer, 224, 150, "This is some debug text.")

		sdl.SetRenderDrawColor(renderer, 51, 102, 255, sdl.AlphaOpaque)
		sdl.RenderDebugText(renderer, 184, 200, "You can do it in different colors.")
		sdl.SetRenderDrawColor(renderer, 255, 255, 255, sdl.AlphaOpaque)

		sdl.SetRenderScale(renderer, 4.0, 4.0)
		sdl.RenderDebugText(renderer, 14, 65, "It can be scaled.")
		sdl.SetRenderScale(renderer, 1.0, 1.0)
		sdl.RenderDebugText(renderer, 64, 350, "This only does ASCII chars. So this laughing emoji won't draw: 🤣")

		sdl.RenderDebugTextFormat(renderer, float32((WindowWidth-charSize*46)/2), 400, "(This program has been running for %d seconds.)", sdl.GetTicks()/1000)

		sdl.RenderPresent(renderer)
	}
}
