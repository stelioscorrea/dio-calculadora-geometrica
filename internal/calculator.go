package geometry

import (
	"fmt"
	"math"
)

type Geometria interface {
	Area() float64
	Perimetro() float64
}

type Quadrado struct {
	Lado float64
}

type Circunferencia struct {
	Raio float64
}

func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

func (q Quadrado) Perimetro() float64 {
	return 4 * q.Lado
}

func (c Circunferencia) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circunferencia) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

func Calculate(g Geometria) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimetro())
}
