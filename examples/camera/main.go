package main

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	if !sdl.Init(sdl.InitVideo | sdl.InitCamera) {
		panic(sdl.GetError())
	}
	defer sdl.Quit()

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/camera", 640, 480, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer renderer.Destroy()
	defer window.Destroy()

	renderer.SetVSync(1)

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
	defer camera.Close()

	var texture *sdl.Texture
	defer func() {
		if texture != nil {
			texture.Destroy()
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
		frame := camera.AcquireFrame(&timestampNS)
		if frame != nil {
			if texture == nil {
				window.SetSize(frame.W, frame.H)
				texture = renderer.CreateTexture(frame.Format, sdl.TextureAccessStreaming, frame.W, frame.H)
			}

			if texture != nil {
				texture.UpdateTexture(nil, frame.Pixels, frame.Pitch)
			}

			camera.ReleaseFrame(frame)
		}

		renderer.SetDrawColor(0x99, 0x99, 0x99, 255)
		renderer.Clear()
		if texture != nil {
			renderer.RenderTexture(texture, nil, nil)
		}
		renderer.Present()
	}
}
