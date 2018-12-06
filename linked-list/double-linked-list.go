package linked_list

type DoubleNode struct {
	pre   *DoubleNode
	next  *DoubleNode
	value interface{}
}

func (s *DoubleNode) Pre() *DoubleNode {
	return s.pre
}

func (s *DoubleNode) Next() *DoubleNode {
	return s.next
}

func (s *DoubleNode) Value() interface{} {
	return s.value
}

type DoubleLinkedList struct {
	len  int
	head *DoubleNode
	tail *DoubleNode
}

func (d *DoubleLinkedList) Init() {
	d.head = new(DoubleNode)
	d.tail = new(DoubleNode)
	d.head.next = d.tail
	d.tail.pre = d.head
	d.len = 0
}

func (d *DoubleLinkedList) insert(newNode, pos *DoubleNode) {
	if nil == newNode || nil == pos || nil == pos.pre {
		return
	}
	pos.pre.next = newNode
	newNode.pre = pos.pre
	newNode.next = pos
	pos.pre = newNode
	d.len++
}

func (d *DoubleLinkedList) delete(pos *DoubleNode) (next *DoubleNode) {
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

func (d *DoubleLinkedList) PushFront(data interface{}) {
	if nil == d.head {
		return
	}
	newNode := new(DoubleNode)
	newNode.value = data
	d.insert(newNode, d.head.next)
}

func (d *DoubleLinkedList) PushBack(data interface{}) {
	if nil == d.tail {
		return
	}
	newNode := new(DoubleNode)
	newNode.value = data
	d.insert(newNode, d.tail)
}

func (d *DoubleLinkedList) Delete(data interface{}) {
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

func (d *DoubleLinkedList) GetNode(pos int) *DoubleNode {
	if pos > d.len {
		return nil
	}
	if pos > d.len/2 {
		node := d.tail
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

func (d *DoubleLinkedList) Len() int {
	return d.len
}
