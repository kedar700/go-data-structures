package questions

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestBSTSubtree(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return 0
	}
	return largestBSTSubtreeHelper(root)[2]
}
func largestBSTSubtreeHelper(node *TreeNode) []int {
	if node == nil {
		return []int{math.MaxInt32, math.MinInt32, 0}
	}
	left := largestBSTSubtreeHelper(node.Left)
	right := largestBSTSubtreeHelper(node.Right)
	if node.Val > left[1] && node.Val < right[0] {
		return []int{int(math.Min(float64(node.Val), float64(left[0]))), int(math.Max(float64(node.Val), float64(right[1]))), left[2] + right[2], 1}
	}
	return [] int{math.MinInt32, math.MaxInt32, int(math.Max(float64(left[2]), float64(right[2])))}
}
