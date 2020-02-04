package zu

type Timer struct {
	count    int
	min, max int
}

func (t Timer) Count() int {
	return t.count - 1
}

func (t Timer) Span(start, end int, f func()) {
	if start <= t.Count() && t.Count() < end {
		f()
	}
}

func (t Timer) ElapsedRatio() float64 {
	return float64(t.Count()) / float64(t.max)
}

func _() {
	t := Timer{}
	t.Span(0, 60, func() {
	})
}
