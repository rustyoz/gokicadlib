package gokicadlib

import "fmt"

// PadTypes ...
type PadTypes int

const (
	ThroughHole PadTypes = iota
	Surface
	Connect
	NotPlatedThroughHole
)

// PadShapes ...
type PadShapes int

const (
	Circle PadShapes = iota
	Rectangle
	Oval
	Trapezoid
)

type PadSlice []Pad

// Pad ...
type Pad struct {
	Number           int
	PinName          string
	Origin           Point
	Shape            PadShapes
	Size             Point
	Drillsize        float64
	Layers           KicadLayerSlice
	SolderMaskMargin float64
	Padtype          PadTypes
}

// ToSExp convert to kicad s-expression
func (pad Pad) ToSExp() string {
	at := NewSExp("at", false, pad.Origin.ToString())
	size := NewSExp("size", false, pad.Size.ToString())
	drill := NewSExp("drill", false, fmt.Sprint(pad.Drillsize))
	layers := NewSExp("layers", false, pad.Layers.String()...)
	soldermaskmargin := NewSExp("solder_mask_margin", false, fmt.Sprintf("%f", pad.SolderMaskMargin))

	return NewSExp("pad "+fmt.Sprint(pad.Number)+" "+pad.Padtype.String()+" "+pad.Shape.String(), true, at, size, drill, layers, soldermaskmargin)
}

func (pad *Pad) String() string {
	return pad.ToSExp()
}

func (ps PadSlice) ToSExp() string {
	var r string
	for _, p := range ps {
		r = r + fmt.Sprintln(p.ToSExp())
	}

	return r
}
