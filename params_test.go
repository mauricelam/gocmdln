package params

import (
    "reflect"
    "testing"
)

func TestStringPtrParsing(t *testing.T) {

    // Test setup

    reset()
    arg1 := String("arg1", false)
    arg2 := String("arg2", false)
    arg3 := String("arg3", false)

    // Test execution

    err := Parse([]string { "arg1", "arg2", "arg3" })

    // Assertions

    if err != nil { t.Errorf("Error is not nil: %v", err) }
    if *arg1 != "arg1" { t.Errorf(`arg1 should be "arg1", but was %s`, *arg1) }
    if *arg2 != "arg2" { t.Errorf(`arg2 should be "arg2", but was %s`, *arg2) }
    if *arg3 != "arg3" { t.Errorf(`arg3 should be "arg3", but was %s`, *arg3) }
}

func TestParamListParsing(t *testing.T) {

    // Test setup

    reset()
    arg1 := String("arg1", false)
    arg2 := String("arg2", false)
    argRest := StringList("argRest", false)
    arg3 := String("arg3", false)

    // Test execution

    err := Parse([]string { "arg1", "arg2", "arg3", "arg4", "arg5" })

    // Assertions

    if err != nil { t.Errorf("Error is not nil: %v", err) }
    if *arg1 != "arg1" { t.Errorf(`arg1 should be "arg1", but was %s`, *arg1) }
    if *arg2 != "arg2" { t.Errorf(`arg2 should be "arg2", but was %s`, *arg2) }
    if !reflect.DeepEqual(*argRest, []string { "arg3", "arg4" }) {
        t.Errorf(`argrest should be ["arg3", "arg4"], but was %s`, *argRest)
    }
    if *arg3 != "arg5" { t.Errorf(`arg3 should be "arg5", but was %s`, *arg3) }
}

func TestCustomParamListParsing(t *testing.T) {

    // Test setup

    reset()
    arg1 := String("arg1", false)
    arg2 := String("arg2", false)
    argRest := StringListCustom("argRest", 1, 2)
    arg3 := String("arg3", true)

    // Test execution

    err := Parse([]string { "arg1", "arg2", "arg3", "arg4", "arg5" })

    // Assertions

    if err != nil { t.Errorf("Error is not nil: %v", err) }
    if *arg1 != "arg1" { t.Errorf(`arg1 should be "arg1", but was %s`, *arg1) }
    if *arg2 != "arg2" { t.Errorf(`arg2 should be "arg2", but was %s`, *arg2) }
    if !reflect.DeepEqual(*argRest, []string { "arg3", "arg4" }) {
        t.Errorf(`argrest should be ["arg3", "arg4"], but was %s`, *argRest)
    }
    if *arg3 != "arg5" { t.Errorf(`arg3 should be "arg5", but was %s`, *arg3) }
}

func TestErrors(t *testing.T) {

    // Test setup

    reset()
    String("arg1", false)

    t.Run("too many arguments", func (t *testing.T) {
        err := Parse([]string { "arg1", "arg2", "arg3" })

        // Assertions
        if err == nil || err.Error() != "Too many arguments. 2 remaining" {
            t.Errorf("Error should be thrown, but was: %v", err)
        }
    })

    t.Run("too many arguments", func (t *testing.T) {
        err := Parse([]string {})

        // Assertions
        if err == nil || err.Error() != `Missing required argument "arg1"` {
            t.Errorf("Error should be thrown, but was: %v", err)
        }
    })
}
