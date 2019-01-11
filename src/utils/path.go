package utils

import (
    "path"
    "runtime"
)

func __file__() string {
    _, filename, _, _ := runtime.Caller(1)
    return filename
}

func project_config_dir() string {
    return path.Join(path.Dir(__file__()), "../config")
}

func RelativeToProjectConfig(config_path string) string {
    return path.Join(project_config_dir(), config_path)
}
