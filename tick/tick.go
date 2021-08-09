package tick

type Tick struct {
	t        int
	delta    int
	start    int
	duration int
}

func (t Tick) Get() int { return t.t }

func (t Tick) Delta() int { return t.delta }

func (t *Tick) Advance(delta int) {
	t.t += delta
	t.delta = delta
}

func (t Tick) Elapsed() int {
	return t.t - t.start
}

func (t Tick) Elapsedf() float64 {
	return float64(t.Elapsed())
}

func (t Tick) ElapsedRate() float64 {
	return float64(t.Elapsed()) / float64(t.duration)
}

// Span calls f only within the span.
func (t Tick) Span(start, duration int, f func(Tick)) Tick {

	// t.start   u.start   v.start
	// |---------|---------|
	//           start     start+duration

	start += t.start
	u := Tick{
		t:        t.t,
		delta:    t.delta,
		start:    start,
		duration: duration,
	}
	v := Tick{
		t:     t.t,
		delta: t.delta,
		start: start + duration,
	}
	if t.t < start {
		return v
	}
	if start+duration <= t.t {
		return v
	}
	f(u)

	return v
}

// From calls f from the start tick.
func (t Tick) From(start int, f func(Tick)) {
	start += t.start
	if t.t < start {
		return
	}
	u := Tick{
		t:     t.t,
		delta: t.delta,
		start: start,
	}
	f(u)
}

// Once calls f only once after the start tick has passed.
func (t Tick) Once(start int, f func()) {
	start += t.start
	if t.t < start {
		return
	}
	if start <= t.t-t.delta {
		return
	}
	f()
}

func (t Tick) Repeat(start, duration int, f func(int, Tick)) {
	start += t.start
	e := t.t - start
	if e < 0 {
		return
	}
	n := e / duration
	u := Tick{
		t:        t.t,
		delta:    t.delta,
		start:    start + n*duration,
		duration: duration,
	}
	f(n, u)
}

func (t Tick) Every(start, duration int, f func(int)) {
	start += t.start
	e := t.t - start
	if e < 0 {
		return
	}
	n := e / duration
	start += n * duration
	if t.t < start {
		return
	}
	if start <= t.t-t.delta {
		return
	}
	f(n)
}
