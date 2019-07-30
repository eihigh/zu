package main

import (
	_ "image/png"
	"net/http"

	"github.com/eihigh/zu"
	"github.com/hajimehoshi/ebiten"
)

var (
	spc = zu.Input{
		Keys: []ebiten.Key{ebiten.KeySpace},
	}
	quit = zu.Input{
		Keys: []ebiten.Key{ebiten.KeyQ},
	}

	fs   = http.Dir(".")
	eimg *ebiten.Image
)

func main() {
	f, err := fs.Open("e.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	eimg, err = zu.DecodeImage(f)
	if err != nil {
		panic(err)
	}

	zu.Main(app)
}

func app() error {
	for zu.Next() {
		if spc.IsDown() {
			popup()
		}
		if quit.IsDown() {
			break
		}
	}
	return nil
}

func popup() {
	zu.PushViewFunc(popupView)
	defer zu.PopView()

	for zu.Next() {
		if spc.IsDown() {
			break
		}
	}
}

func popupView() {
	zu.DrawImage(eimg, nil)
}
