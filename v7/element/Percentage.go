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
	mat "math"
	reg "regexp"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func PercentageClass() PercentageClassLike {
	return percentageClass()
}

// Constructor Methods

func (c *percentageClass_) Percentage(
	float float64,
) PercentageLike {
	return percentage_(float / 100.0)
}

func (c *percentageClass_) PercentageFromInteger(
	integer int,
) PercentageLike {
	var float = float64(integer)
	return percentage_(float / 100.0)
}

func (c *percentageClass_) PercentageFromString(
	string_ string,
) PercentageLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the percentage constructor method: %s",
			string_,
		)
		panic(message)
	}
	var float, _ = stc.ParseFloat(matches[1], 64) // Strip off the '%' suffix.
	return percentage_(float / 100.0)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v percentage_) GetClass() PercentageClassLike {
	return percentageClass()
}

func (v percentage_) GetIntrinsic() float64 {
	return float64(v)
}

// Attribute Methods

// Continuous Methods

func (v percentage_) AsFloat() float64 {
	return float64(v * 100.0)
}

func (v percentage_) IsZero() bool {
	return v == 0
}

func (v percentage_) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

func (v percentage_) IsUndefined() bool {
	return mat.IsNaN(float64(v))
}

func (v percentage_) HasMagnitude() bool {
	return !(v.IsZero() || v.IsInfinite() || v.IsUndefined())
}

// Discrete Methods

func (v percentage_) AsBoolean() bool {
	return v != 0
}

func (v percentage_) AsInteger() int {
	return int(float64(v) * 100.0)
}

// Lexical Methods

func (v percentage_) AsString() string {
	return numberClass().stringFromFloat(float64(v)*100.0) + "%"
}

// Polarized Methods

func (v percentage_) IsNegative() bool {
	return v < 0
}

// PROTECTED INTERFACE

func (v percentage_) String() string {
	return v.AsString()
}

// Private Methods

// Instance Structure

type percentage_ float64

// Class Structure

type percentageClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func percentageClass() *percentageClass_ {
	return percentageClassReference_
}

var percentageClassReference_ = &percentageClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^(" + real_ + ")%"),
}
