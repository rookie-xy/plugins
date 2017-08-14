package codec

import (
    "fmt"
    "errors"

    "github.com/rookie-xy/hubble/src/plugin"
    "github.com/rookie-xy/hubble/src/register"
    "github.com/rookie-xy/hubble/src/pipeline"

  _ "github.com/rookie-xy/plugins/pipeline/slot"
  _ "github.com/rookie-xy/plugins/pipeline/queue"
  _ "github.com/rookie-xy/plugins/pipeline/stream"
)

const Namespace = "plugin.pipeline"

type channelPlugin struct {
    name    string
    factory pipeline.Factory
}

func Plugin(name string, f pipeline.Factory) map[string][]interface{} {
     return plugin.Make(name, channelPlugin{name, f})
}

func init() {
    plugin.MustRegisterLoader(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(channelPlugin)
        if !ok {
            return errors.New("plugin does not match output codec plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Pipeline(b.name, b.factory)

        return
    })
}
