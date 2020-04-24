package questions

/**
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

func ContainsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	mp := make(map[int]int, len(nums))
	for idx, val := range nums {
		if mapIndex, ok := mp[val]; ok && (idx-mapIndex) <= k {
			return true
		}
		mp[val] = idx
	}
	return false
}
