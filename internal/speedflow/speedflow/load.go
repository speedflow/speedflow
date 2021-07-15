package speedflow

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Load loads the Speedflow file
func Load(path string) (err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, &SF)
}
