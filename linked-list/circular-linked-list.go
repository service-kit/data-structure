package linked_list

import "errors"

type SingleCircularLinkedList struct {
	head *SingleNode
	len  int
}

func (s *SingleCircularLinkedList) Init() {
	s.head = new(SingleNode)
	s.head.next = s.head
}

func (s *SingleCircularLinkedList) insert(newNode *SingleNode, pos int) {
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

func (s *SingleCircularLinkedList) Delete(pos int) error {
	if nil == s.head || pos > s.len {
		return errors.New("can not find data")
	}
	node := s.head
	for i := 0; i < pos-1; i++ {
		node = node.next
	}
	delNode := node.next
	node.next = node.next.next
	delNode.next = nil
	s.len--
	return nil
}

func (s *SingleCircularLinkedList) Insert(data interface{}, pos int) {
	newNode := new(SingleNode)
	newNode.value = data
	s.insert(newNode, pos)
}

func (s *SingleCircularLinkedList) Len() int {
	return s.len
}

func (s *SingleCircularLinkedList) GetValue(pos int) interface{} {
	if nil == s.head || pos > s.len {
		return errors.New("can not find data")
	}
	node := s.head
	for i := 0; i < pos; i++ {
		node = node.next
	}
	return node.value
}

type DoubleCircularLinkedList struct {
	len  int
	head *DoubleNode
}

func (d *DoubleCircularLinkedList) Init() {
	d.head = new(DoubleNode)
	d.head.next = d.head
	d.head.pre = d.head
	d.len = 0
}

func (d *DoubleCircularLinkedList) insert(newNode, pos *DoubleNode) {
	if nil == newNode || nil == pos || nil == pos.pre {
		return
	}
	pos.pre.next = newNode
	newNode.pre = pos.pre
	newNode.next = pos
	pos.pre = newNode
	d.len++
}

func (d *DoubleCircularLinkedList) delete(pos *DoubleNode) (next *DoubleNode) {
	if nil == pos.pre || nil == pos.next {
		return
	}
	pos.pre.next = pos.next
	pos.next.pre = pos.pre
	pos.pre = nil
	next = pos.next
	pos.next = nil
	d.len--
	return next
}

func (d *DoubleCircularLinkedList) PushFront(data interface{}) {
	if nil == d.head {
		return
	}
	newNode := new(DoubleNode)
	newNode.value = data
	d.insert(newNode, d.head.next)
}

func (d *DoubleCircularLinkedList) PushBack(data interface{}) {
	if nil == d.head {
		return
	}
	newNode := new(DoubleNode)
	newNode.value = data
	d.insert(newNode, d.head)
}

func (d *DoubleCircularLinkedList) Delete(data interface{}) {
	node := d.head
	for i := 0; i < d.len; i++ {
		if nil == node {
			break
		}
		node = node.next
		if data == node.value {
			node = d.delete(node)
		}
	}
}

func (d *DoubleCircularLinkedList) GetNode(pos int) *DoubleNode {
	if pos > d.len {
		return nil
	}
	node := d.head
	if pos > d.len/2 {
		for i := d.len; i >= pos; i-- {
			node = node.pre
		}
		return node
	} else {
		node := d.head
		for i := 0; i < pos; i++ {
			node = node.next
		}
		return node
	}
}

func (d *DoubleCircularLinkedList) Len() int {
	return d.len
}
