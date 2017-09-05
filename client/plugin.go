package client

import (
    "fmt"
    "errors"

    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"

  _ "github.com/rookie-xy/plugins/client/stdout"
  _ "github.com/rookie-xy/plugins/client/elasticsearch"
  _ "github.com/rookie-xy/plugins/client/kafka"
  _ "github.com/rookie-xy/plugins/client/logstash"
)

const Namespace = "plugin.client"

type clientPlugin struct {
    name    string
    factory proxy.Client
}

func Plugin(name string, f proxy.Client) map[string][]interface{} {
     return plugin.Make(name, clientPlugin{name, f})
}

func init() {
    plugin.Register(Namespace, func(ifc interface{}) (err error) {
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
