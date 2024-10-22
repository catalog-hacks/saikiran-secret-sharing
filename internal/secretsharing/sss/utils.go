package sss

import (
	"math/big"
)

// Polynomial represents a mathematical polynomial
type Polynomial struct {
	Coefficients []*big.Int // Changed to use *big.Int
}

// Evaluate calculates the value of the polynomial at a given x
func (p *Polynomial) Evaluate(x *big.Int) *big.Int {
	y := big.NewInt(0) // Initialize y as a big.Int
	temp := big.NewInt(1) // Initialize temp as 1

	for _, coeff := range p.Coefficients {
		term := new(big.Int).Mul(coeff, temp) // coeff * temp
		y.Add(y, term)                        // y += term
		temp.Mul(temp, x)                     // temp *= x
	}
	return y
}

// Fraction represents a mathematical fraction
type Fraction struct {
	Num *big.Int // Changed to use *big.Int
	Den *big.Int // Changed to use *big.Int
}

// Reduce simplifies the fraction by dividing both num and den by their greatest common divisor (GCD)
func (f *Fraction) Reduce() {
	gcd := GCD(f.Num, f.Den)
	f.Num.Div(f.Num, gcd) // Use big.Int division
	f.Den.Div(f.Den, gcd) // Use big.Int division
}

// GCD calculates the greatest common divisor using the Euclidean algorithm
func GCD(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)
	if b.Cmp(zero) == 0 {
		return a
	}
	return GCD(b, new(big.Int).Mod(a, b)) // Use big.Int Mod for modulus operation
}
