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

/*
Package "collections" declares a set of collection classes that maintain values
of a generic type:
  - Catalog (a sortable map of key-value associations)
  - List (a sortable list)
  - Queue (a blocking FIFO)
  - Set (an ordered set)
  - Stack (a LIFO)

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
package collections

import (
	age "github.com/craterdog/go-component-framework/v7/agents"
	str "github.com/craterdog/go-component-framework/v7/strings"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AssociationClassLike[K comparable, V any] is a class interface that declares
the complete set of class constructors, constants and functions that must be
supported by each concrete association-like class.

An association-like class captures the relationship between a generic typed
key-value pair.
*/
type AssociationClassLike[K comparable, V any] interface {
	// Constructor Methods
	Association(
		key K,
		value V,
	) AssociationLike[K, V]
}

/*
CatalogClassLike[K comparable, V any] is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete catalog-like class.

A catalog-like class maintains a sortable set of generic typed key-value
associations.  Unlike the intrinsic Go map data type, the order of the
associations in a catalog is the order in which they were added to the catalog.
A catalog can also be sorted using either the default "natural" ordering of the
keys or using a custom association ranking function.

The following class functions are also supported:

Extract() returns a new catalog containing only the associations that are in
the specified catalog that have the specified keys.  The associations in the
resulting catalog will be in the same order as the specified keys.

Merge() returns a new catalog containing all of the associations that are in
the specified Catalogs in the order that they appear in each catalog.  If a
key is present in both Catalogs, the value of the key from the second
catalog takes precedence.
*/
type CatalogClassLike[K comparable, V any] interface {
	// Constructor Methods
	Catalog() CatalogLike[K, V]
	CatalogFromArray(
		associations []AssociationLike[K, V],
	) CatalogLike[K, V]
	CatalogFromMap(
		associations map[K]V,
	) CatalogLike[K, V]
	CatalogFromSequence(
		associations str.Sequential[AssociationLike[K, V]],
	) CatalogLike[K, V]

	// Function Methods
	Extract(
		catalog CatalogLike[K, V],
		keys str.Sequential[K],
	) CatalogLike[K, V]
	Merge(
		first CatalogLike[K, V],
		second CatalogLike[K, V],
	) CatalogLike[K, V]
}

/*
ListClassLike[V any] is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete list-like class.

A list-like class maintains a dynamic sequence of values which can grow or
shrink as needed.  Each value is associated with an implicit positive integer
index.  It uses ORDINAL based indexing rather than the more common—and
nonsensical—ZERO based indexing scheme (see the description of what
this means in the str.Accessible[V] interface definition).

The following class functions are supported:

Concatenate() combines two lists into a new list containing all values in both
lists.  The order of the values in each list is preserved in the new list.
*/
type ListClassLike[V any] interface {
	// Constructor Methods
	List() ListLike[V]
	ListFromArray(
		values []V,
	) ListLike[V]
	ListFromSequence(
		values str.Sequential[V],
	) ListLike[V]

	// Function Methods
	Concatenate(
		first ListLike[V],
		second ListLike[V],
	) ListLike[V]
}

/*
QueueClassLike[V any] is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete queue-like class.

A queue-like class is generally used by multiple go-routines at the same time
and therefore enforces synchronized, first-in-first-out (FIFO) symantics for
generic typed values.  An optional queue capacity may be specified.  A request
to add a value to a queue will block when the queue has reached its maximum
capacity.  It will also block on attempts to remove a value when it is empty.
The default capacity for a queue-like class is 16 values.

The following class functions are supported:

Fork() connects the output of the specified input Queue with a number of new
output queues specified by the size parameter and returns a sequence of the new
output queues. Each value added to the input queue will be added automatically
to ALL of the output queues. This pattern is useful when a set of DIFFERENT
operations needs to occur for every value and each operation can be done in
parallel.

Split() connects the output of the specified input Queue with the number of
output queues specified by the size parameter and returns a sequence of the new
output queues. Each value added to the input queue will be added automatically
to ONE of the output queues. This pattern is useful when a SINGLE operation
needs to occur for each value and the operation can be done on the values in
parallel.  The results can then be consolidated later on using the Join()
function.

Join() connects the outputs of the specified sequence of input queues with a new
output queue returns the new output queue. Each value removed from each input
queue will automatically be added to the output queue.  This pattern is useful
when the results of the processing with a Split() function need to be
consolidated into a single queue.
*/
type QueueClassLike[V any] interface {
	// Constructor Methods
	Queue() QueueLike[V]
	QueueWithCapacity(
		capacity uti.Cardinal,
	) QueueLike[V]
	QueueFromArray(
		values []V,
	) QueueLike[V]
	QueueFromSequence(
		values str.Sequential[V],
	) QueueLike[V]

	// Function Methods
	Fork(
		group Synchronized,
		input QueueLike[V],
		size uti.Cardinal,
	) str.Sequential[QueueLike[V]]
	Split(
		group Synchronized,
		input QueueLike[V],
		size uti.Cardinal,
	) str.Sequential[QueueLike[V]]
	Join(
		group Synchronized,
		inputs str.Sequential[QueueLike[V]],
	) QueueLike[V]
}

/*
SetClassLike[V any] is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete set-like class.

A set-like class maintains an ordered sequence of unique generic typed
values—which can grow or shrink as needed.  The order of the values is
determined by a configurable collator agent.

The following class functions are supported:

And() returns a new set containing the values that are both of the specified
sets.

Ior() returns a new set containing the values that are in either of the specified
sets.

San() returns a new set containing the values that are in the first specified
set but not in the second specified set.

Xor() returns a new set containing the values that are in the first specified
set or the second specified set but not both.
*/
type SetClassLike[V any] interface {
	// Constructor Methods
	Set() SetLike[V]
	SetWithCollator(
		collator age.CollatorLike[V],
	) SetLike[V]
	SetFromArray(
		values []V,
	) SetLike[V]
	SetFromSequence(
		values str.Sequential[V],
	) SetLike[V]

	// Function Methods
	And(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	Ior(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	San(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	Xor(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
}

/*
StackClassLike[V any] is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete stack-like class.

A stack-like class supports last-in-first-out (LIFO) symantics for generic
typed values.  An optional stack capacity may be specified.  The default
capacity is 16 values.  Adding a value to a full stack will cause an error.
*/
type StackClassLike[V any] interface {
	// Constructor Methods
	Stack() StackLike[V]
	StackWithCapacity(
		capacity uti.Cardinal,
	) StackLike[V]
	StackFromArray(
		values []V,
	) StackLike[V]
	StackFromSequence(
		values str.Sequential[V],
	) StackLike[V]
}

// INSTANCE DECLARATIONS

/*
AssociationLike[K comparable, V any] is an instance interface that declares
the complete set of principal, attribute and aspect methods that must be
supported by each instance of a concrete association-like class.
*/
type AssociationLike[K comparable, V any] interface {
	// Principal Methods
	GetClass() AssociationClassLike[K, V]

	// Attribute Methods
	GetKey() K
	GetValue() V
	SetValue(
		value V,
	)
}

/*
CatalogLike[K comparable, V any] is an instance interface that declares
the complete set of principal, attribute and aspect methods that must be
supported by each instance of a concrete catalog-like class.
*/
type CatalogLike[K comparable, V any] interface {
	// Principal Methods
	GetClass() CatalogClassLike[K, V]

	// Aspect Interfaces
	Associative[K, V]
	str.Sequential[AssociationLike[K, V]]
	Sortable[AssociationLike[K, V]]
}

/*
ListLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete list-like class.
*/
type ListLike[V any] interface {
	// Principal Methods
	GetClass() ListClassLike[V]

	// Aspect Interfaces
	str.Accessible[V]
	Malleable[V]
	str.Searchable[V]
	str.Sequential[V]
	Sortable[V]
	Updatable[V]
}

/*
QueueLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete queue-like class.
*/
type QueueLike[V any] interface {
	// Principal Methods
	GetClass() QueueClassLike[V]

	// Attribute Methods
	GetCapacity() uti.Cardinal

	// Aspect Interfaces
	Fifo[V]
	str.Sequential[V]
}

/*
SetLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete set-like class.
*/
type SetLike[V any] interface {
	// Principal Methods
	GetClass() SetClassLike[V]

	// Attribute Methods
	GetCollator() age.CollatorLike[V]

	// Aspect Interfaces
	str.Accessible[V]
	Elastic[V]
	str.Searchable[V]
	str.Sequential[V]
}

/*
StackLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete stack-like class.
*/
type StackLike[V any] interface {
	// Principal Methods
	GetClass() StackClassLike[V]

	// Attribute Methods
	GetCapacity() uti.Cardinal

	// Aspect Interfaces
	Lifo[V]
	str.Sequential[V]
}

// ASPECT DECLARATIONS

/*
Associative[K comparable, V any] is an aspect interface that declares a set of
method signatures that must be supported by each instance of an associative
concrete class.

An associative class maintains a sequence of generic typed key-value
associations.
*/
type Associative[K comparable, V any] interface {
	AsMap() map[K]V
	GetValue(
		key K,
	) V
	SetValue(
		key K,
		value V,
	)
	GetKeys() str.Sequential[K]
	GetValues(
		keys str.Sequential[K],
	) str.Sequential[V]
	RemoveValue(
		key K,
	) V
	RemoveValues(
		keys str.Sequential[K],
	) str.Sequential[V]
	RemoveAll()
}

/*
Elastic[V any] is an aspect interface that declares a set of method signatures
that must be supported by each instance of an elastic concrete class.
*/
type Elastic[V any] interface {
	AddValue(
		value V,
	)
	AddValues(
		values str.Sequential[V],
	)
	RemoveValue(
		value V,
	)
	RemoveValues(
		values str.Sequential[V],
	)
	RemoveAll()
}

/*
Fifo[V any] is an aspect interface that declares a set of method signatures
that must be supported by each instance of a synchronized first-in-first-out
channel concrete class.
*/
type Fifo[V any] interface {
	AddValue(
		value V,
	)
	RemoveFirst() (
		first V,
		ok bool,
	)
	RemoveAll()
	CloseChannel()
}

/*
Lifo[V any] is an aspect interface that declares a set of method signatures
that must be supported by each instance of a last-in-first-out class.
*/
type Lifo[V any] interface {
	AddValue(
		value V,
	)
	GetLast() V
	RemoveLast() V
	RemoveAll()
}

/*
Malleable[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a malleable concrete
class.
*/
type Malleable[V any] interface {
	InsertValue(
		slot age.Slot,
		value V,
	)
	InsertValues(
		slot age.Slot,
		values str.Sequential[V],
	)
	AppendValue(
		value V,
	)
	AppendValues(
		values str.Sequential[V],
	)
	RemoveValue(
		index uti.Index,
	) V
	RemoveValues(
		first uti.Index,
		last uti.Index,
	) str.Sequential[V]
	RemoveAll()
}

/*
Sortable[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of a sortable concrete
class.

A sortable class allows its sequence of values to be sorted using a specific
sorting algorithm.
*/
type Sortable[V any] interface {
	SortValues()
	SortValuesWithRanker(
		ranker age.RankingFunction[V],
	)
	ReverseValues()
	ShuffleValues()
}

/*
Synchronized is an aspect interface that declares a set of method signatures
that must be supported by each instance of a synchronized concrete class.
*/
type Synchronized interface {
	Add(
		delta int,
	)
	Wait()
	Done()
}

/*
Updatable[V any] is an aspect interface that declares a set of method
signatures that must be supported by each instance of an updatable concrete
class.
*/
type Updatable[V any] interface {
	SetValue(
		index uti.Index,
		value V,
	)
	SetValues(
		index uti.Index,
		values str.Sequential[V],
	)
}
