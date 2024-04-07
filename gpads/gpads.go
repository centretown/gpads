package gpads

import (
	"fmt"

	// "github.com/centretown/gpads/pad"
	"github.com/centretown/gpads/pad"
	"github.com/holoplot/go-evdev"
)

var _ pad.Pad = NewGPads()

const PAD_MAX = 4
const UNDEFINED = "UNDEFINED"

func badPad(pad int32, padCount int32) bool {
	return pad < 0 || pad >= padCount
}

func badAxis(axis int32) bool {
	return axis < 0 || axis >= RL_AXIS_COUNT
}

func badButton(button int32) bool {
	return button < 0 || button >= RL_BUTTON_COUNT
}

type GPads struct {
	Pads       [PAD_MAX]*GPad
	padCount   int32
	intialized bool
}

func NewGPads() *GPads {
	js := &GPads{}
	js.intialize()
	return js
}

func (js *GPads) intialize() {
	var (
		devicePaths []evdev.InputPath
		err         error
	)

	defer func() { js.intialized = true }()

	devicePaths, err = evdev.ListDevicePaths()
	if err != nil {
		fmt.Println("ListDevicePaths", err)
		devicePaths = make([]evdev.InputPath, 0)
	}

	js.padCount = 0
	for _, p := range devicePaths {
		stg, err := OpenGPad(p.Path)
		if err != nil {
			fmt.Printf("Open %s [%s] %v\n", p.Name, p.Path, err)
			continue
		}
		js.Pads[js.padCount] = stg
		js.padCount++
	}
}

func (js *GPads) BeginPad() {
	for i := range js.padCount {
		js.Pads[i].ReadState()
	}
}

func (js *GPads) IsGamepadAvailable(pad int32) bool {
	return !badPad(pad, js.padCount)
}

func (js *GPads) GetPadCount() int32 {
	return js.padCount
}

func (js *GPads) GetGamepadName(pad int32) string {
	if badPad(pad, js.padCount) {
		return UNDEFINED
	}
	return js.Pads[pad].Name
}

func (js *GPads) GetButtonName(pad int32, button int32) string {
	if badPad(pad, js.padCount) || badButton(button) {
		return UNDEFINED
	}
	return evdev.KEYNames[PS3Buttons[button]]
}

func (js *GPads) IsGamepadButtonPressed(pad int32, button int32) bool {
	if badPad(pad, js.padCount) || badButton(button) {
		return false
	}
	return js.Pads[pad].ButtonPressed(button)
}

func (js *GPads) IsGamepadButtonDown(pad int32, button int32) bool {
	if badPad(pad, js.padCount) || badButton(button) {
		return false
	}
	return js.Pads[pad].ButtonDown(button)
}

func (js *GPads) IsGamepadButtonReleased(pad int32, button int32) bool {
	if badPad(pad, js.padCount) || badButton(button) {
		return false
	}
	return js.Pads[pad].ButtonReleased(button)
}

func (js *GPads) IsGamepadButtonUp(pad int32, button int32) bool {
	if badPad(pad, js.padCount) || badButton(button) {
		return false
	}
	return !js.Pads[pad].ButtonDown(button)
}

func (js *GPads) GetGamepadButtonPressed() int32 {
	return LastPressed
}

func (js *GPads) GetGamepadAxisCount(pad int32) int32 {
	if js.padCount <= pad {
		return 0
	}
	return int32(len(js.Pads[pad].AxisState))
}

func (js *GPads) GetPadButtonCount(pad int32) int32 {
	if badPad(pad, js.padCount) {
		return 0
	}
	return int32(len(js.Pads[pad].ButtonState))
}

func (js *GPads) GetGamepadAxisMovement(pad int32, axis int32) float32 {
	if badPad(pad, js.padCount) || badAxis(axis) {
		return 0
	}
	return js.Pads[pad].AxisMove(axis)
}

func (js *GPads) GetPadAxisValue(pad int32, axis int32) int32 {
	if badPad(pad, js.padCount) || badAxis(axis) {
		return 0
	}
	return js.Pads[pad].AxisValue(axis)
}

func (js *GPads) SetGamepadMappings(mappings string) int32 {
	return 0
}

func (js *GPads) DumpPad() {
	for i := int32(0); i < js.padCount; i++ {
		js.Pads[i].Dump()
		fmt.Println()
	}
}

func (js *GPads) Close() {
	for i := int32(0); i < js.padCount; i++ {
		js.Pads[i].Close()
	}
}
