package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"runtime"
)

// Configuration is a type that stores runtime configuration information for the
// dependency checker service
type Configuration struct {
	ID       string `json:"id"`
	Os       string
	HostName string
}

var (
	// ErrMissingConfigFile happens when the configuration file cannot be found
	ErrMissingConfigFile = errors.New("configuration: the specified configuration file does not exist")

	// ErrParsingConfigFile happens when the configuration file cannot be parsed as json
	ErrParsingConfigFile = errors.New("configuration: cannot parse configuration file contents as json")
)

func readConfiguration(fileName string) ([]byte, error) {

	// read and return the contents of the local config file
	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		return contents, ErrMissingConfigFile
	}

	return contents, err
}

func parseConfiguration(c *Configuration, contents []byte) error {

	// parse the contents into a JSON string
	err := json.Unmarshal(contents, c)

	if err != nil {
		return ErrParsingConfigFile
	}

	// we need to set the Os and Hostname fields
	c.Os = runtime.GOOS
	c.HostName, _ = os.Hostname()
	return err
}

// LoadConfiguration loads the configuration from the file system
func (c *Configuration) loadConfiguration(fileName string) error {

	// read the configuration file
	contents, err := readConfiguration(fileName)

	if err != nil {
		return err
	}

	// parse the configuration and populate the configuration struct
	err = parseConfiguration(c, contents)

	if err != nil {
		return err
	}

	return nil
}
