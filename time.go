package zu

type Time int

func Now() Time {
	return Time(count)
}

type Timer struct {
	from Time
	Max  Time
}

func NewTimer(from Time) Timer {
	return Timer{
		from: from,
	}
}

func (t Timer) Elapsed() Time {
	return Now() - t.from
}

func (t Timer) Ratio() float64 {
	if t.Max == 0 {
		return 0
	}
	return float64(t.Elapsed()) / float64(t.Max)
}

func (t Timer) Once(f func()) {
	if t.Elapsed() == 0 {
		f()
	}
}
