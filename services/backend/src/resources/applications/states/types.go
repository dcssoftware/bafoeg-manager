package states

import "errors"

type ApplicationState string

const (
	StatusInProgress      ApplicationState = "IN_PROGRESS"
	StatusResponseAwaited ApplicationState = "RESPONSE_AWAITED"
	StatusApproved        ApplicationState = "APPROVED"
	StatusDenied          ApplicationState = "DENIED"
	StatusUnknown         ApplicationState = "UNKNOWN"
)

func (s ApplicationState) ToString() string {
	return string(s)
}

func ConvertStrToApplicationState(status string) (ApplicationState, error) {
	switch status {
	case StatusInProgress.ToString():
		return StatusInProgress, nil
	case StatusResponseAwaited.ToString():
		return StatusResponseAwaited, nil
	case StatusApproved.ToString():
		return StatusApproved, nil
	case StatusDenied.ToString():
		return StatusDenied, nil
	default:
		return StatusUnknown, errors.New("cannot find status in defined, available status list")
	}
}
