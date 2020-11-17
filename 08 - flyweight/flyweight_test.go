package flyweight

import (
	"math/rand"
	"testing"
	"time"
)

func TestFlyweight(t *testing.T) {
	fm := map[int]Fruit{0: Apple, 1: Cherry, 2: Peach}
	am := map[int]AgeType{0: Young, 1: Mature, 2: Old}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fruitTrie := NewFruitTree(fm[r.Intn(3)], am[r.Intn(3)])
		fruitTrie.Render(r.Intn(100), r.Intn(100))
	}
}
