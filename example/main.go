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

	left = zu.Input{
		Keys: []ebiten.Key{ebiten.KeyLeft},
	}
	right = zu.Input{
		Keys: []ebiten.Key{ebiten.KeyRight},
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
	v := newPopupView()
	zu.PushView(v)
	defer zu.PopView()

	for zu.Next() {
		if left.IsDown() {
			v.x--
		}
		if right.IsDown() {
			v.x++
		}
		if spc.IsDown() {
			break
		}
	}
}

type popupView struct {
	x    float64
	from zu.Time
}

func newPopupView() *popupView {
	return &popupView{
		from: zu.Now(),
	}
}

func (v *popupView) View() {
	zu.DrawImage(eimg, nil, zu.Translate(v.x, 0))
	zu.NewTimer(v.from).Once(func() {
		zu.DrawImage(eimg, nil, zu.Translate(100, 100))
	})
}
