package sss

import (
	"fmt"
	"math/big"
	"testing"
)

func TestShare(t *testing.T) {
	shamirImpl := Shamir{}
	secret := big.NewInt(123) // Use *big.Int for the secret
	parts := 5
	threshold := 3

	shares, err := shamirImpl.Share(secret, parts, threshold)
	if err != nil {
		t.Fatalf("Error during sharing: %v", err)
	}

	if len(shares) != parts {
		t.Fatalf("Expected %d shares, got %d", parts, len(shares))
	} 

	for i, share := range shares {
		if share.X.Cmp(big.NewInt(0)) == 0 || share.Y.Cmp(big.NewInt(0)) == 0 {
			t.Errorf("Share %d is invalid: X=%d, Y=%d", i+1, share.X, share.Y)
		}
	}

	for i, share := range shares {
		fmt.Printf("Share %d is valid: X=%d, Y=%d\n", i+1, share.X, share.Y)
	}
}

func TestReconstruct(t *testing.T) {
	shamirImpl := Shamir{}
	secret := big.NewInt(123) // Use *big.Int for the secret
	parts := 5
	threshold := 3

	// Generate shares
	shares, err := shamirImpl.Share(secret, parts, threshold)
	if err != nil {
		t.Fatalf("Error during sharing: %v", err)
	}

	// Reconstruct secret from the first 'threshold' shares
	reconstructedSecret, err := shamirImpl.Reconstruct(shares[:threshold], threshold)
	if err != nil {
		t.Fatalf("Error during reconstruction: %v", err)
	}

	if reconstructedSecret.Cmp(secret) != 0 { // Compare big.Ints
		t.Errorf("Expected secret %d, got %d", secret, reconstructedSecret)
	}
}

// Number of keys provided is less than the required number of tests
func TestInvalidReconstruction(t *testing.T) {
	shamirImpl := Shamir{}
	secret := big.NewInt(123) // Use *big.Int for the secret
	parts := 5
	threshold := 3

	// Generate shares
	shares, err := shamirImpl.Share(secret, parts, threshold)
	if err != nil {
		t.Fatalf("Error during sharing: %v", err)
	}

	// Try to reconstruct secret using only 1 share (less than threshold)
	reconstructedSecret, err := shamirImpl.Reconstruct(shares[:1], threshold)
	if err == nil {
		t.Fatal("Expected error during reconstruction with less than threshold shares, got nil")
	}

	if reconstructedSecret != nil && reconstructedSecret.Cmp(secret) == 0 {
		t.Errorf("Expected wrong secret %d, got secret %d", secret, reconstructedSecret)
	}
}

// Threshold value is wrong 
func TestInvalidThreshold(t *testing.T) {
	shamirImpl := Shamir{}
	secret := big.NewInt(123) 
	parts := 5
	threshold := 3

	// Generate shares
	shares, err := shamirImpl.Share(secret, parts, threshold)
	if err != nil {
		t.Fatalf("Error during sharing: %v", err)
	}

	// Try to reconstruct secret using less than required threshold
	res, err := shamirImpl.Reconstruct(shares[:threshold-1], threshold-1)
	if err != nil {
		t.Fatal("error during reconstruction with insufficient shares")
	}

	if res == secret {
		t.Errorf("Expected secret not to match %d, got %d", secret, res)
	}

}

func TestEdgeCases(t *testing.T) {
	shamirImpl := Shamir{}
	secret := big.NewInt(0) 
	parts := 3
	threshold := 2

	// Test with secret = 0
	shares, err := shamirImpl.Share(secret, parts, threshold)
	if err != nil {
		t.Fatalf("Error during sharing with zero secret: %v", err)
	}

	reconstructedSecret, err := shamirImpl.Reconstruct(shares[:threshold], threshold)
	if err != nil {
		t.Fatalf("Error during reconstruction with zero secret: %v", err)
	}

	if reconstructedSecret.Cmp(secret) != 0 { // Compare big.Ints
		t.Errorf("Expected secret %d, got %d", secret, reconstructedSecret)
	}
}
