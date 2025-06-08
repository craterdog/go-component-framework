/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ControllerClass() ControllerClassLike {
	return controllerClass()
}

// Constructor Methods

func (c *controllerClass_) Controller(
	events []Event,
	transitions map[State]Transitions,
) ControllerLike {
	// Validate the constructor arguments.
	if uti.IsUndefined(events) {
		panic("The \"events\" attribute is required by this class.")
	}
	if uti.IsUndefined(transitions) {
		panic("The \"transitions\" attribute is required by this class.")
	}
	var height = len(transitions)
	if height < 3 {
		var message = fmt.Sprintf(
			"The state table must have at least two possible transitions: %v\n",
			height,
		)
		panic(message)
	}
	var width = len(events)
	if width < 2 {
		var message = fmt.Sprintf(
			"The state table must have at least one possible event: %v\n",
			width,
		)
		panic(message)
	}
	for _, row := range transitions {
		if len(row) != width {
			var message = fmt.Sprintf(
				"Each row in the state table must be the same width: %v\n",
				width,
			)
			panic(message)
		}
	}

	// Create a new instance.
	var instance = &controller_{
		// Initialize the instance attributes.
		state_:       1,
		events_:      events,
		transitions_: transitions,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *controller_) GetClass() ControllerClassLike {
	return controllerClass()
}

func (v *controller_) ProcessEvent(
	event Event,
) State {
	if event < 1 || event > Event(len(v.events_)) {
		var message = fmt.Sprintf(
			"Received an invalid event: %v",
			event,
		)
		panic(message)
	}
	var next = v.transitions_[v.state_][event-1]
	if next < 1 || next > State(len(v.transitions_)) {
		var message = fmt.Sprintf(
			"Attempted to transition to an invalid state: %v",
			next,
		)
		panic(message)
	}
	v.state_ = next
	return next
}

// Attribute Methods

func (v *controller_) GetState() State {
	return v.state_
}

func (v *controller_) SetState(
	state State,
) {
	if state == 0 {
		var message = fmt.Sprintf(
			"A valid \"state\" argument is required: %v\n",
			state,
		)
		panic(message)
	}
	v.state_ = state
}

func (v *controller_) GetEvents() []Event {
	return v.events_
}

func (v *controller_) GetTransitions() map[State]Transitions {
	return v.transitions_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type controller_ struct {
	// Declare the instance attributes.
	state_       State
	events_      []Event
	transitions_ map[State]Transitions
}

// Class Structure

type controllerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func controllerClass() *controllerClass_ {
	return controllerClassReference_
}

var controllerClassReference_ = &controllerClass_{
	// Initialize the class constants.
}
