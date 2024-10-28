package mathFunc

import "math"

// Get standard deviation = sqrt(variance)
func GetStandardDeviation(data []float64) float64 {
	return math.Sqrt(float64(GetVariance(data)))
}
