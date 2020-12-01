package state

import (
	"math/rand"
	"testing"
	"time"
)



func TestFSM(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	ts := &Turnstile{
		ID:     1,
		State:  "Locked",
		States: []string{"Locked"},
	}
	fsm := initFSM()

	//推门
	//没刷卡/投币不可进入
	err := fsm.Trigger(ts.State, "Push", ts)
	if err != nil {
		t.Errorf("trigger err: %v", err)
	}

	//推门
	//没刷卡/投币不可进入
	err = fsm.Trigger(ts.State, "Push", ts)
	if err != nil {
		t.Errorf("trigger err: %v", err)
	}

	//刷卡或者投币
	//不容易啊，终于解锁了
	err = fsm.Trigger(ts.State, "Coin", ts)
	if err != nil {
		t.Errorf("trigger err: %v", err)
	}

	//刷卡或者投币
	//无用的投币, 测试Action执行失败
	err = fsm.Trigger(ts.State, "Coin", ts)
	if err != nil {
		t.Logf("trigger err: %v", err)
	}

	//推门
	//这时才能进入，进入后闸门被锁
	err = fsm.Trigger(ts.State, "Push", ts)
	if err != nil {
		t.Errorf("trigger err: %v", err)
	}

	//推门
	//无法进入，闸门已锁
	err = fsm.Trigger(ts.State, "Push", ts)
	if err != nil {
		t.Errorf("trigger err: %v", err)
	}

	lastState := Turnstile{
		ID:         1,
		EventCount: 6,
		CoinCount:  1,
		PassCount:  1,
		State:      "Locked",
		States:     []string{"Locked", "Unlocked", "Locked"},
	}

	if !compareTurnstile(&lastState, ts) {
		t.Errorf("Expected last state: %+v, but got %+v", lastState, ts)
	} else {
		t.Logf("最终的状态: %+v", ts)
	}
}
