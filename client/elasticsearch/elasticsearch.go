package elasticsearch

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
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
/*
func (r *elasticsearch) Post(e event.Event) int {
    return state.Ok
}

func (r *elasticsearch) Delete(e event.Event) int {
    return state.Ok
}

func (r *elasticsearch) Put(e event.Event) int {
    return state.Ok
}

func (r *elasticsearch) Get(e event.Event) types.Object {
    return state.Ok
}
*/

func init() {
    register.Client(Namespace, open)
}
