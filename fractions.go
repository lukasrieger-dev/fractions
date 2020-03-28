package fractions

import (
	"github.com/lukasrieger-dev/utils"
)

type fraction struct {
	numerator   int
	denominator int
}

func New(n, d int) fraction {
	return fraction{n, d}
}

func EqualFractionsValues(fract1, fract2 fraction) bool {
	return fract1.toFloat() == fract2.toFloat()
}

func EqualFractions(fract1, fract2 fraction) bool {
	n := fract1.numerator == fract2.numerator
	d := fract1.denominator == fract2.denominator

	return n && d
}

func (f *fraction) Normalize() {
	divisor := utils.GreatestCommonDivisor(f.numerator, f.denominator)
	f.numerator = f.numerator / divisor
	f.denominator = f.denominator / divisor
}

func (f fraction) toFloat() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

func (f fraction) Add(fract fraction) fraction {
	commonDenominator := utils.LowestCommonMultiple(f.denominator, fract.denominator)

	multiplicator1 := commonDenominator / f.denominator
	a := f.numerator * multiplicator1

	multiplicator2 := commonDenominator / fract.denominator
	b := fract.numerator * multiplicator2

	result := fraction{a + b, commonDenominator}
	result.Normalize()

	return result
}

func (f fraction) Sub(fract2 fraction) fraction {
	return f.Add(New(-1*fract2.numerator, fract2.denominator))
}

func (f fraction) Div(fract fraction) fraction {
	numerator := f.numerator * fract.denominator
	denominator := f.denominator * fract.numerator

	result := New(numerator, denominator)
	//result.Normalize()
	return result
}

func (f fraction) Mult(fract fraction) fraction {
	numerator := f.numerator * fract.numerator
	denominator := f.denominator * fract.denominator

	result := fraction{numerator, denominator}
	//result.Normalize()
	return result
}
