package abstractfactory

import "fmt"

type Mouse interface {
	Click()
}

type Keyboard interface {
	Type()
}

type ComputerHardwareFactory interface {
	CreateMouse() Mouse
	CreateKeyboard() Keyboard
}

type LogitechMouse struct{}

func (lm *LogitechMouse) Click() {
	fmt.Println("LogitechMouse Click")
}

type LogitechKeyboard struct{}

func (lm *LogitechKeyboard) Type() {
	fmt.Println(" LogitechKeyboard Type")
}

type LogitechFactory struct {
}

func (l *LogitechFactory) CreateMouse() Mouse {
	return &LogitechMouse{}
}

func (l *LogitechFactory) CreateKeyboard() Keyboard {
	return &LogitechKeyboard{}
}
