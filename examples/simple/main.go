package main

import (
	sdl "github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	if !sdl.SetHint(sdl.HintRenderVsync, "1") {
		panic(sdl.GetError())
	}

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("Hello, World!", 1280, 720, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer renderer.Destroy()
	defer window.Destroy()

	renderer.SetDrawColor(100, 150, 200, 255)

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
		renderer.Clear()
		renderer.Present()
	}
}
