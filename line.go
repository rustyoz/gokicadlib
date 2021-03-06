package gokicadlib

import (
	"fmt"
	"strings"
)

type Line struct {
	Origin Point
	End    Point
	Layer  Layer
	Width  float64
}

func (line *Line) ToSExp() string {
	start := NewSExp("start", false, line.Origin.ToString())
	end := NewSExp("end", false, line.End.ToString())
	if len(line.Layer) == 0 {
		line.Layer = F_SilkS
	}
	layer := NewSExp("layer", false, string(line.Layer))
	if line.Width == 0 {
		line.Width = DEFAULTLINEWIDTH / 25.4
	}
	width := NewSExp("width", false, fmt.Sprintf("%.4f", line.Width))
	return NewSExp("fp_line", false, start, end, layer, width)
}

type LineSlice []Line

func (ls LineSlice) ToSExp() string {
	var r []string
	for _, l := range ls {
		r = append(r, l.ToSExp())
	}
	return strings.Join(r, "\n")

}

func (l Line) SchemLib() string {
	l.Origin.Y = -l.Origin.Y
	l.End.Y = -l.End.Y
	return fmt.Sprintf("P 2 0 0 %.0f %s %s N\r\n", l.Width, l.Origin.Rounded(), l.End.Rounded())
}
