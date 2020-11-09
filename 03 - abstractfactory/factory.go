package abstractfactory

import (
	"fmt"
)

// Hero 英雄
type Hero interface {
	// InteractWith 和 xx 交互
	InteractWith(Obstacle)
	String() string
}

var _ Hero = (*Frog)(nil)

// Frog 青蛙🐸
type Frog struct {
	Name string
}

func (f *Frog) InteractWith(o Obstacle) {
	act := o.Action()

	fmt.Printf("%s the Frog encounters %s and %s！\n", f, o, act)
}

func (f *Frog) String() string {
	return f.Name
}

var _ Hero = (*Wizard)(nil)

// Wizard 巫师💂
type Wizard struct {
	Name string
}

func (w *Wizard) InteractWith(o Obstacle) {
	act := o.Action()

	fmt.Printf("%s the Wizard battles against %s and %s!\n", w, o, act)
}

func (w *Wizard) String() string {
	return w.Name
}

// Obstacle 障碍物
type Obstacle interface {
	// Action 动作
	Action() string
	String() string
}

var _ Obstacle = (*Bug)(nil)

// Bug 虫子🐛
type Bug struct{}

func (b *Bug) Action() string {
	return "eat it"
}

func (b *Bug) String() string {
	return "a bug"
}

var _ Obstacle = (*Ork)(nil)

// Ork 兽人👹
type Ork struct{}

func (o *Ork) Action() string {
	return "kills it"
}

func (o *Ork) String() string {
	return "an evil ork"
}

// World  世界
type World interface {
	// MakeCharacter 初始化角色
	MakeCharacter() Hero
	// MakeObstacle 初始化障碍物
	MakeObstacle() Obstacle
	String() string
}

var _ World = (*FrogWorld)(nil)

// FrogWorld 青蛙的世界
type FrogWorld struct {
	PlayerName string
}

func (fw *FrogWorld) MakeCharacter() Hero {
	return &Frog{
		Name: fw.PlayerName,
	}
}

func (fw *FrogWorld) MakeObstacle() Obstacle {
	return &Bug{}
}

func (fw *FrogWorld) String() string {
	return "\n\n\t------ Frog World ------"
}

var _ World = (*WizardWorld)(nil)

// WizardWorld 巫师世界
type WizardWorld struct {
	PlayerName string
}

func (ww *WizardWorld) MakeCharacter() Hero {
	return &Wizard{
		Name: ww.PlayerName,
	}
}

func (ww *WizardWorld) MakeObstacle() Obstacle {
	return &Ork{}
}

func (ww *WizardWorld) String() string {
	return "\n\n\t------ Wizard World ------"
}

// GameEnvironment 游戏入口 🕹
type GameEnvironment struct {
	hero     Hero
	obstacle Obstacle
}

func NewGame(world World) GameEnvironment {
	fmt.Println(world)
	return GameEnvironment{
		hero:     world.MakeCharacter(),
		obstacle: world.MakeObstacle(),
	}
}

// Play 开始游戏
func (g GameEnvironment) Play() {
	g.hero.InteractWith(g.obstacle)
}
