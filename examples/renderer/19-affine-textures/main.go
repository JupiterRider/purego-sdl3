package main

import (
	"math"

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
	if !sdl.CreateWindowAndRenderer("examples/renderer/affine-textures", WindowWidth, WindowHeight, 0, &window, &renderer) {
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

		x0 := 0.5 * WindowWidth
		y0 := 0.5 * WindowHeight
		px := math.Min(WindowWidth, WindowHeight) / math.Sqrt(3.0)

		now := sdl.GetTicks()
		rad := (float64((now % 2000)) / 2000.0) * math.Pi * 2
		cos := math.Cos(rad)
		sin := math.Sin(rad)
		k := [3]float64{3.0 / math.Sqrt(50), 4.0 / math.Sqrt(50.0), 5.0 / math.Sqrt(50.0)}
		mat := [9]float64{
			cos + (1.0-cos)*k[0]*k[0], -sin*k[2] + (1.0-cos)*k[0]*k[1], sin*k[1] + (1.0-cos)*k[0]*k[2],
			sin*k[2] + (1.0-cos)*k[0]*k[1], cos + (1.0-cos)*k[1]*k[1], -sin*k[0] + (1.0-cos)*k[1]*k[2],
			-sin*k[1] + (1.0-cos)*k[0]*k[2], sin*k[0] + (1.0-cos)*k[1]*k[2], cos + (1.0-cos)*k[2]*k[2],
		}

		corners := [16]float64{}

		for i := 0; i < 8; i++ {
			var x, y, z float64
			if i&1 > 0 {
				x = -0.5
			} else {
				x = 0.5
			}
			if i&2 > 0 {
				y = -0.5
			} else {
				y = 0.5
			}
			if i&4 > 0 {
				z = -0.5
			} else {
				z = 0.5
			}
			corners[0+2*i] = mat[0]*x + mat[1]*y + mat[2]*z
			corners[1+2*i] = mat[3]*x + mat[4]*y + mat[5]*z
		}

		// as you can see from this, rendering draws over whatever was drawn before it.
		sdl.SetRenderDrawColor(renderer, 0x42, 0x87, 0xf5, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		for i := 1; i < 7; i++ {
			var dir, dirNum int
			if i&4 > 0 {
				dirNum = ^i
			} else {
				dirNum = i
			}
			dir = 3 & dirNum

			odd := (i & 1) ^ ((i & 2) >> 1) ^ ((i & 4) >> 2)

			var oddNum float64
			if odd == 1 {
				oddNum = 1.0
			} else {
				oddNum = -1.0
			}
			if 0 < (oddNum * mat[5+dir]) {
				continue
			}

			originIndex := (1 << ((dir - 1) % 3))
			rightIndex := (1 << ((dir + odd) % 3)) | originIndex
			downIndex := (1 << ((dir + (odd ^ 1)) % 3)) | originIndex
			if odd != 1 {
				originIndex ^= 7
				rightIndex ^= 7
				downIndex ^= 7
			}

			var origin, right, down sdl.FPoint
			origin.X = float32(x0 + px*corners[0+2*originIndex])
			origin.Y = float32(y0 + px*corners[1+2*originIndex])
			right.X = float32(x0 + px*corners[0+2*rightIndex])
			right.Y = float32(y0 + px*corners[1+2*rightIndex])
			down.X = float32(x0 + px*corners[0+2*downIndex])
			down.Y = float32(y0 + px*corners[1+2*downIndex])
			sdl.RenderTextureAffine(renderer, texture, nil, &origin, &right, &down)
		}

		sdl.RenderPresent(renderer)
	}

	sdl.DestroyTexture(texture)
}
