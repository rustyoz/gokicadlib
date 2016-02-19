package gokicadlib

import "fmt"

type Font struct {
	Size      Point
	Thickness float32
}

func (f Font) ToSExp() string {
	s := NewSExp("size", false, f.Size.ToString())
	fmt.Println(f.Size.Y)
	t := NewSExp("thickness", false, fmt.Sprint(f.Thickness))
	return NewSExp("font", false, s, t)
}
