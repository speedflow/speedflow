package step

import (
	actionCopy "github.com/speedflow/speedflow/internal/speedflow/speedflow/flow/step/action/copy"
	"github.com/speedflow/speedflow/internal/speedflow/speedflow/flow/step/statement"
)

// Actions represents the available actions
type Actions struct {
	// ActionCopy is an action to copy files from source to destination
	ActionCopy actionCopy.Copy `yaml:"copy"`

	// ActionPrintError is an action to print an error message to stdErr
	ActionPrintError string `yaml:"printError"`

	// ActionPrintWarning is an action to print a warning message to stdErr
	ActionPrintWarning string `yaml:"printWarning"`

	// ActionPrintSuccess is an action to print a success message to stdOut
	ActionPrintSuccess string `yaml:"printSuccess"`

	// ActionPrintInfo is an action to print an info message to stdOut
	ActionPrintInfo string `yaml:"printInfo"`

	// ActionRun is an action to execute command in a shell
	// TODO: Convert to struct?
	ActionRun string `yaml:"run"`

	// ActionFlow is an action to execute an other flow
	ActionFlow string `yaml:"flow"`
}

// Step represents a flow step
type Step struct {
	// ID is the step identifier (Eg: copy)
	// This field is used to refer to the step (Eg: .steps.copy.out)
	ID string `yaml:"id"`

	// Name is the step name (Eg: My Default Step)
	// This field is used to debug
	Name string `yaml:"name"`

	// Description is the step description (Eg: This step is useful to ...)
	// This field is used to debug
	Description string `yaml:"description"`

	// Actions is the possible action in this step
	Actions

	// Statements are the statements for execute this step
	Statements []statement.Statement
}
