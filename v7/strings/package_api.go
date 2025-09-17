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
Package "strings" provides a framework of aspects and class definitions for a
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
package strings

import (
	age "github.com/craterdog/go-component-framework/v7/agents"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
)

// TYPE DECLARATIONS

/*
Identifier is a constrained type representing a string of the form:
(LOWER | UPPER) (LOWER | UPPER | DIGIT | '-')
*/
type Identifier string

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
		sequence Sequential[byte],
	) BinaryLike
	BinaryFromString(
		source string,
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
		sequence Sequential[Identifier],
	) NameLike
	NameFromString(
		source string,
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
		sequence Sequential[Line],
	) NarrativeLike
	NarrativeFromString(
		source string,
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
		sequence Sequential[Character],
	) PatternLike
	PatternFromString(
		source string,
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
		sequence Sequential[Character],
	) QuoteLike
	QuoteFromString(
		source string,
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
		size uti.Cardinal,
	) TagLike
	TagFromSequence(
		sequence Sequential[byte],
	) TagLike
	TagFromString(
		source string,
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
		ordinals []uti.Ordinal,
	) VersionLike
	VersionFromSequence(
		sequence Sequential[uti.Ordinal],
	) VersionLike
	VersionFromString(
		source string,
	) VersionLike

	// Function Methods
	IsValidNextVersion(
		current VersionLike,
		next VersionLike,
	) bool
	GetNextVersion(
		current VersionLike,
		level uti.Cardinal,
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
	Accessible[byte]
	Searchable[byte]
	Sequential[byte]
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
	Accessible[Identifier]
	Searchable[Identifier]
	Sequential[Identifier]
	Spectral[NameLike]
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
	Accessible[Line]
	Searchable[Line]
	Sequential[Line]
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
	Accessible[Character]
	Searchable[Character]
	Sequential[Character]
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
	Accessible[Character]
	Searchable[Character]
	Sequential[Character]
	Spectral[QuoteLike]
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
	Accessible[byte]
	Searchable[byte]
	Sequential[byte]
}

/*
VersionLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete version-like class.
*/
type VersionLike interface {
	// Principal Methods
	GetClass() VersionClassLike
	AsIntrinsic() []uti.Ordinal
	AsString() string

	// Aspect Interfaces
	Accessible[uti.Ordinal]
	Searchable[uti.Ordinal]
	Sequential[uti.Ordinal]
	Spectral[VersionLike]
}

// ASPECT DECLARATIONS

/*
Accessible[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of an accessible concrete
class.

An accessible class maintains values that can be accessed using indices. The
indices of an accessible sequence are ORDINAL rather than ZERO based—which
never really made sense except for pointer offsets. What is the "zeroth
value" anyway? It's the "first value", right?  So we start fresh...

This approach allows for positive indices starting at the beginning of the
sequence, and negative indices starting at the end of the sequence as follows:

	    1           2           3             N
	[value 1] . [value 2] . [value 3] ... [value N]
	   -N        -(N-1)      -(N-2)          -1

Notice that because the indices are ordinal based, the positive and negative
indices are symmetrical.
*/
type Accessible[V any] interface {
	GetValue(
		index uti.Index,
	) V
	GetValues(
		first uti.Index,
		last uti.Index,
	) Sequential[V]
	GetIndex(
		value V,
	) uti.Index
}

/*
Searchable[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a searchable concrete
class.
*/
type Searchable[V any] interface {
	ContainsValue(
		value V,
	) bool
	ContainsAny(
		values Sequential[V],
	) bool
	ContainsAll(
		values Sequential[V],
	) bool
}

/*
Sequential[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a sequential concrete
class.
*/
type Sequential[V any] interface {
	IsEmpty() bool
	GetSize() uti.Cardinal
	AsArray() []V
	GetIterator() age.IteratorLike[V]
}

/*
Spectral[V any] is an aspect interface that declares a set of method signatures
that must be supported by each instance of a spectral concrete class.
*/
type Spectral[V any] interface {
	AsString() string
	CompareWith(
		value V,
	) age.Rank
}
