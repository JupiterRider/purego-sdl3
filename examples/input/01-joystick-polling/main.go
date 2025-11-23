package main

import (
	"math"
	"math/rand"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

var colors [64]sdl.Color

func main() {
	for i := 0; i < len(colors); i++ {
		colors[i] = sdl.Color{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}
	}

	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
		panic(sdl.GetError())
	}

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo | sdl.InitJoystick) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("Polling Example", 640, 480, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	var joystick *sdl.Joystick

	running := true
	for running {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventJoystickAdded:
				if joystick == nil {
					joystick = sdl.OpenJoystick(event.JDevice().Which)
				}
			case sdl.EventJoystickRemoved:
				if joystick != nil && event.JDevice().Which == sdl.GetJoystickID(joystick) {
					sdl.CloseJoystick(joystick)
					joystick = nil
				}
			case sdl.EventQuit:
				running = false
			case sdl.EventKeyDown:
				if event.Key().Scancode == sdl.ScancodeEscape {
					running = false
				}
			}
		}
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, 255)
		sdl.RenderClear(renderer)
		draw(renderer, window, joystick)
		sdl.RenderPresent(renderer)
	}
}

func draw(renderer *sdl.Renderer, window *sdl.Window, joystick *sdl.Joystick) {
	text := "Plug in a joystick, please."
	if joystick != nil {
		text = sdl.GetJoystickName(joystick)
	}

	var winw, winh int32 = 640, 480
	sdl.GetWindowSize(window, &winw, &winh)

	// note that you can get input as events, instead of polling, which is
	// better since it won't miss button presses if the system is lagging,
	// but often times checking the current state per-frame is good enough,
	// and maybe better if you'd rather _drop_ inputs due to lag.

	if joystick != nil {
		const size = 30.0
		fwinw := float32(winw)
		fwinh := float32(winh)

		// draw axes
		totalAxes := sdl.GetNumJoystickAxes(joystick)
		y := (fwinh - float32(totalAxes)*size) / 2
		x := fwinw / 2
		for i := 0; i < int(totalAxes); i++ {
			color := colors[i%len(colors)]
			val := float32(sdl.GetJoystickAxis(joystick, int32(i))) / math.MaxInt16
			dx := x + val*x
			rectWidth := x - abs(dx)

			rect := sdl.FRect{X: dx, Y: y, W: rectWidth, H: size}
			sdl.SetRenderDrawColor(renderer, color.R, color.G, color.B, color.A)
			sdl.RenderFillRect(renderer, &rect)
			y += size
		}

		// draw buttons
		totalButtons := sdl.GetNumJoystickButtons(joystick)
		x = (fwinw - float32(totalButtons)*size) / 2
		for i := 0; i < int(totalButtons); i++ {
			color := colors[i%len(colors)]
			if sdl.GetJoystickButton(joystick, int32(i)) {
				sdl.SetRenderDrawColor(renderer, color.R, color.G, color.B, color.A)
			} else {
				sdl.SetRenderDrawColor(renderer, 0, 0, 0, 255)
			}

			rect := sdl.FRect{X: x, Y: 0, W: size, H: size}
			sdl.RenderFillRect(renderer, &rect)
			sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
			sdl.RenderRect(renderer, &rect)
			x += size
		}

		// draw hats
		totalHats := sdl.GetNumJoystickHats(joystick)
		x = (fwinw-float32(totalHats)*(size*2))/2 + size/2
		y = fwinh - size
		for i := 0; i < int(totalHats); i++ {
			color := colors[i%len(colors)]
			third := size / float32(3)
			cross := []sdl.FRect{
				{X: x, Y: y + third, W: size, H: third},
				{X: x + third, Y: y, W: third, H: size},
			}

			sdl.SetRenderDrawColor(renderer, 90, 90, 90, 255)
			sdl.RenderFillRects(renderer, cross)

			hat := sdl.GetJoystickHat(joystick, int32(i))
			if hat&sdl.HatUp != 0 {
				rect := sdl.FRect{X: x + third, Y: y, W: third, H: third}
				sdl.SetRenderDrawColor(renderer, color.R, color.G, color.B, color.A)
				sdl.RenderFillRect(renderer, &rect)
			}
			if hat&sdl.HatRight != 0 {
				rect := sdl.FRect{X: x + 2*third, Y: y + third, W: third, H: third}
				sdl.SetRenderDrawColor(renderer, color.R, color.G, color.B, color.A)
				sdl.RenderFillRect(renderer, &rect)
			}
			if hat&sdl.HatDown != 0 {
				rect := sdl.FRect{X: x + third, Y: y + 2*third, W: third, H: third}
				sdl.SetRenderDrawColor(renderer, color.R, color.G, color.B, color.A)
				sdl.RenderFillRect(renderer, &rect)
			}
			if hat&sdl.HatLeft != 0 {
				rect := sdl.FRect{X: x, Y: y + third, W: third, H: third}
				sdl.SetRenderDrawColor(renderer, color.R, color.G, color.B, color.A)
				sdl.RenderFillRect(renderer, &rect)
			}
			x += size * 2
		}
	}

	textX := (int(winw) - len(text)*sdl.DebugTextFontCharacterSize) / 2
	textY := (winh - sdl.DebugTextFontCharacterSize) / 2
	sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
	sdl.RenderDebugText(renderer, float32(textX), float32(textY), text)
}

func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
