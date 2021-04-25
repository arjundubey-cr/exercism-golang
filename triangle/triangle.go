// Package triangle has a function which determines kind of triangle
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides is a function which determined kind of triangle on the basis of length of sides.
func KindFromSides(a, b, c float64) Kind {
	for _, side := range [3]float64{a, b, c} {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}
	switch {
	case a+b < c || a+c < b || c+b < a:
		return NaT
	case a == b && a == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	default:
		return Sca
	}
}
