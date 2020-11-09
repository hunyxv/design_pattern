package factorymethod

type Operator interface {
	SetA(a int)
	SetB(b int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type operatorBase struct {
	a, b int
}

func (o *operatorBase) SetA(a int) {
	o.a = a
}

func (o *operatorBase) SetB(b int) {
	o.b = b
}

type PulsOperator struct {
	*operatorBase
}

func (p *PulsOperator) Result() int {
	return p.a + p.b
}

type PlusOperatorFactory struct{}

func (f PlusOperatorFactory) Create() Operator {
	return &PulsOperator{
		operatorBase: &operatorBase{},
	}
}

type SubOPerator struct {
	*operatorBase
}

func (s *SubOPerator) Result() int {
	return s.a - s.b
}

type SubOPeratorFartory struct{}

func (f SubOPeratorFartory) Create() Operator {
	return &SubOPerator{
		operatorBase: &operatorBase{},
	}
}
