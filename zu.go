package zu

import (
	"fmt"
	"image"
	"io"

	"github.com/hajimehoshi/ebiten"
)

var (
	tick = make(chan struct{})
	tock = make(chan struct{})

	screen *ebiten.Image

	errFinished = fmt.Errorf("finished")

	done = make(chan struct{})
)

func Main(app func() error) error {
	var err error

	go func() {
		err = app()
		close(done)
	}()

	ebiten.Run(update, 320, 240, 2, "title")

	return err
}

func update(s *ebiten.Image) error {
	screen = s
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	select {
	case <-done:
		return errFinished
	default:
	}

	<-tick
	for _, v := range views {
		v.View()
	}
	tock <- struct{}{}
	return nil
}

func Next() bool {
	tick <- struct{}{}
	<-tock
	return true
}

func DecodeImage(r io.Reader) (*ebiten.Image, error) {
	src, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	img, _ := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	return img, nil
}

func DrawImage(src, dst *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	if dst == nil {
		dst = screen
	}
	dst.DrawImage(src, opt)
}
