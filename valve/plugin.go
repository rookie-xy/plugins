package valve

import (
    "fmt"
    "errors"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/valve"

  _ "github.com/rookie-xy/plugins/valve/grok"
)

const Namespace = "plugin.valve"

type valvePlugin struct {
    name     string
    factory  valve.Factory
}

func Plugin(name string, f valve.Factory) map[string][]interface{} {
     return plugin.Make(name, valvePlugin{name, f})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(valvePlugin)
        if !ok {
            return errors.New("plugin does not match valve plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Valve(b.name, b.factory)

        return
    })
}
