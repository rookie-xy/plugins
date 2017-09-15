package sincedb

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
	   "github.com/rookie-xy/hubble/types"
)

const Namespace = "plugin.client.sincedb"

type sincedb struct {
    log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    return &sincedb{
        Log: l,
    }, nil
}

func (r *sincedb) Sender(e event.Event) int {
    return state.Ok
}

func (r *sincedb) Close() int {
    return state.Ok
}

func (r *sincedb) Search() {
    return
}

func init() {
    register.Client(Namespace, open)
}
