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
	if !sdl.CreateWindowAndRenderer("examples/renderer/rectangles", WindowWidth, WindowHeight, 0, &window, &renderer) {
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

		var rects [16]sdl.FRect
		now := sdl.GetTicks()

		// we'll have the rectangles grow and shrink over a few seconds.
		var direction float32
		if now%2000 >= 1000 {
			direction = 1
		} else {
			direction = -1
		}

		scale := (float32(int(now%1000)-500) / 500.0) * direction

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// Rectangles are comprised of set of X and Y coordinates, plus width and
		// height. (0, 0) is the top left of the window, and larger numbers go
		// down and to the right. This isn't how geometry works, but this is
		// pretty standard in 2D graphics.

		// Let's draw a single rectangle (square, really).
		rects[0].X, rects[0].Y = 100, 100
		rects[0].W, rects[0].H = 100+(100*scale), 100+(100*scale)
		sdl.SetRenderDrawColor(renderer, 255, 0, 0, sdl.AlphaOpaque)
		sdl.RenderRect(renderer, &rects[0])

		// Now let's draw several rectangles with one function call.
		for i := 0; i < 3; i++ {
			size := float32(i+1) * 50.0
			rects[i].W, rects[i].H = size+(size*scale), size+(size*scale)
			rects[i].X = (WindowWidth - rects[i].W) / 2
			rects[i].Y = (WindowHeight - rects[i].H) / 2
		}
		sdl.SetRenderDrawColor(renderer, 0, 255, 0, sdl.AlphaOpaque)
		sdl.RenderRects(renderer, rects[:3])

		// those were rectangle _outlines_, really. You can also draw _filled_ rectangles!
		rects[0].X = 400
		rects[0].Y = 50
		rects[0].W = 100 + (100 * scale)
		rects[0].H = 50 + (50 * scale)
		sdl.SetRenderDrawColor(renderer, 0, 0, 255, sdl.AlphaOpaque)
		sdl.RenderFillRect(renderer, &rects[0])

		// ...and also fill a bunch of rectangles at once...
		for i := 0; i < len(rects); i++ {
			w := float32(WindowWidth / len(rects))
			h := float32(i * 8.0)
			rects[i].X = float32(i) * w
			rects[i].Y = WindowHeight - h
			rects[i].W = w
			rects[i].H = h
		}
		sdl.SetRenderDrawColor(renderer, 255, 255, 255, sdl.AlphaOpaque)
		sdl.RenderFillRects(renderer, rects[:])

		sdl.RenderPresent(renderer)
	}
}
