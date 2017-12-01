package filter

import (
    "fmt"
    "errors"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/filter"

  _ "github.com/rookie-xy/plugins/filter/grok"
)

type filterPlugin struct {
    name     string
    factory  filter.Factory
}

func Plugin(name string, f filter.Factory) map[string][]interface{} {
     return plugin.Make(name, filterPlugin{name, f})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(filterPlugin)
        if !ok {
            return errors.New("plugin does not match filter plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Filter(b.name, b.factory)

        return nil
    })
}
