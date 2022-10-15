package collector

type doubleNode[T any] struct {
	data T
	prev *doubleNode[T]
	next *doubleNode[T]
}

func NewLinkList[T comparable]() List[T] {
	return &linkList[T]{}
}

type linkList[T comparable] struct {
	head *doubleNode[T]
	tail *doubleNode[T]
	len  int
}

func (list *linkList[T]) Len() int {
	return list.len
}

func (list *linkList[T]) IsEmpty() bool {
	return list.Len() == 0
}

func (list *linkList[T]) Slice(opts ...EachOpt) []T {
	slice := make([]T, 0, list.Len())
	list.forEach(func(data T) bool {
		slice = append(slice, data)
		return true
	}, opts...)
	return slice
}

func (list *linkList[T]) ForEach(fn func(T) bool, opts ...EachOpt) {
	list.forEach(fn, opts...)
}

func (list *linkList[T]) forEach(fn func(T) bool, opts ...EachOpt) {
	eo := newEachOpt(opts...)
	gon := true
	if eo.reverse {
		node := list.tail
		for node != nil && gon {
			gon = fn(node.data)
			node = node.prev
		}
	} else {
		node := list.head
		for node != nil && gon {
			gon = fn(node.data)
			node = node.next
		}
	}
}

func (list *linkList[T]) Add(data T) {
	list.insert(data, list.Len())
}
func (list *linkList[T]) Insert(data T, i int) {
	list.insert(data, i)
}
func (list *linkList[T]) insert(data T, i int) {
	list.len++
	item := &doubleNode[T]{data: data}
	if list.head == nil {
		list.head, list.tail = item, item
		return
	}

	var (
		next = list.head
		prev *doubleNode[T]
	)
	for i > 0 && next != nil {
		next = next.next
		i--
	}

	if next != nil {
		prev = next.prev
		next.prev = item
	} else {
		prev = list.tail
		list.tail = item
	}
	if prev != nil {
		prev.next, item.prev = item, prev
	} else {
		list.head = item
	}
	item.next = next
}

func (list *linkList[T]) Remove() (data T) {
	return list.delete(list.Len() - 1)
}
func (list *linkList[T]) Delete(i int) T {
	return list.delete(i)
}
func (list *linkList[T]) delete(i int) (data T) {
	if i < 0 || i >= list.Len() {
		return
	}
	list.len--
	node := list.head
	for i > 0 {
		node = node.next
		i--
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}

	node.prev, node.next = nil, nil
	data = node.data
	return
}
