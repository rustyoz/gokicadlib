package gokicadlib

// Module Kicad Module
type Module struct {
	Name string

	Layer     KicadLayer
	Tedit     TimeStamp
	Tstamp    TimeStamp
	Origin    XYZ
	Descr     string
	Tags      []string
	Path      string
	Reference Text
	Value     Text
	Lines     []Line
	Pads      []Pad
	Model     Model
}

// ToSExp convert to kicad s-expression
func (m Module) ToSExp() string {
	module := JoinStrings("module", m.Reference.Text, m.Layer.String(), NewSExp("tedit", false, m.Tedit.String()), NewSExp("tstamp", false, m.Tstamp.String()))
	at := m.Origin.ToSexp()
	tags := NewSExp("tags", false, m.Tags...)
	path := NewSExp("path", false, m.Path)
	reference := m.Reference.ToSExp()
	value := m.Value.ToSExp()
	lines := LineSlice(m.Lines).ToSExp()
	return NewSExp(module, true, at, tags, path, reference, value, lines)
}
