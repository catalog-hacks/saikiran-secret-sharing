package sss

// Polynomial represents a mathematical polynomial
type Polynomial struct {
	Coefficients []int
}

// Evaluate calculates the value of the polynomial at a given x
func (p *Polynomial) Evaluate(x int) int {
	y := 0
	temp := 1
	for _, coeff := range p.Coefficients {
		y += coeff * temp
		temp *= x
	}
	return y
}


type Fraction struct {
	Num int
	Den int
}

// Reduce simplifies the fraction by dividing both num and den by their greatest common divisor (GCD)
func (f *Fraction) Reduce() {
	gcd := GCD(f.Num, f.Den)
	f.Num /= gcd
	f.Den /= gcd
}

// GCD calculates the greatest common divisor using the Euclidean algorithm
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
