package gokicadlib

type Polygon struct {
	Layer  Layer
	Origin Point
	End    Point
	Points []Point
	Width  string
}
