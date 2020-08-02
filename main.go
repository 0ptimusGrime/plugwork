package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/0ptimusGrime/plugwork/device"
	"github.com/0ptimusGrime/plugwork/input"
	"github.com/0ptimusGrime/plugwork/message"
)

func main() {
	readQueue := make(chan message.Instruction, 1)
	readers := []input.Readable{
		input.NewConsoleReader(readQueue),
	}
	for _, reader := range readers {
		go reader.Start()
	}

	deviceSet := device.Set{}
	udpVibrator, err := device.NewUDPVibrator(device.UDPVibratorConfig{
		Endpoint: "127.0.0.1:9999",
	})
	if err != nil {
		fmt.Printf("Couldn't register udp viberator: %s", err)
		os.Exit(1)
	}
	deviceSet.Register(udpVibrator)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
workLoop:
	for {
		select {
		case m := <-readQueue:
			fmt.Printf("sending message [%s] to %d devices...\n", m.Type(), deviceSet.Len())
			deviceSet.Send(m)
		case <-c:
			deviceSet.Stop()
			break workLoop
		}
	}
	fmt.Println("\r- Ctrl+C captured. Shutting down...")
	for _, reader := range readers {
		reader.Stop()
	}
	os.Exit(0)
}
