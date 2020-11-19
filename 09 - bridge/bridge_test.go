package bridge

import "testing"

func TestCommonSMS(t *testing.T) {
	m := NewCommonMessage(ViaSMS())
    m.SendMessage("have a drink?", "bob")
}

func TestCommonEmail(t *testing.T) {
    m := NewCommonMessage(ViaEmail())
    m.SendMessage("have a drink?", "bob")
}

func TestUrgencySMS(t *testing.T) {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
}

func TestUrgencyEmail(t *testing.T) {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
}