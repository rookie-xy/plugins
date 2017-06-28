package plugins

import (
    "github.com/rookie-xy/worker/src/configure"
)

func Configure(name string) configure.ConfigureMethod {

    for _, plugin := range configure.Plugins {
        plugin.GetFile()
    }

    return nil
}