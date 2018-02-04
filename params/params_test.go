package params

import (
    "reflect"
    "testing"
)

func TestStringPtrParsing(t *testing.T) {

    // Test setup

    p := new(ParamSet)
    arg1 := p.String("arg1", false, nil)
    arg2 := p.String("arg2", false, nil)
    arg3 := p.String("arg3", false, nil)

    // Test execution

    err := p.Parse([]string { "arg1", "arg2", "arg3" })

    // Assertions

    if err != nil { t.Errorf("Error is not nil: %v", err) }
    if *arg1 != "arg1" { t.Errorf(`arg1 should be "arg1", but was %s`, *arg1) }
    if *arg2 != "arg2" { t.Errorf(`arg2 should be "arg2", but was %s`, *arg2) }
    if *arg3 != "arg3" { t.Errorf(`arg3 should be "arg3", but was %s`, *arg3) }
}

func TestParamListParsing(t *testing.T) {

    // Test setup

    p := new(ParamSet)
    arg1 := p.String("arg1", false, nil)
    arg2 := p.String("arg2", false, nil)
    argRest := p.StringList("argRest", false, nil)
    arg3 := p.String("arg3", false, nil)

    // Test execution

    err := p.Parse([]string { "arg1", "arg2", "arg3", "arg4", "arg5" })

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

    p := new(ParamSet)
    arg1 := p.String("arg1", false, nil)
    arg2 := p.String("arg2", false, nil)
    argRest := p.StringListCustom("argRest", 1, 2, nil)
    arg3 := p.String("arg3", true, nil)

    // Test execution

    err := p.Parse([]string { "arg1", "arg2", "arg3", "arg4", "arg5" })

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

    p := new(ParamSet)
    p.String("arg1", false, nil)

    t.Run("too many arguments", func (t *testing.T) {
        err := p.Parse([]string { "arg1", "arg2", "arg3" })

        // Assertions
        if err == nil || err.Error() != "Too many arguments. 2 remaining" {
            t.Errorf("Error should be thrown, but was: %v", err)
        }
    })

    t.Run("too many arguments", func (t *testing.T) {
        err := p.Parse([]string {})

        // Assertions
        if err == nil || err.Error() != `Missing required argument "arg1"` {
            t.Errorf("Error should be thrown, but was: %v", err)
        }
    })
}

func TestRequiredArgAfterList(t *testing.T) {
    // Setup
    p := new(ParamSet)
    p.StringList("argList", true, nil)
    p.String("arg2", false, nil)

    // Execution
    err := p.Parse([]string {})

    // Assertions
    if err == nil || err.Error() != `Missing required argument "arg2"` {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestMetadata(t *testing.T) {
    // Setup
    p := new(ParamSet)
    helloArg := p.String("hello", false, "mystring")

    // Execution
    err := p.Parse([]string {"world"})

    // Assertions
    if err != nil { t.Errorf("Unexpected error %v", err) }
    if *helloArg != "world" { t.Errorf(`helloArg should be "world"`, *helloArg) }
    if metadata := (*p)[0].Metadata(); metadata != "mystring" {
        t.Errorf(`Unexpected metadata "%v"`, metadata)
    }
}
