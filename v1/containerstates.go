package containerstates

const (
	Deleting State = "deleting"
	Deleted  State = "deleted"
	Creating State = "creating"
	Created  State = "created"
	Starting State = "starting"
	Up       State = "up"
	Stopping State = "stopping"
	Down     State = "down"
	Failed   State = "failed"
)

var (
	// StateOrder is used by Aggregate to determine the higher state.
	// Lower Index wins over higher Index
	StateOrder = []State{
		Failed,
		Down,
		Deleting,
		Deleted,
		Stopping,
		Creating,
		Starting,
		Up,
	}
)

type State string

func (s *State) String() string {
	return string(*s)
}

// Aggregate returns the 'higher' of the two state, given the following order:
//  ok < stopping < starting < down < failed
func Aggregate(state1, state2 State) State {
	for _, nextState := range StateOrder {
		if state1 == nextState || state2 == nextState {
			return nextState
		}
	}

	panic("unknown state: " + state1 + ", " + state2)
}

// Inactive means an app is not Up.
func IsStateInactive(state State) bool {
	return state != Up
}

// IsStateFinal returns whether the given state is a final state and should not change upon itself.
// E.g. A unit with a state Starting will after some time either switch to Up or Failed,
// thus the state is not final.
// Final states are Up, Down or Failed.
func IsStateFinal(state State) bool {
	return state == Failed || state == Up || state == Down || state == Deleted
}
