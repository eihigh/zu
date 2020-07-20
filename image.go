package zu

import (
	"image"
	"io"

	"github.com/hajimehoshi/ebiten"
)

var (
	screen *Image
	dio    = ebiten.DrawImageOptions{}
)

type Image = ebiten.Image

func DecodeImage(r io.Reader) (*Image, error) {
	src, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	img, _ := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	return img, nil
}

func NewEmptyImage(width, height int) *Image {
	img, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	return img
}

func SubImage(i *Image, r Rectangle) *Image {
	return i.SubImage(r.Image()).(*Image)
}

// Copy copies the src image to the dst while converting it with options.
func Copy(dst, src *Image, opts ...CopyOption) {
	dio.GeoM.Reset()
	dio.ColorM.Reset()
	dio.CompositeMode = ebiten.CompositeModeSourceOver
	if dst == nil {
		dst = screen
	}
	for _, opt := range opts {
		opt.applyCopyOption(src, &dio)
	}
	dst.DrawImage(src, &dio)
}
