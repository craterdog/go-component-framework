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
)

// CLASS INTERFACE

// Access Function

func SymbolClass() SymbolClassLike {
	return symbolClass()
}

// Constructor Methods

func (c *symbolClass_) Symbol(
	runes []rune,
) SymbolLike {
	return symbol_(runes)
}

func (c *symbolClass_) SymbolFromSequence(
	sequence col.Sequential[rune],
) SymbolLike {
	var class = col.ListClass[rune]()
	var list = class.ListFromSequence(sequence)
	return symbol_(list.AsArray())
}

func (c *symbolClass_) SymbolFromString(
	string_ string,
) SymbolLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the symbol constructor method: %s",
			string_,
		)
		panic(message)
	}
	return symbol_(matches[1]) // Strip off the leading "$".
}

// Constant Methods

// Function Methods

func (c *symbolClass_) Concatenate(
	first SymbolLike,
	second SymbolLike,
) SymbolLike {
	return c.Symbol(uti.CombineArrays(first.AsIntrinsic(), second.AsIntrinsic()))
}

// INSTANCE INTERFACE

// Principal Methods

func (v symbol_) GetClass() SymbolClassLike {
	return symbolClass()
}

func (v symbol_) AsIntrinsic() []rune {
	return []rune(v)
}

func (v symbol_) AsString() string {
	return "$" + string(v)
}

// Attribute Methods

// col.Sequential[rune] Methods

func (v symbol_) IsEmpty() bool {
	return len(v) == 0
}

func (v symbol_) GetSize() age.Cardinal {
	return age.Cardinal(len(v.AsArray()))
}

func (v symbol_) AsArray() []rune {
	return []rune(v)
}

func (v symbol_) GetIterator() age.IteratorLike[rune] {
	var class = age.IteratorClass[rune]()
	var iterator = class.Iterator(v.AsArray())
	return iterator
}

// col.Accessible[rune] Methods

func (v symbol_) GetValue(
	index col.Index,
) rune {
	var class = col.ListClass[rune]()
	var list = class.ListFromArray(v.AsArray())
	return list.GetValue(index)
}

func (v symbol_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[rune] {
	var class = col.ListClass[rune]()
	var list = class.ListFromArray(v.AsArray())
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v symbol_) String() string {
	return v.AsString()
}

// Private Methods

// Instance Structure

type symbol_ string

// Class Structure

type symbolClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func symbolClass() *symbolClass_ {
	return symbolClassReference_
}

var symbolClassReference_ = &symbolClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^\\$(" + identifier_ + ")"),
}
