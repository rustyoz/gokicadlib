package gokicadlib

import "strconv"

type Point struct {
	X string
	Y string
}

func (p Point) ToPointFloat() PointFloat {
	x, _ := strconv.ParseFloat(p.X, 64)
	y, _ := strconv.ParseFloat(p.Y, 64)
	return PointFloat{x, y}
}

type PointFloat struct {
	X float64
	Y float64
}

func (p Point) ToString() string {
	return p.X + " " + p.Y
}
