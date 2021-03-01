package calculator

import "errors"

func Plus(a, b float64) float64 {
	return a + b
}

func Minus(a, b float64) float64 {
	return a - b
}

func Times(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}
