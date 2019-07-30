package zu

import "github.com/hajimehoshi/ebiten"

var (
	opt = &ebiten.DrawImageOptions{}
)

type DrawOptionFn func()

func DrawImage(src, dst *ebiten.Image, fns ...DrawOptionFn) {
	opt.GeoM.Reset()
	for _, fn := range fns {
		fn()
	}
	if dst == nil {
		dst = screen
	}
	dst.DrawImage(src, opt)
}

func Translate(x, y float64) DrawOptionFn {
	return func() {
		opt.GeoM.Translate(x, y)
	}
}
