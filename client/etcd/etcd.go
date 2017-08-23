package etcd

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/state"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/register"
    "github.com/rookie-xy/hubble/src/client"
	"github.com/rookie-xy/hubble/src/types"
)

const Namespace = "plugin.client.etcd"

type etcd struct {
    log.Log
}

func open(l log.Log, v types.Value) (client.Client, error) {
    return &etcd{
        Log: l,
    }, nil
}

func (r *etcd) Sender(e event.Event) int {
    return state.Ok
}

func (r *etcd) Close() int {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
