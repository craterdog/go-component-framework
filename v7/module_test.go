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

package module_test

import (
	fmt "fmt"
	fra "github.com/craterdog/go-component-framework/v7"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	cmp "math/cmplx"
	syn "sync"
	tes "testing"
)

func TestModuleFunctions(t *tes.T) {
	fra.Collator[any]()
	fra.CollatorWithMaximumDepth[any](8)
	fra.Iterator[any]([]any{"foo", 5})
	var sorter = fra.Sorter[any]()
	fra.SorterWithRanker[any](sorter.GetRanker())
	fra.List[string]()
	var list = fra.ListFromArray[string]([]string{"A"})
	fra.ListFromSequence[string](list)
	fra.ListClass[string]().Concatenate(list, list)
	var association = fra.Association[string, int]("A", 1)
	var catalog = fra.Catalog[string, int]()
	fra.CatalogFromArray[string, int]([]fra.AssociationLike[string, int]{association})
	fra.CatalogFromMap[string, int](catalog.AsMap())
	fra.CatalogFromSequence[string, int](catalog)
	fra.CatalogClass[string, int]().Extract(catalog, list)
	fra.CatalogClass[string, int]().Merge(catalog, catalog)
	fra.Queue[string]()
	fra.QueueWithCapacity[string](8)
	var queue = fra.QueueFromArray[string](list.AsArray())
	fra.QueueFromSequence[string](queue)
	var group = new(syn.WaitGroup)
	defer group.Wait()
	var queues = fra.QueueClass[string]().Fork(group, queue, 2)
	fra.QueueClass[string]().Split(group, queue, 2)
	fra.QueueClass[string]().Join(group, queues)
	queue.CloseChannel()
	var set = fra.Set[string]()
	fra.SetWithCollator[string](set.GetCollator())
	fra.SetFromArray[string](set.AsArray())
	fra.SetFromSequence[string](set)
	fra.SetClass[string]().And(set, set)
	fra.SetClass[string]().Ior(set, set)
	fra.SetClass[string]().San(set, set)
	fra.SetClass[string]().Xor(set, set)
	fra.Stack[string]()
	fra.StackWithCapacity[string](8)
	fra.StackFromArray[string](list.AsArray())
	fra.StackFromSequence[string](list)
}

func TestModuleExampleCode(t *tes.T) {
	fmt.Println("MODULE EXAMPLE:")

	// Create an empty list.
	var list = fra.List[string]()
	fmt.Printf("An empty list: %v\n", list)
	fmt.Println()

	// Create a list using an intrinsic Go array of values.
	list = fra.ListFromArray[string](
		[]string{"Hello", "World"},
	)
	fmt.Printf("A list: %v\n", list)
	fmt.Println()

	// Create an empty catalog.
	var catalog = fra.Catalog[string, int64]()
	fmt.Printf("An empty catalog: %v\n", catalog)
	fmt.Println()

	// Create a catalog from an intrinsic Go map.
	catalog = fra.CatalogFromMap[string, int64](
		map[string]int64{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		},
	)
	fmt.Printf("A catalog: %v\n", catalog)
	fmt.Println()

	// Create a list of the catalog keys.
	list = fra.ListFromSequence[string](catalog.GetKeys())
	fmt.Printf("A list of keys: %v\n", list)
	fmt.Println()

	// Create a set from an intrinsic Go array of values.
	var set = fra.SetFromArray[string](
		[]string{"a", "b", "r", "a", "c", "a", "d", "a", "b", "r", "a"},
	)
	fmt.Printf("A set: %v\n", set)
	fmt.Println()

	// Create an empty stack with a capacity of 4.
	var stack = fra.StackWithCapacity[string](4)
	fmt.Printf("An empty stack: %v\n", stack)
	fmt.Println()

	// Create a stack containing the values from a list.
	stack = fra.StackFromSequence[string](list)
	fmt.Printf("A stack: %v\n", stack)
	fmt.Println()

	// Create an empty queue with a capacity of 5.
	var queue = fra.QueueWithCapacity[string](5)
	fmt.Printf("An empty queue: %v\n", queue)
	fmt.Println()

	// Create a queue containing the values from a set.
	queue = fra.QueueFromSequence[string](set)
	fmt.Printf("A queue: %v\n", queue)
	fmt.Println()
}

func TestListExampleCode(t *tes.T) {
	fmt.Println("LIST EXAMPLE:")

	// Create a new list from an array.
	var list = fra.ListFromArray[string](
		[]string{"bar", "foo", "bax"},
	)
	fmt.Println("The initialized list:", list)
	fmt.Println()

	// Add some more values to the list.
	list.AppendValue("fuz")
	list.AppendValue("box")
	fmt.Println("The augmented list:", list)
	fmt.Println()

	// Change a value in the list.
	list.SetValue(2, "foz")
	fmt.Println("The updated list:", list)
	fmt.Println()

	// Insert a new value at the beginning of the list (slot 0).
	list.InsertValue(0, "bax")
	fmt.Println("The updated list:", list)
	fmt.Println()

	// Insert a new value at the end of the list (slot N).
	list.InsertValue(6, "bux")
	fmt.Println("The updated list:", list)
	fmt.Println()

	// Sort the values in the list.
	list.SortValues()
	fmt.Println("The sorted list:", list)
	fmt.Println()

	// Randomly shuffle the values in the list.
	list.ShuffleValues()
	fmt.Println("The shuffled list:", list)
	fmt.Println()

	// Remove a value from the list.
	list.RemoveValue(4)
	fmt.Println("The shortened list:", list)
	fmt.Println()

	// Remove all values from the list.
	list.RemoveAll()
	fmt.Println("The empty list:", list)
	fmt.Println()
}

func TestSetExampleCode(t *tes.T) {
	fmt.Println("SET EXAMPLE:")

	// Create two sets with overlapping values.
	var set1 = fra.SetFromArray[string](
		[]string{"alpha", "beta", "gamma"},
	)
	fmt.Println("The first set is:", set1)
	fmt.Println()
	var set2 = fra.SetFromArray[string](
		[]string{"beta", "gamma", "delta"},
	)
	fmt.Println("The second set is:", set2)
	fmt.Println()

	// Find the logical union of the two sets.
	var Set = set1.GetClass()
	fmt.Println("The union of the two sets is:", Set.Ior(set1, set2))
	fmt.Println()

	// Find the logical difference between the two sets.
	fmt.Println("The first set minus the second set is:", Set.San(set1, set2))
	fmt.Println()

	// Find the logical intersection of the two sets.
	fmt.Println("The intersection of the two sets is:", Set.And(set1, set2))
	fmt.Println()

	// Find the logical exclusion of the two sets.
	fmt.Println("The exclusion of the two sets is:", Set.Xor(set1, set2))
	fmt.Println()

	// Add an existing value to the first set.
	set1.AddValue("beta")
	fmt.Println("Adding an existing value to a set does not change it:", set1)
	fmt.Println()
}

func TestStackExampleCode(t *tes.T) {
	fmt.Println("STACK EXAMPLE:")

	// Create a new empty stack.
	var stack = fra.Stack[string]()
	fmt.Println("The empty stack:", stack)
	fmt.Println()

	// Add some values to it.
	stack.AddValue("foo")
	fmt.Println("The stack with one value on it:", stack)
	fmt.Println()
	stack.AddValue("bar")
	fmt.Println("The stack with two values on it:", stack)
	fmt.Println()
	stack.AddValue("baz")
	fmt.Println("The stack with three values on it:", stack)
	fmt.Println()

	// Remove the top value from the stack.
	var top = stack.RemoveLast()
	fmt.Println("The top value was:", top)
	fmt.Println()
	fmt.Println("The stack with only two values on it:", stack)
	fmt.Println()

	// Remove all values from the stack.
	stack.RemoveAll()
	fmt.Println("The stack with no more values on it:", stack)
	fmt.Println()
	var isEmpty = stack.IsEmpty()
	fmt.Println("The stack is now empty?", isEmpty)
	fmt.Println()
}

func TestQueueExampleCode(t *tes.T) {
	fmt.Println("QUEUE EXAMPLE:")

	// Create a wait group for synchronization.
	var wg = new(syn.WaitGroup)
	defer wg.Wait()

	// Create a new queue with a specific capacity.
	var queue = fra.QueueWithCapacity[int](12)
	fmt.Println("The empty queue:", queue)
	fmt.Println()

	// Add some values to the queue.
	for i := 1; i < 10; i++ {
		queue.AddValue(i)
		fmt.Println("Added value:", i)
	}
	fmt.Println()
	fmt.Println("The partially filled queue:", queue)
	fmt.Println()

	// Remove values from the queue in the background.
	wg.Add(1)
	go func() {
		defer wg.Done()
		var value int
		var ok = true
		for i := 1; ok; i++ {
			value, ok = queue.RemoveFirst()
			if ok {
				fmt.Println("Removed value:", value)
			}
		}
		fmt.Println("The closed queue:", queue)
		fmt.Println()
	}()

	// Add some more values to the queue.
	for i := 10; i < 31; i++ {
		queue.AddValue(i)
		fmt.Println("Added value:", i)
	}
	queue.CloseChannel()
}

func TestCatalogExampleCode(t *tes.T) {
	fmt.Println("CATALOG EXAMPLE:")

	// Create a new catalog from a map.
	var catalog = fra.CatalogFromMap[string, int64](
		map[string]int64{
			"foo": 1,
			"bar": 2,
			"baz": 3,
		},
	)
	fmt.Println("The initialized catalog:", catalog)
	fmt.Println()

	// Add a new association to the catalog.
	catalog.SetValue("fuz", 4)
	fmt.Println("The updated catalog:", catalog)
	fmt.Println()

	// Sort the associations in the catalog by key.
	catalog.SortValues()
	fmt.Println("The sorted catalog:", catalog)
	fmt.Println()

	// List the keys for the catalog.
	var keys = catalog.GetKeys()
	fmt.Println("The keys for the catalog:", keys)
	fmt.Println()

	catalog.ReverseValues()
	fmt.Println("The reversed catalog:", catalog)
	fmt.Println()

	// Retrieve a value from the catalog.
	var value = catalog.GetValue("bar")
	fmt.Println("The value for the \"bar\" key is", value)
	fmt.Println()

	// Remove a value from the catalog.
	catalog.RemoveValue("foo")
	fmt.Println("The smaller catalog:", catalog)
	fmt.Println()

	// Change an existing value in the catalog.
	catalog.SetValue("baz", 5)
	fmt.Println("The updated catalog:", catalog)
	fmt.Println()
}

func TestCollatorExampleCode(t *tes.T) {
	fmt.Println("COLLATOR EXAMPLE:")

	// Create a collator with the default maximum depth.
	var collator = fra.Collator[any]()

	// Collate two strings.
	var s1 = "first"
	var s2 = "second"
	fmt.Println(
		"The first and second strings are equal:",
		collator.CompareValues(s1, s2),
	)
	fmt.Println(
		"The first string is ranked before the second strings:",
		collator.RankValues(s1, s2) == fra.LesserRank,
	)
	fmt.Println()

	// Collate two arrays.
	var a1 = []int{1, 2, 3}
	var a2 = []int{1, 3, 2}
	fmt.Println(
		"The first and second arrays are equal:",
		collator.CompareValues(a1, a2),
	)
	fmt.Println(
		"The first array is ranked before the second array:",
		collator.RankValues(a1, a2) == fra.LesserRank,
	)
	fmt.Println()
}

func TestSorterExampleCode(t *tes.T) {
	fmt.Println("SORTER EXAMPLE:")

	// Create a sorter with the default (natural) ranker.
	var sorter = fra.Sorter[string]()

	// Create an array.
	var array = []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
	}
	fmt.Println("The initial ordering of the values:", array)
	fmt.Println()

	// Sort the values in alphabetical order.
	sorter.SortValues(array)
	fmt.Println("The values in alphabetical order:", array)
	fmt.Println()

	// Sort the values in reverse order.
	sorter.ReverseValues(array)
	fmt.Println("The values in reverse order:", array)
	fmt.Println()

	// Shuffle the order of the values.
	sorter.ShuffleValues(array)
	fmt.Println("The values in random order:", array)
	fmt.Println()

	// Sort the values with a custom ranking function.
	sorter = fra.SorterWithRanker[string](
		func(first, second string) fra.Rank {
			switch {
			case first < second:
				return fra.GreaterRank
			case first > second:
				return fra.LesserRank
			default:
				return fra.EqualRank
			}
		},
	)
	sorter.SortValues(array)
	fmt.Println("The values in custom order:", array)
	fmt.Println()
}

func TestIteratorExampleCode(t *tes.T) {
	fmt.Println("ITERATOR EXAMPLE:")

	// Create a list from an array.
	var list = fra.ListFromArray[string](
		[]string{"foo", "bar", "baz"},
	)
	var iterator = list.GetIterator()

	// Iterate over the values in order.
	fmt.Println("The list values in order:")
	for iterator.HasNext() {
		var value = iterator.GetNext()
		fmt.Println("    value:", value)
	}
	fmt.Println()

	// Go to a specific value in the list.
	iterator.SetSlot(2)
	var value = iterator.GetPrevious()
	fmt.Println("The second value in the list is:", value)
	fmt.Println()

	// Iterate over the values in reverse order.
	fmt.Println("The list values in reverse order:")
	iterator.ToEnd()
	for iterator.HasPrevious() {
		var value = iterator.GetPrevious()
		fmt.Println("    value:", value)
	}
	fmt.Println()
}

// AGENT

// Tilde Types
type Boolean bool
type Byte byte
type Integer int
type String string

type Foolish interface {
	GetFoo() int
	GetBar() string
	GetNil() Foolish
}

func FooBar(foo int, bar string, baz Foolish) Foolish {
	return &foobar{foo, bar, baz}
}

// Encapsulated Type
type foobar struct {
	foo int
	bar string
	Baz Foolish
}

func (v *foobar) GetFoo() int     { return v.foo }
func (v foobar) GetFoo2() int     { return v.foo }
func (v *foobar) GetBar() string  { return v.bar }
func (v foobar) GetBar2() string  { return v.bar }
func (v *foobar) GetNil() Foolish { return nil }
func (v foobar) GetNil2() Foolish { return nil }

// Pure Structure
type Fuz struct {
	Bar string
}

func TestRank(t *tes.T) {
	ass.Equal(t, "LesserRank", fra.LesserRank.String())
	ass.Equal(t, "EqualRank", fra.EqualRank.String())
	ass.Equal(t, "GreaterRank", fra.GreaterRank.String())
}

func TestCompareMaximum(t *tes.T) {
	var collator = fra.CollatorClass[any]().CollatorWithMaximumDepth(1)
	var list = fra.ListClass[any]().ListFromArray([]any{"foo", []int{1, 2, 3}})
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The maximum traversal depth was exceeded: 1", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	collator.CompareValues(list, list)
}

func TestRankMaximum(t *tes.T) {
	var collator = fra.CollatorClass[any]().CollatorWithMaximumDepth(1)
	var list = fra.ListClass[any]().ListFromArray([]any{"foo", []int{1, 2, 3}})
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The maximum traversal depth was exceeded: 1", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	collator.RankValues(list, list)
}

func TestComparison(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()

	// Nil
	var ShouldBeNil any

	ass.True(t, collator.CompareValues(nil, nil))
	ass.True(t, collator.CompareValues(nil, ShouldBeNil))
	ass.True(t, collator.CompareValues(ShouldBeNil, ShouldBeNil))
	ass.True(t, collator.CompareValues(ShouldBeNil, nil))

	// Boolean
	var False = false
	var True = true
	var ShouldBeFalse bool

	ass.True(t, collator.CompareValues(ShouldBeFalse, False))
	ass.False(t, collator.CompareValues(True, ShouldBeFalse))

	ass.False(t, collator.CompareValues(False, True))
	ass.True(t, collator.CompareValues(False, False))
	ass.False(t, collator.CompareValues(True, False))
	ass.True(t, collator.CompareValues(True, True))

	// Byte
	var Zero byte = 0x00
	var One byte = 0x01
	var ShouldBeZero byte

	ass.True(t, collator.CompareValues(ShouldBeZero, Zero))
	ass.False(t, collator.CompareValues(One, ShouldBeZero))

	ass.False(t, collator.CompareValues(Zero, One))
	ass.True(t, collator.CompareValues(Zero, Zero))
	ass.False(t, collator.CompareValues(One, Zero))
	ass.True(t, collator.CompareValues(One, One))

	// Integer
	var Zilch = 0
	var Two = 2
	var Three = 3
	var ShouldBeZilch int

	ass.True(t, collator.CompareValues(ShouldBeZilch, Zilch))
	ass.False(t, collator.CompareValues(Two, ShouldBeZilch))

	ass.False(t, collator.CompareValues(Two, Three))
	ass.True(t, collator.CompareValues(Two, Two))
	ass.False(t, collator.CompareValues(Three, Two))
	ass.True(t, collator.CompareValues(Three, Three))

	// Float
	var Negligible = 0.0
	var Fourth = 0.25
	var Half = 0.5
	var ShouldBeNegligible float64

	ass.True(t, collator.CompareValues(ShouldBeNegligible, Negligible))
	ass.False(t, collator.CompareValues(Half, ShouldBeNegligible))

	ass.False(t, collator.CompareValues(Fourth, Half))
	ass.True(t, collator.CompareValues(Fourth, Fourth))
	ass.False(t, collator.CompareValues(Half, Fourth))
	ass.True(t, collator.CompareValues(Half, Half))

	// Complex
	var Origin = 0 + 0i
	var PiOver4 = 1 + 1i
	var PiOver2 = 1 + 0i
	var ShouldBeOrigin complex128

	ass.True(t, collator.CompareValues(ShouldBeOrigin, Origin))
	ass.False(t, collator.CompareValues(PiOver4, ShouldBeOrigin))

	ass.False(t, collator.CompareValues(PiOver4, PiOver2))
	ass.True(t, collator.CompareValues(PiOver4, PiOver4))
	ass.False(t, collator.CompareValues(PiOver2, PiOver4))
	ass.True(t, collator.CompareValues(PiOver2, PiOver2))

	// Rune
	var Null = rune(0)
	var Sad = '☹'
	var Happy = '☺'
	var ShouldBeNull rune

	ass.True(t, collator.CompareValues(ShouldBeNull, Null))
	ass.False(t, collator.CompareValues(Sad, ShouldBeNull))

	ass.False(t, collator.CompareValues(Happy, Sad))
	ass.True(t, collator.CompareValues(Happy, Happy))
	ass.False(t, collator.CompareValues(Sad, Happy))
	ass.True(t, collator.CompareValues(Sad, Sad))

	// String
	var Empty = ""
	var Hello = "Hello"
	var World = "World"
	var ShouldBeEmpty string

	ass.True(t, collator.CompareValues(ShouldBeEmpty, Empty))
	ass.False(t, collator.CompareValues(Hello, ShouldBeEmpty))

	ass.False(t, collator.CompareValues(World, Hello))
	ass.True(t, collator.CompareValues(World, World))
	ass.False(t, collator.CompareValues(Hello, World))
	ass.True(t, collator.CompareValues(Hello, Hello))

	// Array
	var Universe = "Universe"
	var a0 = []any{}
	var a1 = []any{Hello, World}
	var a2 = []any{Hello, Universe}
	var aNil []any

	ass.True(t, collator.CompareValues(aNil, aNil))
	ass.False(t, collator.CompareValues(aNil, a0))
	ass.False(t, collator.CompareValues(a0, aNil))
	ass.True(t, collator.CompareValues(a0, a0))

	ass.False(t, collator.CompareValues(a1, a2))
	ass.True(t, collator.CompareValues(a1, a1))
	ass.False(t, collator.CompareValues(a2, a1))
	ass.True(t, collator.CompareValues(a2, a2))

	// Map
	var m0 = map[any]any{}
	var m1 = map[any]any{
		One: True,
		Two: World}
	var m2 = map[any]any{
		One: True,
		Two: Hello}
	var m3 = map[any]any{
		One: nil,
		Two: Hello}
	var mNil map[any]any

	ass.True(t, collator.CompareValues(mNil, mNil))
	ass.False(t, collator.CompareValues(mNil, m0))
	ass.False(t, collator.CompareValues(m0, mNil))
	ass.True(t, collator.CompareValues(m0, m0))

	ass.False(t, collator.CompareValues(m1, m2))
	ass.True(t, collator.CompareValues(m1, m1))
	ass.False(t, collator.CompareValues(m2, m1))
	ass.True(t, collator.CompareValues(m2, m2))
	ass.False(t, collator.CompareValues(m2, m3))

	// Struct
	var f0 Foolish
	var f1 = FooBar(1, "one", nil)
	var f2 = FooBar(1, "one", nil)
	var f3 = FooBar(2, "two", nil)
	var f4 = Fuz{"two"}
	var f5 = Fuz{"two"}
	var f6 = Fuz{"three"}
	ass.True(t, collator.CompareValues(f0, f0))
	ass.False(t, collator.CompareValues(f0, f1))
	ass.True(t, collator.CompareValues(f1, f1))
	ass.True(t, collator.CompareValues(f1, f2))
	ass.False(t, collator.CompareValues(f2, f3))
	ass.True(t, collator.CompareValues(f4, f4))
	ass.True(t, collator.CompareValues(f4, f5))
	ass.False(t, collator.CompareValues(f5, f6))
	ass.True(t, collator.CompareValues(&f4, &f4))
	ass.True(t, collator.CompareValues(&f4, &f5))
	ass.False(t, collator.CompareValues(&f5, &f6))
}

func TestTildeTypes(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()

	// Boolean
	var False = Boolean(false)
	var True = Boolean(true)
	var ShouldBeFalse Boolean

	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeFalse, ShouldBeFalse))
	ass.Equal(t, fra.LesserRank, collator.RankValues(ShouldBeFalse, True))
	ass.Equal(t, fra.EqualRank, collator.RankValues(False, ShouldBeFalse))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(True, ShouldBeFalse))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeFalse, False))
	ass.Equal(t, fra.LesserRank, collator.RankValues(False, True))
	ass.Equal(t, fra.EqualRank, collator.RankValues(False, False))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(True, False))
	ass.Equal(t, fra.EqualRank, collator.RankValues(True, True))

	// Byte
	var Zero = Byte(0)
	var One = Byte(1)
	var TFF = Byte(255)
	var ShouldBeZero Byte

	ass.True(t, collator.CompareValues(ShouldBeZero, Zero))
	ass.False(t, collator.CompareValues(One, ShouldBeZero))

	ass.False(t, collator.CompareValues(One, TFF))
	ass.True(t, collator.CompareValues(One, One))
	ass.False(t, collator.CompareValues(TFF, One))
	ass.True(t, collator.CompareValues(TFF, TFF))

	// Integer
	var Zilch = Integer(0)
	var Two = Integer(2)
	var Three = Integer(3)
	var ShouldBeZilch Integer

	ass.True(t, collator.CompareValues(ShouldBeZilch, Zilch))
	ass.False(t, collator.CompareValues(Two, ShouldBeZilch))

	ass.False(t, collator.CompareValues(Two, Three))
	ass.True(t, collator.CompareValues(Two, Two))
	ass.False(t, collator.CompareValues(Three, Two))
	ass.True(t, collator.CompareValues(Three, Three))

	// String
	var Empty = String("")
	var Hello = String("Hello")
	var World = String("World")
	var ShouldBeEmpty String

	ass.True(t, collator.CompareValues(ShouldBeEmpty, Empty))
	ass.False(t, collator.CompareValues(Hello, ShouldBeEmpty))

	ass.False(t, collator.CompareValues(World, Hello))
	ass.True(t, collator.CompareValues(World, World))
	ass.False(t, collator.CompareValues(Hello, World))
	ass.True(t, collator.CompareValues(Hello, Hello))
}

func TestCompareRecursiveArrays(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var list = fra.ListClass[any]().ListFromArray(
		[]any{0},
	)
	list.SetValue(1, list) // Now it is recursive.
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The maximum traversal depth was exceeded: 16", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	collator.CompareValues(list, list) // This should panic.
}

func TestCompareRecursiveMaps(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var catalog = fra.CatalogClass[string, any]().CatalogFromMap(
		map[string]any{
			"first": 1,
		},
	)
	catalog.SetValue("first", catalog) // Now it is recursive.
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The maximum traversal depth was exceeded: 16", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	collator.CompareValues(catalog, catalog) // This should panic.
}

func TestNilRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var ShouldBeNil any
	ass.Equal(t, fra.EqualRank, collator.RankValues(nil, nil))
	ass.Equal(t, fra.EqualRank, collator.RankValues(nil, ShouldBeNil))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeNil, ShouldBeNil))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeNil, nil))
}

func TestBooleanRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var False = false
	var True = true
	var ShouldBeFalse bool
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeFalse, ShouldBeFalse))
	ass.Equal(t, fra.LesserRank, collator.RankValues(ShouldBeFalse, True))
	ass.Equal(t, fra.EqualRank, collator.RankValues(False, ShouldBeFalse))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(True, ShouldBeFalse))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeFalse, False))
	ass.Equal(t, fra.LesserRank, collator.RankValues(False, True))
	ass.Equal(t, fra.EqualRank, collator.RankValues(False, False))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(True, False))
	ass.Equal(t, fra.EqualRank, collator.RankValues(True, True))
}

func TestByteRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var Zero byte = 0x00
	var One byte = 0x01
	var ShouldBeZero byte
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeZero, ShouldBeZero))
	ass.Equal(t, fra.LesserRank, collator.RankValues(ShouldBeZero, One))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Zero, ShouldBeZero))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(One, ShouldBeZero))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeZero, Zero))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Zero, One))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Zero, Zero))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(One, Zero))
	ass.Equal(t, fra.EqualRank, collator.RankValues(One, One))
}

func TestIntegerRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var Zilch = 0
	var Two = 2
	var Three = 3
	var ShouldBeZilch int
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeZilch, ShouldBeZilch))
	ass.Equal(t, fra.LesserRank, collator.RankValues(ShouldBeZilch, Two))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Zilch, ShouldBeZilch))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Two, ShouldBeZilch))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeZilch, Zilch))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Two, Three))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Two, Two))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Three, Two))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Three, Three))
}

func TestFloatRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var Negligible = 0.0
	var Fourth = 0.25
	var Half = 0.5
	var ShouldBeNegligible float64
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeNegligible, ShouldBeNegligible))
	ass.Equal(t, fra.LesserRank, collator.RankValues(ShouldBeNegligible, Half))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Negligible, ShouldBeNegligible))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Half, ShouldBeNegligible))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeNegligible, Negligible))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Fourth, Half))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Fourth, Fourth))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Half, Fourth))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Half, Half))
}

func TestComplexRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var Origin complex128
	var One = 1 + 0i
	var Pi = -1 + 0i
	var PiOver2 = 0 + 1i
	var PiOver4 = 1 + 1i

	ass.Equal(t, fra.EqualRank, collator.RankValues(Origin, Origin))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Origin, One))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Origin, Pi))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Origin, PiOver2))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Origin, PiOver4))

	ass.Equal(t, fra.GreaterRank, collator.RankValues(One, Origin))
	ass.Equal(t, fra.EqualRank, collator.RankValues(One, One))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(One, Pi))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(One, PiOver2))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(One, PiOver4))

	ass.Equal(t, fra.GreaterRank, collator.RankValues(Pi, Origin))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Pi, One))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Pi, Pi))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Pi, PiOver2))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Pi, PiOver4))

	ass.Equal(t, fra.GreaterRank, collator.RankValues(PiOver2, Origin))
	ass.Equal(t, fra.LesserRank, collator.RankValues(PiOver2, One))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(PiOver2, Pi))
	ass.Equal(t, fra.EqualRank, collator.RankValues(PiOver2, PiOver2))
	ass.Equal(t, fra.LesserRank, collator.RankValues(PiOver2, PiOver4))

	ass.Equal(t, fra.GreaterRank, collator.RankValues(PiOver4, Origin))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(PiOver4, One))
	ass.Equal(t, fra.LesserRank, collator.RankValues(PiOver4, Pi))
	ass.Equal(t, fra.LesserRank, collator.RankValues(PiOver4, PiOver2))
	ass.Equal(t, fra.EqualRank, collator.RankValues(PiOver4, PiOver4))
}

func TestRuneRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var Null = rune(0)
	var Sad = '☹'
	var Happy = '☺'
	var ShouldBeNull rune
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeNull, ShouldBeNull))
	ass.Equal(t, fra.LesserRank, collator.RankValues(ShouldBeNull, Sad))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Null, ShouldBeNull))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Sad, ShouldBeNull))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeNull, Null))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Sad, Happy))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Sad, Sad))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Happy, Sad))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Happy, Happy))
}

func TestStringRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var Empty = ""
	var Hello = "Hello"
	var World = "World"
	var ShouldBeEmpty string
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeEmpty, ShouldBeEmpty))
	ass.Equal(t, fra.LesserRank, collator.RankValues(ShouldBeEmpty, Hello))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Empty, ShouldBeEmpty))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(Hello, ShouldBeEmpty))
	ass.Equal(t, fra.EqualRank, collator.RankValues(ShouldBeEmpty, Empty))
	ass.Equal(t, fra.LesserRank, collator.RankValues(Hello, World))
	ass.Equal(t, fra.EqualRank, collator.RankValues(Hello, Hello))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(World, Hello))
	ass.Equal(t, fra.EqualRank, collator.RankValues(World, World))
}

func TestArrayRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var Hello = "Hello"
	var World = "World"
	var Universe = "Universe"
	var a0 = []any{}
	var a1 = []any{Hello, World}
	var a2 = []any{Hello, Universe}
	var a3 = []any{Hello, World, Universe}
	var a4 = []any{Hello, Universe, World}
	var aNil []any
	ass.Equal(t, fra.EqualRank, collator.RankValues(aNil, aNil))
	ass.Equal(t, fra.LesserRank, collator.RankValues(aNil, a0))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(a0, aNil))
	ass.Equal(t, fra.EqualRank, collator.RankValues(a0, a0))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(a1, aNil))
	ass.Equal(t, fra.LesserRank, collator.RankValues(a2, a1))
	ass.Equal(t, fra.EqualRank, collator.RankValues(a2, a2))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(a1, a2))
	ass.Equal(t, fra.EqualRank, collator.RankValues(a1, a1))
	ass.Equal(t, fra.LesserRank, collator.RankValues(a2, a3))
	ass.Equal(t, fra.EqualRank, collator.RankValues(a2, a2))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(a3, a2))
	ass.Equal(t, fra.EqualRank, collator.RankValues(a3, a3))
	ass.Equal(t, fra.LesserRank, collator.RankValues(a4, a1))
	ass.Equal(t, fra.EqualRank, collator.RankValues(a4, a4))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(a1, a4))
	ass.Equal(t, fra.EqualRank, collator.RankValues(a1, a1))
}

func TestMapRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var True = true
	var One byte = 0x01
	var Two = 2
	var Three = "three"
	var Hello = "Hello"
	var World = "World"
	var Universe = "Universe"
	var m0 = map[any]any{}
	var m1 = map[any]any{
		One: True,
		Two: World}
	var m2 = map[any]any{
		One: True,
		Two: Hello}
	var m3 = map[any]any{
		One:   True,
		Two:   World,
		Three: Universe}
	var m4 = map[any]any{
		One:   True,
		Two:   Universe,
		Three: World}
	var mNil map[any]any
	ass.Equal(t, fra.EqualRank, collator.RankValues(mNil, mNil))
	ass.Equal(t, fra.LesserRank, collator.RankValues(mNil, m0))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(m0, mNil))
	ass.Equal(t, fra.EqualRank, collator.RankValues(m0, m0))
	ass.Equal(t, fra.LesserRank, collator.RankValues(m2, m1))
	ass.Equal(t, fra.EqualRank, collator.RankValues(m2, m2))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(m1, m2))
	ass.Equal(t, fra.EqualRank, collator.RankValues(m1, m1))
	ass.Equal(t, fra.LesserRank, collator.RankValues(m2, m3))
	ass.Equal(t, fra.EqualRank, collator.RankValues(m2, m2))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(m3, m2))
	ass.Equal(t, fra.EqualRank, collator.RankValues(m3, m3))
	ass.Equal(t, fra.LesserRank, collator.RankValues(m4, m1))
	ass.Equal(t, fra.EqualRank, collator.RankValues(m4, m4))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(m1, m4))
	ass.Equal(t, fra.EqualRank, collator.RankValues(m1, m1))
}

func TestStructRanking(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var f1 = FooBar(1, "one", nil)
	var f2 = FooBar(1, "two", nil)
	var f3 = FooBar(2, "two", nil)
	var f4 = Fuz{"two"}
	var f5 = Fuz{"two"}
	var f6 = Fuz{"three"}
	ass.Equal(t, fra.EqualRank, collator.RankValues(f1, f1))
	ass.Equal(t, fra.LesserRank, collator.RankValues(f1, f2))
	ass.Equal(t, fra.LesserRank, collator.RankValues(f2, f3))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(f3, f1))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(f3, f2))
	ass.Equal(t, fra.EqualRank, collator.RankValues(f4, f4))
	ass.Equal(t, fra.EqualRank, collator.RankValues(f4, f5))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(f5, f6))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(f3, &f4))
	ass.Equal(t, fra.EqualRank, collator.RankValues(&f4, &f4))
	ass.Equal(t, fra.EqualRank, collator.RankValues(&f4, &f5))
	ass.Equal(t, fra.GreaterRank, collator.RankValues(&f5, &f6))
}

func TestTildeArrays(t *tes.T) {
	var collator = fra.CollatorClass[String]().Collator()
	var ranker = collator.RankValues
	var sorter = fra.SorterClass[String]().SorterWithRanker(ranker)
	var alpha = String("alpha")
	var beta = String("beta")
	var gamma = String("gamma")
	var delta = String("delta")
	var array = []String{alpha, beta, gamma, delta}
	sorter.SortValues(array)
	ass.Equal(t, alpha, array[0])
	ass.Equal(t, beta, array[1])
	ass.Equal(t, delta, array[2])
	ass.Equal(t, gamma, array[3])
}

func TestRankRecursiveArrays(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var list = fra.ListClass[any]().ListFromArray(
		[]any{0},
	)
	list.SetValue(1, list) // Now it is recursive.
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The maximum traversal depth was exceeded: 16", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	collator.RankValues(list, list) // This should panic.
}

func TestRankRecursiveMaps(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var catalog = fra.CatalogClass[string, any]().CatalogFromMap(
		map[string]any{
			"first": 1,
		},
	)
	catalog.SetValue("first", catalog) // Now it is recursive.
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The maximum traversal depth was exceeded: 16", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	collator.RankValues(catalog, catalog) // This should panic.
}

func TestIteratorsWithLists(t *tes.T) {
	var list = fra.ListClass[int]().ListFromArray([]int{1, 2, 3, 4, 5})
	list = fra.ListClass[int]().ListFromSequence(list)
	var iterator = list.GetIterator()
	ass.False(t, iterator.IsEmpty())
	ass.True(t, iterator.GetSize() == 5)
	ass.True(t, iterator.GetSlot() == 0)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, 1, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, 1, iterator.GetPrevious())
	iterator.SetSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, 3, iterator.GetNext())
	ass.False(t, iterator.IsEmpty())
	ass.True(t, iterator.GetSize() == 5)
	ass.True(t, iterator.GetSlot() == 3)
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, 5, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, 1, iterator.GetNext())
}

func TestSortingEmpty(t *tes.T) {
	var collator = fra.CollatorClass[any]().Collator()
	var ranker = collator.RankValues
	var sorter = fra.SorterClass[any]().SorterWithRanker(ranker)
	var empty = []any{}
	sorter.SortValues(empty)
}

func TestSortingIntegers(t *tes.T) {
	var collator = fra.CollatorClass[int]().Collator()
	var ranker = collator.RankValues
	var sorter = fra.SorterClass[int]().SorterWithRanker(ranker)
	var unsorted = []int{4, 3, 1, 5, 2}
	var sorted = []int{1, 2, 3, 4, 5}
	sorter.SortValues(unsorted)
	ass.Equal(t, sorted, unsorted)
}

func TestSortingStrings(t *tes.T) {
	var collator = fra.CollatorClass[string]().Collator()
	var ranker = collator.RankValues
	var sorter = fra.SorterClass[string]().SorterWithRanker(ranker)
	var unsorted = []string{"alpha", "beta", "gamma", "delta"}
	var sorted = []string{"alpha", "beta", "delta", "gamma"}
	sorter.SortValues(unsorted)
	ass.Equal(t, sorted, unsorted)
}

var encoder = fra.Encoder()

func TestBase16EmptyRoundTrip(t *tes.T) {
	var bytes = make([]byte, 0)

	// Encode as base 16.
	var base16 = encoder.Base16Encode(bytes)

	// Decode base 16 to bytes.
	var decoded = encoder.Base16Decode(base16)
	ass.Equal(t, bytes, decoded)

	// Encode as base 16 again.
	var encoded = encoder.Base16Encode(decoded)
	ass.Equal(t, base16, encoded)

	// Decode base 16 again.
	var again = encoder.Base16Decode(encoded)
	ass.Equal(t, again, decoded)
}

func TestBase16RoundTrip(t *tes.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index := range bytes {
		bytes[index] = byte(index)
	}

	for index := 0; index < len(bytes); index++ {
		// Encode as base 16.
		var base16 = encoder.Base16Encode(bytes[:index+1])

		// Decode base 16 to bytes.
		var decoded = encoder.Base16Decode(base16)
		ass.Equal(t, bytes[:index+1], decoded)

		// Encode as base 16 again.
		var encoded = encoder.Base16Encode(decoded)
		ass.Equal(t, base16, encoded)

		// Decode base 16 again.
		var again = encoder.Base16Decode(encoded)
		ass.Equal(t, again, decoded)
	}
}

func TestBase32EmptyRoundTrip(t *tes.T) {
	var bytes = make([]byte, 0)

	// Encode as base 32.
	var base32 = encoder.Base32Encode(bytes)

	// Decode base 32 to bytes.
	var decoded = encoder.Base32Decode(base32)
	ass.Equal(t, bytes, decoded)

	// Encode as base 32 again.
	var encoded = encoder.Base32Encode(decoded)
	ass.Equal(t, base32, encoded)

	// Decode base 32 again.
	var again = encoder.Base32Decode(encoded)
	ass.Equal(t, again, decoded)
}

func TestBase32RoundTrip(t *tes.T) {
	// Seed the bytes.
	var bytes = make([]byte, 256)
	for index := range bytes {
		bytes[index] = byte(index)
	}

	for index := 0; index < len(bytes); index++ {
		// Encode as base 32.
		var base32 = encoder.Base32Encode(bytes[:index+1])

		// Decode base 32 to bytes.
		var decoded = encoder.Base32Decode(base32)
		ass.Equal(t, bytes[:index+1], decoded)

		// Encode as base 32 again.
		var encoded = encoder.Base32Encode(decoded)
		ass.Equal(t, base32, encoded)

		// Decode base 32 again.
		var again = encoder.Base32Decode(encoded)
		ass.Equal(t, again, decoded)
	}
}

func TestBase64EmptyRoundTrip(t *tes.T) {
	var bytes = make([]byte, 0)

	// Encode as base 64.
	var base64 = encoder.Base64Encode(bytes)

	// Decode base 64 to bytes.
	var decoded = encoder.Base64Decode(base64)
	ass.Equal(t, bytes, decoded)

	// Encode as base 64 again.
	var encoded = encoder.Base64Encode(decoded)
	ass.Equal(t, base64, encoded)

	// Decode base 64 again.
	var again = encoder.Base64Decode(encoded)
	ass.Equal(t, again, decoded)
}

func TestBase64RoundTrip(t *tes.T) {
	// Seed the bytes.
	var bytes = make([]byte, 240)
	for index := range bytes {
		bytes[index] = byte(index)
	}

	for index := 0; index < len(bytes); index++ {
		// Encode as base 64.
		var base64 = encoder.Base64Encode(bytes[:index+1])

		// Decode base 64 to bytes.
		var decoded = encoder.Base64Decode(base64)
		ass.Equal(t, bytes[:index+1], decoded)

		// Encode as base 64 again.
		var encoded = encoder.Base64Encode(decoded)
		ass.Equal(t, base64, encoded)

		// Decode base 64 again.
		var again = encoder.Base64Decode(encoded)
		ass.Equal(t, again, decoded)
	}
}

var generator = fra.Generator()

func TestRandomBooleans(t *tes.T) {
	var foundFalse uint
	var foundTrue uint
	for i := 0; i < 100; i++ {
		if generator.RandomBoolean() {
			foundTrue++
		} else {
			foundFalse++
		}
	}
	ass.True(t, foundFalse > 25)
	ass.True(t, foundTrue > 25)
}

func TestRandomOrdinals(t *tes.T) {
	var foundZero bool
	var foundFive bool
	for i := 0; i < 100; i++ {
		var random = generator.RandomOrdinal(5)
		if random == 0 {
			foundZero = true
		}
		if random == 5 {
			foundFive = true
		}
	}
	ass.False(t, foundZero)
	ass.True(t, foundFive)
}

func TestRandomProbabilities(t *tes.T) {
	var total float64
	for i := 0; i < 10000; i++ {
		total += generator.RandomProbability()
	}
	ass.True(t, total > 4800)
	ass.True(t, total < 5200)
}

var (
	invalid fra.State = fra.ControllerClass().Invalid()
	state1  fra.State = "$State1"
	state2  fra.State = "$State2"
	state3  fra.State = "$State3"
)

var (
	initialized fra.Event = "$Initialized"
	processed   fra.Event = "$Processed"
	finalized   fra.Event = "$Finalized"
)

func TestController(t *tes.T) {
	var events = []fra.Event{initialized, processed, finalized}
	var transitions = map[fra.State]fra.Transitions{
		state1: fra.Transitions{state2, invalid, invalid},
		state2: fra.Transitions{invalid, state2, state3},
		state3: fra.Transitions{invalid, invalid, invalid},
	}

	var controller = fra.Controller(events, transitions, state1)
	ass.Equal(t, state1, controller.GetState())
	ass.Equal(t, state2, controller.ProcessEvent(initialized))
	ass.Equal(t, state2, controller.ProcessEvent(processed))
	ass.Equal(t, state3, controller.ProcessEvent(finalized))
	controller.SetState(state1)
	ass.Equal(t, state1, controller.GetState())
}

// COLLECTION

func TestCatalogConstructors(t *tes.T) {
	var class = fra.CatalogClass[rune, int64]()
	class.Catalog()
	class.CatalogFromArray([]fra.AssociationLike[rune, int64]{})
	class.CatalogFromMap(map[rune]int64{})
	var sequence = class.CatalogFromMap(map[rune]int64{
		'a': 1,
		'b': 2,
		'c': 3,
	})
	var catalog = class.CatalogFromSequence(sequence)
	ass.Equal(t, sequence.AsArray(), catalog.AsArray())
}

func TestCatalogsWithPrimitivesAndStrings(t *tes.T) {
	var catalog = fra.Catalog[any, string]()

	var binaryString = `'>
    0123456789abcdefghijk
<'`
	var binary = fra.BinaryFromString(binaryString)
	catalog.SetValue(binary, binaryString)
	ass.Equal(t, catalog.GetValue(binary), binary.AsString())

	var nameString = "/foo/5bar"
	var name = fra.NameFromString(nameString)
	catalog.SetValue(name, nameString)
	ass.Equal(t, catalog.GetValue(name), name.AsString())

	var narrativeString = `">
    This is a narrative.
<"`
	var narrative = fra.NarrativeFromString(narrativeString)
	catalog.SetValue(narrative, narrativeString)
	ass.Equal(t, catalog.GetValue(narrative), narrative.AsString())

	var patternString = `"b[aeiou]g"?`
	var pattern = fra.PatternFromString(patternString)
	catalog.SetValue(pattern, patternString)
	ass.Equal(t, catalog.GetValue(pattern), pattern.AsString())

	var quoteString = `"To be or not to be..."`
	var quote = fra.QuoteFromString(quoteString)
	catalog.SetValue(quote, quoteString)
	ass.Equal(t, catalog.GetValue(quote), quote.AsString())

	var tag = fra.TagWithSize(16)
	var tagString = tag.AsString()
	catalog.SetValue(tag, tagString)
	ass.Equal(t, catalog.GetValue(tag), tag.AsString())

	var versionString = "v1.2.3"
	var version = fra.VersionFromString(versionString)
	catalog.SetValue(version, versionString)
	ass.Equal(t, catalog.GetValue(version), version.AsString())
}

func TestCatalogsWithStringsAndIntegers(t *tes.T) {
	var catalogCollator = fra.Collator[fra.CatalogLike[string, int]]()
	var keys = fra.ListClass[string]().ListFromArray([]string{"foo", "bar"})
	var association1 = fra.Association("foo", 1)
	var association2 = fra.Association("bar", 2)
	var association3 = fra.Association("baz", 3)
	var catalog = fra.Catalog[string, int]()
	ass.True(t, catalog.IsEmpty())
	ass.True(t, catalog.GetSize() == 0)
	ass.Equal(t, []string{}, catalog.GetKeys().AsArray())
	ass.Equal(t, []fra.AssociationLike[string, int]{}, catalog.AsArray())
	var iterator = catalog.GetIterator()
	ass.False(t, iterator.HasNext())
	ass.False(t, iterator.HasPrevious())
	iterator.ToStart()
	iterator.ToEnd()
	catalog.SortValues()
	catalog.ShuffleValues()
	catalog.RemoveAll()
	catalog.SetValue(association1.GetKey(), association1.GetValue())
	ass.False(t, catalog.IsEmpty())
	ass.True(t, catalog.GetSize() == 1)
	catalog.SetValue(association2.GetKey(), association2.GetValue())
	catalog.SetValue(association3.GetKey(), association3.GetValue())
	ass.True(t, catalog.GetSize() == 3)
	var catalog2 = fra.CatalogFromSequence(catalog)
	ass.True(t, catalogCollator.CompareValues(catalog, catalog2))
	var m = fra.CatalogClass[string, int]().CatalogFromMap(map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	})
	var associationCollator = fra.Collator[fra.AssociationLike[string, int]]()
	var catalog3 = fra.CatalogFromSequence(m)
	catalog2.SortValues()
	catalog3.SortValuesWithRanker(associationCollator.RankValues)
	ass.True(t, catalogCollator.CompareValues(catalog2, catalog3))
	iterator = catalog.GetIterator()
	ass.True(t, iterator.HasNext())
	ass.False(t, iterator.HasPrevious())
	ass.Equal(t, association1, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	iterator.ToEnd()
	ass.False(t, iterator.HasNext())
	ass.True(t, iterator.HasPrevious())
	ass.Equal(t, association3, iterator.GetPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, []string{"foo", "bar", "baz"}, catalog.GetKeys().AsArray())
	ass.Equal(t, 3, catalog.GetValue("baz"))
	catalog.SetValue("bar", 5)
	ass.Equal(t, []int{1, 5}, catalog.GetValues(keys).AsArray())
	catalog.SortValues()
	ass.Equal(t, []string{"bar", "baz", "foo"}, catalog.GetKeys().AsArray())
	catalog.ReverseValues()
	ass.Equal(t, []string{"foo", "baz", "bar"}, catalog.GetKeys().AsArray())
	catalog.ReverseValues()
	ass.Equal(t, []int{1, 5}, catalog.RemoveValues(keys).AsArray())
	ass.True(t, catalog.GetSize() == 1)
	ass.Equal(t, 3, catalog.RemoveValue("baz"))
	ass.True(t, catalog.IsEmpty())
	ass.True(t, catalog.GetSize() == 0)
	catalog.RemoveAll()
	ass.True(t, catalog.IsEmpty())
	ass.True(t, catalog.GetSize() == 0)
}

func TestCatalogsWithMerge(t *tes.T) {
	var collator = fra.Collator[fra.CatalogLike[string, int]]()
	var association1 = fra.Association("foo", 1)
	var association2 = fra.Association("bar", 2)
	var association3 = fra.Association("baz", 3)
	var catalog1 = fra.Catalog[string, int]()
	catalog1.SetValue(association1.GetKey(), association1.GetValue())
	catalog1.SetValue(association2.GetKey(), association2.GetValue())
	var catalog2 = fra.Catalog[string, int]()
	catalog2.SetValue(association2.GetKey(), association2.GetValue())
	catalog2.SetValue(association3.GetKey(), association3.GetValue())
	var catalog3 = fra.CatalogClass[string, int]().Merge(catalog1, catalog2)
	var catalog4 = fra.Catalog[string, int]()
	catalog4.SetValue(association1.GetKey(), association1.GetValue())
	catalog4.SetValue(association2.GetKey(), association2.GetValue())
	catalog4.SetValue(association3.GetKey(), association3.GetValue())
	ass.True(t, collator.CompareValues(catalog3, catalog4))
}

func TestCatalogsWithExtract(t *tes.T) {
	var keys = fra.ListClass[string]().ListFromArray([]string{"foo", "baz"})
	var association1 = fra.Association("foo", 1)
	var association2 = fra.Association("bar", 2)
	var association3 = fra.Association("baz", 3)
	var catalog1 = fra.Catalog[string, int]()
	catalog1.SetValue(association1.GetKey(), association1.GetValue())
	catalog1.SetValue(association2.GetKey(), association2.GetValue())
	catalog1.SetValue(association3.GetKey(), association3.GetValue())
	var catalog2 = fra.CatalogClass[string, int]().Extract(catalog1, keys)
	var catalog3 = fra.Catalog[string, int]()
	catalog3.SetValue(association1.GetKey(), association1.GetValue())
	catalog3.SetValue(association3.GetKey(), association3.GetValue())
	var collator = fra.Collator[fra.CatalogLike[string, int]]()
	ass.True(t, collator.CompareValues(catalog2, catalog3))
	var catalog4 = fra.CatalogFromArray([]fra.AssociationLike[string, int]{
		association1,
		association2,
		association3,
	})
	ass.True(t, collator.CompareValues(catalog1, catalog4))
}

func TestCatalogsWithEmptyCatalogs(t *tes.T) {
	var keys = fra.ListClass[int]().List()
	var catalog1 = fra.Catalog[int, string]()
	var catalog2 = fra.Catalog[int, string]()
	var catalog3 = fra.CatalogClass[int, string]().Merge(catalog1, catalog2)
	var catalog4 = fra.CatalogClass[int, string]().Extract(catalog1, keys)
	var collator = fra.Collator[fra.CatalogLike[int, string]]()
	ass.True(t, collator.CompareValues(catalog1, catalog2))
	ass.True(t, collator.CompareValues(catalog2, catalog3))
	ass.True(t, collator.CompareValues(catalog3, catalog4))
	ass.True(t, collator.CompareValues(catalog4, catalog1))
}

func TestIntervalIterators(t *tes.T) {
	var glyphs = fra.Interval[fra.GlyphLike](
		fra.Exclusive,
		fra.Glyph(65),
		fra.Glyph(70),
		fra.Exclusive,
	)
	var iterator = glyphs.GetIterator()
	ass.Equal(t, glyphs.GetSize(), iterator.GetSize())
	ass.Equal(t, 'B', iterator.GetNext().AsIntrinsic())
	iterator.ToEnd()
	ass.Equal(t, 'E', iterator.GetPrevious().AsIntrinsic())

	glyphs = fra.Interval[fra.GlyphLike](
		fra.Inclusive,
		fra.Glyph(65),
		fra.Glyph(70),
		fra.Exclusive,
	)
	iterator = glyphs.GetIterator()
	ass.Equal(t, glyphs.GetSize(), iterator.GetSize())
	ass.Equal(t, 'A', iterator.GetNext().AsIntrinsic())
	iterator.ToEnd()
	ass.Equal(t, 'E', iterator.GetPrevious().AsIntrinsic())

	glyphs = fra.Interval[fra.GlyphLike](
		fra.Exclusive,
		fra.Glyph(65),
		fra.Glyph(70),
		fra.Inclusive,
	)
	iterator = glyphs.GetIterator()
	ass.Equal(t, glyphs.GetSize(), iterator.GetSize())
	ass.Equal(t, 'B', iterator.GetNext().AsIntrinsic())
	iterator.ToEnd()
	ass.Equal(t, 'F', iterator.GetPrevious().AsIntrinsic())

	glyphs = fra.Interval[fra.GlyphLike](
		fra.Inclusive,
		fra.Glyph(65),
		fra.Glyph(70),
		fra.Inclusive,
	)
	iterator = glyphs.GetIterator()
	ass.Equal(t, glyphs.GetSize(), iterator.GetSize())
	ass.Equal(t, 'A', iterator.GetNext().AsIntrinsic())
	iterator.ToEnd()
	ass.Equal(t, 'F', iterator.GetPrevious().AsIntrinsic())

}

func TestListConstructors(t *tes.T) {
	fra.List[int64]()
	var sequence = fra.ListFromArray([]int64{1, 2, 3})
	var list = fra.ListFromSequence(sequence)
	ass.Equal(t, sequence.AsArray(), list.AsArray())
}

func TestListsWithStrings(t *tes.T) {
	var collator = fra.CollatorClass[fra.ListLike[string]]().Collator()
	var foo = fra.ListFromArray([]string{"foo"})
	var bar = fra.ListFromArray([]string{"bar"})
	var baz = fra.ListFromArray([]string{"baz"})
	var foz = fra.ListFromArray([]string{"foz"})
	var barbaz = fra.ListFromArray([]string{"bar", "baz"})
	var bazbaz = fra.ListFromArray([]string{"baz", "baz"})
	var foobar = fra.ListFromArray([]string{"foo", "bar"})
	var baxbaz = fra.ListFromArray([]string{"bax", "baz"})
	var baxbez = fra.ListFromArray([]string{"bax", "bez"})
	var barfoobax = fra.ListFromArray([]string{"bar", "foo", "bax"})
	var foobazbar = fra.ListFromArray([]string{"foo", "baz", "bar"})
	var foobarbaz = fra.ListFromArray([]string{"foo", "bar", "baz"})
	var barbazfoo = fra.ListFromArray([]string{"bar", "baz", "foo"})
	var list = fra.List[string]()
	ass.True(t, list.IsEmpty())
	ass.True(t, list.GetSize() == 0)
	ass.False(t, list.ContainsValue("bax"))
	ass.Equal(t, []string{}, list.AsArray())
	var iterator = list.GetIterator()
	ass.False(t, iterator.HasNext())
	ass.False(t, iterator.HasPrevious())
	iterator.ToStart()
	iterator.ToEnd()
	list.ShuffleValues()
	list.SortValues()
	list.RemoveAll()                      //       [ ]
	list.AppendValue("foo")               //       ["foo"]
	ass.False(t, list.IsEmpty())          //       ["foo"]
	ass.True(t, list.GetSize() == 1)      //       ["foo"]
	ass.Equal(t, "foo", list.GetValue(1)) //       ["foo"]
	list.AppendValues(barbaz)             //       ["foo", "bar", "baz"]
	ass.True(t, list.GetSize() == 3)      //       ["foo", "bar", "baz"]
	ass.Equal(t, "foo", list.GetValue(1)) //       ["foo", "bar", "baz"]
	ass.True(t, collator.CompareValues(fra.ListFromArray(list.AsArray()), list))
	ass.Equal(t, barbaz.AsArray(), list.GetValues(2, 3).AsArray())
	ass.Equal(t, foo.AsArray(), list.GetValues(1, 1).AsArray())
	var list2 = fra.ListFromSequence(list)
	ass.True(t, collator.CompareValues(list, list2))
	var array = fra.ListFromArray([]string{"foo", "bar", "baz"})
	var list3 = fra.ListFromSequence(array)
	list2.SortValues()
	list3.SortValues()
	ass.True(t, collator.CompareValues(list2, list3))
	iterator = list.GetIterator()               // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasNext())             // ["foo", "bar", "baz"]
	ass.False(t, iterator.HasPrevious())        // ["foo", "bar", "baz"]
	ass.Equal(t, "foo", iterator.GetNext())     // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasPrevious())         // ["foo", "bar", "baz"]
	iterator.ToEnd()                            // ["foo", "bar", "baz"]
	ass.False(t, iterator.HasNext())            // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasPrevious())         // ["foo", "bar", "baz"]
	ass.Equal(t, "baz", iterator.GetPrevious()) // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasNext())             // ["foo", "bar", "baz"]
	list.ShuffleValues()                        // [ ?, ?, ? ]
	list.RemoveAll()                            // [ ]
	ass.True(t, list.IsEmpty())                 // [ ]
	ass.True(t, list.GetSize() == 0)            // [ ]
	list.InsertValue(0, "baz")                  // ["baz"]
	ass.True(t, list.GetSize() == 1)            // ["baz"]
	ass.Equal(t, "baz", list.GetValue(-1))      // ["baz"]
	list.InsertValues(0, foobar)                // ["foo", "bar", "baz"]
	ass.True(t, list.GetSize() == 3)            // ["foo", "bar", "baz"]
	ass.Equal(t, "foo", list.GetValue(-3))      // ["foo", "bar", "baz"]
	ass.Equal(t, "bar", list.GetValue(-2))      // ["foo", "bar", "baz"]
	ass.Equal(t, "baz", list.GetValue(-1))      // ["foo", "bar", "baz"]
	list.ReverseValues()                        // ["baz", "bar", "foo"]
	ass.Equal(t, "foo", list.GetValue(-1))      // ["baz", "bar", "foo"]
	ass.Equal(t, "bar", list.GetValue(-2))      // ["baz", "bar", "foo"]
	ass.Equal(t, "baz", list.GetValue(-3))      // ["baz", "bar", "foo"]
	list.ReverseValues()                        // ["foo", "bar", "baz"]
	ass.True(t, list.GetIndex("foz") == 0)      // ["foo", "bar", "baz"]
	ass.True(t, list.GetIndex("baz") == 3)      // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsValue("baz"))      // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsValue("bax"))     // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsAny(baxbaz))       // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsAny(baxbez))      // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsAll(barbaz))       // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsAll(baxbaz))      // ["foo", "bar", "baz"]
	list.SetValue(3, "bax")                     // ["foo", "bar", "bax"]
	list.InsertValues(3, baz)                   // ["foo", "bar", "bax", "baz"]
	ass.True(t, list.GetSize() == 4)            // ["foo", "bar", "bax", "baz"]
	ass.Equal(t, "baz", list.GetValue(4))       // ["foo", "bar", "bax", "baz"]
	list.InsertValue(4, "bar")                  // ["foo", "bar", "bax", "baz", "bar"]
	ass.True(t, list.GetSize() == 5)            // ["foo", "bar", "bax", "baz", "bar"]
	ass.Equal(t, "bar", list.GetValue(5))       // ["foo", "bar", "bax", "baz", "bar"]
	list.InsertValue(2, "foo")                  // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.True(t, list.GetSize() == 6)            // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "bar", list.GetValue(2))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "foo", list.GetValue(3))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "bax", list.GetValue(4))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, bar.AsArray(), list.GetValues(6, 6).AsArray())
	list.InsertValues(5, baz)             //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.True(t, list.GetSize() == 7)      //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "bax", list.GetValue(4)) //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "baz", list.GetValue(5)) //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "baz", list.GetValue(6)) //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, barfoobax.AsArray(), list.GetValues(2, -4).AsArray())
	list.SetValues(2, foobazbar) //                        ["foo", "foo", "baz", "bar", "baz", "baz", "bar"]
	ass.Equal(t, foobazbar.AsArray(), list.GetValues(2, -4).AsArray())
	list.SetValues(-1, foz)
	ass.Equal(t, "foz", list.GetValue(-1)) //      ["foo", "foo", "baz", "bar", "baz", "baz", "foz"]
	list.SortValues()                      //      ["bar", "baz", "baz", "baz", "foo", "foo", "foz"]

	ass.Equal(t, bazbaz.AsArray(), list.RemoveValues(2, -5).AsArray()) // ["bar", "baz", "foo", "foo", "foz"]
	ass.Equal(t, barbaz.AsArray(), list.RemoveValues(1, 2).AsArray())  // ["foo", "foo", "foz"]
	ass.Equal(t, "foz", list.RemoveValue(-1))                          // ["foo", "foo"]
	ass.True(t, list.GetSize() == 2)                                   // ["foo", "foo"]
	list.RemoveAll()                                                   // [ ]
	ass.True(t, list.GetSize() == 0)                                   // [ ]
	list.SortValues()                                                  // [ ]
	list.AppendValues(foobarbaz)                                       // ["foo", "bar", "baz"]
	list.SortValues()                                                  // ["bar", "baz", "foo"]
	ass.Equal(t, barbazfoo.AsArray(), list.AsArray())                  // ["bar", "baz", "foo"]
	list.RemoveAll()                                                   // [ ]
	list.AppendValue("foo")                                            // ["foo"]
	list.SortValues()                                                  // ["foo"]
	ass.True(t, list.GetSize() == 1)                                   // ["foo"]
	ass.Equal(t, "foo", list.GetValue(1))                              // ["foo"]
	list.AppendValue("bar")                                            // ["foo", "bar"]
	list.SortValues()                                                  // ["bar", "foo"]
	ass.True(t, list.GetSize() == 2)                                   // ["bar", "foo"]
	ass.Equal(t, "bar", list.GetValue(1))                              // ["bar", "foo"]
}

func TestListsWithTildes(t *tes.T) {
	var array = fra.ListFromArray([]Integer{3, 1, 4, 5, 9, 2})
	var list = fra.ListFromSequence(array)
	ass.False(t, list.IsEmpty())        // [3,1,4,5,9,2]
	ass.True(t, list.GetSize() == 6)    // [3,1,4,5,9,2]
	ass.True(t, list.GetValue(1) == 3)  // [3,1,4,5,9,2]
	ass.True(t, list.GetValue(-1) == 2) // [3,1,4,5,9,2]
	list.SortValues()                   // [1,2,3,4,5,9]
	ass.True(t, list.GetSize() == 6)    // [1,2,3,4,5,9]
	ass.True(t, list.GetValue(3) == 3)  // [1,2,3,4,5,9]
}

func TestListsWithConcatenate(t *tes.T) {
	var collator = fra.CollatorClass[fra.ListLike[int]]().Collator()
	var onetwothree = fra.ListFromArray([]int{1, 2, 3})
	var fourfivesix = fra.ListFromArray([]int{4, 5, 6})
	var onethrusix = fra.ListFromArray([]int{1, 2, 3, 4, 5, 6})
	var list1 = fra.List[int]()
	list1.AppendValues(onetwothree)
	var list2 = fra.List[int]()
	list2.AppendValues(fourfivesix)
	var list3 = fra.ListClass[int]().Concatenate(list1, list2)
	var list4 = fra.List[int]()
	list4.AppendValues(onethrusix)
	ass.True(t, collator.CompareValues(list3, list4))
}

func TestListsWithEmptyLists(t *tes.T) {
	var collator = fra.Collator[fra.ListLike[int]]()
	var empty = fra.List[int]()
	var list = fra.ListClass[int]().Concatenate(empty, empty)
	ass.True(t, collator.CompareValues(empty, empty))
	ass.True(t, collator.CompareValues(list, empty))
	ass.True(t, collator.CompareValues(empty, list))
	ass.True(t, collator.CompareValues(list, list))
}

func TestQueueConstructors(t *tes.T) {
	fra.Queue[int64]()
	fra.QueueWithCapacity[int64](5)
	var sequence = fra.QueueFromArray([]int64{1, 2, 3})
	var queue = fra.QueueFromSequence(sequence)
	ass.Equal(t, sequence.AsArray(), queue.AsArray())
}

func TestQueueWithConcurrency(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with a specific capacity.
	var queue = fra.QueueWithCapacity[int](12)
	ass.True(t, queue.GetCapacity() == 12)
	ass.True(t, queue.IsEmpty())
	ass.True(t, queue.GetSize() == 0)

	// Add some values to the queue.
	for i := 1; i < 10; i++ {
		queue.AddValue(i)
	}
	ass.True(t, queue.GetSize() == 9)

	// Remove values from the queue in the background.
	group.Add(1)
	go func() {
		defer group.Done()
		var value int
		var ok = true
		for i := 1; ok; i++ {
			value, ok = queue.RemoveFirst()
			if ok {
				ass.Equal(t, i, value)
			}
		}
		queue.RemoveAll()
	}()

	// Add some more values to the queue.
	for i := 10; i < 101; i++ {
		queue.AddValue(i)
	}
	queue.CloseChannel()
}

func TestQueueWithFork(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with a fan out of two.
	var input = fra.QueueWithCapacity[int](3)
	var outputs = fra.QueueClass[int]().Fork(group, input, 2)

	// Remove values from the output queues in the background.
	var readOutput = func(output fra.QueueLike[int], name string) {
		defer group.Done()
		var value int
		var ok = true
		for i := 1; ok; i++ {
			value, ok = output.RemoveFirst()
			if ok {
				ass.Equal(t, i, value)
			}
		}
	}
	group.Add(2)
	var iterator = outputs.GetIterator()
	for iterator.HasNext() {
		var output = iterator.GetNext()
		go readOutput(output, "output")
	}

	// Add values to the input queue.
	for i := 1; i < 11; i++ {
		input.AddValue(i)
	}
	input.CloseChannel()
}

func TestQueueWithInvalidFanOut(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with an invalid fan out.
	var input = fra.QueueWithCapacity[int](3)
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The fan out size for a queue must be greater than one.", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	fra.QueueClass[int]().Fork(group, input, 1) // Should panic here.
}

func TestQueueWithSplitAndJoin(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with a split of five outputs and a join back to one.
	var input = fra.QueueWithCapacity[int](3)
	var split = fra.QueueClass[int]().Split(group, input, 5)
	var output = fra.QueueClass[int]().Join(group, split)

	// Remove values from the output queue in the background.
	group.Add(1)
	go func() {
		defer group.Done()
		var value int
		var ok = true
		for i := 1; ok; i++ {
			value, ok = output.RemoveFirst()
			if ok {
				ass.Equal(t, i, value)
			}
		}
	}()

	// Add values to the input queue.
	for i := 1; i < 21; i++ {
		input.AddValue(i)
	}
	input.CloseChannel()
}

func TestQueueWithInvalidSplit(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with an invalid fan out.
	var input = fra.QueueWithCapacity[int](3)
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The size of the split must be greater than one.", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	fra.QueueClass[int]().Split(group, input, 1) // Should panic here.
}

func TestQueueWithInvalidJoin(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with an invalid fan out.
	var inputs = fra.List[fra.QueueLike[int]]()
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The number of input queues for a join must be at least one.", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	fra.QueueClass[int]().Join(group, inputs) // Should panic here.
	defer group.Done()
}

func TestSetConstructors(t *tes.T) {
	var collator = fra.Collator[int64]()
	fra.Set[int64]()
	fra.SetWithCollator(collator)
	var sequence = fra.SetFromArray([]int64{1, 2, 3})
	var set = fra.SetFromSequence(sequence)
	ass.Equal(t, sequence.AsArray(), set.AsArray())
}

func TestSetsWithStrings(t *tes.T) {
	var collator = fra.Collator[fra.SetLike[string]]()
	fra.List[string]()
	var empty = []string{}
	var bazbar = fra.ListFromArray([]string{"baz", "bar"})
	var bazfoo = fra.ListFromArray([]string{"baz", "foo"})
	var baxbaz = fra.ListFromArray([]string{"bax", "baz"})
	var baxbez = fra.ListFromArray([]string{"bax", "bez"})
	var barbaz = fra.ListFromArray([]string{"bar", "baz"})
	var bar = fra.ListFromArray([]string{"bar"})
	var set = fra.Set[string]()                                   // [ ]
	ass.True(t, set.IsEmpty())                                    // [ ]
	ass.True(t, set.GetSize() == 0)                               // [ ]
	ass.False(t, set.ContainsValue("bax"))                        // [ ]
	ass.Equal(t, empty, set.AsArray())                            // [ ]
	var iterator = set.GetIterator()                              // [ ]
	ass.False(t, iterator.HasNext())                              // [ ]
	ass.False(t, iterator.HasPrevious())                          // [ ]
	iterator.ToStart()                                            // [ ]
	iterator.ToEnd()                                              // [ ]
	set.RemoveAll()                                               // [ ]
	set.RemoveValue("foo")                                        // [ ]
	set.AddValue("foo")                                           // ["foo"]
	ass.False(t, set.IsEmpty())                                   // ["foo"]
	ass.True(t, set.GetSize() == 1)                               // ["foo"]
	ass.Equal(t, "foo", set.GetValue(1))                          // ["foo"]
	ass.True(t, set.GetIndex("baz") == 0)                         // ["foo"]
	ass.True(t, set.ContainsValue("foo"))                         // ["foo"]
	ass.False(t, set.ContainsValue("bax"))                        // ["foo"]
	set.AddValues(bazbar)                                         // ["bar", "baz", "foo"]
	ass.True(t, set.GetSize() == 3)                               // ["bar", "baz", "foo"]
	ass.True(t, set.GetIndex("baz") == 2)                         // ["bar", "baz", "foo"]
	ass.Equal(t, "bar", set.GetValue(1))                          // ["bar", "baz", "foo"]
	ass.Equal(t, bazfoo.AsArray(), set.GetValues(2, 3).AsArray()) // ["bar", "baz", "foo"]
	ass.Equal(t, bar.AsArray(), set.GetValues(1, 1).AsArray())    // ["bar", "baz", "foo"]
	var set2 = fra.SetFromSequence(set)                           // ["bar", "baz", "foo"]
	ass.True(t, collator.CompareValues(set, set2))                // ["bar", "baz", "foo"]
	var array = fra.ListFromArray([]string{"foo", "bar", "baz"})  // ["bar", "baz", "foo"]
	var set3 = fra.SetFromSequence(array)                         // ["bar", "baz", "foo"]
	ass.True(t, collator.CompareValues(set2, set3))               // ["bar", "baz", "foo"]
	iterator = set.GetIterator()                                  // ["bar", "baz", "foo"]
	ass.True(t, iterator.HasNext())                               // ["bar", "baz", "foo"]
	ass.False(t, iterator.HasPrevious())                          // ["bar", "baz", "foo"]
	ass.Equal(t, "bar", string(iterator.GetNext()))               // ["bar", "baz", "foo"]
	ass.True(t, iterator.HasPrevious())                           // ["bar", "baz", "foo"]
	iterator.ToEnd()                                              // ["bar", "baz", "foo"]
	ass.False(t, iterator.HasNext())                              // ["bar", "baz", "foo"]
	ass.True(t, iterator.HasPrevious())                           // ["bar", "baz", "foo"]
	ass.Equal(t, "foo", string(iterator.GetPrevious()))           // ["bar", "baz", "foo"]
	ass.True(t, iterator.HasNext())                               // ["bar", "baz", "foo"]
	ass.True(t, set.ContainsValue("baz"))                         // ["bar", "baz", "foo"]
	ass.False(t, set.ContainsValue("bax"))                        // ["bar", "baz", "foo"]
	ass.True(t, set.ContainsAny(baxbaz))                          // ["bar", "baz", "foo"]
	ass.False(t, set.ContainsAny(baxbez))                         // ["bar", "baz", "foo"]
	ass.True(t, set.ContainsAll(barbaz))                          // ["bar", "baz", "foo"]
	ass.False(t, set.ContainsAll(baxbaz))                         // ["bar", "baz", "foo"]
	set.RemoveAll()                                               // [ ]
	ass.True(t, set.IsEmpty())                                    // [ ]
	ass.True(t, set.GetSize() == 0)                               // [ ]
}

func TestSetsWithIntegers(t *tes.T) {
	var array = fra.ListFromArray([]int{3, 1, 4, 5, 9, 2})
	var set = fra.Set[int]()           // [ ]
	set.AddValues(array)               // [1,2,3,4,5,9]
	ass.False(t, set.IsEmpty())        // [1,2,3,4,5,9]
	ass.True(t, set.GetSize() == 6)    // [1,2,3,4,5,9]
	ass.True(t, set.GetValue(1) == 1)  // [1,2,3,4,5,9]
	ass.True(t, set.GetValue(-1) == 9) // [1,2,3,4,5,9]
	set.RemoveValue(6)                 // [1,2,3,4,5,9]
	ass.True(t, set.GetSize() == 6)    // [1,2,3,4,5,9]
	set.RemoveValue(3)                 // [1,2,4,5,9]
	ass.True(t, set.GetSize() == 5)    // [1,2,4,5,9]
	ass.True(t, set.GetValue(3) == 4)  // [1,2,4,5,9]
}

func TestSetsWithTildes(t *tes.T) {
	var array = fra.ListFromArray([]Integer{3, 1, 4, 5, 9, 2})
	var set = fra.Set[Integer]()       // [ ]
	set.AddValues(array)               // [1,2,3,4,5,9]
	ass.False(t, set.IsEmpty())        // [1,2,3,4,5,9]
	ass.True(t, set.GetSize() == 6)    // [1,2,3,4,5,9]
	ass.True(t, set.GetValue(1) == 1)  // [1,2,3,4,5,9]
	ass.True(t, set.GetValue(-1) == 9) // [1,2,3,4,5,9]
	set.RemoveValue(6)                 // [1,2,3,4,5,9]
	ass.True(t, set.GetSize() == 6)    // [1,2,3,4,5,9]
	set.RemoveValue(3)                 // [1,2,4,5,9]
	ass.True(t, set.GetSize() == 5)    // [1,2,4,5,9]
	ass.True(t, set.GetValue(3) == 4)  // [1,2,4,5,9]
}

func TestSetsWithSets(t *tes.T) {
	var array1 = fra.ListFromArray([]int{3, 1, 4, 5, 9, 2})
	var array2 = fra.ListFromArray([]int{7, 1, 4, 5, 9, 2})
	var set1 = fra.Set[int]()
	set1.AddValues(array1)
	var set2 = fra.Set[int]()
	set2.AddValues(array2)
	var set = fra.Set[fra.SetLike[int]]()
	set.AddValue(set1)
	set.AddValue(set2)
	ass.False(t, set.IsEmpty())
	ass.True(t, set.GetSize() == 2)
	ass.Equal(t, set1, set.GetValue(1))
	ass.Equal(t, set2, set.GetValue(-1))
	set.RemoveValue(set1)
	ass.True(t, set.GetSize() == 1)
	set.RemoveAll()
	ass.True(t, set.GetSize() == 0)
}

func TestSetsWithAnd(t *tes.T) {
	var collator = fra.Collator[fra.SetLike[int]]()
	var array1 = fra.ListFromArray([]int{3, 1, 2})
	var array2 = fra.ListFromArray([]int{3, 2, 4})
	var array3 = fra.ListFromArray([]int{3, 2})
	var set1 = fra.Set[int]()
	set1.AddValues(array1)
	var set2 = fra.Set[int]()
	set2.AddValues(array2)
	var set3 = fra.SetClass[int]().And(set1, set2)
	var set4 = fra.Set[int]()
	set4.AddValues(array3)
	ass.True(t, collator.CompareValues(set3, set4))
}

func TestSetsWithSan(t *tes.T) {
	var collator = fra.Collator[fra.SetLike[int]]()
	var array1 = fra.ListFromArray([]int{3, 1, 2})
	var array2 = fra.ListFromArray([]int{3, 2, 4})
	var array3 = fra.ListFromArray([]int{1})
	var set1 = fra.Set[int]()
	set1.AddValues(array1)
	var set2 = fra.Set[int]()
	set2.AddValues(array2)
	var set3 = fra.SetClass[int]().San(set1, set2)
	var set4 = fra.Set[int]()
	set4.AddValues(array3)
	ass.True(t, collator.CompareValues(set3, set4))
}

func TestSetsWithIor(t *tes.T) {
	var collator = fra.Collator[fra.SetLike[int]]()
	var array1 = fra.ListFromArray([]int{3, 1, 5})
	var array2 = fra.ListFromArray([]int{6, 2, 4})
	var array3 = fra.ListFromArray([]int{1, 3, 5, 6, 2, 4})
	var set1 = fra.Set[int]()
	set1.AddValues(array1)
	var set2 = fra.Set[int]()
	set2.AddValues(array2)
	var set3 = fra.SetClass[int]().Ior(set1, set2)
	ass.True(t, set3.ContainsAll(set1))
	ass.True(t, set3.ContainsAll(set2))
	var set4 = fra.Set[int]()
	set4.AddValues(array3)
	ass.True(t, collator.CompareValues(set3, set4))
}

func TestSetsWithXor(t *tes.T) {
	var collator = fra.Collator[fra.SetLike[int]]()
	var array1 = fra.ListFromArray([]int{2, 3, 1, 5})
	var array2 = fra.ListFromArray([]int{6, 2, 5, 4})
	var array3 = fra.ListFromArray([]int{1, 3, 4, 6})
	var set1 = fra.Set[int]()
	set1.AddValues(array1)
	var set2 = fra.Set[int]()
	set2.AddValues(array2)
	var set3 = fra.SetClass[int]().Xor(set1, set2)
	var set4 = fra.Set[int]()
	set4.AddValues(array3)
	ass.True(t, collator.CompareValues(set3, set4))
}

func TestSetsWithEmptySets(t *tes.T) {
	var collator = fra.Collator[fra.SetLike[int]]()
	var set1 = fra.Set[int]()
	var set2 = fra.Set[int]()
	var set3 = fra.SetClass[int]().And(set1, set2)
	var set4 = fra.SetClass[int]().San(set1, set2)
	var set5 = fra.SetClass[int]().Ior(set1, set2)
	var set6 = fra.SetClass[int]().Xor(set1, set2)
	ass.True(t, collator.CompareValues(set3, set4))
	ass.True(t, collator.CompareValues(set4, set5))
	ass.True(t, collator.CompareValues(set5, set6))
	ass.True(t, collator.CompareValues(set6, set1))
}

func TestStackConstructors(t *tes.T) {
	fra.Stack[int64]()
	fra.StackWithCapacity[int64](5)
	var sequence = fra.StackFromArray([]int64{1, 2, 3})
	var stack = fra.StackFromSequence(sequence)
	ass.Equal(t, sequence.AsArray(), stack.AsArray())
}

func TestStackWithSmallCapacity(t *tes.T) {
	var stack = fra.StackWithCapacity[int](1)
	stack.AddValue(1)
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "Attempted to add a value onto a stack that has reached its capacity.", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	stack.AddValue(2) // This should panic.
}

func TestEmptyStackRemoval(t *tes.T) {
	var stack = fra.Stack[int]()
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "Attempted to remove a value from an empty stack!", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	stack.RemoveLast() // This should panic.
}

func TestStacksWithStrings(t *tes.T) {
	var stack = fra.Stack[string]()
	ass.True(t, stack.IsEmpty())
	ass.True(t, stack.GetSize() == 0)
	stack.RemoveAll()
	stack.AddValue("foo")
	stack.AddValue("bar")
	stack.AddValue("baz")
	ass.True(t, stack.GetSize() == 3)
	var last = stack.GetLast()
	ass.Equal(t, last, stack.RemoveLast())
	ass.True(t, stack.GetSize() == 2)
	ass.Equal(t, "bar", stack.RemoveLast())
	ass.True(t, stack.GetSize() == 1)
	stack.RemoveAll()
}

// ELEMENT

func TestUnits(t *tes.T) {
	ass.Equal(t, "Degrees", fra.Degrees.String())
	ass.Equal(t, "Radians", fra.Radians.String())
	ass.Equal(t, "Gradians", fra.Gradians.String())
}

func TestZeroAngles(t *tes.T) {
	var v = fra.Angle(0)
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "~0", v.AsString())
	ass.Equal(t, v, fra.AngleClass().Zero())

	v = fra.Angle(2.0 * mat.Pi)
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "~0", v.AsString())
	ass.Equal(t, v, fra.AngleClass().Zero())

	v = fra.AngleFromString("~0")
	ass.Equal(t, "~0", v.AsString())
	ass.Equal(t, v, fra.AngleClass().Zero())

	v = fra.AngleFromString("~τ")
	ass.Equal(t, "~τ", v.AsString())
	ass.Equal(t, v, fra.AngleClass().Tau())
}

func TestPositiveAngles(t *tes.T) {
	var v = fra.Angle(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())
	ass.Equal(t, v, fra.AngleClass().Pi())

	v = fra.AngleFromString("~π")
	ass.Equal(t, "~π", v.AsString())
	ass.Equal(t, v, fra.AngleClass().Pi())
}

func TestNegativeAngles(t *tes.T) {
	var v = fra.Angle(-mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())
	ass.Equal(t, v, fra.AngleClass().Pi())

	v = fra.Angle(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v.AsFloat())
}

func TestAnglesLibrary(t *tes.T) {
	var class = fra.AngleClass()
	var v0 = class.Zero()
	var v1 = fra.Angle(mat.Pi * 0.25)
	var v2 = fra.Angle(mat.Pi * 0.5)
	var v3 = fra.Angle(mat.Pi * 0.75)
	var v4 = class.Pi()
	var v5 = fra.Angle(mat.Pi * 1.25)
	var v6 = fra.Angle(mat.Pi * 1.5)
	var v7 = fra.Angle(mat.Pi * 1.75)
	var v8 = class.Tau()

	ass.Equal(t, v4, class.Inverse(v0))
	ass.Equal(t, v5, class.Inverse(v1))
	ass.Equal(t, v6, class.Inverse(v2))
	ass.Equal(t, v7, class.Inverse(v3))
	ass.Equal(t, v0, class.Inverse(v4))
	ass.Equal(t, v4, class.Inverse(v8))

	ass.Equal(t, v1, class.Sum(v0, v1))
	ass.Equal(t, v0, class.Difference(v1, v1))
	ass.Equal(t, v3, class.Sum(v1, v2))
	ass.Equal(t, v1, class.Difference(v3, v2))
	ass.Equal(t, v5, class.Sum(v2, v3))
	ass.Equal(t, v2, class.Difference(v5, v3))
	ass.Equal(t, v7, class.Sum(v3, v4))
	ass.Equal(t, v3, class.Difference(v7, v4))
	ass.Equal(t, v1, class.Sum(v8, v1))
	ass.Equal(t, v0, class.Difference(v8, v8))

	ass.Equal(t, v3, class.Scaled(v1, 3.0))
	ass.Equal(t, v0, class.Scaled(v4, 2.0))
	ass.Equal(t, v4, class.Scaled(v4, -1.0))
	ass.Equal(t, v0, class.Scaled(v8, 1.0))

	ass.Equal(t, v0, class.ArcCosine(class.Cosine(v0)))
	ass.Equal(t, v1, class.ArcCosine(class.Cosine(v1)))
	ass.Equal(t, v2, class.ArcCosine(class.Cosine(v2)))
	ass.Equal(t, v3, class.ArcCosine(class.Cosine(v3)))
	ass.Equal(t, v4, class.ArcCosine(class.Cosine(v4)))
	ass.Equal(t, v0, class.ArcCosine(class.Cosine(v8)))

	ass.Equal(t, v0, class.ArcSine(class.Sine(v0)))
	ass.Equal(t, v1, class.ArcSine(class.Sine(v1)))
	ass.Equal(t, v2, class.ArcSine(class.Sine(v2)))
	ass.Equal(t, v6, class.ArcSine(class.Sine(v6)))
	ass.Equal(t, v7, class.ArcSine(class.Sine(v7)))
	ass.Equal(t, v0, class.ArcSine(class.Sine(v8)))

	ass.Equal(t, v0, class.ArcTangent(class.Cosine(v0), class.Sine(v0)))
	ass.Equal(t, v1, class.ArcTangent(class.Cosine(v1), class.Sine(v1)))
	ass.Equal(t, v2, class.ArcTangent(class.Cosine(v2), class.Sine(v2)))
	ass.Equal(t, v3, class.ArcTangent(class.Cosine(v3), class.Sine(v3)))
	ass.Equal(t, v4, class.ArcTangent(class.Cosine(v4), class.Sine(v4)))
	ass.Equal(t, v5, class.ArcTangent(class.Cosine(v5), class.Sine(v5)))
	ass.Equal(t, v0, class.ArcTangent(class.Cosine(v8), class.Sine(v8)))
}

func TestFalseBooleans(t *tes.T) {
	ass.False(t, fra.BooleanClass().False().AsIntrinsic())
	var v = fra.Boolean(false)
	ass.False(t, v.AsIntrinsic())
	v = fra.BooleanFromString("false")
	ass.Equal(t, "false", v.AsString())
	ass.Equal(t, v, fra.BooleanClass().False())
}

func TestTrueBooleans(t *tes.T) {
	ass.True(t, fra.BooleanClass().True().AsIntrinsic())
	var v = fra.Boolean(true)
	ass.True(t, v.AsIntrinsic())
	v = fra.BooleanFromString("true")
	ass.Equal(t, "true", v.AsString())
	ass.Equal(t, v, fra.BooleanClass().True())
}

func TestBooleansLibrary(t *tes.T) {
	var T = fra.Boolean(true)
	var F = fra.Boolean(false)
	var class = fra.BooleanClass()

	var andNot = class.And(class.Not(T), class.Not(T))
	var notIor = class.Not(class.Ior(T, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(T), class.Not(F))
	notIor = class.Not(class.Ior(T, F))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(T))
	notIor = class.Not(class.Ior(F, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(F))
	notIor = class.Not(class.Ior(F, F))
	ass.Equal(t, andNot, notIor)

	var san = class.And(T, class.Not(T))
	ass.Equal(t, san, class.San(T, T))

	san = class.And(T, class.Not(F))
	ass.Equal(t, san, class.San(T, F))

	san = class.And(F, class.Not(T))
	ass.Equal(t, san, class.San(F, T))

	san = class.And(F, class.Not(F))
	ass.Equal(t, san, class.San(F, F))

	var xor = class.Ior(class.San(T, T), class.San(T, T))
	ass.Equal(t, xor, class.Xor(T, T))

	xor = class.Ior(class.San(T, F), class.San(F, T))
	ass.Equal(t, xor, class.Xor(T, F))

	xor = class.Ior(class.San(F, T), class.San(T, F))
	ass.Equal(t, xor, class.Xor(F, T))

	xor = class.Ior(class.San(F, F), class.San(F, F))
	ass.Equal(t, xor, class.Xor(F, F))
}

var DurationClass = fra.DurationClass()

var zero uint = 0
var one uint = 1

func TestZeroDurations(t *tes.T) {
	var v = fra.Duration(0)
	ass.Equal(t, 0, v.AsInteger())
	ass.Equal(t, zero, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsMilliseconds())
	ass.Equal(t, 0.0, v.AsSeconds())
	ass.Equal(t, 0.0, v.AsMinutes())
	ass.Equal(t, 0.0, v.AsHours())
	ass.Equal(t, 0.0, v.AsDays())
	ass.Equal(t, 0.0, v.AsWeeks())
	ass.Equal(t, 0.0, v.AsMonths())
	ass.Equal(t, 0.0, v.AsYears())
	ass.Equal(t, zero, v.GetMilliseconds())
	ass.Equal(t, zero, v.GetSeconds())
	ass.Equal(t, zero, v.GetMinutes())
	ass.Equal(t, zero, v.GetHours())
	ass.Equal(t, zero, v.GetDays())
	ass.Equal(t, zero, v.GetWeeks())
	ass.Equal(t, zero, v.GetMonths())
	ass.Equal(t, zero, v.GetYears())
}

func TestStringDurations(t *tes.T) {
	var duration = fra.DurationFromString("~P1Y2M3DT4H5M6S")
	ass.Equal(t, "~P1Y2M3DT4H5M6S", duration.AsString())
	duration = fra.DurationFromString("~P0W")
	ass.Equal(t, "~P0W", duration.AsString())
}

func TestDurations(t *tes.T) {
	var v = fra.Duration(60000)
	ass.Equal(t, "~PT1M", v.AsString())
	ass.Equal(t, 60000, v.AsInteger())
	ass.Equal(t, uint(60000), v.AsIntrinsic())
	ass.Equal(t, 60000.0, v.AsMilliseconds())
	ass.Equal(t, 60.0, v.AsSeconds())
	ass.Equal(t, 1.0, v.AsMinutes())
	ass.Equal(t, 0.016666666666666666, v.AsHours())
	ass.Equal(t, 0.0006944444444444445, v.AsDays())
	ass.Equal(t, 9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, 2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, 1.9013243104086858e-06, v.AsYears())
	ass.Equal(t, zero, v.GetMilliseconds())
	ass.Equal(t, zero, v.GetSeconds())
	ass.Equal(t, one, v.GetMinutes())
	ass.Equal(t, zero, v.GetHours())
	ass.Equal(t, zero, v.GetDays())
	ass.Equal(t, zero, v.GetWeeks())
	ass.Equal(t, zero, v.GetMonths())
	ass.Equal(t, zero, v.GetYears())
}

var GlyphClass = fra.GlyphClass()

func TestGlyphs(t *tes.T) {
	var v = fra.GlyphFromString("'''")
	ass.Equal(t, "'''", v.AsString())

	v = fra.Glyph('a')
	ass.Equal(t, "'a'", v.AsString())

	v = fra.Glyph('"')
	ass.Equal(t, `'"'`, v.AsString())

	v = fra.Glyph('😊')
	ass.Equal(t, "'😊'", v.AsString())

	v = fra.Glyph('界')
	ass.Equal(t, "'界'", v.AsString())

	v = fra.Glyph('\'')
	ass.Equal(t, "'''", v.AsString())

	v = fra.Glyph('\\')
	ass.Equal(t, "'\\'", v.AsString())

	v = fra.Glyph('\n')
	ass.Equal(t, "'\n'", v.AsString())

	v = fra.Glyph('\t')
	ass.Equal(t, "'\t'", v.AsString())
}

var MomentClass = fra.MomentClass()

func TestIntegerMoments(t *tes.T) {
	var v = fra.Moment(1238589296789)
	ass.False(t, v.IsNegative())
	ass.Equal(t, 1238589296789, v.AsIntrinsic())
	ass.Equal(t, 1238589296789, v.AsInteger())
	ass.Equal(t, 1238589296789.0, v.AsMilliseconds())
	ass.Equal(t, 1238589296.789, v.AsSeconds())
	ass.Equal(t, 20643154.946483333, v.AsMinutes())
	ass.Equal(t, 344052.58244138886, v.AsHours())
	ass.Equal(t, 14335.524268391204, v.AsDays())
	ass.Equal(t, 2047.9320383416004, v.AsWeeks())
	ass.Equal(t, 470.9919881193849, v.AsMonths())
	ass.Equal(t, 39.24933234328208, v.AsYears())
	ass.Equal(t, uint(789), v.GetMilliseconds())
	ass.Equal(t, uint(56), v.GetSeconds())
	ass.Equal(t, uint(34), v.GetMinutes())
	ass.Equal(t, uint(12), v.GetHours())
	ass.Equal(t, uint(1), v.GetDays())
	ass.Equal(t, uint(14), v.GetWeeks())
	ass.Equal(t, uint(4), v.GetMonths())
	ass.Equal(t, uint(2009), v.GetYears())
}

func TestStringMoments(t *tes.T) {
	var v = fra.MomentFromString("<-1-02-03T04:05:06.700>")
	ass.True(t, v.IsNegative())
	ass.Equal(t, "<-1-02-03T04:05:06.700>", v.AsString())
}

func TestMomentsLibrary(t *tes.T) {
	var before = fra.Now()
	var duration = fra.Duration(12345)
	var after = fra.Moment(before.AsInteger() + duration.AsInteger())
	var class = fra.MomentClass()

	ass.Equal(t, duration, class.Duration(before, after))
	ass.Equal(t, duration, class.Duration(after, before))
	ass.Equal(t, after, class.Later(before, duration))
	ass.Equal(t, before, class.Earlier(after, duration))
}

func TestZero(t *tes.T) {
	var v = fra.Number(0 + 0i)
	ass.Equal(t, 0+0i, v.AsIntrinsic())
	ass.True(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.True(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, "0", v.AsString())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, v, fra.NumberClass().Zero())
}

func TestInfinity(t *tes.T) {
	var v = fra.Number(cmp.Inf())
	ass.Equal(t, cmp.Inf(), v.AsIntrinsic())
	ass.False(t, v.IsZero())
	ass.True(t, v.IsInfinite())
	ass.True(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, "∞", v.AsString())
	ass.Equal(t, mat.Inf(1), v.AsFloat())
	ass.Equal(t, mat.Inf(1), v.GetReal())
	ass.Equal(t, mat.Inf(1), v.GetImaginary())
	ass.Equal(t, v, fra.NumberClass().Infinity())
}

func TestUndefined(t *tes.T) {
	var v = fra.Number(cmp.NaN())
	ass.True(t, cmp.IsNaN(v.AsIntrinsic()))
	ass.False(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.False(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.True(t, mat.IsNaN(v.AsFloat()))
	ass.True(t, mat.IsNaN(v.GetReal()))
	ass.True(t, mat.IsNaN(v.GetImaginary()))
}

func TestPositivePureReals(t *tes.T) {
	var v = fra.Number(0.25)
	ass.Equal(t, 0.25+0i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.25, v.AsFloat())
	ass.Equal(t, 0.25, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	var integer = 5
	v = fra.NumberFromInteger(integer)
	ass.Equal(t, 5.0, v.AsFloat())
	var float = 5.0
	v = fra.NumberFromFloat(float)
	ass.Equal(t, 5.0, v.AsFloat())
	v = fra.NumberFromString("1.23456789E+100")
	ass.Equal(t, "1.23456789E+100", v.AsString())
	v = fra.NumberFromString("1.23456789E-10")
	ass.Equal(t, "1.23456789E-10", v.AsString())
}

func TestPositivePureImaginaries(t *tes.T) {
	var v = fra.Number(0.25i)
	ass.Equal(t, 0+0.25i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.25, v.GetImaginary())
}

func TestNegativePureReals(t *tes.T) {
	var v = fra.Number(-0.75)
	ass.Equal(t, -0.75+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -0.75, v.AsFloat())
	ass.Equal(t, -0.75, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestNegativePureImaginaries(t *tes.T) {
	var v = fra.Number(-0.75i)
	ass.Equal(t, 0-0.75i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -0.75, v.GetImaginary())
}

func TestNumberFromPolar(t *tes.T) {
	var v = fra.NumberFromPolar(1.0, mat.Pi)
	ass.Equal(t, -1.0+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetPhase())

	v = fra.NumberFromString("5e^~1i")
	ass.Equal(t, 5.0, v.GetMagnitude())
	ass.Equal(t, 1.0, v.GetPhase())
	ass.Equal(t, "5e^~1i", v.AsPolar())
}

func TestNumberFromString(t *tes.T) {
	var v = fra.NumberFromString("1e^~πi")
	ass.Equal(t, -1.0+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, "-1", v.AsString())
	ass.Equal(t, "1e^~πi", v.AsPolar())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetPhase())

	v = fra.NumberFromString("-1.2-3.4i")
	ass.Equal(t, "-1.2-3.4i", v.AsString())
	ass.Equal(t, -1.2, v.GetReal())
	ass.Equal(t, -3.4, v.GetImaginary())

	v = fra.NumberFromString("-π+τi")
	ass.Equal(t, "-π+τi", v.AsString())
	ass.Equal(t, -3.141592653589793, v.GetReal())
	ass.Equal(t, 6.283185307179586, v.GetImaginary())

	v = fra.NumberFromString("undefined")
	ass.Equal(t, "undefined", v.AsString())
	ass.False(t, v.IsDefined())
	ass.False(t, v.HasMagnitude())

	v = fra.NumberFromString("+infinity")
	ass.Equal(t, "+∞", v.AsString())
	ass.True(t, v.IsMaximum())
	ass.False(t, v.HasMagnitude())

	v = fra.NumberFromString("infinity")
	ass.Equal(t, "∞", v.AsString())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = fra.NumberFromString("-infinity")
	ass.Equal(t, "-∞", v.AsString())
	ass.True(t, v.IsMinimum())
	ass.False(t, v.HasMagnitude())

	v = fra.NumberFromString("∞")
	ass.Equal(t, "∞", v.AsString())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = fra.NumberFromString("-∞")
	ass.Equal(t, "-∞", v.AsString())
	ass.True(t, v.IsMinimum())
	ass.False(t, v.HasMagnitude())

	v = fra.NumberFromString("+1")
	ass.Equal(t, "1", v.AsString())
	ass.Equal(t, 1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, 0.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = fra.NumberFromString("1")
	ass.Equal(t, "1", v.AsString())
	ass.Equal(t, 1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, 0.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = fra.NumberFromString("-π")
	ass.Equal(t, "-π", v.AsString())
	ass.Equal(t, -mat.Pi, v.GetReal())
	ass.Equal(t, mat.Pi, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.True(t, v.IsNegative())

	v = fra.NumberFromString("+1i")
	ass.Equal(t, "1i", v.AsString())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi/2.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = fra.NumberFromString("1i")
	ass.Equal(t, "1i", v.AsString())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi/2.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = fra.NumberFromString("-1i")
	ass.Equal(t, "-1i", v.AsString())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, -mat.Pi/2.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = fra.NumberFromString("-1.2345678E+90")
	ass.Equal(t, "-1.2345678E+90", v.AsString())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.2345678e+90, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())

	v = fra.NumberFromString("-1.2345678E+90i")
	ass.Equal(t, "-1.2345678E+90i", v.AsString())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -1.2345678e+90, v.GetImaginary())

	v = fra.NumberFromString("1.2345678E+90e^~1.2345678E-90i")
	ass.Equal(t, "1.2345678E+90e^~1.2345678E-90i", v.AsPolar())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 1.2345678e+90, v.GetMagnitude())
	ass.Equal(t, 1.2345678e-90, v.GetPhase())
}

func TestNumberLibrary(t *tes.T) {
	var class = fra.NumberClass()
	var zero = class.Zero()
	var i = class.I()
	var minusi = fra.Number(-1i)
	var half = fra.Number(0.5)
	var minushalf = fra.Number(-0.5)
	var one = class.One()
	var minusone = fra.Number(-1)
	var two = fra.Number(2.0)
	var minustwo = fra.Number(-2.0)
	var infinity = class.Infinity()
	var undefined = class.Undefined()

	//	-z
	ass.Equal(t, zero, class.Inverse(zero))
	ass.Equal(t, minushalf, class.Inverse(half))
	ass.Equal(t, minusone, class.Inverse(one))
	ass.Equal(t, minusi, class.Inverse(i))
	ass.Equal(t, infinity, class.Inverse(infinity))
	ass.False(t, class.Inverse(undefined).IsDefined())

	//	z + zero => z
	ass.Equal(t, minusi, class.Sum(minusi, zero))
	ass.Equal(t, minusone, class.Sum(minusone, zero))
	ass.Equal(t, zero, class.Sum(zero, zero))
	ass.Equal(t, one, class.Sum(one, zero))
	ass.Equal(t, i, class.Sum(i, zero))
	ass.Equal(t, infinity, class.Sum(infinity, zero))
	ass.False(t, class.Sum(undefined, zero).IsDefined())

	//	z + infinity => infinity
	ass.Equal(t, infinity, class.Sum(minusi, infinity))
	ass.Equal(t, infinity, class.Sum(minusone, infinity))
	ass.Equal(t, infinity, class.Sum(zero, infinity))
	ass.Equal(t, infinity, class.Sum(one, infinity))
	ass.Equal(t, infinity, class.Sum(i, infinity))
	ass.Equal(t, infinity, class.Sum(infinity, infinity))
	ass.False(t, class.Sum(undefined, infinity).IsDefined())

	//	z - infinity => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Difference(minusi, infinity))
	ass.Equal(t, infinity, class.Difference(minusone, infinity))
	ass.Equal(t, infinity, class.Difference(zero, infinity))
	ass.Equal(t, infinity, class.Difference(one, infinity))
	ass.Equal(t, infinity, class.Difference(i, infinity))
	ass.False(t, class.Difference(infinity, infinity).IsDefined())
	ass.False(t, class.Difference(undefined, infinity).IsDefined())

	//	infinity - z => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Difference(infinity, minusi))
	ass.Equal(t, infinity, class.Difference(infinity, minusone))
	ass.Equal(t, infinity, class.Difference(infinity, zero))
	ass.Equal(t, infinity, class.Difference(infinity, one))
	ass.Equal(t, infinity, class.Difference(infinity, i))
	ass.False(t, class.Difference(infinity, undefined).IsDefined())

	//	z - z => zero  {z != infinity}
	ass.Equal(t, zero, class.Difference(minusi, minusi))
	ass.Equal(t, zero, class.Difference(minusone, minusone))
	ass.Equal(t, zero, class.Difference(zero, zero))
	ass.Equal(t, zero, class.Difference(one, one))
	ass.Equal(t, zero, class.Difference(i, i))
	ass.False(t, class.Difference(infinity, infinity).IsDefined())
	ass.False(t, class.Difference(undefined, undefined).IsDefined())

	//	z * r
	ass.Equal(t, minusi, class.Scaled(minusi, 1.0))
	ass.Equal(t, minushalf, class.Scaled(minusone, 0.5))
	ass.Equal(t, zero, class.Scaled(zero, 5.0))
	ass.Equal(t, half, class.Scaled(one, 0.5))
	ass.Equal(t, i, class.Scaled(i, 1.0))
	ass.Equal(t, infinity, class.Scaled(infinity, 5.0))
	ass.False(t, class.Scaled(undefined, 5.0).IsDefined())

	//	/z
	ass.Equal(t, infinity, class.Reciprocal(zero))
	ass.Equal(t, two, class.Reciprocal(half))
	ass.Equal(t, one, class.Reciprocal(one))
	ass.Equal(t, minushalf, class.Reciprocal(minustwo))
	ass.Equal(t, minusi, class.Reciprocal(i))
	ass.Equal(t, zero, class.Reciprocal(infinity))
	ass.False(t, class.Reciprocal(undefined).IsDefined())

	//	*z
	ass.Equal(t, zero, class.Conjugate(zero))
	ass.Equal(t, one, class.Conjugate(one))
	ass.Equal(t, minusi, class.Conjugate(i))
	ass.Equal(t, i, class.Conjugate(minusi))
	ass.False(t, class.Conjugate(undefined).IsDefined())

	//	z * zero => zero          {z != infinity}
	ass.Equal(t, zero, class.Product(zero, zero))
	ass.Equal(t, zero, class.Product(one, zero))
	ass.Equal(t, zero, class.Product(i, zero))
	ass.False(t, class.Product(infinity, zero).IsDefined())
	ass.False(t, class.Product(undefined, zero).IsDefined())

	//	z * one => z
	ass.Equal(t, zero, class.Product(zero, one))
	ass.Equal(t, one, class.Product(one, one))
	ass.Equal(t, i, class.Product(i, one))
	ass.Equal(t, infinity, class.Product(infinity, one))
	ass.False(t, class.Product(undefined, one).IsDefined())

	//	z * infinity => infinity  {z != zero}
	ass.False(t, class.Product(zero, infinity).IsDefined())
	ass.Equal(t, infinity, class.Product(one, infinity))
	ass.Equal(t, infinity, class.Product(i, infinity))
	ass.Equal(t, infinity, class.Product(infinity, infinity))

	//	zero / z => zero          {z != zero}
	ass.False(t, class.Quotient(zero, zero).IsDefined())
	ass.Equal(t, zero, class.Quotient(zero, one))
	ass.Equal(t, zero, class.Quotient(zero, i))
	ass.Equal(t, zero, class.Quotient(zero, infinity))
	ass.False(t, class.Quotient(zero, undefined).IsDefined())

	//	z / zero => infinity      {z != zero}
	ass.Equal(t, infinity, class.Quotient(one, zero))
	ass.Equal(t, infinity, class.Quotient(i, zero))
	ass.Equal(t, infinity, class.Quotient(infinity, zero))
	ass.False(t, class.Quotient(undefined, zero).IsDefined())

	//	z / infinity => zero      {z != infinity}
	ass.Equal(t, zero, class.Quotient(one, infinity))
	ass.Equal(t, zero, class.Quotient(i, infinity))
	ass.False(t, class.Quotient(infinity, infinity).IsDefined())
	ass.False(t, class.Quotient(undefined, infinity).IsDefined())

	//	infinity / z => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Quotient(infinity, zero))
	ass.Equal(t, infinity, class.Quotient(infinity, one))
	ass.Equal(t, infinity, class.Quotient(infinity, i))
	ass.False(t, class.Quotient(infinity, undefined).IsDefined())

	//	y / z
	ass.Equal(t, one, class.Quotient(one, one))
	ass.Equal(t, one, class.Quotient(i, i))
	ass.Equal(t, i, class.Quotient(i, one))
	ass.Equal(t, two, class.Quotient(one, half))
	ass.Equal(t, one, class.Quotient(half, half))

	//	z ^ zero => one           {by definition}
	ass.Equal(t, one, class.Power(minusi, zero))
	ass.Equal(t, one, class.Power(minusone, zero))
	ass.Equal(t, one, class.Power(zero, zero))
	ass.Equal(t, one, class.Power(one, zero))
	ass.Equal(t, one, class.Power(i, zero))
	ass.Equal(t, one, class.Power(infinity, zero))
	ass.False(t, class.Power(undefined, zero).IsDefined())

	//	zero ^ z => zero          {z != zero}
	ass.Equal(t, zero, class.Power(zero, one))
	ass.Equal(t, zero, class.Power(zero, i))
	ass.Equal(t, zero, class.Power(zero, infinity))
	ass.False(t, class.Power(zero, undefined).IsDefined())

	//	z ^ infinity => zero      {|z| < one}
	//	z ^ infinity => one       {|z| = one}
	//	z ^ infinity => infinity  {|z| > one}
	ass.Equal(t, infinity, class.Power(minustwo, infinity))
	ass.Equal(t, one, class.Power(minusi, infinity))
	ass.Equal(t, one, class.Power(minusone, infinity))
	ass.Equal(t, zero, class.Power(minushalf, infinity))
	ass.Equal(t, zero, class.Power(half, infinity))
	ass.Equal(t, one, class.Power(one, infinity))
	ass.Equal(t, one, class.Power(i, infinity))
	ass.Equal(t, infinity, class.Power(two, infinity))

	//	infinity ^ z => infinity  {z != zero}
	ass.Equal(t, one, class.Power(infinity, zero))
	ass.Equal(t, infinity, class.Power(infinity, one))
	ass.Equal(t, infinity, class.Power(infinity, i))
	ass.Equal(t, infinity, class.Power(infinity, infinity))
	ass.False(t, class.Power(infinity, undefined).IsDefined())

	//	one ^ z => one
	ass.Equal(t, one, class.Power(one, one))
	ass.Equal(t, one, class.Power(one, i))
	ass.Equal(t, one, class.Power(one, minusone))
	ass.Equal(t, one, class.Power(one, minusi))

	//	log(zero, z) => zero
	ass.False(t, class.Logarithm(zero, zero).IsDefined())
	ass.Equal(t, zero, class.Logarithm(zero, i))
	ass.Equal(t, zero, class.Logarithm(zero, one))
	ass.False(t, class.Logarithm(zero, infinity).IsDefined())
	ass.False(t, class.Logarithm(zero, undefined).IsDefined())

	//	log(one, z) => infinity
	ass.Equal(t, infinity, class.Logarithm(one, zero))
	ass.False(t, class.Logarithm(one, one).IsDefined())
	ass.Equal(t, infinity, class.Logarithm(one, infinity))
	ass.False(t, class.Logarithm(one, undefined).IsDefined())

	//	log(infinity, z) => zero
	ass.False(t, class.Logarithm(infinity, zero).IsDefined())
	ass.Equal(t, zero, class.Logarithm(infinity, one))
	ass.False(t, class.Logarithm(infinity, infinity).IsDefined())
	ass.False(t, class.Logarithm(infinity, undefined).IsDefined())
}

func TestZeroPercentages(t *tes.T) {
	var v = fra.Percentage(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestPositivePercentages(t *tes.T) {
	var v = fra.Percentage(25)
	ass.Equal(t, 0.25, v.AsIntrinsic())
	ass.Equal(t, 25.0, v.AsFloat())
}

func TestNegativePercentages(t *tes.T) {
	var v = fra.Percentage(-75)
	ass.Equal(t, -0.75, v.AsIntrinsic())
	ass.Equal(t, -75.0, v.AsFloat())
}

func TestStringPercentages(t *tes.T) {
	var v = fra.PercentageFromString("1.7%")
	//ass.Equal(t, -1.0, v.AsIntrinsic())
	//ass.Equal(t, -100.0, v.AsFloat())
	ass.Equal(t, "1.7%", v.AsString())
}

func TestBooleanProbabilities(t *tes.T) {
	var v1 = fra.ProbabilityFromBoolean(false)
	ass.Equal(t, 0.0, v1.AsFloat())

	var v2 = fra.ProbabilityFromBoolean(true)
	ass.Equal(t, 1.0, v2.AsFloat())
}

func TestZeroProbabilities(t *tes.T) {
	var v = fra.Probability(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestOneProbabilities(t *tes.T) {
	var v = fra.Probability(1.0)
	ass.Equal(t, 1.0, v.AsFloat())
}

func TestRandomProbability(t *tes.T) {
	fra.Random()
}

func TestStringProbabilities(t *tes.T) {
	var v = fra.ProbabilityFromString("p0")
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "p0", v.AsString())

	v = fra.ProbabilityFromString("p0.5")
	ass.Equal(t, 0.5, v.AsIntrinsic())
	ass.Equal(t, 0.5, v.AsFloat())
	ass.Equal(t, "p0.5", v.AsString())

	v = fra.ProbabilityFromString("p1")
	ass.Equal(t, 1.0, v.AsIntrinsic())
	ass.Equal(t, 1.0, v.AsFloat())
	ass.Equal(t, "p1", v.AsString())
}

func TestOtherProbabilities(t *tes.T) {
	var v1 = fra.Probability(0.25)
	ass.Equal(t, 0.25, v1.AsFloat())

	var v2 = fra.Probability(0.5)
	ass.Equal(t, 0.5, v2.AsFloat())

	var v3 = fra.Probability(0.75)
	ass.Equal(t, 0.75, v3.AsFloat())
}

func TestProbabilitieLibrary(t *tes.T) {
	var T = fra.Probability(0.75)
	var F = fra.Probability(0.25)
	var class = fra.ProbabilityClass()

	var andNot = class.And(class.Not(T), class.Not(T))
	var notIor = class.Not(class.Ior(T, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(T), class.Not(F))
	notIor = class.Not(class.Ior(T, F))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(T))
	notIor = class.Not(class.Ior(F, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(F))
	notIor = class.Not(class.Ior(F, F))
	ass.Equal(t, andNot, notIor)

	var san = class.And(T, class.Not(T))
	ass.Equal(t, san, class.San(T, T))

	san = class.And(T, class.Not(F))
	ass.Equal(t, san, class.San(T, F))

	san = class.And(F, class.Not(T))
	ass.Equal(t, san, class.San(F, T))

	san = class.And(F, class.Not(F))
	ass.Equal(t, san, class.San(F, F))

	var xor = class.Probability(class.San(T, T).AsFloat() + class.San(T, T).AsFloat())
	ass.Equal(t, xor, class.Xor(T, T))

	xor = class.Probability(class.San(T, F).AsFloat() + class.San(F, T).AsFloat())
	ass.Equal(t, xor, class.Xor(T, F))

	xor = class.Probability(class.San(F, T).AsFloat() + class.San(T, F).AsFloat())
	ass.Equal(t, xor, class.Xor(F, T))

	xor = class.Probability(class.San(F, F).AsFloat() + class.San(F, F).AsFloat())
	ass.Equal(t, xor, class.Xor(F, F))
}

func TestResource(t *tes.T) {
	var v = fra.Resource("https://craterdog.com/About.html")
	ass.Equal(t, "https://craterdog.com/About.html", v.AsIntrinsic())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPath(t *tes.T) {
	var v = fra.ResourceFromString("<https://craterdog.com/About.html>")
	ass.Equal(t, "<https://craterdog.com/About.html>", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithPath(t *tes.T) {
	var v = fra.ResourceFromString("<mailto:craterdog@google.com>")
	ass.Equal(t, "<mailto:craterdog@google.com>", v.AsString())
	ass.Equal(t, "mailto", v.GetScheme())
	ass.Equal(t, "", v.GetAuthority())
	ass.Equal(t, "", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQuery(t *tes.T) {
	var v = fra.ResourceFromString("<https://craterdog.com/?foo=bar;bar=baz>")
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz>", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndFragment(t *tes.T) {
	var v = fra.ResourceFromString("<https://craterdog.com/#Home>")
	ass.Equal(t, "<https://craterdog.com/#Home>", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQueryAndFragment(t *tes.T) {
	var v = fra.ResourceFromString("<https://craterdog.com/?foo=bar;bar=baz#Home>")
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz#Home>", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestSymbol(t *tes.T) {
	var foobar = "foo-bar"
	var v = fra.Symbol(foobar)
	ass.Equal(t, foobar, v.AsIntrinsic())
}

func TestSymbolFromString(t *tes.T) {
	var foobar = "$foo-bar"
	var v = fra.SymbolFromString(foobar)
	ass.Equal(t, foobar, v.AsString())
}

// STRING

func TestEmptyBinary(t *tes.T) {
	var binary = `'><'`
	var v = fra.BinaryFromString(binary)
	ass.Equal(t, binary, v.AsString())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestBinary(t *tes.T) {
	var b1 = `'>
    abcd1234
<'`
	var b2 = `'>
    abcd
<'`
	var v = fra.BinaryFromString(b1)
	ass.Equal(t, b1, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, int(v.GetSize()))
	ass.Equal(t, byte(0x69), v.GetValue(1))
	ass.Equal(t, byte(0xf8), v.GetValue(-1))
	ass.Equal(t, v.AsArray(), fra.Binary(v.AsArray()).AsArray())
	ass.Equal(t, b2, fra.BinaryFromSequence(v.GetValues(1, 3)).AsString())
	ass.Equal(t, 1, v.GetIndex(0x69))
}

func TestBinaryLibrary(t *tes.T) {
	var b1 = `'>
    abcd
<'`
	var b2 = `'>
    12345678
<'`
	var b3 = `'>
    abcd12345678
<'`
	var v1 = fra.BinaryFromString(b1)
	var v2 = fra.BinaryFromString(b2)
	var class = fra.BinaryClass()
	ass.Equal(t, b3, class.Concatenate(v1, v2).AsString())

	v1 = fra.Binary([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = fra.Binary([]byte{0x03, 0x00, 0x01, 0x02})
	var not = fra.Binary([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = fra.Binary([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = fra.Binary([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = fra.Binary([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = fra.Binary([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = fra.Binary([]byte{0x03, 0x00, 0x01, 0x00, 0x00})

	ass.Equal(t, not, class.Not(v1))
	ass.Equal(t, and, class.And(v1, v2))
	ass.Equal(t, sans, class.San(v1, v2))
	ass.Equal(t, or, class.Ior(v1, v2))
	ass.Equal(t, xor, class.Xor(v1, v2))
	ass.Equal(t, sans2, class.San(v2, v1))
}

func TestName(t *tes.T) {
	var v1 = fra.NameFromString("/bali-nebula/types/abstractions/5String")
	ass.Equal(t, "/bali-nebula/types/abstractions/5String", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 4, int(v1.GetSize()))
	ass.Equal(t, fra.Identifier("bali-nebula"), v1.GetValue(1))
	ass.Equal(t, fra.Identifier("5String"), v1.GetValue(-1))
	var v2 = fra.Name(v1.AsArray())
	ass.Equal(t, v1.AsString(), v2.AsString())
	var v3 = fra.NameFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, 1, v1.GetIndex("bali-nebula"))
	ass.Equal(t, "/bali-nebula/types", v3.AsString())
}

func TestNamesLibrary(t *tes.T) {
	var v1 = fra.NameFromString("/bali-nebula/types/abstractions")
	var v2 = fra.NameFromString("/String")
	ass.Equal(
		t,
		"/bali-nebula/types/abstractions/String",
		fra.NameClass().Concatenate(v1, v2).AsString(),
	)
}

const n0 = `"><"`

const n1 = `">
    abcd本
<"`

const n2 = `">
    1234
	\">
        This is an embedded narrative.
    <\"
<"`

const n3 = `">
    abcd本
    1234
	\">
        This is an embedded narrative.
    <\"
<"`

func TestEmptyNarrative(t *tes.T) {
	var v0 = fra.NarrativeFromString(n0)
	ass.Equal(t, n0, v0.AsString())
	ass.True(t, v0.IsEmpty())
	ass.Equal(t, 0, int(v0.GetSize()))
	ass.Equal(t, 0, len(v0.AsArray()))
}

func TestNarrative(t *tes.T) {
	var v1 = fra.NarrativeFromString(n1)
	ass.Equal(t, n1, v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 1, int(v1.GetSize()))

	var v3 = fra.NarrativeFromString(n3)
	ass.Equal(t, n3, v3.AsString())
	ass.False(t, v3.IsEmpty())
	ass.Equal(t, 5, int(v3.GetSize()))

	ass.Equal(t, n3, fra.Narrative(v3.AsArray()).AsString())
	ass.Equal(t, 5, len(v3.AsArray()))
}

func TestNarrativesLibrary(t *tes.T) {
	var v1 = fra.NarrativeFromString(n1)
	var v2 = fra.NarrativeFromString(n2)
	var v3 = fra.NarrativeClass().Concatenate(v1, v2)
	ass.Equal(t, v1.GetValue(1), v3.GetValue(1))
	ass.Equal(t, v2.GetValue(-1), v3.GetValue(-1))
	ass.Equal(t, n3, v3.AsString())
}

func TestNonePattern(t *tes.T) {
	var v = fra.PatternClass().None()
	ass.Equal(t, `none`, v.AsString())

	v = fra.PatternFromString(`none`)
	ass.Equal(t, `none`, v.AsString())
	ass.Equal(t, v, fra.PatternClass().None())

	var text = ""
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "anything at all..."
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestAnyPattern(t *tes.T) {
	var v = fra.PatternClass().Any()
	ass.Equal(t, `any`, v.AsString())

	v = fra.PatternFromString(`any`)
	ass.Equal(t, `any`, v.AsString())
	ass.Equal(t, v, fra.PatternClass().Any())

	var text = ""
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "anything at all..."
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestSomePattern(t *tes.T) {
	var v = fra.PatternFromString(`"c(.+t)"?`)
	ass.Equal(t, `"c(.+t)"?`, v.AsString())

	var text = "ct"
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "cat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "caaat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "cot"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))
}

func TestEmptyQuote(t *tes.T) {
	var v = fra.Quote([]fra.Character{})
	ass.Equal(t, []fra.Character{}, v.AsIntrinsic())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestQuote(t *tes.T) {
	var v = fra.QuoteFromString(`"abcd本1234"`)
	ass.Equal(t, `"abcd本1234"`, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 9, int(v.GetSize()))
	ass.Equal(t, 'a', rune(v.GetValue(1)))
	ass.Equal(t, '4', rune(v.GetValue(-1)))
	ass.Equal(t, `"d本1"`, fra.QuoteFromSequence(v.GetValues(4, 6)).AsString())
	ass.Equal(t, 8, v.GetIndex('3'))
}

func TestQuotesLibrary(t *tes.T) {
	var v1 = fra.QuoteFromString(`"abcd本"`)
	var v2 = fra.QuoteFromString(`"1234"`)
	ass.Equal(t, `"abcd本1234"`, fra.QuoteClass().Concatenate(v1, v2).AsString())
}

func TestStringTags(t *tes.T) {
	var size uint
	for size = 8; size < 33; size++ {
		var t1 = fra.TagWithSize(size)
		ass.Equal(t, len(t1.AsString()), 1+int(mat.Ceil(float64(size)*8.0/5.0)))
		var s1 = t1.AsString()
		var t2 = fra.TagFromString(s1)
		ass.Equal(t, t1, t2)
		var s2 = t2.AsString()
		ass.Equal(t, s1, s2)
		ass.Equal(t, t1.AsArray(), t2.AsArray())
	}
}

func TestVersion(t *tes.T) {
	var v1 = fra.VersionFromString("v1.2.3")
	ass.Equal(t, "v1.2.3", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 3, int(v1.GetSize()))
	ass.Equal(t, uint(1), v1.GetValue(1))
	ass.Equal(t, uint(3), v1.GetValue(-1))
	var v3 = fra.VersionFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, 2, v1.GetIndex(2))
	ass.Equal(t, "v1.2", v3.AsString())
}

func TestVersionsLibrary(t *tes.T) {
	var v1 = fra.Version([]uint{1})
	var v2 = fra.Version([]uint{2, 3})
	var class = fra.VersionClass()

	var v3 = class.Concatenate(v1, v2)
	ass.Equal(t, []uint{1, 2, 3}, v3.AsIntrinsic())
	ass.False(t, class.IsValidNextVersion(v1, v1))
	ass.Equal(t, "v2", class.GetNextVersion(v1, 0).AsString())
	ass.Equal(t, "v2", class.GetNextVersion(v1, 1).AsString())
	ass.True(t, class.IsValidNextVersion(v1, class.GetNextVersion(v1, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v1, 1), v1))
	ass.Equal(t, "v1.1", class.GetNextVersion(v1, 2).AsString())
	ass.True(t, class.IsValidNextVersion(v1, class.GetNextVersion(v1, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v1, 2), v1))
	ass.Equal(t, "v1.1", class.GetNextVersion(v1, 3).AsString())

	ass.False(t, class.IsValidNextVersion(v2, v2))
	ass.Equal(t, "v3", class.GetNextVersion(v2, 1).AsString())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 1), v2))
	ass.Equal(t, "v2.4", class.GetNextVersion(v2, 0).AsString())
	ass.Equal(t, "v2.4", class.GetNextVersion(v2, 2).AsString())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 2), v2))
	ass.Equal(t, "v2.3.1", class.GetNextVersion(v2, 3).AsString())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 3)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 3), v2))

	ass.False(t, class.IsValidNextVersion(v3, v3))
	ass.Equal(t, "v2", class.GetNextVersion(v3, 1).AsString())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 1), v3))
	ass.Equal(t, "v1.3", class.GetNextVersion(v3, 2).AsString())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 2), v3))
	ass.Equal(t, "v1.2.4", class.GetNextVersion(v3, 0).AsString())
	ass.Equal(t, "v1.2.4", class.GetNextVersion(v3, 3).AsString())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 3)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 3), v3))
	ass.Equal(t, "v1.2.3.1", class.GetNextVersion(v3, 4).AsString())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 4)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 4), v3))
}

func TestIntervalConstructors(t *tes.T) {
	var glyphs = fra.Interval[fra.GlyphLike](
		fra.Inclusive,
		fra.Glyph(65),
		fra.Glyph(70),
		fra.Inclusive,
	)
	ass.Equal(t, 6, int(glyphs.GetSize()))
	ass.Equal(t, "['A'..'F']", fmt.Sprintf("%v", glyphs))

	var durations = fra.IntervalClass[fra.DurationLike]().Interval(
		fra.Exclusive,
		fra.Duration(0),
		fra.DurationFromString("~P4W"),
		fra.Inclusive,
	)
	ass.Equal(t, 2419200000, int(durations.GetSize()))
	ass.Equal(t, "(..~P4W]", fmt.Sprintf("%v", durations))

	durations = fra.Interval[fra.DurationLike](
		fra.Exclusive,
		fra.DurationFromString("~P1D"),
		fra.DurationFromString("~P5D"),
		fra.Inclusive,
	)
	ass.Equal(t, 345600000, int(durations.GetSize()))
	ass.Equal(t, "(~P1D..~P5D]", fmt.Sprintf("%v", durations))

	var moments = fra.Interval[fra.MomentLike](
		fra.Exclusive,
		fra.MomentFromString("<2001-02-03T04:05:06>"),
		fra.MomentFromString("<2001-02-03T04:05:07>"),
		fra.Exclusive,
	)
	ass.Equal(t, 999, int(moments.GetSize()))
	ass.Equal(
		t,
		"(<2001-02-03T04:05:06>..<2001-02-03T04:05:07>)",
		fmt.Sprintf("%v", moments),
	)

	moments = fra.Interval[fra.MomentLike](
		fra.Exclusive,
		fra.Moment(mat.MinInt64),
		fra.Moment(mat.MaxInt64),
		fra.Exclusive,
	)
	ass.Equal(t, uint(0xfffffffffffffffe), moments.GetSize())
	ass.Equal(
		t,
		"(..)",
		fmt.Sprintf("%v", moments),
	)
}

func TestSpectrumConstructors(t *tes.T) {
	var names = fra.Spectrum[fra.NameLike](
		fra.Inclusive,
		fra.NameFromString("/nebula/classes/abstract"),
		fra.NameFromString("/nebula/types"),
		fra.Inclusive,
	)
	ass.Equal(
		t,
		"[/nebula/classes/abstract../nebula/types]",
		fmt.Sprintf("%v", names),
	)

	var quotes = fra.SpectrumClass[fra.QuoteLike]().Spectrum(
		fra.Inclusive,
		fra.QuoteFromString(`"A"`),
		fra.QuoteFromString(`"Fe"`),
		fra.Exclusive,
	)
	ass.Equal(
		t,
		`["A".."Fe")`,
		fmt.Sprintf("%v", quotes),
	)

	var versions = fra.Spectrum[fra.VersionLike](
		fra.Exclusive,
		fra.VersionFromString("v1.2.3"),
		fra.VersionFromString("v2"),
		fra.Exclusive,
	)
	ass.Equal(
		t,
		`(v1.2.3..v2)`,
		fmt.Sprintf("%v", versions),
	)
}

func TestContinuumConstructors(t *tes.T) {
	var numbers = fra.Continuum[fra.NumberLike](
		fra.Exclusive,
		fra.Number(-1.23),
		fra.Number(4.56),
		fra.Exclusive,
	)
	ass.Equal(t, "(-1.23..4.56)", fmt.Sprintf("%v", numbers))

	numbers = fra.Continuum[fra.NumberLike](
		fra.Exclusive,
		fra.NumberClass().Undefined(),
		fra.NumberClass().Zero(),
		fra.Inclusive,
	)
	ass.Equal(t, "(..0]", fmt.Sprintf("%v", numbers))

	numbers = fra.Continuum[fra.NumberLike](
		fra.Inclusive,
		fra.Number(1),
		fra.NumberClass().Undefined(),
		fra.Exclusive,
	)
	ass.Equal(t, "[1..)", fmt.Sprintf("%v", numbers))

	var probabilities = fra.ContinuumClass[fra.ProbabilityLike]().Continuum(
		fra.Inclusive,
		fra.Probability(0),
		fra.Probability(1),
		fra.Inclusive,
	)
	ass.Equal(t, "[p0..p1]", fmt.Sprintf("%v", probabilities))

	var angles = fra.Continuum[fra.AngleLike](
		fra.Inclusive,
		fra.Angle(0),
		fra.AngleClass().Tau(),
		fra.Exclusive,
	)
	ass.Equal(t, "[~0..~τ)", fmt.Sprintf("%v", angles))
}
