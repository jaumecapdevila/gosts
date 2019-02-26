package config

import (
	"io/ioutil"

	"github.com/jaumecapdevila/gosts/logger"
	yaml "gopkg.in/yaml.v2"
)

// YamlLoader loads the configuration from a yaml file
type YamlLoader struct {
	Logger logger.Logger
}

// Load configuration into params
func (l *YamlLoader) Load(name string) Parameters {
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		l.Logger.Fatal(logger.Context{}, "Unable to read configuarion file")
	}

	var params = Parameters{}

	if err = yaml.Unmarshal(contents, &params); err != nil {
		l.Logger.Fatal(logger.Context{}, "Error unmarshalling the configuration file")
	}

	return params
}

// NewLoader constructs a new Loader
func NewLoader(logger logger.Logger) Loader {
	return &YamlLoader{
		Logger: logger,
	}
}
