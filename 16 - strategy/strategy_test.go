package strategy

import (
	"math/rand"
	"testing"
	"time"
)

func TestRemoveRepByMap(t *testing.T) {
	slc := make([]int, 100000)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100000; i++ {
		slc[i] = r.Intn(100000)
	}
	result := RemoveRep(slc)
	t.Log(len(result))
}

func TestRemoveRepByLoop(t *testing.T) {
	slc := make([]int, 1000)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		slc[i] = r.Intn(1000)
	}
	result := RemoveRep(slc)
	t.Log(len(result))
}
