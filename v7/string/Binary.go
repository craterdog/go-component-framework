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
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func BinaryClass() BinaryClassLike {
	return binaryClass()
}

// Constructor Methods

func (c *binaryClass_) Binary(
	bytes []byte,
) BinaryLike {
	return binary_(bytes)
}

func (c *binaryClass_) BinaryFromSequence(
	sequence Sequential[byte],
) BinaryLike {
	return binary_(sequence.AsArray())
}

func (c *binaryClass_) BinaryFromString(
	source string,
) BinaryLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the binary constructor method: %s",
			source,
		)
		panic(message)
	}
	var base64 = matches[1]                   // Strip off the delimiters.
	base64 = sts.ReplaceAll(base64, " ", "")  // Remove all spaces.
	base64 = sts.ReplaceAll(base64, "\n", "") // Remove all newlines.
	var encoder = age.EncoderClass().Encoder()
	var bytes = encoder.Base64Decode(base64)
	return binary_(bytes)
}

// Constant Methods

// Function Methods

func (c *binaryClass_) Not(
	binary BinaryLike,
) BinaryLike {
	var bytes = binary.AsArray()
	var size = len(bytes)
	for i := 0; i < size; i++ {
		bytes[i] = ^bytes[i]
	}
	return c.Binary(bytes)
}

func (c *binaryClass_) And(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] & secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) San(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] &^ secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Ior(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] | secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Xor(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var result []byte
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var size = len(firstBytes)
	if size < len(secondBytes) {
		size = len(secondBytes)
		result = make([]byte, size)
		copy(result, firstBytes)
		firstBytes = result
	} else {
		result = make([]byte, size)
		copy(result, secondBytes)
		secondBytes = result
	}
	for i := 0; i < size; i++ {
		result[i] = firstBytes[i] ^ secondBytes[i]
	}
	return c.Binary(result)
}

func (c *binaryClass_) Concatenate(
	first BinaryLike,
	second BinaryLike,
) BinaryLike {
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var allBytes = make(
		[]byte,
		len(firstBytes)+len(secondBytes),
	)
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return c.Binary(allBytes)
}

// INSTANCE INTERFACE

// Principal Methods

func (v binary_) GetClass() BinaryClassLike {
	return binaryClass()
}

func (v binary_) AsIntrinsic() []byte {
	return []byte(v)
}

func (v binary_) AsString() string {
	var encoder = age.EncoderClass().Encoder()
	var encoded = encoder.Base64Encode(v)
	var length = len(encoded)
	var string_ = "'>"
	if length > 0 {
		string_ += "\n"
		var width = 60
		var indentation = "    "
		var index int
		for index = 0; index+width < length; index += width {
			string_ += indentation + encoded[index:index+width] + "\n"
		}
		string_ += indentation + encoded[index:] + "\n"
	}
	string_ += "<'"
	return string_
}

// Attribute Methods

// Sequential[byte] Methods

func (v binary_) IsEmpty() bool {
	return len(v) == 0
}

func (v binary_) GetSize() uti.Cardinal {
	return uti.Cardinal(len(v))
}

func (v binary_) AsArray() []byte {
	return uti.CopyArray(v)
}

func (v binary_) GetIterator() age.IteratorLike[byte] {
	var array = uti.CopyArray(v)
	var class = age.IteratorClass[byte]()
	var iterator = class.Iterator(array)
	return iterator
}

// Accessible[byte] Methods

func (v binary_) GetValue(
	index uti.Index,
) byte {
	var size = v.GetSize()
	var goIndex = uti.RelativeToZeroBased(index, size)
	return v[goIndex]
}

func (v binary_) GetValues(
	first uti.Index,
	last uti.Index,
) Sequential[byte] {
	var size = v.GetSize()
	var goFirst = uti.RelativeToZeroBased(first, size)
	var goLast = uti.RelativeToZeroBased(last, size)
	return binary_(v[goFirst : goLast+1])
}

func (v binary_) GetIndex(
	value byte,
) uti.Index {
	var index uti.Index
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

func (v binary_) String() string {
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
	alpha_        = "[A-Za-z]"
	alphanumeric_ = alpha_ + "|" + base10_
	base10_       = "[0-9]"
	base64_       = alphanumeric_ + "|[\\+/]"
	space_        = " "
)

// Instance Structure

type binary_ []byte

// Class Structure

type binaryClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func binaryClass() *binaryClass_ {
	return binaryClassReference_
}

var binaryClassReference_ = &binaryClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^'>(" + eol_ + "((?:" + space_ + ")*(?:" + base64_ + "){2,60}" +
			eol_ + ")+(?:" + space_ + ")*)?<'",
	),
}
