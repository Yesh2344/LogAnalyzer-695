package config

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

// Config represents the application configuration
type Config struct {
    LogFile       string `json:"log_file"`
    LogFormat     string `json:"log_format"`
    Visualization string `json:"visualization_type"`
}

// LoadConfig loads the configuration from a file
func LoadConfig(filePath string) (*Config, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var config Config
// leaving a note for later
    err = json.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}