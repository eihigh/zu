package main

import (
	_ "image/png"
	"net/http"

	"github.com/eihigh/zu"
	"github.com/hajimehoshi/ebiten"
)

var (
	spc = zu.Input{
		Keys: []ebiten.Key{ebiten.KeyA},
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
		if quit.IsDown() {
			break
		}
		if spc.OnDown() {
			popup()
		}
	}
	return nil
}

func popup() {
	v := newPopupView()
	zu.PushView(v)
	defer zu.WillRemoveView(v)

	s := newSpawnSystem()
	zu.PushSystems(s)
	defer zu.RemoveSystems(s)

	for zu.Next() {
		if v.opened() {
			break
		}
	}

	for zu.Next() {
		if spc.OnDown() {
			break
		}
		if left.IsDown() {
			v.x--
		}
		if right.IsDown() {
			v.x++
		}
	}
}

/*
func battle() {
	v := newBattleView()
	zu.PushView(v)
	defer zu.WillRemoveView(v)

	// select party action
	for zu.Next() {
		escape := chooseAction(v)
		if escape {
			if ok := escapeBattle(); ok {
				// when successfully escaped
				v.escaped = true
				return
			}
		}
	}
}
*/

type popupView struct {
	x            float64
	from, closed zu.Time
}

func newPopupView() *popupView {
	return &popupView{
		from: zu.Now(),
	}
}

func (v *popupView) opened() bool {
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
