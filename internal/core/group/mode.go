package group

type Mode struct {
	Isolated bool
}

var state Mode

func State() Mode {
	return state
}

func SetState(mode Mode) error {
	state = mode
	return nil
}

func SetIsolated(b bool) error {
	state.Isolated = b
	return nil
}
