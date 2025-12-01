package helpers

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	currTicks   int
	targetTicks int
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currTicks:   0,
		targetTicks: int(d.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (t *Timer) Update() {
	if t.currTicks < t.targetTicks {
		t.currTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currTicks = 0
}
