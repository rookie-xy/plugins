package input

import (
    "fmt"
    "errors"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/register"

  _ "github.com/rookie-xy/plugins/input/log"
    "github.com/rookie-xy/hubble/input"
)

type inputPlugin struct {
    name    string
    factory input.Factory
}

func Plugin(name string, f input.Factory) map[string][]interface{} {
     return plugin.Make(name, inputPlugin{name, f})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(inputPlugin)
        if !ok {
            return errors.New("plugin does not match input plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Input(b.name, b.factory)
        return
    })
}
