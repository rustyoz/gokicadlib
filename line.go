package gokicadlib

import "strings"

type Line struct {
	Origin Point
	End    Point
	Layer  Layer
	Width  string
}

func (line Line) ToSExp() string {
	start := NewSExp("start", false, line.Origin.ToString())
	end := NewSExp("end", false, line.End.ToString())
	layer := NewSExp("layer", false, line.Layer.Name)
	width := NewSExp("width", false, line.Width)
	return NewSExp("fp_line", true, start, end, layer, width)
}

type LineSlice []Line

func (ls LineSlice) ToSExp() string {
	var r []string
	for _, l := range ls {
		r = append(r, l.ToSExp())
	}
	return strings.Join(r, "\n")

}
