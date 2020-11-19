package bridge

import "fmt"

type AbstractMessage interface {
    SendMessage(text, to string)
}

type MessageImplementer interface {
    Send(text, to string)
}

// MessageSMS .
type MessageSMS struct{}

func ViaSMS() MessageImplementer {
    return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
    fmt.Printf("send %s to %s via SMS\n", text, to)
}

// MessageEmail .
type MessageEmail struct{}

func ViaEmail() MessageImplementer {
    return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
    fmt.Printf("send %s to %s via Email\n", text, to)
}

// CommonMessage 普通消息
type CommonMessage struct {
    method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
    return &CommonMessage{
        method: method,
    }
}

func (m *CommonMessage) SendMessage(text, to string) {
    m.method.Send(text, to)
}

// UrgencyMessage 紧急消息
type UrgencyMessage struct {
    method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
    return &UrgencyMessage{
        method: method,
    }
}

func (m *UrgencyMessage) SendMessage(text, to string) {
    m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
