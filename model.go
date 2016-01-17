package gokicadlib

type Model struct {
	Filename string
	origin   XYZ
	scale    XYZ
	rotation XYZ
}

func (m Model) ToSexp() string {
	at := NewSExp("at", false, m.origin.ToSexp())
	scale := NewSExp("scale", false, m.scale.ToSexp())
	rotate := NewSExp("rotate", false, m.rotation.ToSexp())

	return NewSExp("model "+m.Filename, true, at, scale, rotate)
}
