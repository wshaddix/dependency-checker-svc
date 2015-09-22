package main

import "testing"

const ballotX = "\u2717"
const checkMark = "\u2713"

func TestLoadingConfiguration(t *testing.T) {
	fileName := "config.json"

	t.Log("Given the need to test loading configuration values from the filesystem")

	t.Logf("\tWhen checking \"%s\" for the configuration", fileName)

	c := new(configuration)
	err := c.loadConfiguration(fileName)

	if err != nil {
		t.Fatal("\t\t", ballotX, "Should be able to load the configuration: ", err)
	} else {
		t.Log("\t\t", checkMark, "Should be able to load the configuration")
	}
}

func TestParsingConfiguration(t *testing.T) {
	fileName := "config_bad.json"

	t.Log("Given the need to test parse errors from the config file")

	t.Logf("\tWhen checking \"%s\" for the configuration", fileName)

	c := new(configuration)
	err := c.loadConfiguration(fileName)

	if err != nil {
		t.Fatal("\t\t", ballotX, "Should be able to load the configuration: ", err)
	} else {
		t.Log("\t\t", checkMark, "Should be able to load the configuration")
	}
}
