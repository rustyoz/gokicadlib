package gokicadlib

import (
	"bytes"
	"fmt"
)

// PinTypes ...
type PinTypes int

const (
	Input PinTypes = iota
	Output
)

// PinShapes ...
type PinShapes int

const (
	Normal PinShapes = iota
	Inverted
	Clock
	Lowinput
	Outputlow
	Invisible
)

type PinSlice []Pin

// Pin ...
type Pin struct {
	Number   int
	PinName  string
	Origin   Point
	Shape    PinShapes
	Length   float64
	Pintype  PinTypes
	DeMorgan int
	Sizenum  float64
	Sizename float64
	Part     int
}

func (p *Pin) KicadLib() *bytes.Buffer {
	output := &bytes.Buffer{}
	var l string
	// symbol definition
	var orientation string
	var shape string
	var pintype string
	l = fmt.Sprint("X %s %d %f %f %f %s %d %d %d %d %s %s",
		p.PinName,
		p.Number,
		p.Origin.X, p.Origin.Y, p.Length, orientation, p.Sizenum, p.Sizename, p.Part, p.DeMorgan, pintype, shape)

	output.WriteString(l + "\r\n")

	return output
}
