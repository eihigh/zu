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
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return 0, 0
	}

	if dst == nil {
		dst = screen
	}

	p := &printOption{clr: color.Gray{128}}
	for _, opt := range opts {
		opt.applyPrintOption(p)
	}

	ws := make([]fixed.Int26_6, 0, len(lines))
	wholeW := fixed.Int26_6(0)

	h := fface.Metrics().Height
	wholeH := h.Mul(fixed.I(len(lines)))

	for _, line := range lines {
		_, w := font.BoundString(fface, line)
		if wholeW < w {
			wholeW = w
		}
		ws = append(ws, w)
	}

	bounds, _, _ := fface.GlyphBounds('M')
	offsetX := -bounds.Min.X
	offsetY := -bounds.Min.Y

	x := p.x + float64(wholeW.Round())*p.rx + float64(offsetX.Round())
	y := p.y + float64(wholeH.Round())*p.ry + float64(offsetY.Round())

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
