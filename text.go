package gokicadlib

// Text kicad text emelement
type Text struct {
	Text          string
	Layer         Layer
	Origin        Point
	Visible       bool
	Justification string
	Style         string
	width         float64
}

// ToSExp	convert to kicad s-expression format
func (t *Text) ToSExp() string {
	layer := NewSExp("layer", false, t.Layer.Name)
	origin := NewSExp("start", false, t.Origin.ToString())
	return NewSExp("fp_text", false, layer, origin)
}
