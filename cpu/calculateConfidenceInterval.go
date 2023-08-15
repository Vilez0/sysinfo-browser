package cpu

import (
	"math"
)

func CalculateConfidenceInterval(samples []float64) (float64, []float64) {
	arrayLong := len(samples)
	var sum, sDeviation, confidenceLevel float64
	confidenceLevel = 0.95
	//* Calculating samples mean
	for i := 0; i < arrayLong; i++ {
		sum += (samples[i])
	}
	mean := sum / float64(arrayLong)

	//* Calculating the standard deviation
	for j := 0; j < len(samples); j++ {
		sDeviation += math.Pow(samples[j]-mean, 2)
	}
	sDeviation = math.Sqrt(sDeviation / 10)

	//* Calculate confidence interval and return
	s := (confidenceLevel * (sDeviation / math.Sqrt(float64(arrayLong))))
	highest := mean + s
	lowest := mean - s

	result := []float64{lowest, highest}
	// fmt.Printf("lowest: %v, highest: %v\n", lowest, highest,)
	return mean, result
}
