package main

import (
	"encoding/json"
	"io/ioutil"

	"go.thomasd.se/ebooks/slack"
)

// Config contains the deserialized application configuration.
type Config struct {
	Slack slack.Config `json:"slack"`
}

// NewConfig deserializes a configuration file.
func NewConfig(filename string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}
