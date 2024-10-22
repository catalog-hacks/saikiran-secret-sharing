// pkg/secretsharing/vss.go
package secretsharing

import "math/big"

// VSS defines the interface for a Verifiable Secret Sharing scheme
type VSS interface {
	// GenerateShares generates shares and commitments for the secret
	GenerateShares(secret *big.Int, threshold, numShares int) ([][2]*big.Int, []*big.Int, error)

	// VerifyShare verifies a share using the public commitments
	VerifyShare(x, share *big.Int, commitments []*big.Int) bool

	// ReconstructSecret reconstructs the secret from shares using Lagrange interpolation
	ReconstructSecret(shares [][2]*big.Int, modulus *big.Int) *big.Int
}
