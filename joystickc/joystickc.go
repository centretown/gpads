package joystickc

/*
#include "joystick.h"
*/
import "C"

import "github.com/centretown/gpads/pad"

var _ pad.Pad = NewJoyStickC()

type JoystickC struct {
}

func NewJoyStickC() *JoystickC {
	js := &JoystickC{}
	return js
}

func (js *JoystickC) BeginPad() {
	C.BeginJoystick()
}

func (js *JoystickC) IsGamepadAvailable(Joystick int32) bool {
	return bool(C.IsJoystickAvailable(C.int(Joystick)))
}

func (js *JoystickC) GetGamepadName(Joystick int32) string {
	return C.GoString(C.GetJoystickName(C.int(Joystick)))
}

func (js *JoystickC) GetPadCount() int32 {
	return int32(C.GetJoystickCount())
}

func (js *JoystickC) IsGamepadButtonPressed(Joystick int32, button int32) bool {
	return bool(C.IsJoystickButtonPressed(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsGamepadButtonDown(Joystick int32, button int32) bool {
	return bool(C.IsJoystickButtonDown(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsGamepadButtonReleased(Joystick int32, button int32) bool {
	return bool(C.IsJoystickButtonReleased(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsGamepadButtonUp(Joystick int32, button int32) bool {
	return bool(C.IsJoystickButtonUp(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) GetGamepadButtonPressed() int32 {
	return int32(C.GetJoystickButtonPressed())
}

func (js *JoystickC) GetGamepadAxisCount(Joystick int32) int32 {
	return int32(C.GetJoystickAxisCount(C.int(Joystick)))
}

func (js *JoystickC) GetPadButtonCount(Joystick int32) int32 {
	return int32(C.GetJoystickButtonCount(C.int(Joystick)))
}

func (js *JoystickC) GetGamepadAxisMovement(Joystick int32, axis int32) float32 {
	return float32(C.GetJoystickAxisMovement(C.int(Joystick), C.int(axis)))
}

func (js *JoystickC) GetPadAxisValue(Joystick int32, axis int32) int32 {
	return int32(C.GetJoystickAxisValue(C.int(Joystick), C.int(axis)))
}

func (js *JoystickC) SetGamepadMappings(mappings string) int32 {
	return 0
}

func (js *JoystickC) DumpPad() {
	C.DumpJoystick()
}

func (js *JoystickC) GetButtonName(Joystick int32, button int32) string {
	return C.GoString(C.GetButtonName(C.int(Joystick), C.int(button)))
}
