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

func AngleClass() AngleClassLike {
	return angleClass()
}

// Constructor Methods

func (c *angleClass_) Angle(
	radians float64,
) AngleLike {
	return c.angleFromFloat(radians)
}

func (c *angleClass_) AngleFromString(
	string_ string,
) AngleLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the angle constructor method: %s",
			string_,
		)
		panic(message)
	}
	var float float64
	var match = matches[1] // Strip off the leading '~' character.
	switch match {
	case "pi", "π":
		float = mat.Pi
	case "tau", "τ":
		float = mat.Pi * 2.0
	default:
		float, _ = stc.ParseFloat(match, 64)
	}
	return c.angleFromFloat(float)
}

// Constant Methods

func (c *angleClass_) Minimum() AngleLike {
	return c.minimum_
}

func (c *angleClass_) Maximum() AngleLike {
	return c.maximum_
}

func (c *angleClass_) Zero() AngleLike {
	return c.zero_
}

func (c *angleClass_) Pi() AngleLike {
	return c.pi_
}

func (c *angleClass_) Tau() AngleLike {
	return c.tau_
}

// Function Methods

func (c *angleClass_) Inverse(
	angle AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(angle.AsFloat() - angleClass().Pi().AsFloat())
	return result_
}

func (c *angleClass_) Sum(
	first AngleLike,
	second AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(first.AsFloat() + second.AsFloat())
	return result_
}

func (c *angleClass_) Difference(
	first AngleLike,
	second AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(first.AsFloat() - second.AsFloat())
	return result_
}

func (c *angleClass_) Scaled(
	angle AngleLike,
	factor float64,
) AngleLike {
	var result_ = c.angleFromFloat(angle.AsFloat() * factor)
	return result_
}

func (c *angleClass_) Complement(
	angle AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(angleClass().Pi().AsFloat()/2.0 - angle.AsFloat())
	return result_
}

func (c *angleClass_) Supplement(
	angle AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(angleClass().Pi().AsFloat() - angle.AsFloat())
	return result_
}

func (c *angleClass_) Conjugate(
	angle AngleLike,
) AngleLike {
	var result_ = c.angleFromFloat(-angle.AsFloat())
	return result_
}

func (c *angleClass_) Cosine(
	angle AngleLike,
) float64 {
	var result_ float64
	switch angle.AsFloat() {
	case 0.0:
		result_ = 1.0
	case mat.Pi * 0.25:
		result_ = 0.5 * mat.Sqrt2
	case mat.Pi * 0.5:
		result_ = 0.0
	case mat.Pi * 0.75:
		result_ = -0.5 * mat.Sqrt2
	case mat.Pi:
		result_ = -1.0
	case mat.Pi * 1.25:
		result_ = -0.5 * mat.Sqrt2
	case mat.Pi * 1.5:
		result_ = 0.0
	case mat.Pi * 1.75:
		result_ = 0.5 * mat.Sqrt2
	case mat.Pi * 2.0:
		result_ = 1.0
	default:
		result_ = mat.Cos(angle.AsFloat())
	}
	return result_
}

func (c *angleClass_) ArcCosine(
	x float64,
) AngleLike {
	var result_ = c.angleFromFloat(mat.Acos(x))
	return result_
}

func (c *angleClass_) Sine(
	angle AngleLike,
) float64 {
	var result_ float64
	switch angle.AsFloat() {
	case 0.0:
		result_ = 0.0
	case mat.Pi * 0.25:
		result_ = 0.5 * mat.Sqrt2
	case mat.Pi * 0.5:
		result_ = 1.0
	case mat.Pi * 0.75:
		result_ = 0.5 * mat.Sqrt2
	case mat.Pi:
		result_ = 0.0
	case mat.Pi * 1.25:
		result_ = -0.5 * mat.Sqrt2
	case mat.Pi * 1.5:
		result_ = -1.0
	case mat.Pi * 1.75:
		result_ = -0.5 * mat.Sqrt2
	case mat.Pi * 2.0:
		result_ = 0.0
	default:
		result_ = mat.Sin(angle.AsFloat())
	}
	return result_
}

func (c *angleClass_) ArcSine(
	y float64,
) AngleLike {
	var result_ = c.angleFromFloat(mat.Asin(y))
	return result_
}

func (c *angleClass_) Tangent(
	angle AngleLike,
) float64 {
	var result_ float64
	switch angle.AsFloat() {
	case 0.0:
		result_ = 0.0
	case mat.Pi * 0.25:
		result_ = 1.0
	case mat.Pi * 0.5:
		result_ = mat.Inf(1)
	case mat.Pi * 0.75:
		result_ = -1.0
	case mat.Pi:
		result_ = 0.0
	case mat.Pi * 1.25:
		result_ = 1.0
	case mat.Pi * 1.5:
		result_ = mat.Inf(1)
	case mat.Pi * 1.75:
		result_ = -1.0
	case mat.Pi * 2.0:
		result_ = 0.0
	default:
		result_ = mat.Tan(angle.AsFloat())
	}
	return result_
}

func (c *angleClass_) ArcTangent(
	x float64,
	y float64,
) AngleLike {
	var result_ = c.angleFromFloat(mat.Atan2(y, x))
	return result_
}

// INSTANCE INTERFACE

// Principal Methods

func (v angle_) GetClass() AngleClassLike {
	return angleClass()
}

func (v angle_) AsIntrinsic() float64 {
	return float64(v)
}

func (v angle_) AsString() string {
	var result_ = angleClass().stringFromAngle(v)
	return result_
}

func (v angle_) AsUnits(
	units Units,
) float64 {
	var result_ float64
	var radians = v.AsIntrinsic()
	var pi = angleClass().pi_.AsIntrinsic()
	var tau = angleClass().tau_.AsIntrinsic()
	switch units {
	case Degrees:
		result_ = 360.0 * radians / tau
	case Radians:
		result_ = radians
	case Gradians:
		result_ = 200.0 / pi
	}
	return result_
}

// Attribute Methods

func (v angle_) GetParts() (
	x float64,
	y float64,
) {
	var complex_ = v.AsIntrinsic()
	x = mat.Cos(complex_)
	y = mat.Sin(complex_)
	return
}

// Continuous Methods

func (v angle_) AsFloat() float64 {
	return float64(v)
}

func (v angle_) IsZero() bool {
	return v == angleClass().zero_ || v == angleClass().tau_
}

func (v angle_) IsInfinite() bool {
	return false
}

func (v angle_) IsUndefined() bool {
	return false
}

func (v angle_) HasMagnitude() bool {
	return !v.IsZero()
}

// PROTECTED INTERFACE

func (v angle_) String() string {
	return v.AsString()
}

// Private Methods

func (c *angleClass_) angleFromFloat(float float64) angle_ {
	float = c.normalizeValue(float)
	float = c.lockPhase(float)
	return angle_(float)
}

func (c *angleClass_) lockPhase(value float64) float64 {
	var pi = angleClass().Pi().AsIntrinsic()
	var value32 = float32(value)
	switch {
	case mat.Abs(value) <= 1.2246467991473515e-16:
		value = 0
	case value32 == float32(0.5*pi):
		value = 0.5 * pi
	case value32 == float32(pi):
		value = pi
	case value32 == float32(1.5*pi):
		value = 1.5 * pi
	}
	return value
}

func (c *angleClass_) normalizeValue(value float64) float64 {
	var tau = angleClass().Tau().AsIntrinsic()
	if value < -tau || value >= tau {
		// Normalize the value to the range [-τ..τ).
		value = mat.Remainder(value, tau)
	}
	if value < 0.0 {
		// Normalize the value to the range [0..τ).
		value = value + tau
	}
	return value
}

func (c *angleClass_) stringFromAngle(angle angle_) string {
	var string_ string
	switch angle {
	case c.pi_:
		string_ = "~π"
	case c.tau_:
		string_ = "~τ"
	default:
		string_ = "~" + stc.FormatFloat(float64(angle), 'G', -1, 64)
	}
	return string_
}

// Instance Structure

type angle_ float64

// Class Structure

type angleClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
	minimum_ AngleLike
	maximum_ AngleLike
	zero_    AngleLike
	pi_      AngleLike
	tau_     AngleLike
}

// Class Reference

func angleClass() *angleClass_ {
	return angleClassReference_
}

var angleClassReference_ = &angleClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^~(0|" + amplitude_ + ")"),
	zero_:    angle_(0.0),
	pi_:      angle_(mat.Pi),
	tau_:     angle_(2.0 * mat.Pi),
}
