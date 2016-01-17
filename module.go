package gokicadlib

type Module struct {
	name string

	layer     KicadLayer
	tedit     TimeStamp
	tstamp    TimeStamp
	origin    XYZ
	descr     string
	tags      []string
	path      string
	reference Text
	value     Text
	lines     []Line
	pads      []Pad
	model     Model
}

func (m Module) ToSexp() string {
	module := JoinStrings("module", m.reference.Text, m.layer.String(), NewSExp("tedit", false, m.tedit.String()), NewSExp("tstamp", false, m.tstamp.String()))
	at := m.origin.ToSexp()
	tags := NewSExp("tags", false, m.tags...)
	path := NewSExp("path", false, m.path)
	reference := m.reference.ToSExp()
	value := m.reference.ToSExp()
	lines := LineSlice(m.lines).ToSExp()
	return NewSExp(module, true, at, tags, path, reference, value, lines)
}
