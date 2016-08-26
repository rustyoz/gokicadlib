package gokicadlib

import "fmt"

type Polygon struct {
	Layer  Layer
	Points []Point
	Width  float64
}

func (p *Polygon) ToSExp() string {
	var points string
	for _, s := range p.Points {
		points = points + s.ToSExp()
	}
	pts := NewSExp("pts", true, points)
	if len(p.Layer) == 0 {
		p.Layer = F_SilkS
	}
	layer := NewSExp("layer", false, string(p.Layer))
	if p.Width == 0 {
		p.Width = DEFAULTLINEWIDTH / 25.4
	}
	width := NewSExp("width", false, fmt.Sprintf("%.4f", p.Width))
	return NewSExp("fp_poly", false, pts, layer, width)
}
