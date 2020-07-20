// +build example

package main

import (
	"image/color"

	"github.com/eihigh/zu"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

const (
	vw = 800.0
	vh = 600.0
)

type app struct {
	mplus font.Face
}

func newApp() (*app, error) {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return nil, err
	}
	const dpi = 72
	return &app{
		mplus: truetype.NewFace(tt, &truetype.Options{
			Size:    48,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}, nil
}

func (a *app) Update() error {
	return nil
}

func (a *app) Draw() {
	text := "こんにちは\n世界。"
	zu.Print(nil, text, a.mplus, zu.AlignLeft(), zu.Color(color.White))
	zu.Print(nil, text, a.mplus, zu.AlignCenter(), zu.Move(vw/2, vh/2), zu.Center(), zu.Color(color.White))
	zu.Print(nil, text, a.mplus, zu.AlignRight(), zu.Move(vw, vh), zu.Rel(-1, -1), zu.Color(color.White))
}

func (a *app) Layout(w, h int) (int, int) {
	return vw, vh
}

func main() {
	ebiten.SetWindowSize(vw, vh)
	a, err := newApp()
	if err != nil {
		panic(err)
	}
	zu.Run(a)
}
