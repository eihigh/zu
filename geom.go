package zu

import (
	"image"
	"math"
	"strconv"
)

func Wave(min, max, rate float64) float64 {
	sin := math.Sin(rate * math.Pi * 2)
	return min + (max-min)*(sin+1)/2
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

// In reports whether p is in r.
func (p Point) In(r Rect) bool {
	return r.Min.X <= p.X && p.X < r.Max.X &&
		r.Min.Y <= p.Y && p.Y < r.Max.Y
}

// Equal reports whether p and q are equal.
func (p Point) Equal(q Point) bool {
	return p == q
}

func (p Point) Dot(q Point) float64 {
	return p.X*q.Y + p.Y*q.Y
}

func (p Point) Cross(q Point) float64 {
	return p.X*q.Y - p.Y*q.X
}

// Pt is shorthand for Point{X, Y}.
func Pt(X, Y float64) Point {
	return Point{X, Y}
}

type Line struct {
	Start, End Point
}

// Dx returns l's width.
func (l *Line) Dx() float64 {
	return l.End.X - l.Start.X
}

// Dy returns l's height.
func (l *Line) Dy() float64 {
	return l.End.Y - l.Start.Y
}

func (l *Line) Length() float64 {
	return (l.Start.X-l.End.X)*2 + (l.Start.Y-l.End.Y)*2
}

func (l *Line) Angle() float64 {
	return math.Atan2(l.Dy(), l.Dx())
}

func _() {
	var l Line
	l.Angle()
	l.Length()
	l.Intersect(l)
}

func (l *Line) Intersect(m Line) bool {
	lsms := m.Start.Sub(l.Start)
	lsme := m.End.Sub(l.Start)
	msls := l.Start.Sub(m.Start)
	msle := l.End.Sub(m.Start)
	return lsms.Cross(lsme) < 0 && msls.Cross(msle) < 0
}

func (l *Line) Bounds() Rect {
	return Rect{
		Min: l.Start,
		Max: l.Start.Add(Pt(l.Dx(), l.Dy())),
	}
}

type Rect struct {
	Min, Max Point
}

// String returns a string representation of r like "(3,4)-(6,5)".
func (r Rect) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Dx returns r's width.
func (r Rect) Dx() float64 {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r Rect) Dy() float64 {
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

func (r Rect) Image() image.Rectangle {
	return image.Rect(int(r.Min.X), int(r.Min.Y), int(r.Max.X), int(r.Max.Y))
}

// Intersection returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r Rect) Intersection(s Rect) Rect {
	if r.Min.X < s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y < s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X > s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y > s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	// Letting r0 and s0 be the values of r and s at the time that the method
	// is called, this next line is equivalent to:
	//
	// if max(r0.Min.X, s0.Min.X) >= min(r0.Max.X, s0.Max.X) || likewiseForY { etc }
	// if r.Empty() {
	// 	return Rect{}
	// }
	return r
}
