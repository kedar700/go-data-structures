package trees

type bstTree struct {
	Left  *bstTree
	Value int
	Right *bstTree
}

func newBstTree(left *bstTree, value int, right *bstTree) *bstTree {
	return &bstTree{Left: left, Value: value, Right: right}
}

func NewV(t *bstTree, v int) *bstTree {
	if t == nil {
		return newBstTree(nil, v, nil)
	}
	if v < t.Value {
		t.Left = NewV(t.Left, v)
		return t
	}

	t.Right = NewV(t.Right, v)
	return t
}

func Walk(t *bstTree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *bstTree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

//noinspection GoUnreachableCode
func findByDFS(t *bstTree, v int) bool {
	if t.Value == v {
		return true
	}
	if v < t.Value {
		return findByDFS(t.Left, v)
	} else {
		return findByDFS(t.Right, v)
	}
	return false
}

func findByBFS(t *bstTree, v int) bool {
	queue := make([]*bstTree, 0)
	queue = append(queue, t)
	for len(queue) != 0 {
		nextT := queue[0]
		queue = queue[1:]
		if nextT.Value == v {
			return true
		}
		if v < nextT.Value && nextT.Left != nil {
			queue = append(queue, nextT.Left)
		}

		if v > nextT.Value && nextT.Right != nil {
			queue = append(queue, nextT.Right)
		}
	}
	return false
}

