package questions

/*
Input:  [1,2,3,4]
Output: [24,12,8,6]
*/
func ProductExceptSelf(nums []int) []int {

	outputArr := make([]int, len(nums))
	outputArr[0] = 1
	for i := 1; i < len(nums); i++ {
		outputArr[i] = nums[i-1] * outputArr[i-1]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		outputArr[i] *= right
		right *= nums[i]
	}
	return outputArr
}
