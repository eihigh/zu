package zu

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type CopyOption interface {
	applyCopyOption(*Image, *ebiten.DrawImageOptions)
}

type PrintOption interface {
	applyPrintOption(*printOption)
}

// Move represents a translation.
func Move(x, y float64) Mover { return Mover(Pt(x, y)) }

// MovePt represents a translation with the Point.
func MovePt(p Point) Mover { return Mover(p) }

// Rel represents a translation relative to the image size.
func Rel(x, y float64) Reler { return Reler(Pt(x, y)) }

// Center translates the image that moves it to the center,
// which is equivalent to Rel(-0.5, -0.5).
func Center() Reler { return Reler{-0.5, -0.5} }

// Scale scales the image.
func Scale(x, y float64) Scaler { return Scaler(Pt(x, y)) }

// Scale scales the image with the Point.
func ScalePt(p Point) Scaler { return Scaler(p) }

// Fit scales the image to fit the specified size.
func Fit(w, h int) Fitter { return Fitter{float64(w), float64(h)} }

// Rotate represents a rotation.
func Rotate(theta float64) Rotater { return Rotater(theta) }

func Color(clr color.Color) Colorizer { return Colorizer{clr} }
func RGBA(r, g, b, a uint8) Colorizer { return Colorizer{color.RGBA{r, g, b, a}} }

func BlendLighter() Blender { return Blender(ebiten.CompositeModeLighter) }

func AlignCenter() Aligner { return Aligner(alignCenter) }
func AlignLeft() Aligner   { return Aligner(alignLeft) }
func AlignRight() Aligner  { return Aligner(alignRight) }

type Mover Point

func (m Mover) applyCopyOption(_ *Image, opt *ebiten.DrawImageOptions) {
	opt.GeoM.Translate(m.X, m.Y)
}

func (m Mover) applyPrintOption(opt *printOption) {
	opt.x += m.X
	opt.y += m.Y
}

type Reler Point

func (r Reler) applyCopyOption(i *Image, opt *ebiten.DrawImageOptions) {
	b := i.Bounds()
	opt.GeoM.Translate(float64(b.Dx())*r.X, float64(b.Dy())*r.Y)
}

func (r Reler) applyPrintOption(opt *printOption) {
	opt.rx = r.X
	opt.ry = r.Y
}

type Scaler Point

func (s Scaler) applyCopyOption(_ *Image, opt *ebiten.DrawImageOptions) {
	opt.GeoM.Scale(s.X, s.Y)
}

type Fitter Point

func (f Fitter) applyCopyOption(i *Image, opt *ebiten.DrawImageOptions) {
	b := i.Bounds()
	sx := f.X / float64(b.Dx())
	sy := f.Y / float64(b.Dy())
	opt.GeoM.Scale(sx, sy)
}

type Rotater float64

func (r Rotater) applyCopyOption(_ *Image, opt *ebiten.DrawImageOptions) {
	opt.GeoM.Rotate(float64(r))
}

type Blender ebiten.CompositeMode

func (b Blender) applyCopyOption(_ *Image, opt *ebiten.DrawImageOptions) {
	opt.CompositeMode = ebiten.CompositeMode(b)
}

type Aligner align // text alignment

func (a Aligner) applyPrintOption(opt *printOption) {
	opt.align = align(a)
}

type Colorizer struct {
	clr color.Color
}

func (c Colorizer) applyCopyOption(_ *Image, opt *ebiten.DrawImageOptions) {
	opt.ColorM.Apply(c.clr)
}

func (c Colorizer) applyPrintOption(opt *printOption) {
	opt.clr = c.clr
}

type CopyOptions []CopyOption

func (opts CopyOptions) applyCopyOption(i *Image, dio *ebiten.DrawImageOptions) {
	for _, opt := range opts {
		opt.applyCopyOption(i, dio)
	}
}
