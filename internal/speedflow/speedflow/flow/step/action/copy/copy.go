package copy

// Copy is a step that's helps to copy files from source to destination
// TODO: Doc dir to dir§ / file to file / etc.
type Copy struct {
	// Source is the source of files (Eg: ./bin/)
	Source string `yaml:"source"`

	// Destination is the destination of files (Eg: ./bin-saved/)
	Destination string `yaml:"destination"`
}
