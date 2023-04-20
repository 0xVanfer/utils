package utils

import (
	"math"
)

// Solve the equation(within cubic).
func SolveSimpleEquations(params ...float64) []float64 {
	if len(params) < 2 {
		return nil
	}
	// ax+b=0
	if len(params) == 2 {
		return solveEquationLinear(params[0], params[1])
	}
	// ax^2+bx+c=0
	if len(params) == 3 {
		return solveEquationSquare(params[0], params[1], params[2])
	}
	// ax^3+bx^2+cx+d=0
	if len(params) == 4 {
		return solveEquationCubic(params[0], params[1], params[2], params[3])
	}
	// Numerical solutions need more params.
	return nil
}

// Solve the function `ax+b=0`.
func solveEquationLinear(a, b float64) []float64 {
	// No solution or any $x \in \R$ both return nil.
	if a == 0 {
		return nil
	}
	if b == 0 {
		return []float64{0}
	}
	return []float64{-b / a}
}

// Solve the function `ax^2+bx+c=0`.
func solveEquationSquare(a, b, c float64) []float64 {
	if a == 0 {
		return solveEquationLinear(b, c)
	}
	delta := b*b - 4*a*c
	if delta < 0 {
		return nil
	}
	x1 := (-b + math.Sqrt(delta)) / a / 2
	x2 := (-b - math.Sqrt(delta)) / a / 2
	res := []float64{x1, x2}
	return res
}

// Solve the function `ax^3+bx^2+cx+d=0`.
//
// Only gives the real solutions.
// Not sure about the accuracy. Not that sure the result is true when delta>0.
func solveEquationCubic(a, b, c, d float64) []float64 {
	if a == 0 {
		return solveEquationSquare(b, c, d)
	}
	// x^3+px+q=0
	p := (3*a*c - math.Pow(b, 2)) / (3 * math.Pow(a, 2))
	q := (2*math.Pow(b, 3) - 9*a*b*c + 27*math.Pow(a, 2)*d) / (27 * math.Pow(a, 3))
	// delta = (q/2)^2+(p/3)^3
	delta := math.Pow(q/2, 2) + math.Pow(p/3, 3)
	// fmt.Println("p:", p, "q:", q, "delta:", delta)

	if delta > 0 { // 1 real, 2 complex roots
		var u, v float64
		if q > 0 {
			u = -math.Pow(q/2+math.Sqrt(delta), 1.0/3)
			if q/2-math.Sqrt(delta) > 0 {
				v = -math.Pow(q/2-math.Sqrt(delta), 1.0/3)
			} else {
				v = math.Pow(-q/2+math.Sqrt(delta), 1.0/3)
			}
		} else {
			if q/2+math.Sqrt(delta) > 0 {
				u = -math.Pow(q/2+math.Sqrt(delta), 1.0/3)
			} else {
				u = math.Pow(-q/2-math.Sqrt(delta), 1.0/3)
			}
			v = math.Pow(-q/2+math.Sqrt(delta), 1.0/3)
		}
		A := u + v
		x1 := A - b/(3*a)
		return []float64{x1}
	} else if delta == 0 { // 3 real roots (at least two are equal)
		x1 := -2*math.Pow(q/2, 1.0/3) - b/(3*a)
		x2 := math.Pow(q/2, 1.0/3) - b/(3*a)
		x3 := math.Pow(q/2, 1.0/3) - b/(3*a)
		res := []float64{x1, x2, x3}
		return res
	} else { // 3 real roots
		var phi float64
		if q > 0 {
			phi = math.Atan(math.Sqrt(-delta) / (-q / 2))
		} else {
			phi = math.Atan(math.Sqrt(-delta)/(-q/2)) + math.Pi
		}
		r := 2 * math.Sqrt(-p/3)
		x1 := r*math.Cos(phi/3) - b/(3*a)
		x2 := r*math.Cos((phi+2*math.Pi)/3) - b/(3*a)
		x3 := r*math.Cos((phi+4*math.Pi)/3) - b/(3*a)
		res := []float64{x1, x2, x3}
		return res
	}
}
