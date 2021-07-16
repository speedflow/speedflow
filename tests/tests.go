package tests

import (
	"bytes"

	"github.com/speedflow/speedflow/internal/speedflow/command"
)

// Execute executes a real CLI test
func Execute(args ...string) (string, error) {
	bufIn := new(bytes.Buffer)
	bufOut := new(bytes.Buffer)
	command.Exitable = false
	_, err := command.ExecuteC(bufIn, bufOut, bufOut, args...)
	return bufOut.String(), err
}
