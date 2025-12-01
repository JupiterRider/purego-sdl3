package main

import (
	"math"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/renderer/clear", 640, 480, 0, &window, &renderer) {
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

		now := float64(sdl.GetTicks()) / 1000.0
		// choose the color for the frame we will draw. The sine wave trick makes it fade between colors smoothly.
		red := 0.5 + 0.5*float32(math.Sin(now))
		green := 0.5 + 0.5*float32(math.Sin(now+math.Pi*2/3))
		blue := 0.5 + 0.5*float32(math.Sin(now+math.Pi*4/3))

		sdl.SetRenderDrawColorFloat(renderer, red, green, blue, 1)

		// clear the window to the draw color.
		sdl.RenderClear(renderer)

		// put the newly-cleared rendering on the screen.
		sdl.RenderPresent(renderer)
	}
}
