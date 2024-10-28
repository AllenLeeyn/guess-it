package mathFunc

// Get variance. variance = variability / data size.
// variability = sum of square of data point distance from average.
func GetVariance(data []float64) float64 {
	var variability float64
	dataSize := len(data)
	avg := GetAvg(data)

	for _, dataPoint := range data {
		viaPoint := dataPoint - avg
		variability = variability + float64((viaPoint * viaPoint))
	}

	return variability / float64(dataSize)
}
