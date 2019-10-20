package trees

import "strings"

type Node struct {
	Tag      string
	Text     string
	Id       string
	Class    string
	Children []*Node
}

// This is the breadth first search method wherein we will first look at the neighbors before then looking at the children.
func findById(root *Node, id string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextVal := queue[0]
		queue = queue[1:]
		if nextVal.Id == id {
			return nextVal
		}
		if len(nextVal.Children) != 0 {
			// This will append the children to the slice.
			queue = append(queue, nextVal.Children...)
		}
	}
	return nil
}

func findByIdDFS(root *Node, id string) *Node {

	if root.Id == id {
		return root
	}
	for _, child := range root.Children {
		findByIdDFS(child, id)
	}
	return nil
}

func (n *Node) hasClass(className string) bool {
	classes := strings.Fields(n.Class)
	for _, class := range classes {
		if class == className {
			return true
		}
	}
	return false
}
