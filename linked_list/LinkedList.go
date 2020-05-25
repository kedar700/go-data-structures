package linked_list

type ListData struct {
	Info string
	Date int64
	Next *ListData
}
type ListNode struct {
	Length int
	Start  *ListData
	End    *ListData
}

func (l *ListNode) Append(data *ListData) {
	if l.Length == 0 {
		l.Start = data
		l.End = data
	} else {
		lastNode := l.End
		lastNode.Next = data
		l.End = data
	}
	l.Length++
}

func (l *ListNode) Remove(data *ListData) {
	if l.Length == 0 {
		panic("The list has no elements!!")
	}
	prevNode := &ListData{}
	foundNode := l.Start
	for foundNode.Date != data.Date {
		if foundNode == nil {
			panic("Node you want to remove was not found")
		}
		prevNode = foundNode
		foundNode = foundNode.Next
	}
	prevNode.Next = foundNode.Next
	l.Length--

}

func (l *ListNode) AppendWithSort(data *ListData) {
	if l.Length == 0 {
		panic("The list has no elements!!")
	}
	prevNode := &ListData{}
	foundNode := l.Start
	for foundNode.Date < data.Date {
		prevNode = foundNode
		foundNode = foundNode.Next
	}
	prevNode.Next = data
	data.Next = foundNode

	l.Length++

}
