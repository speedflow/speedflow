package print

type Format string

type Print struct {
	// Message is the message to print (Eg: Hello!)
	Message string `yaml:"message"`

	// Format is the format of message (Eg: info, success, warn, error)
	Format Format `yaml:"format"`
}
