package main

import "github.com/jupiterrider/purego-sdl3/sdl"

func main() {
	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo | sdl.InitCamera) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/camera", 640, 480, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	sdl.SetRenderVSync(renderer, 1)

	cameras := sdl.GetCameras()
	if cameras == nil {
		panic(sdl.GetError())
	}

	if len(cameras) == 0 {
		panic("couldn't find any connected camera device")
	}

	camera := sdl.OpenCamera(cameras[0], nil)
	if camera == nil {
		panic(sdl.GetError())
	}
	defer sdl.CloseCamera(camera)

	var texture *sdl.Texture
	defer func() {
		if texture != nil {
			sdl.DestroyTexture(texture)
		}
	}()

Outer:
	for {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				break Outer
			case sdl.EventKeyDown:
				if event.Key().Scancode == sdl.ScancodeEscape {
					break Outer
				}
			}
		}

		var timestampNS uint64
		frame := sdl.AcquireCameraFrame(camera, &timestampNS)
		if frame != nil {
			if texture == nil {
				sdl.SetWindowSize(window, frame.W, frame.H)
				texture = sdl.CreateTexture(renderer, frame.Format, sdl.TextureAccessStreaming, frame.W, frame.H)
			}

			if texture != nil {
				sdl.UpdateTexture(texture, nil, frame.Pixels, frame.Pitch)
			}

			sdl.ReleaseCameraFrame(camera, frame)
		}

		sdl.SetRenderDrawColor(renderer, 0x99, 0x99, 0x99, 255)
		sdl.RenderClear(renderer)
		if texture != nil {
			sdl.RenderTexture(renderer, texture, nil, nil)
		}
		sdl.RenderPresent(renderer)
	}
}
