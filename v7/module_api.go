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
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│             This "module_api.go" file was automatically generated.           │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-component-framework/wiki
*/
package module

import (
	age "github.com/craterdog/go-component-framework/v7/agent"
	col "github.com/craterdog/go-component-framework/v7/collection"
	ele "github.com/craterdog/go-component-framework/v7/element"
	ser "github.com/craterdog/go-component-framework/v7/series"
	uri "net/url"
)

// TYPE ALIASES

// Element

type (
	Units = ele.Units
)

const (
	Degrees  = ele.Degrees
	Radians  = ele.Radians
	Gradians = ele.Gradians
)

type (
	AngleClassLike       = ele.AngleClassLike
	BooleanClassLike     = ele.BooleanClassLike
	CitationClassLike    = ele.CitationClassLike
	DurationClassLike    = ele.DurationClassLike
	GlyphClassLike       = ele.GlyphClassLike
	MomentClassLike      = ele.MomentClassLike
	NumberClassLike      = ele.NumberClassLike
	PercentageClassLike  = ele.PercentageClassLike
	ProbabilityClassLike = ele.ProbabilityClassLike
	ResourceClassLike    = ele.ResourceClassLike
)

type (
	AngleLike       = ele.AngleLike
	BooleanLike     = ele.BooleanLike
	CitationLike    = ele.CitationLike
	DurationLike    = ele.DurationLike
	GlyphLike       = ele.GlyphLike
	MomentLike      = ele.MomentLike
	NumberLike      = ele.NumberLike
	PercentageLike  = ele.PercentageLike
	ProbabilityLike = ele.ProbabilityLike
	ResourceLike    = ele.ResourceLike
)

type (
	Continuous = ele.Continuous
	Discrete   = ele.Discrete
	Factored   = ele.Factored
	Polarized  = ele.Polarized
	Temporal   = ele.Temporal
)

// Series

type (
	Identifier  = ser.Identifier
	Instruction = ser.Instruction
	Line        = ser.Line
)

type (
	BinaryClassLike    = ser.BinaryClassLike
	BytecodeClassLike  = ser.BytecodeClassLike
	NameClassLike      = ser.NameClassLike
	NarrativeClassLike = ser.NarrativeClassLike
	PatternClassLike   = ser.PatternClassLike
	QuoteClassLike     = ser.QuoteClassLike
	SymbolClassLike    = ser.SymbolClassLike
	TagClassLike       = ser.TagClassLike
	VersionClassLike   = ser.VersionClassLike
)

type (
	BinaryLike    = ser.BinaryLike
	BytecodeLike  = ser.BytecodeLike
	NameLike      = ser.NameLike
	NarrativeLike = ser.NarrativeLike
	PatternLike   = ser.PatternLike
	QuoteLike     = ser.QuoteLike
	SymbolLike    = ser.SymbolLike
	TagLike       = ser.TagLike
	VersionLike   = ser.VersionLike
)

// Collection

type (
	Index = col.Index
)

type (
	AssociationClassLike[K comparable, V any] = col.AssociationClassLike[K, V]
	CatalogClassLike[K comparable, V any]     = col.CatalogClassLike[K, V]
	ListClassLike[V any]                      = col.ListClassLike[V]
	QueueClassLike[V any]                     = col.QueueClassLike[V]
	SetClassLike[V any]                       = col.SetClassLike[V]
	StackClassLike[V any]                     = col.StackClassLike[V]
)

type (
	AssociationLike[K comparable, V any] = col.AssociationLike[K, V]
	CatalogLike[K comparable, V any]     = col.CatalogLike[K, V]
	ListLike[V any]                      = col.ListLike[V]
	QueueLike[V any]                     = col.QueueLike[V]
	SetLike[V any]                       = col.SetLike[V]
	StackLike[V any]                     = col.StackLike[V]
)

type (
	Accessible[V any]                = col.Accessible[V]
	Associative[K comparable, V any] = col.Associative[K, V]
	Elastic[V any]                   = col.Elastic[V]
	Fifo[V any]                      = col.Fifo[V]
	Lifo[V any]                      = col.Lifo[V]
	Malleable[V any]                 = col.Malleable[V]
	Searchable[V any]                = col.Searchable[V]
	Sequential[V any]                = col.Sequential[V]
	Sortable[V any]                  = col.Sortable[V]
	Synchronized                     = col.Synchronized
	Updatable[V any]                 = col.Updatable[V]
)

// Agent

type (
	Cardinal    = age.Cardinal
	Event       = age.Event
	Ordinal     = age.Ordinal
	Rank        = age.Rank
	Slot        = age.Slot
	State       = age.State
	Transitions = age.Transitions
)

const (
	LesserRank  = age.LesserRank
	EqualRank   = age.EqualRank
	GreaterRank = age.GreaterRank
)

type (
	RankingFunction[V any] = age.RankingFunction[V]
)

type (
	CollatorClassLike[V any] = age.CollatorClassLike[V]
	ControllerClassLike      = age.ControllerClassLike
	EncoderClassLike         = age.EncoderClassLike
	GeneratorClassLike       = age.GeneratorClassLike
	IteratorClassLike[V any] = age.IteratorClassLike[V]
	SorterClassLike[V any]   = age.SorterClassLike[V]
)

type (
	CollatorLike[V any] = age.CollatorLike[V]
	ControllerLike      = age.ControllerLike
	EncoderLike         = age.EncoderLike
	GeneratorLike       = age.GeneratorLike
	IteratorLike[V any] = age.IteratorLike[V]
	SorterLike[V any]   = age.SorterLike[V]
)

// CLASS ACCESSORS

// Element

func AngleClass() AngleClassLike {
	return ele.AngleClass()
}

func Angle(
	radians float64,
) AngleLike {
	return AngleClass().Angle(
		radians,
	)
}

func AngleFromString(
	string_ string,
) AngleLike {
	return AngleClass().AngleFromString(
		string_,
	)
}

func BooleanClass() BooleanClassLike {
	return ele.BooleanClass()
}

func Boolean(
	boolean bool,
) BooleanLike {
	return BooleanClass().Boolean(
		boolean,
	)
}

func BooleanFromString(
	string_ string,
) BooleanLike {
	return BooleanClass().BooleanFromString(
		string_,
	)
}

func CitationClass() CitationClassLike {
	return ele.CitationClass()
}

func Citation(
	string_ string,
) CitationLike {
	return CitationClass().Citation(
		string_,
	)
}

func DurationClass() DurationClassLike {
	return ele.DurationClass()
}

func Duration(
	milliseconds int,
) DurationLike {
	return DurationClass().Duration(
		milliseconds,
	)
}

func DurationFromString(
	string_ string,
) DurationLike {
	return DurationClass().DurationFromString(
		string_,
	)
}

func GlyphClass() GlyphClassLike {
	return ele.GlyphClass()
}

func Glyph(
	rune_ rune,
) GlyphLike {
	return GlyphClass().Glyph(
		rune_,
	)
}

func GlyphFromInteger(
	integer int,
) GlyphLike {
	return GlyphClass().GlyphFromInteger(
		integer,
	)
}

func GlyphFromString(
	string_ string,
) GlyphLike {
	return GlyphClass().GlyphFromString(
		string_,
	)
}

func MomentClass() MomentClassLike {
	return ele.MomentClass()
}

func Moment(
	milliseconds int,
) MomentLike {
	return MomentClass().Moment(
		milliseconds,
	)
}

func MomentFromString(
	string_ string,
) MomentLike {
	return MomentClass().MomentFromString(
		string_,
	)
}

func NumberClass() NumberClassLike {
	return ele.NumberClass()
}

func Number(
	complex_ complex128,
) NumberLike {
	return NumberClass().Number(
		complex_,
	)
}

func NumberFromPolar(
	magnitude float64,
	phase float64,
) NumberLike {
	return NumberClass().NumberFromPolar(
		magnitude,
		phase,
	)
}

func NumberFromRectangular(
	real_ float64,
	imaginary float64,
) NumberLike {
	return NumberClass().NumberFromRectangular(
		real_,
		imaginary,
	)
}

func NumberFromString(
	string_ string,
) NumberLike {
	return NumberClass().NumberFromString(
		string_,
	)
}

func PercentageClass() PercentageClassLike {
	return ele.PercentageClass()
}

func Percentage(
	float float64,
) PercentageLike {
	return PercentageClass().Percentage(
		float,
	)
}

func PercentageFromInteger(
	integer int,
) PercentageLike {
	return PercentageClass().PercentageFromInteger(
		integer,
	)
}

func PercentageFromString(
	string_ string,
) PercentageLike {
	return PercentageClass().PercentageFromString(
		string_,
	)
}

func ProbabilityClass() ProbabilityClassLike {
	return ele.ProbabilityClass()
}

func Probability(
	float float64,
) ProbabilityLike {
	return ProbabilityClass().Probability(
		float,
	)
}

func ProbabilityFromBoolean(
	boolean bool,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromBoolean(
		boolean,
	)
}

func ProbabilityFromString(
	string_ string,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromString(
		string_,
	)
}

func ResourceClass() ResourceClassLike {
	return ele.ResourceClass()
}

func Resource(
	string_ string,
) ResourceLike {
	return ResourceClass().Resource(
		string_,
	)
}

func ResourceFromUri(
	url *uri.URL,
) ResourceLike {
	return ResourceClass().ResourceFromUri(
		url,
	)
}

// Series

func BinaryClass() BinaryClassLike {
	return ser.BinaryClass()
}

func Binary(
	bytes []byte,
) BinaryLike {
	return BinaryClass().Binary(
		bytes,
	)
}

func BinaryFromSequence(
	sequence col.Sequential[byte],
) BinaryLike {
	return BinaryClass().BinaryFromSequence(
		sequence,
	)
}

func BinaryFromString(
	string_ string,
) BinaryLike {
	return BinaryClass().BinaryFromString(
		string_,
	)
}

func BytecodeClass() BytecodeClassLike {
	return ser.BytecodeClass()
}

func Bytecode(
	instructions []ser.Instruction,
) BytecodeLike {
	return BytecodeClass().Bytecode(
		instructions,
	)
}

func BytecodeFromSequence(
	sequence col.Sequential[ser.Instruction],
) BytecodeLike {
	return BytecodeClass().BytecodeFromSequence(
		sequence,
	)
}

func BytecodeFromString(
	string_ string,
) BytecodeLike {
	return BytecodeClass().BytecodeFromString(
		string_,
	)
}

func NameClass() NameClassLike {
	return ser.NameClass()
}

func Name(
	identifiers []ser.Identifier,
) NameLike {
	return NameClass().Name(
		identifiers,
	)
}

func NameFromSequence(
	sequence col.Sequential[ser.Identifier],
) NameLike {
	return NameClass().NameFromSequence(
		sequence,
	)
}

func NameFromString(
	string_ string,
) NameLike {
	return NameClass().NameFromString(
		string_,
	)
}

func NarrativeClass() NarrativeClassLike {
	return ser.NarrativeClass()
}

func Narrative(
	lines []ser.Line,
) NarrativeLike {
	return NarrativeClass().Narrative(
		lines,
	)
}

func NarrativeFromSequence(
	sequence col.Sequential[ser.Line],
) NarrativeLike {
	return NarrativeClass().NarrativeFromSequence(
		sequence,
	)
}

func NarrativeFromString(
	string_ string,
) NarrativeLike {
	return NarrativeClass().NarrativeFromString(
		string_,
	)
}

func PatternClass() PatternClassLike {
	return ser.PatternClass()
}

func Pattern(
	string_ string,
) PatternLike {
	return PatternClass().Pattern(
		string_,
	)
}

func PatternFromSequence(
	sequence col.Sequential[rune],
) PatternLike {
	return PatternClass().PatternFromSequence(
		sequence,
	)
}

func PatternFromString(
	string_ string,
) PatternLike {
	return PatternClass().PatternFromString(
		string_,
	)
}

func QuoteClass() QuoteClassLike {
	return ser.QuoteClass()
}

func Quote(
	string_ string,
) QuoteLike {
	return QuoteClass().Quote(
		string_,
	)
}

func QuoteFromSequence(
	sequence col.Sequential[rune],
) QuoteLike {
	return QuoteClass().QuoteFromSequence(
		sequence,
	)
}

func QuoteFromString(
	string_ string,
) QuoteLike {
	return QuoteClass().QuoteFromString(
		string_,
	)
}

func SymbolClass() SymbolClassLike {
	return ser.SymbolClass()
}

func Symbol(
	string_ string,
) SymbolLike {
	return SymbolClass().Symbol(
		string_,
	)
}

func SymbolFromSequence(
	sequence col.Sequential[rune],
) SymbolLike {
	return SymbolClass().SymbolFromSequence(
		sequence,
	)
}

func SymbolFromString(
	string_ string,
) SymbolLike {
	return SymbolClass().SymbolFromString(
		string_,
	)
}

func TagClass() TagClassLike {
	return ser.TagClass()
}

func Tag(
	bytes []byte,
) TagLike {
	return TagClass().Tag(
		bytes,
	)
}

func TagWithSize(
	size Cardinal,
) TagLike {
	return TagClass().TagWithSize(
		size,
	)
}

func TagFromSequence(
	sequence col.Sequential[byte],
) TagLike {
	return TagClass().TagFromSequence(
		sequence,
	)
}

func TagFromString(
	string_ string,
) TagLike {
	return TagClass().TagFromString(
		string_,
	)
}

func VersionClass() VersionClassLike {
	return ser.VersionClass()
}

func Version(
	ordinals []Ordinal,
) VersionLike {
	return VersionClass().Version(
		ordinals,
	)
}

func VersionFromSequence(
	sequence col.Sequential[Ordinal],
) VersionLike {
	return VersionClass().VersionFromSequence(
		sequence,
	)
}

func VersionFromString(
	string_ string,
) VersionLike {
	return VersionClass().VersionFromString(
		string_,
	)
}

// Collection

func AssociationClass[K comparable, V any]() AssociationClassLike[K, V] {
	return col.AssociationClass[K, V]()
}

func Association[K comparable, V any](
	key K,
	value V,
) AssociationLike[K, V] {
	return AssociationClass[K, V]().Association(
		key,
		value,
	)
}

func CatalogClass[K comparable, V any]() CatalogClassLike[K, V] {
	return col.CatalogClass[K, V]()
}

func Catalog[K comparable, V any]() CatalogLike[K, V] {
	return CatalogClass[K, V]().Catalog()
}

func CatalogFromArray[K comparable, V any](
	associations []col.AssociationLike[K, V],
) CatalogLike[K, V] {
	return CatalogClass[K, V]().CatalogFromArray(
		associations,
	)
}

func CatalogFromMap[K comparable, V any](
	associations map[K]V,
) CatalogLike[K, V] {
	return CatalogClass[K, V]().CatalogFromMap(
		associations,
	)
}

func CatalogFromSequence[K comparable, V any](
	associations col.Sequential[col.AssociationLike[K, V]],
) CatalogLike[K, V] {
	return CatalogClass[K, V]().CatalogFromSequence(
		associations,
	)
}

func ListClass[V any]() ListClassLike[V] {
	return col.ListClass[V]()
}

func List[V any]() ListLike[V] {
	return ListClass[V]().List()
}

func ListFromArray[V any](
	values []V,
) ListLike[V] {
	return ListClass[V]().ListFromArray(
		values,
	)
}

func ListFromSequence[V any](
	values col.Sequential[V],
) ListLike[V] {
	return ListClass[V]().ListFromSequence(
		values,
	)
}

func QueueClass[V any]() QueueClassLike[V] {
	return col.QueueClass[V]()
}

func Queue[V any]() QueueLike[V] {
	return QueueClass[V]().Queue()
}

func QueueWithCapacity[V any](
	capacity Cardinal,
) QueueLike[V] {
	return QueueClass[V]().QueueWithCapacity(
		capacity,
	)
}

func QueueFromArray[V any](
	values []V,
) QueueLike[V] {
	return QueueClass[V]().QueueFromArray(
		values,
	)
}

func QueueFromSequence[V any](
	values col.Sequential[V],
) QueueLike[V] {
	return QueueClass[V]().QueueFromSequence(
		values,
	)
}

func SetClass[V any]() SetClassLike[V] {
	return col.SetClass[V]()
}

func Set[V any]() SetLike[V] {
	return SetClass[V]().Set()
}

func SetWithCollator[V any](
	collator age.CollatorLike[V],
) SetLike[V] {
	return SetClass[V]().SetWithCollator(
		collator,
	)
}

func SetFromArray[V any](
	values []V,
) SetLike[V] {
	return SetClass[V]().SetFromArray(
		values,
	)
}

func SetFromSequence[V any](
	values col.Sequential[V],
) SetLike[V] {
	return SetClass[V]().SetFromSequence(
		values,
	)
}

func StackClass[V any]() StackClassLike[V] {
	return col.StackClass[V]()
}

func Stack[V any]() StackLike[V] {
	return StackClass[V]().Stack()
}

func StackWithCapacity[V any](
	capacity Cardinal,
) StackLike[V] {
	return StackClass[V]().StackWithCapacity(
		capacity,
	)
}

func StackFromArray[V any](
	values []V,
) StackLike[V] {
	return StackClass[V]().StackFromArray(
		values,
	)
}

func StackFromSequence[V any](
	values col.Sequential[V],
) StackLike[V] {
	return StackClass[V]().StackFromSequence(
		values,
	)
}

// Agent

func CollatorClass[V any]() CollatorClassLike[V] {
	return age.CollatorClass[V]()
}

func Collator[V any]() CollatorLike[V] {
	return CollatorClass[V]().Collator()
}

func CollatorWithMaximumDepth[V any](
	maximumDepth age.Cardinal,
) CollatorLike[V] {
	return CollatorClass[V]().CollatorWithMaximumDepth(
		maximumDepth,
	)
}

func ControllerClass() ControllerClassLike {
	return age.ControllerClass()
}

func Controller(
	events []age.Event,
	transitions map[State]age.Transitions,
) ControllerLike {
	return ControllerClass().Controller(
		events,
		transitions,
	)
}

func EncoderClass() EncoderClassLike {
	return age.EncoderClass()
}

func Encoder() EncoderLike {
	return EncoderClass().Encoder()
}

func GeneratorClass() GeneratorClassLike {
	return age.GeneratorClass()
}

func Generator() GeneratorLike {
	return GeneratorClass().Generator()
}

func IteratorClass[V any]() IteratorClassLike[V] {
	return age.IteratorClass[V]()
}

func Iterator[V any](
	array []V,
) IteratorLike[V] {
	return IteratorClass[V]().Iterator(
		array,
	)
}

func SorterClass[V any]() SorterClassLike[V] {
	return age.SorterClass[V]()
}

func Sorter[V any]() SorterLike[V] {
	return SorterClass[V]().Sorter()
}

func SorterWithRanker[V any](
	ranker age.RankingFunction[V],
) SorterLike[V] {
	return SorterClass[V]().SorterWithRanker(
		ranker,
	)
}

// GLOBAL FUNCTIONS
