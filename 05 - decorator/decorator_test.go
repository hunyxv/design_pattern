package decorator

import (
	"testing"
	"time"
)

func runFib(n int, fib Fibonacci) (t time.Duration) {
	start := time.Now()
	for i := 0; i<n; i++ {
		_ = fib.Calc(n)
	}
	t = start.Sub(time.Now())
	return
}

func TestDecorator(t *testing.T) {
	fib := NewFib()
	t.Log(runFib(30, fib))

	fibPlus := NewFibPlus(fib)
	t.Log(runFib(30, fibPlus))
}

func TestDecoratorFunc(t *testing.T) {
	t.Log(runFib(30, FibFunc(Fibfunc)))
	
	fibFunc := Decorator(Fibfunc)
	t.Log(runFib(30, FibFunc(fibFunc)))
}