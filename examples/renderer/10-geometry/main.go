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
	if !sdl.CreateWindowAndRenderer("examples/renderer/geometry", WindowWidth, WindowHeight, 0, &window, &renderer) {
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

		now := sdl.GetTicks()

		// we'll have some textures move around over a few seconds.
		var direction float32
		if now%2000 >= 1000 {
			direction = 1
		} else {
			direction = -1
		}

		scale := (float32(int(now%1000)-500) / 500.0) * direction
		size := 200.0 + 200.0*scale

		var vertices [4]sdl.Vertex

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// Draw a single triangle with a different color at each vertex. Center this one and make it grow and shrink.
		// You always draw triangles with this, but you can string triangles together to form polygons.
		vertices[0].Position.X = float32(WindowWidth) / 2
		vertices[0].Position.Y = float32(WindowHeight-size) / 2
		vertices[0].Color.R = 1
		vertices[0].Color.A = 1
		vertices[1].Position.X = float32(WindowWidth+size) / 2
		vertices[1].Position.Y = float32(WindowHeight+size) / 2
		vertices[1].Color.G = 1
		vertices[1].Color.A = 1
		vertices[2].Position.X = float32(WindowWidth-size) / 2
		vertices[2].Position.Y = float32(WindowHeight+size) / 2
		vertices[2].Color.B = 1
		vertices[2].Color.A = 1

		sdl.RenderGeometry(renderer, nil, vertices[:3], nil)

		// you can also map a texture to the geometry! Texture coordinates go from 0.0f to 1.0f. That will be the location
		// in the texture bound to this vertex.
		vertices = [4]sdl.Vertex{}
		vertices[0].Position.X = 10
		vertices[0].Position.Y = 10
		vertices[0].Color.R = 1
		vertices[0].Color.G = 1
		vertices[0].Color.B = 1
		vertices[0].Color.A = 1
		vertices[0].TexCoord.X = 0
		vertices[0].TexCoord.Y = 0
		vertices[1].Position.X = 150
		vertices[1].Position.Y = 10
		vertices[1].Color.R = 1
		vertices[1].Color.G = 1
		vertices[1].Color.B = 1
		vertices[1].Color.A = 1
		vertices[1].TexCoord.X = 1
		vertices[1].TexCoord.Y = 0
		vertices[2].Position.X = 10
		vertices[2].Position.Y = 150
		vertices[2].Color.R = 1
		vertices[2].Color.G = 1
		vertices[2].Color.B = 1
		vertices[2].Color.A = 1
		vertices[2].TexCoord.X = 0
		vertices[2].TexCoord.Y = 1
		sdl.RenderGeometry(renderer, texture, vertices[:3], nil)

		// Did that only draw half of the texture? You can do multiple triangles sharing some vertices,
		// using indices, to get the whole thing on the screen:

		// Let's just move this over so it doesn't overlap...
		for i := 0; i < 3; i++ {
			vertices[i].Position.X += 450
		}

		// we need one more vertex, since the two triangles can share two of them.
		vertices[3].Position.X = 600
		vertices[3].Position.Y = 150
		vertices[3].Color.R = 1
		vertices[3].Color.G = 1
		vertices[3].Color.B = 1
		vertices[3].Color.A = 1
		vertices[3].TexCoord.X = 1
		vertices[3].TexCoord.Y = 1

		// And an index to tell it to reuse some of the vertices between triangles...
		// 4 vertices, but 6 actual places they used. Indices need less bandwidth to transfer and can reorder vertices easily!
		indices := []int32{0, 1, 2, 1, 2, 3}
		sdl.RenderGeometry(renderer, texture, vertices[:], indices)

		sdl.RenderPresent(renderer)
	}

	sdl.DestroyTexture(texture)
}
