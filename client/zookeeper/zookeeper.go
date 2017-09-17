package zookeeper

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
)

const Namespace = "plugin.client.zookeeper"

type zookeeper struct {
    log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    return &zookeeper{
        Log: l,
    }, nil
}

func (r *zookeeper) Post(e event.Event) int {
    return state.Ok
}

func (r *zookeeper) Delete(e event.Event) int {
    return state.Ok
}

func (r *zookeeper) Put(e event.Event) int {
    return state.Ok
}

func (r *zookeeper) Get(e event.Event) types.Object {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
