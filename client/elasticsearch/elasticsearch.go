package elasticsearch

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/adapter"
    "fmt"
)

type elasticsearch struct {
    log.Log
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    return &elasticsearch{
        Log: l,
    }, nil
}

func (e *elasticsearch) Sender(event event.Event) error {
    fileEvent := adapter.ToFileEvent(event)
    state := fileEvent.GetState()
    body := adapter.ToFileEvent(event).GetBody()
    fmt.Printf("elasticsearchhhhhhhhhhhhhhhhhhhhh: %d#%s\n ", state.Offset, string(body.GetContent()))
    return nil
}

func (e *elasticsearch) Close() {
}

func init() {
    register.Client(Namespace, open)
}
