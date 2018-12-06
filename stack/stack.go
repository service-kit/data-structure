package stack

import "errors"

type Stack struct {
	capacity int
	size     int
	top      int
	array    []interface{}
}

func (s *Stack) Init() {
	s.array = make([]interface{}, 1024)
	s.capacity = 1024
	s.size = 0
	s.top = -1
}

func (s *Stack) ensureCapacity() {
	s.capacity *= 2
	newArray := make([]interface{}, s.capacity)
	copy(newArray, s.array)
	s.array = newArray
}

func (s *Stack) Push(obj interface{}) {
	if s.size == s.capacity {
		s.ensureCapacity()
	}
	s.top++
	s.array[s.top] = obj
	s.size++
}

func (s *Stack) Pop() (data interface{}, err error) {
	if 0 == s.size {
		return nil, errors.New("stack empty")
	}
	data = s.array[s.top]
	s.top--
	s.size--
	return
}

func (s *Stack) Peek() (data interface{}, err error) {
	if 0 == s.size {
		return nil, errors.New("stack empty")
	}
	data = s.array[s.top]
	return
}
