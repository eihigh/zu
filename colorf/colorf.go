package colorf

import "image/color"

const uint16max = 0xffff

func RGBAf(clr color.Color) (r, g, b, a float64) {
	R, G, B, A := clr.RGBA()
	return utof(R), utof(G), utof(B), utof(A)
}

// colorf.Color != color.Colorf
type Colorf interface {
	RGBAf() (r, g, b, a float64)
}

type RGBA struct {
	R, G, B, A float64
}

func (c RGBA) RGBA() (r, g, b, a uint32) {
	return ftou(c.R), ftou(c.G), ftou(c.B), ftou(c.A)
}

func (c RGBA) RGBAf() (r, g, b, a float64) {
	return c.R, c.G, c.B, c.A
}

type Alpha struct {
	A float64
}

func (c Alpha) RGBA() (r, g, b, a uint32) {
	A := ftou(c.A)
	return A, A, A, A
}

func (c Alpha) RGBAf() (r, g, b, a float64) {
	return c.A, c.A, c.A, c.A
}

type Gray struct {
	Y float64
}

func (c Gray) RGBA() (r, g, b, a uint32) {
	y := ftou(c.Y)
	return y, y, y, uint16max
}

func (c Gray) RGBAf() (r, g, b, a float64) {
	return c.Y, c.Y, c.Y, 1
}

func utof(n uint32) float64 {
	return float64(n) / uint16max
}

func ftou(f float64) uint32 {
	return uint32(f * uint16max)
}
