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
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func QuoteClass() QuoteClassLike {
	return quoteClass()
}

// Constructor Methods

func (c *quoteClass_) Quote(
	characters []Character,
) QuoteLike {
	return quote_(stc.Quote(string(characters)))
}

func (c *quoteClass_) QuoteFromSequence(
	sequence Sequential[Character],
) QuoteLike {
	return c.Quote(sequence.AsArray())
}

func (c *quoteClass_) QuoteFromString(
	source string,
) QuoteLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the quote constructor method: %s",
			source,
		)
		panic(message)
	}
	return quote_(source)
}

// Constant Methods

// Function Methods

func (c *quoteClass_) Concatenate(
	first QuoteLike,
	second QuoteLike,
) QuoteLike {
	return c.Quote(uti.CombineArrays(first.AsIntrinsic(), second.AsIntrinsic()))
}

// INSTANCE INTERFACE

// Principal Methods

func (v quote_) GetClass() QuoteClassLike {
	return quoteClass()
}

func (v quote_) AsIntrinsic() []Character {
	var unquoted, _ = stc.Unquote(string(v)) // Strip off the double quotes.
	return []Character(unquoted)
}

func (v quote_) AsString() string {
	return string(v)
}

// Attribute Methods

// Spectral[Quote] Methods

func (v quote_) CompareWith(
	value QuoteLike,
) age.Rank {
	switch sli.Compare(v.AsIntrinsic(), value.AsIntrinsic()) {
	case -1:
		return age.LesserRank
	case 1:
		return age.GreaterRank
	default:
		return age.EqualRank
	}
}

// Searchable[Character] Methods

func (v quote_) ContainsValue(
	value Character,
) bool {
	return sli.Index(v.AsIntrinsic(), value) > -1
}

func (v quote_) ContainsAny(
	values Sequential[Character],
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

func (v quote_) ContainsAll(
	values Sequential[Character],
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

// Sequential[Character] Methods

func (v quote_) IsEmpty() bool {
	return len(v.AsIntrinsic()) == 0
}

func (v quote_) GetSize() uint {
	return uint(len(v.AsIntrinsic()))
}

func (v quote_) AsArray() []Character {
	return v.AsIntrinsic()
}

func (v quote_) GetIterator() age.IteratorLike[Character] {
	return age.IteratorClass[Character]().Iterator(v.AsIntrinsic())
}

// Accessible[Character] Methods

func (v quote_) GetValue(
	index int,
) Character {
	var characters = v.AsIntrinsic()
	var size = uint(len(characters))
	var goIndex = uti.RelativeToCardinal(index, size)
	return characters[goIndex]
}

func (v quote_) GetValues(
	first int,
	last int,
) Sequential[Character] {
	var characters = v.AsIntrinsic()
	var size = uint(len(characters))
	var goFirst = uti.RelativeToCardinal(first, size)
	var goLast = uti.RelativeToCardinal(last, size)
	return quoteClass().Quote(characters[goFirst : goLast+1])
}

func (v quote_) GetIndex(
	value Character,
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

func (v Character) String() string {
	return fmt.Sprintf("%c", rune(v))
}

func (v quote_) String() string {
	return v.AsString()
}

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string quotes for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	base16_    = base10_ + "|[a-f]"
	character_ = escape_ + "|\\\\\"|[^\"" + control_ + "]"
	control_   = "\\p{Cc}"
	escape_    = "\\\\(?:" + unicode_ + "|[abfnrtv\\\\])"
	unicode_   = "u(?:" + base16_ + "){4}|U(?:" + base16_ + "){8}"
)

// Instance Structure

type quote_ string // This type must support the "comparable" type contraint.

// Class Structure

type quoteClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func quoteClass() *quoteClass_ {
	return quoteClassReference_
}

var quoteClassReference_ = &quoteClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^\"((?:" + character_ + ")*)\""),
}
