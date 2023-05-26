package user

type Rectangulo struct {
	Descripcion string
	Largo       float64
	Alto        float64
}

func (R Rectangulo) Area() (string, float64) {
	return R.Descripcion, R.Largo * R.Alto
}
