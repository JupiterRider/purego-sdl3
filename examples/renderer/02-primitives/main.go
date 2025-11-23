package main

import (
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
	if !sdl.CreateWindowAndRenderer("examples/renderer/primitives", 640, 480, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	points := [500]sdl.FPoint{}

	for i := 0; i < len(points); i++ {
		points[i].X = (rand.Float32() * 440) + 100
		points[i].Y = (rand.Float32() * 280) + 100
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

		var rect sdl.FRect

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 33, 33, 33, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// draw a filled rectangle in the middle of the canvas.
		sdl.SetRenderDrawColor(renderer, 0, 0, 255, sdl.AlphaOpaque)
		rect.X, rect.Y = 100, 100
		rect.W = 440
		rect.H = 280
		sdl.RenderFillRect(renderer, &rect)

		// draw some points across the canvas.
		sdl.SetRenderDrawColor(renderer, 255, 0, 0, sdl.AlphaOpaque)
		sdl.RenderPoints(renderer, points[:])

		// draw a unfilled rectangle in-set a little bit.
		sdl.SetRenderDrawColor(renderer, 0, 255, 0, sdl.AlphaOpaque)
		rect.X += 30
		rect.Y += 30
		rect.W -= 60
		rect.H -= 60
		sdl.RenderRect(renderer, &rect)

		// draw two lines in an X across the whole canvas.
		sdl.SetRenderDrawColor(renderer, 255, 255, 0, sdl.AlphaOpaque)
		sdl.RenderLine(renderer, 0, 0, 640, 480)
		sdl.RenderLine(renderer, 0, 480, 640, 0)

		sdl.RenderPresent(renderer)
	}
}
