package gokicadlib

type Table struct {
	Columns []string
}

type RGBA [4]float32

type Library struct{}
type LibraryTable struct {
	Libraries []Library
}

type Part struct{}

type PartList struct {
	Table Table
	Parts []Part
}

type Property struct {
	Name   string
	Value  string
	Origin Point
	Effect Effect
}
type Effect struct {
	Origin        Point
	Font          Font
	Justification string
	Color         RGBA
}

type Hint struct{}

type SchRectangle struct {
	Start     Point
	End       Point
	LineWidth float32
	Fill      string
}

type SchCircle struct {
	Center    Point
	Radius    float32
	LineWidth float32
	Fill      string
}

type Bezier struct {
	Points    []Point
	LineWidth float32
	Fill      string
}

type Image struct {
	Type   string
	Origin Point
	Alpha  float32
	Scale  struct {
		Hor  float32
		Vert float32
	}
	Data     []byte
	Filename string
}

type Component struct {
	Ref      string
	Basename string
}
type Wire struct{}
type Bus struct{}
type Label struct{}
type GlobalLabel struct{}
type HierachicalLabel struct{}
type Junction struct{}
type NoConnect struct{}
type Sheet struct{}
type Circuit struct{}

type Schematic struct {
	LibTable  LibraryTable
	PartsList PartList

	Properties        []Property
	Hints             []Hint
	Lines             []Line
	Rectangles        []SchRectangle
	Circles           []SchCircle
	Arcs              []Arc
	Bezier            []Bezier
	Text              []Text
	Images            []Image
	Components        []Component
	Wires             []Wire
	Buses             []Bus
	Labels            []Label
	GlobalLables      []GlobalLabel
	HierachicalLabels []HierachicalLabel
	Junctions         []Junction
	NoConnects        []NoConnect
	Sheets            []Sheet
	Circuits          []Circuit
}
