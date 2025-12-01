package main

import (
	"math/rand"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

const (
	WindowWidth  = 640
	WindowHeight = 480

	NumPoints          = 500
	MinPixelsPerSecond = 30
	MaxPixelsPerSecond = 60
)

var (
	points      [NumPoints]sdl.FPoint
	pointSpeeds [NumPoints]float32
	lastTime    uint64
)

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/renderer/points", WindowWidth, WindowHeight, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	for i := 0; i < len(points); i++ {
		points[i].X = rand.Float32() * WindowWidth
		points[i].Y = rand.Float32() * WindowHeight
		pointSpeeds[i] = MinPixelsPerSecond + (rand.Float32() * (MaxPixelsPerSecond - MinPixelsPerSecond))
	}

	lastTime = sdl.GetTicks()

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

		now := sdl.GetTicks()
		elapsed := float32(now-lastTime) / 1000.0

		// let's move all our points a little for a new frame.
		for i := 0; i < len(points); i++ {
			distance := elapsed * pointSpeeds[i]
			points[i].X += distance
			points[i].Y += distance
			if points[i].X >= WindowWidth || points[i].Y >= WindowHeight {
				// off the screen; restart it elsewhere!
				if rand.Intn(2) != 0 {
					points[i].X = rand.Float32() * float32(WindowWidth)
					points[i].Y = 0.0
				} else {
					points[i].X = 0.0
					points[i].Y = rand.Float32() * float32(WindowHeight)
				}
				pointSpeeds[i] = MinPixelsPerSecond + (rand.Float32() * (MaxPixelsPerSecond - MinPixelsPerSecond))
			}
		}

		lastTime = now

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		sdl.SetRenderDrawColor(renderer, 255, 255, 255, sdl.AlphaOpaque)
		sdl.RenderPoints(renderer, points[:])

		// You can also draw single points with SDL_RenderPoint(), but it's
		// cheaper (sometimes significantly so) to do them all at once.

		sdl.RenderPresent(renderer)
	}
}
