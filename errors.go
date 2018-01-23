package params

import "fmt"

// ArgumentError is the error returned when arguments cannot be assigned correctly according to the
// specifications
type ArgumentError struct {
    err error
}

func argumentErrorf(format string, a ...interface{}) *ArgumentError {
    return &ArgumentError{ fmt.Errorf(format, a...) }
}

func (a *ArgumentError) Error() string {
    return a.err.Error()
}

// Cause (inheritDoc from causer interface)
func (a *ArgumentError) Cause() error {
    return a.err
}
