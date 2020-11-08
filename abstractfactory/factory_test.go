package abstractfactory

import "testing"

func TestFrogWorld(t *testing.T) {
	var world World = &FrogWorld{PlayerName: "Billy"}

	var game GameEnvironment = NewGame(world)

	game.Play()
}

func TestWizardWorld(t *testing.T) {
	var world World = &WizardWorld{PlayerName: "Charles"}

	var game GameEnvironment = NewGame(world)

	game.Play()
}
