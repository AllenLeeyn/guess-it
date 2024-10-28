package mathFunc

// Get Median by finding mid point.
// If mid point is between two data point,
// add them up and divide by 2.
func GetMedian(data []float64) float64 {
	dataSize := len(data)
	midpointOne := (dataSize / 2)

	if dataSize%2 == 0 {
		return (data[midpointOne] + data[midpointOne-1]) / 2
	}
	return data[midpointOne]
}
