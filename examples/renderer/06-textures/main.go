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
	if !sdl.CreateWindowAndRenderer("examples/renderer/textures", WindowWidth, WindowHeight, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

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

		var dstRect sdl.FRect
		now := sdl.GetTicks()

		// we'll have some textures move around over a few seconds.
		var direction float32
		if now%2000 >= 1000 {
			direction = 1
		} else {
			direction = -1
		}

		scale := (float32(int(now%1000)-500) / 500) * direction

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// Just draw the static texture a few times. You can think of it like a
		// stamp, there isn't a limit to the number of times you can draw with it.

		// top left
		dstRect.X = 100.0 * scale
		dstRect.Y = 0.0
		dstRect.W = float32(texture.W)
		dstRect.H = float32(texture.H)
		sdl.RenderTexture(renderer, texture, nil, &dstRect)

		// center this one.
		dstRect.X = float32(WindowWidth-texture.W) / 2.0
		dstRect.Y = float32(WindowHeight-texture.H) / 2.0
		sdl.RenderTexture(renderer, texture, nil, &dstRect)

		// bottom right.
		dstRect.X = float32(WindowWidth-texture.W) - 100.0*scale
		dstRect.Y = float32(WindowHeight - texture.H)
		sdl.RenderTexture(renderer, texture, nil, &dstRect)

		sdl.RenderPresent(renderer)
	}

	sdl.DestroyTexture(texture)
}
