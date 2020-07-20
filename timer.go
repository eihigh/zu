package zu

var (
	tick = 0
)

func Now() int {
	return tick
}

type Timer struct {
	from, to int
}

func NewTimer() Timer {
	return Timer{from: Now() + 1}
}

func (t *Timer) Reset() {
	t.from = Now() + 1
}

func (t Timer) Elapsed() int {
	return Now() - t.from
}

func (t Timer) Elapsedf() float64 {
	return float64(t.Elapsed())
}

func (t Timer) Rate() float64 {
	if t.to == 0 {
		return 0
	}
	return float64(Now()-t.from) / float64(t.to-t.from)
}

func (t Timer) Span(offset, duration int, f func(Timer)) Timer {
	// endless
	if duration < 0 {
		u := Timer{from: t.from + offset}
		if t.Elapsed() < offset {
			return u
		}
		f(u)
		return u
	}

	u := Timer{from: t.from + offset + duration}
	if t.Elapsed() < offset {
		return u
	}
	if offset+duration <= t.Elapsed() {
		return u
	}
	v := Timer{
		from: t.from + offset,
		to:   t.from + offset + duration,
	}
	f(v)
	return u
}

func (t Timer) Once(f func()) {
	if t.Elapsed() == 0 {
		f()
	}
}

func (t Timer) Repeat(offset, duration int, f func(Timer)) {
	if t.Elapsed() < offset {
		return
	}
	e := int((t.Elapsed() - offset) / duration)
	from := t.from + offset + e*duration
	f(Timer{
		from: from,
		to:   from + duration,
	})
}

func (t Timer) Every(offset, interval int, f func()) {
	if interval == 0 {
		return
	}
	if t.Elapsed() < offset {
		return
	}
	if (t.Elapsed()-offset)%interval != 0 {
		return
	}
	f()
}

type State struct {
	Timer
	state string

	reservedState string
	reserveFrom   int
}

func NewState(initializer string) State {
	s := State{}
	s.from = Now()
	s.state = initializer
	return s
}

func (s *State) Reserve(state string) {
	s.reservedState = state
	s.reserveFrom = Now() + 1
}

func (s *State) Continue(next string) {
	s.state = next
}

func (s *State) Get() string {
	if s.reserveFrom != 0 && s.reserveFrom <= Now() {
		s.state = s.reservedState
		s.from = Now()
		s.reserveFrom = 0
	}
	return s.state
}
