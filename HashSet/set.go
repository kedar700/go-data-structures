package HashSet

import "sync"

type ItemSet struct {
	items map[interface{}]bool
	lock sync.RWMutex
}

func (s *ItemSet) Add(t interface{}) *ItemSet {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[interface{}]bool)
	}

	if _, ok := s.items[t]; !ok {
		s.items[t] = true
	}
	return s
}

func (s *ItemSet) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = make(map[interface{}]bool)
}

func (s *ItemSet) Delete(item interface{}) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

func (s *ItemSet) Contains(item interface{}) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[item]
	return ok
}

func (s *ItemSet) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.items)
}

func (s *ItemSet) Items() []interface{} {
	//s.lock.Lock()
	//defer s.lock.Unlock()
	items := make([]interface{}, s.Size())
	for _, val := range s.items {
		items = append(items, val)
	}
	return items
}

