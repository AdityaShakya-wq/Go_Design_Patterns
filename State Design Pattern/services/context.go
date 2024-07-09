package services

type StateContext struct {
	currentTvState TvState
}

func GetState() *StateContext {
	return &StateContext{
		currentTvState: &Off{},
	}
}

func (sc *StateContext) SetState(state TvState) {
	sc.currentTvState = state
}

func (sc *StateContext) GetState() {
	sc.currentTvState.State()
}
