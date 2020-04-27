package questions

type TreeNode2 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type traveller struct {
	maxDepth int
}

func (t *traveller) traverse(root *TreeNode2, currentDepth int) {
	if root == nil {
		return
	}
	if t.maxDepth < currentDepth {
		t.maxDepth = currentDepth
	}
	if root.Left != nil {
		t.traverse((*TreeNode2)(root.Left), currentDepth+1)
	}
	if root.Right != nil {
		t.traverse((*TreeNode2)(root.Right), currentDepth+1)
	}
}

func maxDepth(root *TreeNode2) int {

	if root == nil {
		return 0
	}
	t:= traveller{}
	t.traverse(root,1)
	return t.maxDepth
}
