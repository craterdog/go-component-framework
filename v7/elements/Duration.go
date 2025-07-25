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

package elements

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v7"
	mat "math"
	reg "regexp"
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func DurationClass() DurationClassLike {
	return durationClass()
}

// Constructor Methods

func (c *durationClass_) Duration(
	milliseconds int,
) DurationLike {
	if milliseconds < 0 {
		milliseconds = -milliseconds
	}
	return duration_(milliseconds)
}

func (c *durationClass_) DurationFromString(
	source string,
) DurationLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the duration constructor method: %s",
			source,
		)
		panic(message)
	}
	return duration_(c.durationFromMatches(matches))
}

// Constant Methods

func (c *durationClass_) Undefined() DurationLike {
	return c.undefined_
}

func (c *durationClass_) MillisecondsPerSecond() int {
	return c.millisecondsPerSecond_
}

func (c *durationClass_) MillisecondsPerMinute() int {
	return c.millisecondsPerMinute_
}

func (c *durationClass_) MillisecondsPerHour() int {
	return c.millisecondsPerHour_
}

func (c *durationClass_) MillisecondsPerDay() int {
	return c.millisecondsPerDay_
}

func (c *durationClass_) MillisecondsPerWeek() int {
	return c.millisecondsPerWeek_
}

func (c *durationClass_) MillisecondsPerMonth() int {
	return c.millisecondsPerMonth_
}

func (c *durationClass_) MillisecondsPerYear() int {
	return c.millisecondsPerYear_
}

func (c *durationClass_) DaysPerMonth() float64 {
	return c.daysPerMonth_
}

func (c *durationClass_) DaysPerYear() float64 {
	return c.daysPerYear_
}

func (c *durationClass_) WeeksPerMonth() float64 {
	return c.weeksPerMonth_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v duration_) GetClass() DurationClassLike {
	return durationClass()
}

func (v duration_) AsIntrinsic() int {
	return int(v)
}

// Attribute Methods

// Discrete Methods

func (v duration_) AsString() string {
	var builder sts.Builder
	builder.WriteString("~")
	if v.IsNegative() {
		builder.WriteString("-")
	}
	builder.WriteString("P")
	var float = mat.Abs(v.AsWeeks())
	var weeks = int(float)
	if float64(weeks) == float {
		// It is an exact number of weeks.
		builder.WriteString(stc.FormatInt(int64(weeks), 10))
		builder.WriteString("W")
		return builder.String()
	}
	var years = v.GetYears()
	if years > 0 {
		builder.WriteString(stc.FormatInt(int64(years), 10))
		builder.WriteString("Y")
	}
	var months = v.GetMonths()
	if months > 0 {
		builder.WriteString(stc.FormatInt(int64(months), 10))
		builder.WriteString("M")
	}
	var days = v.GetDays()
	if days > 0 {
		builder.WriteString(stc.FormatInt(int64(days), 10))
		builder.WriteString("D")
	}
	var hours = v.GetHours()
	var minutes = v.GetMinutes()
	var seconds = v.GetSeconds()
	var milliseconds = v.GetMilliseconds()
	if hours+minutes+seconds+milliseconds == 0 {
		// There is no time part of the duration.
		return builder.String()
	}
	builder.WriteString("T")
	if hours > 0 {
		builder.WriteString(stc.FormatInt(int64(hours), 10))
		builder.WriteString("H")
	}
	if minutes > 0 {
		builder.WriteString(stc.FormatInt(int64(minutes), 10))
		builder.WriteString("M")
	}
	if seconds+milliseconds > 0 {
		builder.WriteString(stc.FormatInt(int64(seconds), 10))
		if milliseconds > 0 {
			builder.WriteString(".")
			builder.WriteString(stc.FormatInt(int64(milliseconds), 10))
		}
		builder.WriteString("S")
	}
	return builder.String()
}

func (v duration_) AsInteger() int {
	return int(v)
}

func (v duration_) IsDefined() bool {
	return v >= 0
}

func (v duration_) IsMinimum() bool {
	return v == 0
}

func (v duration_) IsZero() bool {
	return v == 0
}

func (v duration_) IsMaximum() bool {
	return v == mat.MaxInt64
}

// Polarized Methods

func (v duration_) IsNegative() bool {
	return v < 0
}

// Temporal Methods

func (v duration_) AsMilliseconds() float64 {
	return float64(v)
}

func (v duration_) AsSeconds() float64 {
	return float64(v) / float64(durationClass().millisecondsPerSecond_)
}

func (v duration_) AsMinutes() float64 {
	return float64(v) / float64(durationClass().millisecondsPerMinute_)
}

func (v duration_) AsHours() float64 {
	return float64(v) / float64(durationClass().millisecondsPerHour_)
}

func (v duration_) AsDays() float64 {
	return float64(v) / float64(durationClass().millisecondsPerDay_)
}

func (v duration_) AsWeeks() float64 {
	return float64(v) / float64(durationClass().millisecondsPerWeek_)
}

func (v duration_) AsMonths() float64 {
	return float64(v) / float64(durationClass().millisecondsPerMonth_)
}

func (v duration_) AsYears() float64 {
	return float64(v) / float64(durationClass().millisecondsPerYear_)
}

// Factored Methods

func (v duration_) GetMilliseconds() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off everything but the milliseconds.
	milliseconds = milliseconds % durationClass().millisecondsPerSecond_
	return milliseconds
}

func (v duration_) GetSeconds() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass().millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass().millisecondsPerDay_)

	// Strip off the hours.
	milliseconds = milliseconds - (v.GetHours() * durationClass().millisecondsPerHour_)

	// Strip off the minutes.
	milliseconds = milliseconds - (v.GetMinutes() * durationClass().millisecondsPerMinute_)

	// Convert to seconds.
	var seconds = milliseconds / durationClass().millisecondsPerSecond_
	return seconds
}

func (v duration_) GetMinutes() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass().millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass().millisecondsPerDay_)

	// Strip off the hours.
	milliseconds = milliseconds - (v.GetHours() * durationClass().millisecondsPerHour_)

	// Convert to minutes.
	var minutes = milliseconds / durationClass().millisecondsPerMinute_
	return minutes
}

func (v duration_) GetHours() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass().millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass().millisecondsPerDay_)

	// Convert to hours.
	var hours = milliseconds / durationClass().millisecondsPerHour_
	return hours
}

func (v duration_) GetDays() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass().millisecondsPerMonth_)

	// Convert to days.
	var days = milliseconds / durationClass().millisecondsPerDay_
	return days
}

func (v duration_) GetWeeks() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Convert to weeks.
	var weeks = milliseconds / durationClass().millisecondsPerWeek_
	return weeks
}

func (v duration_) GetMonths() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Convert to months.
	var months = milliseconds / durationClass().millisecondsPerMonth_
	return months
}

func (v duration_) GetYears() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Convert to years.
	var years = milliseconds / durationClass().millisecondsPerYear_
	return years
}

// PROTECTED INTERFACE

func (v duration_) String() string {
	return v.AsString()
}

// Private Methods

func (c *durationClass_) durationFromMatches(matches []string) int {
	var milliseconds = 0.0
	if len(matches[1]) > 0 {
		// The duration is in weeks.
		var float, _ = stc.ParseFloat(matches[1], 64)
		milliseconds += float * float64(c.millisecondsPerWeek_)
		return int(milliseconds)
	}
	if len(matches[2]) > 0 {
		// The duration has a years component.
		var float, _ = stc.ParseFloat(matches[2], 64)
		milliseconds += float * float64(c.millisecondsPerYear_)
	}
	if len(matches[3]) > 0 {
		// The duration has a months component.
		var float, _ = stc.ParseFloat(matches[3], 64)
		milliseconds += float * float64(c.millisecondsPerMonth_)
	}
	if len(matches[4]) > 0 {
		// The duration has a days component.
		var float, _ = stc.ParseFloat(matches[4], 64)
		milliseconds += float * float64(c.millisecondsPerDay_)
	}
	if len(matches[5]) > 0 {
		// The duration has an hours component.
		var float, _ = stc.ParseFloat(matches[5], 64)
		milliseconds += float * float64(c.millisecondsPerHour_)
	}
	if len(matches[6]) > 0 {
		// The duration has a minutes component.
		var float, _ = stc.ParseFloat(matches[6], 64)
		milliseconds += float * float64(c.millisecondsPerMinute_)
	}
	if len(matches[7]) > 0 {
		// The duration has a seconds component.
		var float, _ = stc.ParseFloat(matches[7], 64)
		milliseconds += float * float64(c.millisecondsPerSecond_)
	}
	return int(milliseconds)
}

func (c *durationClass_) magnitude(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	days_     = "(" + timespan_ + ")D"
	hours_    = "(" + timespan_ + ")H"
	minutes_  = "(" + timespan_ + ")M"
	months_   = "(" + timespan_ + ")M"
	seconds_  = "(" + timespan_ + ")S"
	timespan_ = "0|" + ordinal_ + "(?:" + fraction_ + ")?"
	weeks_    = "(" + timespan_ + ")W"
	years_    = "(" + timespan_ + ")Y"
)

// Instance Structure

type duration_ int

// Class Structure

type durationClass_ struct {
	// Declare the class constants.
	matcher_               *reg.Regexp
	undefined_             DurationLike
	millisecondsPerSecond_ int
	millisecondsPerMinute_ int
	millisecondsPerHour_   int
	millisecondsPerDay_    int
	millisecondsPerWeek_   int
	millisecondsPerMonth_  int
	millisecondsPerYear_   int
	daysPerMonth_          float64
	daysPerYear_           float64
	weeksPerMonth_         float64
}

// Class Reference

func durationClass() *durationClass_ {
	return durationClassReference_
}

var durationClassReference_ = &durationClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^~P(?:(?:" + weeks_ + ")|(?:(?:" + years_ + ")?(?:" + months_ +
			")?(?:" + days_ + ")?(?:T(?:" + hours_ + ")?(?:" + minutes_ +
			")?(?:" + seconds_ + ")?)?))",
	),
	undefined_: duration_(-1),

	// These are locked to the Earth's daily revolutions.
	millisecondsPerSecond_: 1000,
	millisecondsPerMinute_: 60000,
	millisecondsPerHour_:   3600000,
	millisecondsPerDay_:    86400000,
	millisecondsPerWeek_:   604800000,

	// These are locked to the Earth's yearly orbit around the sun.
	millisecondsPerMonth_: 2629746000, // An average but exact value.
	millisecondsPerYear_:  31556952000,

	// Tying the two together is where things get messy.
	daysPerMonth_:  30.436875,
	daysPerYear_:   365.2425,
	weeksPerMonth_: 4.348125,
}
