package default_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/speedflow/speedflow/internal/speedflow/command"
)

func executeCLI(args ...string) (string, error) {
	bufIn := new(bytes.Buffer)
	bufOut := new(bytes.Buffer)
	_, err := command.ExecuteC(bufIn, bufOut, bufOut, args...)
	return bufOut.String(), err
}

func TestVersionText(t *testing.T) {
	output, err := executeCLI("version")
	assert.NoError(t, err)
	assert.Equal(t, "Version:        dev\nCommit:         n/a\nBuild date:     n/a\n", output)
}

func TestVersionJSON(t *testing.T) {
	output, err := executeCLI("version", "-o", "json")
	assert.NoError(t, err)
	assert.Equal(t, `{"version":"dev","commit":"n/a","date":"n/a"}`, output)
}

func TestVersionYAML(t *testing.T) {
	output, err := executeCLI("version", "-o", "yaml")
	assert.NoError(t, err)
	assert.Equal(t, "version: dev\ncommit: n/a\ndate: n/a\n", output)
}

func TestList(t *testing.T) {
	output, err := executeCLI("-l")
	assert.NoError(t, err)
	assert.Equal(t, "Flow     Name        \ndefault  Default flow\n", output)
}

func TestDefault(t *testing.T) {
	output, err := executeCLI()
	assert.NoError(t, err)
	assert.Equal(t, "Hello World!\n", output)
}
