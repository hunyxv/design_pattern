package template

import "testing"


func TestDotsStyleTemplate(t *testing.T) {
	dotsStyle := NewDotsStyle()
	dotsStyle.ShowBanner("dots stype test")
}

func TestASCIIStyleTemplate(t *testing.T) {
	asciiStyle := NewASCIIStyle()
	asciiStyle.ShowBanner("ascii style test")
}