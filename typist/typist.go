package typist

import (
	"strings"

	"github.com/rivo/uniseg"
	"golang.org/x/image/font"
)

type Line struct {
	Text  string
	Width int
}

func Measure(f font.Face, s string, wrapWidth int) (maxWidth int, lines []*Line) {
	if wrapWidth >= 0 {
		return measureWithWrap(f, s, wrapWidth)
	}
	return measureWithoutWrap(f, s)
}

func measureWithWrap(f font.Face, s string, wrapWidth int) (maxWidth int, lines []*Line) {
	ss := strings.Split(s, "\n")
	lines = make([]*Line, 0, len(ss))
	for _, s := range ss {
		g := uniseg.NewGraphemes(s)
		l := &Line{}
		lines = append(lines, l)
		for g.Next() {
			_, gadv := font.BoundString(f, g.Str())
			gw := gadv.Round()
			if l.Width > 0 && l.Width+gw > wrapWidth {
				l = &Line{Text: g.Str(), Width: gw}
				lines = append(lines, l)
				continue
			}
			l.Text += g.Str()
			l.Width += gw
		}
	}
	for _, line := range lines {
		if line.Width > maxWidth {
			maxWidth = line.Width
		}
	}
	return
}

func measureWithoutWrap(f font.Face, s string) (maxWidth int, lines []*Line) {
	ss := strings.Split(s, "\n")
	lines = make([]*Line, 0, len(ss))
	for _, s := range ss {
		_, adv := font.BoundString(f, s)
		w := adv.Round()
		if w > maxWidth {
			maxWidth = w
		}
		lines = append(lines, &Line{Text: s, Width: w})
	}
	return
}
