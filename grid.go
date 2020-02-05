package zu

type Grid struct {
	Rect Rect
	Rows []Length
	Cols []Length
}

type Length struct {
}

func NewGrid(columns string, rows ...string) *Grid {
	return nil
}

func (g *Grid) At(x, y int) Rect {
	return Rect{}
}

func (g *Grid) AtIndex(i int) Rect {
	x := i % len(g.Cols)
	y := int(i / len(g.Cols))
	return g.At(x, y)
}

func (g *Grid) Range(ax, ay, bx, by int) Rect {
	return Rect{}
}

func (g *Grid) Lookup(name string) Rect {
	return Rect{}
}
