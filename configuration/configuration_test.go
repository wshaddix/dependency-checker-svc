package configuration

import (
	"os"
	"testing"
)

const ballotX = "\u2717"
const checkMark = "\u2713"

func TestLoadingConfiguration(t *testing.T) {
	fileName := "../config.json"

	t.Log("Given the need to test loading configuration values from the filesystem")

	t.Logf("\tWhen checking \"%s\" for the configuration", fileName)

	c := new(Configuration)
	err := c.LoadConfiguration(fileName)

	if err != nil {
		t.Fatal("\t\t", ballotX, "Should be able to load the configuration: ", err)
	} else {
		t.Log("\t\t", checkMark, "Configuration file is loaded and parsed")
	}
}

func TestGettingConfiguration(t *testing.T) {
	fileName := "../config.json"

	t.Log("Given the need to test retrieving configuration values from the filesystem")

	t.Logf("\tWhen checking \"%s\" for the configuration", fileName)

	c := new(Configuration)
	err := c.LoadConfiguration(fileName)

	if err != nil {
		t.Fatal("\t\t", ballotX, "Should be able to load the configuration: ", err)
	} else {

		// verify that the ID field is unmarshaled correctly
		if c.ID != "9893cf53-4f33-4946-afdc-01c61c1835fd" {
			t.Fatalf("\t\t%s The id is wrong:%s", ballotX, c.ID)
		} else {
			t.Log("\t\t", checkMark, "The ID field is unmarshaled correctly: ", c.ID)
		}

		// verify that the Os field is set correctly
		if c.Os != "windows" {
			t.Fatalf("\t\t%s The Os is wrong:%s", ballotX, c.Os)
		} else {
			t.Log("\t\t", checkMark, "The Os field is unmarshaled correctly: ", c.Os)
		}

		// verify that the HostName field is set correctly
		hostname, _ := os.Hostname()

		if c.HostName != hostname {
			t.Fatalf("\t\t%s The Hostname is wrong:%s", ballotX, c.HostName)
		} else {
			t.Log("\t\t", checkMark, "The Hostname field is unmarshaled correctly: ", c.HostName)
		}
	}
}

func TestMalformedConfiguration(t *testing.T) {
	fileName := "../config_bad.json"

	t.Log("Given the need to test parse errors from the config file")

	t.Logf("\tWhen checking \"%s\" for the configuration", fileName)

	c := new(Configuration)
	err := c.LoadConfiguration(fileName)

	if err == nil {
		t.Fatal("\t\t", ballotX, "Should have returned an error")
	} else {
		switch err {
		case ErrParsingConfigFile:
			t.Log("\t\t", checkMark, "Should return a parsing error")
			return
		default:
			t.Fatal("\t\t", ballotX, "Should have returned a parsing error")
		}
	}
}

func TestMissingConfiguration(t *testing.T) {
	fileName := "../config_missing.json"

	t.Log("Given the need to test loading a missing configuration file")

	t.Logf("\tWhen checking \"%s\" for the configuration", fileName)

	c := new(Configuration)
	err := c.LoadConfiguration(fileName)

	if err == nil {
		t.Fatal("\t\t", ballotX, "Should have returned an error")
	} else {
		switch err {
		case ErrMissingConfigFile:
			t.Log("\t\t", checkMark, "Should return a missing config file error")
			return
		default:
			t.Fatal("\t\t", ballotX, "Should have returned a missing config file error")
		}
	}
}
