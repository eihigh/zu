package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	vw, vh = 800, 600
)

type app struct {
}

func newApp() (*app, error) {
	a := &app{}
	return a, nil
}

func (a *app) Update() error {
	return nil
}

func (a *app) Draw(screen *ebiten.Image) {

}

func (a *app) Layout(ow, oh int) (int, int) {
	return vw, vh
}

func main() {
	ebiten.SetWindowSize(vw, vh)
	app, err := newApp()
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(app); err != nil {
		panic(err)
	}
}
