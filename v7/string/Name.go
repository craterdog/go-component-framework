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

func NameClass() NameClassLike {
	return nameClass()
}

// Constructor Methods

func (c *nameClass_) Name(
	identifiers []Identifier,
) NameLike {
	return name_(identifiers)
}

func (c *nameClass_) NameFromSequence(
	sequence col.Sequential[Identifier],
) NameLike {
	var class = col.ListClass[Identifier]()
	var list = class.ListFromSequence(sequence)
	return name_(list.AsArray())
}

func (c *nameClass_) NameFromString(
	string_ string,
) NameLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the name constructor method: %s",
			string_,
		)
		panic(message)
	}
	var name = matches[0]
	var strings = sts.Split(name, "/")[1:] // Extract the identifiers.
	var identifiers = make([]Identifier, len(strings))
	for index, identifier := range strings {
		identifiers[index] = Identifier(identifier)
	}
	return name_(identifiers)
}

// Constant Methods

// Function Methods

func (c *nameClass_) Concatenate(
	first NameLike,
	second NameLike,
) NameLike {
	var firstIdentifiers = first.AsArray()
	var secondIdentifiers = second.AsArray()
	var allIdentifiers = make(
		[]Identifier,
		len(firstIdentifiers)+len(secondIdentifiers),
	)
	copy(allIdentifiers, firstIdentifiers)
	copy(allIdentifiers[len(firstIdentifiers):], secondIdentifiers)
	return c.Name(allIdentifiers)
}

// INSTANCE INTERFACE

// Principal Methods

func (v name_) GetClass() NameClassLike {
	return nameClass()
}

func (v name_) AsIntrinsic() []Identifier {
	return []Identifier(v)
}

func (v name_) AsString() string {
	var string_ string
	for _, identifier := range v {
		string_ += "/" + string(identifier)
	}
	return string_
}

// Attribute Methods

// col.Sequential[Identifier] Methods

func (v name_) IsEmpty() bool {
	return len(v) == 0
}

func (v name_) GetSize() age.Cardinal {
	return age.Cardinal(len(v))
}

func (v name_) AsArray() []Identifier {
	return uti.CopyArray(v)
}

func (v name_) GetIterator() age.IteratorLike[Identifier] {
	var array = uti.CopyArray(v)
	var class = age.IteratorClass[Identifier]()
	var iterator = class.Iterator(array)
	return iterator
}

// col.Accessible[Identifier] Methods

func (v name_) GetValue(
	index col.Index,
) Identifier {
	var class = col.ListClass[Identifier]()
	var list = class.ListFromArray(v)
	return list.GetValue(index)
}

func (v name_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[Identifier] {
	var class = col.ListClass[Identifier]()
	var list = class.ListFromArray(v)
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v name_) String() string {
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
	digit_      = "\\p{Nd}"
	identifier_ = "(?:" + letter_ + ")(?:" + letter_ + "|" + digit_ + "|-)*"
	letter_     = lower_ + "|" + upper_
	lower_      = "\\p{Ll}"
	upper_      = "\\p{Lu}"
)

// Instance Structure

type name_ []Identifier

// Class Structure

type nameClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func nameClass() *nameClass_ {
	return nameClassReference_
}

var nameClassReference_ = &nameClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^(?:/" + identifier_ + ")+",
	),
}
