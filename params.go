package params

import (
    "fmt"
)

// Value interface similar to flag.Value
type Value interface {
    Set(string) error
}

// ValueReceiver receives a list of strings to that can be received by this receiver
type ValueReceiver interface {
    Set([]string) error
}

type valueContainer struct {
    value Value
}

func (vc valueContainer) Set(vals []string) error {
    if len(vals) != 1 {
        return fmt.Errorf("Cannot set multiple values %v", vals)
    }
    return vc.value.Set(vals[0])
}

type ParamSet []ParamSpec

// func (ps *ParamSet) 

var defaultParamSet = new(ParamSet)

// DefaultParamSet gets the default ParamSet used when the Var / String etc functions are called
// on the package directly.
func DefaultParamSet() *ParamSet {
    return defaultParamSet
}

// ParamSpec is an interface for capturing positional arguments. The argument passing is done in
// 2 passes, the first pass will call MinLength() to find out the minimum number of values each
// argument captures. The second pass will call CaptureLength with the longest argv slice that it
// can capture to determine how many arguments it wants to capture.
type ParamSpec interface {

    // MinLength returns the minimum number of arguments to capture. For most common uses, this
    // should return 1 for required arguments and 0 for optional arguments
    MinLength() int

    // CaptureLength is called with the longest slice of argv that it can capture, and it can use
    // that to determine how many arguments it wants to capture. For a typical required or optional
    // argument, this will return 1. For a "variable length" argument, this will return
    // len(argvSlice).
    //
    // Optionally, this argument can check for a special "end of argument list" token
    // (typically "--") and stop capturing there.
    //
    // If the required arguments cannot be captured, an error should be returned.
    CaptureLength(argvSlice []string) (int, error)

    // Sets the given string slice, which has the length returned in CaptureLength.
    Set([]string) error

    Name() string

    // Metadata field to contain arbitrary data associated with this parameter. Typically this
    // would contain the usage string and other help values, but it can contain any arbitrary value.
    Metadata() interface{}
}

type commonParamSpec struct {
    name string
    minLength int
    maxLength int
    set func([]string) error
    metadata interface{}
}

var _ ParamSpec = (*commonParamSpec)(nil)

func (param *commonParamSpec) MinLength() int {
    return param.minLength
}

func (param *commonParamSpec) CaptureLength(argvSlice []string) (int, error) {
    if sliceLen := len(argvSlice); param.maxLength == -1 || param.maxLength > sliceLen {
        return sliceLen, nil
    }
    return param.maxLength, nil
}

func (param *commonParamSpec) Set(args []string) error {
    return param.set(args)
}

func (param *commonParamSpec) Name() string {
    // TODO: use Stringer instead?
    return param.name
}

func (param *commonParamSpec) Metadata() interface{} {
    return param.metadata
}

func (ps *ParamSet) VarValue(value Value, name string, optional bool, metadata interface{}) {
    minLength := 0
    if !optional { minLength = 1 }
    ps.VarListCustom(valueContainer{value}, name, minLength, 1, metadata)
}

// Var defines a parameter with a Value interface to receive the value
func VarValue(value Value, name string, optional bool, metadata interface{}) {
    defaultParamSet.VarValue(value, name, optional, metadata)
}

func (ps *ParamSet) Var(value ValueReceiver, name string, optional bool, metadata interface{}) {
    minLength := 0
    if !optional { minLength = 1 }
    ps.VarListCustom(value, name, minLength, 1, metadata)
}

func Var(value ValueReceiver, name string, optional bool, metadata interface{}) {
    defaultParamSet.Var(value, name, optional, metadata)
}

func (ps *ParamSet) VarList(value ValueReceiver, name string, optional bool, metadata interface{}) {
    minLength := 0
    if !optional { minLength = 1 }
    ps.VarListCustom(value, name, minLength, -1, metadata)
}

// VarList creates a parameter using ValueReceiver that captures all the remaining arguments.
func VarList(value ValueReceiver, name string, optional bool, metadata interface{}) {
    defaultParamSet.VarList(value, name, optional, metadata)
}

func (ps *ParamSet) VarListCustom(value ValueReceiver, name string, minLength int, maxLength int, metadata interface{}) {
    paramSpec := &commonParamSpec{
        name: name,
        minLength: minLength,
        maxLength: maxLength,
        set: value.Set,
        metadata: metadata,
    }
    *ps = append(*ps, paramSpec)
}

// VarListCustom creates a parameter using ValueReceiver that captures a list of the specified min
// and max length from the remaining arguments
func VarListCustom(value ValueReceiver, name string, minLength int, maxLength int, metadata interface{}) {
    defaultParamSet.VarListCustom(value, name, minLength, maxLength, metadata)
}

func (ps *ParamSet) Param(paramSpec ParamSpec) {
    *ps = append(*ps, paramSpec)
}

func Param(paramSpec ParamSpec) {
    defaultParamSet.Param(paramSpec)
}

// Parse parses the given string according to the parameters specs defined earlier using Var(...),
// StringParam(...) etc.
func Parse(argv []string) error {
    return defaultParamSet.Parse(argv)
}

func (paramSet *ParamSet) Parse(argv []string) error {
    if paramSet == nil {
        panic("paramSet cannot be null")
    }
    // First pass determines the min length of all the arguments
    minLengths := make([]int, len(*paramSet))
    minArgCount := 0
    for i, paramSpec := range *paramSet {
        minLengths[i] = paramSpec.MinLength()
        minArgCount += minLengths[i]
    }

    // Iterate over the arguments again to capture the variable length arguments
    // In this pass, the remaining arguments are allocated 
    remainingMinArg := minArgCount
    argvIndex := 0
    for i, paramSpec := range *paramSet {
        ml := minLengths[i]
        sliceEnd := len(argv) - remainingMinArg + ml
        if sliceEnd <= argvIndex {
            if ml > 0 {
                // Pass an empty slice to CaptureLength, even though we don't have enough arguments
                // so that the error message can be generated
                // sliceEnd = argvIndex
                // FIXME: comments
                return argumentErrorf(`Missing required argument "%s"`, paramSpec.Name())
            } else {
                // No argument available for parsing, but this param is not required
                continue
            }
        }
        l, err := paramSpec.CaptureLength(argv[argvIndex:sliceEnd])
        if err != nil { return &ArgumentError{ err } }
        if l < ml {
            return argumentErrorf(`Argument "%s" captured less than min length`, paramSpec.Name())
        }
        if argvIndex > argvIndex + l {
            return argumentErrorf(`Argument "%s" captured more than it should`, paramSpec.Name())
        }
        if argvIndex < argvIndex + l {
            // Don't call Set if the slice is empty, to avoid initializing pointers when no value
            // will be added
            paramSpec.Set(argv[argvIndex:argvIndex + l])
        }
        argvIndex += l
        remainingMinArg -= ml
    }

    if argvIndex < len(argv) {
        return argumentErrorf(`Too many arguments. %d remaining`, len(argv) - argvIndex)
    }

    return nil
}
