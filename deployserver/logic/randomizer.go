package logic

import (
	"math/rand"
	"sync"
	"time"
)

type randomizer struct {
	rand *rand.Rand
	mu sync.Mutex
}

func newRandomizer() *randomizer {
	return &randomizer{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *randomizer) Intn(n int) (v int) {
	r.mu.Lock()
	v = r.rand.Intn(n)
	r.mu.Unlock()
	return
}