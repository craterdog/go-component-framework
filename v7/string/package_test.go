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

package string_test

import (
	age "github.com/craterdog/go-component-framework/v7/agent"
	str "github.com/craterdog/go-component-framework/v7/string"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

var BinaryClass = str.BinaryClass()

func TestEmptyBinary(t *tes.T) {
	var binary = `'> <'`
	var v = BinaryClass.BinaryFromString(binary)
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
	var v = BinaryClass.BinaryFromString(b1)
	ass.Equal(t, b1, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, int(v.GetSize()))
	ass.Equal(t, byte(0x69), v.GetValue(1))
	ass.Equal(t, byte(0xf8), v.GetValue(-1))
	ass.Equal(t, v.AsArray(), BinaryClass.Binary(v.AsArray()).AsArray())
	ass.Equal(t, b2, BinaryClass.BinaryFromSequence(v.GetValues(1, 3)).AsString())
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
	var v1 = BinaryClass.BinaryFromString(b1)
	var v2 = BinaryClass.BinaryFromString(b2)
	ass.Equal(t, b3, BinaryClass.Concatenate(v1, v2).AsString())

	v1 = BinaryClass.Binary([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = BinaryClass.Binary([]byte{0x03, 0x00, 0x01, 0x02})
	var not = BinaryClass.Binary([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = BinaryClass.Binary([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = BinaryClass.Binary([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = BinaryClass.Binary([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = BinaryClass.Binary([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = BinaryClass.Binary([]byte{0x03, 0x00, 0x01, 0x00, 0x00})
	ass.Equal(t, not, BinaryClass.Not(v1))
	ass.Equal(t, and, BinaryClass.And(v1, v2))
	ass.Equal(t, sans, BinaryClass.San(v1, v2))
	ass.Equal(t, or, BinaryClass.Ior(v1, v2))
	ass.Equal(t, xor, BinaryClass.Xor(v1, v2))
	ass.Equal(t, sans2, BinaryClass.San(v2, v1))
}

var BytecodeClass = str.BytecodeClass()

func TestBytecode(t *tes.T) {
	var bytecode = `'abcd'`
	var v = BytecodeClass.BytecodeFromString(bytecode)
	ass.Equal(t, bytecode, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 1, int(v.GetSize()))
	ass.Equal(t, v.AsArray(), BytecodeClass.Bytecode(v.AsArray()).AsArray())

	bytecode = `'abcd 1234'`
	v = BytecodeClass.BytecodeFromString(bytecode)
	ass.Equal(t, bytecode, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 2, int(v.GetSize()))
	ass.Equal(t, v.AsArray(), BytecodeClass.Bytecode(v.AsArray()).AsArray())
}

var NameClass = str.NameClass()

func TestName(t *tes.T) {
	var v1 = NameClass.NameFromString("/bali-nebula/types/abstractions/String")
	ass.Equal(t, "/bali-nebula/types/abstractions/String", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 4, int(v1.GetSize()))
	ass.Equal(t, str.Identifier("bali-nebula"), v1.GetValue(1))
	ass.Equal(t, str.Identifier("String"), v1.GetValue(-1))
	var v2 = NameClass.Name(v1.AsArray())
	ass.Equal(t, v1.AsString(), v2.AsString())
	var v3 = NameClass.NameFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, "/bali-nebula/types", v3.AsString())
}

func TestNamesLibrary(t *tes.T) {
	var v1 = NameClass.NameFromString("/bali-nebula/types/abstractions")
	var v2 = NameClass.NameFromString("/String")
	ass.Equal(
		t,
		"/bali-nebula/types/abstractions/String",
		NameClass.Concatenate(v1, v2).AsString(),
	)
}

var NarrativeClass = str.NarrativeClass()

const n0 = `"><"`

const n1 = `">
    abcd本
<"`

const n2 = `">
    1234
<"`

const n3 = `">
    abcd本
    1234
<"`

func TestEmptyNarrative(t *tes.T) {
	var v0 = NarrativeClass.NarrativeFromString(n0)
	ass.Equal(t, n0, v0.AsString())
	ass.True(t, v0.IsEmpty())
	ass.Equal(t, 0, int(v0.GetSize()))
	ass.Equal(t, 0, len(v0.AsArray()))
}

func TestNarrative(t *tes.T) {
	var v1 = NarrativeClass.NarrativeFromString(n1)
	ass.Equal(t, n1, v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 1, int(v1.GetSize()))

	var v3 = NarrativeClass.NarrativeFromString(n3)
	ass.Equal(t, n3, v3.AsString())
	ass.False(t, v3.IsEmpty())
	ass.Equal(t, 2, int(v3.GetSize()))

	ass.Equal(t, n3, NarrativeClass.Narrative(v3.AsArray()).AsString())
	ass.Equal(t, 2, len(v3.AsArray()))
}

func TestNarrativesLibrary(t *tes.T) {
	var v1 = NarrativeClass.NarrativeFromString(n1)
	var v2 = NarrativeClass.NarrativeFromString(n2)
	var v3 = NarrativeClass.Concatenate(v1, v2)
	ass.Equal(t, v1.GetValue(1), v3.GetValue(1))
	ass.Equal(t, v2.GetValue(-1), v3.GetValue(-1))
	ass.Equal(t, n3, v3.AsString())
}

var PatternClass = str.PatternClass()

func TestNonePattern(t *tes.T) {
	var v = PatternClass.PatternFromString(`none`)
	ass.Equal(t, `none`, v.AsString())

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
	var v = PatternClass.PatternFromString(`any`)
	ass.Equal(t, `any`, v.AsString())

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
	var v = PatternClass.PatternFromString(`"c(.+t)"?`)
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

var QuoteClass = str.QuoteClass()

func TestEmptyQuote(t *tes.T) {
	var v = QuoteClass.Quote([]str.Character{})
	ass.Equal(t, []str.Character{}, v.AsIntrinsic())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestQuote(t *tes.T) {
	var v = QuoteClass.QuoteFromString(`"abcd本1234"`)
	ass.Equal(t, `"abcd本1234"`, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 9, int(v.GetSize()))
	ass.Equal(t, 'a', rune(v.GetValue(1)))
	ass.Equal(t, '4', rune(v.GetValue(-1)))
	ass.Equal(t, `"d本1"`, QuoteClass.QuoteFromSequence(v.GetValues(4, 6)).AsString())
}

func TestQuotesLibrary(t *tes.T) {
	var v1 = QuoteClass.QuoteFromString(`"abcd本"`)
	var v2 = QuoteClass.QuoteFromString(`"1234"`)
	ass.Equal(t, `"abcd本1234"`, QuoteClass.Concatenate(v1, v2).AsString())
}

var SymbolClass = str.SymbolClass()

func TestSymbol(t *tes.T) {
	var foobar = "$foo-bar"
	var v = SymbolClass.SymbolFromString(foobar)
	ass.Equal(t, foobar, v.AsString())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 7, int(v.GetSize()))
	ass.Equal(t, []str.Character("foo-bar"), v.AsArray())
}

var TagClass = str.TagClass()

func TestStringTags(t *tes.T) {
	var size age.Cardinal
	for size = 8; size < 33; size++ {
		var t1 = TagClass.TagWithSize(size)
		ass.Equal(t, len(t1.AsString()), 1+int(mat.Ceil(float64(size)*8.0/5.0)))
		var s1 = t1.AsString()
		var t2 = TagClass.TagFromString(s1)
		ass.Equal(t, t1, t2)
		var s2 = t2.AsString()
		ass.Equal(t, s1, s2)
		ass.Equal(t, t1.AsArray(), t2.AsArray())
	}
}

var VersionClass = str.VersionClass()

func TestVersion(t *tes.T) {
	var v1 = VersionClass.VersionFromString("v1.2.3")
	ass.Equal(t, "v1.2.3", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 3, int(v1.GetSize()))
	ass.Equal(t, age.Ordinal(1), v1.GetValue(1))
	ass.Equal(t, age.Ordinal(3), v1.GetValue(-1))
	var v3 = VersionClass.VersionFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, "v1.2", v3.AsString())
}

func TestVersionsLibrary(t *tes.T) {
	var v1 = VersionClass.Version([]age.Ordinal{1})
	var v2 = VersionClass.Version([]age.Ordinal{2, 3})
	var v3 = VersionClass.Concatenate(v1, v2)
	ass.Equal(t, []age.Ordinal{1, 2, 3}, v3.AsIntrinsic())

	ass.False(t, VersionClass.IsValidNextVersion(v1, v1))
	ass.Equal(t, "v2", VersionClass.GetNextVersion(v1, 0).AsString())
	ass.Equal(t, "v2", VersionClass.GetNextVersion(v1, 1).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v1, VersionClass.GetNextVersion(v1, 1)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v1, 1), v1))
	ass.Equal(t, "v1.1", VersionClass.GetNextVersion(v1, 2).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v1, VersionClass.GetNextVersion(v1, 2)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v1, 2), v1))
	ass.Equal(t, "v1.1", VersionClass.GetNextVersion(v1, 3).AsString())

	ass.False(t, VersionClass.IsValidNextVersion(v2, v2))
	ass.Equal(t, "v3", VersionClass.GetNextVersion(v2, 1).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v2, VersionClass.GetNextVersion(v2, 1)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v2, 1), v2))
	ass.Equal(t, "v2.4", VersionClass.GetNextVersion(v2, 0).AsString())
	ass.Equal(t, "v2.4", VersionClass.GetNextVersion(v2, 2).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v2, VersionClass.GetNextVersion(v2, 2)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v2, 2), v2))
	ass.Equal(t, "v2.3.1", VersionClass.GetNextVersion(v2, 3).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v2, VersionClass.GetNextVersion(v2, 3)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v2, 3), v2))

	ass.False(t, VersionClass.IsValidNextVersion(v3, v3))
	ass.Equal(t, "v2", VersionClass.GetNextVersion(v3, 1).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v3, VersionClass.GetNextVersion(v3, 1)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v3, 1), v3))
	ass.Equal(t, "v1.3", VersionClass.GetNextVersion(v3, 2).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v3, VersionClass.GetNextVersion(v3, 2)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v3, 2), v3))
	ass.Equal(t, "v1.2.4", VersionClass.GetNextVersion(v3, 0).AsString())
	ass.Equal(t, "v1.2.4", VersionClass.GetNextVersion(v3, 3).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v3, VersionClass.GetNextVersion(v3, 3)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v3, 3), v3))
	ass.Equal(t, "v1.2.3.1", VersionClass.GetNextVersion(v3, 4).AsString())
	ass.True(t, VersionClass.IsValidNextVersion(v3, VersionClass.GetNextVersion(v3, 4)))
	ass.False(t, VersionClass.IsValidNextVersion(VersionClass.GetNextVersion(v3, 4), v3))
}
