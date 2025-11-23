package main

import (
	"math"
	"math/rand"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/renderer/lines", 640, 480, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	line_points := []sdl.FPoint{
		{X: 100, Y: 354}, {X: 220, Y: 230}, {X: 140, Y: 230}, {X: 320, Y: 100}, {X: 500, Y: 230},
		{X: 420, Y: 230}, {X: 540, Y: 354}, {X: 400, Y: 354}, {X: 100, Y: 354},
	}

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

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 100, 100, 100, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// You can draw lines, one at a time, like these brown ones...
		sdl.SetRenderDrawColor(renderer, 127, 49, 32, sdl.AlphaOpaque)
		sdl.RenderLine(renderer, 240, 450, 400, 450)
		sdl.RenderLine(renderer, 240, 356, 400, 356)
		sdl.RenderLine(renderer, 240, 356, 240, 450)
		sdl.RenderLine(renderer, 400, 356, 400, 450)

		// You can also draw a series of connected lines in a single batch...
		sdl.SetRenderDrawColor(renderer, 0, 255, 0, sdl.AlphaOpaque)
		sdl.RenderLines(renderer, line_points[:])

		// here's a bunch of lines drawn out from a center point in a circle.
		// we randomize the color of each line, so it functions as animation.
		for i := 0; i < 360; i++ {
			var size, x, y float32
			size = 30.0
			x = 320.0
			y = 95.0 - (size / 2.0)
			sdl.SetRenderDrawColor(renderer,
				uint8(rand.Intn(256)),
				uint8(rand.Intn(256)),
				uint8(rand.Intn(256)),
				sdl.AlphaOpaque)
			sdl.RenderLine(renderer, x, y,
				x+float32(math.Sin(float64(i)))*size,
				y+float32(math.Cos(float64(i)))*size)
		}

		sdl.RenderPresent(renderer)
	}
}
