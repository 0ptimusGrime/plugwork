package message

import (
	"fmt"

	"github.com/0ptimusGrime/plugwork/capability"
)

// Instruction defines a generic interface for representing device instructions.
// Specific types of instructions simply need to satisfy this interface.
type Instruction interface {
	HasPayload() bool
	Payload() interface{}
	Type() string
}

// VibrateInstruction will instruct a device to turn its vibration function on/off
// we assume the actual device keeps track of its own state, so "vibrate vibrate vibrate"
// will turn the device on/off/on
type VibrateInstruction struct{}

// Type returns the capability type
func (v VibrateInstruction) Type() string {
	return capability.Vibrate
}

// HasPayload returns false; we don't have payloads for this instruction type
func (v VibrateInstruction) HasPayload() bool {
	return false
}

// Payload returns nil for this instruction
func (v VibrateInstruction) Payload() interface{} {
	fmt.Println("this instruction doesn't have payloads. how did you get here")
	return nil
}
