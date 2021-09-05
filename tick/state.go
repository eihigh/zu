package tick

// State represents a single Tick and a single state string.
type State struct {
	Tick
	state string
}

func NewState(initial string) State {
	return State{state: initial}
}

// State returns the current state.
func (s *State) State() string { return s.state }

// Change changes the state and resets the Tick.
func (s *State) Change(name string) {
	s.state = name
	s.Reset()
}
