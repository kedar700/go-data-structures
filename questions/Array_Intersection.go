package questions

func Intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	returnmp := make(map[int]int, 0)
	mp := make(map[int]bool, 0)
	for _, val := range nums1 {
		mp[val] = true
	}
	for _, val := range nums2 {
		if _, ok := mp[val]; ok {
			returnmp[val] = 1
		}
	}
	returnArr := make([]int, len(returnmp))
	i := 0
	for k, _ := range returnmp {
		returnArr[i] = k
		i++
	}
	return returnArr
}

func ExpensiveIntersection(nums1 []int, nums2 []int) []int {

	if len(nums1) < len(nums2) {
		return CalculateIntersection(nums2, nums1)
	}
	return CalculateIntersection(nums1, nums2)
}

func CalculateIntersection(big []int, small []int) []int {
	set := make([]int, 0)

	for _, smallVal := range small {
		for _, bigVal := range big {
			if smallVal == bigVal && !Contains(set, smallVal) {
				set = append(set, smallVal)
			}
		}
	}
	return set
}

func Contains(set []int, p int) bool {
	for _, v := range set {
		if v == p {
			return true
		}
	}
	return false
}
