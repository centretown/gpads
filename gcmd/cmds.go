package gcmd

import (
	"fmt"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/try"
)

func LastButtonPressed(cmd *GCmd) {
	button := js.GetGamepadButtonPressed()
	fmt.Printf("[%2d]\r", button)
}

func IsButtonUp(cmd *GCmd) {
	up := js.IsGamepadButtonUp(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, try.As[int](up))
}

func IsButtonDown(cmd *GCmd) {
	down := js.IsGamepadButtonDown(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, try.As[int](down))
}

func IsButtonReleased(cmd *GCmd) {
	released := js.IsGamepadButtonReleased(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, try.As[int](released))
}

func IsButtonPressed(cmd *GCmd) {
	pressed := js.IsGamepadButtonPressed(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, try.As[int](pressed))
}

func GetAxisValues(cmd *GCmd) {
	count := js.GetGamepadAxisCount(cmd.Pad)
	fmt.Print("axes:  ")
	for i := range count {
		value := js.GetPadAxisValue(cmd.Pad, i)
		fmt.Printf("[%d:%6d] ", i, value)
	}
	fmt.Print("\r")
}

func GetAxisMovement(cmd *GCmd) {
	fmt.Print("axes:  ")
	count := js.GetGamepadAxisCount(cmd.Pad)
	for i := range count {
		move := js.GetGamepadAxisMovement(cmd.Pad, i)
		fmt.Printf("[%d:%6.0f] ", i, move)
	}
	fmt.Print("\r")
}

func DumpPad(cmd *GCmd) {
	js.DumpPad()
}

const MAX_BUTTONS = 64

func TestKeys(cmd *GCmd) {
	count := gpads.RL_RightThumb - gpads.RL_Unknown + 1
	for i := range count {
		down := js.IsGamepadButtonDown(cmd.Pad, i)
		fmt.Printf("[%x:%2d]", i, try.As[int](down))
	}
	fmt.Print("\r")
}

func TestAxes(cmd *GCmd) {
	count := js.GetGamepadAxisCount(cmd.Pad)
	for i := range count {
		val := js.GetPadAxisValue(cmd.Pad, i)
		mov := js.GetGamepadAxisMovement(cmd.Pad, i)
		fmt.Printf("[%x:%03x:%3.0f]", i, val, mov)
	}
	fmt.Print("  \r")
}
