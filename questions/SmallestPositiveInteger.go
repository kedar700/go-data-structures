package questions


func SmallestPositiveInteger(A []int) int {
	length := len(A)
	for i := 0; i < length; i++ {
		if A[i] <= 0 || A[i] > length {
			A[i] = 0
		} else if A[i] == i+1 {
			continue
		} else if A[i] == A[A[i]-1] {
			A[i] = 0
		} else if A[i] != A[A[i]-1] {
			A[A[i]-1], A[i] = A[i], A[A[i]-1]
			i--
		}
	}

	for i, v := range A {
		if v == 0 {
			return i+1
		}
	}
	return length+1
}
