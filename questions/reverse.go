package questions

import "math"

func Reverse(x int) int {
	max := math.MaxInt32
	min := math.MinInt32
	if x > max || x < min {
		return 0
	}
	reversed := 0
	for x != 0 {
		rem := x % 10
		reversed = reversed*10 + rem
		x /= 10
	}
	if reversed > max || reversed < min {
		return reversed
	}
	return 0
}
