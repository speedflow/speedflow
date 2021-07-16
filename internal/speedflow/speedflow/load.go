package speedflow

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"
)

// Load loads the Speedflow file
func Load(path string) (err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("file \"%s\" does not exists", path),
		)
	}

	return errors.Wrap(
		yaml.Unmarshal(b, &SF),
		fmt.Sprintf("unable to convert \"%s\" file", path),
	)
}
