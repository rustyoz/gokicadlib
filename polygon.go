package gokicadlib

import "fmt"

type Polygon struct {
	Layer  Layer
	Origin Point
	End    Point
	Points []Point
	Width  float64
}

func (p *Polygon) ToSExp() string {
	start := NewSExp("start", false, p.Origin.ToString())
	end := NewSExp("end", false, p.End.ToString())
	if len(p.Layer) == 0 {
		p.Layer = F_SilkS
	}
	layer := NewSExp("layer", false, string(p.Layer))
	if p.Width == 0 {
		p.Width = DEFAULTLINEWIDTH / 25.4
	}
	width := NewSExp("width", false, fmt.Sprintf("%.4f", p.Width))
	return NewSExp("polygon", false, start, end, layer, width)
}
