package zu

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Input struct {
	Keys []Key
}

func (i *Input) IsDown() bool {
	for _, k := range i.Keys {
		if ebiten.IsKeyPressed(ebiten.Key(k)) {
			return true
		}
	}
	return false
}

func (i *Input) OnDown() bool {
	for _, k := range i.Keys {
		if inpututil.IsKeyJustPressed(ebiten.Key(k)) {
			return true
		}
	}
	return false
}
