package device

import (
	"fmt"
	"net"

	"github.com/0ptimusGrime/plugwork/message"

	"github.com/0ptimusGrime/plugwork/capability"
)

// UDPVibrator is a set device. It simply accepts instructions and writes them out to a udp connection
type UDPVibrator struct {
	config          UDPVibratorConfig
	conn            net.Conn
	implementations ImplementationSet
	name            string
}

// UDPVibratorConfig configures the UdpVibrator
// `Endpoint` an "ip:port" address for the device. e.g: "127.0.0.1:9999"
type UDPVibratorConfig struct {
	Endpoint string
}

// NewUDPVibrator returns a fully-initialized UdpVibrator ready for use
func NewUDPVibrator(conf UDPVibratorConfig) (*UDPVibrator, error) {
	vibrator := &UDPVibrator{
		config: conf,
		name:   "UDPVibrator",
	}

	addr, err := net.ResolveUDPAddr("udp4", conf.Endpoint)
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return nil, err
	}
	vibrator.conn = conn

	vibrator.implementations = StringifyImplementationSet(vibrator.Capabilities()...)
	return vibrator, nil
}

// Capabilities returns the capabilities supported by this device.
// This device only supports vibrate on/off, we assume the device tracks
// its own internal state, so sending "vibrate, vibrate, vibrate" will result
// in the device turning on, then off, then on.
func (vibrator *UDPVibrator) Capabilities() []string {
	return []string{capability.Vibrate}
}

// CapabilitySubset takes a list of requested capabilities and returns those which this device
// supports
func (vibrator *UDPVibrator) CapabilitySubset(requestedCaps ...string) []string {
	supported := []string{}
	for _, requestedCap := range requestedCaps {
		if _, found := vibrator.implementations[requestedCap]; found {
			supported = append(supported, requestedCap)
		}
	}
	return supported
}

// HasCapabilities returns true/false if all capabilities are supported by this device
func (vibrator *UDPVibrator) HasCapabilities(requestedCaps ...string) bool {
	for _, requestedCap := range requestedCaps {
		if _, found := vibrator.implementations[requestedCap]; !found {
			return false
		}
	}
	return true
}

// Name returns the vibrator name
func (vibrator *UDPVibrator) Name() string {
	return vibrator.name
}

// Send a message to the vibrator
func (vibrator *UDPVibrator) Send(m message.Instruction) {
	fmt.Printf("sending messsage %s to device %s", m.Type(), vibrator.Name())
	deviceSpecificMessage := vibrator.implementations[m.Type()](m)
	vibrator.conn.Write([]byte(deviceSpecificMessage.(string)))
}

// Stop closes the network connection to the vibrator
func (vibrator *UDPVibrator) Stop() {
	vibrator.conn.Close()
}
