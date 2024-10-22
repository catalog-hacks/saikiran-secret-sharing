// internal/secretsharing/vss/vss.go
package vss

import (
	"math/big"

	"github.com/SaiKiranMatta/secret-sharing/pkg/secretsharing" // Import the interface from the external package
)

// FeldmanVSSParams represents the public parameters for the Feldman VSS scheme
type FeldmanVSSParams struct {
	Gen        *big.Int // Generator of the group G
	GroupOrder *big.Int // Prime order of the group G
}

// NewFeldmanVSSParams creates new Feldman VSS parameters
func NewFeldmanVSSParams(generator, primeOrder *big.Int) *FeldmanVSSParams {
	return &FeldmanVSSParams{
		Gen:        generator,
		GroupOrder: primeOrder,
	}
}

// GenerateShares generates shares and commitments for the secret
func (params *FeldmanVSSParams) GenerateShares(secret *big.Int, threshold, numShares int) ([][2]*big.Int, []*big.Int, error) {
	// Create polynomial with random coefficients
	poly, err := NewPolynomialForShamir(threshold, secret.BitLen(), secret)
	if err != nil {
		return nil, nil, err
	}

	// Generate shares
	shares := make([][2]*big.Int, numShares)
	for i := 1; i <= numShares; i++ {
		participantID := big.NewInt(int64(i))
		shareVal := poly.Evaluate(participantID)
		shareVal.Mod(shareVal, params.GroupOrder)
		shares[i-1] = [2]*big.Int{participantID, shareVal}
	}

	// Generate commitments
	commitments := params.generateCommitments(poly)

	return shares, commitments, nil
}

// generateCommitments generates commitments for the polynomial coefficients
func (params *FeldmanVSSParams) generateCommitments(poly *Polynomial) []*big.Int {
	commitments := make([]*big.Int, len(poly.Coefficients))
	for i, coef := range poly.Coefficients {
		commitments[i] = ModExp(params.Gen, coef, params.GroupOrder)
	}
	return commitments
}

// VerifyShare verifies a share using the public commitments
func (params *FeldmanVSSParams) VerifyShare(x, share *big.Int, commitments []*big.Int) bool {
	// Calculate LHS = Gen^share mod GroupOrder
	leftHandSide := ModExp(params.Gen, share, params.GroupOrder)

	// Calculate RHS = Product(commitment_i^(x^i) mod GroupOrder)
	rightHandSide := big.NewInt(1)
	xPower := big.NewInt(1)

	for _, commitment := range commitments {
		term := ModExp(commitment, xPower, params.GroupOrder)
		rightHandSide.Mul(rightHandSide, term)
		rightHandSide.Mod(rightHandSide, params.GroupOrder)
		xPower.Mul(xPower, x)
		xPower.Mod(xPower, params.GroupOrder)
	}

	return leftHandSide.Cmp(rightHandSide) == 0
}

// ReconstructSecret reconstructs the secret from shares using Lagrange interpolation
func (params *FeldmanVSSParams) ReconstructSecret(shares [][2]*big.Int, modulus *big.Int) *big.Int {
	return LagrangeInterpolationZero(shares, modulus)
}

// Ensure that FeldmanVSSParams implements the VSS interface
var _ secretsharing.VSS = (*FeldmanVSSParams)(nil)
