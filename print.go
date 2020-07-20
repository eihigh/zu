package zu

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type align int

const (
	alignLeft align = iota
	alignCenter
	alignRight
)

type printOption struct {
	x, y   float64
	rx, ry float64
	clr    color.Color
	align  align
}

// Print prints the text on the dst with the options.
func Print(dst *ebiten.Image, s string, fface font.Face, opts ...PrintOption) (width, height int) {
	if dst == nil {
		dst = screen
	}

	p := &printOption{clr: color.Gray{128}}
	for _, opt := range opts {
		opt.applyPrintOption(p)
	}

	bounds, _, _ := fface.GlyphBounds('.')
	offsetX := -bounds.Min.X
	offsetY := -bounds.Min.Y
	h := bounds.Max.Y - bounds.Min.Y

	lines := strings.Split(s, "\n")
	ws := make([]fixed.Int26_6, 0, len(lines))
	wholeW := fixed.Int26_6(0)
	wholeH := h.Mul(fixed.I(len(lines))) // TODO consider line-height

	for _, line := range lines {
		w := font.MeasureString(fface, line)
		if wholeW < w {
			wholeW = w
		}
		ws = append(ws, w)
	}

	x := p.x + float64(offsetX.Round()) + float64(wholeW.Round())*p.rx
	y := p.y + float64(offsetY.Round()) + float64(wholeH.Round())*p.ry

	for i, line := range lines {
		w := ws[i]
		u := x
		switch p.align {
		case alignCenter:
			u += float64((wholeW - w).Round()) / 2
		case alignRight:
			u += float64((wholeW - w).Round())
		}
		text.Draw(dst, line, fface, int(u+0.5), int(y+0.5), p.clr)
		y += float64(h.Round())
	}

	return wholeW.Round(), wholeH.Round()
}
