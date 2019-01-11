package utils

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "reflect"
)

const (
    DefaultConfigFile = "config.json"
)

func ConfigFile(path string) string {
    return fmt.Sprintf("%s/%s", path, DefaultConfigFile)
}

func GetConfigFatal(config interface{}, path *string) {
    var file string
    if path == nil {
        file = DefaultConfigFile
    } else {
        file = *path
    }
    getConfigFatal(config, file)
}

func getConfigFatal(config interface{}, path string) {
    err := getConfig(config, path)
    if err != nil {
        message := fmt.Sprintf("Could not initialize %s Config", reflect.TypeOf(config))
        log.Fatalln(message, err)
    }
}

//Usage: GetConfig(&config struct, path_relative_to_project_root_config_dir) returns error
func getConfig(config interface{}, path string) error {
    bytes, err := ioutil.ReadFile(RelativeToProjectConfig(path))
    if err != nil {
        return err
    }
    err = json.Unmarshal(bytes, config)
    return err
}
