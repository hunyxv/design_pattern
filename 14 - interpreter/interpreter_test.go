package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	parser := &Parser{}

	err := parser.Parse("2 × (2 + 2) ÷ 2 + 1 × ((3 + 9) ÷ 6) - 1")
	if err != nil {
		t.Fatal(err)
	}

	n, err := parser.Result()
	if err != nil {
		t.Fatal(err)
	}

	if 2*(2+2)/2+1*((3+9)/6-1) != n.Interpret() {
		t.Fatalf("The result should be %d, not %d", 2*(2+2)/2+1*((3+9)/6), n.Interpret())
	}
}