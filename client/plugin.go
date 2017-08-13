package client

import (
    "fmt"
    "errors"

    "github.com/rookie-xy/hubble/src/plugin"
    "github.com/rookie-xy/hubble/src/register"
    "github.com/rookie-xy/hubble/src/output"
    "github.com/rookie-xy/hubble/src/client"

  _ "github.com/rookie-xy/plugins/client/stdout"
)

const Namespace = "plugin.client"

type clientPlugin struct {
    name    string
    factory client.Factory
}

func Plugin(name string, f output.Factory) map[string][]interface{} {
     return plugin.Make(name, clientPlugin{name, f})
}

func init() {
    plugin.MustRegisterLoader(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(clientPlugin)
        if !ok {
            return errors.New("plugin does not match output codec plugin type")
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
