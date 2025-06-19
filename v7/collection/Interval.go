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

package collection

import (
	fmt "fmt"
	age "github.com/craterdog/go-component-framework/v7/agent"
	ele "github.com/craterdog/go-component-framework/v7/element"
	str "github.com/craterdog/go-component-framework/v7/string"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ref "reflect"
	sts "strings"
	syn "sync"
)

// CLASS INTERFACE

// Access Function

func IntervalClass[V ele.Discrete]() IntervalClassLike[V] {
	return intervalClass[V]()
}

// Constructor Methods

func (c *intervalClass_[V]) Interval(
	left Bracket,
	minimum V,
	maximum V,
	right Bracket,
) IntervalLike[V] {
	var instance = &interval_[V]{
		// Initialize the instance attributes.
		left_:    left,
		minimum_: minimum,
		maximum_: maximum,
		right_:   right,
	}
	instance.validateInterval()
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *interval_[V]) GetClass() IntervalClassLike[V] {
	return intervalClass[V]()
}

// Accessible[V] Methods

func (v *interval_[V]) GetValue(
	index uti.Index,
) V {
	var size = v.effectiveSize()
	var goIndex = uti.RelativeToZeroBased(index, size)
	var integer = v.effectiveMinimum() + goIndex
	return v.valueOf(integer)
}

func (v *interval_[V]) GetValues(
	first uti.Index,
	last uti.Index,
) str.Sequential[V] {
	var firstValue = v.GetValue(first)
	var lastValue = v.GetValue(last)
	return v.GetClass().Interval(
		Inclusive,
		firstValue,
		lastValue,
		Inclusive,
	)
}

func (v *interval_[V]) GetIndex(
	value V,
) uti.Index {
	var index = uti.Index(1)
	var iterator = v.GetIterator()
	for iterator.HasNext() {
		var candidate = iterator.GetNext()
		if candidate.AsInteger() == value.AsInteger() {
			return index
		}
		index++
	}
	return uti.Index(0)
}

// Bounded[V] Methods

func (v *interval_[V]) GetLeft() Bracket {
	return v.left_
}

func (v *interval_[V]) SetLeft(left Bracket) {
	v.left_ = left
	v.validateInterval()
}

func (v *interval_[V]) GetMinimum() V {
	return v.minimum_
}

func (v *interval_[V]) SetMinimum(minimum V) {
	v.minimum_ = minimum
	v.validateInterval()
}

func (v *interval_[V]) GetMaximum() V {
	return v.maximum_
}

func (v *interval_[V]) SetMaximum(maximum V) {
	v.maximum_ = maximum
	v.validateInterval()
}

func (v *interval_[V]) GetRight() Bracket {
	return v.right_
}

func (v *interval_[V]) SetRight(right Bracket) {
	v.right_ = right
	v.validateInterval()
}

// Searchable[V] Methods

func (v *interval_[V]) ContainsValue(
	value V,
) bool {
	if v.minimum_.IsDefined() && value.AsInteger() < v.minimum_.AsInteger() {
		return false
	}
	if v.maximum_.IsDefined() && value.AsInteger() > v.maximum_.AsInteger() {
		return false
	}
	return true
}

func (v *interval_[V]) ContainsAny(
	values str.Sequential[V],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This interval contains at least one of the values.
			return true
		}
	}
	// This interval does not contain any of the values.
	return false
}

func (v *interval_[V]) ContainsAll(
	values str.Sequential[V],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This interval is missing at least one of the values.
			return false
		}
	}
	// This interval does contains all of the values.
	return true
}

// str.Sequential[V] Methods

func (v *interval_[V]) IsEmpty() bool {
	return false
}

func (v *interval_[V]) GetSize() uti.Cardinal {
	var size = v.maximum_.AsInteger() - v.minimum_.AsInteger() + 1
	size -= int(v.left_)
	size -= int(v.right_)
	return uti.Cardinal(size)
}

func (v *interval_[V]) AsArray() []V {
	var size = int(v.GetSize())
	if size > 256 {
		// Limit the size to something reasonable.
		size = 256
	}
	var array = make([]V, size)
	for index := 0; index < size; index++ {
		var integer = v.minimum_.AsInteger() + index
		var value = v.valueOf(integer)
		array[index] = value
	}
	return array
}

func (v *interval_[V]) GetIterator() age.IteratorLike[V] {
	var array = v.AsArray()
	var iterator = age.IteratorClass[V]().Iterator(array)
	return iterator
}

// PROTECTED INTERFACE

func (v *interval_[V]) String() string {
	var string_ string
	switch v.left_ {
	case Inclusive:
		string_ += "["
	case Exclusive:
		string_ += "("
	}
	if v.minimum_.IsDefined() {
		string_ += v.minimum_.AsString()
	}
	string_ += ".."
	if v.maximum_.IsDefined() {
		string_ += v.maximum_.AsString()
	}
	switch v.right_ {
	case Inclusive:
		string_ += "]"
	case Exclusive:
		string_ += ")"
	}
	return string_
}

// Private Methods

func (v *interval_[V]) effectiveMaximum() int {
	var maximum = v.maximum_.AsInteger()
	maximum -= int(v.right_)
	return maximum
}

func (v *interval_[V]) effectiveMinimum() int {
	var minimum = v.minimum_.AsInteger()
	minimum += int(v.left_)
	return minimum
}

func (v *interval_[V]) effectiveSize() uti.Cardinal {
	var size = v.effectiveMaximum()
	size = size - v.effectiveMinimum() + 1
	return uti.Cardinal(size)
}

// This method ensures that the endpoints are valid.
func (v *interval_[V]) validateInterval() {
	// Validate the left bracket.
	switch v.left_ {
	case Inclusive:
	case Exclusive:
	default:
		var message = fmt.Sprintf(
			"Received an invalid left bracket for an interval: %v",
			v.left_,
		)
		panic(message)
	}

	// Validate the right bracket.
	switch v.right_ {
	case Inclusive:
	case Exclusive:
	default:
		var message = fmt.Sprintf(
			"Received an invalid right bracket for an interval: %v",
			v.right_,
		)
		panic(message)
	}

	// Validate the endpoints.
	if v.minimum_.IsDefined() && v.maximum_.IsDefined() {
		var collator = age.CollatorClass[V]().Collator()
		if collator.RankValues(v.minimum_, v.maximum_) != age.LesserRank {
			var message = fmt.Sprintf(
				"The minimum %v in an interval must be less than the maximum %v.",
				v.minimum_,
				v.maximum_,
			)
			panic(message)
		}
		var size = v.effectiveSize()
		if size <= 0 {
			var message = fmt.Sprintf(
				"The effective size of an interval must be greater than zero: %v.",
				size,
			)
			panic(message)
		}
	}
}

func (v *interval_[V]) valueOf(integer int) V {
	// We cannot get access to the constructor we need without reflection.
	var integerRef = ref.ValueOf(integer)
	var resultRef ref.Value
	var valueRef = ref.ValueOf(v.minimum_)
	var method = valueRef.MethodByName("GetClass")
	var classRef = method.Call([]ref.Value{})[0]
	var typeRef = classRef.Type()
	var count = typeRef.NumMethod()
	for index := 0; index < count; index++ {
		var name = typeRef.Method(index).Name
		if sts.HasSuffix(name, "FromInteger") {
			resultRef = classRef.MethodByName(name).Call([]ref.Value{integerRef})[0]
			break
		}
	}
	return resultRef.Interface().(V)
}

// Instance Structure

type interval_[V ele.Discrete] struct {
	// Declare the instance attributes.
	left_    Bracket
	minimum_ V
	maximum_ V
	right_   Bracket
}

// Class Structure

type intervalClass_[V ele.Discrete] struct {
	// Declare the class constants.
}

// Class Reference

var intervalMap_ = map[string]any{}
var intervalMutex_ syn.Mutex

func intervalClass[V ele.Discrete]() *intervalClass_[V] {
	// Generate the name of the bound class type.
	var class *intervalClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	intervalMutex_.Lock()
	var value = intervalMap_[name]
	switch actual := value.(type) {
	case *intervalClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &intervalClass_[V]{
			// Initialize the class constants.
		}
		intervalMap_[name] = class
	}
	intervalMutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}
