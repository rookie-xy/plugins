package console

import (
    "fmt"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/adapter"
    "github.com/rookie-xy/hubble/register"
)

type console struct {
    log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    return &console{
        Log: l,
    }, nil
}

func (c *console) Sender(e event.Event) error {
    fileEvent := adapter.ToFileEvent(e)
    state := fileEvent.GetFooter()
    body := adapter.ToFileEvent(e).GetBody()
    fmt.Printf("consoleeeeeeeeeeeeeeeeeeeeeeeeeeee: %d#%s\n ", state.Offset, string(body.GetContent()))
    return nil
}

func (c *console) Close() {
}

func init() {
    register.Client(Namespace, open)
}
