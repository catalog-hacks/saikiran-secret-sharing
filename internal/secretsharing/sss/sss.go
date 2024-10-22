package sss

import (
	"errors"
	"math/big"
	"math/rand"
	"time"

	"github.com/SaiKiranMatta/secret-sharing/pkg/secretsharing"
)

// ShamirSecretSharing implementation using big.Int
type Shamir struct{}

// Generate polynomial and share the secret into N parts with threshold K
func (s *Shamir) Share(secret *big.Int, parts, threshold int) ([]secretsharing.Share, error) {
	if threshold > parts {
		return nil, errors.New("threshold cannot be greater than the number of parts")
	}

	// Initialize polynomial of degree threshold - 1
	poly := make([]*big.Int, threshold)
	poly[0] = new(big.Int).Set(secret) // secret as the first coefficient

	rand.Seed(time.Now().UnixNano())
	for i := 1; i < threshold; i++ {
		poly[i] = big.NewInt(rand.Int63n(997)) // Random coefficients mod a prime (997)
	}

	shares := make([]secretsharing.Share, parts)
	for i := 1; i <= parts; i++ {
		x := big.NewInt(int64(i))
		y := calculateY(x, poly)
		shares[i-1] = secretsharing.Share{
			X: x,
			Y: y,
		}
	}

	return shares, nil
}

// Reconstruct secret from shares using Lagrange interpolation
func (s *Shamir) Reconstruct(shares []secretsharing.Share, k int) (*big.Int, error) {
	if len(shares) == 0 {
		return nil, errors.New("no shares provided")
	} else if len(shares) < k {
		return nil, errors.New("not enough keys provided")
	}

	secret := big.NewInt(0)
	for i := 0; i < k; i++ {
		num := big.NewInt(1)
		den := big.NewInt(1)
		for j := 0; j < len(shares); j++ {
			if i != j {
				negXj := new(big.Int).Neg(shares[j].X)
				num.Mul(num, negXj) // num *= -shares[j].X

				diff := new(big.Int).Sub(shares[i].X, shares[j].X)
				den.Mul(den, diff) // den *= (shares[i].X - shares[j].X)
			}
		}

		// Fraction for Lagrange basis polynomial evaluation
		fNum := new(big.Int).Mul(shares[i].Y, num) // shares[i].Y * num
		fraction := new(big.Int).Div(fNum, den)    // fNum / den (integer division)

		// Add to the secret
		secret.Add(secret, fraction)
	}

	return secret, nil
}

// Helper function to calculate y = poly[0] + x*poly[1] + x^2*poly[2] + ...
func calculateY(x *big.Int, poly []*big.Int) *big.Int {
	y := big.NewInt(0)
	temp := big.NewInt(1)

	for _, coeff := range poly {
		term := new(big.Int).Mul(coeff, temp) // coeff * temp
		y.Add(y, term)                        // y += term
		temp.Mul(temp, x)                     // temp *= x
	}

	return y
}
