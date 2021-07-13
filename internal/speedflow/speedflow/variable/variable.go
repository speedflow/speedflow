package variable

// Variable represents the variable struct
type Variable struct {
	// Name is the name of the variable
	Name string `yaml:"name"`

	// Value is the content of the variable
	Value string `yaml:"value"`
}
