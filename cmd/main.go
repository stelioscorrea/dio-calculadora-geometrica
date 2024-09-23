package main

import (
	geometry "geometry/internal"
)

func main() {

	q1 := geometry.Quadrado{Lado: 10.0}
	c1 := geometry.Circunferencia{Raio: 5.0}

	geometry.Calculate(q1)
	geometry.Calculate(c1)
}
