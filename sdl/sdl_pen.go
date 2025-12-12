package sdl

// [PenAxis] defines the pen axis indices.
//
// [PenAxis]: https://wiki.libsdl.org/SDL3/SDL_PenAxis
type PenAxis uint32

const (
	PenAxisPressure           PenAxis = iota // Pen pressure. Unidirectional: 0 to 1.0.
	PenAxisXTilt                             // Pen horizontal tilt angle.  Bidirectional: -90.0 to 90.0 (left-to-right).
	PenAxisYTilt                             // Pen vertical tilt angle.  Bidirectional: -90.0 to 90.0 (top-to-down).
	PenAxisDistance                          // Pen distance to drawing surface. Unidirectional: 0.0 to 1.0.
	PenAxisRotation                          // Pen barrel rotation. Bidirectional: -180 to 179.9 (clockwise, 0 is facing up, -180.0 is facing down).
	PenAxisSlider                            // Pen finger wheel or slider (e.g., Airbrush Pen). Unidirectional: 0 to 1.0.
	PenAxisTangentialPressure                // Pressure from squeezing the pen ("barrel pressure").
	PenAxisCount                             // Total known pen axis types in this version of SDL. This number may grow in future releases!.
)

// [PenID] defines the SDL pen instance IDs.
//
// [PenID]: https://wiki.libsdl.org/SDL3/SDL_PenID
type PenID uint32

// [PenInputFlags] defines the pen input flags, as reported by various pen events' pen_state field.
//
// [PenInputFlags]: https://wiki.libsdl.org/SDL3/SDL_PenInputFlags
type PenInputFlags uint32

const (
	PenInputDown      PenInputFlags = 1 << iota // Pen is pressed down.
	PenInputButton1                             // Button 1 is pressed.
	PenInputButton2                             // Button 2 is pressed.
	PenInputButton3                             // Button 3 is pressed.
	PenInputButton4                             // Button 4 is pressed.
	PenInputButton5                             // Button 5 is pressed.
	PenInputEraserTip PenInputFlags = 1 << 30   // Eraser tip is used.
)
