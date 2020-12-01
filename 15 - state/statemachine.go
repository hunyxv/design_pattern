package state 

import (
	"errors"
	"fmt"
	"log"

	"github.com/smallnest/gofsm"
)

// Turnstile 旋转栅门
type Turnstile struct {
	ID         uint64
	EventCount uint64
	CoinCount  uint64
	PassCount  uint64
	State      string
	States     []string
}

// TurnstileEventProcessor is used to handle turnstile actions.
type TurnstileEventProcessor struct{}

func (p *TurnstileEventProcessor) OnExit(fromState string, args []interface{}) {
	t := args[0].(*Turnstile)
	if t.State != fromState {
		panic(fmt.Errorf("转门 %v 的状态与期望的状态 %s 不一致，可能在状态机外被改变了", t, fromState))
	}

	log.Printf("转门 %d 从状态 %s 改变", t.ID, fromState)
}

func (p *TurnstileEventProcessor) Action(action string, fromState string, toState string, args []interface{}) error {
	t := args[0].(*Turnstile)
	t.EventCount++

	switch action {
	case "pass": //用户通过的action
		t.PassCount++
	case "check", "repeat-check": //刷卡或者投币的action
		if t.CoinCount > 0 { // repeat-check
			return errors.New("转门暂时故障")
		}

		t.CoinCount++
	default: //其它action
	}

	return nil
}

func (p *TurnstileEventProcessor) OnEnter(toState string, args []interface{}) {
	t := args[0].(*Turnstile)
	t.State = toState
	t.States = append(t.States, toState)

	log.Printf("转门 %d 的状态改变为 %s ", t.ID, toState)
}

func (p *TurnstileEventProcessor) OnActionFailure(action string, fromState string, toState string, args []interface{}, err error) {
	t := args[0].(*Turnstile)

	log.Printf("转门 %d 的状态从 %s to %s 改变失败， 原因: %v", t.ID, fromState, toState, err)
}


func compareTurnstile(t1 *Turnstile, t2 *Turnstile) bool {
	if t1.ID != t2.ID || t1.CoinCount != t2.CoinCount || t1.EventCount != t2.EventCount || t1.PassCount != t2.PassCount ||
		t1.State != t2.State {
		return false
	}

	return fmt.Sprint(t1.States) == fmt.Sprint(t2.States)
}

func initFSM() *fsm.StateMachine {
	delegate := &fsm.DefaultDelegate{P: &TurnstileEventProcessor{}}

	transitions := []fsm.Transition{
		{From: "Locked", Event: "Coin", To: "Unlocked", Action: "check"},
		{From: "Locked", Event: "Push", To: "Locked", Action: "invalid-push"},
		{From: "Unlocked", Event: "Push", To: "Locked", Action: "pass"},
		{From: "Unlocked", Event: "Coin", To: "Unlocked", Action: "repeat-check"},
	}

	return fsm.NewStateMachine(delegate, transitions...)
}