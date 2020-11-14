package adapter

import "testing"

func TestAdapter(t *testing.T) {
	objects := make([]Executer, 0)
	objects = append(objects, NewExecuter("Asus"))

	player := NewSynthesizer("moog")
	objects = append(objects, NewAdapter(player, "Play"))

	human := NewHuman("Bob")
	objects = append(objects, NewAdapter(human, "Speak"))

	for _, obj := range objects {
		t.Logf("%s %s", obj, obj.Execute())
	}
}