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

package element

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func CitationClass() CitationClassLike {
	return citationClass()
}

// Constructor Methods

func (c *citationClass_) Citation(
	string_ string,
) CitationLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the citation constructor method: %s",
			string_,
		)
		panic(message)
	}
	return citation_(matches[0])
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v citation_) GetClass() CitationClassLike {
	return citationClass()
}

func (v citation_) AsIntrinsic() string {
	return string(v)
}

func (v citation_) AsString() string {
	return string(v)
}

// Attribute Methods

func (v citation_) GetName() string {
	var index = sts.LastIndex(string(v), "@")
	return string(v[:index])
}

func (v citation_) GetVersion() string {
	var index = sts.LastIndex(string(v), "@")
	return string(v[index+1:])
}

// PROTECTED INTERFACE

func (v citation_) String() string {
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
	identifier_ = "(?:" + letter_ + ")(?:" + letter_ + "|" + digit_ + ")*"
	letter_     = lower_ + "|" + upper_
	lower_      = "\\p{Ll}"
	name_       = "(?:/" + identifier_ + ")+"
	upper_      = "\\p{Lu}"
	version_    = "v" + ordinal_ + "(?:\\." + ordinal_ + ")*"
)

// Instance Structure

type citation_ string

// Class Structure

type citationClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func citationClass() *citationClass_ {
	return citationClassReference_
}

var citationClassReference_ = &citationClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^(" + name_ + ")@(" + version_ + ")"),
}
