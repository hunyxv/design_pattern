package iterator

import "testing"

func TestIterator(t *testing.T) {
	var list Aggregate = NewList()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	iterator := list.GetIterator()

	t.Logf("list Head: %d, Tail: %d \n", iterator.Head(), iterator.Tail())

	for iterator.HasNext() {
		t.Log(iterator.Next())
	}
}
