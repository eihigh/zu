// Tick represents a loop counter and can be used as an alternative to a coroutine in Go.
package tick

type Tick struct {
	t        int
	delta    int
	duration int
}

func (t Tick) Sub(u int) Tick {
	return Tick{
		t:        t.t - u,
		delta:    t.delta,
		duration: t.duration,
	}
}

// Delta returns the value of delta that was last passed to Advance.
func (t Tick) Delta() int { return t.delta }

// Advance advances the tick, and delta can take a value
// greater than 1.
func (t *Tick) Advance(delta int) {
	t.t += delta
	t.delta = delta
}

// Reset resets the tick and the delta.
func (t *Tick) Reset() {
	t.t = 0
	t.delta = 0
}

// Elapsed returns the elapsed tick.
func (t Tick) Elapsed() int {
	return t.t
}

// Elapsedf returns the elapsed tick as float64.
func (t Tick) Elapsedf() float64 {
	return float64(t.Elapsed())
}

// ElapsedRate returns the rate at which the tick has
// elapsed during the Span or Repeat period, in the range [0, 1).
func (t Tick) ElapsedRate() float64 {
	return float64(t.Elapsed()) / float64(t.duration)
}

// Rate is an alias for ElapsedRate.
func (t Tick) Rate() float64 { return t.ElapsedRate() }

// Span calls f only within the span.
// The returned value is a Tick that starts from 0 after
// the end of the span, and can be method chained to
// represent seqential spans.
func (t Tick) Span(start, duration int, f func(Tick)) Tick {

	// t=0       u=0       v=0
	// |---------|---------|
	//           start     start+duration

	u := Tick{
		t:        t.t - start,
		delta:    t.delta,
		duration: duration,
	}
	v := Tick{
		t:     t.t - start - duration,
		delta: t.delta,
	}
	if u.t < 0 {
		return v
	}
	if 0 <= v.t {
		return v
	}
	f(u)

	return v
}

// From calls f from the start tick.
func (t Tick) From(start int, f func(Tick)) {
	if t.t < start {
		return
	}
	u := Tick{
		t:     t.t - start,
		delta: t.delta,
	}
	f(u)
}

// Once calls f only once after the start tick has passed.
func (t Tick) Once(start int, f func()) {
	if t.delta == 0 {
		if t.t == start {
			f()
		}
	} else {
		if t.t-t.delta < start && start <= t.t {
			f()
		}
	}
}

// Repeat continues to call function f after the start tick has elapsed,
// where t is the tick to be reset for each duration tick and
// n is the number of iterations.
func (t Tick) Repeat(start, duration int, f func(int, Tick)) {
	e := t.t - start
	if e < 0 {
		return
	}
	n := e / duration
	u := Tick{
		t:        t.t - start - n*duration,
		delta:    t.delta,
		duration: duration,
	}
	f(n, u)
}

// Every calls function f every duration tick from after the start tick.
// The argument n is the number of calls.
func (t Tick) Every(start, duration int, f func(n int)) {
	e := t.t - start
	if e < 0 {
		return
	}
	n := e / duration
	start += n * duration

	if t.delta == 0 {
		if t.t == start {
			f(n)
		}
	} else {
		if t.t-t.delta < start && start <= t.t {
			f(n)
		}
	}
}

func (t Tick) RepeatFor(start, duration, times int, f func(int, Tick)) Tick {
	return t.Span(start, duration*times, func(t Tick) {
		t.Repeat(0, duration, func(n int, t Tick) {
			f(n, t)
		})
	})
}

func (t Tick) EveryFor(start, duration, times int, f func(int)) Tick {
	return t.Span(start, duration*times, func(t Tick) {
		t.Every(0, duration, func(n int) {
			f(n)
		})
	})
}
