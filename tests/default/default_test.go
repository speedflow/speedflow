package default_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/speedflow/speedflow/tests"
)

func TestVersionText(t *testing.T) {
	output, err := tests.Execute("version")
	assert.NoError(t, err)
	assert.Equal(t, "Version:        dev\nCommit:         n/a\nBuild date:     n/a\n", output)
}

func TestVersionJSON(t *testing.T) {
	output, err := tests.Execute("version", "-o", "json")
	assert.NoError(t, err)
	assert.Equal(t, `{"version":"dev","commit":"n/a","date":"n/a"}`, output)
}

func TestVersionYAML(t *testing.T) {
	output, err := tests.Execute("version", "-o", "yaml")
	assert.NoError(t, err)
	assert.Equal(t, "version: dev\ncommit: n/a\ndate: n/a\n", output)
}

// func TestList(t *testing.T) {
// 	output, err := tests.Execute("-l")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Flow     Name        \ndefault  Default flow\n", output)
// }

// func TestDefault(t *testing.T) {
// 	output, err := tests.Execute()
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Hello World!\n", output)
// }
