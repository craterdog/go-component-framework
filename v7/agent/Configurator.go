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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ConfiguratorClass() ConfiguratorClassLike {
	return configuratorClass()
}

// Constructor Methods

func (c *configuratorClass_) Configurator(
	file string,
) ConfiguratorLike {
	if uti.IsUndefined(file) {
		panic("The \"file\" attribute is required by this class.")
	}
	var instance = &configurator_{
		// Initialize the instance attributes.
		file_: file,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *configurator_) GetClass() ConfiguratorClassLike {
	return configuratorClass()
}

func (v *configurator_) ConfigurationExists() bool {
	return uti.PathExists(v.file_)
}

func (v *configurator_) LoadConfiguration() string {
	return uti.ReadFile(v.file_)
}

func (v *configurator_) StoreConfiguration(
	configuration string,
) {
	uti.WriteFile(v.file_, configuration)
}

func (v *configurator_) DeleteConfiguration() {
	uti.RemovePath(v.file_)
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type configurator_ struct {
	// Declare the instance attributes.
	file_ string
}

// Class Structure

type configuratorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func configuratorClass() *configuratorClass_ {
	return configuratorClassReference_
}

var configuratorClassReference_ = &configuratorClass_{
	// Initialize the class constants.
}
