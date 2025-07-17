package main

import (
	"fmt"
	"os"

	"github.com/jupiterrider/purego-sdl3/img"
	"github.com/jupiterrider/purego-sdl3/sdl"
)

// Size at which the sprite should be rendered
// Decrease to draw more sprites and better see the impact of batch drawing compared to individual draw calls
var spriteSize float32 = 32

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	// sdl.SetHint(sdl.HintRenderVSync, "1") // VSync shouldn't be used to see performance impact of batched vs individual drawing

	const width, height = 1600, 900
	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("Sprite batching", width, height, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	// Load sprite into surface and render it into texture
	gopherSurface := img.Load("gopher-happy.png")
	if gopherSurface == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroySurface(gopherSurface)
	gopherTexture := sdl.CreateTextureFromSurface(renderer, gopherSurface)
	if gopherTexture == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyTexture(gopherTexture)
	sdl.SetTextureScaleMode(gopherTexture, sdl.ScaleModeNearest) // Required for pixel-art sprites

	// Precalculate positions at which the sprites should be rendered
	xStep, yStep := spriteSize+2, spriteSize+2 // 2 px padding between sprites
	positions := make([]sdl.FPoint, 0)
	for yPos := float32(0); yPos < height; yPos += yStep {
		for xPos := float32(0); xPos < width; xPos += xStep {
			positions = append(positions, sdl.FPoint{X: xPos, Y: yPos})
		}
	}

	// Create arrays required for batched drawing
	// These arrays can be inside the render loop but it would impact the performance
	white := sdl.FColor{R: 1, G: 1, B: 1, A: 1}       // Just a white color
	xy := make([]sdl.FPoint, 0, len(positions)*4)     // Positions and sizes at which sprites should be rendered
	colors := make([]sdl.FColor, 0, len(positions)*4) // Color modulation of each of the vertex of the sprite that will be rendered
	uv := make([]sdl.FPoint, 0, len(positions)*4)     // Normalized texture coordinates of a sprite position and size inside texture atlas
	indices := make([]int32, 0, len(positions)*6)     // Indexes of the vertices that create a triangle
	for i := int32(0); i < int32(len(positions)); i++ {
		p := positions[i]
		xy = append(xy, []sdl.FPoint{
			{X: p.X, Y: p.Y},                           // Top left
			{X: p.X + spriteSize, Y: p.Y},              // Top right
			{X: p.X, Y: p.Y + spriteSize},              // Bottom left
			{X: p.X + spriteSize, Y: p.Y + spriteSize}, // Bottom right
		}...)

		colors = append(colors, []sdl.FColor{
			white, // For each of the vertex use color white (don't modify the sprite)
			white,
			white,
			white,
		}...)

		uv = append(uv, []sdl.FPoint{
			{X: 0, Y: 0}, // Normalized texture coordinate of top left sprite corner, it can be calculated like this: {X: atlas_texture.W / sprite_position_x, Y: atlas_texture.H / sprite_position_y}
			{X: 1, Y: 0}, // Normalized texture coordinate of top right sprite corner, it can be calculated like this: {X: atlas_texture.W / (sprite_position_x + sprite_width), Y: atlas_texture.H / sprite_position_y}
			{X: 0, Y: 1}, // Normalized texture coordinate of bottom left sprite corner, it can be calculated like this: {X: atlas_texture.W / sprite_position_x, Y: atlas_texture.H / (sprite_position_y + sprite_height)}
			{X: 1, Y: 1}, // Normalized texture coordinate of bottom right sprite corner, it can be calculated like this: {X: atlas_texture.W / (sprite_position_x + sprite_width), Y: atlas_texture.H / (sprite_position_y + sprite_height)}
		}...)

		indices = append(indices, []int32{
			(4 * i) + 1, (4 * i) + 0, (4 * i) + 2, // First triangle
			(4 * i) + 1, (4 * i) + 2, (4 * i) + 3, // Second triangle
		}...)
	}

	fmt.Println("Hold space to use individual draw calls instead of batched drawing")
	fmt.Printf("Rendering %d sprites of size %.fx%.f px\n", len(positions), spriteSize, spriteSize)

	running, individualRendering := true, false
	frameStart, perfFreq, frameCount := sdl.GetPerformanceCounter(), float32(sdl.GetPerformanceFrequency()), 0
	for running {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit, sdl.EventWindowCloseRequested:
				running = false
			case sdl.EventKeyDown:
				e := event.Key()
				if e.Key == sdl.KeycodeSpace {
					individualRendering = true
				}
			case sdl.EventKeyUp:
				e := event.Key()
				if e.Key == sdl.KeycodeSpace {
					individualRendering = false
				}
			}
		}

		sdl.SetRenderDrawColor(renderer, 30, 30, 30, 255)
		sdl.RenderClear(renderer)

		if individualRendering {
			// Draw everything using individual draw calls
			for i := 0; i < len(positions); i++ {
				p := positions[i]
				if !sdl.RenderTexture(renderer, gopherTexture, nil, &sdl.FRect{X: p.X, Y: p.Y, W: spriteSize, H: spriteSize}) {
					fmt.Println(sdl.GetError())
				}
			}
		} else {
			// Draw everything using single call
			if !sdl.RenderGeometryRaw(renderer, gopherTexture, xy, colors, uv, indices) {
				fmt.Println(sdl.GetError())
			}
		}

		// Performance metrics, FPS and frame time
		frameEnd := sdl.GetPerformanceCounter()
		frameCount++
		dt := float32(frameEnd-frameStart) / perfFreq
		if dt >= 1 {
			fps := float32(frameCount) / dt
			frameTime := float32(1000) / fps
			frameCount = 0
			frameStart = frameEnd
			fmt.Fprintf(os.Stdout, " %.0f FPS, frame time: %.4f ms, using %s drawing mode          \r", fps, frameTime, Tif(individualRendering, "individual", "batched"))
		}

		sdl.RenderPresent(renderer)
	}

	fmt.Println() // Additional print line for the last performance metrics to be left in console
}

// Ternary if statement
func Tif[T any](condition bool, vTrue, vFalse T) T {
	if condition {
		return vTrue
	}
	return vFalse
}
