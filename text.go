package gokicadlib

import (
	"bytes"
	"fmt"
	"strings"
)

// Text kicad text emelement
type Text struct {
	Text          string
	Type          string
	Layer         Layer
	Origin        Point
	Visible       bool
	Justification string
	Style         string
	Width         float64
	Font          Font
	Orientation   TextOrientation
}

type TextOrientation int

const (
	Horizontal TextOrientation = 0
	Vertical                   = 1
)

func (to TextOrientation) String() string {
	if to == Horizontal {
		return "0"
	}
	return "1"
}

type TextSlice []Text

// ToSExp	convert to kicad s-expression format
func (t *Text) ToSExp() string {
	if t.Type == "" {
		t.Type = "user"
	}
	layer := NewSExp("layer", false, string(t.Layer))
	origin := NewSExp("at", false, t.Origin.ToString())
	if t.Type == "user" {
		effects := NewSExp("effects", false, t.Font.ToSExp())
		return NewSExp(fmt.Sprintf("fp_text %s %s %s%s", t.Type, t.Text, origin, layer), true, effects)
	} else {
		return fmt.Sprintf("(fp_text %s %s %s %s)", t.Type, t.Text, origin, layer)
	}
}

func (ts TextSlice) ToSExp() string {
	var r []string
	for _, t := range ts {
		r = append(r, t.ToSExp())
	}
	return strings.Join(r, "\n")

}

func (t Text) SchemLib() string {
	var out string
	out = "T " + t.Orientation.String() + " " + t.Origin.ToString() + "0 0 0" + t.Text + "\r\n"

	return out
}

func (ts TextSlice) Schemlib() string {
	var o bytes.Buffer
	for _, t := range ts {
		o.WriteString(t.SchemLib())
	}
	return o.String()

}
