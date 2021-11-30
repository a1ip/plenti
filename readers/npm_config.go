package readers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// NpmConfig is the NPM configuration file values from package.json.
type NpmConfig struct {
	Dependencies map[string]string `json:"dependencies"`
	Module       string            `json:"module"`
}

// GetNpmConfig reads the prod dependencies from package.json.
func GetNpmConfig(path string) NpmConfig {

	var npmConfig NpmConfig

	// Read NPM file for the project.
	configFile, _ := ioutil.ReadFile(path)
	err := json.Unmarshal(configFile, &npmConfig)
	if err != nil {
		fmt.Printf("Unable to read npm config file: %s\n", err)
	}

	return npmConfig
}
