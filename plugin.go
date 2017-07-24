package plugins

import (
    "fmt"
    "github.com/rookie-xy/worker/src/state"
    "github.com/rookie-xy/worker/src/plugin"

    "github.com/rookie-xy/worker/src/plugin/codec"
)

func Codec(name string) plugin.Codec {
    fmt.Println("uuuuuuuuuuuuuuu", name)
    for _, plugin := range codec.Plugins {
        if plugin.Type(name) == state.Ok {
            fmt.Println("hhhhhhhhhhhhhhhhh")
            return plugin.Clone()
        }
    }

    fmt.Println("nnnnnnnnnnnnnnnnnnnnnnnnnnn")

    return nil
}
