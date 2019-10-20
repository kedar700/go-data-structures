package questions

func twoSum(numbers []int, target int) []int {
	returnArr := make([]int, 0)
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			returnArr = append(returnArr, left+1, right+1)
			break
		} else if sum < target {
			left++
		} else {
			right--
		}

	}
	return returnArr
}
