package secretsharing

import (
	"math/big"
)

type Share struct {
	X *big.Int 
	Y *big.Int 
}

type ShamirSecretSharing interface {
	Share(secret *big.Int, parts int, threshold int) ([]Share, error) 
	Reconstruct(shares []Share, k int) (*big.Int, error)               
}
