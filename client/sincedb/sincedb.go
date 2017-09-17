package sincedb

import (
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "fmt"
)

const Namespace = "plugin.client.sincedb"

type sinceDB struct {
    log log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    fmt.Println("sincedbbbbbbbbbbbbbb")
    return &sinceDB{
        log: l,
    }, nil
}

func (r *sinceDB) Sender(e event.Event) int {
    return state.Ok
}

func (r *sinceDB) Add() int {
    return state.Ok
}

func (r *sinceDB) Find() types.Object {
    return state.Ok
}

func (r *sinceDB) Close() int {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
