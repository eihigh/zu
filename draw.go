package zu

import "github.com/hajimehoshi/ebiten"

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

// var (
// 	opt = &ebiten.DrawImageOptions{}
// )
//
// type DrawOptionFn func()
//
// func DrawImage(src, dst *ebiten.Image, fns ...DrawOptionFn) {
// 	opt.GeoM.Reset()
// 	for _, fn := range fns {
// 		fn()
// 	}
// 	if dst == nil {
// 		dst = screen
// 	}
// 	dst.DrawImage(src, opt)
// }
//
// func Translate(x, y float64) DrawOptionFn {
// 	return func() {
// 		opt.GeoM.Translate(x, y)
// 	}
// }
