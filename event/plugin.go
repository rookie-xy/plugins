package event

import (
    "fmt"
    "errors"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/register"

  _ "github.com/rookie-xy/plugins/event/yaml"
  _ "github.com/rookie-xy/plugins/event/json"
  _ "github.com/rookie-xy/plugins/event/line"
  _ "github.com/rookie-xy/plugins/event/multiline"
)

const Namespace = "plugin.event"

type eventPlugin struct {
    name    string
    factory event.Factory
}

func Plugin(name string, f event.Factory) map[string][]interface{} {
     return plugin.Make(name, eventPlugin{name, f})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(eventPlugin)
        if !ok {
            return errors.New("plugin does not match event plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Event(b.name, b.factory)

        return
    })
}
