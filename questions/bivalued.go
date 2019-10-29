package questions

import "sort"

func Bivalued(A []int) int {
	if len(A) == 0 {
		return 0
	}
	if len(A) <= 2 {
		return 2
	}

	commonA := A[0:2]
	for i := 2; i < len(A); i++ {
		if A[i] != commonA[0] && A[i] != commonA[1] {
			if len(commonA) == 2 {
				commonA[0], commonA[1] = commonA[1], commonA[0]
				commonA = commonA[:1]
				commonA = append(commonA, A[i])
			}
		} else {
			commonA = append(commonA, A[i])
		}
	}
	return len(commonA)
}

type pr struct {
	First  int
	Second int
}

func ImprovedBivalued(A []int) int {
	if len(A) == 0 {
		return 0
	}
	if len(A) <= 2 {
		return 2
	}
	rA := make([]int, 0)
	for inx, val := range A {
		if inx+1 < len(A) {
			commonA := []int{val, A[inx+1]}
			count := 2
			for i := 2; i < len(A); i = i + 2 {
				if A[i] == commonA[0] || A[i] == commonA[1] {

					if A[i] == commonA[0] {
						if (i+1) < len(A) && A[i+1] == commonA[1] {
							count = count + 2
						}
					}
					if A[i] != commonA[0] {
						if (i+1) < len(A) && A[i+1] == commonA[0] {
							count = count + 2
						}
					}
					count++
				}
			}
			rA = append(rA, count)
		}
	}
	sort.Ints(rA)
	return rA[len(rA)-1]
}
