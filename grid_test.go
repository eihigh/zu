package zu_test

import (
	"testing"

	"github.com/eihigh/zu"
	"github.com/hajimehoshi/ebiten"
)

func TestHoge(t *testing.T) {
	t.Log("a")
	g := zu.NewGrid(
		"     32px*8",
		"1fr  nav ",
		"20px ",
	)
	t.Log(g.Lookup("nav"))
	var img *ebiten.Image
	img.SubImage(g.At(0, 0).Image())
}
