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
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
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
	age "github.com/craterdog/go-component-framework/v7/agents"
	col "github.com/craterdog/go-component-framework/v7/collections"
	ele "github.com/craterdog/go-component-framework/v7/elements"
	ran "github.com/craterdog/go-component-framework/v7/ranges"
	str "github.com/craterdog/go-component-framework/v7/strings"
	uti "github.com/craterdog/go-missing-utilities/v7"
	uri "net/url"
)

// TYPE ALIASES

// Agents

type (
	Event       = age.Event
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

// Collections

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
	Associative[K comparable, V any] = col.Associative[K, V]
	Elastic[V any]                   = col.Elastic[V]
	Fifo[V any]                      = col.Fifo[V]
	Lifo[V any]                      = col.Lifo[V]
	Malleable[V any]                 = col.Malleable[V]
	Sortable[V any]                  = col.Sortable[V]
	Synchronized                     = col.Synchronized
	Updatable[V any]                 = col.Updatable[V]
)

// Elements

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
	DurationClassLike    = ele.DurationClassLike
	GlyphClassLike       = ele.GlyphClassLike
	MomentClassLike      = ele.MomentClassLike
	NumberClassLike      = ele.NumberClassLike
	PercentageClassLike  = ele.PercentageClassLike
	ProbabilityClassLike = ele.ProbabilityClassLike
	ResourceClassLike    = ele.ResourceClassLike
	SymbolClassLike      = ele.SymbolClassLike
)

type (
	AngleLike       = ele.AngleLike
	BooleanLike     = ele.BooleanLike
	DurationLike    = ele.DurationLike
	GlyphLike       = ele.GlyphLike
	MomentLike      = ele.MomentLike
	NumberLike      = ele.NumberLike
	PercentageLike  = ele.PercentageLike
	ProbabilityLike = ele.ProbabilityLike
	ResourceLike    = ele.ResourceLike
	SymbolLike      = ele.SymbolLike
)

type (
	Continuous = ele.Continuous
	Discrete   = ele.Discrete
	Factored   = ele.Factored
	Polarized  = ele.Polarized
	Temporal   = ele.Temporal
)

// Ranges

type (
	Bracket = ran.Bracket
)

const (
	Inclusive = ran.Inclusive
	Exclusive = ran.Exclusive
)

type (
	ContinuumClassLike[V ele.Continuous] = ran.ContinuumClassLike[V]
	IntervalClassLike[V ele.Discrete]    = ran.IntervalClassLike[V]
	SpectrumClassLike[V str.Spectral[V]] = ran.SpectrumClassLike[V]
)

type (
	ContinuumLike[V ele.Continuous] = ran.ContinuumLike[V]
	IntervalLike[V ele.Discrete]    = ran.IntervalLike[V]
	SpectrumLike[V str.Spectral[V]] = ran.SpectrumLike[V]
)

type (
	Bounded[V any] = ran.Bounded[V]
)

// Strings

type (
	Identifier = str.Identifier
	Line       = str.Line
	Character  = str.Character
)

type (
	BinaryClassLike    = str.BinaryClassLike
	NameClassLike      = str.NameClassLike
	NarrativeClassLike = str.NarrativeClassLike
	PatternClassLike   = str.PatternClassLike
	QuoteClassLike     = str.QuoteClassLike
	TagClassLike       = str.TagClassLike
	VersionClassLike   = str.VersionClassLike
)

type (
	BinaryLike    = str.BinaryLike
	NameLike      = str.NameLike
	NarrativeLike = str.NarrativeLike
	PatternLike   = str.PatternLike
	QuoteLike     = str.QuoteLike
	TagLike       = str.TagLike
	VersionLike   = str.VersionLike
)

type (
	Accessible[V any] = str.Accessible[V]
	Searchable[V any] = str.Searchable[V]
	Sequential[V any] = str.Sequential[V]
	Spectral[V any]   = str.Spectral[V]
)

// CLASS ACCESSORS

// Agents

func CollatorClass[V any]() CollatorClassLike[V] {
	return age.CollatorClass[V]()
}

func Collator[V any]() CollatorLike[V] {
	return CollatorClass[V]().Collator()
}

func CollatorWithMaximumDepth[V any](
	maximumDepth uti.Cardinal,
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
	initialState age.State,
) ControllerLike {
	return ControllerClass().Controller(
		events,
		transitions,
		initialState,
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

// Collections

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
	associations str.Sequential[col.AssociationLike[K, V]],
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
	values str.Sequential[V],
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
	capacity uti.Cardinal,
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
	values str.Sequential[V],
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
	values str.Sequential[V],
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
	capacity uti.Cardinal,
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
	values str.Sequential[V],
) StackLike[V] {
	return StackClass[V]().StackFromSequence(
		values,
	)
}

// Elements

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
	source string,
) AngleLike {
	return AngleClass().AngleFromString(
		source,
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
	source string,
) BooleanLike {
	return BooleanClass().BooleanFromString(
		source,
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
	source string,
) DurationLike {
	return DurationClass().DurationFromString(
		source,
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
	source string,
) GlyphLike {
	return GlyphClass().GlyphFromString(
		source,
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
	source string,
) MomentLike {
	return MomentClass().MomentFromString(
		source,
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

func NumberFromInteger(
	integer int,
) NumberLike {
	return NumberClass().NumberFromInteger(
		integer,
	)
}

func NumberFromFloat(
	float float64,
) NumberLike {
	return NumberClass().NumberFromFloat(
		float,
	)
}

func NumberFromString(
	source string,
) NumberLike {
	return NumberClass().NumberFromString(
		source,
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
	source string,
) PercentageLike {
	return PercentageClass().PercentageFromString(
		source,
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
	source string,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromString(
		source,
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

func ResourceFromString(
	source string,
) ResourceLike {
	return ResourceClass().ResourceFromString(
		source,
	)
}

func ResourceFromUri(
	url *uri.URL,
) ResourceLike {
	return ResourceClass().ResourceFromUri(
		url,
	)
}

func SymbolClass() SymbolClassLike {
	return ele.SymbolClass()
}

func Symbol(
	string_ string,
) SymbolLike {
	return SymbolClass().Symbol(
		string_,
	)
}

func SymbolFromString(
	source string,
) SymbolLike {
	return SymbolClass().SymbolFromString(
		source,
	)
}

// Ranges

func ContinuumClass[V ele.Continuous]() ContinuumClassLike[V] {
	return ran.ContinuumClass[V]()
}

func Continuum[V ele.Continuous](
	left ran.Bracket,
	minimum V,
	maximum V,
	right ran.Bracket,
) ContinuumLike[V] {
	return ContinuumClass[V]().Continuum(
		left,
		minimum,
		maximum,
		right,
	)
}

func IntervalClass[V ele.Discrete]() IntervalClassLike[V] {
	return ran.IntervalClass[V]()
}

func Interval[V ele.Discrete](
	left ran.Bracket,
	minimum V,
	maximum V,
	right ran.Bracket,
) IntervalLike[V] {
	return IntervalClass[V]().Interval(
		left,
		minimum,
		maximum,
		right,
	)
}

func SpectrumClass[V str.Spectral[V]]() SpectrumClassLike[V] {
	return ran.SpectrumClass[V]()
}

func Spectrum[V str.Spectral[V]](
	left ran.Bracket,
	minimum V,
	maximum V,
	right ran.Bracket,
) SpectrumLike[V] {
	return SpectrumClass[V]().Spectrum(
		left,
		minimum,
		maximum,
		right,
	)
}

// Strings

func BinaryClass() BinaryClassLike {
	return str.BinaryClass()
}

func Binary(
	bytes []byte,
) BinaryLike {
	return BinaryClass().Binary(
		bytes,
	)
}

func BinaryFromSequence(
	sequence str.Sequential[byte],
) BinaryLike {
	return BinaryClass().BinaryFromSequence(
		sequence,
	)
}

func BinaryFromString(
	source string,
) BinaryLike {
	return BinaryClass().BinaryFromString(
		source,
	)
}

func NameClass() NameClassLike {
	return str.NameClass()
}

func Name(
	identifiers []str.Identifier,
) NameLike {
	return NameClass().Name(
		identifiers,
	)
}

func NameFromSequence(
	sequence str.Sequential[str.Identifier],
) NameLike {
	return NameClass().NameFromSequence(
		sequence,
	)
}

func NameFromString(
	source string,
) NameLike {
	return NameClass().NameFromString(
		source,
	)
}

func NarrativeClass() NarrativeClassLike {
	return str.NarrativeClass()
}

func Narrative(
	lines []str.Line,
) NarrativeLike {
	return NarrativeClass().Narrative(
		lines,
	)
}

func NarrativeFromSequence(
	sequence str.Sequential[str.Line],
) NarrativeLike {
	return NarrativeClass().NarrativeFromSequence(
		sequence,
	)
}

func NarrativeFromString(
	source string,
) NarrativeLike {
	return NarrativeClass().NarrativeFromString(
		source,
	)
}

func PatternClass() PatternClassLike {
	return str.PatternClass()
}

func Pattern(
	characters []str.Character,
) PatternLike {
	return PatternClass().Pattern(
		characters,
	)
}

func PatternFromSequence(
	sequence str.Sequential[str.Character],
) PatternLike {
	return PatternClass().PatternFromSequence(
		sequence,
	)
}

func PatternFromString(
	source string,
) PatternLike {
	return PatternClass().PatternFromString(
		source,
	)
}

func QuoteClass() QuoteClassLike {
	return str.QuoteClass()
}

func Quote(
	characters []str.Character,
) QuoteLike {
	return QuoteClass().Quote(
		characters,
	)
}

func QuoteFromSequence(
	sequence str.Sequential[str.Character],
) QuoteLike {
	return QuoteClass().QuoteFromSequence(
		sequence,
	)
}

func QuoteFromString(
	source string,
) QuoteLike {
	return QuoteClass().QuoteFromString(
		source,
	)
}

func TagClass() TagClassLike {
	return str.TagClass()
}

func Tag(
	bytes []byte,
) TagLike {
	return TagClass().Tag(
		bytes,
	)
}

func TagWithSize(
	size uti.Cardinal,
) TagLike {
	return TagClass().TagWithSize(
		size,
	)
}

func TagFromSequence(
	sequence str.Sequential[byte],
) TagLike {
	return TagClass().TagFromSequence(
		sequence,
	)
}

func TagFromString(
	source string,
) TagLike {
	return TagClass().TagFromString(
		source,
	)
}

func VersionClass() VersionClassLike {
	return str.VersionClass()
}

func Version(
	ordinals []uti.Ordinal,
) VersionLike {
	return VersionClass().Version(
		ordinals,
	)
}

func VersionFromSequence(
	sequence str.Sequential[uti.Ordinal],
) VersionLike {
	return VersionClass().VersionFromSequence(
		sequence,
	)
}

func VersionFromString(
	source string,
) VersionLike {
	return VersionClass().VersionFromString(
		source,
	)
}

// GLOBAL FUNCTIONS

func Now() MomentLike {
	return ele.MomentClass().Now()
}

func Random() ProbabilityLike {
	return ele.ProbabilityClass().Random()
}
