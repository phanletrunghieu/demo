package errors

import (
	"bytes"
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

// Errors .
type Errors []error

// Err .
func (e Errors) Err() error {
	if len(e) == 0 {
		return nil
	}

	return e
}

func (e Errors) Error() string {
	var buf bytes.Buffer

	if n := len(e); n == 1 {
		buf.WriteString("1 error: ")
	} else {
		fmt.Fprintf(&buf, "%d errors: ", n)
	}

	for i, err := range e {
		if i != 0 {
			buf.WriteString("; ")
		}

		buf.WriteString(err.Error())
	}

	return buf.String()
}

// Slice .
func (e Errors) Slice() []error {
	return []error(e)
}

// New This is convenience method so we don't have to fight with package imports.
func New(message string) error {
	return pkgerrors.New(message)
}

// Errorf .
func Errorf(format string, args ...interface{}) error {
	return pkgerrors.Errorf(format, args)
}
