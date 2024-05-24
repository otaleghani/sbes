package terminalinput

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

var originalState *term.State

// EnableRawMode sets the terminal to raw mode.
func EnableRawMode() {
	var err error
	fd := int(os.Stdin.Fd())
	originalState, err = term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}

	// Handle interrupt signals to restore original terminal state
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		DisableRawMode()
		os.Exit(0)
	}()
}

// DisableRawMode restores the terminal to its original state.
func DisableRawMode() {
	fd := int(os.Stdin.Fd())
	term.Restore(fd, originalState)
}

// ReadInput reads input from the terminal with real-time backspace handling.
func ReadInput(prompt string) string {
	EnableRawMode()
	defer DisableRawMode()

	fmt.Print(prompt)

	var input []byte
	reader := bufio.NewReader(os.Stdin)

	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		if char == 3 { // Ctrl+C to exit
			fmt.Println("^C")
			os.Exit(0)
		} else if char == 13 { // Enter key
			fmt.Println()
			break
		} else if char == 127 { // Backspace key
			if len(input) > 0 {
				input = input[:len(input)-1]
				fmt.Print("\b \b")
			}
		} else {
			input = append(input, char)
			fmt.Print(string(char))
		}
	}

	fmt.Print("\r \r")

	return string(input)
}
