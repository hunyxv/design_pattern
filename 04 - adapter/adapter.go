package adapter

import (
	"fmt"
	"reflect"
)

// Executer 适配的接口
type Executer interface {
	Execute() string
}

var _ Executer = (*Computer)(nil)

type Computer struct {
	Name string
}

func NewExecuter(name string) Executer {
	return &Computer{Name: name}
}

func (c *Computer) Execute() string {
	return "executes a program"
}

func (c *Computer) String() string {
	return fmt.Sprintf("the %s computer", c.Name)
}

// Player 被适配的接口
type Player interface {
	Play() string
}

var _ Player = (*Synthesizer)(nil)

type Synthesizer struct {
	Name string
}

func NewSynthesizer(name string) *Synthesizer {
	return &Synthesizer{Name: name}
}

func (s *Synthesizer) Play() string {
	return "is playing an electronic song"
}

func (s *Synthesizer) String() string {
	return fmt.Sprintf("the %s synthesizer", s.Name)
}

// Speaker 被适配的接口
type Speaker interface {
	Speak() string
}

var _ Speaker = (*Human)(nil)

type Human struct {
	Name string
}

func NewHuman(name string) *Human {
	return &Human{Name: name}
}

func (h *Human) Speak() string {
	return "says hello"
}

func (h *Human) String() string {
	return fmt.Sprintf("%s the human", h.Name)
}

var _ Executer = (*Adapter)(nil)

// Adapter 适配器
type Adapter struct {
	obj    interface{}
	method string
}

func NewAdapter(obj interface{}, method string) Executer {
	return &Adapter{
		obj:    obj,
		method: method,
	}
}

func (a *Adapter) Execute() string {
	objValue := reflect.ValueOf(a.obj)
	fun := objValue.MethodByName(a.method)
	return fun.Call(nil)[0].String()
}

func (a *Adapter) String() string {
	s := (a.obj).(fmt.Stringer)
	return fmt.Sprint(s)
}
