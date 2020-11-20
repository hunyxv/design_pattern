package proxy

import "fmt"

type House interface {
	Sail()
}

type SmailHouse struct {

}

func (SmailHouse) Sail() {
	fmt.Println("100 万")
}

type Proxy struct {
	house SmailHouse
}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p Proxy) Sail() {
	var result string = "签订成功"
	p.Before()
	p.house.Sail()
	p.After()
	fmt.Println(result)
}

func (Proxy) Before() {
	fmt.Println("代理之前的一些检查")
}

func (Proxy) After() {
	fmt.Println("代理之后的一些检查")
}
