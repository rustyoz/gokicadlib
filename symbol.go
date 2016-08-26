package gokicadlib

import (
	"bytes"
	"fmt"
)

// Symbol Kicad Symbol
type Symbol struct {
	Name string

	Layer      Layer
	Tedit      TimeStamp
	Tstamp     TimeStamp
	Origin     Point
	Descr      string
	Attributes []string
	Tags       []string
	Path       string
	Reference  Text
	Value      Text
	Arcs       []Arc
	Lines      []Line
	Pins       []Pin
	Alias      []string
	FootPrints []string
	Text       []Text
}

func (s *Symbol) KicadLib() *bytes.Buffer {
	output := &bytes.Buffer{}
	var l string
	// symbol definition
	l = "DEF " + s.Name + " " + s.Reference.Text + " " + s.Origin.ToString() + " Y Y 1 L N" + "\r\n"
	output.WriteString(l)

	// symbolic prefix text
	l = "F0 \"" + s.Reference.Text + "\" " + s.Reference.Origin.ToString() + " " + fmt.Sprint(s.Reference.width) + " H V C N" + "\r\n"
	output.WriteString(l)
	// symbolic prefix text
	l = "F1 \"" + s.Value.Text + "\" " + s.Value.Origin.ToString() + " " + fmt.Sprint(s.Value.width) + " H V C N" + "\r\n"
	output.WriteString(l)

	l = "ALIAS " + fmt.Sprint(s.Alias) + "\r\n"
	output.WriteString(l)

	l = "DRAW\r\n"
	output.WriteString(l)

	for _, p := range s.Pins {
		p.KicadLib().WriteTo(output)
	}

	l = "ENDDRAW\r\nENDDEF\r\n"
	output.WriteString(l)

	return output
}
