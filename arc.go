package gokicadlib

import (
	"fmt"
	"math"
	"strings"
)

type Arc struct {
	Layer Layer
	Start Point
	End   Point
	Angle float64
	Width float64
}

type ArcSlice []Arc

func (arc Arc) ToSExp() string {
	start := NewSExp("start", false, arc.Start.ToString())
	end := NewSExp("end", false, arc.End.ToString())
	angle := NewSExp("angle", false, fmt.Sprintf("%d", int(arc.Angle)))
	layer := NewSExp("layer", false, string(arc.Layer))
	if arc.Width == 0 {
		arc.Width = DEFAULTLINEWIDTH / 25.4
	}
	width := NewSExp("width", false, fmt.Sprint(arc.Width))
	return NewSExp("fp_arc", false, start, end, angle, layer, width)
}

func Vector(p Point, radius float64, angle float64) Point {
	vector := Point{radius * math.Cos(angle*180/math.Pi), radius * math.Sin(angle*180/math.Pi), 0}
	return vector
}

func (as ArcSlice) ToSExp() string {
	var r []string
	for _, a := range as {
		r = append(r, a.ToSExp())
	}
	return strings.Join(r, "\n")
}
