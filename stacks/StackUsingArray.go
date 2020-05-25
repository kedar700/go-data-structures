package stacks

import "sync"


// ItemStack the stack of Items
type ItemStack struct {
	items []interface{}
	lock  sync.RWMutex
}

// New creates a new ItemStack
func (s *ItemStack) New() *ItemStack {
	s.items = make([]interface{},0)
	return s
}

// Push adds an Item to the top of the stack
func (s *ItemStack) Push(t interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

// Pop removes an Item from the top of the stack
func (s *ItemStack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return &item
}

// Get the length of the stack
func (s *ItemStack) Length() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.items)
}

// Get the top element in the stack
func (s *ItemStack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	return &s.items[len(s.items)-1]
}
