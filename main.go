package main

import (
	"fmt"

	"github.com/robertobouses/w2v-api/user"
)

type Calculador interface {
	Area() (string, float64)
}

func Mayor(C1, C2 Calculador) (string, float64) {
	nombre1, valor1 := C1.Area()
	nombre2, valor2 := C2.Area()
	if valor1 > valor2 {
		return nombre1, valor1
	}
	return nombre2, valor2
}

func main() {
	fmt.Printf("hello world")

	Circulo1 :=
		user.Circulo{"círculo", 5}
	Rectagulo1 := user.Rectangulo{"rectángulo", 10, 3}

	nombre, resultado := Mayor(Circulo1, Rectagulo1)
	fmt.Println("El área mayor es", nombre, resultado)
}
