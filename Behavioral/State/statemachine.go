package main

import (
	"fmt"
)

// iState is the interface to define a state.
type iState interface {
	Enter()
	Handle()
	Exit()
}

// StateMachine is the context to execute states
type StateMachine struct {
	states       map[string]iState
	currentState string
	lastState    string
}

func (sm *StateMachine) Exec() error {
	if sm.states == nil {
		return fmt.Errorf("the statemachine is empty")
	}
	sm.states[sm.currentState].Handle()
	return nil
}

func (sm *StateMachine) ChangeState(newState string) error {
	state, ok := sm.states[newState]

	if !ok {
		return fmt.Errorf("%s is not a valid state", newState)
	}

	sm.states[sm.currentState].Exit()
	sm.lastState, sm.currentState = sm.currentState, newState
	state.Enter()

	return nil
}
