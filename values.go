// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// LICENSE file: https://github.com/golang/go/blob/master/LICENSE
// Copied from https://github.com/golang/go/blob/master/src/flag/flag.go

package params

import (
	"time"
	"strconv"
)

// -- bool Value
type BoolValue bool

func NewBoolValue(val bool, p *bool) *BoolValue {
	*p = val
	return (*BoolValue)(p)
}

func (b *BoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = BoolValue(v)
	return err
}

func (b *BoolValue) Get() interface{} { return bool(*b) }

func (b *BoolValue) String() string { return strconv.FormatBool(bool(*b)) }

func (b *BoolValue) IsBoolFlag() bool { return true }

// -- int Value
type IntValue int

func NewIntValue(val int, p *int) *IntValue {
	*p = val
	return (*IntValue)(p)
}

func (i *IntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	*i = IntValue(v)
	return err
}

func (i *IntValue) Get() interface{} { return int(*i) }

func (i *IntValue) String() string { return strconv.Itoa(int(*i)) }

// -- int64 Value
type Int64Value int64

func NewInt64Value(val int64, p *int64) *Int64Value {
	*p = val
	return (*Int64Value)(p)
}

func (i *Int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = Int64Value(v)
	return err
}

func (i *Int64Value) Get() interface{} { return int64(*i) }

func (i *Int64Value) String() string { return strconv.FormatInt(int64(*i), 10) }

// -- uint Value
type UintValue uint

func NewUintValue(val uint, p *uint) *UintValue {
	*p = val
	return (*UintValue)(p)
}

func (i *UintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	*i = UintValue(v)
	return err
}

func (i *UintValue) Get() interface{} { return uint(*i) }

func (i *UintValue) String() string { return strconv.FormatUint(uint64(*i), 10) }

// -- uint64 Value
type Uint64Value uint64

func NewUint64Value(val uint64, p *uint64) *Uint64Value {
	*p = val
	return (*Uint64Value)(p)
}

func (i *Uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = Uint64Value(v)
	return err
}

func (i *Uint64Value) Get() interface{} { return uint64(*i) }

func (i *Uint64Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// -- string Value
type StringValue string

func NewStringValue(val string, p *string) *StringValue {
	*p = val
	return (*StringValue)(p)
}

func (s *StringValue) Set(val string) error {
	*s = StringValue(val)
	return nil
}

func (s *StringValue) Get() interface{} { return string(*s) }

func (s *StringValue) String() string { return string(*s) }

// -- float64 Value
type Float64Value float64

func NewFloat64Value(val float64, p *float64) *Float64Value {
	*p = val
	return (*Float64Value)(p)
}

func (f *Float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = Float64Value(v)
	return err
}

func (f *Float64Value) Get() interface{} { return float64(*f) }

func (f *Float64Value) String() string { return strconv.FormatFloat(float64(*f), 'g', -1, 64) }

// -- time.Duration Value
type DurationValue time.Duration

func NewDurationValue(val time.Duration, p *time.Duration) *DurationValue {
	*p = val
	return (*DurationValue)(p)
}

func (d *DurationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = DurationValue(v)
	return err
}

func (d *DurationValue) Get() interface{} { return time.Duration(*d) }

func (d *DurationValue) String() string { return (*time.Duration)(d).String() }
