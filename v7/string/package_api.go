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

/*
Package "string" provides a framework of aspects and class definitions for a
rich set of primitive data types that can be iterated over.  All primitive types
are immutable and—for better performance—are implemented as extensions to
existing Go primitive types.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-component-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package string

import (
	age "github.com/craterdog/go-component-framework/v7/agent"
	col "github.com/craterdog/go-component-framework/v7/collection"
	reg "regexp"
)

// TYPE DECLARATIONS

/*
Identifier is a constrained type representing a string of the form:
(LOWER | UPPER) (LOWER | UPPER | DIGIT | '-')
*/
type Identifier string

/*
Instruction is a constrained type representing a single bytecode instruction.
*/
type Instruction uint16

/*
Line is a constrained type representing a single line of a narrative.
*/
type Line string

/*
Character is a constrained type representing a single character of a quote.
*/
type Character rune

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
BinaryClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
binary-like concrete class.
*/
type BinaryClassLike interface {
	// Constructor Methods
	Binary(
		bytes []byte,
	) BinaryLike
	BinaryFromSequence(
		sequence col.Sequential[byte],
	) BinaryLike
	BinaryFromString(
		string_ string,
	) BinaryLike

	// Function Methods
	Not(
		binary BinaryLike,
	) BinaryLike
	And(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	San(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Ior(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Xor(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
	Concatenate(
		first BinaryLike,
		second BinaryLike,
	) BinaryLike
}

/*
BytecodeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
bytecode-like concrete class.
*/
type BytecodeClassLike interface {
	// Constructor Methods
	Bytecode(
		instructions []Instruction,
	) BytecodeLike
	BytecodeFromSequence(
		sequence col.Sequential[Instruction],
	) BytecodeLike
	BytecodeFromString(
		string_ string,
	) BytecodeLike

	// Function Methods
	Concatenate(
		first BytecodeLike,
		second BytecodeLike,
	) BytecodeLike
}

/*
NameClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
name-like concrete class.
*/
type NameClassLike interface {
	// Constructor Methods
	Name(
		identifiers []Identifier,
	) NameLike
	NameFromSequence(
		sequence col.Sequential[Identifier],
	) NameLike
	NameFromString(
		string_ string,
	) NameLike

	// Function Methods
	Concatenate(
		first NameLike,
		second NameLike,
	) NameLike
}

/*
NarrativeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
narrative-like concrete class.
*/
type NarrativeClassLike interface {
	// Constructor Methods
	Narrative(
		lines []Line,
	) NarrativeLike
	NarrativeFromSequence(
		sequence col.Sequential[Line],
	) NarrativeLike
	NarrativeFromString(
		string_ string,
	) NarrativeLike

	// Function Methods
	Concatenate(
		first NarrativeLike,
		second NarrativeLike,
	) NarrativeLike
}

/*
PatternClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
pattern-like concrete class.
*/
type PatternClassLike interface {
	// Constructor Methods
	Pattern(
		characters []Character,
	) PatternLike
	PatternFromSequence(
		sequence col.Sequential[Character],
	) PatternLike
	PatternFromString(
		string_ string,
	) PatternLike

	// Constant Methods
	None() PatternLike
	Any() PatternLike

	// Function Methods
	Concatenate(
		first PatternLike,
		second PatternLike,
	) PatternLike
}

/*
QuoteClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
quote-like concrete class.
*/
type QuoteClassLike interface {
	// Constructor Methods
	Quote(
		characters []Character,
	) QuoteLike
	QuoteFromSequence(
		sequence col.Sequential[Character],
	) QuoteLike
	QuoteFromString(
		string_ string,
	) QuoteLike

	// Function Methods
	Concatenate(
		first QuoteLike,
		second QuoteLike,
	) QuoteLike
}

/*
TagClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
tag-like concrete class.
*/
type TagClassLike interface {
	// Constructor Methods
	Tag(
		bytes []byte,
	) TagLike
	TagWithSize(
		size age.Cardinal,
	) TagLike
	TagFromSequence(
		sequence col.Sequential[byte],
	) TagLike
	TagFromString(
		string_ string,
	) TagLike

	// Function Methods
	Concatenate(
		first TagLike,
		second TagLike,
	) TagLike
}

/*
VersionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
version-like concrete class.
*/
type VersionClassLike interface {
	// Constructor Methods
	Version(
		ordinals []age.Ordinal,
	) VersionLike
	VersionFromSequence(
		sequence col.Sequential[age.Ordinal],
	) VersionLike
	VersionFromString(
		string_ string,
	) VersionLike

	// Function Methods
	IsValidNextVersion(
		current VersionLike,
		next VersionLike,
	) bool
	GetNextVersion(
		current VersionLike,
		level age.Ordinal,
	) VersionLike
	Concatenate(
		first VersionLike,
		second VersionLike,
	) VersionLike
}

// INSTANCE DECLARATIONS

/*
BinaryLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete binary-like class.
*/
type BinaryLike interface {
	// Principal Methods
	GetClass() BinaryClassLike
	AsIntrinsic() []byte
	AsString() string

	// Aspect Interfaces
	col.Accessible[byte]
	col.Sequential[byte]
}

/*
BytecodeLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete bytecode-like class.
*/
type BytecodeLike interface {
	// Principal Methods
	GetClass() BytecodeClassLike
	AsIntrinsic() []Instruction
	AsString() string

	// Aspect Interfaces
	col.Accessible[Instruction]
	col.Sequential[Instruction]
}

/*
NameLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete name-like class.
*/
type NameLike interface {
	// Principal Methods
	GetClass() NameClassLike
	AsIntrinsic() []Identifier
	AsString() string

	// Aspect Interfaces
	col.Accessible[Identifier]
	col.Sequential[Identifier]
}

/*
NarrativeLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete narrative-like class.
*/
type NarrativeLike interface {
	// Principal Methods
	GetClass() NarrativeClassLike
	AsIntrinsic() []Line
	AsString() string

	// Aspect Interfaces
	col.Accessible[Line]
	col.Sequential[Line]
}

/*
PatternLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a pattern-like elemental class.
*/
type PatternLike interface {
	// Principal Methods
	GetClass() PatternClassLike
	AsIntrinsic() []Character
	AsString() string
	AsRegexp() *reg.Regexp
	MatchesText(
		text string,
	) bool
	GetMatches(
		text string,
	) []string

	// Aspect Interfaces
	col.Accessible[Character]
	col.Sequential[Character]
}

/*
QuoteLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete quote-like class.
*/
type QuoteLike interface {
	// Principal Methods
	GetClass() QuoteClassLike
	AsIntrinsic() []Character
	AsString() string

	// Aspect Interfaces
	col.Accessible[Character]
	col.Sequential[Character]
}

/*
TagLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete tag-like class.
*/
type TagLike interface {
	// Principal Methods
	GetClass() TagClassLike
	AsIntrinsic() []byte
	AsString() string
	GetHash() uint64

	// Aspect Interfaces
	col.Accessible[byte]
	col.Sequential[byte]
}

/*
VersionLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete version-like class.
*/
type VersionLike interface {
	// Principal Methods
	GetClass() VersionClassLike
	AsIntrinsic() []age.Ordinal
	AsString() string

	// Aspect Interfaces
	col.Accessible[age.Ordinal]
	col.Sequential[age.Ordinal]
}

// ASPECT DECLARATIONS
