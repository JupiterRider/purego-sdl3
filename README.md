# purego-sdl3
[![Go Reference](https://pkg.go.dev/badge/github.com/jupiterrider/purego-sdl3/sdl.svg)](https://pkg.go.dev/github.com/jupiterrider/purego-sdl3/sdl)

A cgo-free SDL3 binding.

## About
This library doesn't require cgo. It uses [dynamic loading](https://en.wikipedia.org/wiki/Dynamic_loading) with the help of [purego](https://github.com/ebitengine/purego).

## Status
This project is in a very, very early stage and currently only a proof of concept.

## Requirements
You need to have SDL3 installed as shared library. That means at runtime, it is trying to load SDL:
- macOS: `libSDL3.dylib`
- Linux and FreeBSD: `libSDL3.so.0`
- Windows: `SDL3.dll`

Only the above-mentioned operating systems with AMD64 or ARM64 architecture are supported.

## Example
This simple example just opens a resizable window with a blue background:

```golang
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
```
