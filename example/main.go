package main

import (
	_ "image/png"
	"io"
	"os"

	"github.com/eihigh/zu"
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	pos, spd, acc zu.Point
}

var (
	e      *ebiten.Image
	pl     Player
	grav   = 0.3  // 重力加速度
	abyssY = 220. // 下限Y

	right = zu.Input{
		Keys: []ebiten.Key{ebiten.KeyRight},
	}
	left = zu.Input{
		Keys: []ebiten.Key{ebiten.KeyLeft},
	}
	quit = zu.Input{
		Keys: []ebiten.Key{ebiten.KeyQ},
	}

	q = zu.Pt(0, 0)
	t zu.Timer
)

func main() {
	f, _ := os.Open("./e.png")
	e, _ = zu.DecodeImage(f)
	zu.Run(update, draw, 320, 240, 2, "title")
}

func update() error {
	if quit.IsDown() {
		return io.EOF
	}
	pl.acc.Y = grav

	switch {
	case right.IsDown():
		pl.spd.X = 4
	case left.IsDown():
		pl.spd.X = -4
	default:
		pl.spd.X = 0
	}

	pl.spd = pl.spd.Add(pl.acc)
	pl.pos = pl.pos.Add(pl.spd)
	if pl.pos.Y > abyssY {
		pl.pos.Y = abyssY
	}

	t.Update()
	t.Repeat(0, 60, func(u zu.Timer) {
		q.X = zu.Wave(0, 100, u.Ratio())
	})
	return nil
}

func draw() {
	zu.Copy(nil, e, zu.Center(e), zu.MoveP(pl.pos))
	zu.Copy(nil, e, zu.MoveP(q))
}
