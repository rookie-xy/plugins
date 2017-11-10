package file

import (
    "fmt"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/adapter"
    "github.com/rookie-xy/hubble/register"
)

type file struct {
    log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    return &file{
        Log: l,
    }, nil
}

func (f *file) Sender(e event.Event) error {
    fileEvent := adapter.ToFileEvent(e)
    state := fileEvent.GetState()
    body := adapter.ToFileEvent(e).GetBody()
    fmt.Printf("fileeeeeeeeeeeeeeeeeeeeeeeeeeeeeee: %d#%s\n ", state.Offset, string(body.GetContent()))
    return nil
}

func (f *file) Close() {
}

func init() {
    register.Client(Namespace, open)
}
