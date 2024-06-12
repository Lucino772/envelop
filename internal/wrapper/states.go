package wrapper

type ProcessStatusState struct {
	Description string `json:"description"`
}

func (state ProcessStatusState) GetStateName() string {
	return "/process/status"
}

func (state ProcessStatusState) Equals(otherState WrapperState) bool {
	if val, ok := otherState.(ProcessStatusState); ok {
		return state.Description == val.Description
	}
	return false
}
