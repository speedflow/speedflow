package default_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/speedflow/speedflow/tests"
)

func TestNoFile(t *testing.T) {
	output, err := tests.Execute("-f", "no-file")
	assert.NoError(t, err)
	assert.Contains(t, output, "Unable to open the Speedflow file\n➡ Create an new Speedflow file? speedflow init\n")
}

func TestBadFile(t *testing.T) {
	output, err := tests.Execute("-f", ".bad.speedflow.yml")
	assert.NoError(t, err)
	assert.Contains(t, output, "Unable to read and parse the Speedflow file\n➡ Call the doctor? speedflow doctor\n")
}
