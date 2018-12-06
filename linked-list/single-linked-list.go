package linked_list

import "errors"

type SingleNode struct {
	next  *SingleNode
	value interface{}
}

func (s *SingleNode) Next() *SingleNode {
	return s.next
}

func (s *SingleNode) Value() interface{} {
	return s.value
}

type SingleLinkedList struct {
	head *SingleNode
	len  int
}

func (s *SingleLinkedList) Init() {
	s.head = new(SingleNode)
}

func (s *SingleLinkedList) insert(newNode *SingleNode, pos int) {
	if nil == s.head || pos-1 > s.len {
		return
	}
	node := s.head
	for i := 0; i < pos-1; i++ {
		node = node.next
	}
	newNode.next = node.next
	node.next = newNode
	s.len++
}

func (s *SingleLinkedList) Delete(pos int) error {
	if nil == s.head || pos > s.len {
		return errors.New("can not find data")
	}
	node := s.head
	for i := 0; i < pos-1; i++ {
		node = node.next
	}
	delNode := node.next
	if nil != delNode.next {
		node.next = node.next.next
		delNode.next = nil
	}
	s.len--
	return nil
}

func (s *SingleLinkedList) Insert(data interface{}, pos int) {
	newNode := new(SingleNode)
	newNode.value = data
	s.insert(newNode, pos)
}

func (s *SingleLinkedList) Len() int {
	return s.len
}

func (s *SingleLinkedList) GetValue(pos int) interface{} {
	if nil == s.head || pos > s.len {
		return errors.New("can not find data")
	}
	node := s.head
	for i := 0; i < pos; i++ {
		node = node.next
	}
	return node.value
}
