package codec

import (
    "fmt"
    "errors"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/source"
    "github.com/rookie-xy/hubble/register"

  _ "github.com/rookie-xy/plugins/source/log"
)

const Namespace = "plugin.source"

type sourcePlugin struct {
    name    string
    factory source.Factory
}

func Plugin(name string, f source.Factory) map[string][]interface{} {
     return plugin.Make(name, sourcePlugin{name, f})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(sourcePlugin)
        if !ok {
            return errors.New("plugin does not match source plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Source(b.name, b.factory)
        return
    })
}
