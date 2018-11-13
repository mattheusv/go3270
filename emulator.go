package go3270

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// These constants represent the keyboard keys
const (
	Enter = "Enter"
	Tab   = "Tab"
	F1    = "PF(1)"
	F2    = "PF(2)"
	F3    = "PF(3)"
	F4    = "PF(4)"
	F5    = "PF(5)"
	F6    = "PF(6)"
	F7    = "PF(7)"
	F8    = "PF(8)"
	F9    = "PF(9)"
	F10   = "PF(10)"
	F11   = "PF(11)"
	F12   = "PF(12)"
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

//GetRows returns the number of rows in the saved screen image.
func (e *Emulator) GetRows() (int, error) {
	s, err := e.execCommandOutput("Snap(Rows)")
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error from x3270 to get numbers of row: %v", err)
	}
	return i, nil
}

//GetColumns returns the number of columns in the saved screen image.
func (e *Emulator) GetColumns() (int, error) {
	s, err := e.execCommandOutput("Snap(Cols)")
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error from x3270 to get numbers of columns: %v", err)
	}
	return i, nil
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
	s, err := e.query("ConnectionState")
	if err != nil || len(strings.TrimSpace(s)) == 0 {
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
	if !e.IsConnected() {
		e.execCommand("quit")
		log.Fatalf("error to connect in %s", e.hostname())
	}
	return nil
}

//Disconnect close connection with x3270
func (e *Emulator) Disconnect() error {
	if e.IsConnected() {
		return e.execCommand("quit")
	}
	return nil
}

//query returns state information from x3270
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
