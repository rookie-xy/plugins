package channel

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/types"

)

type channel struct {
    log.Log
    channel chan event.Event
}

func open(l log.Log, v types.Value) (pipeline.Queue, error) {
    return &channel{
        Log: l,
        channel: make(chan event.Event, 1024),
    }, nil
}

// TODO 确定如何保证并发
func (r *channel) Clone() pipeline.Queue {
    return r
}

func (r *channel) Enqueue(e event.Event) error {
    r.channel <- e
    return nil
}

func (r *channel) Dequeue(size int) (event.Event, error) {
    event, open := <- r.channel
    if !open {
        return nil/*, state.Done*/, nil
    }

    return event, /*state.Ok*/ nil
}

func (r *channel) Requeue(e event.Event) error {
    return r.Enqueue(e)
}

func (r *channel) Close() int {
    return state.Ok
}

func (r *channel) On() {
    return
}

func (r *channel) Off() {
    return
}

func init() {
    register.Pipeline(Namespace, open)
}
