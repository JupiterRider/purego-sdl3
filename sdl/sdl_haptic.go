package sdl

// [Haptic] is a structure used to identify an SDL haptic.
//
// [Haptic]: https://wiki.libsdl.org/SDL3/SDL_Haptic
type Haptic struct{}

const HapticInfinity = 4294967295 // Used to play a device an infinite number of times.

// [HapticEffectType] describes type of haptic effect.
//
// [HapticEffectType]: https://wiki.libsdl.org/SDL3/SDL_HapticEffectType
type HapticEffectType uint16

const (
	HapticConstant_    HapticEffectType = 1 << iota    // Constant effect supported.
	HapticSine                                         // Sine wave effect supported.
	HapticSquare                                       // Square wave effect supported.
	HapticTriangle                                     // Triangle wave effect supported.
	HapticSawtoothup                                   // Sawtoothup wave effect supported.
	HapticSawtoothdown                                 // Sawtoothdown wave effect supported.
	HapticRamp_                                        // Ramp effect supported.
	HapticSpring                                       // Spring effect supported - uses axes position.
	HapticDamper                                       // Damper effect supported - uses axes velocity.
	HapticInertia                                      // Inertia effect supported - uses axes acceleration.
	HapticFriction                                     // Friction effect supported - uses axes movement.
	HapticLeftright                                    // Left/Right effect supported.
	HapticReserved1                                    // Reserved for future use.
	HapticReserved2                                    // Reserved for future use.
	HapticReserved3                                    // Reserved for future use.
	HapticCustom_                                      // Custom effect is supported.
	HapticGain         uint32           = 1<<iota + 16 // Device can set global gain.
	HapticAutocenter                                   // Device can set autocenter.
	HapticStatus                                       // Device can be queried for effect status.
	HapticPause                                        // Device can be paused.
)

// [HapticDirectionType] describes type of coordinates used for haptic direction.
//
// [HapticDirectionType]: https://wiki.libsdl.org/SDL3/SDL_HapticDirectionType
type HapticDirectionType uint8

const (
	HapticPolar        HapticDirectionType = 0 // Uses polar coordinates for the direction.
	HapticCartesian    HapticDirectionType = 1 // Uses cartesian coordinates for the direction.
	HapticSpherical    HapticDirectionType = 2 // Uses spherical coordinates for the direction.
	HapticSteeringAxis HapticDirectionType = 3 // Use this value to play an effect on the steering wheel axis.
)

// [HapticEffectID] describes ID for haptic effects.
//
// [HapticEffectID]: https://wiki.libsdl.org/SDL3/SDL_HapticEffectID
type HapticEffectID int32

// [HapticDirection] is a structure that represents a haptic direction.
//
// [HapticDirection]: https://wiki.libsdl.org/SDL3/SDL_HapticDirection
type HapticDirection struct {
	Type HapticDirectionType // The type of encoding.
	Dir  [3]int32            // The encoded direction.
}

// [HapticConstant] is a structure containing a template for a Constant effect.
//
// [HapticConstant]: https://wiki.libsdl.org/SDL3/SDL_HapticConstant
type HapticConstant struct {
	Type         HapticEffectType // [HapticConstant_].
	Direction    HapticDirection  // Direction of the effect.
	Length       uint32           // Duration of the effect.
	Delay        uint16           // Delay before starting the effect.
	Button       uint16           // Button that triggers the effect.
	Interval     uint16           // How soon it can be triggered again after button.
	Level        int16            // Strength of the constant effect.
	AttackLength uint16           // Duration of the attack.
	AttackLevel  uint16           // Level at the start of the attack.
	FadeLength   uint16           // Duration of the fade.
	FadeLevel    uint16           // Level at the end of the fade.
}

// [HapticPeriodic] is a structure containing a template for a Periodic effect.
//
// [HapticPeriodic]: https://wiki.libsdl.org/SDL3/SDL_HapticPeriodic
type HapticPeriodic struct {
	Type         HapticEffectType // [HapticSine], [HapticSquare], [HapticTriangle], [HapticSawtoothup] or [HapticSawtoothdown].
	Direction    HapticDirection  // Direction of the effect.
	Length       uint32           // Duration of the effect.
	Delay        uint16           // Delay before starting the effect.
	Button       uint16           // Button that triggers the effect.
	Interval     uint16           // How soon it can be triggered again after button.
	Period       uint16           // Period of the wave.
	Magnitude    int16            // Peak value; if negative, equivalent to 180 degrees extra phase shift.
	Offset       int16            // Mean value of the wave.
	Phase        uint16           // Positive phase shift given by hundredth of a degree.
	AttackLength uint16           // Duration of the attack.
	AttackLevel  uint16           // Level at the start of the attack.
	FadeLength   uint16           // Duration of the fade.
	FadeLevel    uint16           // Level at the end of the fade.
}

// [HapticCondition] is a structure containing a template for a Condition effect.
//
// [HapticCondition]: https://wiki.libsdl.org/SDL3/SDL_HapticCondition
type HapticCondition struct {
	Type       HapticEffectType // [HapticSpring], [HapticDamper], [HapticInertia] or [HapticFriction].
	Direction  HapticDirection  // Direction of the effect.
	Length     uint32           // Duration of the effect.
	Delay      uint16           // Delay before starting the effect.
	Button     uint16           // Button that triggers the effect.
	Interval   uint16           // How soon it can be triggered again after button.
	RightSat   [3]uint16        // Level when joystick is to the positive side; max 0xFFFF.
	LeftSat    [3]uint16        // Level when joystick is to the negative side; max 0xFFFF.
	RightCoeff [3]int16         // How fast to increase the force towards the positive side.
	LeftCoeff  [3]int16         // How fast to increase the force towards the negative side.
	Deadband   [3]uint16        // Size of the dead zone; max 0xFFFF: whole axis-range when 0-centered.
	Center     [3]int16         // Position of the dead zone.
}

// [HapticRamp] is a structure containing a template for a Ramp effect.
//
// [HapticRamp]: https://wiki.libsdl.org/SDL3/SDL_HapticRamp
type HapticRamp struct {
	Type         HapticEffectType // [HapticRamp_].
	Direction    HapticDirection  // Direction of the effect.
	Length       uint32           // Duration of the effect.
	Delay        uint16           // Delay before starting the effect.
	Button       uint16           // Button that triggers the effect.
	Interval     uint16           // How soon it can be triggered again after button.
	Start        int16            // Beginning strength level.
	End          int16            // Ending strength level.
	AttackLength uint16           // Duration of the attack.
	AttackLevel  uint16           // Level at the start of the attack.
	FadeLength   uint16           // Duration of the fade.
	FadeLevel    uint16           // Level at the end of the fade.
}

// [HapticLeftRight] is a structure containing a template for a Left/Right effect.
//
// [HapticLeftRight]: https://wiki.libsdl.org/SDL3/SDL_HapticLeftRight
type HapticLeftRight struct {
	Type           HapticEffectType // [HapticLeftRight].
	Length         uint32           // Duration of the effect in milliseconds.
	LargeMagnitude uint16           // Control of the large controller motor.
	SmallMagnitude uint16           // Control of the small controller motor.
}

// [HapticCustom] is a structure containing a template for the SDL_HAPTIC_CUSTOM effect.
//
// [HapticCustom]: https://wiki.libsdl.org/SDL3/SDL_HapticCustom
type HapticCustom struct {
	Type         HapticEffectType // [HapticCustom_].
	Direction    HapticDirection  // Direction of the effect.
	Length       uint32           // Duration of the effect.
	Delay        uint16           // Delay before starting the effect.
	Button       uint16           // Button that triggers the effect.
	Interval     uint16           // How soon it can be triggered again after button.
	Channels     uint8            // Axes to use, minimum of one.
	Period       uint16           // Sample periods.
	Samples      uint16           // Amount of samples.
	Data         *uint16          // Should contain channels*samples items.
	AttackLength uint16           // Duration of the attack.
	AttackLevel  uint16           // Level at the start of the attack.
	FadeLength   uint16           // Duration of the fade.
	FadeLevel    uint16           // Level at the end of the fade.
}

// [HapticEffect] is a structure describing generic template for any haptic effect.
//
// [HapticEffect]: https://wiki.libsdl.org/SDL3/SDL_HapticEffect
type HapticEffect struct {
	Type      HapticEffectType // Effect type.
	Constant  HapticConstant   // Constant effect.
	Periodic  HapticPeriodic   // Periodic effect.
	Condition HapticCondition  // Condition effect.
	Ramp      HapticRamp       // Ramp effect.
	Leftright HapticLeftRight  // Left/Right effect.
	Custom    HapticCustom     // Custom effect.
}

// [HapticID] is a unique ID for a haptic device for the time it is connected to the system, and is never reused for the lifetime of the application.
//
// [HapticID]: https://wiki.libsdl.org/SDL3/SDL_HapticID
type HapticID uint32

// func CloseHaptic(haptic *Haptic)  {
//	sdlCloseHaptic(haptic)
// }

// func CreateHapticEffect(haptic *Haptic, effect *HapticEffect) int32 {
//	return sdlCreateHapticEffect(haptic, effect)
// }

// func DestroyHapticEffect(haptic *Haptic, effect int32)  {
//	sdlDestroyHapticEffect(haptic, effect)
// }

// func GetHapticEffectStatus(haptic *Haptic, effect int32) bool {
//	return sdlGetHapticEffectStatus(haptic, effect)
// }

// func GetHapticFeatures(haptic *Haptic) uint32 {
//	return sdlGetHapticFeatures(haptic)
// }

// func GetHapticFromID(instance_id HapticID) *Haptic {
//	return sdlGetHapticFromID(instance_id)
// }

// func GetHapticID(haptic *Haptic) HapticID {
//	return sdlGetHapticID(haptic)
// }

// func GetHapticName(haptic *Haptic) string {
//	return sdlGetHapticName(haptic)
// }

// func GetHapticNameForID(instance_id HapticID) string {
//	return sdlGetHapticNameForID(instance_id)
// }

// func GetHaptics(count *int32) *HapticID {
//	return sdlGetHaptics(count)
// }

// func GetMaxHapticEffects(haptic *Haptic) int32 {
//	return sdlGetMaxHapticEffects(haptic)
// }

// func GetMaxHapticEffectsPlaying(haptic *Haptic) int32 {
//	return sdlGetMaxHapticEffectsPlaying(haptic)
// }

// func GetNumHapticAxes(haptic *Haptic) int32 {
//	return sdlGetNumHapticAxes(haptic)
// }

// func HapticEffectSupported(haptic *Haptic, effect *HapticEffect) bool {
//	return sdlHapticEffectSupported(haptic, effect)
// }

// func HapticRumbleSupported(haptic *Haptic) bool {
//	return sdlHapticRumbleSupported(haptic)
// }

// func InitHapticRumble(haptic *Haptic) bool {
//	return sdlInitHapticRumble(haptic)
// }

// func IsJoystickHaptic(joystick *Joystick) bool {
//	return sdlIsJoystickHaptic(joystick)
// }

// func IsMouseHaptic() bool {
//	return sdlIsMouseHaptic()
// }

// func OpenHaptic(instance_id HapticID) *Haptic {
//	return sdlOpenHaptic(instance_id)
// }

// func OpenHapticFromJoystick(joystick *Joystick) *Haptic {
//	return sdlOpenHapticFromJoystick(joystick)
// }

// func OpenHapticFromMouse() *Haptic {
//	return sdlOpenHapticFromMouse()
// }

// func PauseHaptic(haptic *Haptic) bool {
//	return sdlPauseHaptic(haptic)
// }

// func PlayHapticRumble(haptic *Haptic, strength float32, length uint32) bool {
//	return sdlPlayHapticRumble(haptic, strength, length)
// }

// func ResumeHaptic(haptic *Haptic) bool {
//	return sdlResumeHaptic(haptic)
// }

// func RunHapticEffect(haptic *Haptic, effect int32, iterations uint32) bool {
//	return sdlRunHapticEffect(haptic, effect, iterations)
// }

// func SetHapticAutocenter(haptic *Haptic, autocenter int32) bool {
//	return sdlSetHapticAutocenter(haptic, autocenter)
// }

// func SetHapticGain(haptic *Haptic, gain int32) bool {
//	return sdlSetHapticGain(haptic, gain)
// }

// func StopHapticEffect(haptic *Haptic, effect int32) bool {
//	return sdlStopHapticEffect(haptic, effect)
// }

// func StopHapticEffects(haptic *Haptic) bool {
//	return sdlStopHapticEffects(haptic)
// }

// func StopHapticRumble(haptic *Haptic) bool {
//	return sdlStopHapticRumble(haptic)
// }

// func UpdateHapticEffect(haptic *Haptic, effect int32, data *HapticEffect) bool {
//	return sdlUpdateHapticEffect(haptic, effect, data)
// }
