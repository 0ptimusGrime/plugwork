package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/0ptimusGrime/plugwork/capability"
	"github.com/0ptimusGrime/plugwork/message"
)

// Readable defines an interface which anything that can consume control signals should satisfy.
type Readable interface {
	Start()
	Stop()
}

// ConsoleReader will read text from the console, convert it into `messages.Instruction` types
// and emit them into a read queue
type ConsoleReader struct {
	queue chan<- message.Instruction
	done  chan bool
}

// NewConsoleReader will return a console reader configured to emit messages on the
// provided `messages.Instruction` channel
func NewConsoleReader(providedQueue chan<- message.Instruction) *ConsoleReader {
	return &ConsoleReader{
		queue: providedQueue,
		done:  make(chan bool, 1),
	}
}

// CreateInstruction creates a device.Instruction from the specified text input
func (c *ConsoleReader) CreateInstruction(text string) (message.Instruction, error) {
	parts := strings.Split(text, " ")
	cmd := parts[0]
	switch cmd {
	case capability.Vibrate:
		return message.VibrateInstruction{}, nil
	}
	return nil, fmt.Errorf("instruction type %s unsupported", cmd)
}

// Start starts the console reader. This should be wrapped in a gofunc
// TODO should probably use readrune and define a tokenizer or something
func (c *ConsoleReader) Start() {
	fmt.Println("Starting console reader. ctrl+c to exit")
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-c.done:
			return
		default:
			fmt.Print("> ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			instruction, err := c.CreateInstruction(text)
			if err != nil {
				fmt.Println(err)
				continue
			}
			c.queue <- instruction
		}
	}
}

// Stop will shut down the console reader
func (c *ConsoleReader) Stop() {
	c.done <- true
}
