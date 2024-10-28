package mathFunc

// Get average. (total sum)/ (data size)
func GetAvg(data []float64) float64 {
	var totalSum float64
	dataSize := len(data)

	for _, dataPoint := range data {
		totalSum += dataPoint
	}

	return totalSum / float64(dataSize)
}
