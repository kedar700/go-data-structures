package linked_list

type Node struct {
	Info        string
	PublishDate int64
	Next        *Node
}

type linked_l struct {
	Length int
	Start  *Node
	End    *Node
}

func (l *linked_l) Append(newNode *Node) {
	if l.Length == 0 {
		l.Start = newNode
		l.End = newNode
	} else {
		lastNode := l.End
		lastNode.Next = newNode
		l.End = newNode
	}
	l.Length++
}

func (l *linked_l) Remove(publishDate int64) {
	if l.Length == 0 {
		panic("The length of the list is zero")
	} else {
		var lastNode = new(Node)
		newNode := l.Start
		for newNode.PublishDate != publishDate {
			if newNode.Next == nil {
				panic("Reached the end of the list and the element is not found")
			}
			lastNode = newNode
			newNode = newNode.Next
		}
		lastNode.Next = newNode.Next
		l.Length--
	}
}

func (l *linked_l) AppendWithSort(newNode *Node) {
	if l.Length == 0 {
		l.Start = newNode
		l.End = newNode
	} else {
		var lastNode = new(Node)
		nextNode := l.Start
		for nextNode.PublishDate < newNode.PublishDate {
			lastNode = nextNode
			nextNode = nextNode.Next
		}
		lastNode.Next = newNode
		newNode.Next = nextNode
	}
	l.Length++
}
