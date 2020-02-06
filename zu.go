package zu

import (
	"image"
	"io"

	"github.com/hajimehoshi/ebiten"
)

var (
	Screen *ebiten.Image
	count  = 0
)

func Run(update func() error, draw func(), width, height int, scale float64, title string) error {
	u := func(s *ebiten.Image) error {
		Screen = s
		count++
		if err := update(); err != nil {
			return err
		}
		if !ebiten.IsDrawingSkipped() {
			draw()
		}
		return nil
	}

	return ebiten.Run(u, width, height, scale, title)
}

func DecodeImage(r io.Reader) (*ebiten.Image, error) {
	src, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	img, _ := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	return img, nil
}
