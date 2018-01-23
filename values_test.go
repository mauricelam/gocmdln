package params

import (
    "reflect"
    "testing"
    "time"
)

func TestBoolValueList(t *testing.T) {
    bl := new(BoolValueList)
    err := bl.Set([]string { "true", "1", "0", "false" })

    if err != nil { t.Errorf("Error: %v", err) }
    if !reflect.DeepEqual(*bl, BoolValueList([]bool { true, true, false, false })) {
        t.Errorf("Unexpected bl: %v", bl)
    }
}

func TestIntValueList(t *testing.T) {
    il := new(IntValueList)
    err := il.Set([]string { "123", "456", "-654", "0" })

    if err != nil { t.Errorf("Error: %v", err) }
    if !reflect.DeepEqual(*il, IntValueList([]int { 123, 456, -654, 0 })) {
        t.Errorf("Unexpected il: %v", il)
    }
}

func TestInt64ValueList(t *testing.T) {
    il := new(Int64ValueList)
    err := il.Set([]string { "123456789123456", "456", "-654", "0" })

    if err != nil { t.Errorf("Error: %v", err) }
    if !reflect.DeepEqual(*il, Int64ValueList([]int64 { 123456789123456, 456, -654, 0 })) {
        t.Errorf("Unexpected il: %v", il)
    }
}

func TestUintValueList(t *testing.T) {
    il := new(UintValueList)
    err := il.Set([]string { "123", "456", "0" })

    if err != nil { t.Errorf("Error: %v", err) }
    if !reflect.DeepEqual(*il, UintValueList([]uint { 123, 456, 0 })) {
        t.Errorf("Unexpected il: %v", il)
    }
}

func TestUint64ValueList(t *testing.T) {
    il := new(Uint64ValueList)
    err := il.Set([]string { "123456789123456", "456", "0" })

    if err != nil { t.Errorf("Error: %v", err) }
    if !reflect.DeepEqual(*il, Uint64ValueList([]uint64 { 123456789123456, 456, 0 })) {
        t.Errorf("Unexpected il: %v", il)
    }
}

func TestStringValueList(t *testing.T) {
    il := new(StringValueList)
    err := il.Set([]string { "abc", "def" })

    if err != nil { t.Errorf("Error: %v", err) }
    if !reflect.DeepEqual(*il, StringValueList([]string { "abc", "def" })) {
        t.Errorf("Unexpected il: %v", il)
    }
}

func TestFloat64ValueList(t *testing.T) {
    il := new(Float64ValueList)
    err := il.Set([]string { "123", "3.14", "-2.2" })

    if err != nil { t.Errorf("Error: %v", err) }
    if !reflect.DeepEqual(*il, Float64ValueList([]float64 { 123, 3.14, -2.2 })) {
        t.Errorf("Unexpected il: %v", il)
    }
}

func TestDurationValueList(t *testing.T) {
    il := new(DurationValueList)
    err := il.Set([]string { "300ms", "-1.5h", "2h45m" })

    if err != nil { t.Errorf("Error: %v", err) }
    if !reflect.DeepEqual(*il, DurationValueList([]time.Duration {
        time.Duration(300) * time.Millisecond,
        time.Duration(-90) * time.Minute,
        time.Duration(165) * time.Minute,
    })) {
        t.Errorf("Unexpected il: %v", il)
    }
}
