package factorymethod

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var factory OperatorFactory
	
	factory = PlusOperatorFactory{}
	if compute(factory, 2, 2) != 4 {
		t.Fatal("error with factory method pattern")
	}

	factory = SubOPeratorFartory{}
	if compute(factory, 3, 1) != 2 {
		t.Fatal("error with factory method pattern")
	}
}