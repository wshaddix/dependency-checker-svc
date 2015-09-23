package configuration

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"runtime"
)

// ConfigSettings is a type that stores runtime configuration information for the
// dependency checker service
type ConfigSettings struct {
	ID       string `json:"id"`
	Os       string
	HostName string
	MD5Hash  string
	APIURL   string
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

func parseConfiguration(c *ConfigSettings, contents []byte) error {

	// parse the contents into a JSON string
	err := json.Unmarshal(contents, c)

	if err != nil {
		return ErrParsingConfigFile
	}

	// we need to set the Os, Hostname, MD5Hash and optionally the api url fields
	c.Os = runtime.GOOS
	c.HostName, _ = os.Hostname()
	c.MD5Hash = hashConfiguration(contents)

	if len(c.APIURL) == 0 {
		c.APIURL = "api.dependencychecker.com"
	}
	return err
}

func hashConfiguration(contents []byte) string {
	hasher := md5.New()
	hasher.Write(contents)
	return hex.EncodeToString(hasher.Sum(nil))
}

// LoadConfiguration loads the configuration from the file system
func (c *ConfigSettings) LoadConfiguration(fileName string) error {

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
