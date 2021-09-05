// Package hsm provides a simple Hierarchical State Machine.
package hsm

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/eihigh/zu/tick"
)

var (
	ErrUndefinedState = errors.New("undefined state")
	ErrNotACleanPath  = errors.New("not a clean path")
)

// HSM is a simple hierarchical state machine.
type HSM struct {
	ticks   []tick.Tick
	history []string
	current string
	states  []*State
}

// State represents a state and its callbacks.
type State struct {
	Name                      string
	Enter, Exit, Update, Draw func(*HSM)
}

func NewHSM(states []*State, initial string, historyCap int) *HSM {
	mustClean(initial)
	for _, s := range states {
		mustClean(s.Name)
	}

	if historyCap < 1 {
		panic("hsm: historyCap must be greater than 0")
	}
	h := &HSM{states: states}

	// enter the root
	h.ticks = append(h.ticks, tick.Tick{})
	h.history = make([]string, 0, historyCap)
	h.current = "/"
	h.callEnter()

	// initialize
	h.Change(initial)

	return h
}

func (h *HSM) Current() string {
	return h.current
}

func (h *HSM) Prev() string {
	if len(h.history) >= 1 {
		return h.history[len(h.history)-1]
	}
	return ""
}

func (h *HSM) Tick() tick.Tick {
	return h.ticks[len(h.ticks)-1]
}

func (h *HSM) TickOf(target string) tick.Tick {
	mustClean(target)
	cur := h.current

	for i := 0; i < len(h.ticks); i++ {
		t := h.ticks[len(h.ticks)-1-i]
		if cur == target {
			return t
		}
		cur = path.Dir(cur)
	}

	panic(fmt.Errorf("hsm: state %q: %w", target, ErrUndefinedState))
}

func (h *HSM) Update() {
	for i := range h.ticks {
		h.ticks[i].Advance(1)
	}

	state := h.state(h.current)
	if state == nil {
		panic(fmt.Errorf("hsm: state %q: %w", h.current, ErrUndefinedState))
	}
	if state.Update == nil {
		return
	}
	state.Update(h)
}

func (h *HSM) Draw() {
	state := h.state(h.current)
	if state == nil {
		panic(fmt.Errorf("hsm: state %q: %w", h.current, ErrUndefinedState))
	}
	if state.Draw == nil {
		return
	}
	state.Draw(h)
}

func (h *HSM) Change(next string) {
	mustClean(next)

	cur := h.current

	// append cur into history before changing
	h.appendHistory(cur)

	// special case: cur == next
	if cur == next {
		h.callExit()
		h.ticks[len(h.ticks)-1].Reset()
		h.callEnter()
		return
	}

	// up & exit
	// /foo/bar/baz => /foo
	// => exit baz, exit bar
	for {
		if strings.HasPrefix(next, cur) {
			break
		}
		h.callExit()
		h.ticks = h.ticks[:len(h.ticks)-1]
		cur = path.Dir(cur)
		h.current = cur
	}

	// down & enter
	// /foo => /foo/x/y
	// => enter x, enter y
	names := []string{}
	for {
		if next == cur {
			break
		}
		names = append(names, path.Base(next))
		next = path.Dir(next)
	}

	for i := len(names) - 1; i >= 0; i-- {
		name := names[i]
		next = path.Join(next, name)
		h.current = next
		h.ticks = append(h.ticks, tick.Tick{})
		h.callEnter()
	}
}

func (h *HSM) appendHistory(s string) {
	if len(h.history) == cap(h.history) {
		l := len(h.history)
		copy(h.history[0:l-1], h.history[1:l])
		h.history[l-1] = s
	} else {
		h.history = append(h.history, s)
	}
}

func (h *HSM) state(name string) *State {
	for _, s := range h.states {
		if s.Name == name {
			return s
		}
	}
	return nil
}

func (h *HSM) callEnter() {
	state := h.state(h.current)
	if state == nil {
		return
	}
	if state.Enter == nil {
		return
	}
	state.Enter(h)
}

func (h *HSM) callExit() {
	state := h.state(h.current)
	if state == nil {
		return
	}
	if state.Exit == nil {
		return
	}
	state.Exit(h)
}

func mustClean(s string) {
	if s != path.Clean(path.Join("/", s)) {
		panic(fmt.Errorf("hsm: state %q: %w", s, ErrNotACleanPath))
	}
}
