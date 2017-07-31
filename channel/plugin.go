package codec

import (
    "fmt"
    "errors"

    "github.com/rookie-xy/worker/src/plugin"
    "github.com/rookie-xy/worker/src/register"
    "github.com/rookie-xy/worker/src/channel"

  _ "github.com/rookie-xy/plugins/codec/yaml"
  _ "github.com/rookie-xy/plugins/codec/json"
)

type channelPlugin struct {
    name   string
    method channel.Channel
}

const Name = "channel"

func Plugin(name string, m channel.Channel) map[string][]interface{} {
     return plugin.Make(name, channelPlugin{name, m})
}

func init() {
    plugin.MustRegisterLoader(Name, func(ifc interface{}) (err error) {
        b, ok := ifc.(channelPlugin)
        if !ok {
            return errors.New("plugin does not match output codec plugin type")
        }

        defer func() {
            if msg := recover(); msg != nil {
                err = fmt.Errorf("%s", msg)
	           }
        }()

        register.Channel(b.name, b.method)

        return
    })
}
