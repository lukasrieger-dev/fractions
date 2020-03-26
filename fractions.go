package fractions

import (
	"github.com/lukasrieger-dev/utils"
)

type Fraction struct {
	numerator   int
	denominator int
}

func (f *Fraction) Normalize() {
	divisor := utils.GreatestCommonDivisor(f.numerator, f.denominator)
	f.numerator = f.numerator / divisor
	f.denominator = f.denominator / divisor
}

func (f Fraction) toFloat() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

func Div(fract1, fract2 Fraction) Fraction {
	numerator := fract1.numerator * fract2.denominator
	denominator := fract1.denominator * fract2.numerator

	result := Fraction{numerator, denominator}
	//result.Normalize()
	return result
}

func Mult(fract1, fract2 Fraction) Fraction {
	numerator := fract1.numerator * fract2.numerator
	denominator := fract1.denominator * fract2.denominator

	result := Fraction{numerator, denominator}
	//result.Normalize()
	return result
}
