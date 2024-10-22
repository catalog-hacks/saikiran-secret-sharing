// File: internal/secretsharing/vss/vss_test.go

package vss

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/SaiKiranMatta/secret-sharing/pkg/secretsharing" // Import the VSS interface from the pkg
)

// setupFeldmanVSSParams initializes the test setup for Feldman VSS
func setupFeldmanVSSParams(secret *big.Int, threshold, numShares int) (secretsharing.VSS, *big.Int, [][2]*big.Int, []*big.Int, error) {
	g := big.NewInt(2)
	q, err := GeneratePrime(256)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Create VSS parameters
	params := NewFeldmanVSSParams(g, q)

	// Generate shares and commitments
	shares, commitments, err := params.GenerateShares(secret, threshold, numShares)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return params, q, shares, commitments, nil
}

// TestFeldmanVerifiableSecretSharing tests the core functionality of Feldman VSS
func TestFeldmanVerifiableSecretSharing(t *testing.T) {
	// Test parameters
	secret := big.NewInt(986743267)
	threshold := 2
	numShares := 5

	// Setup VSS scheme
	vssScheme, q, shares, commitments, err := setupFeldmanVSSParams(secret, threshold, numShares)
	if err != nil {
		t.Fatalf("Failed to setup VSS: %v", err)
	}

	// Verify all shares
	for i, share := range shares {
		if !vssScheme.VerifyShare(share[0], share[1], commitments) {
			t.Errorf("Share %d failed verification", i+1)
		}
	}

	// Reconstruct secret
	reconstructedSecret := vssScheme.ReconstructSecret(shares[:threshold], q)
	if reconstructedSecret == nil {
		t.Fatal("Failed to reconstruct secret")
	}

	fmt.Printf("Reconstructed Secret: %v\n", reconstructedSecret)

	if reconstructedSecret.Cmp(secret) != 0 {
		t.Errorf("Reconstructed secret %v doesn't match original secret %v",
			reconstructedSecret, secret)
	}
}

// TestCorruptedShares checks that corrupted shares fail verification
func TestCorruptedShares(t *testing.T) {
	secret := big.NewInt(986743267)
	threshold := 2
	numShares := 5

	// Setup VSS scheme
	vssScheme, _, shares, commitments, err := setupFeldmanVSSParams(secret, threshold, numShares)
	if err != nil {
		t.Fatalf("Failed to setup VSS: %v", err)
	}

	// Corrupt the commitments
	corruptedCommitments := make([]*big.Int, len(commitments))
	for i, commitment := range commitments {
		corruptedCommitments[i] = new(big.Int).Add(commitment, big.NewInt(1))
	}

	// Verify shares with corrupted commitments - should fail
	for i, share := range shares {
		if vssScheme.VerifyShare(share[0], share[1], corruptedCommitments) {
			t.Errorf("Share %d unexpectedly passed verification with corrupted commitments", i+1)
		}
	}
}
