package secretsharing

// Share represents a single share with X and Y values
type Share struct {
	X int
	Y int
}

// ShamirSecretSharing interface defines the operations for Shamir's Secret Sharing
type ShamirSecretSharing interface {
	Share(secret int, parts int, threshold int) ([]Share, error)
	Reconstruct(shares []Share) (int, error)
}
