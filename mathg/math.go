// Package mathg provides mathematical functions for graphics / games.
package mathg

import "math"

const (
	Tau     = math.Pi * 2
	Deg2Rad = math.Pi / 180
	Rad2Deg = 180 / math.Pi
)

func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func Yoyo(v float64) float64 {
	v *= 2
	n := int(v)
	if n%2 == 0 {
		return v - float64(n)
	}
	return float64(n) - v
}

// e := ease.InOutQuad(mathg.Yoyo(t.Rate()))
// rad := mathg.Lerp(-mathg.Tau/8, mathg.Tau/8, e)
// aim := vel.Rotate(rad)
