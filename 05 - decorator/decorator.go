package decorator

type Fibonacci interface {
	Calc(int) int
}

type Fib struct {
}

func NewFib() *Fib {
	return &Fib{}
}

func (f *Fib) Calc(n int) int {
	if n < 0 {
		panic("n must be >= 0")
	}

	if n == 0 || n == 1 {
		return n
	}

	return f.Calc(n-1) + f.Calc(n-2)
}

type FibPlus struct {
	Fibonacci
	cach map[int]int
}

func NewFibPlus(fib Fibonacci) *FibPlus {
	return &FibPlus{
		Fibonacci: fib,
		cach:      map[int]int{0: 0, 1: 1},
	}
}

func (fp *FibPlus) Calc(n int) int {
	if r, ok := fp.cach[n]; ok {
		return r
	}
	fp.cach[n] = fp.Fibonacci.Calc(n)
	return fp.cach[n]
}

/*
	使用装饰器函数
*/

type FibFunc func(int) int

func (f FibFunc) Calc(n int) int {
	return f(n)
}

func Fibfunc(n int) int {
	if n < 0 {
		panic("n must be >= 0")
	}

	if n == 0 || n == 1 {
		return n
	}

	return Fibfunc(n-1) + Fibfunc(n-2)
}

// Decorator 装饰器函数
func Decorator(f FibFunc) FibFunc {
	cach := map[int]int{0: 0, 1: 1}
	return func(n int) int {
		if r, ok := cach[n]; ok {
			return r
		}
		cach[n] = f(n)
		return cach[n]
	}
}
