package speedflow

import (
	"github.com/speedflow/speedflow/internal/speedflow/speedflow/flow"
	"github.com/speedflow/speedflow/internal/speedflow/speedflow/variable"
)

var (
	SF = &Speedflow{}
)

// Speedflow represents the YAML file
type Speedflow struct {
	// Version is the .speedflow.yaml schema version
	Version string `yaml:"version" default:"1"`

	// Imports is the others imported flows
	Imports map[string]string `yaml:"imports"`

	// Variables is the default variables
	Variables []variable.Variable `yaml:"vars"`

	// Flows is the flows used
	Flows map[string]flow.Flow `yaml:"flows"`
}
