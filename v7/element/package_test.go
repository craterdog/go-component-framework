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

package element_test

import (
	ele "github.com/craterdog/go-component-framework/v7/element"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	cmp "math/cmplx"
	tes "testing"
)

var AngleClass = ele.AngleClass()

func TestZeroAngles(t *tes.T) {
	var v = AngleClass.Angle(0)
	ass.Equal(t, 0.0, v.GetIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())

	v = AngleClass.AngleFromString("~0")
	ass.Equal(t, "~0", v.AsString())
}

func TestPositiveAngles(t *tes.T) {
	var v = AngleClass.Angle(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())

	v = AngleClass.AngleFromString("~π")
	ass.Equal(t, "~π", v.AsString())
}

func TestNegativeAngles(t *tes.T) {
	var v = AngleClass.Angle(-mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())

	v = AngleClass.Angle(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v.AsFloat())
}

func TestAnglesLibrary(t *tes.T) {
	var v0 = AngleClass.Zero()
	var v1 = AngleClass.Angle(mat.Pi * 0.25)
	var v2 = AngleClass.Angle(mat.Pi * 0.5)
	var v3 = AngleClass.Angle(mat.Pi * 0.75)
	var v4 = AngleClass.Pi()
	var v5 = AngleClass.Angle(mat.Pi * 1.25)
	var v6 = AngleClass.Angle(mat.Pi * 1.5)
	var v7 = AngleClass.Angle(mat.Pi * 1.75)
	var v8 = AngleClass.Tau()

	ass.Equal(t, v4, AngleClass.Inverse(v0))
	ass.Equal(t, v5, AngleClass.Inverse(v1))
	ass.Equal(t, v6, AngleClass.Inverse(v2))
	ass.Equal(t, v7, AngleClass.Inverse(v3))
	ass.Equal(t, v0, AngleClass.Inverse(v4))
	ass.Equal(t, v4, AngleClass.Inverse(v8))

	ass.Equal(t, v1, AngleClass.Sum(v0, v1))
	ass.Equal(t, v0, AngleClass.Difference(v1, v1))
	ass.Equal(t, v3, AngleClass.Sum(v1, v2))
	ass.Equal(t, v1, AngleClass.Difference(v3, v2))
	ass.Equal(t, v5, AngleClass.Sum(v2, v3))
	ass.Equal(t, v2, AngleClass.Difference(v5, v3))
	ass.Equal(t, v7, AngleClass.Sum(v3, v4))
	ass.Equal(t, v3, AngleClass.Difference(v7, v4))
	ass.Equal(t, v1, AngleClass.Sum(v8, v1))
	ass.Equal(t, v0, AngleClass.Difference(v8, v8))

	ass.Equal(t, v3, AngleClass.Scaled(v1, 3.0))
	ass.Equal(t, v0, AngleClass.Scaled(v4, 2.0))
	ass.Equal(t, v4, AngleClass.Scaled(v4, -1.0))
	ass.Equal(t, v0, AngleClass.Scaled(v8, 1.0))

	ass.Equal(t, v0, AngleClass.ArcCosine(AngleClass.Cosine(v0)))
	ass.Equal(t, v1, AngleClass.ArcCosine(AngleClass.Cosine(v1)))
	ass.Equal(t, v2, AngleClass.ArcCosine(AngleClass.Cosine(v2)))
	ass.Equal(t, v3, AngleClass.ArcCosine(AngleClass.Cosine(v3)))
	ass.Equal(t, v4, AngleClass.ArcCosine(AngleClass.Cosine(v4)))
	ass.Equal(t, v0, AngleClass.ArcCosine(AngleClass.Cosine(v8)))

	ass.Equal(t, v0, AngleClass.ArcSine(AngleClass.Sine(v0)))
	ass.Equal(t, v1, AngleClass.ArcSine(AngleClass.Sine(v1)))
	ass.Equal(t, v2, AngleClass.ArcSine(AngleClass.Sine(v2)))
	ass.Equal(t, v6, AngleClass.ArcSine(AngleClass.Sine(v6)))
	ass.Equal(t, v7, AngleClass.ArcSine(AngleClass.Sine(v7)))
	ass.Equal(t, v0, AngleClass.ArcSine(AngleClass.Sine(v8)))

	ass.Equal(t, v0, AngleClass.ArcTangent(AngleClass.Cosine(v0), AngleClass.Sine(v0)))
	ass.Equal(t, v1, AngleClass.ArcTangent(AngleClass.Cosine(v1), AngleClass.Sine(v1)))
	ass.Equal(t, v2, AngleClass.ArcTangent(AngleClass.Cosine(v2), AngleClass.Sine(v2)))
	ass.Equal(t, v3, AngleClass.ArcTangent(AngleClass.Cosine(v3), AngleClass.Sine(v3)))
	ass.Equal(t, v4, AngleClass.ArcTangent(AngleClass.Cosine(v4), AngleClass.Sine(v4)))
	ass.Equal(t, v5, AngleClass.ArcTangent(AngleClass.Cosine(v5), AngleClass.Sine(v5)))
	ass.Equal(t, v0, AngleClass.ArcTangent(AngleClass.Cosine(v8), AngleClass.Sine(v8)))
}

var BooleanClass = ele.BooleanClass()

func TestFalseBooleans(t *tes.T) {
	ass.False(t, BooleanClass.Minimum().GetIntrinsic())
	ass.False(t, BooleanClass.False().GetIntrinsic())
	var v = BooleanClass.Boolean(false)
	ass.False(t, v.AsBoolean())
	v = BooleanClass.BooleanFromString("false")
	ass.Equal(t, "false", v.AsString())
}

func TestTrueBooleans(t *tes.T) {
	ass.True(t, BooleanClass.Maximum().GetIntrinsic())
	ass.True(t, BooleanClass.True().GetIntrinsic())
	var v = BooleanClass.Boolean(true)
	ass.True(t, v.AsBoolean())
	v = BooleanClass.BooleanFromString("true")
	ass.Equal(t, "true", v.AsString())
}

func TestBooleansLibrary(t *tes.T) {
	var T = BooleanClass.Boolean(true)
	var F = BooleanClass.Boolean(false)

	var andNot = BooleanClass.And(BooleanClass.Not(T), BooleanClass.Not(T))
	var notIor = BooleanClass.Not(BooleanClass.Ior(T, T))
	ass.Equal(t, andNot, notIor)

	andNot = BooleanClass.And(BooleanClass.Not(T), BooleanClass.Not(F))
	notIor = BooleanClass.Not(BooleanClass.Ior(T, F))
	ass.Equal(t, andNot, notIor)

	andNot = BooleanClass.And(BooleanClass.Not(F), BooleanClass.Not(T))
	notIor = BooleanClass.Not(BooleanClass.Ior(F, T))
	ass.Equal(t, andNot, notIor)

	andNot = BooleanClass.And(BooleanClass.Not(F), BooleanClass.Not(F))
	notIor = BooleanClass.Not(BooleanClass.Ior(F, F))
	ass.Equal(t, andNot, notIor)

	var san = BooleanClass.And(T, BooleanClass.Not(T))
	ass.Equal(t, san, BooleanClass.San(T, T))

	san = BooleanClass.And(T, BooleanClass.Not(F))
	ass.Equal(t, san, BooleanClass.San(T, F))

	san = BooleanClass.And(F, BooleanClass.Not(T))
	ass.Equal(t, san, BooleanClass.San(F, T))

	san = BooleanClass.And(F, BooleanClass.Not(F))
	ass.Equal(t, san, BooleanClass.San(F, F))

	var xor = BooleanClass.Ior(BooleanClass.San(T, T), BooleanClass.San(T, T))
	ass.Equal(t, xor, BooleanClass.Xor(T, T))

	xor = BooleanClass.Ior(BooleanClass.San(T, F), BooleanClass.San(F, T))
	ass.Equal(t, xor, BooleanClass.Xor(T, F))

	xor = BooleanClass.Ior(BooleanClass.San(F, T), BooleanClass.San(T, F))
	ass.Equal(t, xor, BooleanClass.Xor(F, T))

	xor = BooleanClass.Ior(BooleanClass.San(F, F), BooleanClass.San(F, F))
	ass.Equal(t, xor, BooleanClass.Xor(F, F))
}

var CitationClass = ele.CitationClass()

func TestCitation(t *tes.T) {
	var v1 = CitationClass.Citation("/bali/types/abstractions/String@v1.2.3")
	ass.Equal(t, "/bali/types/abstractions/String@v1.2.3", v1.GetIntrinsic())
	ass.Equal(t, "/bali/types/abstractions/String@v1.2.3", v1.AsString())
	ass.Equal(t, "/bali/types/abstractions/String", v1.GetName())
	ass.Equal(t, "v1.2.3", v1.GetVersion())
}

var DurationClass = ele.DurationClass()

func TestZeroDurations(t *tes.T) {
	var v = DurationClass.Duration(0)
	ass.Equal(t, 0, v.AsInteger())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0, v.GetIntrinsic())
	ass.Equal(t, 0.0, v.AsMilliseconds())
	ass.Equal(t, 0.0, v.AsSeconds())
	ass.Equal(t, 0.0, v.AsMinutes())
	ass.Equal(t, 0.0, v.AsHours())
	ass.Equal(t, 0.0, v.AsDays())
	ass.Equal(t, 0.0, v.AsWeeks())
	ass.Equal(t, 0.0, v.AsMonths())
	ass.Equal(t, 0.0, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 0, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 0, v.GetDays())
	ass.Equal(t, 0, v.GetWeeks())
	ass.Equal(t, 0, v.GetMonths())
	ass.Equal(t, 0, v.GetYears())
}

func TestStringDurations(t *tes.T) {
	var duration = DurationClass.DurationFromString("~-P1Y2M3DT4H5M6S")
	ass.Equal(t, "~-P1Y2M3DT4H5M6S", duration.AsString())
	duration = DurationClass.DurationFromString("~P0W")
	ass.Equal(t, "~P0W", duration.AsString())
}

func TestPositiveDurations(t *tes.T) {
	var v = DurationClass.Duration(60000)
	ass.Equal(t, "~PT1M", v.AsString())
	ass.Equal(t, 60000, v.AsInteger())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 60000, v.GetIntrinsic())
	ass.Equal(t, 60000.0, v.AsMilliseconds())
	ass.Equal(t, 60.0, v.AsSeconds())
	ass.Equal(t, 1.0, v.AsMinutes())
	ass.Equal(t, 0.016666666666666666, v.AsHours())
	ass.Equal(t, 0.0006944444444444445, v.AsDays())
	ass.Equal(t, 9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, 2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, 1.9013243104086858e-06, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 1, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 0, v.GetDays())
	ass.Equal(t, 0, v.GetWeeks())
	ass.Equal(t, 0, v.GetMonths())
	ass.Equal(t, 0, v.GetYears())
}

func TestNegativeDurations(t *tes.T) {
	var v = DurationClass.Duration(-60000)
	ass.Equal(t, -60000, v.AsInteger())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -60000.0, v.AsMilliseconds())
	ass.Equal(t, -60.0, v.AsSeconds())
	ass.Equal(t, -1.0, v.AsMinutes())
	ass.Equal(t, -0.016666666666666666, v.AsHours())
	ass.Equal(t, -0.0006944444444444445, v.AsDays())
	ass.Equal(t, -9.92063492063492e-05, v.AsWeeks())
	ass.Equal(t, -2.2815891724904232e-05, v.AsMonths())
	ass.Equal(t, -1.9013243104086858e-06, v.AsYears())
	ass.Equal(t, 0, v.GetMilliseconds())
	ass.Equal(t, 0, v.GetSeconds())
	ass.Equal(t, 1, v.GetMinutes())
	ass.Equal(t, 0, v.GetHours())
	ass.Equal(t, 0, v.GetDays())
	ass.Equal(t, 0, v.GetWeeks())
	ass.Equal(t, 0, v.GetMonths())
	ass.Equal(t, 0, v.GetYears())
}

var GlyphClass = ele.GlyphClass()

func TestGlyphs(t *tes.T) {
	var v = GlyphClass.GlyphFromString("'''")
	ass.Equal(t, "'''", v.AsString())

	v = GlyphClass.Glyph('a')
	ass.Equal(t, "'a'", v.AsString())

	v = GlyphClass.Glyph('"')
	ass.Equal(t, `'"'`, v.AsString())

	v = GlyphClass.Glyph('😊')
	ass.Equal(t, "'😊'", v.AsString())

	v = GlyphClass.Glyph('界')
	ass.Equal(t, "'界'", v.AsString())

	v = GlyphClass.Glyph('\'')
	ass.Equal(t, "'''", v.AsString())

	v = GlyphClass.Glyph('\\')
	ass.Equal(t, "'\\'", v.AsString())

	v = GlyphClass.Glyph('\n')
	ass.Equal(t, "'\n'", v.AsString())

	v = GlyphClass.Glyph('\t')
	ass.Equal(t, "'\t'", v.AsString())
}

var MomentClass = ele.MomentClass()

func TestIntegerMoments(t *tes.T) {
	var v = MomentClass.Moment(1238589296789)
	ass.Equal(t, 1238589296789, v.GetIntrinsic())
	ass.Equal(t, 1238589296789, v.AsInteger())
	ass.Equal(t, 1238589296789.0, v.AsMilliseconds())
	ass.Equal(t, 1238589296.789, v.AsSeconds())
	ass.Equal(t, 20643154.946483333, v.AsMinutes())
	ass.Equal(t, 344052.58244138886, v.AsHours())
	ass.Equal(t, 14335.524268391204, v.AsDays())
	ass.Equal(t, 2047.9320383416004, v.AsWeeks())
	ass.Equal(t, 470.9919881193849, v.AsMonths())
	ass.Equal(t, 39.24933234328208, v.AsYears())
	ass.Equal(t, 789, v.GetMilliseconds())
	ass.Equal(t, 56, v.GetSeconds())
	ass.Equal(t, 34, v.GetMinutes())
	ass.Equal(t, 12, v.GetHours())
	ass.Equal(t, 1, v.GetDays())
	ass.Equal(t, 14, v.GetWeeks())
	ass.Equal(t, 4, v.GetMonths())
	ass.Equal(t, 2009, v.GetYears())
}

func TestStringMoments(t *tes.T) {
	var v = MomentClass.MomentFromString("<-1-02-03T04:05:06.700>")
	ass.Equal(t, "<-1-02-03T04:05:06.700>", v.AsString())
}

func TestMomentsLibrary(t *tes.T) {
	var before = MomentClass.Now()
	var duration = DurationClass.Duration(12345)
	var after = MomentClass.Moment(before.AsInteger() + duration.AsInteger())

	ass.Equal(t, duration, MomentClass.Duration(before, after))
	ass.Equal(t, after, MomentClass.Later(before, duration))
	ass.Equal(t, before, MomentClass.Earlier(after, duration))
}

var NumberClass = ele.NumberClass()

func TestZero(t *tes.T) {
	var v = NumberClass.Number(0 + 0i)
	ass.Equal(t, 0+0i, v.GetIntrinsic())
	ass.True(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.False(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestInfinity(t *tes.T) {
	var v = NumberClass.Number(cmp.Inf())
	ass.Equal(t, cmp.Inf(), v.GetIntrinsic())
	ass.False(t, v.IsZero())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, mat.Inf(1), v.AsFloat())
	ass.Equal(t, mat.Inf(1), v.GetReal())
	ass.Equal(t, mat.Inf(1), v.GetImaginary())
}

func TestUndefined(t *tes.T) {
	var v = NumberClass.Number(cmp.NaN())
	ass.True(t, cmp.IsNaN(v.GetIntrinsic()))
	ass.False(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.True(t, v.IsUndefined())
	ass.False(t, v.IsNegative())
	ass.True(t, mat.IsNaN(v.AsFloat()))
	ass.True(t, mat.IsNaN(v.GetReal()))
	ass.True(t, mat.IsNaN(v.GetImaginary()))
}

func TestPositivePureReals(t *tes.T) {
	var v = NumberClass.Number(0.25)
	ass.Equal(t, 0.25+0i, v.GetIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.25, v.AsFloat())
	ass.Equal(t, 0.25, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestPositivePureImaginaries(t *tes.T) {
	var v = NumberClass.Number(0.25i)
	ass.Equal(t, 0+0.25i, v.GetIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.25, v.GetImaginary())
}

func TestNegativePureReals(t *tes.T) {
	var v = NumberClass.Number(-0.75)
	ass.Equal(t, -0.75+0i, v.GetIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -0.75, v.AsFloat())
	ass.Equal(t, -0.75, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestNegativePureImaginaries(t *tes.T) {
	var v = NumberClass.Number(-0.75i)
	ass.Equal(t, 0-0.75i, v.GetIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -0.75, v.GetImaginary())
}

func TestNumberFromPolar(t *tes.T) {
	var v = NumberClass.NumberFromPolar(1.0, mat.Pi)
	ass.Equal(t, -1.0+0i, v.GetIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetPhase())
}

func TestNumberFromString(t *tes.T) {
	var v = NumberClass.NumberFromString("1e^~πi")
	ass.Equal(t, -1.0+0i, v.GetIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, "-1", v.AsString())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetPhase())

	v = NumberClass.NumberFromString("-1.2-3.4i")
	ass.Equal(t, "-1.2-3.4i", v.AsString())
	ass.Equal(t, -1.2, v.GetReal())
	ass.Equal(t, -3.4, v.GetImaginary())

	v = NumberClass.NumberFromString("undefined")
	ass.Equal(t, "undefined", v.AsString())
	ass.True(t, v.IsUndefined())
	ass.False(t, v.HasMagnitude())

	v = NumberClass.NumberFromString("+infinity")
	ass.Equal(t, "∞", v.AsString())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = NumberClass.NumberFromString("infinity")
	ass.Equal(t, "∞", v.AsString())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = NumberClass.NumberFromString("∞")
	ass.Equal(t, "∞", v.AsString())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = NumberClass.NumberFromString("-∞")
	ass.Equal(t, "∞", v.AsString())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = NumberClass.NumberFromString("+1")
	ass.Equal(t, "1", v.AsString())
	ass.Equal(t, 1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, 0.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = NumberClass.NumberFromString("1")
	ass.Equal(t, "1", v.AsString())
	ass.Equal(t, 1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, 0.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = NumberClass.NumberFromString("-π")
	ass.Equal(t, "-π", v.AsString())
	ass.Equal(t, -mat.Pi, v.GetReal())
	ass.Equal(t, mat.Pi, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.True(t, v.IsNegative())

	v = NumberClass.NumberFromString("+1i")
	ass.Equal(t, "1i", v.AsString())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi/2.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = NumberClass.NumberFromString("1i")
	ass.Equal(t, "1i", v.AsString())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi/2.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = NumberClass.NumberFromString("-1i")
	ass.Equal(t, "-1i", v.AsString())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, -mat.Pi/2.0, v.GetPhase())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())
}

func TestNumberLibrary(t *tes.T) {
	var zero = NumberClass.Zero()
	var i = NumberClass.I()
	var minusi = NumberClass.Number(-1i)
	var half = NumberClass.Number(0.5)
	var minushalf = NumberClass.Number(-0.5)
	var one = NumberClass.One()
	var minusone = NumberClass.Number(-1)
	var two = NumberClass.Number(2.0)
	var minustwo = NumberClass.Number(-2.0)
	var infinity = NumberClass.Infinity()
	var undefined = NumberClass.Undefined()

	//	-z
	ass.Equal(t, zero, NumberClass.Inverse(zero))
	ass.Equal(t, minushalf, NumberClass.Inverse(half))
	ass.Equal(t, minusone, NumberClass.Inverse(one))
	ass.Equal(t, minusi, NumberClass.Inverse(i))
	ass.Equal(t, infinity, NumberClass.Inverse(infinity))
	ass.True(t, NumberClass.Inverse(undefined).IsUndefined())

	//	z + zero => z
	ass.Equal(t, minusi, NumberClass.Sum(minusi, zero))
	ass.Equal(t, minusone, NumberClass.Sum(minusone, zero))
	ass.Equal(t, zero, NumberClass.Sum(zero, zero))
	ass.Equal(t, one, NumberClass.Sum(one, zero))
	ass.Equal(t, i, NumberClass.Sum(i, zero))
	ass.Equal(t, infinity, NumberClass.Sum(infinity, zero))
	ass.True(t, NumberClass.Sum(undefined, zero).IsUndefined())

	//	z + infinity => infinity
	ass.Equal(t, infinity, NumberClass.Sum(minusi, infinity))
	ass.Equal(t, infinity, NumberClass.Sum(minusone, infinity))
	ass.Equal(t, infinity, NumberClass.Sum(zero, infinity))
	ass.Equal(t, infinity, NumberClass.Sum(one, infinity))
	ass.Equal(t, infinity, NumberClass.Sum(i, infinity))
	ass.Equal(t, infinity, NumberClass.Sum(infinity, infinity))
	ass.True(t, NumberClass.Sum(undefined, infinity).IsUndefined())

	//	z - infinity => infinity  {z != infinity}
	ass.Equal(t, infinity, NumberClass.Difference(minusi, infinity))
	ass.Equal(t, infinity, NumberClass.Difference(minusone, infinity))
	ass.Equal(t, infinity, NumberClass.Difference(zero, infinity))
	ass.Equal(t, infinity, NumberClass.Difference(one, infinity))
	ass.Equal(t, infinity, NumberClass.Difference(i, infinity))
	ass.True(t, NumberClass.Difference(infinity, infinity).IsUndefined())
	ass.True(t, NumberClass.Difference(undefined, infinity).IsUndefined())

	//	infinity - z => infinity  {z != infinity}
	ass.Equal(t, infinity, NumberClass.Difference(infinity, minusi))
	ass.Equal(t, infinity, NumberClass.Difference(infinity, minusone))
	ass.Equal(t, infinity, NumberClass.Difference(infinity, zero))
	ass.Equal(t, infinity, NumberClass.Difference(infinity, one))
	ass.Equal(t, infinity, NumberClass.Difference(infinity, i))
	ass.True(t, NumberClass.Difference(infinity, undefined).IsUndefined())

	//	z - z => zero  {z != infinity}
	ass.Equal(t, zero, NumberClass.Difference(minusi, minusi))
	ass.Equal(t, zero, NumberClass.Difference(minusone, minusone))
	ass.Equal(t, zero, NumberClass.Difference(zero, zero))
	ass.Equal(t, zero, NumberClass.Difference(one, one))
	ass.Equal(t, zero, NumberClass.Difference(i, i))
	ass.True(t, NumberClass.Difference(infinity, infinity).IsUndefined())
	ass.True(t, NumberClass.Difference(undefined, undefined).IsUndefined())

	//	z * r
	ass.Equal(t, minusi, NumberClass.Scaled(minusi, 1.0))
	ass.Equal(t, minushalf, NumberClass.Scaled(minusone, 0.5))
	ass.Equal(t, zero, NumberClass.Scaled(zero, 5.0))
	ass.Equal(t, half, NumberClass.Scaled(one, 0.5))
	ass.Equal(t, i, NumberClass.Scaled(i, 1.0))
	ass.Equal(t, infinity, NumberClass.Scaled(infinity, 5.0))
	ass.True(t, NumberClass.Scaled(undefined, 5.0).IsUndefined())

	//	/z
	ass.Equal(t, infinity, NumberClass.Reciprocal(zero))
	ass.Equal(t, two, NumberClass.Reciprocal(half))
	ass.Equal(t, one, NumberClass.Reciprocal(one))
	ass.Equal(t, minushalf, NumberClass.Reciprocal(minustwo))
	ass.Equal(t, minusi, NumberClass.Reciprocal(i))
	ass.Equal(t, zero, NumberClass.Reciprocal(infinity))
	ass.True(t, NumberClass.Reciprocal(undefined).IsUndefined())

	//	*z
	ass.Equal(t, zero, NumberClass.Conjugate(zero))
	ass.Equal(t, one, NumberClass.Conjugate(one))
	ass.Equal(t, minusi, NumberClass.Conjugate(i))
	ass.Equal(t, i, NumberClass.Conjugate(minusi))
	ass.True(t, NumberClass.Conjugate(undefined).IsUndefined())

	//	z * zero => zero          {z != infinity}
	ass.Equal(t, zero, NumberClass.Product(zero, zero))
	ass.Equal(t, zero, NumberClass.Product(one, zero))
	ass.Equal(t, zero, NumberClass.Product(i, zero))
	ass.True(t, NumberClass.Product(infinity, zero).IsUndefined())
	ass.True(t, NumberClass.Product(undefined, zero).IsUndefined())

	//	z * one => z
	ass.Equal(t, zero, NumberClass.Product(zero, one))
	ass.Equal(t, one, NumberClass.Product(one, one))
	ass.Equal(t, i, NumberClass.Product(i, one))
	ass.Equal(t, infinity, NumberClass.Product(infinity, one))
	ass.True(t, NumberClass.Product(undefined, one).IsUndefined())

	//	z * infinity => infinity  {z != zero}
	ass.True(t, NumberClass.Product(zero, infinity).IsUndefined())
	ass.Equal(t, infinity, NumberClass.Product(one, infinity))
	ass.Equal(t, infinity, NumberClass.Product(i, infinity))
	ass.Equal(t, infinity, NumberClass.Product(infinity, infinity))

	//	zero / z => zero          {z != zero}
	ass.True(t, NumberClass.Quotient(zero, zero).IsUndefined())
	ass.Equal(t, zero, NumberClass.Quotient(zero, one))
	ass.Equal(t, zero, NumberClass.Quotient(zero, i))
	ass.Equal(t, zero, NumberClass.Quotient(zero, infinity))
	ass.True(t, NumberClass.Quotient(zero, undefined).IsUndefined())

	//	z / zero => infinity      {z != zero}
	ass.Equal(t, infinity, NumberClass.Quotient(one, zero))
	ass.Equal(t, infinity, NumberClass.Quotient(i, zero))
	ass.Equal(t, infinity, NumberClass.Quotient(infinity, zero))
	ass.True(t, NumberClass.Quotient(undefined, zero).IsUndefined())

	//	z / infinity => zero      {z != infinity}
	ass.Equal(t, zero, NumberClass.Quotient(one, infinity))
	ass.Equal(t, zero, NumberClass.Quotient(i, infinity))
	ass.True(t, NumberClass.Quotient(infinity, infinity).IsUndefined())
	ass.True(t, NumberClass.Quotient(undefined, infinity).IsUndefined())

	//	infinity / z => infinity  {z != infinity}
	ass.Equal(t, infinity, NumberClass.Quotient(infinity, zero))
	ass.Equal(t, infinity, NumberClass.Quotient(infinity, one))
	ass.Equal(t, infinity, NumberClass.Quotient(infinity, i))
	ass.True(t, NumberClass.Quotient(infinity, undefined).IsUndefined())

	//	y / z
	ass.Equal(t, one, NumberClass.Quotient(one, one))
	ass.Equal(t, one, NumberClass.Quotient(i, i))
	ass.Equal(t, i, NumberClass.Quotient(i, one))
	ass.Equal(t, two, NumberClass.Quotient(one, half))
	ass.Equal(t, one, NumberClass.Quotient(half, half))

	//	z ^ zero => one           {by definition}
	ass.Equal(t, one, NumberClass.Power(minusi, zero))
	ass.Equal(t, one, NumberClass.Power(minusone, zero))
	ass.Equal(t, one, NumberClass.Power(zero, zero))
	ass.Equal(t, one, NumberClass.Power(one, zero))
	ass.Equal(t, one, NumberClass.Power(i, zero))
	ass.Equal(t, one, NumberClass.Power(infinity, zero))
	ass.True(t, NumberClass.Power(undefined, zero).IsUndefined())

	//	zero ^ z => zero          {z != zero}
	ass.Equal(t, zero, NumberClass.Power(zero, one))
	ass.Equal(t, zero, NumberClass.Power(zero, i))
	ass.Equal(t, zero, NumberClass.Power(zero, infinity))
	ass.True(t, NumberClass.Power(zero, undefined).IsUndefined())

	//	z ^ infinity => zero      {|z| < one}
	//	z ^ infinity => one       {|z| = one}
	//	z ^ infinity => infinity  {|z| > one}
	ass.Equal(t, infinity, NumberClass.Power(minustwo, infinity))
	ass.Equal(t, one, NumberClass.Power(minusi, infinity))
	ass.Equal(t, one, NumberClass.Power(minusone, infinity))
	ass.Equal(t, zero, NumberClass.Power(minushalf, infinity))
	ass.Equal(t, zero, NumberClass.Power(half, infinity))
	ass.Equal(t, one, NumberClass.Power(one, infinity))
	ass.Equal(t, one, NumberClass.Power(i, infinity))
	ass.Equal(t, infinity, NumberClass.Power(two, infinity))

	//	infinity ^ z => infinity  {z != zero}
	ass.Equal(t, one, NumberClass.Power(infinity, zero))
	ass.Equal(t, infinity, NumberClass.Power(infinity, one))
	ass.Equal(t, infinity, NumberClass.Power(infinity, i))
	ass.Equal(t, infinity, NumberClass.Power(infinity, infinity))
	ass.True(t, NumberClass.Power(infinity, undefined).IsUndefined())

	//	one ^ z => one
	ass.Equal(t, one, NumberClass.Power(one, one))
	ass.Equal(t, one, NumberClass.Power(one, i))
	ass.Equal(t, one, NumberClass.Power(one, minusone))
	ass.Equal(t, one, NumberClass.Power(one, minusi))

	//	log(zero, z) => zero
	ass.True(t, NumberClass.Logarithm(zero, zero).IsUndefined())
	ass.Equal(t, zero, NumberClass.Logarithm(zero, i))
	ass.Equal(t, zero, NumberClass.Logarithm(zero, one))
	ass.True(t, NumberClass.Logarithm(zero, infinity).IsUndefined())
	ass.True(t, NumberClass.Logarithm(zero, undefined).IsUndefined())

	//	log(one, z) => infinity
	ass.Equal(t, infinity, NumberClass.Logarithm(one, zero))
	ass.True(t, NumberClass.Logarithm(one, one).IsUndefined())
	ass.Equal(t, infinity, NumberClass.Logarithm(one, infinity))
	ass.True(t, NumberClass.Logarithm(one, undefined).IsUndefined())

	//	log(infinity, z) => zero
	ass.True(t, NumberClass.Logarithm(infinity, zero).IsUndefined())
	ass.Equal(t, zero, NumberClass.Logarithm(infinity, one))
	ass.True(t, NumberClass.Logarithm(infinity, infinity).IsUndefined())
	ass.True(t, NumberClass.Logarithm(infinity, undefined).IsUndefined())
}

var PercentageClass = ele.PercentageClass()

func TestZeroPercentages(t *tes.T) {
	var v = PercentageClass.Percentage(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestPositivePercentages(t *tes.T) {
	var v = PercentageClass.Percentage(25)
	ass.Equal(t, 0.25, v.GetIntrinsic())
	ass.Equal(t, 25, v.AsInteger())
	ass.Equal(t, 25.0, v.AsFloat())
}

func TestNegativePercentages(t *tes.T) {
	var v = PercentageClass.Percentage(-75)
	ass.Equal(t, -0.75, v.GetIntrinsic())
	ass.Equal(t, -75, v.AsInteger())
	ass.Equal(t, -75.0, v.AsFloat())
}

func TestStringPercentages(t *tes.T) {
	var v = PercentageClass.PercentageFromString("-100.0%")
	ass.Equal(t, -1.0, v.GetIntrinsic())
	ass.Equal(t, -100, v.AsInteger())
	ass.Equal(t, -100.0, v.AsFloat())
	ass.Equal(t, "-100%", v.AsString())
}

var ProbabilityClass = ele.ProbabilityClass()

func TestBooleanProbabilities(t *tes.T) {
	var v1 = ProbabilityClass.ProbabilityFromBoolean(false)
	ass.Equal(t, 0.0, v1.AsFloat())

	var v2 = ProbabilityClass.ProbabilityFromBoolean(true)
	ass.Equal(t, 1.0, v2.AsFloat())
}

func TestZeroProbabilities(t *tes.T) {
	var v = ProbabilityClass.Probability(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestOneProbabilities(t *tes.T) {
	var v = ProbabilityClass.Probability(1.0)
	ass.Equal(t, 1.0, v.AsFloat())
}

func TestRandomProbabilities(t *tes.T) {
	ProbabilityClass.Random()
}

func TestStringProbabilities(t *tes.T) {
	var v = ProbabilityClass.ProbabilityFromString("p0")
	ass.Equal(t, 0.0, v.GetIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "p0", v.AsString())

	v = ProbabilityClass.ProbabilityFromString("p0.5")
	ass.Equal(t, 0.5, v.GetIntrinsic())
	ass.Equal(t, 0.5, v.AsFloat())
	ass.Equal(t, "p0.5", v.AsString())

	v = ProbabilityClass.ProbabilityFromString("p1")
	ass.Equal(t, 1.0, v.GetIntrinsic())
	ass.Equal(t, 1.0, v.AsFloat())
	ass.Equal(t, "p1", v.AsString())
}

func TestOtherProbabilities(t *tes.T) {
	var v1 = ProbabilityClass.Probability(0.25)
	ass.Equal(t, 0.25, v1.AsFloat())

	var v2 = ProbabilityClass.Probability(0.5)
	ass.Equal(t, 0.5, v2.AsFloat())

	var v3 = ProbabilityClass.Probability(0.75)
	ass.Equal(t, 0.75, v3.AsFloat())
}

func TestProbabilitieLibrary(t *tes.T) {
	var T = ProbabilityClass.Probability(0.75)
	var F = ProbabilityClass.Probability(0.25)

	var andNot = ProbabilityClass.And(ProbabilityClass.Not(T), ProbabilityClass.Not(T))
	var notIor = ProbabilityClass.Not(ProbabilityClass.Ior(T, T))
	ass.Equal(t, andNot, notIor)

	andNot = ProbabilityClass.And(ProbabilityClass.Not(T), ProbabilityClass.Not(F))
	notIor = ProbabilityClass.Not(ProbabilityClass.Ior(T, F))
	ass.Equal(t, andNot, notIor)

	andNot = ProbabilityClass.And(ProbabilityClass.Not(F), ProbabilityClass.Not(T))
	notIor = ProbabilityClass.Not(ProbabilityClass.Ior(F, T))
	ass.Equal(t, andNot, notIor)

	andNot = ProbabilityClass.And(ProbabilityClass.Not(F), ProbabilityClass.Not(F))
	notIor = ProbabilityClass.Not(ProbabilityClass.Ior(F, F))
	ass.Equal(t, andNot, notIor)

	var san = ProbabilityClass.And(T, ProbabilityClass.Not(T))
	ass.Equal(t, san, ProbabilityClass.San(T, T))

	san = ProbabilityClass.And(T, ProbabilityClass.Not(F))
	ass.Equal(t, san, ProbabilityClass.San(T, F))

	san = ProbabilityClass.And(F, ProbabilityClass.Not(T))
	ass.Equal(t, san, ProbabilityClass.San(F, T))

	san = ProbabilityClass.And(F, ProbabilityClass.Not(F))
	ass.Equal(t, san, ProbabilityClass.San(F, F))

	var xor = ProbabilityClass.Probability(ProbabilityClass.San(T, T).AsFloat() + ProbabilityClass.San(T, T).AsFloat())
	ass.Equal(t, xor, ProbabilityClass.Xor(T, T))

	xor = ProbabilityClass.Probability(ProbabilityClass.San(T, F).AsFloat() + ProbabilityClass.San(F, T).AsFloat())
	ass.Equal(t, xor, ProbabilityClass.Xor(T, F))

	xor = ProbabilityClass.Probability(ProbabilityClass.San(F, T).AsFloat() + ProbabilityClass.San(T, F).AsFloat())
	ass.Equal(t, xor, ProbabilityClass.Xor(F, T))

	xor = ProbabilityClass.Probability(ProbabilityClass.San(F, F).AsFloat() + ProbabilityClass.San(F, F).AsFloat())
	ass.Equal(t, xor, ProbabilityClass.Xor(F, F))
}

var ResourceClass = ele.ResourceClass()

func TestResourceWithAuthorityAndPath(t *tes.T) {
	var v = ResourceClass.Resource("<https://craterdog.com/About.html>")
	ass.Equal(t, "<https://craterdog.com/About.html>", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithPath(t *tes.T) {
	var v = ResourceClass.Resource("<mailto:craterdog@google.com>")
	ass.Equal(t, "<mailto:craterdog@google.com>", v.AsString())
	ass.Equal(t, "mailto", v.GetScheme())
	ass.Equal(t, "", v.GetAuthority())
	ass.Equal(t, "", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQuery(t *tes.T) {
	var v = ResourceClass.Resource("<https://craterdog.com/?foo=bar;bar=baz>")
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz>", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndFragment(t *tes.T) {
	var v = ResourceClass.Resource("<https://craterdog.com/#Home>")
	ass.Equal(t, "<https://craterdog.com/#Home>", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQueryAndFragment(t *tes.T) {
	var v = ResourceClass.Resource("<https://craterdog.com/?foo=bar;bar=baz#Home>")
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz#Home>", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}
