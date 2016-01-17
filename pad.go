package gokicadlib

import (
	"fmt"
	"strconv"
)

// PadTypes ...
type PadTypes int

const (
	th PadTypes = iota
	smd
	connect
	npth
)

// PadShapes ...
type PadShapes int

const (
	circle PadShapes = iota
	rect
	oval
	trapezoid
)

// Pad ...
type Pad struct {
	Number           int
	PinName          string
	Origin           Point
	Style            string
	OriginalStyle    string
	OriginalNumber   int
	Shape            PadShapes
	Angle            int
	Size             Point
	Drillsize        string
	Layers           KicadLayerSlice
	SolderMaskMargin float64
	Padtype          PadTypes
}

// ToSExp convert to kicad s-expression
func (pad Pad) ToSExp() string {
	at := NewSExp("at", false, pad.Origin.ToString(), strconv.FormatInt(int64(pad.Angle), 10))
	size := NewSExp("size", false, pad.Size.ToString())
	drill := NewSExp("drill", false, pad.Drillsize)
	layers := NewSExp("layers", false, pad.Layers.String()...)
	soldermaskmargin := NewSExp("solder_mask_margin", false, fmt.Sprintf("%f", pad.SolderMaskMargin))

	return NewSExp("pad "+fmt.Sprint(pad.Number)+" "+pad.Padtype.String()+" "+pad.Shape.String(), true, at, size, drill, layers, soldermaskmargin)
}
