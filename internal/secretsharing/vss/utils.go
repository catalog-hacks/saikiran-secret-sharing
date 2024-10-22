// File: internal/secretsharing/vss/utils.go

package vss

import (
	"crypto/rand"
	"math/big"
)

// Polynomial represents a mathematical polynomial with big.Int coefficients
type Polynomial struct {
    Coefficients []*big.Int
}

// NewPolynomialForShamir creates a new random polynomial for Shamir's scheme
func NewPolynomialForShamir(threshold int, secretBits int, secret *big.Int) (*Polynomial, error) {
    coefficients := make([]*big.Int, threshold)
    coefficients[0] = new(big.Int).Set(secret)

    max := new(big.Int).Lsh(big.NewInt(1), uint(secretBits))
    for i := 1; i < threshold; i++ {
        coef, err := rand.Int(rand.Reader, max)
        if err != nil {
            return nil, err
        }
        coefficients[i] = coef
    }

    return &Polynomial{Coefficients: coefficients}, nil
}

// Evaluate computes the value of the polynomial at point x
func (p *Polynomial) Evaluate(x *big.Int) *big.Int {
    result := big.NewInt(0)
    xPow := big.NewInt(1)

    for _, coef := range p.Coefficients {
        term := new(big.Int).Mul(coef, xPow)
        result.Add(result, term)
        xPow.Mul(xPow, x)
    }

    return result
}

// ModExp calculates (base^exponent) mod modulus efficiently
func ModExp(base, exponent, modulus *big.Int) *big.Int {
    return new(big.Int).Exp(base, exponent, modulus)
}

// GeneratePrime generates a prime number of the given bit size
func GeneratePrime(bitSize int) (*big.Int, error) {
    return rand.Prime(rand.Reader, bitSize)
}

// ModInv calculates the modular multiplicative inverse
func ModInv(a, m *big.Int) *big.Int {
    return new(big.Int).ModInverse(a, m)
}

// LagrangeInterpolationZero performs Lagrange interpolation at x=0
func LagrangeInterpolationZero(points [][2]*big.Int, modulus *big.Int) *big.Int {
    secret := big.NewInt(0)

    for i := range points {
        numerator := big.NewInt(1)
        denominator := big.NewInt(1)

        for j := range points {
            if i != j {
                // x_diff = (modulus - x_j) % modulus
                xDiff := new(big.Int).Sub(modulus, points[j][0])
                xDiff.Mod(xDiff, modulus)
                
                // denominator = denominator * (x_i - x_j) % modulus
                denomDiff := new(big.Int).Sub(points[i][0], points[j][0])
                denomDiff.Mod(denomDiff, modulus)
                
                numerator.Mul(numerator, xDiff)
                numerator.Mod(numerator, modulus)
                denominator.Mul(denominator, denomDiff)
                denominator.Mod(denominator, modulus)
            }
        }

        invDenominator := new(big.Int).ModInverse(denominator, modulus)
        if invDenominator == nil {
            return nil
        }

        term := new(big.Int).Mul(points[i][1], numerator)
        term.Mul(term, invDenominator)
        term.Mod(term, modulus)
        
        secret.Add(secret, term)
        secret.Mod(secret, modulus)
    }

    return secret
}