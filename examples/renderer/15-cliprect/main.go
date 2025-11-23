package main

import (
	"math"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

const (
	WindowWidth  = 640
	WindowHeight = 480

	ClipRectSize  = 250
	ClipRectSpeed = 200
)

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/renderer/cliprect", WindowWidth, WindowHeight, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	var clipRectPosition sdl.FPoint
	clipRectDirection := sdl.FPoint{X: 1, Y: 1}

	lastTime := sdl.GetTicks()

	// Textures are pixel data that we upload to the video hardware for fast drawing. Lots of 2D
	// engines refer to these as "sprites." We'll do a static texture (upload once, draw many
	// times) with data from a bitmap file.

	// SDL_Surface is pixel data the CPU can access. SDL_Texture is pixel data the GPU can access.
	// Load a .bmp into a surface, move it to a texture from there.
	surface := sdl.LoadBMP(sdl.GetBasePath() + "sample.bmp")
	if surface == nil {
		panic(sdl.GetError())
	}

	texture := sdl.CreateTextureFromSurface(renderer, surface)
	if texture == nil {
		panic(sdl.GetError())
	}

	sdl.DestroySurface(surface)

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

		cliprect := sdl.Rect{
			X: int32(math.Round(float64(clipRectPosition.X))),
			Y: int32(math.Round(float64(clipRectPosition.Y))),
			W: ClipRectSize,
			H: ClipRectSize,
		}
		now := sdl.GetTicks()
		elapsed := (float32(now - lastTime)) / 1000
		distance := elapsed * ClipRectSpeed

		// Set a new clipping rectangle position
		clipRectPosition.X += distance * clipRectDirection.X
		if clipRectPosition.X < 0 {
			clipRectPosition.X = 0
			clipRectDirection.X = 1
		} else if clipRectPosition.X >= (WindowWidth - ClipRectSize) {
			clipRectPosition.X = WindowWidth - ClipRectSize - 1
			clipRectDirection.X = -1
		}

		clipRectPosition.Y += distance * clipRectDirection.Y
		if clipRectPosition.Y < 0 {
			clipRectPosition.Y = 0
			clipRectDirection.Y = 1
		} else if clipRectPosition.Y >= (WindowHeight - ClipRectSize) {
			clipRectPosition.Y = WindowHeight - ClipRectSize - 1
			clipRectDirection.Y = -1
		}
		sdl.SetRenderClipRect(renderer, &cliprect)

		lastTime = now

		// okay, now draw!

		// Note that SDL_RenderClear is _not_ affected by the clipping rectangle!
		sdl.SetRenderDrawColor(renderer, 33, 33, 33, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// stretch the texture across the entire window. Only the piece in the
		// clipping rectangle will actually render, though!
		sdl.RenderTexture(renderer, texture, nil, nil)

		sdl.RenderPresent(renderer)
	}

	sdl.DestroyTexture(texture)
}
