package codec

import (
    "fmt"
    "errors"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/register"

  _ "github.com/rookie-xy/plugins/codec/yaml"
  _ "github.com/rookie-xy/plugins/codec/json"
  _ "github.com/rookie-xy/plugins/codec/byte"
  _ "github.com/rookie-xy/plugins/codec/word"
  _ "github.com/rookie-xy/plugins/codec/rune"
  _ "github.com/rookie-xy/plugins/codec/line"
  _ "github.com/rookie-xy/plugins/codec/multiline"
)

type codecPlugin struct {
    name    string
    factory codec.Factory
}

func Plugin(name string, f codec.Factory) map[string][]interface{} {
     return plugin.Make(name, codecPlugin{name, f})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
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

        return nil
    })
}
