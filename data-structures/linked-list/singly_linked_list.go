package list

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func newNode(i int) *Node {
	return &Node{data: i}
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

func (list *SinglyLinkedList) InsertFirst(i int) {
	node := newNode(i)
	if list.head != nil {
		node.next = list.head
	}
	list.head = node
}

func (list *SinglyLinkedList) InsertLast(i int) {
	node := newNode(i)
	if list.head == nil {
		list.head = node
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (list *SinglyLinkedList) GetFirst() (i int, ok bool) {
	 if list.head != nil {
	 	return list.head.data, true
	 }

	 return 0, false
}

func (list *SinglyLinkedList) GetLast() (i int, ok bool) {
	if list.head == nil {
		return 0, false
	}

	p := list.head
	for p.next != nil {
		p = p.next
	}

	return p.data, true
}

func (list *SinglyLinkedList) GetSize() (size int) {
	if list.head == nil {
		return
	}

	p := list.head
	for p.next != nil {
		p = p.next
		size += 1
	}

	return size + 1
}

func (list *SinglyLinkedList) GetItems() (items []int) {
	if list.head == nil {
		return
	}

	p := list.head
	for p != nil {
		items = append(items, p.data)
		p = p.next
	}

	return items
}

func (list *SinglyLinkedList) RemoveByValue(i int) bool {
	if list.head == nil {
		return false
	}

	p := list.head

	if p.data == i {
		list.head = p.next
	}

	for p.next != nil {
		if p.next.data == i {
			p.next = p.next.next
			return true
		}
		p = p.next
	}

	return false
}

func (list *SinglyLinkedList) RemoveByIndex(i int) bool {
	if list.head == nil || i < 0 {
		return false
	}

	p := list.head
	if i == 0 {
		list.head = p.next
		return true
	}

	for j := 1; j < i; j++ {
		if p.next.next == nil {
			return false
		}
		p = p.next
	}

	p.next = p.next.next
	return true
}

func (list *SinglyLinkedList) SearchValue(i int) bool {
	if list.head == nil {
		return false
	}

	p := list.head
	for p != nil {
		if p.data == i {
			return true
		}
		p = p.next
	}

	return false
}