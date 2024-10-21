package sss

// Polynomial represents a mathematical polynomial
type Polynomial struct {
	Coefficients []int
}

// Evaluate calculates the value of the polynomial at a given x
func (p *Polynomial) Evaluate(x int) int {
	y := 0
	temp := 1
	for _, coeff := range p.Coefficients {
		y += coeff * temp
		temp *= x
	}
	return y
}
