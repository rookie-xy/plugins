package output

import (
    "fmt"
    "errors"
    "github.com/rookie-xy/worker/src/plugin"
    "github.com/rookie-xy/worker/src/register"
    "github.com/rookie-xy/worker/src/output"

  _ "github.com/rookie-xy/plugins/output/stdout"
)

const Namespace = "plugin.output"

type outputPlugin struct {
    name    string
    factory output.Factory
}

func Plugin(name string, f output.Factory) map[string][]interface{} {
     return plugin.Make(name, outputPlugin{name, f})
}

func init() {
    plugin.MustRegisterLoader(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(outputPlugin)
        if !ok {
            return errors.New("plugin does not match output codec plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Output(b.name, b.factory)

        return
    })
}
