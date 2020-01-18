package questions

import (
	"github.com/golang-collections/collections/stack"
	"strings"
)

type LNode struct {
	C        string
	Children map[string]*LNode
	End      int
}

func NewLNode(c string) *LNode {
	return &LNode{C: c, Children: make(map[string]*LNode, 0)}
}

type Trie struct {
	Words []string
	Root  *LNode
}

func NewTrie() *Trie {
	return &Trie{Root: NewLNode("0")}
}

func (f *Trie) insert(word string, index int) {
	cur := f.Root
	charArray := strings.Split(word, "")
	for _, val := range charArray {
		cur.Children[val] = NewLNode(val)
	}
	cur.End = index
}

func (f *Trie) dfs() string {
	ans := ""
	stck := stack.New()
	stck.Push(f.Root)
	for stck.Len() != 0 {
		LNode := stck.Pop().(*LNode)
		if LNode.End > 0 || LNode == f.Root {
			if LNode != f.Root {
				word := f.Words[LNode.End-1]
				if len(word) > len(ans) || len(word) == len(ans) && strings.Compare(word, ans) < 0 {
					ans = word
				}
			}

			for _, value := range LNode.Children {
				stck.Push(value)
			}
		}
	}

	return ans
}
func LongestWord(words []string) string {
	tri := NewTrie()
	index := 0
	for _, word := range words {
		index++
		tri.insert(word, index)
	}
	tri.Words = words
	return tri.dfs()
}
