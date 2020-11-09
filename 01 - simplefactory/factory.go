package simplefactory

import (
	"fmt"
)

type T int

const (
	Hi T = iota
	Hello
)

type API interface {
	Say(name string) string
}

func NewAPI(t T) API {
	switch t {
	case Hi:
		return &hiApi{}
	case Hello:
		return &helloApi{}
	}
	return nil
}

type hiApi struct{}

func (h *hiApi) Say(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}

type helloApi struct{}

func (h *helloApi) Say(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
