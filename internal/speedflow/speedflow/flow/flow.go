package flow

import "github.com/speedflow/speedflow/internal/speedflow/speedflow/flow/step"

const (
	// DefaultFlowName represents the default flow name
	DefaultFlowName = "default"
)

// Flow is a flow structure
type Flow struct {
	// Name is the flow name (Eg: demo)
	Name string `yaml:"name"`

	// Description is the flow description (Eg: This is an amazing demo)
	Description string `yaml:"description"`

	// Steps are the steps to execute for the flow
	Steps []step.Step `yaml:"steps"`

	// Internal is a boolean to declare the flow internal
	Internal bool `yaml:"internal"`

	// Image is used to run steps in a container produced from this image
	Image string `yaml:"image"`
}
