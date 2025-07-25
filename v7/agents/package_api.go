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

/*
Package "agents" declares a set of agents that operate on values that have a
generic type.  They are used by the collection classes declared in this Go
module.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-component-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package agents

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// TYPE DECLARATIONS

/*
Event is a constrained type representing an event type in a state machine.
Using a string type for an event makes it easier to print out in a human
readable way.
*/
type Event string

/*
Rank is a constrained type representing the possible rankings for two values.
*/
type Rank uint8

const (
	LesserRank Rank = iota
	EqualRank
	GreaterRank
)

/*
Slot is a constrained type representing a slot between values in a sequence.
*/
type Slot uint

/*
State is a constrained type representing a state in a state machine.  Using a
string type for a state makes it easier to print out in a human readable way.
*/
type State string

/*
Transitions is a constrained type representing a row of states in a state machine.
*/
type Transitions []State

// FUNCTIONAL DECLARATIONS

/*
RankingFunction[V any] is a functional type that declares the signature for any
function that can determine the relative ranking of two values.
*/
type RankingFunction[V any] func(
	first V,
	second V,
) Rank

// CLASS DECLARATIONS

/*
CollatorClassLike[V any] is a class interface that declares the complete set
of class constructors, constants and functions that must be supported by each
concrete collator-like class.

A collator-like class is capable of recursively comparing and ranking two values
of any type.  An optional maximum depth may be specified that limits the depth
of the structures being collated to avoid possible infinite recursion.

The default maximum depth is 16.
*/
type CollatorClassLike[V any] interface {
	// Constructor Methods
	Collator() CollatorLike[V]
	CollatorWithMaximumDepth(
		maximumDepth uti.Cardinal,
	) CollatorLike[V]
}

/*
ControllerClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
controller-like class.

A controller-like class implements a state machine based on a finite state
machine and possible event types. It enforces the possible states of the state
machine and allowed transitions between states given a finite set of possible
event types. It implements a finite state machine with the following table
structure:

	                    events:
	        -------------------------------
	        [event1,  event2,  ... eventM ]

	                 transitions:
	        -------------------------------
	state1: [invalid, state2,  ... invalid]
	state2: [state3,  stateN,  ... invalid]
	                    ...
	stateN: [state1,  invalid, ... state3 ]

The first row of the state machine defines the possible events that can occur.
Each subsequent row defines a state and the possible transitions from that
state to the next state for each possible event. Transitions marked as "invalid"
cannot occur. The state machine always starts in the first state of the finite
state machine (e.g. state1).
*/
type ControllerClassLike interface {
	// Constructor Methods
	Controller(
		events []Event,
		transitions map[State]Transitions,
		initialState State,
	) ControllerLike

	// Constant Methods
	Invalid() State
}

/*
EncoderClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
encoder-like class.

A encoder-like class implements encoding and decoding algorithms for the
following:
  - Base 16 [0-9][a-f]
  - Base 32 [0-9][A-D][F-H][J-N][P-T][V-Z]  {excludes "EIOU"}
  - Base 64 [0-9][A-Z][a-z][+/]
*/
type EncoderClassLike interface {
	// Constructor Methods
	Encoder() EncoderLike
}

/*
GeneratorClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
generator-like class.

A generator-like class generates various types of cryptographically secure
random values.
*/
type GeneratorClassLike interface {
	// Constructor Methods
	Generator() GeneratorLike
}

/*
IteratorClassLike[V any] is a class interface that declares the complete set
of class constructors, constants and functions that must be supported by each
concrete iterator-like class.

An iterator-like class can be used to move forward and backward over the values
in an array.  It implements the Gang of Four (GoF) Iterator Design Pattern:
  - https://en.wikipedia.org/wiki/Iterator_pattern

An iterator agent locks into the slots that reside between each value in the
sequence:

	  . [value 1] . [value 2] . [value 3] ... [value N] .
	  ^           ^           ^                         ^
	slot 0      slot 1      slot 2                    slot N

It moves from slot to slot and has access to the values (if they exist) on each
side of the slot.  At each slot an iterator has access to the previous value
and next value in the array (assuming they exist). The slot at the start of
the array has no PREVIOUS value, and the slot at the end of the array has no
NEXT value.  The size of the array is static so that its values can be modified
during iteration.
*/
type IteratorClassLike[V any] interface {
	// Constructor Methods
	Iterator(
		array []V,
	) IteratorLike[V]
}

/*
SorterClassLike[V any] is a class interface that declares the complete set
of class constructors, constants and functions that must be supported by each
concrete sorter-like class.

A sorter-like class implements a specific sorting algorithm.  It uses a ranking
function to correlate the values.  If no ranking function is specified the
values are sorted into their "natural" ordering by type of value.
*/
type SorterClassLike[V any] interface {
	// Constructor Methods
	Sorter() SorterLike[V]
	SorterWithRanker(
		ranker RankingFunction[V],
	) SorterLike[V]
}

// INSTANCE DECLARATIONS

/*
CollatorLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete collator-like class.
*/
type CollatorLike[V any] interface {
	// Principal Methods
	GetClass() CollatorClassLike[V]
	CompareValues(
		first V,
		second V,
	) bool
	RankValues(
		first V,
		second V,
	) Rank

	// Attribute Methods
	GetMaximumDepth() uti.Cardinal
}

/*
ControllerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete controller-like class.
*/
type ControllerLike interface {
	// Principal Methods
	GetClass() ControllerClassLike
	ProcessEvent(
		event Event,
	) State

	// Attribute Methods
	GetState() State
	SetState(
		state State,
	)
	GetEvents() []Event
	GetTransitions() map[State]Transitions
}

/*
EncoderLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete encoder-like class.
*/
type EncoderLike interface {
	// Principal Methods
	GetClass() EncoderClassLike
	Base16Encode(
		bytes []byte,
	) string
	Base16Decode(
		encoded string,
	) []byte
	Base32Encode(
		bytes []byte,
	) string
	Base32Decode(
		encoded string,
	) []byte
	Base64Encode(
		bytes []byte,
	) string
	Base64Decode(
		encoded string,
	) []byte
}

/*
GeneratorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete generator-like class.
*/
type GeneratorLike interface {
	// Principal Methods
	GetClass() GeneratorClassLike
	RandomBoolean() bool
	RandomOrdinal(
		maximum uti.Ordinal,
	) uti.Ordinal
	RandomProbability() float64
	RandomBytes(
		size uti.Cardinal,
	) []byte
}

/*
IteratorLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete iterator-like class.
*/
type IteratorLike[V any] interface {
	// Principal Methods
	GetClass() IteratorClassLike[V]
	IsEmpty() bool
	ToStart()
	ToEnd()
	HasPrevious() bool
	GetPrevious() V
	HasNext() bool
	GetNext() V

	// Attribute Methods
	GetSize() uti.Cardinal
	GetSlot() Slot
	SetSlot(
		slot Slot,
	)
}

/*
SorterLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete sorter-like class.
*/
type SorterLike[V any] interface {
	// Principal Methods
	GetClass() SorterClassLike[V]
	SortValues(
		values []V,
	)
	ReverseValues(
		values []V,
	)
	ShuffleValues(
		values []V,
	)

	// Attribute Methods
	GetRanker() RankingFunction[V]
}

// ASPECT DECLARATIONS
