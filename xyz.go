package gokicadlib

import "fmt"

type XYZ struct {
	x, y, z float64
}

func (a XYZ) ToSexp() string {
	return "xyz" + fmt.Sprint(a.x, a.y, a.z)
}
