package chainofresponsibility

import "fmt"

// Event 事件
type Event int

const (
	// AskLeave 请假
	AskLeave Event = iota
	// Reimburse 报销
	Reimburse
	// Other 其他
	Other
)

// Chain .
type Chain interface {
	SetNextHandler(Chain)
}

// Handler .
type Handler interface {
	Handle(Event)
}

var _ Handler = (*HRDepartment)(nil)

// HRDepartment .
type HRDepartment struct {
	next Chain
}

// SetNextHandler .
func (hr *HRDepartment) SetNextHandler(chain Chain) {
	hr.next = chain
}

// Handler .
func (h *HRDepartment) Handle(event Event) {
	if event == AskLeave {
		fmt.Println("批准请假")
		return
	} else if h.next != nil {
		if handler, ok := (h.next).(Handler); ok {
			handler.Handle(event)
			return
		}
	}

	fmt.Println("该事件无法处理")
}

var _ Handler = (*FinanceDepartment)(nil)

// FinanceDepartment .
type FinanceDepartment struct {
	next Chain
}

// SetNextHandler .
func (f *FinanceDepartment) SetNextHandler(chain Chain) {
	f.next = chain
}

// Handler .
func (f *FinanceDepartment) Handle(event Event) {
	if event == Reimburse {
		fmt.Println("批准报销")
		return
	} else if f.next != nil {
		if handler, ok := f.next.(Handler); ok {
			handler.Handle(event)
			return
		}
	}

	fmt.Println("该事件无法处理")
}
