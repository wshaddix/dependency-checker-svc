package main

import (
	"fmt"

	"github.com/wshaddix/dependency-checker-svc/configuration"
)

const configFileName = "config.json"

func main() {
	// we need to do the following tasks at startup
	// 1. read the local config.json file to pull in our settings
	// 2. hash the contents of the config.json file for comparison against the api server
	// 3. check in with the dependency checker api to send a heartbeat
	// 4. compare the dependencies hash returned from the heartbeat with our
	//    local hash. If they do not match we need to update our configuration
	// 5. (optional) Update our local configuration from the api server
	// 6. parse the dependencies from config.json and start monitoring every
	//    minute
	// 7. setup a timer to send a heartbeat every minute

	// 1. read the local config.json file to pull in our settings
	config := new(configuration.ConfigSettings)
	err := config.LoadConfiguration(configFileName)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(config.ID, err)
}
