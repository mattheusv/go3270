package go3270

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// These constants represent the keyboard keys
const (
	Enter = "Enter"
	Tab   = "Tab"
)

//Emulator base struct to x3270 terminal emulator
type Emulator struct {
	Host       string
	Port       int
	ScriptPort string
}

//moveCursor move cursor to especific row(x) and column(y)
func (e *Emulator) moveCursor(x, y int) error {
	command := fmt.Sprintf("MoveCursor(%d,%d)", x, y)
	return e.execCommand(command)
}

//SetString fill field with value passed by parameter
//setString will fill the field that the cursor is marked
func (e *Emulator) SetString(value string) error {
	command := fmt.Sprintf("String(%s)", value)
	return e.execCommand(command)
}

//FillString fill the field by your position
//row(x), column(y) value(value to be filled)
func (e *Emulator) FillString(x, y int, value string) error {
	if err := e.moveCursor(x, y); err != nil {
		return fmt.Errorf("error to move cursor: %v", err)
	}
	return e.SetString(value)
}

//Press press a keyboard key
func (e *Emulator) Press(key string) error {
	if !e.validaKeyboard(key) {
		return fmt.Errorf("invalid key %s", key)
	}
	return e.execCommand(key)
}

//validaKeyboard valid if key passed by parameter if a key valid
func (e *Emulator) validaKeyboard(key string) bool {
	switch key {
	case Tab:
		return true
	case Enter:
		return true
	default:
		return false
	}
}

//IsConnected check if a connection with host exist
func (e *Emulator) IsConnected() bool {
	_, err := e.query("ConnectionState")
	if err != nil {
		return false
	}
	return true
}

//GetValue return a value by position
//row(x), column(y), lenght(l)
func (e *Emulator) GetValue(x, y, l int) (string, error) {
	command := fmt.Sprintf("Ascii(%d,%d,%d)", x, y, l)
	return e.execCommandOutput(command)
}

//CursorPosition return actual position by cursor
func (e *Emulator) CursorPosition() (string, error) {
	return e.query("cursor")
}

//Connect open an connection with x3270 and host
func (e *Emulator) Connect() error {
	if e.Host == "" {
		return errors.New("Host need to be filled")
	}
	if e.IsConnected() {
		return errors.New("addres already use")
	}
	if e.ScriptPort == "" {
		e.ScriptPort = "5000"
	}
	e.createApp()
	return nil
}

//Disconnect close connection with x3270
func (e *Emulator) Disconnect() error {
	if e.IsConnected() {
		return e.execCommand("quit")
	}
	return nil
}

func (e *Emulator) query(keyword string) (string, error) {
	command := fmt.Sprintf("query(%s)", keyword)
	return e.execCommandOutput(command)
}

//createApp create a connection of 3270 with host
func (e *Emulator) createApp() {
	cmd := exec.Command("x3270", "-scriptport", e.ScriptPort, e.hostname())
	go func() {
		if err := cmd.Run(); err != nil {
			log.Fatalf("error to create an instance of 3270\n%v\n", err)
		}
	}()
	time.Sleep(6 * time.Second)
}

//hostname return hostname formatted
func (e *Emulator) hostname() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

//execCommand executes a command on the connected x3270 instance
func (e *Emulator) execCommand(command string) error {
	cmd := exec.Command("x3270if", "-t", e.ScriptPort, command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	return nil
}

//execCommand executes a command on the connected x3270 instance and return output
func (e *Emulator) execCommandOutput(command string) (string, error) {
	cmd := exec.Command("x3270if", "-t", e.ScriptPort, command)
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
