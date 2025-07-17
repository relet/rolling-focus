package main

import (
	"fmt"
	"time"

	"go.bug.st/serial"
)

var serialMode = &serial.Mode{
	BaudRate: 9600,
	DataBits: 8,
	Parity:   serial.NoParity,
	StopBits: serial.OneStopBit,
}

const (
	AxisX   = "X"
	AxisY   = "Y"
	AxisZ   = "Z"
	DirUp   = "+"
	DirDown = "-"
)

func motorMove(port serial.Port, axis string, direction string) error {
	fmt.Println("Moving motor:", axis, direction)
	if port == nil {
		return nil // No port to send command to
	}

	_, err := port.Write([]byte(axis + direction + "\n"))
	if err != nil {
		return err
	}
	// sleep for a short duration to allow the movement to settle
	time.Sleep(100 * time.Millisecond)
	return nil
}
