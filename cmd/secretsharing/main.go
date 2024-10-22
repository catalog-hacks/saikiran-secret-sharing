package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/SaiKiranMatta/secret-sharing/internal/secretsharing/sss" // Ensure this path is correct based on your project structure
	"github.com/SaiKiranMatta/secret-sharing/internal/secretsharing/vss" // Ensure this path is correct based on your project structure
)

func main() {
	// Example usage of Shamir's Secret Sharing
	shamir := &sss.Shamir{} 
	secretSSS := big.NewInt(123456) 
	partsSSS := 5
	thresholdSSS := 3

	// Share the secret using Shamir's Secret Sharing
	sharesSSS, err := shamir.Share(secretSSS, partsSSS, thresholdSSS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during sharing (SSS): %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Shamir's Secret Sharing - Shares generated:")
	for _, share := range sharesSSS {
		fmt.Printf("X: %s, Y: %s\n", share.X.String(), share.Y.String())
	}

	// Reconstruct the secret from Shamir's shares
	reconstructedSecretSSS, err := shamir.Reconstruct(sharesSSS[:thresholdSSS], thresholdSSS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during reconstruction (SSS): %v\n", err)
		os.Exit(1)
	}

	if reconstructedSecretSSS.Cmp(secretSSS) != 0 {
		fmt.Printf("Reconstructed secret (SSS) %s does not match the original secret %s\n", reconstructedSecretSSS.String(), secretSSS.String())
	} else {
		fmt.Printf("Reconstructed secret (SSS) matches the original secret: %s\n", reconstructedSecretSSS.String())
	}

	// Example usage of Verifiable Secret Sharing (VSS)
	g := big.NewInt(2) 
	q, err := vss.GeneratePrime(256) 
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate prime: %v\n", err)
		os.Exit(1)
	}
	vssParams := vss.NewFeldmanVSSParams(g, q)

	secretVSS := big.NewInt(654321) 
	partsVSS := 5
	thresholdVSS := 3

	// Generate shares and commitments using VSS
	vssShares, vssCommitments, err := vssParams.GenerateShares(secretVSS, thresholdVSS, partsVSS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during VSS sharing: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Verifiable Secret Sharing - VSS Shares generated:")
	for _, share := range vssShares {
		fmt.Printf("X: %s, Y: %s\n", share[0].String(), share[1].String())
	}

	// Verify the shares using VSS
	for i, share := range vssShares {
		if !vssParams.VerifyShare(share[0], share[1], vssCommitments) {
			fmt.Printf("VSS Share %d failed verification\n", i+1)
		} else {
			fmt.Printf("VSS Share %d passed verification\n", i+1)
		}
	}

	// Reconstruct the secret from VSS shares
	reconstructedSecretVSS := vssParams.ReconstructSecret(vssShares[:thresholdVSS], q)
	if reconstructedSecretVSS.Cmp(secretVSS) != 0 {
		fmt.Printf("Reconstructed secret (VSS) %s does not match the original secret %s\n", reconstructedSecretVSS.String(), secretVSS.String())
	} else {
		fmt.Printf("Reconstructed secret (VSS) matches the original secret: %s\n", reconstructedSecretVSS.String())
	}
}
