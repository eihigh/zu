package zu

import "github.com/hajimehoshi/ebiten"

type App interface {
	Update() error
	Draw()
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type game struct {
	app App
}

func (g *game) Update(s *ebiten.Image) error {
	screen = s
	tick++
	return g.app.Update()
}

func (g *game) Draw(_ *ebiten.Image) {
	g.app.Draw()
}

func (g *game) Layout(vw, vh int) (int, int) {
	return g.app.Layout(vw, vh)
}

func Run(a App) error {
	return ebiten.RunGame(&game{a})
}
