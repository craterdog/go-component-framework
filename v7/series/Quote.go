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

package series

import (
	fmt "fmt"
	age "github.com/craterdog/go-component-framework/v7/agent"
	col "github.com/craterdog/go-component-framework/v7/collection"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
)

// CLASS INTERFACE

// Access Function

func QuoteClass() QuoteClassLike {
	return quoteClass()
}

// Constructor Methods

func (c *quoteClass_) Quote(
	string_ string,
) QuoteLike {
	return quote_(string_)
}

func (c *quoteClass_) QuoteFromSequence(
	sequence col.Sequential[rune],
) QuoteLike {
	var class = col.ListClass[rune]()
	var list = class.ListFromSequence(sequence)
	return quote_(list.AsArray())
}

func (c *quoteClass_) QuoteFromString(
	string_ string,
) QuoteLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the quote constructor method: %s",
			string_,
		)
		panic(message)
	}
	return quote_(matches[1]) // Strip off the double quotes.
}

// Constant Methods

// Function Methods

func (c *quoteClass_) Concatenate(
	first QuoteLike,
	second QuoteLike,
) QuoteLike {
	return c.Quote(first.GetIntrinsic() + second.GetIntrinsic())
}

// INSTANCE INTERFACE

// Principal Methods

func (v quote_) GetClass() QuoteClassLike {
	return quoteClass()
}

func (v quote_) GetIntrinsic() string {
	return string(v)
}

func (v quote_) AsString() string {
	return `"` + v.GetIntrinsic() + `"`
}

// Attribute Methods

// col.Sequential[rune] Methods

func (v quote_) IsEmpty() bool {
	return len(v) == 0
}

func (v quote_) GetSize() uti.Cardinal {
	return uti.Cardinal(len(v.AsArray()))
}

func (v quote_) AsArray() []rune {
	return []rune(v)
}

func (v quote_) GetIterator() age.IteratorLike[rune] {
	var class = age.IteratorClass[rune]()
	var iterator = class.Iterator(v.AsArray())
	return iterator
}

// col.Accessible[rune] Methods

func (v quote_) GetValue(
	index col.Index,
) rune {
	var class = col.ListClass[rune]()
	var list = class.ListFromArray(v.AsArray())
	return list.GetValue(index)
}

func (v quote_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[rune] {
	var class = col.ListClass[rune]()
	var list = class.ListFromArray(v.AsArray())
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

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
	character_ = escape_ + "|\\\\\"|[^\"" + control_ + "]"
	control_   = "\\p{Cc}"
	escape_    = "\\\\(?:" + unicode_ + "|[abfnrtv\\\\])"
	unicode_   = "u(?:" + base16_ + "){4}|U(?:" + base16_ + "){8}"
)

// Instance Structure

type quote_ string

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
