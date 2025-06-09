/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	ran "crypto/rand"
	big "math/big"
)

// CLASS INTERFACE

// Access Function

func GeneratorClass() GeneratorClassLike {
	return generatorClass()
}

// Constructor Methods

func (c *generatorClass_) Generator() GeneratorLike {
	var instance = &generator_{
		// Initialize the instance attributes.
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *generator_) GetClass() GeneratorClassLike {
	return generatorClass()
}

func (v *generator_) RandomBoolean() bool {
	// A random boolean is in the range [0..1].
	var random, _ = ran.Int(ran.Reader, big.NewInt(int64(2)))
	return random.Int64() > 0
}

func (v *generator_) RandomOrdinal(
	maximum Ordinal,
) Ordinal {
	// A random integer is in the range [0..maximum).
	var random, _ = ran.Int(ran.Reader, big.NewInt(int64(maximum)))
	// Convert [0..maximum) to [1..maximum].
	return Ordinal(random.Uint64() + 1)
}

func (v *generator_) RandomProbability() float64 {
	// Use 53 bits for the sign and mantissa only.
	var maximum = Ordinal(1 << 53)
	// A random probability is in the range (0.0..1.0] since something with
	// zero probability will never occur so we use [1..maximum]/maximum.
	return float64(v.RandomOrdinal(maximum)) / float64(maximum)
}

func (v *generator_) RandomBytes(
	size Cardinal,
) []byte {
	var bytes = make([]byte, size)
	_, _ = ran.Read(bytes) // This call should never fail.
	return bytes
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type generator_ struct {
	// Declare the instance attributes.
}

// Class Structure

type generatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func generatorClass() *generatorClass_ {
	return generatorClassReference_
}

var generatorClassReference_ = &generatorClass_{
	// Initialize the class constants.
}
