package elasticsearch

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/state"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/register"
    "github.com/rookie-xy/hubble/src/proxy"
	   "github.com/rookie-xy/hubble/src/types"
)

const Namespace = "plugin.client.elasticsearch"

type elasticsearch struct {
    log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    return &elasticsearch{
        Log: l,
    }, nil
}

func (r *elasticsearch) Sender(e event.Event) int {
    return state.Ok
}

func (r *elasticsearch) Close() int {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
