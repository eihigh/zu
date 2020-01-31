package zu

import (
	"math"
	"strconv"
)

// Pi2 returns math.Pi * 2 * value.
func Pi2(value float64) float64 {
	return math.Pi * 2 * value
}

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return "(" + strconv.FormatFloat(p.X, 'g', -1, 64) + "," + strconv.FormatFloat(p.Y, 'g', -1, 64) + ")"
}

// Add returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Sub returns the vector p-q.
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// Mul returns the vector p*k.
func (p Point) Mul(k float64) Point {
	return Point{p.X * k, p.Y * k}
}

// Div returns the vector p/k.
func (p Point) Div(k float64) Point {
	return Point{p.X / k, p.Y / k}
}

// Eq reports whether p and q are equal.
func (p Point) Eq(q Point) bool {
	return p == q
}

// Pt is shorthand for Point{X, Y}.
func Pt(X, Y float64) Point {
	return Point{X, Y}
}

type Rect struct {
	Min, Max Point
}

// String returns a string representation of r like "(3,4)-(6,5)".
func (r Rect) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Dx returns r's width.
func (r Rect) Dx() int {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r Rect) Dy() int {
	return r.Max.Y - r.Min.Y
}

// Size returns r's width and height.
func (r Rect) Size() Point {
	return Point{
		r.Max.X - r.Min.X,
		r.Max.Y - r.Min.Y,
	}
}

// Add returns the rectangle r translated by p.
func (r Rect) Add(p Point) Rect {
	return Rect{
		Point{r.Min.X + p.X, r.Min.Y + p.Y},
		Point{r.Max.X + p.X, r.Max.Y + p.Y},
	}
}

// Sub returns the rectangle r translated by -p.
func (r Rect) Sub(p Point) Rect {
	return Rect{
		Point{r.Min.X - p.X, r.Min.Y - p.Y},
		Point{r.Max.X - p.X, r.Max.Y - p.Y},
	}
}
