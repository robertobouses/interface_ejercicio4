package user

import "math"

type Circulo struct {
	Descripcion string
	Radio       float64
}

func (C Circulo) Area() (string, float64) {
	return C.Descripcion, math.Pi * math.Pow(C.Radio, 2)
}
