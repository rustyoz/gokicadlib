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
	layer := NewSExp("layer", false, string(arc.Layer))
	width := NewSExp("width", false, arc.Width)
	return NewSExp("fp_line", true, start, end, angle, layer, width)
}

func (arc *Arc) CalculateEndPoint() {
	vector := Point{arc.Radius * math.Cos(arc.StartAngle*180/math.Pi), arc.Radius * math.Sin(arc.StartAngle*180/math.Pi)}
	arc.End = Point{arc.Origin.X + vector.X, arc.Origin.Y + vector.Y}

}
