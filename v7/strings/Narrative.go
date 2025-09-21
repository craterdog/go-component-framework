/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package strings

import (
	fmt "fmt"
	age "github.com/craterdog/go-component-framework/v7/agents"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
	sli "slices"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func NarrativeClass() NarrativeClassLike {
	return narrativeClass()
}

// Constructor Methods

func (c *narrativeClass_) Narrative(
	lines []Line,
) NarrativeLike {
	var source = "\">"
	if len(lines) > 0 {
		for _, line := range lines {
			var encoded = sts.ReplaceAll(string(line), `">`, `\">`)
			encoded = sts.ReplaceAll(encoded, `<"`, `<\"`)
			source += "\n" + encoded
		}
		source += "\n"
	}
	source += "<\""
	return narrative_(source)
}

func (c *narrativeClass_) NarrativeFromSequence(
	sequence Sequential[Line],
) NarrativeLike {
	return c.Narrative(sequence.AsArray())
}

func (c *narrativeClass_) NarrativeFromString(
	source string,
) NarrativeLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the narrative constructor method: %s",
			source,
		)
		panic(message)
	}
	return narrative_(source)
}

// Constant Methods

// Function Methods

func (c *narrativeClass_) Concatenate(
	first NarrativeLike,
	second NarrativeLike,
) NarrativeLike {
	var firstLines = first.AsArray()
	var secondLines = second.AsArray()
	var allLines = make(
		[]Line,
		len(firstLines)+len(secondLines),
	)
	copy(allLines, firstLines)
	copy(allLines[len(firstLines):], secondLines)
	return c.Narrative(allLines)
}

// INSTANCE INTERFACE

// Principal Methods

func (v narrative_) GetClass() NarrativeClassLike {
	return narrativeClass()
}

func (v narrative_) AsIntrinsic() []Line {
	var narrative = string(v)
	var decoded = sts.ReplaceAll(narrative[2:len(v)-2], `\">`, `">`)
	decoded = sts.ReplaceAll(decoded, `<\"`, `<"`)
	var strings = sts.Split(decoded, "\n")
	strings = strings[1:] // Ignore the first empty line.
	var size = len(strings)
	if size > 0 {
		size--
		strings = strings[:size] // Ignore the last empty line.
	}
	var lines = make([]Line, size)
	for index, line := range strings {
		lines[index] = Line(line)
	}
	return []Line(lines)
}

func (v narrative_) AsString() string {
	return string(v)
}

// Attribute Methods

// Searchable[Line] Methods

func (v narrative_) ContainsValue(
	value Line,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v narrative_) ContainsAny(
	values Sequential[Line],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This set contains at least one of the values.
			return true
		}
	}
	// This set does not contain any of the values.
	return false
}

func (v narrative_) ContainsAll(
	values Sequential[Line],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This set is missing at least one of the values.
			return false
		}
	}
	// This set does contains all of the values.
	return true
}

// Sequential[Line] Methods

func (v narrative_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v narrative_) GetSize() uint {
	return uint(len(v.AsIntrinsic()))
}

func (v narrative_) AsArray() []Line {
	return v.AsIntrinsic()
}

func (v narrative_) GetIterator() age.IteratorLike[Line] {
	return age.IteratorClass[Line]().Iterator(v.AsIntrinsic())
}

// Accessible[Line] Methods

func (v narrative_) GetValue(
	index int,
) Line {
	var lines = v.AsIntrinsic()
	var size = uint(len(lines))
	var goIndex = uti.RelativeToCardinal(index, size)
	return lines[goIndex]
}

func (v narrative_) GetValues(
	first int,
	last int,
) Sequential[Line] {
	var lines = v.AsIntrinsic()
	var size = uint(len(lines))
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return narrativeClass().Narrative(lines[goFirst : goLast+1])
}

func (v narrative_) GetIndex(
	value Line,
) uint {
	var index uint
	var iterator = v.GetIterator()
	for iterator.HasNext() {
		index++
		var candidate = iterator.GetNext()
		if candidate == value {
			// Found the value.
			return index
		}
	}
	// The value was not found.
	return 0
}

// PROTECTED INTERFACE

func (v narrative_) String() string {
	return v.AsString()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each narrative to lessen the chance of a narrative collision with other private Go
// class constants in this package.
const (
	any_ = "." // This does NOT include newline characters.
	eol_ = "\\r?\\n"
)

// Instance Structure

type narrative_ string // This type must support the "comparable" type contraint.

// Class Structure

type narrativeClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func narrativeClass() *narrativeClass_ {
	return narrativeClassReference_
}

var narrativeClassReference_ = &narrativeClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^\">((?:" + any_ + "|" + eol_ + ")*?)<\"",
	),
}
