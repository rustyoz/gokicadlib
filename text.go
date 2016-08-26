package gokicadlib

import (
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
	width         float64
	Font          Font
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
