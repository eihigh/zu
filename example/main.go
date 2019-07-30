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
	for !quit.IsDown() {
		zu.Next()
		if spc.IsDown() {
			popup()
		}
	}
	return nil
}

func popup() {
	v := newPopupView()
	zu.PushView(v)
	defer zu.WillRemoveView(v)

	for !v.Opened() {
		zu.Next()
	}

	for !spc.IsDown() {
		zu.Next()
		if left.IsDown() {
			v.x--
		}
		if right.IsDown() {
			v.x++
		}
	}
}

type popupView struct {
	x            float64
	from, closed zu.Time
}

func newPopupView() *popupView {
	return &popupView{
		from: zu.Now(),
	}
}

func (v *popupView) Opened() bool {
	return true
}

func (v *popupView) Close() {
	v.closed = zu.Now()
}

func (v *popupView) Done() bool {
	return zu.Now()-v.closed > 60
}

func (v *popupView) View() {
	zu.DrawImage(eimg, nil, zu.Translate(v.x, 0))
	zu.NewTimer(v.from).Once(func() {
		zu.DrawImage(eimg, nil, zu.Translate(100, 100))
	})

	if v.closed != 0 {
		zu.DrawImage(eimg, nil, zu.Translate(float64(zu.Now()-v.closed), 200))
	}
}
