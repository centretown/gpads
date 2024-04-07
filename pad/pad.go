package pad

type Pad interface {
	IsGamepadAvailable(Joystick int32) bool
	GetGamepadName(Joystick int32) string
	IsGamepadButtonPressed(Joystick int32, button int32) bool
	IsGamepadButtonDown(Joystick int32, button int32) bool
	IsGamepadButtonReleased(Joystick int32, button int32) bool
	IsGamepadButtonUp(Joystick int32, button int32) bool
	GetGamepadButtonPressed() int32
	GetGamepadAxisCount(Joystick int32) int32
	GetGamepadAxisMovement(Joystick int32, axis int32) float32
	SetGamepadMappings(mappings string) int32
}

type PadG interface {
	Pad
	BeginPad()
	GetPadButtonCount(Joystick int32) int32
	GetPadCount() int32
	GetButtonName(Joystick int32, button int32) string
	GetPadAxisValue(Joystick int32, axis int32) int32
	DumpPad()
}
