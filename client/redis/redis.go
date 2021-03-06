package redis

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/adapter"
    "fmt"
)

type redis struct {
    log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    return &redis{
        Log: l,
    }, nil
}

func (k *redis) Sender(event event.Event) error {
    fileEvent := adapter.ToFileEvent(event)
    state := fileEvent.GetFooter()
    body := adapter.ToFileEvent(event).GetBody()
    fmt.Printf("redisaaaaaaaaaaaaaaaaaaaaaaaaaaaa: %d#%s\n ", state.Offset, string(body.GetContent()))
    return nil
}

func (k *redis) Close() {
}

func init() {
    register.Client(Namespace, open)
}
