package proxy

import "testing"

func TestProxy(t *testing.T) {
    var house House = NewProxy()
    house.Sail()
}