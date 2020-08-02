package device

import (
	"fmt"

	"github.com/0ptimusGrime/plugwork/message"
)

// CapabilityImplementation implements an instruction
type CapabilityImplementation func(message.Instruction) interface{}

// ImplementationSet maps a feature to a function which can implement it for this device
type ImplementationSet map[string]CapabilityImplementation

// Generic provides a generic interface for interacting with devices
type Generic interface {
	// Capabilities returns the capabilities supported by this device
	Capabilities() []string

	// HasCapabilities returns `bool` if a device supports all the specified capabilities
	HasCapabilities(...string) bool

	// CapabilitySubset takes a slice of 'requested' capabilities and returns the subset supported by
	// this device
	CapabilitySubset(...string) []string

	// Name just returns the name of the device
	Name() string

	// Send an instruction to the device
	Send(message.Instruction)

	// Stop cleans up whatever resources we might have opened when connecting to the device
	Stop()
}

// Set is a container for an arbitrary number of devices
type Set struct {
	devices []Generic
}

// Send will pass the `message.Instruction` along to any contained devices which support the
// instruction type. This could be heavily optimized.
func (s *Set) Send(m message.Instruction) {
	for _, device := range s.devices {
		if device.HasCapabilities(m.Type()) {
			device.Send(m)
		}
	}
}

// Stop sends a shutdown signal to every registered device
func (s *Set) Stop() {
	for _, device := range s.devices {
		device.Stop()
	}
}

// Register adds a device to the device `Set`
func (s *Set) Register(d Generic) {
	s.devices = append(s.devices, d)
}

// Len returns the number of devices in the device set
func (s *Set) Len() int {
	return len(s.devices)
}

// StringifyImplementationSet will bind the PrintOp implementatin to every specified capabilitiy
func StringifyImplementationSet(capabilities ...string) ImplementationSet {
	implementations := make(ImplementationSet)
	for _, cap := range capabilities {
		implementations[cap] = StringOp
	}
	return implementations
}

// StringOp simply returns the type of the received message as a string
func StringOp(i message.Instruction) interface{} {
	return fmt.Sprintf("this is a [%s] message", i.Type())
}
