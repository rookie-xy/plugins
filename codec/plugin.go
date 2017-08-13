package codec

import (
    "fmt"
    "errors"
    "github.com/rookie-xy/hubble/src/plugin"
    "github.com/rookie-xy/hubble/src/codec"
    "github.com/rookie-xy/hubble/src/register"

  _ "github.com/rookie-xy/plugins/codec/yaml"
  _ "github.com/rookie-xy/plugins/codec/json"
)

const Namespace = "plugin.codec"

type codecPlugin struct {
    name    string
    factory codec.Factory
}

func Plugin(name string, f codec.Factory) map[string][]interface{} {
     return plugin.Make(name, codecPlugin{name, f})
}

func init() {
    plugin.MustRegisterLoader(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(codecPlugin)
        if !ok {
            return errors.New("plugin does not match output codec plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Codec(b.name, b.factory)

        return
    })
}
