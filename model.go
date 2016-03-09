package gokicadlib

type Model struct {
	Filename string
	origin   XYZ
	scale    XYZ
	rotation XYZ
}

func (m Model) ToSExp() string {
	at := NewSExp("at", false, m.origin.ToSExp())
	scale := NewSExp("scale", false, m.scale.ToSExp())
	rotate := NewSExp("rotate", false, m.rotation.ToSExp())

	return NewSExp("model "+m.Filename, true, at, scale, rotate)
}
