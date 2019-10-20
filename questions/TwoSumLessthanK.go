package questions

import (
	"math"
	"sort"
)

func TwoSumLessThanK(A []int, K int) int {
	max := -1
	left := 0
	sort.Ints(A)
	right := len(A) - 1
	for left < right {
		sum := A[left] + A[right]
		if sum < K {
			max = int(math.Max(float64(sum),float64(max)))
			left++
		} else {
			right--
		}

	}
	return max
}
