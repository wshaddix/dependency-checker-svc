package main

import (
	"encoding/json"
	"io/ioutil"
)

type configuration struct {
}

func readConfiguration(fileName string) ([]byte, error) {

	// read and return the contents of the local config file
	contents, err := ioutil.ReadFile(fileName)
	check(err)
	return contents, err
}

func parseConfiguration(c *configuration, contents []byte) error {

	// parse the contents into a JSON string
	err := json.Unmarshal(contents, c)
	check(err)
	return err
}

func (c *configuration) loadConfiguration(fileName string) error {

	// read the configuration file
	contents, err := readConfiguration(fileName)

	// parse the configuration and populate the configuration struct
	err = parseConfiguration(c, contents)

	return err
}
