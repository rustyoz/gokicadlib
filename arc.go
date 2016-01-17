package gokicadlib

import (
	"fmt"
	"math"
)

type Arc struct {
	Layer      Layer
	Origin     Point
	Radius     float64
	StartAngle float64
	SweepAngle float64
	Width      string
	End        Point
}

func (arc Arc) ToSExp() string {
	start := NewSExp("start", false, arc.Origin.ToString())
	arc.CalculateEndPoint()
	end := NewSExp("end", false, arc.End.ToString())
	angle := NewSExp("angle", false, fmt.Sprintf("%f", arc.StartAngle))
	layer := NewSExp("layer", false, arc.Layer.Name)
	width := NewSExp("width", false, arc.Width)
	return NewSExp("fp_line", true, start, end, angle, layer, width)
}

func (arc *Arc) CalculateEndPoint() {
	origin := arc.Origin.ToPointFloat()
	vector := PointFloat{arc.Radius * math.Cos(arc.StartAngle*180/math.Pi), arc.Radius * math.Sin(arc.StartAngle*180/math.Pi)}
	end := PointFloat{origin.X + vector.X, origin.Y + vector.Y}
	arc.End = Point{fmt.Sprintf("%f.4", end.X), fmt.Sprintf("%f.4", end.Y)}
}
