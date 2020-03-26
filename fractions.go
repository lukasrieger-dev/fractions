package fractions

type Fraction struct {
	numerator   int
	denominator int
}

func (f Fraction) toFloat() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

func Div(fract1, fract2 Fraction) Fraction {
	numerator := fract1.numerator * fract2.denominator
	denominator := fract1.denominator * fract2.numerator

	return Fraction{numerator, denominator}
}

func Mult(fract1, fract2 Fraction) Fraction {
	numerator := fract1.numerator * fract2.numerator
	denominator := fract1.denominator * fract2.denominator

	return Fraction{numerator, denominator}
}
