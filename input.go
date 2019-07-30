package zu

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Input struct {
	Keys []ebiten.Key
}

func (i *Input) IsDown() bool {
	for _, k := range i.Keys {
		if ebiten.IsKeyPressed(k) {
			return true
		}
	}
	return false
}

func (i *Input) OnDown() bool {
	for _, k := range i.Keys {
		if inpututil.IsKeyJustPressed(k) {
			return true
		}
	}
	return false
}
