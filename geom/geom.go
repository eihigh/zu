package geom

import (
	"image"
	"image/color"
	"math"
	"strconv"
)

// A Point is an X, Y coordinate pair.
type Point struct {
	X, Y float64
}

func PtFromRect(r, theta float64) Point {
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)
	return Point{x, y}
}

func (p Point) String() string {
	return "(" + strconv.FormatFloat(p.X, 'g', -1, 64) + "," + strconv.FormatFloat(p.Y, 'g', -1, 64) + ")"
}

func (p Point) LengthSq() float64 {
	return p.X*p.X + p.Y*p.Y
}

func (p Point) Normalize() Point {
	l := math.Sqrt(p.LengthSq())
	return Point{p.X / l, p.Y / l}
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
func (p Point) In(r Rectangle) bool {
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

func (p Point) Angle() float64 {
	return math.Atan2(p.Y, p.X)
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

func (l *Line) Intersects(m Line) bool {
	lsms := m.Start.Sub(l.Start)
	lsme := m.End.Sub(l.Start)
	msls := l.Start.Sub(m.Start)
	msle := l.End.Sub(m.Start)
	return lsms.Cross(lsme) < 0 && msls.Cross(msle) < 0
}

func (l *Line) Bounds() Rectangle {
	return Rectangle{
		Min: l.Start,
		Max: l.Start.Add(Pt(l.Dx(), l.Dy())),
	}
}

type Rectangle struct {
	Min, Max Point
}

func Rect(x0, y0, x1, y1 float64) Rectangle {
	return Rectangle{Pt(x0, y0), Pt(x1, y1)}
}

// String returns a string representation of r like "(3,4)-(6,5)".
func (r Rectangle) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Dx returns r's width.
func (r Rectangle) Dx() float64 {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r Rectangle) Dy() float64 {
	return r.Max.Y - r.Min.Y
}

// Size returns r's width and height.
func (r Rectangle) Size() Point {
	return Point{
		r.Max.X - r.Min.X,
		r.Max.Y - r.Min.Y,
	}
}

// Add returns the rectangle r translated by p.
func (r Rectangle) Add(p Point) Rectangle {
	return Rectangle{
		Point{r.Min.X + p.X, r.Min.Y + p.Y},
		Point{r.Max.X + p.X, r.Max.Y + p.Y},
	}
}

// Sub returns the rectangle r translated by -p.
func (r Rectangle) Sub(p Point) Rectangle {
	return Rectangle{
		Point{r.Min.X - p.X, r.Min.Y - p.Y},
		Point{r.Max.X - p.X, r.Max.Y - p.Y},
	}
}

// Inset returns the rectangle r inset by n, which may be negative. If either
// of r's dimensions is less than 2*n then an empty rectangle near the center
// of r will be returned.
func (r Rectangle) Inset(n float64) Rectangle {
	if r.Dx() < 2*n {
		r.Min.X = (r.Min.X + r.Max.X) / 2
		r.Max.X = r.Min.X
	} else {
		r.Min.X += n
		r.Max.X -= n
	}
	if r.Dy() < 2*n {
		r.Min.Y = (r.Min.Y + r.Max.Y) / 2
		r.Max.Y = r.Min.Y
	} else {
		r.Min.Y += n
		r.Max.Y -= n
	}
	return r
}

// Intersect returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r Rectangle) Intersect(s Rectangle) Rectangle {
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
	if r.Empty() {
		return Rectangle{}
	}
	return r
}

// Union returns the smallest rectangle that contains both r and s.
func (r Rectangle) Union(s Rectangle) Rectangle {
	if r.Empty() {
		return s
	}
	if s.Empty() {
		return r
	}
	if r.Min.X > s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y > s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X < s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y < s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	return r
}

// Empty reports whether the rectangle contains no points.
func (r Rectangle) Empty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

// Eq reports whether r and s contain the same set of points. All empty
// rectangles are considered equal.
func (r Rectangle) Eq(s Rectangle) bool {
	return r == s || r.Empty() && s.Empty()
}

// Overlaps reports whether r and s have a non-empty intersection.
func (r Rectangle) Overlaps(s Rectangle) bool {
	return !r.Empty() && !s.Empty() &&
		r.Min.X < s.Max.X && s.Min.X < r.Max.X &&
		r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}

// In reports whether every point in r is in s.
func (r Rectangle) In(s Rectangle) bool {
	if r.Empty() {
		return true
	}
	// Note that r.Max is an exclusive bound for r, so that r.In(s)
	// does not require that r.Max.In(s).
	return s.Min.X <= r.Min.X && r.Max.X <= s.Max.X &&
		s.Min.Y <= r.Min.Y && r.Max.Y <= s.Max.Y
}

// Canon returns the canonical version of r. The returned rectangle has minimum
// and maximum coordinates swapped if necessary so that it is well-formed.
func (r Rectangle) Canon() Rectangle {
	if r.Max.X < r.Min.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}
	if r.Max.Y < r.Min.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
	return r
}

// At implements the Image interface.
func (r Rectangle) At(x, y int) color.Color {
	if (Point{float64(x), float64(y)}).In(r) {
		return color.Opaque
	}
	return color.Transparent
}

// Bounds implements the Image interface.
func (r Rectangle) Bounds() image.Rectangle {
	return image.Rect(
		int(r.Min.X+0.5),
		int(r.Min.Y+0.5),
		int(r.Max.X+0.5),
		int(r.Max.Y+0.5),
	)
}

// ColorModel implements the Image interface.
func (r Rectangle) ColorModel() color.Model {
	return color.Alpha16Model
}

func (r Rectangle) Image() image.Rectangle {
	return image.Rect(int(r.Min.X+0.5), int(r.Min.Y+0.5), int(r.Max.X+0.5), int(r.Max.Y+0.5))
}
