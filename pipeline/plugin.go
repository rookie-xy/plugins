package pipeline

import (
    "fmt"
    "errors"

    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/pipeline"

  _ "github.com/rookie-xy/plugins/pipeline/channel"
  _ "github.com/rookie-xy/plugins/pipeline/queue"
  _ "github.com/rookie-xy/plugins/pipeline/stream"
)

type channelPlugin struct {
    name    string
    factory pipeline.Factory
}

func Plugin(name string, f pipeline.Factory) map[string][]interface{} {
     return plugin.Make(name, channelPlugin{name, f})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
        b, ok := ifc.(channelPlugin)
        if !ok {
            return errors.New("plugin does not match pipeline plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	        }
        }()

        register.Pipeline(b.name, b.factory)
        return nil
    })
}
