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
	Offset     int
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

func (s *Symbol) KicadLib() string {
	output := &bytes.Buffer{}
	var l string

	if s.Reference.Width == 0 {
		s.Reference.Width = 50
	}

	if s.Value.Width == 0 {
		s.Value.Width = 50
	}
	if s.Offset == 0 {
		s.Offset = 50
	}
	// symbol definition
	l = fmt.Sprintf("DEF %s %s 0 %d Y Y 1 L N\r\n", s.Name, s.Reference.Text, s.Offset)
	output.WriteString(l)

	// symbolic prefix text
	l = fmt.Sprintf("F0 \"%s\" %s %.0f H V C CNN\r\n", s.Reference.Text, s.Reference.Origin.Rounded(), s.Reference.Width)
	output.WriteString(l)
	// symbolic prefix text
	l = "F1 \"" + s.Value.Text + "\" " + s.Value.Origin.Rounded() + " " + fmt.Sprint(s.Value.Width) + " H V C CNN" + "\r\n"
	output.WriteString(l)

	l = "ALIAS " + fmt.Sprint(s.Alias) + "\r\n"
	output.WriteString(l)

	l = "DRAW\r\n"
	output.WriteString(l)

	for _, p := range s.Pins {
		output.WriteString(p.KicadLib())
	}
	for _, l := range s.Lines {
		output.WriteString(l.SchemLib())
	}

	l = "ENDDRAW\r\nENDDEF\r\n"
	output.WriteString(l)

	return output.String()
}
