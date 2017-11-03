package elasticsearch

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
    "fmt"
)

const Namespace = "plugin.client.elasticsearch"

type elasticsearch struct {
    log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    fmt.Println("iniiiiiiiiiiiiiiiiiiiiiiiiiiielasticsearch")
    return &elasticsearch{
        Log: l,
    }, nil
}

func (r *elasticsearch) Sender(e event.Event) error {
    fmt.Println("elasticsearchhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhxxxxxxxxxxxxxxxxxxxxxxxxxxx")
    return nil
}

func (r *elasticsearch) Close() int {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
