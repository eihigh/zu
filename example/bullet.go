package main

import "github.com/eihigh/zu"

var (
	bullets []*bullet
)

type bullet struct {
	pos, vec complex128
}

func spawn() {
	if zu.Now()%5 == 0 {
		pos := complex(100, 100)
		vec := complex(0, 1)
		b := &bullet{
			pos: pos, vec: vec,
		}
		bullets = append(bullets, b)
	}
}

type spawnSystem struct {
	from zu.Time
}

func newSpawnSystem() *spawnSystem {
	s := &spawnSystem{
		from: zu.Now(),
	}
	zu.PushSystems(s)
	return s
}

func (s *spawnSystem) Update() {

}
