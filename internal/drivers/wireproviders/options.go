package wireproviders

import (
	"github.com/jessevdk/go-flags"
)

type Options struct {
	File string `short:"f" description:"The config file to load"`
}

// NewOptions parsers arguments provided to the application and returns an Options struct.
func NewOptions() (Options, error) {
	options := Options{}

	if _, err := flags.Parse(&options); err != nil {
		return Options{}, err
	}

	return options, nil
}
