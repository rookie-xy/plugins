package output

import (
    "fmt"
    "errors"

    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"

  _ "github.com/rookie-xy/plugins/output/elasticsearch"
  _ "github.com/rookie-xy/plugins/output/sincedb"
)

const Namespace = "plugin.output"

type outputPlugin struct {
    name    string
    factory proxy.Client
}

func Plugin(name string, c proxy.Client) map[string][]interface{} {
     return plugin.Make(name, outputPlugin{name, c})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(outputPlugin)
        if !ok {
            return errors.New("plugin does not match output plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Client(b.name, b.factory)

        return
    })
}
