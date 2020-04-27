package questions

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[0 : l-1]
	return v
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

type MaxStack struct {
	stack    Stack
	maxStack Stack
}

/** initialize your data structure here. */
func EConstructor() MaxStack {
	return MaxStack{}
}

func (this *MaxStack) Push(x int) {
	this.stack.Push(x)
	if this.maxStack != nil && len(this.maxStack) > 0 {
		currentMax := this.maxStack.Peek()
		if x > currentMax {
			this.maxStack.Push(x)
		} else {
			this.maxStack.Push(currentMax)
		}
	} else {
		this.maxStack.Push(x)
	}
}

func (this *MaxStack) Pop() int {
	this.maxStack.Pop()
	return this.stack.Pop()
}

func (this *MaxStack) Top() int {
	return this.stack.Peek()
}

func (this *MaxStack) PeekMax() int {
	return this.maxStack.Peek()
}

func (this *MaxStack) PopMax() int {
	var buffer Stack
	currentMax := this.maxStack.Peek()
	for {
		v := this.Top()
		if v == currentMax {
			break
		}
		buffer.Push(this.Pop())
	}
	this.Pop()
	for {
		if buffer == nil || len(buffer) == 0 {
			break
		}
		this.Push(buffer.Pop())
	}
	return currentMax
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
