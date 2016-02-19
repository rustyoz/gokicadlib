package gokicadlib

import (
	"fmt"
	"strconv"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) FromString(x string, y string) {
	p.X, _ = strconv.ParseFloat(x, 64)
	p.Y, _ = strconv.ParseFloat(y, 64)
}

func (p Point) ToString() string {
	return fmt.Sprintf("%.4f %.4f", p.X, p.Y)
}
