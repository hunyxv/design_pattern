package iterator

// Value .
type Value int

// Aggregate 聚合接口
type Aggregate interface {
	Add(Value)
	GetIterator() Iterator
}

type node struct {
	val  Value
	next *node
}

// List 链表
type List struct {
	head *node
	tail *node
}

// NewList .
func NewList() *List {
	n := &node{}
	return &List{
		head: n,
		tail: n,
	}
}

// Add .
func (l *List) Add(val Value) {
	newNode := &node{val: val}
	l.tail.next = newNode
	l.tail = newNode
}

// GetIterator .
func (l *List) GetIterator() Iterator {
	return newConcreteIterator(l)
}

// Iterator 迭代器接口
type Iterator interface {
	Head() Value
	Tail() Value
	HasNext() bool
	Next() Value
}

// ConcreteIterator 迭代器实体
type ConcreteIterator struct {
	list *List

	cursor *node
}

func newConcreteIterator(l *List) *ConcreteIterator {
	return &ConcreteIterator{list: l, cursor: l.head}
}

// Head 链表头的值
func (i *ConcreteIterator) Head() Value {
	return i.list.head.next.val
}

// Tail 链表末尾的值
func (i *ConcreteIterator) Tail() Value {
	return i.list.tail.val
}

// HasNext 是否有下一个节点
func (i *ConcreteIterator) HasNext() bool {
	return i.cursor.next != nil
}

// Next 返回下一个节点的值
func (i *ConcreteIterator) Next() Value {
	i.cursor = i.cursor.next
	return i.cursor.val
}
