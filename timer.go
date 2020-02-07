package zu

type Timer struct {
	count1   int // count + 1
	min, max int
}

func (t *Timer) Update() {
	t.count1++
}

func (t Timer) Count() int {
	return t.count1 - 1 - t.min
}

func (t *Timer) setCount(c int) {
	t.count1 = c + 1
}

func (t Timer) Span(start, end int, f func(Timer)) Timer {
	if start <= t.Count() && t.Count() < end {
		u := Timer{
			count1: t.count1,
			min:    start,
			max:    end,
		}
		f(u)
	}
	return t
}

func (t Timer) Repeat(offset, duration int, f func(Timer)) {
	e := (t.Count() + offset) % duration
	u := Timer{
		min: 0,
		max: duration,
	}
	u.setCount(e)
	f(u)
}

func (t Timer) Every(offset, interval, f func(Timer)) {

}

func (t Timer) ElapsedCount() int {
	return t.Count() - t.min
}

func (t Timer) Ratio() float64 {
	if t.max-t.min == 0 {
		return 0
	}
	return float64(t.Count()-t.min) / float64(t.max-t.min)
}

func _() {
	var t Timer
	t.Update()
	t.Span(0, 100, func(u Timer) {
		if u.Ratio() < 0.5 {

		}
	})
	t.Repeat(5, 0, func(u Timer) {
		u.Repeat(2, 0, func(_ Timer) {

		})
	})
}
