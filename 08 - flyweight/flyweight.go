package flyweight

import (
	"fmt"
)

type Fruit int

func (t Fruit) String() string {
	switch t {
	case Apple:
		return "apple"
	case Cherry:
		return "cherry"
	case Peach:
		return "peach"
	default:
		return ""
	}
}

const (
	Apple Fruit = iota
	Cherry
	Peach
)

type AgeType int

func (a AgeType) String() string {
	switch a {
	case Young:
		return "young"
	case Mature:
		return "mature"
	case Old:
		return "old"
	default:
		return ""
	}
}

const (
	Young AgeType = iota
	Mature
	Old
)

type Render interface {
	Render(int, int)
}

var treeCache = make(map[AgeType]*tree, 0)

type tree struct {
	AgeType AgeType
}

func newTree(ageType AgeType) *tree {
	if t, ok := treeCache[ageType]; ok {
		return t
	}
	treeCache[ageType] = &tree{ageType}
	return treeCache[ageType]
}

var fruitTreeCache = make(map[Fruit]*FruitTree, 0)

type FruitTree struct {
	*tree
	Fruit Fruit
}

func NewFruitTree(fruit Fruit, age AgeType) *FruitTree {
	if t, ok := fruitTreeCache[fruit]; ok {
		return t
	}

	tree := newTree(age)
	fruitTreeCache[fruit] = &FruitTree{tree: tree, Fruit: fruit}
	return fruitTreeCache[fruit]
}

func (t *FruitTree) Render(x, y int) {
	fmt.Printf("render a %s %s tree at (%d, %d)\n", t.AgeType, t.Fruit, x, y)
}

type OrnamentalTree struct{}
