package capability

const (
	// Vibrate device can vibrate
	Vibrate = "vibrate"

	// VibrateParametric device can vibrate with optional modes/speeds
	VibrateParametric = "vibrate_parametric"

	// LinearActuation device support autonomous linear actuation
	LinearActuation = "linear_actuation"

	// LinearActuationParametric devices support linear actuation with differnt modes
	LinearActuationParametric = "linear_actuation_parametric"
)

// Supported returns all capabilities supported by this library
func Supported() map[string]bool {
	return map[string]bool{
		Vibrate:                   true,
		VibrateParametric:         true,
		LinearActuation:           true,
		LinearActuationParametric: true,
	}
}
