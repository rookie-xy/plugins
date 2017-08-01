package codec

import (
    "fmt"
    "errors"

    "github.com/rookie-xy/worker/src/plugin"
    "github.com/rookie-xy/worker/src/register"
    "github.com/rookie-xy/worker/src/channel"

  _ "github.com/rookie-xy/plugins/channel/pipeline"
  _ "github.com/rookie-xy/plugins/channel/queue"
  _ "github.com/rookie-xy/plugins/channel/stream"
)

const Namespace = "plugin.channel"

type channelPlugin struct {
    name    string
    factory channel.Factory
}

func Plugin(name string, f channel.Factory) map[string][]interface{} {
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

        register.Channel(b.name, b.factory)

        return
    })
}
