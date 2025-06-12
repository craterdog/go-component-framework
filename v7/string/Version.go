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
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func VersionClass() VersionClassLike {
	return versionClass()
}

// Constructor Methods

func (c *versionClass_) Version(
	ordinals []age.Ordinal,
) VersionLike {
	return version_(ordinals)
}

func (c *versionClass_) VersionFromSequence(
	sequence col.Sequential[age.Ordinal],
) VersionLike {
	var class = col.ListClass[age.Ordinal]()
	var list = class.ListFromSequence(sequence)
	return version_(list.AsArray())
}

func (c *versionClass_) VersionFromString(
	string_ string,
) VersionLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the version constructor method: %s",
			string_,
		)
		panic(message)
	}
	var match = matches[1] // Strip off the leading "v".
	var levels = sts.Split(match, ".")
	var ordinals = make([]age.Ordinal, len(levels))
	for index, level := range levels {
		var ordinal, _ = stc.ParseUint(level, 10, 64)
		ordinals[index] = age.Ordinal(ordinal)
	}
	return version_(ordinals)
}

// Constant Methods

// Function Methods

func (c *versionClass_) IsValidNextVersion(
	current VersionLike,
	next VersionLike,
) bool {
	// Make sure the version sizes are compatible.
	var currentOrdinals = current.AsArray()
	var currentSize = len(currentOrdinals)
	var nextOrdinals = next.AsArray()
	var nextSize = len(nextOrdinals)
	if nextSize > currentSize+1 {
		return false
	}

	// Iterate through the versions comparing level values.
	var class = age.IteratorClass[age.Ordinal]()
	var currentIterator = class.Iterator(current.AsArray())
	var nextIterator = class.Iterator(next.AsArray())
	for currentIterator.HasNext() && nextIterator.HasNext() {
		var currentLevel = currentIterator.GetNext()
		var nextLevel = nextIterator.GetNext()
		if currentLevel == nextLevel {
			// So far the level values are the same.
			continue
		}
		// The last level for the next version must be one more.
		return !nextIterator.HasNext() && nextLevel == currentLevel+1
	}
	// The last level for the next version must be one.
	return nextIterator.HasNext() && nextIterator.GetNext() == 1
}

func (c *versionClass_) GetNextVersion(
	current VersionLike,
	level age.Ordinal,
) VersionLike {
	// Adjust the size of the ordinals as needed.
	var ordinals = current.AsArray()
	var size = age.Ordinal(len(ordinals))
	switch {
	case level == 0:
		level = size // Normalize the level to the current size.
	case level < size:
		// The next version will require fewer levels.
		ordinals = ordinals[:level]
	case level > size:
		// The next version will require another level.
		size++
		level = size // Normalize the level to the new size.
		ordinals = append(ordinals, 0)
	}

	// Increment the specified version level.
	var index = level - 1 // Convert to zero based indexing.
	ordinals[index]++

	var version = c.Version(ordinals)
	return version
}

func (c *versionClass_) Concatenate(
	first VersionLike,
	second VersionLike,
) VersionLike {
	var firstOrdinals = first.AsArray()
	var secondOrdinals = second.AsArray()
	var allOrdinals = make(
		[]age.Ordinal,
		len(firstOrdinals)+len(secondOrdinals),
	)
	copy(allOrdinals, firstOrdinals)
	copy(allOrdinals[len(firstOrdinals):], secondOrdinals)
	return c.Version(allOrdinals)
}

// INSTANCE INTERFACE

// Principal Methods

func (v version_) GetClass() VersionClassLike {
	return versionClass()
}

func (v version_) AsIntrinsic() []age.Ordinal {
	return []age.Ordinal(v)
}

func (v version_) AsString() string {
	var index = 0
	var string_ = "v" + stc.Itoa(int(v[index]))
	for index++; index < len(v); index++ {
		string_ += "." + stc.Itoa(int(v[index]))
	}
	return string_
}

// Attribute Methods

// col.Sequential[age.Ordinal] Methods

func (v version_) IsEmpty() bool {
	return len(v) == 0
}

func (v version_) GetSize() age.Cardinal {
	return age.Cardinal(len(v))
}

func (v version_) AsArray() []age.Ordinal {
	return uti.CopyArray(v)
}

func (v version_) GetIterator() age.IteratorLike[age.Ordinal] {
	var array = uti.CopyArray(v)
	var class = age.IteratorClass[age.Ordinal]()
	var iterator = class.Iterator(array)
	return iterator
}

// col.Accessible[age.Ordinal] Methods

func (v version_) GetValue(
	index col.Index,
) age.Ordinal {
	var class = col.ListClass[age.Ordinal]()
	var list = class.ListFromArray(v)
	return list.GetValue(index)
}

func (v version_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[age.Ordinal] {
	var class = col.ListClass[age.Ordinal]()
	var list = class.ListFromArray(v)
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v version_) String() string {
	return v.AsString()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each version to lessen the chance of a version collision with other private Go
// class constants in this package.
const (
	ordinal_ = "[1-9](?:" + base10_ + ")*"
)

// Instance Structure

type version_ []age.Ordinal

// Class Structure

type versionClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func versionClass() *versionClass_ {
	return versionClassReference_
}

var versionClassReference_ = &versionClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^v(" + ordinal_ + "(?:\\." + ordinal_ + ")*)",
	),
}
