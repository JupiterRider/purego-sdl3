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
	if !sdl.CreateWindowAndRenderer("examples/renderer/viewport", WindowWidth, WindowHeight, 0, &window, &renderer) {
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

		dstRect := sdl.FRect{X: 0, Y: 0, W: float32(texture.W), H: float32(texture.H)}

		var viewport sdl.Rect

		// Setting a viewport has the effect of limiting the area that rendering
		// can happen, and making coordinate (0, 0) live somewhere else in the
		// window. It does _not_ scale rendering to fit the viewport.

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// Draw once with the whole window as the viewport.
		viewport.X = 0
		viewport.Y = 0
		viewport.W = WindowWidth / 2
		viewport.H = WindowHeight / 2
		sdl.SetRenderViewport(renderer, nil)
		sdl.RenderTexture(renderer, texture, nil, &dstRect)

		// top right quarter of the window.
		viewport.X = WindowWidth / 2
		viewport.Y = WindowHeight / 2
		sdl.SetRenderViewport(renderer, &viewport)
		sdl.RenderTexture(renderer, texture, nil, &dstRect)

		// bottom 20% of the window. Note it clips the width!
		viewport.X = 0
		viewport.Y = WindowHeight - WindowHeight/5
		viewport.W = WindowWidth / 5
		viewport.H = WindowHeight / 5
		sdl.SetRenderViewport(renderer, &viewport)
		sdl.RenderTexture(renderer, texture, nil, &dstRect)

		// what happens if you try to draw above the viewport? It should clip!
		viewport.X = 100
		viewport.Y = 200
		viewport.W = WindowWidth
		viewport.H = WindowHeight
		sdl.SetRenderViewport(renderer, &viewport)
		dstRect.Y = -50
		sdl.RenderTexture(renderer, texture, nil, &dstRect)

		sdl.RenderPresent(renderer)
	}

	sdl.DestroyTexture(texture)
}
