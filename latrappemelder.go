package latrappemelder

import (
	"fmt"

	"github.com/DenBeke/la-trappe-melder/register"
)

// LaTrappeMelder 🍻 wraps all the latrappe-melder logic
type LaTrappeMelder struct {
	r      *register.Register
	config *Config
}

// New creates a new LaTrappeMelder with the given config
func New(config *Config) (*LaTrappeMelder, error) {

	r, err := register.New(config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't create register: %w", err)
	}

	return &LaTrappeMelder{r: r, config: config}, nil

}
