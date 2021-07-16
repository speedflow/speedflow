package default_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/speedflow/speedflow/tests"
)

func TestNoFile(t *testing.T) {
	output, err := tests.Execute("-f", "no-file")
	assert.NoError(t, err)
	assert.Contains(t, output, `file "no-file" does not exists`)
}

func TestBadFile(t *testing.T) {
	output, err := tests.Execute("-f", ".bad.speedflow.yml")
	assert.NoError(t, err)
	assert.Contains(t, output, "unmarshal errors")
}
