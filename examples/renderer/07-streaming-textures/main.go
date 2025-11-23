package main

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
)

const (
	WindowWidth  = 640
	WindowHeight = 480
	TextureSize  = 150
)

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/renderer/streaming-textures", WindowWidth, WindowHeight, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	texture := sdl.CreateTexture(renderer, sdl.PixelFormatRGBA8888, sdl.TextureAccessStreaming, TextureSize, TextureSize)
	if texture == nil {
		panic(sdl.GetError())
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

		var dstRect sdl.FRect
		now := sdl.GetTicks()
		var surface *sdl.Surface

		// we'll have some textures move around over a few seconds.
		var direction float32
		if now%2000 >= 1000 {
			direction = 1
		} else {
			direction = -1
		}

		scale := (float32(int(now%1000)-500) / 500) * direction

		// To update a streaming texture, you need to lock it first. This gets you access to the pixels.
		// Note that this is considered a _write-only_ operation: the buffer you get from locking
		// might not acutally have the existing contents of the texture, and you have to write to every
		// locked pixel!

		// You can use SDL_LockTexture() to get an array of raw pixels, but we're going to use
		// SDL_LockTextureToSurface() here, because it wraps that array in a temporary SDL_Surface,
		// letting us use the surface drawing functions instead of lighting up individual pixels.
		if sdl.LockTextureToSurface(texture, nil, &surface) {
			var r sdl.Rect
			sdl.FillSurfaceRect(surface, nil, sdl.MapRGB(sdl.GetPixelFormatDetails(surface.Format), nil, 0, 0, 0))
			r.W = TextureSize
			r.H = TextureSize / 10
			r.X = 0
			r.Y = int32(float32(TextureSize-r.H) * ((scale + 1.0) / 2.0))
			sdl.FillSurfaceRect(surface, &r, sdl.MapRGB(sdl.GetPixelFormatDetails(surface.Format), nil, 0, 255, 0))
			sdl.UnlockTexture(texture)
		}

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 66, 66, 66, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// Just draw the static texture a few times. You can think of it like a
		// stamp, there isn't a limit to the number of times you can draw with it.

		// center this one.
		dstRect.X = float32(WindowWidth-TextureSize) / 2.0
		dstRect.Y = float32(WindowHeight-TextureSize) / 2.0
		dstRect.W, dstRect.H = float32(TextureSize), float32(TextureSize)
		sdl.RenderTexture(renderer, texture, nil, &dstRect)

		sdl.RenderPresent(renderer)
	}

	sdl.DestroyTexture(texture)
}
