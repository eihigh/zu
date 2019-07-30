package zu

import "github.com/hajimehoshi/ebiten"

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
