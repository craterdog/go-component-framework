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

package string

import (
	fmt "fmt"
	age "github.com/craterdog/go-component-framework/v7/agent"
	col "github.com/craterdog/go-component-framework/v7/collection"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
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
	return narrative_(lines)
}

func (c *narrativeClass_) NarrativeFromSequence(
	sequence col.Sequential[Line],
) NarrativeLike {
	var class = col.ListClass[Line]()
	var list = class.ListFromSequence(sequence)
	return narrative_(list.AsArray())
}

func (c *narrativeClass_) NarrativeFromString(
	string_ string,
) NarrativeLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the narrative constructor method: %s",
			string_,
		)
		panic(message)
	}
	var narrative = matches[1]               // Strip off the delimiters.
	var strings = sts.Split(narrative, "\n") // Extract the lines.
	strings = strings[1:]                    // Ignore the first empty line.
	var size = len(strings)
	if size > 0 {
		size--
		strings = strings[:size] // Ignore the last empty line.
	}
	var lines = make([]Line, size)
	for index, line := range strings {
		lines[index] = Line(line)
	}
	return narrative_(lines)
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
	return []Line(v)
}

func (v narrative_) AsString() string {
	var string_ = "\">"
	if len(v) > 0 {
		for _, line := range v {
			string_ += "\n" + string(line)
		}
		string_ += "\n"
	}
	string_ += "<\""
	return string_
}

// Attribute Methods

// col.Sequential[Line] Methods

func (v narrative_) IsEmpty() bool {
	return len(v) == 0
}

func (v narrative_) GetSize() age.Cardinal {
	return age.Cardinal(len(v))
}

func (v narrative_) AsArray() []Line {
	return uti.CopyArray(v)
}

func (v narrative_) GetIterator() age.IteratorLike[Line] {
	var array = uti.CopyArray(v)
	var class = age.IteratorClass[Line]()
	var iterator = class.Iterator(array)
	return iterator
}

// col.Accessible[Line] Methods

func (v narrative_) GetValue(
	index col.Index,
) Line {
	var class = col.ListClass[Line]()
	var list = class.ListFromArray(v)
	return list.GetValue(index)
}

func (v narrative_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[Line] {
	var class = col.ListClass[Line]()
	var list = class.ListFromArray(v)
	return list.GetValues(first, last)
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

type narrative_ []Line

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
