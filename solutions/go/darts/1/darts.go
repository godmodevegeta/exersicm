package darts

import "math"

func Score(x, y float64) int {
	var score int
	distance_squared := (x * x) + (y * y)
	distance := math.Sqrt(distance_squared)
	switch {
	case distance <= 10.0 && distance > 5.0 :
		score += 1	
	case distance <= 5.0 && distance > 1.0 :
		score += 5
	case distance <= 1.0 :
		score += 10
	default:
		score += 0
	}
	return score

}
