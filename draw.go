package zu

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

var (
	dio = &ebiten.DrawImageOptions{}
	dto = &ebiten.DrawTrianglesOptions{}
)

type DrawOptionFn func()

func Move(x, y float64) DrawOptionFn {
	return func() {
		dio.GeoM.Translate(x, y)
	}
}

func MoveP(p Point) DrawOptionFn {
	return func() {
		dio.GeoM.Translate(p.X, p.Y)
	}
}

func Center(i image.Image) DrawOptionFn {
	return func() {
		s := i.Bounds().Size()
		dio.GeoM.Translate(float64(-s.X)/2, float64(-s.Y)/2)
	}
}

func Scale(x, y float64) DrawOptionFn {
	return func() {
		dio.GeoM.Scale(x, y)
	}
}

func ScaleP(p Point) DrawOptionFn {
	return func() {
		dio.GeoM.Scale(p.X, p.Y)
	}
}

func Rotate(a float64) DrawOptionFn {
	return func() {
		dio.GeoM.Rotate(a)
	}
}

func Copy(src, dst *ebiten.Image, opts ...DrawOptionFn) {
	dio.GeoM.Reset()
	for _, opt := range opts {
		opt()
	}
	if dst == nil {
		dst = Screen
	}
	dst.DrawImage(src, dio)
}
