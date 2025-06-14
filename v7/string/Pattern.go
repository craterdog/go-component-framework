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
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func PatternClass() PatternClassLike {
	return patternClass()
}

// Constructor Methods

func (c *patternClass_) Pattern(
	characters []Character,
) PatternLike {
	reg.MustCompile(string(characters))
	return pattern_(characters)
}

func (c *patternClass_) PatternFromSequence(
	sequence col.Sequential[Character],
) PatternLike {
	var class = col.ListClass[Character]()
	var list = class.ListFromSequence(sequence)
	return pattern_(list.AsArray())
}

func (c *patternClass_) PatternFromString(
	string_ string,
) PatternLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the pattern constructor method: %s",
			string_,
		)
		panic(message)
	}
	switch matches[0] {
	case "none":
		return c.none_
	case "any":
		return c.any_
	default:
		var unquoted, _ = stc.Unquote(matches[0]) // Strip off the double quotes.
		reg.MustCompile(unquoted)
		return pattern_(unquoted)
	}
}

// Constant Methods

func (c *patternClass_) None() PatternLike {
	return c.none_
}

func (c *patternClass_) Any() PatternLike {
	return c.any_
}

// Function Methods

func (c *patternClass_) Concatenate(
	first PatternLike,
	second PatternLike,
) PatternLike {
	return c.Pattern(uti.CombineArrays(first.AsIntrinsic(), second.AsIntrinsic()))
}

// INSTANCE INTERFACE

// Principal Methods

func (v pattern_) GetClass() PatternClassLike {
	return patternClass()
}

func (v pattern_) AsIntrinsic() []Character {
	return []Character(v)
}

func (v pattern_) AsString() string {
	var string_ string
	switch v {
	case `^none$`:
		string_ = `none`
	case `.*`:
		string_ = `any`
	default:
		string_ = stc.Quote(string(v)) + "?"
	}
	return string_
}

func (v pattern_) AsRegexp() *reg.Regexp {
	return reg.MustCompile(string(v))
}

func (v pattern_) MatchesText(
	text string,
) bool {
	var matcher = reg.MustCompile(string(v))
	return matcher.MatchString(text)
}

func (v pattern_) GetMatches(
	text string,
) []string {
	var matcher = reg.MustCompile(string(v))
	return matcher.FindStringSubmatch(text)
}

// Attribute Methods

// col.Sequential[Character] Methods

func (v pattern_) IsEmpty() bool {
	return len(v) == 0
}

func (v pattern_) GetSize() age.Cardinal {
	return age.Cardinal(len(v.AsArray()))
}

func (v pattern_) AsArray() []Character {
	return []Character(v)
}

func (v pattern_) GetIterator() age.IteratorLike[Character] {
	var class = age.IteratorClass[Character]()
	var iterator = class.Iterator(v.AsArray())
	return iterator
}

// col.Accessible[Character] Methods

func (v pattern_) GetValue(
	index col.Index,
) Character {
	var class = col.ListClass[Character]()
	var list = class.ListFromArray(v.AsArray())
	return list.GetValue(index)
}

func (v pattern_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[Character] {
	var class = col.ListClass[Character]()
	var list = class.ListFromArray(v.AsArray())
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v pattern_) String() string {
	return v.AsString()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	regex_ = "\"((?:" + character_ + ")+)\"\\\\?"
)

// Instance Structure

type pattern_ string

// Class Structure

type patternClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
	none_    PatternLike
	any_     PatternLike
}

// Class Reference

func patternClass() *patternClass_ {
	return patternClassReference_
}

var patternClassReference_ = &patternClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^" + regex_ + "|any|none"),
	none_:    pattern_(`^none$`),
	any_:     pattern_(`.*`),
}
