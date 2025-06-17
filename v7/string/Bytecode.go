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

func BytecodeClass() BytecodeClassLike {
	return bytecodeClass()
}

// Constructor Methods

func (c *bytecodeClass_) Bytecode(
	instructions []Instruction,
) BytecodeLike {
	return bytecode_(instructions)
}

func (c *bytecodeClass_) BytecodeFromSequence(
	sequence Sequential[Instruction],
) BytecodeLike {
	return bytecode_(sequence.AsArray())
}

func (c *bytecodeClass_) BytecodeFromString(
	string_ string,
) BytecodeLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the bytecode constructor method: %s",
			string_,
		)
		panic(message)
	}
	var base16 = matches[1]                   // Strip off the "'" delimiters.
	base16 = sts.ReplaceAll(base16, " ", "")  // Remove all spaces.
	base16 = sts.ReplaceAll(base16, "\n", "") // Remove all newlines.
	var encoder = age.EncoderClass().Encoder()
	var bytes = encoder.Base16Decode(base16)
	var instructions = make(
		[]Instruction,
		len(bytes)/2,
	)
	var index int
	for index < len(bytes)-1 {
		var firstByte = Instruction(bytes[index]) << 8
		index++
		var secondByte = Instruction(bytes[index])
		index++
		instructions[index/2-1] = firstByte | secondByte
	}
	return bytecode_(instructions)
}

// Constant Methods

// Function Methods

func (c *bytecodeClass_) Concatenate(
	first BytecodeLike,
	second BytecodeLike,
) BytecodeLike {
	var firstInstructions = first.AsArray()
	var secondInstructions = second.AsArray()
	var allInstructions = make(
		[]Instruction,
		len(firstInstructions)+len(secondInstructions),
	)
	copy(allInstructions, firstInstructions)
	copy(allInstructions[len(firstInstructions):], secondInstructions)
	return c.Bytecode(allInstructions)
}

// INSTANCE INTERFACE

// Principal Methods

func (v bytecode_) GetClass() BytecodeClassLike {
	return bytecodeClass()
}

func (v bytecode_) AsIntrinsic() []Instruction {
	return []Instruction(v)
}

func (v bytecode_) AsString() string {
	var string_ = "'"
	var size = len(v)
	if size > 0 {
		var index = 0
		var instruction = v[index]
		string_ += instruction.String()
		for index++; index < size; index++ {
			instruction = v[index]
			string_ += " " + instruction.String()
		}
	}
	string_ += "'"
	return string_
}

// Attribute Methods

// Sequential[Instruction] Methods

func (v bytecode_) IsEmpty() bool {
	return len(v) == 0
}

func (v bytecode_) GetSize() uti.Cardinal {
	return uti.Cardinal(len(v))
}

func (v bytecode_) AsArray() []Instruction {
	return uti.CopyArray(v)
}

func (v bytecode_) GetIterator() age.IteratorLike[Instruction] {
	var array = uti.CopyArray(v)
	var class = age.IteratorClass[Instruction]()
	var iterator = class.Iterator(array)
	return iterator
}

// Accessible[Instruction] Methods

func (v bytecode_) GetValue(
	index uti.Index,
) Instruction {
	var goIndex = uti.RelativeToZeroBased(v, index)
	return v[goIndex]
}

func (v bytecode_) GetValues(
	first uti.Index,
	last uti.Index,
) Sequential[Instruction] {
	var goFirst = uti.RelativeToZeroBased(v, first)
	var goLast = uti.RelativeToZeroBased(v, last)
	return bytecode_(v[goFirst : goLast+1])
}

// PROTECTED INTERFACE

func (v Instruction) String() string {
	return fmt.Sprintf("%04x", uint16(v))
}

func (v bytecode_) String() string {
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
	base16_      = base10_ + "|[a-f]"
	instruction_ = "(?:" + base16_ + "){4}"
	space_       = " "
)

// Instance Structure

type bytecode_ []Instruction

// Class Structure

type bytecodeClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func bytecodeClass() *bytecodeClass_ {
	return bytecodeClassReference_
}

var bytecodeClassReference_ = &bytecodeClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^'(" + instruction_ + "(?:" + space_ +
			instruction_ + ")*)'",
	),
}
