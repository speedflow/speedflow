package speedflow

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// ErrFile return the file error.
type ErrFile struct {
	Err     error
	Message string
	Path    string
}

// Error returns the error
func (e *ErrFile) Error() string {
	return fmt.Sprintf("error caused due to %v; file \"%s\" does not exists", e.Err, e.Path)
}

// Unwrap is used to make it work with errors.Is, errors.As.
func (e *ErrFile) Unwrap() error {
	return e.Err
}

// ErrParse returns the parsing error.
type ErrParse struct {
	Err     error
	Message string
	Path    string
}

// Error returns the error
func (e *ErrParse) Error() string {
	return fmt.Sprintf("error caused due to %v; unable to convert \"%s\" file", e.Err, e.Path)
}

// Unwrap is used to make it work with errors.Is, errors.As.
func (e *ErrParse) Unwrap() error {
	return e.Err
}

// Load loads the Speedflow file
func Load(path string) (err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return &ErrFile{Err: err, Path: path}
	}

	if err = yaml.Unmarshal(b, &SF); err != nil {
		return &ErrParse{Err: err, Path: path}
	}

	return nil
}
