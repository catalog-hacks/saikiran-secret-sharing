package sss

// Fraction represents a rational number as a numerator and denominator
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
