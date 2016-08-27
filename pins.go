package gokicadlib

import "fmt"

// PinTypes ...
type PinTypes string

const (
	Input         PinTypes = "I"
	Output                 = "O"
	Bidirectional          = "B"
	Tristate               = "T"
	Passive                = "P"
	OpenCollector          = "O"
	OpenEmitter            = "E"
	NonConnected           = "N"
	Unspecificied          = "U"
	PowerInput             = "W"
	PowerOutput            = "w"
)

// PinShapes ...
type PinShapes string

const (
	Normal    PinShapes = ""
	Inverted            = "I"
	Clock               = "C"
	Lowinput            = "L"
	Outputlow           = "V"
	Invisible           = "N"
)

type PinSlice []Pin

type PinOrientation string

const (
	Up    PinOrientation = "U"
	Down                 = "D"
	Left                 = "L"
	Right                = "R"
)

// Pin ...
type Pin struct {
	Number      int
	PinName     string
	Origin      Point
	Shape       PinShapes
	Length      float64
	Type        PinTypes
	DeMorgan    int
	Sizenum     float64
	Sizename    float64
	Part        int
	Orientation PinOrientation
}

func (p *Pin) KicadLib() string {
	var l string
	// symbol definition
	if len(p.Type) == 0 {
		p.Type = Unspecificied
	}
	l = fmt.Sprintf("X %s %d %.0f %.0f %.0f %s %.0f %.0f %d %d %s %s \r\n",
		p.PinName,
		p.Number,

		p.Origin.X, p.Origin.Y, p.Length, p.Orientation, p.Sizenum, p.Sizename, p.Part, p.DeMorgan, p.Type, p.Shape)

	return l
}
