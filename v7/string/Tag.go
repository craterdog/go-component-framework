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
	bin "encoding/binary"
	fmt "fmt"
	age "github.com/craterdog/go-component-framework/v7/agent"
	col "github.com/craterdog/go-component-framework/v7/collection"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
)

// CLASS INTERFACE

// Access Function

func TagClass() TagClassLike {
	return tagClass()
}

// Constructor Methods

func (c *tagClass_) Tag(
	bytes []byte,
) TagLike {
	return tag_(bytes)
}

func (c *tagClass_) TagWithSize(
	size age.Cardinal,
) TagLike {
	if size < 4 {
		var message = fmt.Sprintf(
			"A tag must be at least four bytes long: %v",
			size,
		)
		panic(message)
	}
	var generator = age.GeneratorClass().Generator()
	var bytes = generator.RandomBytes(size)
	return tag_(bytes)
}

func (c *tagClass_) TagFromSequence(
	sequence col.Sequential[byte],
) TagLike {
	var class = col.ListClass[byte]()
	var list = class.ListFromSequence(sequence)
	return tag_(list.AsArray())
}

func (c *tagClass_) TagFromString(
	string_ string,
) TagLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the tag constructor method: %s",
			string_,
		)
		panic(message)
	}
	var base32 = matches[1] // Strip off the leading "#".
	var encoder = age.EncoderClass().Encoder()
	var bytes = encoder.Base32Decode(base32)
	return tag_(bytes)
}

// Constant Methods

// Function Methods

func (c *tagClass_) Concatenate(
	first TagLike,
	second TagLike,
) TagLike {
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var allBytes = make(
		[]byte,
		len(firstBytes)+len(secondBytes),
	)
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return c.Tag(allBytes)
}

// INSTANCE INTERFACE

// Principal Methods

func (v tag_) GetClass() TagClassLike {
	return tagClass()
}

func (v tag_) AsIntrinsic() []byte {
	return []byte(v)
}

func (v tag_) AsString() string {
	var encoder = age.EncoderClass().Encoder()
	return "#" + encoder.Base32Encode(v)
}

// Attribute Methods

func (v tag_) GetHash() uint64 {
	return bin.BigEndian.Uint64(v)
}

// col.Sequential[byte] Methods

func (v tag_) IsEmpty() bool {
	return len(v) == 0
}

func (v tag_) GetSize() age.Cardinal {
	return age.Cardinal(len(v))
}

func (v tag_) AsArray() []byte {
	return uti.CopyArray(v)
}

func (v tag_) GetIterator() age.IteratorLike[byte] {
	var array = uti.CopyArray(v)
	var class = age.IteratorClass[byte]()
	var iterator = class.Iterator(array)
	return iterator
}

// col.Accessible[byte] Methods

func (v tag_) GetValue(
	index col.Index,
) byte {
	var class = col.ListClass[byte]()
	var list = class.ListFromArray(v)
	return list.GetValue(index)
}

func (v tag_) GetValues(
	first col.Index,
	last col.Index,
) col.Sequential[byte] {
	var class = col.ListClass[byte]()
	var list = class.ListFromArray(v)
	return list.GetValues(first, last)
}

// PROTECTED INTERFACE

func (v tag_) String() string {
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
	base32_ = base10_ + "|[A-DF-HJ-NP-TV-Z]"
)

// Instance Structure

type tag_ []byte

// Class Structure

type tagClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func tagClass() *tagClass_ {
	return tagClassReference_
}

var tagClassReference_ = &tagClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^#((?:" + base32_ + ")+)"),
}
