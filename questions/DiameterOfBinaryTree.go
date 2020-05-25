package questions

//Definition for a binary tree node.
type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

/**
func diameterOfBinaryTree(root *TreeNode) int {
    ans := 0
    helper(root, &ans)
    return ans
}

func helper(root *TreeNode, ans *int) int{
    if root == nil{
        return 0
    }
    l:= helper(root.Left, ans)
    r:= helper(root.Right, ans)
    fmt.Println("ans ##", ans)
    *ans = max(*ans, l + r)
    return 1 + max(l, r)
}

*/
func diameterOfBinaryTree(root *TreeNode3) int {
	answer := 0
	heightOfTheTree(root, &answer)
	return answer
}

func heightOfTheTree(root *TreeNode3, answer *int) int {
	if root == nil {
		return 0
	}
	leftHeight := heightOfTheTree(root.Left, answer)
	rightHeight := heightOfTheTree(root.Right, answer)
	*answer = max(*answer, leftHeight+rightHeight)
	return 1 + max(leftHeight, rightHeight)
}

func max(left int, right int) int {
	if left > right {
		return left
	}
	return right
}
