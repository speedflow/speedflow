package statement

// Statement represents a statement
type Statement struct {
	// FileExists checks if file exists
	FileExists string `yaml:"fileExists"`

	// FileNotExists checks if file not exists
	FileNotExists string `yaml:"fileNotExists"`

	// DirExists checks if dir exists
	DirExists string `yaml:"dirExists"`

	// DirNotExists checks if dir not exists
	DirNotExists string `yaml:"dirNotExists"`

	// IsTrue checks if statement is true
	IsTrue string `yaml:"if"`

	// IsFalse checks if statement is false
	IsFalse string `yaml:"ifNot"`
}
