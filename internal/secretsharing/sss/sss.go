package sss

import (
	"errors"
	"math/rand"
	"time"

	"github.com/SaiKiranMatta/secret-sharing/pkg/secretsharing"
)

// ShamirSecretSharing implementation
type Shamir struct{}

// Generate polynomial and share the secret into N parts with threshold K
func (s *Shamir) Share(secret, parts, threshold int) ([]secretsharing.Share, error) {
	if threshold > parts {
		return nil, errors.New("threshold cannot be greater than the number of parts")
	}

	// Initialize polynomial of degree threshold - 1
	poly := make([]int, threshold)
	poly[0] = secret

	rand.Seed(time.Now().UnixNano())
	for i := 1; i < threshold; i++ {
		poly[i] = rand.Intn(997) // Random coefficients mod a prime (997)
	}

	shares := make([]secretsharing.Share, parts)
	for i := 1; i <= parts; i++ {
		x := i
		y := calculateY(x, poly)
		shares[i-1] = secretsharing.Share{X: x, Y: y}
	}

	return shares, nil
}

// Reconstruct secret from shares using Lagrange interpolation
func (s *Shamir) Reconstruct(shares []secretsharing.Share, k int) (int, error) {
	if len(shares) == 0 {
		return 0, errors.New("no shares provided")
	}else if (len(shares) < k) {
		return 0, errors.New("not enough keys provided")	
	}

	secret := 0
	for i := 0; i < k; i++ {
		num, den := 1, 1
		for j := 0; j < len(shares); j++ {
			if i != j {
				num *= -shares[j].X
				den *= (shares[i].X - shares[j].X)
			}
		}
		f := Fraction{Num: shares[i].Y * num, Den: den}
		f.Reduce()
		secret += f.Num / f.Den
	}
	return secret, nil
}

// Helper function to calculate y = poly[0] + x*poly[1] + x^2*poly[2] + ...
func calculateY(x int, poly []int) int {
	y := 0
	temp := 1
	for _, coeff := range poly {
		y += coeff * temp
		temp *= x
	}
	return y
}
