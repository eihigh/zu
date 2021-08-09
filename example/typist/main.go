package main

import (
	"image/color"

	"github.com/eihigh/zu/typist"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	vw, vh = 800, 600
)

type app struct {
	mplus24 font.Face
}

func newApp() (*app, error) {
	a := &app{}

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return nil, err
	}
	a.mplus24, err = opentype.NewFace(tt, &opentype.FaceOptions{Size: 24, DPI: 72, Hinting: font.HintingFull})
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *app) Update() error {
	return nil
}

func (a *app) Draw(screen *ebiten.Image) {
	s := "ゴーシュは町の活動写真館でセロを弾く係りでした。\nけれどもあんまり上手でないという評判でした。\n上手でないどころではなく実は仲間の楽手のなかではいちばん下手でしたから、いつでも楽長にいじめられるのでした。"
	maxW, lines := typist.Measure(a.mplus24, s, 600)
	met := a.mplus24.Metrics()
	dot := met.Ascent.Round()
	height := len(lines) * met.Height.Round()

	// draw on left-top
	x, y := 0, 0
	y += dot
	for _, line := range lines {
		text.Draw(screen, line.Text, a.mplus24, x, y, color.White)
		y += met.Height.Round()
	}

	// draw on center
	x, y = vw/2, vh/2
	x -= maxW / 2
	y += dot
	y -= height / 2
	for _, line := range lines {
		ofs := (maxW - line.Width) / 2
		text.Draw(screen, line.Text, a.mplus24, x+ofs, y, color.White)
		y += met.Height.Round()
	}

	// draw on right-bottom
	x, y = vw, vh
	y += dot
	y -= height
	for _, line := range lines {
		text.Draw(screen, line.Text, a.mplus24, x-line.Width, y, color.White)
		y += met.Height.Round()
	}

	ebitenutil.DrawLine(screen, 0, 300, 800, 300, color.White)
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
