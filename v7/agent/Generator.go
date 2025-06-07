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
	var random, _ = ran.Int(ran.Reader, big.NewInt(int64(2)))
	return int(random.Int64()) > 0
}

func (v *generator_) RandomOrdinal(
	maximum Ordinal,
) Ordinal {
	var random, _ = ran.Int(ran.Reader, big.NewInt(int64(maximum)))
	return Ordinal(random.Int64() + 1)
}

func (v *generator_) RandomProbability() Probability {
	var maximum = Ordinal(1 << 53) // 53 bits for the sign and mantissa.
	return Probability(float64(v.RandomOrdinal(maximum)) / float64(maximum))
}

func (v *generator_) RandomBytes(
	size Cardinal,
) []byte {
	var bytes = make([]byte, size)
	_, _ = ran.Read(bytes) // This call can never fail.
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
