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

// A container for Value to implement ValueReceiver
type valueContainer struct {
    value Value
}

func (vc valueContainer) Set(vals []string) error {
    if len(vals) != 1 {
        return fmt.Errorf("Cannot set multiple values %v", vals)
    }
    return vc.value.Set(vals[0])
}

// ParamSet represents a set of parameter specifications. Typically, most applications do not need
// to use this, but can use the functions defined in this "params" package directly, which forwards
// the methods to DefaultParamSet.
type ParamSet []ParamSpec

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

    // Metadata field to contain arbitrary data associated with this parameter. Typically this
    // would contain the usage string and other help values, but it can contain any arbitrary value.
    Metadata() interface{}

    fmt.Stringer
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

func (param *commonParamSpec) String() string {
    return param.name
}

func (param *commonParamSpec) Metadata() interface{} {
    return param.metadata
}

// VarValue defines a parameter with a Value interface to receive the value.
func (ps *ParamSet) VarValue(value Value, name string, optional bool, metadata interface{}) {
    ps.Param(NewValueParamSpec(value, name, optional, metadata))
}

// VarValue defines a parameter with a Value interface to receive the value on the DefaultParamSet.
func VarValue(value Value, name string, optional bool, metadata interface{}) {
    defaultParamSet.VarValue(value, name, optional, metadata)
}

// NewValueParamSpec defines a new ParamSpec with a Value interface to receive the value.
func NewValueParamSpec(value Value, name string, optional bool, metadata interface{}) ParamSpec {
    return NewParamSpec(valueContainer{value}, name, optional, metadata)
}

// Var defines a parameter with a ValueReceiver to receive the value. The ValueReceiver used in this
// method is only expected to receive 1 value in the string slice.
func (ps *ParamSet) Var(value ValueReceiver, name string, optional bool, metadata interface{}) {
    ps.Param(NewParamSpec(value, name, optional, metadata))
}

// Var defines a parameter with a ValueReceiver to receive the value on the DefaultParamSet. The
// ValueReceiver used in this function is only expected to receive 1 value in the string slice.
func Var(value ValueReceiver, name string, optional bool, metadata interface{}) {
    defaultParamSet.Var(value, name, optional, metadata)
}

// NewParamSpec defines a new ParamSpec with a ValueReceiver to receive the value. The ValueReceiver
// used in this function is only expected to receive 1 value in the string slice.
func NewParamSpec(value ValueReceiver, name string, optional bool, metadata interface{}) ParamSpec {
    minLength := 0
    if !optional { minLength = 1 }
    return NewCustomParamSpec(value, name, minLength, 1, metadata)
}

// VarList defines a parameter list using ValueReceiver that captures all the remaining arguments.
func (ps *ParamSet) VarList(value ValueReceiver, name string, optional bool, metadata interface{}) {
    ps.Param(NewListParamSpec(value, name, optional, metadata))
}

// VarList defines a parameter list using ValueReceiver that captures all the remaining arguments on
// the DefaultParamSet.
func VarList(value ValueReceiver, name string, optional bool, metadata interface{}) {
    defaultParamSet.VarList(value, name, optional, metadata)
}

// NewListParamSpec defines a parameter list using ValueReceiver that captures all the remaining
// arguments.
func NewListParamSpec(value ValueReceiver, name string, optional bool, metadata interface{}) ParamSpec {
    minLength := 0
    if !optional { minLength = 1 }
    return NewCustomParamSpec(value, name, minLength, -1, metadata)
}

// VarListCustom adds a parameter list spec using ValueReceiver that captures a list of the
// specified min and max length from the remaining arguments.
func (ps *ParamSet) VarListCustom(value ValueReceiver, name string, minLength int, maxLength int, metadata interface{}) {
    ps.Param(NewCustomParamSpec(value, name, minLength, maxLength, metadata))
}

// VarListCustom adds a parameter list spec using ValueReceiver that captures a list of the
// specified min and max length from the remaining arguments on the DefaultParamSet.
func VarListCustom(value ValueReceiver, name string, minLength int, maxLength int, metadata interface{}) {
    defaultParamSet.VarListCustom(value, name, minLength, maxLength, metadata)
}

// NewCustomParamSpec creates a parameter list spec using ValueReceiver that captures a list of the
// specified min and max length from the remaining arguments.
func NewCustomParamSpec(value ValueReceiver, name string, minLength int, maxLength int, metadata interface{}) ParamSpec {
    return &commonParamSpec{
        name: name,
        minLength: minLength,
        maxLength: maxLength,
        set: value.Set,
        metadata: metadata,
    }
}

// Param adds a ParamSpec to the ParamSet
func (ps *ParamSet) Param(paramSpec ParamSpec) {
    *ps = append(*ps, paramSpec)
}

// Param adds a ParamSpec to the DefaultParamSet
func Param(paramSpec ParamSpec) {
    defaultParamSet.Param(paramSpec)
}

// Parse parses the given string according to the parameters specs defined earlier using Var(...),
// StringParam(...) etc.
func Parse(argv []string) error {
    return defaultParamSet.Parse(argv)
}

// Parse parses the given string list as the arguments, according to the ParamSpecs previously
// added to the ParamSet.
// See the documentation on ParamSpec for details on the parsing.
func (ps *ParamSet) Parse(argv []string) error {
    if ps == nil {
        // No parameter set, just return
        return nil
    }
    // First pass determines the min length of all the arguments
    minLengths := make([]int, len(*ps))
    minArgCount := 0
    for i, paramSpec := range *ps {
        minLengths[i] = paramSpec.MinLength()
        minArgCount += minLengths[i]
    }

    // Iterate over the arguments again to capture the variable length arguments
    // In this pass, the remaining arguments are allocated 
    remainingMinArg := minArgCount
    argvIndex := 0
    for i, paramSpec := range *ps {
        ml := minLengths[i]
        sliceEnd := len(argv) - remainingMinArg + ml
        if sliceEnd <= argvIndex {
            if ml > 0 {
                // Argument is required but missing. Print error message and return.
                return argumentErrorf(`Missing required argument "%s"`, paramSpec.String())
            }
            // No argument available for parsing, but this param is not required
            continue
        }
        l, err := paramSpec.CaptureLength(argv[argvIndex:sliceEnd])
        if err != nil { return &ArgumentError{ err } }
        if l < ml {
            return argumentErrorf(`Argument "%s" captured less than min length`, paramSpec.String())
        }
        if l > 0 {
            // Don't call Set if the slice is empty, to avoid initializing pointers when no values
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
