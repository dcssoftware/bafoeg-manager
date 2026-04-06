package states

func IsDoneString(statusIdentifier string) (bool, error) {
	state, err := ConvertStrToApplicationState(statusIdentifier)
	if err != nil {
		return true, err
	}

	return IsDone(state), nil
}

func IsDone(status ApplicationState) bool {
	switch status {
	case StatusApproved:
		return true
	case StatusDenied:
		return true
	default:
		return false
	}
}
